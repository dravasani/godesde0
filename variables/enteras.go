package variables

import "fmt"

func MuestroEnteros(){
	var intcomun int
	intde32 := int32(10)
	intde64 := int64(100)
	fmt.Println("int comun es igual ", intcomun)
	fmt.Println("int 32 es igual ", intde32)
	fmt.Println("int 64 es igual ", intde64)
}