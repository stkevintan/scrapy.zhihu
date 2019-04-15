# Scrapy for Zhihu

一个使用 go 写的知乎 topic 爬虫

1. 使用 chromedp 无头浏览器获取登录 cookie
2. 使用 mysql 储存数据

## Config

需要指定一个`Config.toml`文件，example:

```toml
# 旅游， 互联网
TopicNames = [ "19551556", "19550517" ]

# 知乎帐号信息，可省略
[Account]
Username = "username"
Password = "password"

# mysql连接信息
[MysqlConfig]
DBName = "scrapy" # 数据库名，默认scrapy
TableName = "topic" # 表名， 默认topic
DataSourceName = "root:password@tcp(127.0.0.1:3306)/" # DSN，必须
```
