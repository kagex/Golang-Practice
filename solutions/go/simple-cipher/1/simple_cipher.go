package cipher

import "strings"

type shift int

type vigenere struct {
	key string
}

func NewCaesar() Cipher {
	return NewShift(3)
}

func NewShift(distance int) Cipher {
	if distance <= -26 || distance >= 26 || distance == 0 {
		return nil
	}
	
	d := distance % 26
	if d < 0 {
		d += 26
	}
	
	return shift(d)
}

func NewVigenere(key string) Cipher {
	if len(key) == 0 {
		return nil
	}
	
	allA := true
	for _, r := range key {
		if r < 'a' || r > 'z' {
			return nil
		}
		if r != 'a' {
			allA = false
		}
	}
	
	if allA {
		return nil
	}
	
	return vigenere{key: key}
}

func (c shift) Encode(input string) string {
	return process(input, string(rune('a'+int(c))), true)
}

func (c shift) Decode(input string) string {
	return process(input, string(rune('a'+int(c))), false)
}

func (v vigenere) Encode(input string) string {
	return process(input, v.key, true)
}

func (v vigenere) Decode(input string) string {
	return process(input, v.key, false)
}

func process(input, key string, encode bool) string {
	var sb strings.Builder
	keyLen := len(key)
	keyIdx := 0

	for _, r := range input {
		if r >= 'A' && r <= 'Z' {
			r = r - 'A' + 'a'
		}

		if r < 'a' || r > 'z' {
			continue
		}

		shiftVal := int(key[keyIdx%keyLen] - 'a')
		
		if !encode {
			shiftVal = -shiftVal
		}

		shifted := int(r-'a') + shiftVal
		
		shifted %= 26
		if shifted < 0 {
			shifted += 26
		}
		sb.WriteByte(byte('a' + shifted))
		keyIdx++
	}

	return sb.String()
}