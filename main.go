package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
)

var flag bool

func main() {
	var AmmunitionChan chan string
	var AimChan chan string
	var LaunchChan chan string
	var resultChan chan string
	var Flag chan bool
	var Flag1 chan bool
	var Flag2 chan bool
	exitChan := make(chan bool, 4) //标志main是否退出
	AmmunitionChan = make(chan string, 10)
	AimChan = make(chan string, 5)
	LaunchChan = make(chan string, 3)
	resultChan = make(chan string, 20)
	//记录瞄准加入结果的管道
	Flag = make(chan bool)
	Flag1 = make(chan bool)
	Flag2 = make(chan bool)
	flag = true

	go Ammunition(AmmunitionChan, Flag, exitChan)
	go Aim(AimChan, Flag1, exitChan)
	go Launch(LaunchChan, Flag2, exitChan)
	go result(AmmunitionChan, LaunchChan, AimChan, resultChan, Flag, Flag1, Flag2, exitChan)
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		close(resultChan)
	}()
	for {
		i := 0
		v, ok := <-resultChan
		if !ok {
			break
		}
		fmt.Printf("%s", v)
		i++
		if i%3 == 0 {
			fmt.Println()
		}
	}
	char, _, err := keyboard.GetSingleKey()
	if err != nil {
		panic(err)
	}
	if string(char) == "q" {
		flag = false
	}

}

func Ammunition(AmmunitionChan chan string, Flag, exitChan chan bool) {
	for {
		if !flag {
			break
		}
		AmmunitionChan <- "装弹->"
		Flag <- true
	}
	exitChan <- true
	close(Flag)
	close(AmmunitionChan)
}

func Aim(AimChan chan string, Flag1, exitChan chan bool) {
	for {
		if !flag {
			break
		}

		AimChan <- "瞄准->"
		Flag1 <- true
	}
	exitChan <- true
	close(Flag1)
	close(AimChan)
}

func Launch(LaunchChan chan string, Flag3, exitChan chan bool) {
	for {
		if !flag {
			break
		}

		LaunchChan <- "发射!"
		Flag3 <- true
	}
	exitChan <- true
	close(Flag3)
	close(LaunchChan)
}

func result(AmmunitionChan, LaunchChan, Aim, resultChan chan string, Flag, Flag1, Flag2, exitChan chan bool) {
	for {
		if !flag {
			break
		}
		<-Flag
		a := <-AmmunitionChan
		resultChan <- a
		<-Flag1
		aim := <-Aim
		resultChan <- aim
		<-Flag2
		l := <-LaunchChan
		resultChan <- l
	}
	exitChan <- true
}
