package basic

func IsNumber(r rune) bool {
	return r >= '0' && r <= '9'
}

func RuneToInt(r rune) int {
	return int(r - '0')
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}

	return num
}

func MinUint8(numA, numB uint8) uint8 {
	if numA < numB {
		return numA
	}

	return numB
}

func MaxUint8(numA, numB uint8) uint8 {
	if numA > numB {
		return numA
	}

	return numB
}
