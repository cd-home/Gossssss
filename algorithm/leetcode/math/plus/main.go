package main

func Plus(a, b int) int {
	if a == 0 && b == 0 {
		return 0
	}
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	return Plus((a&b)<<1, a^b)
}

func main() {

}
