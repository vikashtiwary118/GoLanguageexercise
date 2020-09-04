package main 


import "fmt"

func main() {
	numSlice:=[]int{5,4,3,2,1}

	slice2:=make([]int,5,6)

	copy(slice2,numSlice)
	fmt.Println(slice2)
	fmt.Println(numSlice)
	slice3:=append(numSlice,3,0,-1)
	fmt.Println(slice3)
}