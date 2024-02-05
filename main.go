package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	inputTimes := 0

	rand.NewSource(time.Now().UnixNano())
	answer := rand.Intn(100) + 1

	fmt.Printf("debug: answer=%d\n", answer)

	fmt.Println("数当てゲームをしましょう。正解は１から１００までのランダムな数字です。")
	fmt.Println("数字（半角）を入力してください。")

game:
	for {
		scanner := bufio.NewScanner(os.Stdin)

		if !scanner.Scan() {
			fmt.Println("読み込みに失敗しました。もう一度お試しください。")
			continue
		}

		text := strings.TrimSpace(scanner.Text())
		num, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("入力が正しくありません。数字を入力してください。")
			continue
		}

		inputTimes += 1

		switch {
		case answer < num:
			fmt.Println("残念！もっと小さい数字です。もう一度トライしてみてください。")
		case num < answer:
			fmt.Println("残念！もっと大きい数字です。もう一度トライしてみてください。")
		default:
			fmt.Printf("おめでとうございます！%d回目で正解しました。\n", inputTimes)
			break game
		}
	}
}
