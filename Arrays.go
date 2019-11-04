package main

import(
  "fmt"
)

func main(){
  //Declarar arreglo    var arreglo[Numero_de_Elementos]int

  //Inicializar arreglo
  array := [3]int{1,2,3}
  fmt.Println(array)

  //Tamañode arreglo (len)
  fmt.Println(len(array))

  //Arreglos multimencionales     var matriz[][]int
  var matriz[2][2]int
fmt.Println(matriz)  
}
