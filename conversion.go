package mumblelink

// stringTo256Uint16Slice converts a string to a uint16 slice that has a length of 256.
func stringTo256Uint16Slice(input string) [256]uint16 {
	if len(input) > 256 {
		panic("tried passing an input larger than 256")
	}

	var out [256]uint16
	for i, c := range input {
		out[i] = uint16(c)
	}

	return out
}

// stringTo256ByteSlice converts a string to a byte slice that has a length of 256.
func stringTo256ByteSlice(input string) [256]byte {
	if len(input) > 256 {
		panic("tried passing an input larger than 256")
	}

	var out [256]byte
	for i, c := range input {
		out[i] = byte(c)
	}

	return out
}

// stringTo2048Uint16Slice converts a string to a uint16 slice that has a length of 2048.
func stringTo2048Uint16Slice(input string) [2048]uint16 {
	if len(input) > 2048 {
		panic("tried passing an input larger than 2048")
	}

	var out [2048]uint16
	for i, c := range input {
		out[i] = uint16(c)
	}

	return out
}
