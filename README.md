

1.安装golang，参看：[安装golang](http://www.fancyecommerce.com/2017/12/28/centos6-%e5%ae%89%e8%a3%85-golang-1-9/)

2.安装gin框架

```
go get github.com/gin-gonic/gin
```


详细参看:[安装gin框架](http://www.fancyecommerce.com/2017/12/28/centos6-%e5%ae%89%e8%a3%85go%e6%a1%86%e6%9e%b6gin%e7%9a%84%e6%ad%a5%e9%aa%a4%ef%bc%8c%e4%bb%a5%e5%8f%8a%e4%b8%ad%e9%97%b4%e9%81%87%e5%88%b0%e7%9a%84%e5%9d%91/)

3.下载：



```
cd $GOPATH
git clone https://github.com/fancyecommerce/go_test.git src
```


`vim src/main/customer.go`

修改代码,修改ip：
```
r.Run("120.24.37.249:3000")
``` 

4.运行

4.1 直接运行main包: `go run src/main/customer.go`

可以访问 120.24.37.249:3000/ping

4.2 安装main包: `go install  main`
，然后进入 `$GOPATH/bin`  ,执行  `./main`

可以访问 120.24.37.249:3000/ping
















