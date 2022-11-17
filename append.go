package main

import "fmt"


func appendTest() {

  var array = [10]int{1,2,3,4,5,6,7,8,9,10}
  var numeros = array[0:0]

  numeros = append(numeros,999)
  numeros = append(numeros,666)
  numeros = append(numeros,777)
  

  fmt.Println("Array",array)
  fmt.Println("Números",numeros)

}