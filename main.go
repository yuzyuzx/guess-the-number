package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	printGameStartPrompt()
	answer := generateAnswer()
	answerCount := 0

	for {

		userAnswer, err := readUserAnswer()
		if err != nil {
			fmt.Println(err)
			continue
		}

    answerCount += 1

		result := checkAnswer(answer, userAnswer)
		if result == 1 {
			fmt.Println("残念！もっと小さい数字です。もう一度トライしてみてください。")
			continue

		} else if result == 2 {
			fmt.Println("残念！もっと大きい数字です。もう一度トライしてみてください。")
			continue

		} else {
			fmt.Printf("おめでとうございます！%d回目で正解しました。\n", answerCount)
			if isRestart() {
        answerCount = 0
        answer = generateAnswer()
      } else {
        break
      }

		}
	}
}

/**
* ゲーム開始のメッセージ
 */
func printGameStartPrompt() {
	fmt.Println("数当てゲームをしましょう。正解は１から１００までのランダムな数字です。")
	fmt.Println("数字（半角）を入力してください。")
}

/*
* 答えをランダムに生成する
 */
func generateAnswer() int {
	rand.NewSource(time.Now().UnixNano())
	answer := rand.Intn(100) + 1

	fmt.Printf("debug: answer=%d\n", answer)
	return answer
}

/*
*
標準入力からユーザーが入力した値を読み込む
*/
func readUserAnswer() (int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return -1, errors.New("読み込みに失敗しました。もう一度お試しください。")
	}

	text := strings.TrimSpace(scanner.Text())
	num, err := strconv.Atoi(text)
	if err != nil {
		return -1, errors.New("入力が正しくありません。数字を入力してください。")
	}

	return num, nil
}

/**
* 答え合わせ
 */
func checkAnswer(answer, userAnswer int) int {
	switch {
	case answer < userAnswer:
		return 1
	case userAnswer < answer:
		return 2
	default:
		return 3
	}
}

/**
* 再スタート
 */
func isRestart() bool {
	fmt.Println("もう一度ゲームをしますか？")
	fmt.Println("y:もう一度する, n:終わる")
  scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
	}

	text := strings.TrimSpace(scanner.Text())
  fmt.Println(text)

  if(text == "y") {
    fmt.Println("もう一度開始")
    return true
  } 

  fmt.Println("終了します")
  return false
}

func restart() {

}
