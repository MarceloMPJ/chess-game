package basic

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
