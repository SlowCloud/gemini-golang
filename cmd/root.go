/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gemini-golang",
	Short: "gemini LLM과 간단한 대화가 가능한 CLI",
	Long: `gemini LLM과 대화가 가능한 CLI 도구입니다.

사용하기 위해선 GEMINI_API_KEY 환경변수에 api key를 등록해주는 작업이 필요합니다.
	
윈도우의 경우에는 set GEMINI_API_KEY=<your own key>를,
리눅스의 경우에는 GEMINI_API_KEY=<your own key>를 입력해주세요.
	
등록했는지 확인하려면,
윈도우는 echo %%GEMINI_API_KEY%%를,
리눅스는 env | grep GEMINI를 입력해주세요.`,
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gemini-golang.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}
