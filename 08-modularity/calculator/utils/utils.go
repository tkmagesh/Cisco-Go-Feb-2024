package utils

func IsEven(no int) bool {
	return no%2 == 0
}

func IsOdd(no int) bool {
	return !IsEven(no)
}
