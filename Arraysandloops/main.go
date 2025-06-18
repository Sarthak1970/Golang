package main

import(
	"fmt"
)

func main(){
	var arr [5]int32; //int32=4 bytes thus // 5*4=20 bytes memory allocated
	arr[0] = 10

	// for i:= 0; i < 5; i++ {
	// 	fmt.Println(arr[i])
	// }
	fmt.Println(arr[1:5])

	// arr2:= [5]int32{1, 2, 3, 4, 5} // array with 5 elements
	var arr2[3]int32 = [3]int32{1, 2, 3} // array with 3 elements
	fmt.Println(arr2)

	//Slice
	slice1 := []int32{1, 2, 3, 4, 5} // slice with 5 elements
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	slice1 = append(slice1, 6) // appending an element to the slice
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	//map-key value pair
	mymap := map[string]uint8{"Adam": 23, "Eve": 22} // map with string keys and uint8 values
	fmt.Println(mymap)
	mymap["John"] = 30 // adding a new key-value pair to the map
	fmt.Println(mymap["John"])

	delete(mymap,"John") // deleting a key-value pair from the map
	fmt.Println(mymap)

	//Iterating over a map
	for name,age :=range mymap{
		fmt.Printf("Name : %s Age : %d\n", name, age)
	}
}