package main

func main() {
	/*
		fmt.Println("Func1-1 :", Func1(5, 2))
		fmt.Println("Func1-2 :", Func1(12, 10))
		fmt.Println("Func2-1 :", Func2(5, 2, 5, "yourass"))
		fmt.Println("Func2-1 :", Func2(5, 2, 5, "bol"))
		fmt.Println("Func2-1 :", Func2(5, 2, 5, "tvoyamama"))
	*/

}

func Func1(a, b int) int {
	return a * b
}

func Func2(a, b, c int, s string) int {
	switch s {
	case "lol":
		return a * b * c
	case "bol":
		return a + b + c
	case "yourass":
		return ((a * b) / c) + a + b*c
	default:
		return 0
	}
}
