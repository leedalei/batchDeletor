package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
breakHere:
	var fileName string
	var isContinue string
	fmt.Println("请输入待删除的文件列表，默认为exe同级目录unused-files.json")
	fmt.Println("如需要自定义json文件路径，请输入，否则回车略过")
	fmt.Scanln(&fileName)
	fmt.Println("您输入的路径为：" + fileName)
	batchDel(fileName)
	fmt.Println("输入1继续删除，输入其他退出程序...")
	fmt.Scanln(&isContinue)
	if isContinue == "1" {
		goto breakHere
	}
}

func batchDel(fileName string) {
	var jsonStr string
	if fileName == "" {
		jsonStr = ReadFile("unused-files.json")
	} else {
		jsonStr = ReadFile(fileName)
	}
	var list []string
	if err := json.Unmarshal([]byte(jsonStr), &list); err == nil {
		fmt.Println(list)
	}
	for _, item := range list {
		// 排除ts文件
		if strings.HasSuffix(item, ".ts") {
			fmt.Println("跳过ts文件：" + item)
		} else {
			DelFile(item)
		}
	}
}
