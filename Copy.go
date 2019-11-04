package main

import(
  "fmt"
)

func main(){
  slice := []int{1,2,3,4}
  copia := make([]int,4)

  //Imprimir slice y copia antes de la funcion copy
  fmt.Printf("slice = %d\ncopia = %d\n",slice, copia)

  /*
  Estructura del la funcion copy
  copy(Destino, origen)
  */

  //Imprimir slice y copia despues de utilizar la funcion copy
  copy(copia,slice)
  fmt.Printf("slice = %d\ncopia = %d\n",slice, copia)


}
