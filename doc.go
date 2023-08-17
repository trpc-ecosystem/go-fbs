// Package fbs provides parsing facilities for flatbuffers .fbs file.
// It can read .fbs file to generate flatbuffers descriptor.
//
// `lexer` is used to generate token stream for .fbs file.
// `parser` is used to generate AST and descriptor from token stream.
// `linker` is used to resolve type references in descriptors.
package fbs
