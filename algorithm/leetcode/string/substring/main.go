package main

func FirstSubStringIndex(source, target string) (i int) {
	if source == "" || target == "" {
		return -1
	}
	for i = 0; i < len(source)-len(target)+1; i++ {
		for j := 0; j < len(target); j++ {
			if source[i+j] == target[j] {
				return i
			}
		}
	}
	return -1
}

func main() {

}
