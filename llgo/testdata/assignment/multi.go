package main

func xyz() (int, int) {
	return 123, 456
}

func abc() (int, int) {
	return xyz()
}

func main() {
	a, b := xyz()
	println(a, b)
	b, a = abc()
	println(a, b)
}
