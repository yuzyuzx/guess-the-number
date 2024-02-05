package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
  var isGameEnd = false
  var inputNumber string
  inputTimes := 0

  rand.NewSource(time.Now().UnixNano())
  answer := rand.Intn(100) + 1

  fmt.Printf("debug: answer=%d\n", answer)

  fmt.Println("数当てゲームです。正解は１から１００までの数字です。")
  fmt.Println()
  fmt.Println("１から１００までの数字を入力してください")

  for {
    fmt.Scanf("%s", &inputNumber)
    num, err := strconv.Atoi(inputNumber)
    if err != nil {
      fmt.Println(err)
      fmt.Println("入力が正しくありません")
      continue
    }



    if 0 <= num && num <= 100 {
      inputTimes += 1
    }

    switch {
    case answer == num:
      fmt.Printf("正解です！！！%d回目で正解しました\n", inputTimes)
      isGameEnd = true
    case answer < num:
      fmt.Println("もっと小さい数字です")
    case num < answer:
      fmt.Println("もっと大きい数字です")
    default:
      fmt.Println("無効な数字です。１から１００までの数字を入力してください")
    }

    if isGameEnd {
      break
    }
  }
}
