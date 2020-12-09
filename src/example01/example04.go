package example01

import "fmt"

func Example04() {
	/*题目：输入某年某月某日，判断这一天是这一年的第几天？
	1.程序分析：以3月5日为例，应该先把前两个月的加起来，然后再加上5天即本年的第几天，
	特殊情况，闰年且输入月份大于3时需考虑多加一天。*/
	var y, m, d int = 0, 0, 0
	var days int = 0
	fmt.Scanf("%d%d%d\n", &y, &m, &d)
	fmt.Printf("%d年%d月%d日", y, m, d)
	switch m {
	case 12:
		days += d
		d = 30
		fallthrough
	case 11:
		days += d
		d = 31
		fallthrough
	case 10:
		days += d
		d = 30
		fallthrough
	case 9:
		days += d
		d = 31
		fallthrough
	case 8:
		days += d
		d = 31
		fallthrough
	case 7:
		days += d
		d = 30
		fallthrough
	case 6:
		days += d
		d = 31
		fallthrough
	case 5:
		days += d
		d = 30
		fallthrough
	case 4:
		days += d
		d = 31
		fallthrough
	case 3:
		days += d
		d = 28
		if (y%400 == 0) || (y%4 == 0 && y%100 != 0) {
			d += 1
		}
		fallthrough
	case 2:
		days += d
		d = 31
		fallthrough
	case 1:
		days += d
	}
	fmt.Printf("是今年的第%d天!\n", days)
}

//func Example041() {
//	var y, m, d int = 0, 0, 0
//	var days int = 0
//	fmt.Scanf("%d%d%d\n", &y, &m, &d)
//	fmt.Printf("%d年%d月%d日", y, m, d)
//
//	if m > 2 {
//		dd := 28
//		if (y%400 == 0) || (y%4 == 0 && y%100 != 0) {
//			dd += 1
//		}
//	}
//
//}

