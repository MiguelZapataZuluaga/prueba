package main
//o run main.go
/*
func main() {
  println("it's over 9000!")
}
*/
/* USANDO PAQUETE DE MAIN
import (
  "fmt"
  "os"
)

func main() {
  if len(os.Args) != 2 {
    os.Exit(1)
  }
  fmt.Println("It's over ", os.Args[1])
}
*/

//vARIABLES


import (
  "fmt"
)

func main() {
  var power int
  power = 9000
  fmt.Printf("It's over %d\n", power)
}