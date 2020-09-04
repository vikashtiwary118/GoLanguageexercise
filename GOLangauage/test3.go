package main 


import "fmt"

type car struct {
	gas_padel uint16 //min 0 max 65535
	break_padel uint16 steering_wheel int16 // -32k - +32k
	top_speed_kmh float64 
}


func main() {
	a_car:= car{gas_padel: 22341, 
		        break_padel: 0, 
		        steering_wheel: 12561,
		        top_speed_kmh: 225.0}

	// a_car:=car{22341,0,12562,225.0}  we can also write that way

    fmt.Println(a_car.gas_padel)
}