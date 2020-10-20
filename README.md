## 前言
主要用于和[useless-files-webpack-plugin](https://github.com/Viyozc/useless-files-webpack-plugin)结合使用删除项目中的没有被引用的代码。
如果您并非和useless-files-webpack-plugin一起使用，那么需要删除的文件列表json文件格式如如`unused-files.json`所示。
> 注：useless-files-webpack-plugin对于ts的检测不太友好，所以程序中已经跳过.ts后缀的删除。

## 启动
go run .

## 构建
go build .