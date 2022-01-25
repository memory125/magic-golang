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
- DefaultQuery()若参数不村则，返回默认值，Query()若不存在，返回空串
