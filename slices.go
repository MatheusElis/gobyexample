/*
Slices são um tipo de dado importante em Go que fornece uma interface mais
poderosa para sequencias do que arrays.

A maior diferença de arrays para slices é que slices são tipados apenas pelos
elementos que eles contem e não pelo numero de elementos.

Um slice não inicializado é igual a nil e tem comprimento 0.

Para criar um slices vazio com comprimento difente de zero, temos que usar o
metodo builtin make.

Podemos adicionar valores no slices da mesma maneira que para os arrays
*/
package main

import (
	"fmt"
	"slices"
)

func main() {
  var s []string
  fmt.Println("uninit:",s,s==nil, len(s)==0)

  s = make([]string, 3)
  fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

  s[0] = "a"
  s[1] = "b"
  s[2] = "c"
  fmt.Println("set:", s)
  fmt.Println("get:", s[2])

  fmt.Println("len:", len(s))

  s = append(s, "d")
  s = append(s, "e", "f")
  fmt.Println("apd:", s)

  c := make([]string, len(s))
  copy(c, s)
  fmt.Println("cpy:", c)

  l := s[2:5]
  fmt.Println("sl1:",l)

  l = s[:5]
  fmt.Println("sl2:", l)

  l = s[2:]
  fmt.Println("sl3:", l)

  t := []string{"g", "h", "i"}
  fmt.Println("dcl:", t)

  t2 := []string{"g", "h", "i"}
  if slices.Equal(t, t2) {
    fmt.Println("t == t2")
  }

  twoD := make([][]int, 3)
  for i:=0; i<3; i++ {
    innerLen := i + 1
    twoD[i] = make([]int, innerLen)
    for j := 0; j < innerLen; j++ {
      twoD[i][j] = i + j
    }
  }
  fmt.Println("2d: ", twoD)
}
