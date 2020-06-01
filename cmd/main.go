package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const (
	NumberRange = 45
	ElectCount  = 6
)

func main() {
	fmt.Println("로또 번호를 생성합니다.")

	rand.Seed(time.Now().UnixNano())
	var lucky645 []int
	i := 0
	for i < ElectCount {
		time.Sleep(1 * time.Second)
		randNum := rand.Intn(NumberRange)
		if randNum == 0 || exists(randNum, lucky645) {
			continue
		} else {
			lucky645 = append(lucky645, randNum)
			sort.Sort(sort.IntSlice(lucky645))
			fmt.Printf("두구두구둥 %2d %s %v\n", randNum, " => ", lucky645)
			i++
		}
	}

	fmt.Println("\n*** 오늘의 행운 번호는", lucky645)

}

func exists(num int, arr []int) bool {
	for _, e := range arr {
		if num == e {
			return true
		}
	}
	return false
}
