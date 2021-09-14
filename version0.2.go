package main

import (
	"fmt"
	"math/rand"
	"time" // seed 생성용 패키지
)

// 추출된 난수 소수 판정 프로그램 0.2
func main() {
	seed := time.Now().Unix()
	rand.Seed(seed)

	count := 0
	number := rand.Intn(150) + 2 // 0과1 제외하고 2~151 수
	fmt.Println("임의로 추출된 수 : ", number)

	for i := 2; i < number; i++ { // 1과 number값일때 반복문이 실행되지 않음
		if number%i == 0 {
			count = count + 1 // count++
		}
	}

	if count == 0 {
		fmt.Println(number, "는(은) 소수입니다.")
	} else {
		fmt.Println(number, "는(은) 소수가 아닙니다.")
	}
}
