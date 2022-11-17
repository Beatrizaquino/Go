package main

import "fmt"


func slice() {

  var array = [10]int{1,2,3,4,5,6,7,8,9,10}
  
  var a = array[0:5]
  var b = array[5:10]

  fmt.Println(array)
    
  fmt.Println("Array A;",a)
  fmt.Println("Array B;",b)

}