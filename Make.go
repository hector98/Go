package main

import(
  "fmt"
)

func main(){
                  //Longitud, Capacidad
  slice := make([]int,3,5)

  //Imprimir el slice
  fmt.Printf("El Slice es: %d\n", slice)

  //Imprimir la longitud
  fmt.Printf("Longitud %d\n", len(slice))

  //Imprimir la Capacidad
  fmt.Printf("Capacidad %d\n", cap(slice))

  //Agregar nuevo elemnto al slice
  slice = append(slice,2)
  fmt.Printf("Nuevo elemento '2', %d\n",slice)

}
