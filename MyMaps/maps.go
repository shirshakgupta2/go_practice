package main
import "fmt"

func main()  {
	fmt.Println("Maps in goLang")

	/*syntax of defining maps*/
// 	var a = map[KeyType]ValueType{key1:value1, key2:value2,...}
// b := map[KeyType]ValueType{key1:value1, key2:value2,...}
	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)


	language := make(map[string]string)//map[key]value this is the syntax for map
	// language := make(map[int]string)//map[key]value this is the syntax for map

	language["JS"]="javascript"
	language["RB"]="Ruby"
	language["PY"]="pyhton"

	fmt.Println("List of all languages:  ",language)
	fmt.Println("JS Stands for:  ",language["JS"])
	
	delete(language,"RB")
	fmt.Println("List of all languages:  ",language)

	//loops in goLang

	for key,value := range language{
		fmt.Printf("For key: %v,Value is : %v",key,value)
	}

}