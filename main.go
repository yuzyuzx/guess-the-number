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

	fmt.Println("数当てゲームをしましょう。正解は１から１００までの数字です。")
	fmt.Println("数字を半角で入力してください")

	game:
		for {
			scanner := bufio.NewScanner(os.Stdin)

			if !scanner.Scan() {
				fmt.Println("読み込みに失敗しました")
				continue
			}

			text := strings.TrimSpace(scanner.Text())
			num, err := strconv.Atoi(text)
			if err != nil {
				fmt.Println("入力が正しくありません")
				continue
			}

			inputTimes += 1

			switch {
			case answer < num:
				fmt.Println("もっと小さい数字です")
			case num < answer:
				fmt.Println("もっと大きい数字です")
			default:
				fmt.Printf("正解です！%d回目で正解しました\n", inputTimes)
				break game
			}
		}
}
