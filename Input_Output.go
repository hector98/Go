package main

import(
  "fmt"
  "bufio"
  "os"
)

func main()  {
/*
Verbos para imprimir y leer tipos de datos
%v = cualquier tipo.
%d = Enteros.
%t = Booleano
%b, %e, %f = Puntos flotantes
%s = Strings
*/
var edad int
fmt.Println("Escribe tu edad: ")
fmt.Scanf("%d\n",&edad)
fmt.Printf("Mi edad es: %d\n",edad)


//Leer datos con el metodo Bufio
reader := bufio.NewReader(os.Stdin)
fmt.Println("Ingresa tu nombre ")
nombre, err := reader.ReadString('\n')
 if err != nil{
   fmt.Println(err)
 }else{
   fmt.Println("Hola "+nombre)
 }
}
