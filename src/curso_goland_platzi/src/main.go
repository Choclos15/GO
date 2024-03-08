package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {

	fmt.Println("Hola Mundo")

	//Declaración de variables
	const pi float64 = 3.14

	const pi2 float64 = 3.15

	fmt.Println("pi: ", pi)
	fmt.Println("pi2: ", pi2)

	//Declaración de Variables Enteras
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	//Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	//Area de un Cuadrado

	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area del Cuadrado: ", areaCuadrado)

	x := 10
	y := 50

	//Suma
	result := x + y
	fmt.Println("SUMA: ", result)

	//RESTA
	result = x - y
	fmt.Println("RESTA: ", result)

	//MULTIPLICACION
	result = x * y
	fmt.Println("MULTIPLICACION: ", result)

	//DIVISION
	result = y / x
	fmt.Println("DIVISION: ", result)

	//MODULO
	result = x % y
	fmt.Println("MODULO: ", result)

	//INCREMENTAL
	x++
	fmt.Println("INCREMENTAL: ", x)

	//DECREMENTAL
	x--
	fmt.Println("DECREMENTAL: ", x)

	//RETO
	//AREAS DE RECTANGULO, TRAPECIO Y CIRCULO
	areaRectangulo := base * altura
	fmt.Println("Area Rectangulo: ", areaRectangulo)

	//Paquete fmt
	//Declaración de variables
	helloMessage := "Hello"
	wordlMessage := "World"

	//Println
	fmt.Println(helloMessage, wordlMessage)
	fmt.Println(helloMessage, wordlMessage)

	//Printf
	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos)
	fmt.Printf("%v tiene más de %v cursos\n", nombre, cursos)

	//Sprintf
	message := fmt.Sprintf("%s tiene más de %d cursos", nombre, cursos)
	fmt.Println(message)

	//Tipo de dato
	fmt.Printf("Tipo de variable de helloMessage: %T\n", helloMessage)
	fmt.Printf("Tipo de variable de cursos: %T\n", cursos)

	//Uso de Funciones
	NormalFunction()
	NormalFunction2("Hola Mundo")
	TripeArgument(1, 2, "Hola Mundo")
	TripeArgument2(1, 2, "Hola Mundo")
	value := ReturnValue(2)
	fmt.Println("Value: ", value)

	value1, value2 := DoubleReturn(2)
	fmt.Println("Value1: ", value1)
	fmt.Println("Value2: ", value2)

	value3, _ := DoubleReturn(2)
	fmt.Println("Value3: ", value3)

	_, value4 := DoubleReturn(2)
	fmt.Println("Value4: ", value4)

	//Ciclos
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println()
	//For While
	cont := 0
	for cont < 10 {
		fmt.Println(cont)
		cont++
	}

	//For forever
	/*
		counterForever := 0
		for {
			fmt.Println(counterForever)
			counterForever++
		}
	*/

	//Condición if
	var1 := 1
	var2 := 2

	if var1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	if var2 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	if var1 == 2 {
		fmt.Println("Es 2")
	} else {
		fmt.Println("No es 2")
	}

	if var2 == 2 {
		fmt.Println("Es 2")
	} else {
		fmt.Println("No es 2")
	}

	//With and
	if var1 == 1 && var2 == 2 {
		fmt.Println("Es verdad")
	} else {
		fmt.Println("Es mentira")
	}

	if var1 == 2 && var2 == 1 {
		fmt.Println("Es verdad")
	} else {
		fmt.Println("Es mentira")
	}

	if var1 == 1 || var2 == 2 {
		fmt.Println("Es verdad")
	} else {
		fmt.Println("Es mentira")
	}

	if var1 == 2 || var2 == 1 {
		fmt.Println("Es verdad")
	} else {
		fmt.Println("Es mentira")
	}

	//Convertir texto a int
	valueTxt, err := strconv.Atoi("53")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(valueTxt)

	//Múltiple condiciones anidadas con Switch

	switch modulo := 5 % 2; modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	valueInt := 1
	switch {
	case valueInt > 100:
		fmt.Println("Es mayor a 100")
	case valueInt < 0:
		fmt.Println("Es menor a 0")
	default:
		fmt.Println("No condición")
	}

	//defer
	fmt.Println("Hola")
	fmt.Println("Mundo")

	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	//continue y  break

	for i := 0; i < 10; i++ {
		fmt.Println(i)

		//Continue
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}

		if i == 6 {
			fmt.Println("Es 6")
			break
		}
	}

	//Arrays
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array, len(array), cap(array))

	//Slice
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	//Metodos en el slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	//Append
	slice = append(slice, 7)
	fmt.Println(slice)

	//Append nueva lista
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice)

	sliceRange := []string{"Hola", "que", "haces"}

	for i := range sliceRange {
		fmt.Println(i)
	}

	IsPalindromo("amor a roma")
	IsPalindromo("casas")
}

func IsPalindromo(text string) {
	var textReverse string

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if text == textReverse {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es palindromo")
	}

	m := make(map[string]int)

	m["Jose"] = 14
	m["Pepito"] = 20
	fmt.Println(m)

	for i, v := range m {
		fmt.Println(i, v)
	}

	valueMap := m["Jose"]
	fmt.Println(valueMap)
	valueMap = m["Josep"]
	fmt.Println(valueMap)
	valueMap2, ok := m["Jose"]
	fmt.Println(valueMap2, ok)
	valueMap2, ok = m["Josep"]
	fmt.Println(valueMap2, ok)

}

func NormalFunction() {
	fmt.Println("Hola Mundo")
}

func NormalFunction2(message string) {
	fmt.Println(message)
}

func TripeArgument(a int, b int, c string) {
	fmt.Println(a, b, c)
}

func TripeArgument2(a, b int, c string) {
	fmt.Println(a, b, c)
}

func ReturnValue(a int) int {
	return a * 2
}

func DoubleReturn(a int) (c, d int) {
	return a, a * 2
}
