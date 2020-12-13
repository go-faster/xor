package xor

// Bytes xors the bytes in a and b. The destination should have enough
// space, otherwise Bytes will panic. Returns the number of bytes xor'd.
func Bytes(dst, a, b []byte) int {
	return xorBytes(dst, a, b)
}

// Words XORs multiples of 4 or 8 bytes (depending on architecture.)
// The slice arguments a and b are assumed to be of equal length.
func Words(dst, a, b []byte) {
	xorWords(dst, a, b)
}
