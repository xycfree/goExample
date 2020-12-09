package example01

import (
	"bufio"
	"fmt"
	"os"
)

type B bool

func Example08() {
	var score int = 0
	fmt.Println("输入分数:")
	fmt.Scanf("%d", &score)
	fmt.Println(B(score >= 90).C("A", B(score >= 60).C("B", "C"))) // 三元表达式
}

func (b B) C(t, f interface{}) interface{} {
	/*
		三元表达式，利用条件运算符的嵌套来完成此题：学习成绩>=90 分的同学用 A 表示，60-89 分之
		间的用 B 表示，60分以下的用C表示。
		1.程序分析：(a>b)?a:b 这是条件运算符的基本例子。
	*/

	if bool(b) == true {
		return t
	}
	return f
}

func Example08_1() {
	/*
		最大公约数和最小公倍数,输入两个正整数 m 和 n，求其最大公约数和最小公倍数。
	*/
	var m, n, r, x int
	fmt.Println("输入二位正整数:")
	fmt.Scanf("%d%d", &m, &n)
	x = m * n
	for n != 0 {
		r = m % n
		fmt.Println("r:", r)
		m = n
		n = r
	}
	fmt.Printf("最大公约数:%d, 最小公倍数:%d \n", m, x/m)

}

func Example08_2() {
	/* 统计字符个数, 输入一行字符，分别统计出其中英文字母、空格、数字和其它字符的个数。
	1.程序分析：利用 while 语句, 条件为输入的字符不为' \n'.
	*/
	var i, j, k, l int = 0, 0, 0, 0
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	for _, rune := range input {
		switch {
		case rune >= 'A' && rune <= 'Z':
			i++
		case rune >= 'a' && rune <= 'z':
			i++
		case rune == ' ' || rune == '\t':
			j++
		case rune >= '0' && rune <= '9':
			k++
		default:
			l++
		}
	}
	fmt.Printf("all in all:char=%d space=%d digit=%d others=%d\n", i, j, k, l)
}


func Example08_3() {
	/*
	因子之和, 一个数如果恰好等于它的因子之和，这个数就称为“完数”。例如 6=1＋2＋3.编程找出 1000 以内的所有完数。
	*/
	for n := 2; n<1000 ; n++ {
		m := n
		for i :=1; i<n; i++ {
			if n % i == 0 {
				m -= i
			}
		}
		if m == 0 {
			fmt.Println("完数：", n)
		}
	}
}


func Example08_4() {
	/*
	自由落体
	一球从100米高度自由落下，每次落地后反跳回原高度的一半；再落下，求它在第10次落地时，共经过多少米？第10次反弹多高？
	*/
	s := 100.0
	h := s / 2
	for i := 2; i <= 10; i++ {
		s += 2 * h
		h /= 2
	}
	fmt.Printf("the total of road is %f\n", s)
	fmt.Printf("the tenth is %f meter\n", h)
}

func Example08_5() {
	/* 猴子吃桃
	题目：猴子吃桃问题：猴子第一天摘下若干个桃子，当即吃了一半，还不瘾，又多吃了一个。
	第二天早上又将剩下的桃子吃掉一半，又多吃了一个。以后每天早上都吃了前一天剩下的一半零一个。
	到第 10 天早上想再吃时，见只剩下一个桃子了。求第一天共摘了多少。
	*/
	var x1, x2, day int = 0, 1, 9
	for day > 0 {
		x1 = (x2 + 1) * 2
		x2 = x1
		day--
	}
	fmt.Printf("the total is %d\n", x1)
}

func Example08_6() {
	// 递归方法求阶层 fn = fn_1*4!
	fmt.Printf("%d!=%d\n", 5, fact(5))
}

func fact(n int) (sum  int) {
	if n == 0 {
		return 1
	}
	sum = n * fact(n-1)
	return
}