package main
import ("fmt")

func main() {
	var i = 15.5
	var txt = "Hello World!"
    

	// %v	Prints the value in the default format
	// %#v	Prints the value in Go-syntax format
	// %T	Prints the type of the value
	// %%	Prints the % sign




	fmt.Printf("%v\n", i)
	fmt.Printf("%#v\n", i)
	fmt.Printf("%v%%\n", i)
	fmt.Printf("%T\n", i)
  
	fmt.Printf("%v\n", txt)
	fmt.Printf("%#v\n", txt)
	fmt.Printf("%T\n", txt)

	// %s	Prints the value as plain string
	// %q	Prints the value as a double-quoted string
	// %8s	Prints the value as plain string (width 8, right justified)
	// %-8s	Prints the value as plain string (width 8, left justified)
   
   txt = "Hello"
  
  fmt.Printf("%s\n", txt)
  fmt.Printf("%q\n", txt)
  fmt.Printf("%8s\n", txt)
  fmt.Printf("%-8sshirshak\n", txt)
  

  x := true
  var j = false

  fmt.Printf("%v\n", x)
  fmt.Printf("%t\n", j)//%t can be used for boolean data type and  %v is also valid
}