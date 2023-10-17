//
//
// Tencent is pleased to support the open source community by making tRPC available.
//
// Copyright (C) 2023 THL A29 Limited, a Tencent company.
// All rights reserved.
//
// If you have downloaded a copy of the tRPC source code from Tencent,
// please note that tRPC source code is licensed under the Apache 2.0 License,
// A copy of the Apache 2.0 License is included in this file.
//
//

// Package fbs provides parsing facilities for flatbuffers .fbs file.
// It can read .fbs file to generate flatbuffers descriptor.
//
// `lexer` is used to generate token stream for .fbs file.
// `parser` is used to generate AST and descriptor from token stream.
// `linker` is used to resolve type references in descriptors.
package fbs
