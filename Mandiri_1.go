package main
import "fmt"

func main() {
	var motor, orang int
	fmt.Scan(&orang)
	if orang%2 == 1 {
		orang = orang + 1
	}
	motor = orang / 2
	fmt.Print(motor)
}
