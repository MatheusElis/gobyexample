/* 
No Go, variaveis são explicitamente declaradas e usadas pelo compilador para
checar o tipo corretamente para as chamadas das funções.

`var` declara 1 ou mais variaveis

você pode declarar multiplas variaveis e inicializar elas ao declarar

Uma sintax muito comum é `:=` para declarar e inicializar as variaveis
*/

package main

import "fmt"

func main() {
  var a = "initial"
  fmt.Println(a)

  var b, c int = 1, 2
  fmt.Println(b, c)

  var d = true
  fmt.Println(d)

  var e int
  fmt.Println(e)

  f := "apple"
  fmt.Println(f)
}
