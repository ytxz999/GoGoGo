项目框架
```
memo/
├── main.go                  # 入口：初始化数据库 → 注册路由 → 启动 HTTP 服务
├── go.mod / go.sum          # 依赖管理
│
├── model/                   # 领域模型（数据表映射）
│   ├── user.go              #   用户实体
│   └── todo.go              #   待办事项实体 + 状态常量
│
├── database/                # 数据库连接
│   └── mysql.go             #   GORM 连接初始化（单例 DB 句柄）
│
├── dao/                     # 数据访问层（Data Access Object）
│   ├── user_dao.go          #   用户表：创建 / 按用户名查询
│   └── todo_dao.go          #   待办表：创建 / 分页查询 / 更新 / 删除
│
├── service/                 # 业务逻辑层
│   ├── user_service.go      #   注册 / 登录（含 bcrypt 加密 & JWT 签发）
│   └── todo_service.go      #   待办 CRUD（含参数校验 & 业务规则）
│
├── controller/              # 控制器层（HTTP 请求处理）
│   ├── user_controller.go   #   POST /api/register, POST /api/login
│   └── todo_controller.go   #   POST/GET/PUT/DELETE /api/todos
│
├── router/                  # 路由注册
│   └── router.go            #   Gin 引擎配置、路由分组、Swagger 挂载
│
├── middleware/               # 中间件
│   └── auth.go              #   JWT 鉴权中间件（从 Token 提取 userId）
│
├── common/                  # 公共工具
│   └── response.go          #   统一 HTTP 响应格式（Success / BadReq / Unau / NotFound / Internal）
│
├── utils/jwt/               # JWT 工具
│   └── jwt.go               #   Token 生成 & 解析（HS256, 24h 过期）
│
└── docs/                    # Swagger 文档（swaggo 自动生成）
    ├── docs.go
    ├── swagger.json
    └── swagger.yaml
```
项目主要实现：
- 用户注册
- 用户登录
- JWT 身份认证
- Todo 创建
- Todo 查询（分页、状态筛选、关键词搜索）
- Todo 更新
- Todo 删除
- Swagger API 文档


采用了三层架构
controller：处理HTTP请求和响应
service：    处理业务逻辑
dao：         负责数据库操作

采用了gin+gorm

共有以下路由
POST  /api/register
POST  /api/login
GET    /api/todos
POST /api/todos
PUT   /api/todos
DELETE /api/todos
### 链路流程
开始介绍链路
我的整个项目从终端的go run main.go开始
### main
最先使用gorm，利用mysql.go的init初始化，连接数据库以及本地服务器，保证服务器内的动作能实施到数据库中
使用router层的SetupRouter创建并配置路由，返回路由引擎，启动端口8080.
### router
在这个层注册user和todo的各种路由并且调用controller层的函数
在这一层比较特殊的主要是加入了一个swagger路由以及使用路由分组给组内的路由调用中间件AuthMiddleware
**关于swagger**，
	这个我是在项目搞完最后弄的，在contrller层的每个函数前加入特定的注释，swagger会读取并且自动生成一个docs的文件夹，里面包含着接口文档，当启动服务器并且访问/swagger/\*any会出现以下内容，如图所示
![[Pasted image 20260727210914.png]]

关于**AuthMiddleware**，
	由于在备忘录中，不同的用户只能够查看自己的事项，不能查看其他用户的事项。
	所以我在b站上了解到关于cookie和token的知识，以下是我依稀还记得的内容：cookie和token都是来存储专门用户信息的，不过一个是存储在服务器中，一个是存储在浏览器，关于cookie的东西我记不清了，（还有一个有关session）。以下讲一下我对**token的理解**
		首先服务器确认用户的身份是依靠userId的，但是如果直接用userID，那其他用户可以通过修改浏览器对服务器的请求直接登录其他用户的账号，所以引入了token这个东西，token的构成是head（存储加密算法和类型的64编码），claim（存储关键数据和存在时间的64编码），signature（利用前两者的64位编码采用前面说的加密算法加入服务器本身的密钥签名而成），这便是JWT，在jwt文件夹中有两个函数，一个是生成token的也就是前面关于token构成的内容，一个是解析token的，分别应用login功能和AuthMiddleware功能。
		*最初的一个疑惑：为什么还要解析token呢，只要进行生成再比对不就好了，但是深入去想发现，如果要进行比对，那token将需要保存在服务器中，严重影响性能，并且存在时间过了就比不了了*
	所以使用AuthMiddleware，在Header中找到Authorization的用户信息确保后续todo操作是用户自己操控自己的
### controller
在这个层中主要处理来自gin封装好的上下文c，也就是html里面的body，共有两个主要的controller，分别是用户的和事项的
用户的分别是注册和登录功能，事项的分别是创建，查询，更新，删除，但是我感觉其实结构大差不差：
	先创建一个结构体分别与每个要接受Http的请求的参数一致，用来保存用户输入的信息，值得将的是使用ShouldBind(&req)这个函数直接获取。又或者直接c.Query某个值。然后进行分析，通过response（Success or Fail）发出Http状态和关键数据到浏览器返回响应，在发送前使用service层对数据进行业务逻辑处理和判断
### service和dao
在这个层中主要处理核心业务逻辑，判断从controller输入的数据是否符合格式以及数据正确性，再调用dao层将正确的数据通过gorm对数据库进行CRUD操作
其中的dao层，他的参数需要传入用户以及事项的model，也就是数据库表中各列的值的结构体
保证传入
	其中关于用户service层，用户如果将密码存入数据库，如果数据库泄露所有用户的账号将会有危险，所以密码存入使用bcrypt对密码加密，他是将密码和一个随机的salt值进行组合加密，整个过程是单向的，也就是password ---> database.password ,
						 database.password -/->password
	十分的安全。
	在登录过程中，使用bcrypt.CompareHashAndPassword来对密码进行匹配对比
![[Pasted image 20260727223639.png]]


整个过程的调用是自main->router->controller->service->dao的
但是更像是一层包着一层的嵌套


缺点：几乎所有数据都是通过body传入，没有通过url
以上文档几乎全为手搓，如果碰到搞笑的说法错误请见谅

