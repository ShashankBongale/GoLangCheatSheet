package main

import (
	"fmt"
	"strconv"
)

func main() {

	/*fmt.Println("Hello lets see how it goes this time")

	//simple datatypes

	var intVar int = 10
	fmt.Println("intVar ", intVar)

	var intVar2 = 20 //initialising using inferred data types
	fmt.Println("intVar2 ", intVar2)

	intVar3 := 30                                                            //initialising using shorthand notation
	fmt.Println("intVar3 ", intVar3, " lets see if this is valid ", intVar2) //just checking how to print multiple variable in single print statement

	var rawStr string = `let see what \n happend here with raw string` //prints exactly as seen here
	fmt.Println(rawStr)

	var intepStr string = "let see what \n happens here with interpreted string" //Prints after interpreting the special characters
	fmt.Println(intepStr)

	//go do not support imlplicit type conversion
	var intStr string = "10"
	intInt, err := strconv.Atoi(intStr)
	if err == nil {
		fmt.Println("Converted int from string ", intInt)
	}

	var ptr *int = &intVar
	fmt.Println("Pointer address ", ptr, " pointer value ", *ptr) */

	//Reading cmdline inputs
	//inputHandler := bufio.NewReader(os.Stdin)
	//input, _ := inputHandler.ReadString('\n')
	//input = strings.TrimSpace(input)
	//fmt.Println("Input from user", input)

	//aggregate datatype => array, slice, maps and structs
	/*var myArr [4]int
	fmt.Println(myArr)

	var myArr2 [3]int = [3]int{10, 20, 30}
	fmt.Println(myArr2)

	var myslice []int = []int{10, 20, 30} //slice is a reference data type
	myslice = append(myslice, 40, 50, 60)

	fmt.Println("My slice:", myslice)

	var myMap map[int]string = map[int]string{10: "abc", 20: "def", 30: "xyz"} //maps are also reference type.
	fmt.Println(myMap)

	val, isPresent := myMap[40]
	fmt.Println(val, isPresent)

	//nested data structures
	var nestedDs map[int][]string = make(map[int][]string) //have to assign map, otherwise the entry will be nil. This is because map refers to a memory and if we do not assign the memory will be null
	nestedDs[10] = []string{"abc", "def", "xyz"}
	fmt.Println("Nested data structure", nestedDs)

	//nested maps
	var nestedMaps map[string]map[int]string = make(map[string]map[int]string)
	nestedMaps["firstEntry"] = map[int]string{10: "abc"}
	fmt.Println("Nested map", nestedMaps)

	//working with struct
	type customStruct struct { //definig custom stuctures
		name string
		id   int
	}

	var s1 customStruct
	s1.name = "abc"
	s1.id = 10
	fmt.Println(s1)*/

	//looping
	/*var infinitItr int = 0   //infinite loop
	for {
		fmt.Println(infinitItr)
		infinitItr += 1
	}*/

	/*var itr int = 0
	for itr < 3 {
		fmt.Println(itr)
		itr += 1
	}

	for iterator := 0; iterator < 3; iterator++ {
		fmt.Println(iterator)
	}

	//lvar nestedDs map[int][]string = make(map[int][]string) ooping through collections
	var collections map[int]string = map[int]string{10: "abc", 20: "def", 30: "xyz"}
	for key, value := range collections { //range is important here
		fmt.Println(key, value)
	}

	//working with functions
	var num int = TrialFun(10, 20)
	fmt.Println("Value returned from fun", num)

	returnStr := TrialFun2("abc", "def", 10)
	fmt.Println("String returned", returnStr)

	fmt.Println("Passing string list to function")
	TrialFun3("abc", "def", "xyz")

	fmt.Println("Calling function from a different package")
	result := Helper.AddNumber(100, 200)
	fmt.Println("Added number", result)

	result = Helper.AddNumber2(100, 200, 300, 400)
	fmt.Println("Range sum", result)

	//Here CustomInt is a data type and IsEven is a method which is tightly coupled with the data type
	var myCustomInt Helper.CustomInt = 10
	isEven := myCustomInt.IsEven()
	fmt.Println("Is Even", isEven)

	//working with interfaces and assertions
	var printerObj Helper.Printer = Helper.Book{AuthorName: "someone", Lang: "eng"}
	printedStatement := printerObj.Print()
	fmt.Println(printedStatement)

	printerObj = Helper.Vehicle{TypeOfVehicle: "suv", CountryOfOrigin: "IND"}
	printedStatement = printerObj.Print()
	fmt.Println(printedStatement)

	//getting the variables back using assertion
	convertedBackVar, ok := printerObj.(Helper.Vehicle) //if the type case is not success we will get ok as false. Similar to dynamic typecasting in C++
	fmt.Println("Converted back variable", convertedBackVar, "is ok", ok) */

	//working with go routine
	//var wg sync.WaitGroup
	channel := make(chan string)

	//wg.Add(1)

	/*go func() {
		fmt.Println("Lets do this asynch")
		channel <- "This is a message sent over channel"
	}()*/

	//channel <- "This is a message sent over channel"

	go func() {
		fmt.Println("Asynch2")
		recvMsg := <-channel
		fmt.Println("Msg received", recvMsg)
		//wg.Done()
	}()
	channel <- "This is a message sent over channel"

	fmt.Println("Statement after go routine")
	//wg.Wait()
}

func TrialFun(num1 int, num2 int) int {
	fmt.Println("Inside trialfun")
	return num1 + num2
}

func TrialFun2(string1, string2 string, num1 int) string {

	str := string1 + string2 + strconv.Itoa(num1)
	return str
}

func TrialFun3(strlist ...string) {

	for _, val := range strlist {
		fmt.Println(val)
	}
}
