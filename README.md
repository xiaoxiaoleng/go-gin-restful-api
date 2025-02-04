# go-gin-restful-api

### 一个gin的项目Demo,主要用于RESTFul Api和 gin的实践.项目结构和代码用于新项目的参考项目

1. #### RESTful api路由
2. #### gorm 作为 mysql orm
3. #### redis 数据库
4. #### ratelimit 限流
5. #### Viper 配置文件,动态修改生效
6. #### swagger api文档生成
7. #### jwt 登录鉴权
8. #### 集成阿里云OSS对象存储
9. #### 集成腾讯云SMS 发送短信验证码
10. #### 集成 wechat SDK,小程序登录,发送通知等

````
 ██████╗  ██████╗        ██████╗ ██╗███╗   ██╗      ██████╗ ███████╗███████╗████████╗███████╗██╗   ██╗██╗       █████╗ ██████╗ ██╗
██╔════╝ ██╔═══██╗      ██╔════╝ ██║████╗  ██║      ██╔══██╗██╔════╝██╔════╝╚══██╔══╝██╔════╝██║   ██║██║      ██╔══██╗██╔══██╗██║
██║  ███╗██║   ██║█████╗██║  ███╗██║██╔██╗ ██║█████╗██████╔╝█████╗  ███████╗   ██║   █████╗  ██║   ██║██║█████╗███████║██████╔╝██║
██║   ██║██║   ██║╚════╝██║   ██║██║██║╚██╗██║╚════╝██╔══██╗██╔══╝  ╚════██║   ██║   ██╔══╝  ██║   ██║██║╚════╝██╔══██║██╔═══╝ ██║
╚██████╔╝╚██████╔╝      ╚██████╔╝██║██║ ╚████║      ██║  ██║███████╗███████║   ██║   ██║     ╚██████╔╝███████╗ ██║  ██║██║     ██║
 ╚═════╝  ╚═════╝        ╚═════╝ ╚═╝╚═╝  ╚═══╝      ╚═╝  ╚═╝╚══════╝╚══════╝   ╚═╝   ╚═╝      ╚═════╝ ╚══════╝ ╚═╝  ╚═╝╚═╝     ╚═╝
 ````

### 一.使用步骤

1. 复制 config-example.yaml 并重命名为 config.yaml,填上自己的配置
2. go run main.go
3. 访问 http://localhost:8000/swagger/index.html 即可看到接口文档

### 二.接口列表,路由参考

#### token

1. POST /token/username-pwd 用户名密码登录
2. POST /token/phone-pwd 手机号密码登录
3. POST /token/phone-code 手机号验证码登录
4. POST /token/open-id 小程序 openid 登录

#### user

1. POST /users 用户普通注册
2. PATCH /users/password 新旧密码修改密码
3. PATCH /users/password/phone-code 手机号验证码修改密码
4. PATCH /users/authority 更改用户权限
5. GET /users 获取用户列表

#### verification

1. GET /verification/sms 手机号获取短信验证码

### 三.配置文件

````yaml
app:
  pageSize: 10

server:
  mode: debug
  addr: 8000
  readTimeout: 60
  writeTimeout: 60

  aliOSS:
    endpoint: oss-cn-hangzhou.aliyuncs.com
    accessKeyId: [ 阿里云accessKeyId ]
    accessKeySecret: [ 阿里云accessKeySecret ]
    bucketName: [ bucket ]
    bucketUrl: [ bucket url ]

database:
  type: mysql
  user: xxx
  password: xxxx
  host: xx.xx.xxx.xx:3306
  name: [ 数据库名 ]

tencent:
  secretId: [ 腾讯云secretId ]
  secretKey: [ 腾讯云secretId ]

redisDB:
  addr: xx.xxx.xx.xx:6379
  password: xxxxx
  db: 0

jwt:
  signingKey: [ jwt signingKey ]
  expiresTime: [ 604800, token过期时长 ]
  bufferTime: [ bufferTime ]

zap:
  level: 'info'
  format: 'console'
  prefix: '[AKA_ZWZ]'
  director: 'log'
  linkName: 'latest_log'
  showLine: true
  encodeLevel: 'LowercaseColorLevelEncoder'
  stacktraceKey: 'stacktrace'
  logInConsole: true
````

### 四.鸣谢

[![jetbrains](./public/jetbrains.png)](https://www.jetbrains.com/?from=go-gin-restful-api)
