package triangle

type Kind string

const (
	NaT Kind = "not a triangle"
	Equ Kind = "equilateral"
	Iso Kind = "isosceles"
	Sca Kind = "scalene"
)

func KindFromSides(a, b, c float64) Kind {
    if a <= 0 || b <= 0 || c <= 0 || !(a+b > c && a+c > b && b+c > a){
    	return NaT
    }

    if a == b && b == c {
        return Equ
    }

    if a == b || a == c || b == c {
        return Iso
    }

    return Sca
}
