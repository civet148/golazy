# golazy说明

golazy 是一个基于go-zero的api为输入并生成基于gin的web代码工具。
api文件声明路由和接口格式兼容go-zero，但不支持import导入其他api文件和api中声明请求响应结构体。

# 开始使用

## go install安装golazy

```shell
$ go install github.com/civet148/golazy@latest
```

## 使用方法

- **环境准备**

```shell
# golazy代码编译安装
$ gitclone https://github.com/civet148/golazy.git
$ cd golazy && make install
```

- **编译api文件**

```shell
# -o 指定代码文件输出位置（默认当前目录.）
# -s 指定代码文件生成风格（go_lazy：蛇形 goLazy：小驼峰 GoLazy：大驼峰）
$ cd test
$ golazy api go -f test.api -o . -s go_lazy
```

- **运行**

```shell
$ go mod tidy && go run test.go
```

## 生成文件树

```shell
.
├── etc
│   └── test.yaml
├── go.mod
├── internal
│   ├── config # 配置文件目录
│   │   └── config.go
│   ├── gen_models.bat # 内置数据库生成models和dao脚本
│   ├── handler # 路由handlers
│   │   ├── routes.go # 路由注册文件
│   ├── logic # 业务逻辑代码
│   ├── middleware # 中间件代码
│   │   ├── cors_middleware.go # 内置跨域中间件
│   │   ├── jwtauth_middleware.go # 内置JWT鉴权中间件
│   │   └── validator_middleware.go # 自定义中间件
│   ├── svc
│   │   └── service_context.go # 服务上下文
│   └── types # 请求和响应结构体目录
├── test.api # api定义文件
└── test.go # main文件
```


## 定义API文件

- api文件声明路由和接口格式兼容go-zero，但不支持import导入其他api文件和api中声明请求响应结构体。
- prefix 路由前缀（支持/开头也支持不带/）
- group 生成的代码进行分组
- middleware 中间件（Cors内置跨域中间件，JwtAuth内置的JWT鉴权中间件）

```api
@server (
    prefix:     /v1
    middleware: Cors
)
service api {
    @doc "用户登录"
    @handler UserSignIn
    post /sign_in (UserSignInReq) returns (UserSignInRsp)

    @doc "用户注册"
    @handler UserSignUp
    post /sign_up (UserSignUpReq) returns (UserSignUpRsp)
}

@server (
    prefix:     /v1
    middleware: JwtAuth
)
service api {
    @doc "用户退出登录"
    @handler UserSignOut
    post /sign_out (UserSignOutReq) returns (UserSignOutRsp)
}

@server (
    prefix:     /v1/user
    middleware: JwtAuth,Validator
)
service api {
    @doc "用户列表"
    @handler GetUserList
    get /list (GetUserListReq) returns (GetUserListRsp)

    @doc "添加用户"
    @handler AddUser
    put /add (AddUserReq) returns (AddUserRsp)

    @doc "修改用户"
    @handler EditUser
    post /edit (EditUserReq) returns (EditUserRsp)

    @doc "删除用户"
    @handler DeleteUser
    delete /delete (DeleteUserReq) returns (DeleteUserRsp)

    @doc "根据ID查询用户"
    @handler GetUserById
    get /:id (GetUserByIdReq) returns (GetUserByIdRsp)
}


@server (
    prefix:     /v1/ws
)
service api {
    @doc "市场行情（websocket方式）"
    @handler WsMarketList
    get /market (gin.Context) returns (nil)
}

```

