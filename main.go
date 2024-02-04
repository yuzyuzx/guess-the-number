package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var isGameEnd = false
	var inputNumber int
	inputTimes := 0

	rand.NewSource(time.Now().UnixNano())
	answer := rand.Intn(100) + 1

	for {
		fmt.Println("１から１００までの数字を入力してください")
		fmt.Scanln(&inputNumber)
		if 0 <= inputNumber && inputNumber <= 100 {
      inputTimes += 1
		} else {
			fmt.Println("無効な数字です。１から１００までの数字を入力してください")
    }

		fmt.Printf("debug: answer=%d, input=%d\n", answer, inputNumber)

		for {

			switch {
			case answer == inputNumber:
				fmt.Printf("正解です！！！%d回目で正解しました", inputTimes)
				isGameEnd = true
			case answer < inputNumber:
				fmt.Println("もっと小さい数字です")
			case inputNumber < answer:
				fmt.Println("もっと大きい数字です")
			default:
				fmt.Println("無効な数字です。１から１００までの数字を入力してください")
			}

			if isGameEnd {
				break
			}

			fmt.Scanln(&inputNumber)
			inputTimes += 1
		}

		if isGameEnd {
			break
		}
	}
}
