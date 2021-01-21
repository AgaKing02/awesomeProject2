package main

import "strconv"

type G struct {
}

//var g = G{}
//var gg G

var (
	aName  string
	BigBro string
)

func main() {
	i := 2
	s := "1000"
	if len(s) > 1 {
		//print("in")
		i, _ := strconv.Atoi(s)
		i = i + 5
		print(i)
	}
	print(i)

	//x:=5
	//x, _ = f()
	//
	//print(x)
	//print(x,y)
	str := []string{"sascasc"}
	var p *[]string
	p = &str
	getArray(p)
}

//
//func f() (x,y int) {
//	return 2,3
//}
//func getStrings() (string, *string) {
//	return "", nil
//}
//func getArray(string2 *[]string) {
//	str := string2
//	print(str)
//}
