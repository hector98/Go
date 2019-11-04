package main

import(
  "fmt"
)

func main(){

//  Declarar Slice vacio
  var slice_vacio[]int
  fmt.Println(slice_vacio)


  //Inicializa Slice
  slice :=[]int{1,2,3,4}
  fmt.Println(slice)

  // Tamaño del slice
  fmt.Println(len(slice))

  // Partir arreglo con slice
  a := [4] int {1,2,3,4}

      //De la pocision 0 a la pocision 3
      //[Pocisoion_inicial:Pocision_Final]
  sl := a[:3]
  fmt.Println(sl)
}
