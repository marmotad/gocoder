package main

import "fmt"

func rules(n,base int64) int64 {
	//n:底数，base : 指数； n的base次幂
var rulesNum int64
rulesNum = n * n
//声明一个int64类型的变量，定义初始值为n的2次方
	for base > 2  {
		//若指数> 2 进入循环
		base --
		rulesNum = rulesNum * n
	}
	return rulesNum
}
func main()  {
	fmt.Println(rules(63,2))
}

