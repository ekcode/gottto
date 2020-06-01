package lotto

import (
	"log"
	"math/rand"
	"sort"
	"time"
)

const (
	NumberRange = 45
	ElectCount  = 6
)

func Generate() []int {

	log.Println("로또 번호를 생성합니다.")

	rand.Seed(time.Now().UnixNano())
	var lucky645 []int
	i := 0
	for i < ElectCount {
		//time.Sleep(1 * time.Second)
		randNum := rand.Intn(NumberRange)
		if randNum == 0 || exists(randNum, lucky645) {
			continue
		} else {
			lucky645 = append(lucky645, randNum)
			sort.Sort(sort.IntSlice(lucky645))
			log.Printf("두구두구둥 %2d %s %v\n", randNum, " => ", lucky645)
			i++
		}
	}

	return lucky645

}

func exists(num int, arr []int) bool {
	for _, e := range arr {
		if num == e {
			return true
		}
	}
	return false
}
