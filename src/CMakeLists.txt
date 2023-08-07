# The Flutter tooling requires that developers have CMake 3.10 or later
# installed. You should not increase this version, as doing so will cause
# the plugin to fail to compile for some customers of the plugin.
cmake_minimum_required(VERSION 3.10)

project(catweb_parser_library VERSION 0.0.1 LANGUAGES C)

add_library(catweb_parser SHARED
  "catweb_parser.c"
)

set_target_properties(catweb_parser PROPERTIES
  PUBLIC_HEADER catweb_parser.h
  OUTPUT_NAME "catweb_parser"
)

target_compile_definitions(catweb_parser PUBLIC DART_SHARED_LIB)
