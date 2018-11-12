## 起步
> go mod需要**1.11**以后的版本才能使用
>
> 使用需要设置环境变量 `GO111MODULE=on`（可设置三个值，分别是`auto`, `on`, `off`）
>
> 新特性的使用暂时开关，还不确定以后会不会有

注意`go.mod`是可以自己手动编辑的

## 开始
```shell
//开启mod特性开关
export GO111MODULE=on

//拉取例子仓库
git clone https://github.com/flxxyz/go-mod-example.git && cd go-mod-example

//进入例子模块目录
cd src/example

//处理依赖
go mod tidy

//编译可执行文件（可选系统版本 mac,linux,linux64,windows,windows64）
make

//进入编译完成目录
cd ../../bin

//执行
./example -h
```

## mod命令
- `go help mod` 查看mod命令的帮助
- `go mod init <packageName>` 初始化模块，目录下生成`go.mod`
- `go mod tidy` 根据`go.mod`文件来处理依赖关系
- `go list -m all` 显示依赖文件
- `go mod download <path@version>` 下载依赖（path是包的路径，version是包的版本）
- `go mod vendor` 复制所有依赖到模块下的vendor目录
- `go mod graph` 输出模块所需的依赖表
- `go mod verify` 验证模块的完整性

## 来源
[Go 1.11 Modules Wiki](https://github.com/golang/go/wiki/Modules)
