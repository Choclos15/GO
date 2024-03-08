package main

import "fmt"

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

}
