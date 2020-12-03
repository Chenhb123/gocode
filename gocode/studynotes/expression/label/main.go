package main

func main() {
outer:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j > 2 {
				println("")
				continue outer
			}
			if i > 2 {
				break outer
			}
			print(i, ":", j, " ")
		}
	}
}
