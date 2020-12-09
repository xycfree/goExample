package example01

import (
	"fmt"
	"math"
)

func Example03() {
	/*题目：一个整数，它加上 100 后是一个完全平方数，再加上 168 又是一个完全平方数，请问该数是多少？
	1.程序分析：在10万以内判断，先将该数加上100后再开方，再将该数加上268后再开方，如果开方后的结果满足如下条件，
	即是结果*/
	for i :=0; i< 100000; i++{
		x := int(math.Sqrt(float64(i + 100)))
		y := int(math.Sqrt(float64(i + 100 + 168)))
		if x*x == i + 100 && y*y == i +100 + 168 {
			fmt.Printf("这个数字是%d\n", i)
		}
	}
}