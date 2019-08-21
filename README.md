# 沈师课程成绩查询助手

实验室开发的微信小程序后端，通过和内网的爬虫服务交互来展示成绩课程等等数据，爬虫服务参照[我的另一个项目](https://github.com/czarhao/sync_crawler)，整体基于golang，使用gin框架，返回json格式的数据与小程序端交互。

可以通过docker部署，前后端使用json交互，部署服务直接：

```shell
docker build -t applets .
docker run -p 8080:8080 -d --name="applet" applets
```

提供如下功能

http://127.0.0.1:8080/registered/sno/spw	// 注册服务

http://127.0.0.1:8080/info/sno/spw	// 返回学生信息

http://127.0.0.1:8080/schedule/sno/spw	// 返回学生课程表

http://127.0.0.1:8080/grade/sno/spw	// 返回成绩

http://127.0.0.1:8080/refresh/sno/spw	// 返回成绩

http://127.0.0.1:8080/refresh/types/sno/spw	// 刷新信息

http://127.0.0.1:8080/article/:start	// 返回文章