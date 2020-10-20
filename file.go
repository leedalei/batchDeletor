package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// ReadFile 读取文件
func ReadFile(fileName string) string {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(content)
}

// DelFile 删除文件
func DelFile(fileName string) {
	err := os.Remove(fileName)
	if err != nil {
		fmt.Println("删除失败:" + fileName)
	} else {
		fmt.Println("删除成功:" + fileName)
	}
}
