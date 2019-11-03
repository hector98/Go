package main

import(
  "fmt"
)

func main()  {
  x := 11
  y := 10

  if x > y{
    fmt.Printf("%d es mayor que %d\n",x,y)
  }else if x < y{
    fmt.Printf("%d es menor que %d",x,y)
  }else{
    fmt.Printf("%d y %d son iguales\n",x,y)
  }

}
