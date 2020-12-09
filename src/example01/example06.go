package example01

import "fmt"

func Example06() {
	/*
	题目：输出 9*9 口诀。
	1.程序分析：分行与列考虑，共 9 行 9 列，i 控制行，j 控制列。
	*/
	for i:=1; i<10; i++ {
		for j:=1; j<=i; j++ {
			fmt.Printf("%d*%d=%-3d", j, i, i*j) // %-3d 表示左对齐， 占3位
		}
		fmt.Printf("\n")
	}
}


func Example06_1() {
	/*
	题目：要求输出国际象棋棋盘。
	1.程序分析：用 i 控制行， j 来控制列，根据i+j的和的变化来控制输出黑方格，还是白方格。
	*/
	for i := 0; i<8; i++ {
		for j := 0; j < 8; j++ {
			if (i+j) % 2 == 0 {
				fmt.Printf("■")
			} else {
				fmt.Printf("□")
			}
		}
		fmt.Printf("\n")
	}
}
