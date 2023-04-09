# 功能陆续迁移到 [CoralBot](https://github.com/BoyChai/CoralBot/)

# BotPlugin

## 概述

基于[CoralBot](https://github.com/BoyChai/CoralBot)的插件开发库

## 兼容性

| BotPlugin | CoralBot |
| --------- | -------- |
| v1.0.0    | v0.3.1   |

## 安装

```
$ go get -u github.com/BoyChai/BotPlugin
```

## 使用

创建示例文件并写入以下代码

```
$ touch example.go
```

```go
package main

import (
	"github.com/BoyChai/BotPlugin"
	"github.com/BoyChai/CoralBot"
)

func main() {
	e := BotPlugin.GetEvent()
	h := BotPlugin.GetHandler()
	//设置插件信息
	BotPlugin.SetInfo(CoralBot.PluginInfo{
		Name:      "demo",
		Summary:   "测试插件demo",
		Version:   "v0.0.1",
		Developer: "BoyChai",
		Email:     "1972567225@qq.com",
	})
    // 创建触发器，这里的触发器需要添加一个RunName参数
	BotPlugin.NewTask(CoralBot.Task{
		Condition: []CoralBot.Condition{
			{
				Key:   &e.Message,
				Value: "插件测试",
			},
		},
		RunName: "demo",
	})
	
    // 将触发之后的命令添加到命令执行器中，这里的第一个值需要和上面的RunName值一样
	BotPlugin.NewHandles("demo", func(event *CoralBot.Event) error {
		_, err := h.Reply(*e, CoralBot.Msg{Message: "测试成功"})
		if err != nil {
			return err
		}
		return nil
	})
    // 构建
	BotPlugin.BuildPlugin()
}
```

CoralBot.Task.RunName要和BotPlugin.NewHandles的第一个参数一样。

## 构建

执行命令

```
$ go build .\example.go
```

构建好的文件名称应该是example.exe,我们把它的后缀修改成coral,例如："example.coral"，文件名称是什么都可以后缀必须为coral。

## 装载

在Coral的运行目录下创建一个"plugin"目录把改好名字的插件丢进去，并创建一个config.yaml。

config.yaml的内容为：

```yaml
handler:
  - name: "all"
    host: "192.168.1.1:5700"
    agreement: "http"
  - name: "demo"
    host: "127.0.0.1:5700"
    agreement: "http"
```

config.yaml格式为:

```yaml
handler:
  - name: ""
    host: ""
    agreement: ""
```

其中name要和插件名称一样(BotPlugin.SetInfo的name)，host是cqhttp的地址，agreement为http。名字如果是all的话就代表全部的插件都可以使用这个配置，如果插件配置存在并且all的配置也存在那么优先使用插件名字的配置。
