package main
import ("fmt")

func main() {
  var i string = "Hello"
  var j int = 15
//%T will print the type of the object
//%v will print the value of the object

  fmt.Println(i,"\n",j)
  fmt.Println(i," ",j)
  fmt.Print(i," ",j)
  fmt.Printf("i has value: %v and type: %T\n", i, i)
  fmt.Printf("j has value: %v and type: %T", j, j)
}