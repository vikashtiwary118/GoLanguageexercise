package main 


import "fmt"

func main() {
	const pi float64=3.14641
	x:=5
	isbool:=true
	var name string ="Arg Paul"
	fmt.Printf("%f\n",pi)
	fmt.Printf("%t\n",isbool)
	fmt.Printf("%d\n",x)
	fmt.Printf("%T\n",name)//give type of variable
	fmt.Printf("%b\n",25)//convert integer to binary
	fmt.Printf("%c\n",33)//
	fmt.Printf("%x\n",15)
	fmt.Printf("%e\n",pi)
}