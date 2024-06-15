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
	// ランダム数を生成するためのシードを設定
	rand.Seed(time.Now().UnixNano())
	// 1から10までのランダムな数を生成
	target := rand.Intn(10) + 1

	fmt.Println("数当てゲームへようこそ！")
	fmt.Println("1から10までの数を当てて下さい")

	// ユーザーに難易度を選択させる
	difficulty := getDifficulty()

	// ゲームを開始
	maxAttempts := setAttempts(difficulty)
	playGame(target, maxAttempts)
}

// ユーザーに難易度を選択させる関数
func getDifficulty() string {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("難易度を選択して下さい（簡単、普通、難しい）： ")
		difficulty, _ := reader.ReadString('\n')
		difficulty = strings.TrimSpace(difficulty)
		if difficulty == "簡単" || difficulty == "普通" || difficulty == "難しい" {
			return difficulty
		}
		fmt.Println("無効な難易度です。再度選択して下さい。")
	}
}

// 難易度に応じて回答回数を設定する関数
func setAttempts(difficulty string) int {
	switch difficulty {
	case "簡単":
		return 10
	case "普通":
		return 5
	case "難しい":
		return 3
	default:
		return 5
	}
}

// ゲームのロジックを処理
func playGame(target, maxAttempts int) {
	reader := bufio.NewReader(os.Stdin)
	attempts := 0      // 回答回数をカウントする変数
	isCorrect := false // 正解したかどうか

	for attempts < maxAttempts {
		fmt.Print("あなたの予想を入力して下さい： ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// 入力を整数に変換
		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("無効な入力です。数字を入力して下さい。")
			continue
		}

		// ユーザーの入力をチェック
		if guess < target {
			fmt.Println("小さすぎます！")
		} else if guess > target {
			fmt.Println("大きすぎます！")
		} else {
			fmt.Println("おめでとうございます！正解です！")
			isCorrect = true
			break
		}

		attempts++ // 回答回数をカウント
		fmt.Printf("残りの回答回数： %d回\n", maxAttempts-attempts)
	}

	// 回答回数の上限に達した場合のメッセージ
	if !isCorrect {
		fmt.Println("残念！正解に辿り着けませんでした。")
		fmt.Printf("正解は %d でした \n", target)
	}
}
