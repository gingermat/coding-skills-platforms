package strand

func ToRNA(dna string) string {
	rna := make([]byte, len(dna))

	for i := range dna {
		switch r := dna[i]; r {
		case 'G':
			rna[i] = 'C'
		case 'C':
			rna[i] = 'G'
		case 'T':
			rna[i] = 'A'
		case 'A':
			rna[i] = 'U'
		default:
			rna[i] = r
		}
	}

	return string(rna)
}
