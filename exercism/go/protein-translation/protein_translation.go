package protein

import (
	"errors"
)

var (
	ErrStop        = errors.New("found stop codon")
	ErrInvalidBase = errors.New("invalid codon")
)

func FromCodon(codon string) (string, error) {
	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	}

	return "", ErrInvalidBase
}

func FromRNA(chain string) ([]string, error) {
	var sequence []string

	for i := 0; len(chain) > i; i += 3 {
		chunk := chain[i : i+3]

		switch protein, err := FromCodon(chunk); err {
		case ErrStop:
			return sequence, nil
		case ErrInvalidBase:
			return sequence, err
		default:
			sequence = append(sequence, protein)
		}
	}

	return sequence, nil
}
