package main


import "fmt"

// func add(a,b int) int{  // this is also valid
func add(a int, b int) int{
	return a+b
}

func getLanguages() []string{
	return []string{"go","python","java"}
}
// in golang we return generally 2 values one is the value and other is error
// func getLanguages() (string , string , bool){
// 	return "go","python",true
// }

// function is given as an argument to another function
func processIt(fn func(a int)int){
	fn(5);
}

// now if i want to return a function from another function

func processIt2() func(a int)int{
    return func(a int)int {
		return 4
	 }
}

func main(){
	// lang1 , lang2 , lang3 := getLanguages()
	// fmt.Println(lang1 , lang2 , lang3)
	// to print only lang 1 and lang 2
	// lang1 , lang2 , _ := getLanguages() // _ is used to ignore/suppress the value
	// fmt.Println(lang1 , lang2)
	fmt.Println(getLanguages())
     sum :=add(2,3)
	 fmt.Println(sum)

    //  anonymous function no name function
	 fn:= func(a int)int{
		return 2
	 }
	 processIt(fn)

	 fn1 :=processIt2()
	 fn1(5)





}