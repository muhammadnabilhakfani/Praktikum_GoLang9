package main
import "fmt"

func main() {
	var bilangan int
	fmt.Scan(&bilangan)
	if bilangan < 0 {
		if bilangan  %2 == 0 {
			fmt.Print("genap negatif")
		}
	}
	if bilangan >= 0 {
		fmt.Print("bukan")
	}
}
