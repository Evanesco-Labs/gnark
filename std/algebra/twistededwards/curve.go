/*
Copyright © 2020 ConsenSys

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package twistededwards

import (
	"errors"
	"math/big"

	"github.com/consensys/gnark-crypto/ecc"
	frbls12377 "github.com/consensys/gnark-crypto/ecc/bls12-377/fr"
	edbls12377 "github.com/consensys/gnark-crypto/ecc/bls12-377/twistededwards"
	frbls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381/fr"
	edbls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381/twistededwards"
	frbls24315 "github.com/consensys/gnark-crypto/ecc/bls24-315/fr"
	edbls24315 "github.com/consensys/gnark-crypto/ecc/bls24-315/twistededwards"
	frbn254 "github.com/consensys/gnark-crypto/ecc/bn254/fr"
	edbn254 "github.com/consensys/gnark-crypto/ecc/bn254/twistededwards"
	frbw6633 "github.com/consensys/gnark-crypto/ecc/bw6-633/fr"
	edbw6633 "github.com/consensys/gnark-crypto/ecc/bw6-633/twistededwards"
	frbw6761 "github.com/consensys/gnark-crypto/ecc/bw6-761/fr"
	edbw6761 "github.com/consensys/gnark-crypto/ecc/bw6-761/twistededwards"
	"github.com/consensys/gnark/frontend"
)

// EdCurve stores the info on the chosen edwards curve
type EdCurve struct {
	A, D, Cofactor, Order, BaseX, BaseY, Modulus big.Int
	ID                                           ecc.ID
}

var newTwistedEdwards map[ecc.ID]func() EdCurve

func init() {
	newTwistedEdwards = make(map[ecc.ID]func() EdCurve)
	newTwistedEdwards[ecc.BLS12_381] = newEdBLS381
	newTwistedEdwards[ecc.BN254] = newEdBN254
	newTwistedEdwards[ecc.BLS12_377] = newEdBLS377
	newTwistedEdwards[ecc.BW6_761] = newEdBW761
	newTwistedEdwards[ecc.BLS24_315] = newEdBLS315
	newTwistedEdwards[ecc.BW6_633] = newEdBW633
}

// NewEdCurve returns an Edwards curve parameters
func NewEdCurve(id ecc.ID) (EdCurve, error) {
	if constructor, ok := newTwistedEdwards[id]; ok {
		return constructor(), nil
	}
	return EdCurve{}, errors.New("unknown curve id")
}

// -------------------------------------------------------------------------------------------------
// constructors

func newEdBN254() EdCurve {

	edcurve := edbn254.GetEdwardsCurve()
	var cofactorReg big.Int
	edcurve.Cofactor.ToBigInt(&cofactorReg)

	res := EdCurve{
		A:        frontend.FromInterface(edcurve.A),
		D:        frontend.FromInterface(edcurve.D),
		Cofactor: frontend.FromInterface(cofactorReg),
		Order:    frontend.FromInterface(edcurve.Order),
		BaseX:    frontend.FromInterface(edcurve.Base.X),
		BaseY:    frontend.FromInterface(edcurve.Base.Y),
		ID:       ecc.BN254,
	}
	res.Modulus.Set(frbn254.Modulus())

	return res

}

func newEdBLS381() EdCurve {

	edcurve := edbls12381.GetEdwardsCurve()
	var cofactorReg big.Int
	edcurve.Cofactor.ToBigInt(&cofactorReg)

	res := EdCurve{
		A:        frontend.FromInterface(edcurve.A),
		D:        frontend.FromInterface(edcurve.D),
		Cofactor: frontend.FromInterface(cofactorReg),
		Order:    frontend.FromInterface(edcurve.Order),
		BaseX:    frontend.FromInterface(edcurve.Base.X),
		BaseY:    frontend.FromInterface(edcurve.Base.Y),
		ID:       ecc.BLS12_381,
	}
	res.Modulus.Set(frbls12381.Modulus())

	return res
}

func newEdBLS377() EdCurve {

	edcurve := edbls12377.GetEdwardsCurve()
	var cofactorReg big.Int
	edcurve.Cofactor.ToBigInt(&cofactorReg)

	res := EdCurve{
		A:        frontend.FromInterface(edcurve.A),
		D:        frontend.FromInterface(edcurve.D),
		Cofactor: frontend.FromInterface(cofactorReg),
		Order:    frontend.FromInterface(edcurve.Order),
		BaseX:    frontend.FromInterface(edcurve.Base.X),
		BaseY:    frontend.FromInterface(edcurve.Base.Y),
		ID:       ecc.BLS12_377,
	}
	res.Modulus.Set(frbls12377.Modulus())

	return res
}

func newEdBW761() EdCurve {

	edcurve := edbw6761.GetEdwardsCurve()
	var cofactorReg big.Int
	edcurve.Cofactor.ToBigInt(&cofactorReg)

	res := EdCurve{
		A:        frontend.FromInterface(edcurve.A),
		D:        frontend.FromInterface(edcurve.D),
		Cofactor: frontend.FromInterface(cofactorReg),
		Order:    frontend.FromInterface(edcurve.Order),
		BaseX:    frontend.FromInterface(edcurve.Base.X),
		BaseY:    frontend.FromInterface(edcurve.Base.Y),
		ID:       ecc.BW6_761,
	}
	res.Modulus.Set(frbw6761.Modulus())

	return res
}

func newEdBLS315() EdCurve {

	edcurve := edbls24315.GetEdwardsCurve()
	var cofactorReg big.Int
	edcurve.Cofactor.ToBigInt(&cofactorReg)

	res := EdCurve{
		A:        frontend.FromInterface(edcurve.A),
		D:        frontend.FromInterface(edcurve.D),
		Cofactor: frontend.FromInterface(cofactorReg),
		Order:    frontend.FromInterface(edcurve.Order),
		BaseX:    frontend.FromInterface(edcurve.Base.X),
		BaseY:    frontend.FromInterface(edcurve.Base.Y),
		ID:       ecc.BLS24_315,
	}
	res.Modulus.Set(frbls24315.Modulus())

	return res
}

func newEdBW633() EdCurve {

	edcurve := edbw6633.GetEdwardsCurve()
	var cofactorReg big.Int
	edcurve.Cofactor.ToBigInt(&cofactorReg)

	res := EdCurve{
		A:        frontend.FromInterface(edcurve.A),
		D:        frontend.FromInterface(edcurve.D),
		Cofactor: frontend.FromInterface(cofactorReg),
		Order:    frontend.FromInterface(edcurve.Order),
		BaseX:    frontend.FromInterface(edcurve.Base.X),
		BaseY:    frontend.FromInterface(edcurve.Base.Y),
		ID:       ecc.BW6_633,
	}
	res.Modulus.Set(frbw6633.Modulus())

	return res
}
