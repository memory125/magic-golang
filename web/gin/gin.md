# Gin

## 介绍
- Gin是一个golang的微框架，封装比较优雅，API友好，源码注释比较明确，具有快速灵活，容错方便等特点
- 对于golang而言，web框架的依赖要远比Python，Java之类的要小。自身的net/http足够简单，性能也非常不错
- 借助框架开发，不仅可以省去很多常用的封装带来的时间，也有助于团队的编码风格和形成规范

## 安装
- 首先需要安装Go（需要1.10+版本），然后可以使用下面的Go命令安装Gin。
`go get -u github.com/gin-gonic/gin`
- 将其导入您的代码中：
`import "github.com/gin-gonic/gin"`
- (可选)导入```net/http```。例如，如果使用常量，则需要这样做http.StatusOK。
`import "net/http"`
  
## 基本路由
- gin 框架中采用的路由库是基于```httprouter```做的
- 地址为：`https://github.com/julienschmidt/httprouter`

## Restful API
即Representational State Transfer的缩写。直接翻译的意思是"表现层状态转化"，是一种互联网应用程序的API设计理念：URL定位资源，用HTTP描述操作
- 获取文章 /blog/getXxx Get blog/Xxx
- 添加 /blog/addXxx POST blog/Xxx
- 修改 /blog/updateXxx PUT blog/Xxx
- 删除 /blog/delXxxx DELETE blog/Xxx

## API参数
- 可以通过Context的Param方法来获取API参数

## URL参数
- URL参数可以通过DefaultQuery()或Query()方法获取
- DefaultQuery()若参数不存在，返回默认值，Query()若不存在，返回空串

## 表单参数
- 表单传输为post请求，http常见的传输格式为四种： 
  - application/json
  - application/x-www-form-urlencoded
  - application/xml
  - multipart/form-data
- 表单参数可以通过PostForm()方法获取，该方法默认解析的是x-www-form-urlencoded或from-data格式的参数

## 文件上传
- `multipart/form-data`格式用于文件上传
- gin文件上传与原生的`net/http`方法类似，不同在于gin把原生的request封装到`c.Request`中

## routers group(路由组)
- routes group是为了管理一些相同的URL

## gin数据解析和绑定
- Json 数据解析和绑定
  - 客户端传参，后端接收并解析到结构体  
- 表单数据解析和绑定
- URI数据解析和绑定

## gin渲染
- 各种数据格式的响应
  - Json、结构体、XML、YAML类似于java的properties、ProtoBuf  
- HTML模板渲染
  - gin支持加载HTML模板, 然后根据模板参数进行配置并返回相应的数据，本质上就是字符串替换
  - `LoadHTMLGlob()`方法可以加载模板文件
- 重定向 - `gin.context.redirect()`
- 同步异步
  - `goroutine`机制可以方便地实现异步处理
  - 在启动新的`goroutine`时，不应该使用原始上下文，必须使用它的只读副本
  
## gin中间件
- 全局中间件
  - 所有请求都经过此中间件
- Next()方法
- 局部中间件

## 会话控制
- Cookie介绍
  - HTTP是无状态协议，服务器不能记录浏览器的访问状态，也就是说服务器不能区分两次请求是否由同一个客户端发出
  - Cookie就是解决HTTP协议无状态的方案之一，中文是小甜饼的意思
  - Cookie实际上就是服务器保存在浏览器上的一段信息。浏览器有了Cookie之后，每次向服务器发送请求时都会同时将该信息发送给服务器，服务器收到请求后，就可以根据该信息处理请求
  - Cookie由服务器创建，并发送给浏览器，最终由浏览器保存
  - Cookie的用途 
    - 测试服务端发送Cookie给客户端，客户端请求时携带Cookie
- Cookie的使用
  - 测试服务端发送Cookie给客户端，客户端请求时携带Cookie
- Cookie练习
  - 模拟实现权限验证中间件
    - 有2个路由，login和home
    - login用于设置Cookie
    - home是访问查看信息的请求
    - 在请求home之前，先跑中间件代码，检验是否存在Cookie
  - 访问home，会显示错误，因为权限校验未通过
- Cookie的缺点
  - 不安全，明文
  - 增加带宽消耗
  - 可以被禁用
  - Cookie有上限
- Sessions
  - gorilla/sessions为自定义session后端提供cookie和文件系统session以及基础结构。
  - 主要功能是： 
    - 简单的API：将其用作设置签名（以及可选的加密）cookie的简便方法。
    - 内置的后端可将session存储在cookie或文件系统中。
    - Flash消息：一直持续读取的session值。
    - 切换session持久性（又称“记住我”）和设置其他属性的便捷方法。
    - 旋转身份验证和加密密钥的机制。
    - 每个请求有多个session，即使使用不同的后端也是如此。
    - 自定义session后端的接口和基础结构：可以使用通用API检索并批量保存来自不同商店的session。



