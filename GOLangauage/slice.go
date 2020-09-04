package main 


import "fmt"

func main() {
	numSlice:=[]int{5,4,3,2,1}

	sliced:=numSlice[3:5]
	sliced1:=numSlice[:5]
	sliced2:=numSlice[0:5]

	fmt.Println(sliced)
	fmt.Println(sliced1)
	fmt.Println(sliced2)

}