
准备工作
-----------

1.安装golang，参看：[安装golang](http://www.fancyecommerce.com/2017/12/28/centos6-%e5%ae%89%e8%a3%85-golang-1-9/)

2.安装gin框架

```
go get github.com/gin-gonic/gin
```

详细参看:[安装gin框架](http://www.fancyecommerce.com/2017/12/28/centos6-%e5%ae%89%e8%a3%85go%e6%a1%86%e6%9e%b6gin%e7%9a%84%e6%ad%a5%e9%aa%a4%ef%bc%8c%e4%bb%a5%e5%8f%8a%e4%b8%ad%e9%97%b4%e9%81%87%e5%88%b0%e7%9a%84%e5%9d%91/)

2.2 安装go mysql 包

```
go get github.com/go-sql-driver/mysql
```

详细参看:[安装go mysql](https://github.com/go-sql-driver/mysql/)


安装示例包
-------


3.下载当前示例：

```
// $GOPATH：按照上面的安装，就是/root/go
cd $GOPATH  
git clone https://github.com/fancyecommerce/go_test.git src
```

3.2修改配置

`vim src/main/customer.go`

修改代码,修改ip：(将这个ip改成您自己的ip)
```
r.Run("120.24.37.249:3000")
``` 

3.3 mysql配置

进入文件夹：`src/fecshop.com/common/cdb`， 将文件名
`mysql_example.go` 改成  `mysql.go`
,打开`mysql.go` ，去掉注释，添加你的mysql信息

3.4 在3.3步骤配置对应的数据库中添加数据表：

```
CREATE TABLE IF NOT EXISTS `user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(20) DEFAULT '',
  `age` int(11) DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=11 ;

--
-- 转存表中的数据 `user`
--

INSERT INTO `user` (`id`, `name`, `age`) VALUES
(1, '111', 111),
(2, 'terry', 44),
(3, 'terry', 55),
(4, 'terry', 44),
(5, 'terry', 44),
(6, 'terry', 44),
(7, 'terry', 44),
(8, 'terry', 44),
(10, 'terry', 66);
```


4.运行

4.1 直接运行main包: `go run src/main/customer.go`



4.2 安装main包: `go install  main`
，然后进入 `$GOPATH/bin`  ,执行  `./main`

5.访问

用google浏览器的postman来访问api测试：

5.1 查询

GET: http://120.24.37.249:3000/v1/users

5.2 新增

POST:http://120.24.37.249:3000/v1/users

body json:
```
{
	"name":"terry",
	"age":19
}
```

5.3 更新id为8的数据行

PATCH:http://120.24.37.249:3000/v1/users/8

body json:
```
{
	"name":"terry",
	"age":119
}
```
5.4 更新id为5的数据行

DELETE:http://120.24.37.249:3000/v1/users/5

6.具体可以查看 `src/main/customer.go`的`main`方法













