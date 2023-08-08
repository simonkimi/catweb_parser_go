import 'dart:io';

import 'package:catweb_parser/catweb_parser.dart' as catweb_parser;

void main() async {
  var html = File(r"G:\personal\catweb\test\list.htm").readAsStringSync();
  var parser = File(r"G:\personal\catweb\test\ffi_request.json").readAsStringSync();
  var a = await catweb_parser.parseHtmlAsync(html, "ListParser", parser);
  print(a);
}
