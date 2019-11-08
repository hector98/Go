package main

import (
  "fmt"
)

type User struct{
  edad int
  nombre, apellido string
}

func main(){
  /*
Formas de declarar los valor de la Estructura:
1. hector := User {nombre:"Hector", apellido:"Barrios"}
2. hector := User{21,"Hector","Barrios"}
3. fmt.Println()

  */
hector := User{nombre:"Hector", apellido:"Barrios"}
 fmt.Println(hector)


 //Con punteros
 puntero := new(User)
 puntero.nombre = "Hector"
 fmt.Println(puntero)
}
