/*
Em Go for é o unico construtor de loop.

O loop for mais basico é o com uma unica condição

O classico é Inicio/Condição/Depois `for` loop

Outro caminho é o "faça isso N vezes" é utilizar `range` em um inteiro

Um `for` sem condição irá repetir infinitamente até encontrar um `break`

Também é possivel usar um `continue` para a proxima interação do loop
*/

package main

import "fmt"

func main() {
  i:= 1
  for i<= 3{
    fmt.Println(i)
    i = i + 1
  }

  for j:= 0; j<3;j++{
    fmt.Println(j) 
  }

  for i := range 3{
    fmt.Println("range", i)
  }

  for {
    fmt.Println("loop")
    break
  }

  for n := range 6 {
    if n%2 == 0{
      continue
    }
    fmt.Println(n)
  }

}
