package main
import "fmt"
func main() {
	var x, y int
	fmt.Print("Masukkan nilai x: ")
	fmt.Scanln(&x)
	fmt.Print("Masukkan nilai y: ")
	fmt.Scanln(&y)
	result := power(x, y)
	fmt.Println("Keluaran:", result)
}
func power(x, y int) int {
	if y == 0 {
		return 1
	}
	return x * power(x, y-1)
}