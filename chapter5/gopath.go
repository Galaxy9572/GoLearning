package chapter5

// 依赖管理三阶段：GOPATH -> GOVENDOR -> GOMOD

// GOPATH默认在 ~/go (unix/linux),%USERPROFILE\go (windows)
// 官方推荐所有项目和第三方库都放在同一个GOPATH下
// 也可以放在不同目录

// go get命令下载第三方库，无法下载的使用第三方工具gopm下载

// go build命令编译文件
// go install产生pkg文件和可执行文件，可执行文件会放到bin目录下
// go run直接编译运行

// GOPATH下的目录结构
// src--源代码
// pkg--中间结果
// bin--可执行文件
