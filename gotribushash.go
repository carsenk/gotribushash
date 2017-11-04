// Package gotribushash wraps the C Tribus hash implementation to make it accessible in Go. Tribus isn't
// a single hash function but a series of 3 successive functions: echo, jh, keccak. Its only real usefulness is for
// alternate cryptocurrencies like Denarius.
package gotribushash

// #include "denarius.h"
import "C"

// Hash the provided data, returning a slice of the [32]byte containing the resulting hash.
func Sum(data []byte) []byte {
  var cresstr [32]C.char
  C.denarius_hash(C.CString(string(data)), &cresstr[0])
  return []byte(C.GoStringN(&cresstr[0], 32))
}
