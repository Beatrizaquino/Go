package main

import "fmt"


func main() {

  var array = make([]int,59,100)

  for i:=0; i < len(array); i++ {
    array[i] = i

  }
  
  fmt.Println(array)
  

}