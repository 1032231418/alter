# Alter 配置文件

#Mail  configuration
```
mail:
  username: 'xxxxx@163.com'
  password: 'xxxxx'
  host:     'smtp.163.com'
  port:     "465"
```
#Mail to   configuration
```
mailto:
  mails: "xxxxxx@163.com,xxxx@qq.com"
```
#Dingding to configuration
```
dingding:
  webhook: ''
```
#Weixin to configuration
```
weixin:
  webhook: ''
```
#Image to configuration
```
image:
  base64: ''
```

# Alter 运行参数传递
```
  #go run main.go  config.yaml  mail 这是一封测试邮件 测试邮件内容
```