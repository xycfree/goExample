package example01

import (
	"fmt"
	"math"
)

func Example07() {
	/**
	题目：古典问题：有一对兔子，从出生后第3个月起每个月都生一对兔子，小兔子长到第三个月后每个月又生一对兔子，
	假如兔子都不死，问每个月的兔子总数为多少？
	1.程序分析： 兔子的规律为数列 1,1,2,3,5,8,13,21....   斐波那契数列
	*/
	var m1, m2 int = 1, 1
	for i :=1; i<12; i++ {
		fmt.Println(m1, m2)
		m1 += m2
		m2 += m1
	}
}

func Example07_1() {
	/*
	题目：判断 101-200 之间有多少个素数，并输出所有素数。
	1.程序分析：判断素数的方法：用一个数分别去除2到sqrt(这个数)，如果能被整除，则表明此数不是素数，反之是素数。
	*/
	count := 0
	var j int = 0
	for i := 101; i <= 200; i++ {
		k := int(math.Sqrt(float64(i)))
		for j=2; j <= k; j++ {
			if i % j == 0 {
				break
			}
		}
		if j == k + 1 {
			fmt.Println(i)
			count++
		}
	}
	fmt.Printf("total:%d", count)
}

func Example07_2() {
	/*
	题目：打印出所有的“水仙花数”，所谓“水仙花数”是指一个三位数，其各位数字立方和等于该数本身。
	例如：153 是一个“水仙花数”，因为 153=1 的三次方＋5 的三次方＋3 的三次方。
	1.程序分析：利用for循环控制100-999个数，每个数分解出个位，十位，百位。
	*/
	for i := 100; i<1000; i++ {
		j := i / 100
		k := i /10 % 10
		m := i % 10
		if j*j*j + k*k*k + m*m*m == i {
			fmt.Printf("%d^3+%d^3+%d^3=%d\n", j, k, m, i)
		}
	}
}