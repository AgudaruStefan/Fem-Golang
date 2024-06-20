package collections

import "fmt"

var Countries [10]string
var Slice []int
var Codes map[int]string

func init() {
	Countries[0] = "Argentina"
	Countries[1] = "Australia"
	Countries[8] = "Austria"

	//qty := len(Countries)
	fmt.Println("countries saved")
}
