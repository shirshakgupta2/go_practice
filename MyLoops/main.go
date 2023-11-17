package main
import "fmt"

func main() {
	fmt.Println("Welcome to  loops of goLang")

	days := []string{"Sun", "Mon", "Tue", "Wed", "Thu",  "Fri", "Sat"}// its the slice of days

	fmt.Println(days)
	// type one of loop
	for i:=0;i<len(days);i++ {
		fmt.Println(days[i])
	}
	
	//type two of loop
	
	for i := range days{// range key will autometically loop over any slice or array  
		// i gives up the index of the slice
		fmt.Println(days[i])
		
	}

	for index,day:=range days{// range key will autometically loop over any slice or array}

		fmt.Printf(" index is %v and value is %v\n",index,day)
	}
	for _,day:=range days{// range key will autometically loop over any slice or array}

		fmt.Printf("value is %v\n",day)
	}

	randomValue :=1

	for randomValue<20{
		if randomValue==6{
			randomValue++
			continue
		} 
		if randomValue==9{
			goto shir
		}

		fmt.Println("Value is: ",randomValue)
		randomValue++
	}


	// define a label  which is used by goto 

	shir:
		fmt.Println("hello shirshak you have reached to the random value of 9")
		
}