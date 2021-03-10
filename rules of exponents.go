package main

import "fmt"

func rules(n,base int64) int64 {
	//n:底数，base : 指数； n的base次幂
var rulesNum int64
rulesNum = n * n
	for base > 2  {
		base --
		rulesNum = rulesNum * n
	}
	return rulesNum
}
func main()  {
	fmt.Println(rules(63,2))
}
