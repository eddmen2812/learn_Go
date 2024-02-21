package main

import (
	"fmt"
	"reflect"
	"container/list"
)

func main() {
	fmt.Println("Hola y te cuento que iniciare mi mundo en Go de malware")
	//variables
	var myString string = "Es  una cadena de texto"
	fmt.Println(myString)

	myString = "es una nuevo texto"
	fmt.Println(myString)

	var myString2 = "Esto es una cdaena2 de texto"
	fmt.Println(myString2)

	var myInt int = 16
	myInt = myInt + 4
	fmt.Println(myInt)
	fmt.Println(myInt - 1)
	fmt.Println(myInt)

	fmt.Println(myString, myInt)

	fmt.Println(reflect.TypeOf(myInt))

	var myFloat float64 = 6.5
	fmt.Println(myFloat)
	fmt.Println(reflect.TypeOf(myFloat))

	var myBool bool = false
	myBool = true
	fmt.Println(myBool)

	myString3 := "Otra cadena de texto"
	fmt.Println(myString3)

	//Constantes
	const myCost = "Es mi constante"
	fmt.Println(myCost)

	// control de flujo
	if myInt == 20 && myString == "hola hola 20" {
		fmt.Println("El valor es 20")
	} else if myInt == 15 || myString == "hola hola 15" {
		fmt.Println("es el numero 15")
	} else {
		fmt.Println("eSTO NO ES 20")
	}

	// Array
	var myArray [3]int
	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3
	//myArray[3]=4 //Error
	fmt.Println(myArray[2])
	//fmt.Println(myArray[3])//error

	//Map
	
	myMap := make(map[string]int)
	myMap["Edison"] = 29
	myMap["edd2812"] = 35
	myMap["mendoz"]= 25
	fmt.Println(myMap)

	myMap2 := map[string]int{"EDISON": 26,"EDI2812": 85,"MEN": 45}
	fmt.Println(myMap2)

	var myArray2 [3]string
	fmt.Println(myArray2)

	//Listas

	myList := list.New()
	myList.PushBack(1)
	myList.PushBack(2)
	myList.PushBack(3)
	fmt.Println(myList.Back().Value)

	//Bucles

	for index := 0; index < len(myArray); index ++{
		fmt.Println(myArray[index])
	}

	for index, value := range myArray{
		fmt.Println(index,value)
	}

	//fUNCI[ON]
	myFuction()
	fmt.Println(myF2())

	//Estructuras
	type MyStruct struct{
		name string
		age int
	}

	myStruct := MyStruct{"EDISON",29}
	fmt.Println(myStruct)

}


func myFuction(){
	fmt.Println("Mi Funcion")
}

func myF2()string{
	return "Hola funcion 2"
}

//Este es un comentario, no se ejecuta por el compilador
