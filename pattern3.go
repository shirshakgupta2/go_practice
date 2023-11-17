package main

import "fmt"

func main() {
	fmt.Println("print the pattern ")

	for i:=1;i<=5;i++{
		for j:=1;j<=i;j++{
			fmt.Println(" ")
			
		} 
		for k:=1;k<=(10-i);k++ {
			if i==5||k==(10-i)||k==1||i==1{
				fmt.Print("*");
			} else{
				fmt.Print(" ");
			}
				
		}
		fmt.Println("")

	}
}
