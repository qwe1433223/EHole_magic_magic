/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"ehole/module/finger"
	"ehole/module/finger/source"
	"os"
	"fmt"
	"bufio"
	"strings"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

// fingerCmd represents the finger command
var fingerCmd = &cobra.Command{
	Use:   "finger",
	Short: "ehole的指纹识别模块",
	Long:  `从fofa或者本地文件获取资产进行指纹识别，支持单条url识别。`,
	Run: func(cmd *cobra.Command, args []string) {
		color.RGBStyleFromString("105,187,92").Println("")
		if localfile != "" {
			urls := removeRepeatedElement(source.LocalFile(localfile))
			s := finger.NewScan(urls, thread, output, proxy)
			s.StartScan()
			os.Exit(1)
		}
		if fofaip != "" {
			urls := removeRepeatedElement(source.Fofaip(fofaip))
			s := finger.NewScan(urls, thread, output, proxy)
			s.StartScan()
			os.Exit(1)
		}
		if hunterip != "" {
			urls := removeRepeatedElement(source.Hunterip(hunterip))
			s := finger.NewScan(urls, thread, output, proxy)
			s.StartScan()
			os.Exit(1)
		}
		if fofasearche != "" {
			urls := removeRepeatedElement(source.Fafaall(fofasearche))
			s := finger.NewScan(urls, thread, output, proxy)
			s.StartScan()
			os.Exit(1)
		}
		if huntersearche != "" {
			urls := removeRepeatedElement(source.Hunterall(huntersearche))
			s := finger.NewScan(urls, thread, output, proxy)
			s.StartScan()
			os.Exit(1)
		}
		if urla != "" {
			s := finger.NewScan([]string{urla}, thread, output, proxy)
			s.StartScan()
			os.Exit(1)
		}
		//如果localfile、fofaip、fofasearche、hunterip、huntersearche、urla都为空
		if localfile == "" && urla == "" {
			var urls []string

			// 使用 bufio 包读取标准输入
			reader := bufio.NewScanner(os.Stdin)
		
			for reader.Scan() {
				// 读取管道符传递的每一行内容
				input := reader.Text()
		
				// 去除末尾的换行符（如果有）
				input = strings.TrimSpace(input)
	
				// 将接收到的 URL 添加到列表中
				urls = append(urls, input)
			}
		
			if err := reader.Err(); err != nil {
				fmt.Println("读取输入错误:", err)
				return
			}
	
			s := finger.NewScan(urls, thread, output, proxy)
			s.StartScan()
			os.Exit(1)
		}

	},
}

var (
	fofaip        string
	fofasearche   string
	localfile     string
	urla          string
	thread        int
	output        string
	proxy         string
	hunterip      string
	huntersearche string
)

func init() {
	rootCmd.AddCommand(fingerCmd)
	fingerCmd.Flags().StringVarP(&hunterip, "hip", "a", "", "从hunter提取资产，进行指纹识别，仅仅支持ip或者ip段，例如：192.168.1.1 | 192.168.1.0/24")
	fingerCmd.Flags().StringVarP(&huntersearche, "hunter", "b", "", "从hunter提取资产，进行指纹识别，支持hunter所有语法")

	fingerCmd.Flags().StringVarP(&fofaip, "fip", "f", "", "从fofa提取资产，进行指纹识别，仅仅支持ip或者ip段，例如：192.168.1.1 | 192.168.1.0/24")
	fingerCmd.Flags().StringVarP(&fofasearche, "fofa", "s", "", "从fofa提取资产，进行指纹识别，支持fofa所有语法")
	fingerCmd.Flags().StringVarP(&localfile, "local", "l", "", "从本地文件读取资产，进行指纹识别，支持无协议，列如：192.168.1.1:9090 | http://192.168.1.1:9090")
	fingerCmd.Flags().StringVarP(&urla, "url", "u", "", "识别单个目标。")
	fingerCmd.Flags().StringVarP(&output, "output", "o", "", "输出所有结果，当前仅支持json和xlsx后缀的文件。")
	fingerCmd.Flags().IntVarP(&thread, "thread", "t", 100, "指纹识别线程大小。")
	fingerCmd.Flags().StringVarP(&proxy, "proxy", "p", "", "指定访问目标时的代理，支持http代理和socks5，例如：http://127.0.0.1:8080、socks5://127.0.0.1:8080")
}

func removeRepeatedElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}
