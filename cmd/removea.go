// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"Agenda/service"
	"fmt"

	"github.com/spf13/cobra"
)

// removeaCmd represents the removea command
var removeaCmd = &cobra.Command{
	Use:   "removea",
	Short: "删除会议参与者",
	Long: `该指令用于删除会议参与者
	
	格式: $removea -t [title] -p [participator].`,
	Run: func(cmd *cobra.Command, args []string) {

		//读取参数
		title, _ := cmd.Flags().GetString("title")
		particName, _ := cmd.Flags().GetString("participator")

		//调用服务
		success, errorMsg := service.RemoveParticipator(title, particName)

		if success {
			fmt.Println("操作成功.")
		} else {
			fmt.Println("操作失败: " + errorMsg)
		}

	},
}

func init() {
	rootCmd.AddCommand(removeaCmd)

	// Here you will define your flags and configuration settings.
	removeaCmd.Flags().StringP("title", "t", "", "标题")
	removeaCmd.MarkFlagRequired("title")
	removeaCmd.Flags().StringP("participator", "p", "", "参与者")
	removeaCmd.MarkFlagRequired("participator")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeaCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
