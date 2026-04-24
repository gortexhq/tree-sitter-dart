package tree_sitter_dart

// #cgo CFLAGS: -I${SRCDIR}/../../src -std=c11 -fPIC
// #include "../../src/parser.c"
// #include "../../src/scanner.c"
import "C"

import "unsafe"

// Language returns the tree-sitter Language for Dart.
func Language() unsafe.Pointer {
	return unsafe.Pointer(C.tree_sitter_dart())
}
