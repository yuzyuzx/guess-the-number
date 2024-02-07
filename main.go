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

type AnswerResult struct {
	NumberResult answerResultType
	Message      string
}

type answerResultType int

const (
	isLess    answerResultType = 1
	isGreater answerResultType = 2
	isCorrect answerResultType = 3
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

		result := checkAnswer(answer, userAnswer, answerCount)
		fmt.Println(result.Message)
		if result.NumberResult == isCorrect {
			if isRestart() {
				answerCount = 0
				answer = generateAnswer()
			} else {
				break
			}
		}
	}
}

func promptMessage() map[string]string {
	msg := map[string]string{}

	msg["1001"] = "数当てゲームをしましょう(^^)"
	msg["1002"] = "正解は１から１００までのランダムな数字です。"
	msg["1003"] = "数字（半角）を入力してください"
	msg["1004"] = "もう一度ゲームをプレイしますか？"
	msg["1005"] = "y:もう一度プレイする n:ゲームを終了する"
	msg["1006"] = "もう一度ゲームを始めます(^^)"
	msg["1007"] = "ありがとうございました(⌒▽⌒)"

	msg["2001"] = "残念(><) もっと小さい数字です。\nもう一度トライしてみてください。"
	msg["2002"] = "残念(><) もっと大きい数字です。\nもう一度トライしてみてください。"
	msg["2003"] = "おめでとうございます(≧∇≦) 正解です！"

	msg["9001"] = "読み込みに失敗しました(T_T)\nもう一度お試しください"
	msg["9002"] = "入力が正しくありません(T_T)\n数字を入力してください"

	return msg
}

/**
* ゲーム開始のメッセージ
 */
func printGameStartPrompt() {
	msg := promptMessage()
	fmt.Printf("%s\n%s\n\n%s\n", msg["1001"], msg["1002"], msg["1003"])
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
	msg := promptMessage()
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return -1, errors.New(msg["9001"])
	}

	text := strings.TrimSpace(scanner.Text())
	num, err := strconv.Atoi(text)
	if err != nil {
		return -1, errors.New(msg["9002"])
	}

	return num, nil
}

/**
* 答え合わせ
 */
func checkAnswer(answer, userAnswer, answerCount int) AnswerResult {
	msg := promptMessage()
	result := AnswerResult{}

	switch {
	case answer < userAnswer:
		result.NumberResult = isLess
		result.Message = msg["2001"]
	case userAnswer < answer:
		result.NumberResult = isGreater
		result.Message = msg["2002"]
	default:
		result.NumberResult = isCorrect
		result.Message = fmt.Sprintf(
			"%s\n%d回目で正解しました。\n",
			msg["2003"],
			answerCount,
		)
	}

	return result
}

/**
* 再スタート
 */
func isRestart() bool {
	msg := promptMessage()
	fmt.Println(msg["1004"])
	fmt.Println(msg["1005"])

	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
	}

	reply := strings.TrimSpace(scanner.Text())

	if reply == "y" {
		fmt.Println(msg["1006"])
		return true
	}

	fmt.Println(msg["1007"])
	return false
}
