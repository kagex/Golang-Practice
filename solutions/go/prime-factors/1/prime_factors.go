package primefactors

func Factors(n int64) []int64 {
    factors := []int64{}

    for n%2 == 0 {
        factors = append(factors, 2)
        n /= 2
    }

    for n%3 == 0 {
        factors = append(factors, 3)
        n /= 3
    }

    for d := int64(5); d*d <= n; d += 6 {
        for n%d == 0 {
            factors = append(factors, d)
            n /= d
        }
        for n%(d+2) == 0 {
            factors = append(factors, d+2)
            n /= d+2
        }
    }

    if n > 1 {
        factors = append(factors, n)
    }

    return factors
}
