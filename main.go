// @Title  面試題
// @Description  依照每種肉的數量以及處理時間分別給五位員工處理（使用並行）
// @Author 王信棟
package main

import (
	"fmt"
	"math/rand"
	"time"
)

//meat 肉類類別 定義每種肉含有個數量以及處理時間
type meat struct {
	meatQuan int //數量
	meatSec  int //處理時間
}

var beef = meat{
	meatQuan: 10,
	meatSec:  1,
}
var pork = meat{
	meatQuan: 7,
	meatSec:  2,
}
var chicken = meat{
	meatQuan: 5,
	meatSec:  3,
}

// @Title main
// @Description 主程式利用多執行緒並行開始工作
func main() {
	ch := make(chan string)
	go empolyeeWork("A", ch)
	go empolyeeWork("B", ch)
	go empolyeeWork("C", ch)
	go empolyeeWork("D", ch)
	go empolyeeWork("E", ch)
	<-ch
	<-ch
	<-ch
	<-ch
	<-ch
}

// @Title empolyeeWork
// @Description 先計算總共有多少份肉 並重骰亂數的種子碼，使其每次重新執行時都會拿不一樣的肉。
//再來是依照亂數去分配員工拿到什麼肉，而如果抽到該肉數量為0，就會把i減回去並重新跑一次，直至跑道有數量或是跑到上限。
// @param empolyee string "員工名稱" c chan string "routine 的變數"
func empolyeeWork(empolyee string, c chan string) {
	for i := 0; i < beef.meatQuan+pork.meatQuan+chicken.meatQuan; i++ {
		rand.Seed(time.Now().UnixNano())
		takeNum := rand.Intn(3)
		switch takeNum {
		case 0:
			if beef.meatQuan == 0 {
				i--
			} else {
				takeMeat(beef.meatSec, "beef", empolyee)
				cutMeat("beef", empolyee)
				beef.meatQuan = beef.meatQuan - 1
			}
		case 1:
			if pork.meatQuan == 0 {
				i--
			} else {
				takeMeat(pork.meatSec, "pork", empolyee)
				cutMeat("pork", empolyee)
				pork.meatQuan = pork.meatQuan - 1
			}
		case 2:
			if chicken.meatQuan == 0 {
				i--
			} else {
				takeMeat(chicken.meatSec, "chicken", empolyee)
				cutMeat("chicken", empolyee)
				chicken.meatQuan = chicken.meatQuan - 1
			}
		}
	}
	c <- "FINISH"
}

// @Title takeMeat
// @Description 此處為拿肉並輸出文字顯示該員工拿的肉類是什麼。
// @params t int "處理肉的時間" meat string "肉類名稱" empolyee string "員工名稱"
func takeMeat(t int, meat string, empolyee string) {
	fmt.Println(empolyee, "在", time.Now().Year(), "-", time.Now().Month(), "-", time.Now().Day(),
		" ", time.Now().Hour(), ":", time.Now().Minute(), ":", time.Now().Second(), " ",
		"取得", meat)
	time.Sleep(time.Duration(t) * time.Second)
}

// @Title takeMeat
// @Description 此處為處理肉並輸出文字顯示該員工處理的肉類是什麼。
// @params meat string "肉類名稱" empolyee string "員工名稱"
func cutMeat(meat string, empolyee string) {
	fmt.Println(empolyee, "在", time.Now().Year(), "-", time.Now().Month(), "-", time.Now().Day(),
		" ", time.Now().Hour(), ":", time.Now().Minute(), ":", time.Now().Second(), " ",
		"處理完", meat)
}
