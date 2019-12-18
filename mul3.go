////////////////////////////////////////////////////////////////////////////////
// Copyright © 2020 Privategrity Corporation                                   /
//                                                                             /
// All rights reserved.                                                        /
////////////////////////////////////////////////////////////////////////////////
package cryptops

import "gitlab.com/elixxir/crypto/cyclic"

// It would be pretty easy to make this take a variable number of parameters
type Mul3Prototype func(g *cyclic.Group, x, y *cyclic.Int,
	out *cyclic.Int) *cyclic.Int

// Multiplies 3 numbers in a cyclic group within the cryptops interface.
// Sets `out = x*y*out mod p` and returns out.
var Mul3 Mul3Prototype = func(g *cyclic.Group, x, y, out *cyclic.Int) *cyclic.Int {
	g.Mul(out, x, out)
	g.Mul(out, y, out)
	return out
}

// Returns the function name for debugging.
func (Mul3Prototype) GetName() string {
	return "Mul3"
}

// Returns the input size; used in safety checks.
func (Mul3Prototype) GetInputSize() uint32 {
	return 1
}
