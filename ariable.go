package main
import ("fmt")

func main() {
  var a, b, c, d int = 1, 3, 5, 7

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
  fmt.Println(d)


  var e, f = 6, "Hello"
  g, h := 7, "World!"

  fmt.Println(e)
  fmt.Println(f)
  fmt.Println(g)
  fmt.Println(h)


  var (
	i int
	j int = 1
	k string = "hello"
  )

 fmt.Println(i)
 fmt.Println(j)
 fmt.Println(k)
erf()

 
}

const A int = 1

func erf() {
  fmt.Println(A)
}

// func main() {
// 	const A = 1
// 	A = 2
// 	fmt.Println(A)