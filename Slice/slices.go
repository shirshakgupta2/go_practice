package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slice data structure!")

	var fruitList = []string {}
	fmt.Printf("Type of fruit list is %T\n", fruitList)
	var fruitList1 = []string {"apple", "orange", "mango"}
	fmt.Println("fruit list is:=" ,fruitList1)

	fruitList1 = append(fruitList1, "banana","watermelon")
	fmt.Println(fruitList1)


	fruitList1=append(fruitList1[1:])
	fmt.Println(fruitList1)

	fruitList1=append(fruitList1[1:2])
	fmt.Println(fruitList1)

	fmt.Println(fruitList1)


	highscores :=make([]int,4) 
	highscores[0] =987
	highscores[1] =457
	highscores[2] =627
	highscores[3] =888

	fmt.Println(highscores)

	highscores=append(highscores, 555,66,222)

	fmt.Println(highscores)

	fmt.Println("is the slice defined sorted",sort.IntsAreSorted(highscores))
	sort.Ints(highscores)
	fmt.Println(highscores)
	fmt.Println("is the slice defined sorted",sort.IntsAreSorted(highscores))


	// HOW TO REMOVE VALUE FROM SLICES BASED ON INDEX

	var courses=[]string{"reactjs", "js", "swift", "ruby", "pyhton", "java"}
	fmt.Println(courses)
	var index = 2;
	courses = append(courses[:index],courses[index+1:]... )//'.' is the index 
	fmt.Println(courses)	
	



}