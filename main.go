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

type Game struct {}

func (g Game) Start() {
	fmt.Printf("%s\n%s\n\n%s\n", msg.Print("1001"), msg.Print("1002"), msg.Print("1003"))
}

func (g Game) ContinueQuestion() {
        fmt.Println(msg.Print("1004"))
        fmt.Println(msg.Print("1005"))
}

func (g Game) isRestart() bool {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
	}

	reply := strings.TrimSpace(scanner.Text())

	if reply == "y" {
		return true
	}

	return false
}

type Message struct {}

func (m Message) Print(number string) string {
	msg := map[string]string{}

	msg["1001"] = "数当てゲームをしましょう(^^)"
	msg["1002"] = "正解は１から１００までのランダムな数字です。"
	msg["1003"] = "数字（半角）を入力してください"
	msg["1004"] = "もう一度ゲームをプレイしますか？"
	msg["1005"] = "y:もう一度プレイする y以外:ゲームを終了する"
	msg["1006"] = "もう一度ゲームを始めます(^^)"
	msg["1007"] = "ありがとうございました(⌒▽⌒)"

	msg["2001"] = "残念(><) もっと小さい数字です。\nもう一度トライしてみてください。"
	msg["2002"] = "残念(><) もっと大きい数字です。\nもう一度トライしてみてください。"
	msg["2003"] = "おめでとうございます(≧∇≦) 正解です！"

	msg["9001"] = "読み込みに失敗しました(T_T)\nもう一度お試しください"
	msg["9002"] = "入力が正しくありません(T_T)\n数字を入力してください"

  return msg[number]
}

type AnswerCount struct {
	count int
}

func ResetCountToZero() *AnswerCount {
	return &AnswerCount{}
}

func (c *AnswerCount) Increment() {
	c.count += 1
}

func (c AnswerCount) Get() int {
	return c.count
}

var (
  msg Message
  game Game
)

func init() {
  msg = Message{}
  game = Game{}
}

func main() {
  game.Start()
	answer := generateAnswer()
	userInputCount := ResetCountToZero()

	for {

		userAnswer, err := readUserAnswer()
		if err != nil {
			fmt.Println(err)
			continue
		}

		userInputCount.Increment()

		result := checkAnswer(answer, userAnswer, userInputCount.Get())
		fmt.Println(result.Message)
		if result.NumberResult == isCorrect {
      game.ContinueQuestion()
			if game.isRestart() {
        fmt.Println(msg.Print("1006"))
				userInputCount = ResetCountToZero()
				answer = generateAnswer()
			} else {
        fmt.Println(msg.Print("1007"))
				break
			}
		}
	}
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
		return -1, errors.New(msg.Print("9001"))
	}

	text := strings.TrimSpace(scanner.Text())
	num, err := strconv.Atoi(text)
	if err != nil {
		return -1, errors.New(msg.Print("9002"))
	}

	return num, nil
}

/**
* 答え合わせ
 */
func checkAnswer(answer, userAnswer, answerCount int) AnswerResult {
  var number answerResultType
  var message string

	switch {
	case answer < userAnswer:
    number = isLess
    message = msg.Print("2001")
	case userAnswer < answer:
    number = isGreater
    message = msg.Print("2002")
	default:
		number = isCorrect
		message = fmt.Sprintf(
			"%s\n%d回目で正解しました。\n",
			msg.Print("2003"),
			answerCount,
		)
	}

	result := AnswerResult{
    NumberResult: number,
    Message: message,
  }

	return result
}

