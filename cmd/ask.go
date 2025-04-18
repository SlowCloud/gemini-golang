/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/SlowCloud/gemini-golang/gemini"
	"github.com/spf13/cobra"
)

var (
	geminiModel *gemini.Gemini
	longFlag    *bool
)

var askCmd = &cobra.Command{
	Use:   "ask",
	Short: "간단한 질문 수행",
	Long: `간단한 1줄 질문을 수행할 수 있습니다.
옵션을 통해 이미지나 PDF를 등록하는 등의 작업을 추가로 수행할 수 있습니다.(예정)`,
	Run: func(cmd *cobra.Command, args []string) {

		var words string

		if *longFlag {
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				words += scanner.Text()
			}
		} else {
			words = strings.Join(args, " ")
		}

		answer, err := geminiModel.Ask(words)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(answer)
	},
}

func init() {

	var err error
	geminiModel, err = gemini.New()
	if err != nil {
		log.Fatal(err)
	}

	longFlag = askCmd.Flags().BoolP("long", "l", false, "긴 입력이 필요하다면, 해당 옵션을 입력해주세요. 명령어 인자로 들어간 입력은 무시됩니다.")

	rootCmd.AddCommand(askCmd)
}
