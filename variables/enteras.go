package variables

import "fmt"

func ShowIntegers() {
	var intcomun int // esto va a tener 0 por defecto
	intde32 := int32(10)
	intde64 := int64(100)

	fmt.Println(intcomun)
	fmt.Println(intde32)
	fmt.Println(intde64)

}
