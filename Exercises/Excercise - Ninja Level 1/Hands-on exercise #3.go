package Excersise___Ninja_Level_1

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main()	{
		s := fmt.Sprintf("%v\t%v\t%t\v", x, y, z)
		fmt.Println(s)

}