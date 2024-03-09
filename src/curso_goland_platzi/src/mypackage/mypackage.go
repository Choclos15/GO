package mypackage

import "fmt"

type CarPublic struct {
	Brand string
	Year  int
}

type CarPrivate struct {
	brend string
	year  int
}

func PrintMessage() {
	fmt.Println("Hola")
}

func PrintMessage2(message string) {
	fmt.Println(message)
}
