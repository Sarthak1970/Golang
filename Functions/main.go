package main
import (
	"fmt"
	"bufio"
	"os"
)

func passByValue(x int) {
	x = 10 // This will not change the original value of x in main
	fmt.Println("Inside passByValue:", x)
}

func passByReference(x *int) {
	*x = 20 // This will change the original value of x in main
	fmt.Println("Inside passByReference:", *x)
}

func main(){
	fmt.Println("Functions in Go")
	reader := bufio.NewReader(os.Stdin) //'bufio' is used to read different things like string, slice etc
									//os.Stdin is used to tell the 'reader' where to read from

	fmt.Println("Enter a number:")
	input, _ := reader.ReadString('\n')
	var x int
	fmt.Sscanf(input, "%d", &x)

	fmt.Println("Original value of x:", x)
	passByValue(x) // Call by value
	fmt.Println("After passByValue, original x:", x)
	passByReference(&x) // Call by reference
	fmt.Println("After passByReference, original x:", x)	
}