package complexnumbers

import "math"

type Number struct {
	real      float64
	imaginary float64
}

func (n Number) Real() float64      { return n.real }
func (n Number) Imaginary() float64 { return n.imaginary }

func (n1 Number) Add(n2 Number) Number {
	return Number{
		real:      n1.real + n2.real,
		imaginary: n1.imaginary + n2.imaginary,
	}
}

func (n1 Number) Subtract(n2 Number) Number {
	return Number{
		real:      n1.real - n2.real,
		imaginary: n1.imaginary - n2.imaginary,
	}
}

func (n1 Number) Multiply(n2 Number) Number {
	return Number{
		real:      n1.real*n2.real - n1.imaginary*n2.imaginary,
		imaginary: n1.real*n2.imaginary + n1.imaginary*n2.real,
	}
}

func (n Number) Times(factor float64) Number {
	return Number{
		real:      n.real * factor,
		imaginary: n.imaginary * factor,
	}
}

func (n1 Number) Divide(n2 Number) Number {
	denom := n2.real*n2.real + n2.imaginary*n2.imaginary
	return Number{
		real:      (n1.real*n2.real + n1.imaginary*n2.imaginary) / denom,
		imaginary: (n1.imaginary*n2.real - n1.real*n2.imaginary) / denom,
	}
}

func (n Number) Conjugate() Number {
	return Number{real: n.real, imaginary: -n.imaginary}
}

func (n Number) Abs() float64 {
	return math.Hypot(n.real, n.imaginary)
}

func (n Number) Exp() Number {
	expA := math.Exp(n.real)
	return Number{
		real:      expA * math.Cos(n.imaginary),
		imaginary: expA * math.Sin(n.imaginary),
	}
}