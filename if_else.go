/*
Fazer ramificações `if` e `else` é bem simples em go

Você pode ter um if sem utilizar em else

Operadores logicos como && e || é bem util para formar condicionais

Uma intrução pode preceder um condicional if, além disso uma variavel declarada
nessa intrução estará disponivel na ramificação atual e nas proximas ramificações

No Go não é necessario usar parenteses em volta da condição do if, entretanto
é necessario utilizar chaves.

Não existe if in line no Go, é necessario usar um if completo até para coisas
mais simples
*/

package main

import "fmt"

func main() {
  if 7%2 == 0 {
    fmt.Println("7 is even")
  } else {
    fmt.Println("7 is odd")
  }

  if 8%4 == 0 {
    fmt.Println("8 is divisible by 4")
  }

  if 8%2 == 0 || 7%2 == 0 {
    fmt.Println("Either 8 or 7 are even")
  }

  if num := 9; num < 0 {
    fmt.Println(num, "is negative")
  } else if num < 10 {
    fmt.Println(num, "has 1 digit")
  } else {
    fmt.Println(num, "has multiple digits")
  }
}
