import 'dart:async';
import 'dart:ffi';
import 'dart:io';
import 'dart:isolate';
import 'package:ffi/ffi.dart';

import 'catweb_parser_bindings_generated.dart';

Future<String> parseHtmlAsync(
  String content,
  String parserType,
  String parser,
) async {
  final SendPort helperIsolateSendPort = await _helperIsolateSendPort;
  final int requestId = _nextSumRequestId++;
  final _ParseHtmlRequest request =
      _ParseHtmlRequest(requestId, content, parserType, parser);
  final completer = Completer<String>();
  helperIsolateSendPort.send(request);
  _parseRequests[requestId] = completer;
  return completer.future;
}

const String _libName = 'catweb_parser';

final DynamicLibrary _dylib = () {
  if (Platform.isMacOS || Platform.isIOS) {
    return DynamicLibrary.open('$_libName.framework/$_libName');
  }
  if (Platform.isAndroid || Platform.isLinux) {
    return DynamicLibrary.open('lib$_libName.so');
  }
  if (Platform.isWindows) {
    try {
      return DynamicLibrary.open('$_libName.dll');
    } catch (_) {
      return DynamicLibrary.open('../windows/$_libName.dll');
    }
  }

  throw UnsupportedError('Unknown platform: ${Platform.operatingSystem}');
}();

final CatwebParserBindings _bindings = CatwebParserBindings(_dylib);

class _ParseHtmlRequest {
  final int id;
  final String context;
  final String parserType;
  final String parser;

  const _ParseHtmlRequest(this.id, this.context, this.parserType, this.parser);
}

class _ParseHtmlResponse {
  final int id;
  final String result;

  const _ParseHtmlResponse(this.id, this.result);
}

int _nextSumRequestId = 0;

final Map<int, Completer<String>> _parseRequests = <int, Completer<String>>{};

Future<SendPort> _helperIsolateSendPort = () async {
  final Completer<SendPort> completer = Completer<SendPort>();

  final ReceivePort receivePort = ReceivePort()
    ..listen((dynamic data) {
      if (data is SendPort) {
        completer.complete(data);
        return;
      }
      if (data is _ParseHtmlResponse) {
        final Completer<String> completer = _parseRequests[data.id]!;
        _parseRequests.remove(data.id);
        completer.complete(data.result);
        return;
      }
      throw UnsupportedError('Unsupported message type: ${data.runtimeType}');
    });

  await Isolate.spawn((SendPort sendPort) async {
    final ReceivePort helperReceivePort = ReceivePort()
      ..listen((dynamic data) {
        if (data is _ParseHtmlRequest) {
          final context = data.context.toNativeUtf8().cast<Char>();
          final parserType = data.parserType.toNativeUtf8().cast<Char>();
          final parser = data.parser.toNativeUtf8().cast<Char>();
          final result = _bindings.ParseHtml(context, parserType, parser);
          final response =
              _ParseHtmlResponse(data.id, result.cast<Utf8>().toDartString());
          sendPort.send(response);
          return;
        }
        throw UnsupportedError('Unsupported message type: ${data.runtimeType}');
      });

    sendPort.send(helperReceivePort.sendPort);
  }, receivePort.sendPort);
  return completer.future;
}();
