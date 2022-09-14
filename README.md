# ShoppingList: 简陋的Go webserver API项目

简陋的ShoppingList webserver：

* 仅有api，无页面，通过HTTP request进行交互；
* 实现了新增、读取和删除功能；
* 无数据库，数据暂存在变量中。

依赖`github.com/google/uuid` 和 `github.com/gorilla/mux`

FileTree

```tree /F >filetree.txt
...\GOWORKSPACE\SRC\GO-SHOPPINGLIST
│  go.mod
│  go.sum
│  main.go
│  README.md
│  shopping.rest
│  
└─api
        server.go
        
```

编写整个项目的过程与简陋的[go-server项目](https://github.com/Colt-M1873/go-server)类似

## api

此项目没有具体html的页面，仅实现了api，使用了一个单独的api文件夹，其中server.go文件存放着server的具体实现，比[go-server项目](https://github.com/Colt-M1873/go-server)中一切以函数形式写在main里要更贴近实践。

## Server类

使用了面向对象的思想，server.go中的内容围绕`Server`类来进行，实质上是对`Server`结构体定义了一些方法。

关于Go的不完整的面向对象特性：[Is Go an object-oriented language?](https://go.dev/doc/faq#Is_Go_an_object-oriented_language)

## mux

mux是 HTTP request multiplexer 的简写，意为HTTP请求多路复用器

此项目用性能更强的[`github.com/gorilla/mux`](https://github.com/gorilla/mux)代替了原生的`net/http`来完成路由功能

## 数据存储功能雏形

具有数据存储功能，尽管数据只是以list的形式暂存在`Server`类本身的`shoppingItems`变量里，但这样的实现是一个有益的雏形，只要将与`shoppingItems`交互数据的代码拓展成与SQL交互的代码，即可在项目中加入数据库

### json struct tags

在api斤进行交互的时候，输入与输出的都是application/json形式的数据

经过json struct tags的encoding才能正常输出为json形式的数据

关于json struct tags，参考： 

> [How To Use Struct Tags in Go](https://www.digitalocean.com/community/tutorials/how-to-use-struct-tags-in-go)
>
> [golang struct json tag的使用及深入理解](https://blog.csdn.net/xz_studying/article/details/106012535)

### uuid

在为shoppinglist中的item进行编码的时候，此项目使用了[`github.com/google/uuid`](https://github.com/google/uuid)库

Google的uuid能自动分配不重复的id，相较于如MySQL数据库自带的自增主键ID，有着适合大规模分布式数据（多个MySQL表中相同的ID容易冲突，而uuid不会），便于迁移（同上，不会冲突）和保护信息（不能通过uuid判定当前信息的位置和数据库的规模）的优势

## API Tesing

使用了VScode插件 REST client来进行HTTP请求的调试，非常方便

shopping.rest文件中存放着进行调试的请求代码

### javascript发送请求

如果不使用调试工具，可以浏览器F12直接用js代码发请求

向`http://localhost:8080/shopping-items` POST一个json格式的请求

```js
fetch(new Request('http://localhost:8080/shopping-items',{
    method:'POST', 
    headers: {'Content-Type': 'application/json'},
    body:JSON.stringify({"name":"Pasta"})
})).then((resp)=>{console.log(resp)})
```

向`http://localhost:8080/shopping-items` 发起一个GET请求

```js
var url = "http://localhost:8080/shopping-items";
var xhr = new XMLHttpRequest();
xhr.open("GET", url, true);
xhr.onload = function (e) {
  if (xhr.readyState === 4) {
    if (xhr.status === 200) {
      console.log(xhr.responseText);
    } else {
      console.error(xhr.statusText);
    }
  }
};
xhr.onerror = function (e) {
  console.error(xhr.statusText);
};
xhr.send(null);
```

### 关于http请求：

| 序号 | 方法    | 描述                                                         |
| ---- | ------- | ------------------------------------------------------------ |
| 01   | HEAD    | 请求指定的页面信息，并返回实体主体。                         |
| 02   | GET     | 类似于 `GET` 请求，只不过返回的响应中没有具体的内容，用于获取报头。 |
| 03   | POST    | 向指定资源提交数据进行处理请求（例如提交表单或者上传文件）。数据被包含在请求体中。`POST` 请求可能会导致新的资源的建立或已有资源的修改。 |
| 04   | PUT     | 从客户端向服务器传送的数据取代指定的文档的内容。             |
| 05   | DELETE  | 请求服务器删除指定的页面。                                   |
| 06   | CONNECT | HTTP/1.1 协议中预留给能够将连接改为管道方式的代理服务器。    |
| 07   | OPTIONS | 允许客户端查看服务器的性能。                                 |
| 08   | TRACE   | 回显服务器收到的请求，主要用于测试或诊断。                   |
| 09   | PATCH   | 是对 PUT 方法的补充，用来对已知资源进行局部更新。            |



## 参考

[Your first HTTP Server in Go - Go Web Basics #1](https://www.youtube.com/watch?v=5BIylxkudaE&ab_channel=GoWebExamples)

上个项目: [Go simple web server](https://github.com/Colt-M1873/go-server)

[Is Go an object-oriented language?](https://go.dev/doc/faq#Is_Go_an_object-oriented_language)

[github.com/gorilla/mux](https://github.com/gorilla/mux)

[How To Use Struct Tags in Go](https://www.digitalocean.com/community/tutorials/how-to-use-struct-tags-in-go)

[golang struct json tag的使用及深入理解](https://blog.csdn.net/xz_studying/article/details/106012535)

[github.com/google/uuid](https://github.com/google/uuid)

[都2020年了，还理不清GET请求和POST请求区别？](https://juejin.cn/post/6844904097091420174)

