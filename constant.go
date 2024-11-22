/*
Go suporta constantes de caracteres, string, booleanos e valores numericos.

`const` declara uma constante

uma constante pode aparecer em qualquer lugar que uma variavel tambem pode

Uma constante vai sempre performar operações matematicas com precisão arbitraria

*/

package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
  fmt.Println(s)

  const n = 500000000

  const d = 3e20 / n
  fmt.Println(d)

  fmt.Println(int64(d))

  fmt.Println(math.Sin(n))
}
