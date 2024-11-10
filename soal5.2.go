package main
import "fmt"
func printStars(n int) {
	if n == 0 {
		return
	}
	fmt.Print("*")
	printStars(n - 1)
}
func main() {
	var n int
	fmt.Print("Masukkan nilai N: ")
	fmt.Scanln(&n)
	for i := 1; i <= n; i++ {
		printStars(i)
		fmt.Println()
	}
}