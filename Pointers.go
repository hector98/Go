package  main

import (
  "fmt"
)

func main(){
  /*
    1. Es una direccion de memoria.
    2. En lugar de valor tenemos la direccion.
    3. X,Y => aa12b => 5.
    4. X => aa12b => 6.
    5. Y  ¿? => 6.
    *T => *int, *string, *Struct
    var puntero *int
  */
  var x, y *int
  entero := 5
  x = &entero
  y = &entero

   fmt.Printf("Direcciones de memoria: \n%d\n%d\n",x,y)

//Alterar valor directamente desde la direccion de memoria
*x = 8
   fmt.Printf("Valor en la direccion de memoria: \n%d\n%d\n",*x,*y)
}
