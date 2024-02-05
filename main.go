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

  fmt.Printf("debug: answer=%d\n", answer)

  for {
    fmt.Println("数当てゲームです。正解は１から１００までの数字です。")
    fmt.Println("なるべく少ない回数で当ててみましょう。")
    fmt.Println()
    fmt.Println("１から１００までの数字を入力してください")
    // fmt.Scanln(&inputNumber)
    _, err := fmt.Scanf("%d", &inputNumber)
    if err != nil {
      //errors.New("不正な入力")
      fmt.Printf("%s", "不正な入力")
    }

    if 0 <= inputNumber && inputNumber <= 100 {
      inputTimes += 1
    }

    for {

      // fmt.Printf("%v", inputNumber)

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

      // fmt.Scanln(&inputNumber)
      _, err := fmt.Scanf("%d", &inputNumber)
      if err != nil {
        fmt.Printf("%s", "不正な入力")
      }
      if 0 <= inputNumber && inputNumber <= 100 {
        inputTimes += 1
      }
      fmt.Println(inputNumber)
    }

    if isGameEnd {
      break
    }
  }
}
