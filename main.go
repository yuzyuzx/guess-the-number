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

	for answerCount := 1; ; answerCount++ {

		userAnswer, err := readUserAnswer()
		if err != nil {
			fmt.Println(err)
			continue
		}

		if checkAnswer(answer, userAnswer, answerCount) {
			break
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
func checkAnswer(answer, userAnswer, answerCount int) bool {
	switch {
	case answer < userAnswer:
		fmt.Println("残念！もっと小さい数字です。もう一度トライしてみてください。")
		return false
	case userAnswer < answer:
		fmt.Println("残念！もっと大きい数字です。もう一度トライしてみてください。")
		return false
	default:
		fmt.Printf("おめでとうございます！%d回目で正解しました。\n", answerCount)
		return true
	}
}
