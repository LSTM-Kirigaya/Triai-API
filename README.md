## Triai 反向工程

拿到项目，先编译：

```bash
$ go build -o triai.exe
```

使用方法：

```bash
$ triai.exe -userToken 650069b5c88f5af5f34eb58c1f374b5dfb71f7 -image ./images/stardust.png -prompt "美人，园林" -negPrompt "空白" -outputDir ./results
```

其中 `userToken` 获取方式如下：

1. 进入网站，登录：https://www.trikai.com/apps/trikwebapp/create
2. 进入创作，随便生成一两张
3. 打开控制台，全局搜索 `V-User-Token`，它的值就是 `userToken`