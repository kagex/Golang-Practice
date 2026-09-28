package proteintranslation

import "errors"

var (
	ErrStop        = errors.New("stop codon")
	ErrInvalidBase = errors.New("invalid base")
)

var codons = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
}

var stopCodons = map[string]bool{
	"UAA": true,
	"UAG": true,
	"UGA": true,
}

func FromRNA(rna string) ([]string, error) {
	result := []string{}
    
	if len(rna) == 0 {
		return result, nil
	}

	for startPos := 0; startPos < len(rna); startPos += 3 {
		endPos := startPos + 3
		if endPos > len(rna) {
			return nil, ErrInvalidBase
		}

		amino, err := FromCodon(rna[startPos:endPos])
		if errors.Is(err, ErrStop) {
			return result, nil
		}
        
		if err != nil {
			return nil, err
		}
		result = append(result, amino)
	}

	return result, nil
}

func FromCodon(codon string) (string, error) {
	if stopCodons[codon] {
		return "", ErrStop
	}
	if v, ok := codons[codon]; ok {
		return v, nil
	}
	return "", ErrInvalidBase
}
