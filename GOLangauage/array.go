package main 

import "fmt"

func main() {
	var StudentsCount[10] int

	for i := 0; i < 10; i++ {
		StudentsCount[i]=i+1

		fmt.Println(StudentsCount[i])

		
	}
	fmt.Println(StudentsCount)
}