/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"clops/utils"

	"github.com/spf13/cobra"
)

// base64Cmd represents the base64 command
var base64Cmd = &cobra.Command{
	Use:   "base64",
	Short: "Do a encode/decode of strings",
	Long: `Do a encode/decode of strings to base64.
  Use example: 
      encode: ./clops base64 --e "string"
      decode: ./clops base64 --d "string" 
  `,
	Run: func(cmd *cobra.Command, args []string) {
		encodeStr, _ := cmd.Flags().GetString("e")
		decodeStr, _ := cmd.Flags().GetString("d")
		if encodeStr != "" {
			encode := utils.EncodeString(encodeStr)
			cmd.Println(encode)
		} else if decodeStr != "" {
			decode := utils.DecodeString(decodeStr)
			cmd.Println(decode)
		}
	},
}

func init() {
	rootCmd.AddCommand(base64Cmd)

	base64Cmd.PersistentFlags().String("e", "", "Encode String")
	base64Cmd.PersistentFlags().String("d", "", "Decode String")
}
