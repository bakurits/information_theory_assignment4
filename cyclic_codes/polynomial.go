package cyclic_codes

import (
	"errors"
)

type Polynomial struct {
	Base         int
	Degree       int
	Coefficients []int
}

func InitZeroesPolynomial(n int, p int) Polynomial {
	var pol Polynomial
	pol.Base = p
	pol.Degree = n
	pol.Coefficients = make([]int, n+1)
	return pol
}

func (pol *Polynomial) CoefficientDegree(i int) int {
	return pol.Degree - i
}

func (pol *Polynomial) AddDegree(deg int, coefficient int) {
	if deg > pol.Degree {
		return
	}
	pol.Coefficients[pol.Degree-deg] = (pol.Coefficients[pol.Degree-deg] + coefficient) % pol.Base
}

func (pol *Polynomial) IsZero() bool {
	if pol.Degree > 0 {
		return false
	}
	return pol.Coefficients[0] == 0
}

func (pol *Polynomial) Trim() Polynomial {
	st := len(pol.Coefficients)
	for i := 0; i < len(pol.Coefficients); i++ {
		if pol.Coefficients[i] != 0 {
			st = i
			break
		}
	}
	if st == len(pol.Coefficients) {
		return InitZeroesPolynomial(0, pol.Base)
	}

	return Polynomial{
		Base:         pol.Base,
		Degree:       pol.Degree - st,
		Coefficients: pol.Coefficients[st:],
	}

}

func (pol *Polynomial) Divide(polDiv *Polynomial) (Polynomial, Polynomial, error) {
	a := pol.Trim()
	b := polDiv.Trim()
	if pol.Base != b.Base {
		return Polynomial{}, Polynomial{}, errors.New("bases don't match")
	}
	p := pol.Base
	if pol.Degree < b.Degree {
		return InitZeroesPolynomial(0, p), a, nil
	}

	dv := InitZeroesPolynomial(a.Degree, p)
	rm := InitZeroesPolynomial(a.Degree, p)

	for i := 0; i < len(a.Coefficients); i++ {
		for i < len(a.Coefficients) {
			if a.Coefficients[i] != 0 {
				break
			}
			i++
		}
		if i == len(a.Coefficients) {
			break
		}
		if a.CoefficientDegree(i) < b.Degree {
			for j := i; j < len(a.Coefficients); j++ {
				rm.Coefficients[j] = a.Coefficients[j]
			}
			break
		}

		mul := int((int64(a.Coefficients[i]) * int64(findInverse(b.Coefficients[0], p))) % int64(p))
		for j := 0; j < len(b.Coefficients); j++ {
			a.Coefficients[i+j] -= int(int64(mul) * int64(b.Coefficients[j]) % int64(p))
			a.Coefficients[i+j] = (a.Coefficients[i+j] + p) % p
		}
		dv.Coefficients[len(dv.Coefficients)-(a.CoefficientDegree(i)-b.Degree)-1] = mul
	}
	return dv.Trim(), rm.Trim(), nil
}

func (pol *Polynomial) Multiply(polMul *Polynomial) (Polynomial, error) {

	if pol.Base != polMul.Base {
		return Polynomial{}, errors.New("bases don't match")
	}
	p := pol.Base
	a := pol.Trim()
	b := polMul.Trim()
	res := Polynomial{
		Base:         p,
		Degree:       a.Degree + b.Degree,
		Coefficients: nil,
	}
	res.Coefficients = make([]int, res.Degree+1)
	for i := 0; i < len(a.Coefficients); i++ {
		for j := 0; j < len(b.Coefficients); j++ {
			curVal := int((int64(a.Coefficients[i]) * int64(b.Coefficients[j])) % int64(p))
			curDeg := a.CoefficientDegree(i) + b.CoefficientDegree(j)
			res.AddDegree(curDeg, curVal)
		}
	}
	return res, nil
}
