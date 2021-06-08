package main

import (
	"fmt"
	"math"
)

func main()  {
	// int
	maxInt8 := math.MaxInt8
	maxUint8 := math.MaxUint8
	maxInt16 := math.MaxInt16
	maxUint16 := math.MaxUint16
	maxInt32 := math.MaxInt32
	maxUint32 := math.MaxUint32
	maxInt64 := math.MaxInt64
	var maxUint64 uint64 = math.MaxUint64

	/*
		maxInt8 = 127
		maxUint8 = 255
		maxInt16 = 32767
		maxUint16 = 65535
		maxInt32 = 2147483647
		maxUint32 = 4294967295
		maxInt64 = 9223372036854775807
		maxUint64 = 18446744073709551615
	*/
	fmt.Println("===================== Int ===============================")
	fmt.Printf("maxInt8 = %d\n", maxInt8)
	fmt.Printf("maxUint8 = %d\n", maxUint8)
	fmt.Printf("maxInt16 = %d\n", maxInt16)
	fmt.Printf("maxUint16 = %d\n", maxUint16)
	fmt.Printf("maxInt32 = %d\n", maxInt32)
	fmt.Printf("maxUint32 = %d\n", maxUint32)
	fmt.Printf("maxInt64 = %d\n", maxInt64)
	fmt.Printf("maxUint64 = %d\n", maxUint64)

	// float
	maxFloat32 := math.MaxFloat32
	maxFloat64 := math.MaxFloat64

	/*
		maxFloat32 = 340282346638528859811704183484516925440.000000
		maxFloat64 = 179769313486231570814527423731704356798070567525844996598917476803157260780028538760589558632766878171540458953514382464234321326889464182768467546703537516986049910576551282076245490090389328944075868508455133942304583236903222948165808559332123348274797826204144723168738177180919299881250404026184124858368.000000
	*/
	fmt.Println("===================== Float ===============================")
	fmt.Printf("maxFloat32 = %f\n", maxFloat32)
	fmt.Printf("maxFloat64 = %f\n", maxFloat64)

	// float默认是float64
	f1 := 1.56677
	/*
		Type f1 = float64
	 */
	fmt.Printf("Type f1 = %T\n", f1)

}
