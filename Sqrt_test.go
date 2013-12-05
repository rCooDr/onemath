package onemath

import "testing"

func TestSqrt(t *testing.T) {
     const in, out = 4, 2
     if x := Sqrt(in); x != out {
         t.Errorf("Sqrt(%v) = %v, want %v\n", in, x, out)
     }
}
/*

(M-1 go test)

-|
PASS
ok  	github.com/rCooDr/onemath	0.003s

*/