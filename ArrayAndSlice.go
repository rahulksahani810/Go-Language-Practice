package main

import "fmt"


func lengthofSlice(){
	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitslice := fruitarray[1:3]
	fmt.Printf("length of slice %d capacity %d", len(fruitslice), cap(fruitslice))
}
func slice(){
	numa := [3]int{78, 79 ,80}
	nums1 := numa[:] //creates a slice which contains all elements of the array
	nums2 := numa[:]
	fmt.Println("array before change 1",numa)
	nums1[0] = 100
	fmt.Println("array after modification to slice nums1", numa)
	nums2[1] = 101
	fmt.Println("array after modification to slice nums2", numa)
}
func iterationOfArra(){
	a := [...]float64{67.7, 89.8, 21, 78}
	for i := 0; i < len(a); i++ { //looping from 0 to the length of the array
		fmt.Printf("%d th element of a is %.2f\n", i, a[i])
	}
}
 func lengthOfArray(){
	a := [...]float64{67.7, 89.8, 21, 78}
	fmt.Println("length of a is",len(a))
 }
func main() {
	a := [...]string{"USA", "China", "India", "Germany", "France"}
	b := a // a copy of a is assigned to b
	b[0] = "Singapore"
	fmt.Println("a is ", a)
	fmt.Println("b is ", b)	
	fmt.Println("**********************")
	iterationOfArra()
	fmt.Println("**********************")
	lengthOfArray()
	fmt.Println("**********************")
	slice()
	fmt.Println("**********************")
	lengthofSlice()

}