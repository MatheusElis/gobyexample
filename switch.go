/*
Switch condicionais expressa diversos condicionais atravez de muitas branchs

# Temos casos em que temos um switch basico com apenas uma condicional simples

Podemos utilizar virgulas para separa multiplas expressões na mesma instrução
case.

É possivel tambem usar o switch sem utilizar uma expressão inicial. É uma
maneira semelhante a um if. E as expressoes não necessariamente precisam ser
questões.

É possivel usar o switch para analizar o tipo da variavel no lugar de analizar
o valor das variaveis
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Print("one")
	case 2:
		fmt.Print("two")
	case 3:
		fmt.Print("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
