# go-simplecobra

> 本项目是一个简单的命令行应用框架，是对`hugo`的命令行实现的拙劣模仿，获取更加准确的正确使用方式请前往[hugo](https://github.com/gohugoio/hugo)进行查看

应用框架应该具备以下功能：
- 应用命令行框架 ☑️
- 命令行参数解析 ☑️
- 配置文件解析 ❌

使用`cobra`, `viper`, `pflag`库实现。

遵循模式 `APPNAME VERB NOUN --ADJECTIVE` || `APPNAME COMMAND ARG --FLAG`

## install
```bash
go get github.com/chhz0/gosimplecobra
```

## demo

可以在`demo`目录里查看使用示例，在使用上更加推荐使用`SimpleCommander`接口

直接执行 `gosimplecobra` 命令, 由于设置了`--conf`标志为必须的，所以输出错误信息
```bash
$ gosimplecobra

gosimplecobra init func
gosimplecobra prerun func
panic: required flag(s) "conf" not set
```
补充`--conf`标志信息即可正常执行
```bash
$ gosimplecobra --conf defaule.toml

gosimplecobra init func
gosimplecobra prerun func
gosimplecobra run func
```


执行`-h`命令获取帮助信息
```bash
$ gosimplecobra -h

this is a long description for gosimplecobra

Usage:
  gosimplecobra [command] [flags]
  gosimplecobra [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  echo        Echo anything to the screen
  help        Help about any command
  print       Print anything to the screen

Flags:
  -c, --conf string       loading config name (default "defaule.toml")
  -d, --confPath string   loading config from ? (default ".")
  -h, --help              help for gosimplecobra
  -v, --version           version for gosimplecobra

Use "gosimplecobra [command] --help" for more information about a command.
```

查看`echo`命令的帮助信息
```bash
$ gosimplecobra echo -h

echo is for echoing anything back to the screen.
For many years people have echoed back to the screen.

Usage:
  gosimplecobra echo [command] [flags]
  gosimplecobra echo [command]

Available Commands:
  times       Echo anything to the screen more times

Flags:
  -h, --help   help for echo

Global Flags:
  -c, --conf string   loading config name (default "defaule.toml")

Use "gosimplecobra echo [command] --help" for more information about a command.
```

## Inspiration

- [cobra](https://github.com/spf13/cobra)
- [simplecobra](https://github.com/bep/simplecobra)
- [hugo](https://github.com/gohugoio/hugo)
- [iam/pkg/app](https://github.com/marmotedu/iam)