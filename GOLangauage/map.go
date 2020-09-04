package main 


import "fmt"

func main() {
	studentAge:=make(map[string] int)

	studentAge["Arya"]=23
	studentAge["Saurabh"]=27
	studentAge["Prerna"]=21
	studentAge["Akriti"]=19
	studentAge["Sahitri"]=42
	studentAge["Kriti"]=22

	fmt.Println(studentAge)
	fmt.Println(studentAge["Kriti"])
	fmt.Println(len(studentAge))
}