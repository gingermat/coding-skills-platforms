package reverse_bits

func reverseBits(num uint32) uint32 {
	var r uint32

	for i := 0; i < 32; i++ {
		r = 2*r + num%2
		num /= 2
	}

	return r
}
