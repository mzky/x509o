// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build js && wasm
// +build js,wasm

package x509o

// Possible certificate files; stop after finding one.
var certFiles = []string{}

// Possible directories with certificate files; stop after successfully
// reading at least one file from a directory.
var certDirectories = []string{}
