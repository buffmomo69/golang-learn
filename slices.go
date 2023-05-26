//slices are like array but can be resized when needed
//we use .append method to add to slices

package main

import "fmt"


func slic () {
  var firstslice []int

  firstslice = append(firstslice, 1)
  firstslice = append(firstslice,2)
 
  //shorthand method to declare slices
  secondslice := []int{1,2,3,4}
 
  //methods
  //len(slice) -> to find the length of the slice
  //cap(slice) -> to find the capacity of the slice

  fmt.Println(firstslice)
  fmt.Println(secondslice)

  //exercise
  exerciseSlice := []string {"hello","world","!"}
  fmt.Println(exerciseSlice)
}
