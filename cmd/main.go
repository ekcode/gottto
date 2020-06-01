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
	lucky645 := make([]int, ElectCount)
	i := 0
	for i < ElectCount {
		randNum := rand.Intn(NumberRange)
		if exists(randNum, lucky645) {
			continue
		} else {
			lucky645[i] = randNum
			i++
		}
	}

	sort.Sort(sort.IntSlice(lucky645))
	fmt.Println("오늘의 행운 번호는", lucky645)

}

func exists(num int, arr []int) bool {
	for _, e := range arr {
		if num == e {
			return true
		}
	}
	return false
}
