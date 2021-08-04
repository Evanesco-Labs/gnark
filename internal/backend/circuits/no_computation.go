package circuits

import (
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/frontend"
)

// test circuit with no computational constraints
type noComputationCircuit struct {
	A frontend.Variable `gnark:",public"`
	B frontend.Variable
}

func (c *noComputationCircuit) Define(curveID ecc.ID, cs *frontend.ConstraintSystem) error {
	cs.AssertIsEqual(c.A, c.B)
	return nil
}

func init() {

	var circuit, good, bad, public noComputationCircuit

	good.A.Assign(42)
	good.B.Assign(42)

	bad.A.Assign(42)
	bad.B.Assign(8000)

	public.A.Assign(42)

	addEntry("noComputation", &circuit, &good, &bad, &public)
}
