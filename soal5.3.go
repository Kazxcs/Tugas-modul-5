package main
import "fmt"
func findFactors(n int) []int {
	if n == 1 {
		return []int{1}
	}
	factors := []int{}
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			factors = append(factors, i)
		}
	}
	return factors
}
func main() {
	var n int
	fmt.Print("Masukkan nilai N: ")
	fmt.Scanln(&n)
	factors := findFactors(n)
	fmt.Println("Faktor bilangan dari", n, "adalah:", factors)
}