## 前言
主要用于和[useless-files-webpack-plugin](https://github.com/Viyozc/useless-files-webpack-plugin)结合使用删除项目中的没有被引用的代码。
如果您并非和useless-files-webpack-plugin一起使用，那么需要删除的文件列表`unused-files.json`文件格式如下所示
。
```json
[
  "D:\\Work\\src\\assets\\images\\test1.png", 
  "D:\\Work\\src\\assets\\images\\test2.png"
]
```

> 注：useless-files-webpack-plugin对于ts的检测不太友好，所以程序中已经跳过.ts后缀的删除。

## 启动
go run .

## 构建
- 构建mac程序
```
go env -w SET CGO_ENABLED=0
go env -w SET GOOS=darwin
go env -w SET GOARCH=amd64
go build main.go
```
- 构建linux程序
```
go env -w SET CGO_ENABLED=0
go env -w SET GOOS=linux
go env -w SET GOARCH=amd64
go build main.go
```
- 构建windows程序
```
go env -w SET CGO_ENABLED=0
go env -w SET GOOS=windows
go env -w SET GOARCH=amd64
go build main.go
```