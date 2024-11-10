package main
import "fmt"
func main() {
	var N int
	fmt.Print("Masukkan nilai N: ")
	fmt.Scanln(&N)
	fmt.Println("Keluaran:")
	printOddSequence(1, N)
}
func printOddSequence(start, N int) {
	if start > N {
		return
	}
	fmt.Print(start, " ")
	printOddSequence(start+2, N)
}