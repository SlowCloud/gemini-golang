/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/SlowCloud/gemini-golang/gemini"
	"github.com/spf13/cobra"
)

var askCmd = &cobra.Command{
	Use:   "ask",
	Short: "간단한 질문 수행",
	Long: `간단한 1줄 질문을 수행할 수 있습니다.
옵션을 통해 이미지나 PDF를 등록하는 등의 작업을 추가로 수행할 수 있습니다.(예정)`,
	Run: func(cmd *cobra.Command, args []string) {
		gemini, err := gemini.New()
		if err != nil {
			log.Fatal(err)
		}

		words := strings.Join(args, " ")

		answer, err := gemini.Ask(words)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(answer)
	},
}

func init() {
	rootCmd.AddCommand(askCmd)
}
