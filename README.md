# go-simplecobra

> 本项目是一个简单的命令行应用框架，是对`hugo`的commandline实现的拙劣模仿，正确使用方式请前往[hugo](https://github.com/gohugoio/hugo)进行查看

应用框架应该具备以下功能：
- 命令行参数解析
- 配置文件解析
- 应用命令行框架

使用`cobra`, `viper`, `pflag`库实现。

遵循模式 `APPNAME VERB NOUN --ADJECTIVE` || `APPNAME COMMAND ARG --FLAG`

## Inspiration

- [cobra](https://github.com/spf13/cobra)
- [simplecobra](https://github.com/bep/simplecobra)
- [hugo](https://github.com/gohugoio/hugo)
- [iam/pkg/app](https://github.com/marmotedu/iam)