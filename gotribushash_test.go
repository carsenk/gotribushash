package gotribushash

import (
  "bytes"
  "testing"

  "github.com/carsenk/gotribushash"
)


func TestTribusHash(t *testing.T) {
  tribush  := gotribushash.Sum(
            []byte("00000000000000000000000000000000000000000000000000000000000000000000000000000000"))
  exp   := [32]byte{84, 183, 251, 186, 152, 73, 8, 109, 107, 58, 125, 219, 129, 97, 234, 243, 241,
                    195, 39, 242, 189, 238, 17, 253, 216, 217, 31, 243, 5, 113, 8, 170}
  if bytes.Compare(tribush, exp[:]) != 0 {
    t.Error("Hashing produced", tribush, "expected", exp)
  }
}
