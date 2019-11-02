package main

import(
  "fmt"
  "strconv"
  )

func main()  {
  //Convertir Entero a String
 edad := 21

 fmt.Println("Mi edad es " + strconv.Itoa(edad) + " años")

 //Convertir String a Entero
 a := "23"
 a_int,_ := strconv.Atoi(a)

 fmt.Println(a_int + 12)
}
