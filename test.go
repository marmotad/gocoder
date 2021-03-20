package main
import "fmt"
func main () {
	var score int
	fmt.Println("请输入成绩:")
	fmt.Scan(&score)
	switch {
	case score >= 90 :
		fmt.Println("A")
	case score >= 80 :
		fmt.Println("B")
	case score >= 70 :
		fmt.Println("C")
	case score >= 60:
		fmt.Println("D")
	default :
		fmt.Println("E")
	}
}