package cyclic_codes

import "errors"

func ParityCheck(pol Polynomial) (Polynomial, error) {

	a := InitZeroesPolynomial(pol.Degree+1, pol.Base)
	a.Coefficients[len(a.Coefficients)-1] = pol.Base - 1
	a.Coefficients[0] = 1

	dv, rem, err := a.Divide(&pol)
	if err != nil {
		return Polynomial{}, err
	}

	if !rem.IsZero() {
		return Polynomial{}, errors.New("divided with reminder")
	}

	return dv, nil

}
