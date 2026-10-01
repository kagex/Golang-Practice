package diffiehellman

import (
    "crypto/rand"
    "math/big"
)
// Diffie-Hellman-Merkle key exchange
// Private keys should be generated randomly.

func PrivateKey(p *big.Int) *big.Int {
	max := new(big.Int).Sub(p, big.NewInt(2))
	a, err := rand.Int(rand.Reader, max)
	if err != nil {
		panic(err)
	}
	return a.Add(a, big.NewInt(2))
}

func PublicKey(private, p *big.Int, g int64) *big.Int {
	base := big.NewInt(g)
	result := new(big.Int)
	result.Exp(base, private, p)
	return result
}

func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	private := PrivateKey(p)
	public := PublicKey(private, p, g)
	return private, public
}

func SecretKey(private1, public2, p *big.Int) *big.Int {
	result := new(big.Int)
	result.Exp(public2, private1, p)
	return result
}
