# airad
Beego based AirAd API service

which supports mobile app as cloud service
![](https://github.com/rubinliudongpo/airad/blob/master/pictures/airad.png)

## 特性

- Air Quality API 
- Access Token, User Auth

## 依赖

- [github.com/astaxie/beego](https://github.com/astaxie/beego)
- [github.com/astaxie/beego/context](https://github.com/astaxie/beego/context)
- [github.com/astaxie/beego/orm](https://github.com/astaxie/beego/orm)
- [github.com/dgrijalva/jwt-go](https://github.com/dgrijalva/jwt-go)
- [github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)


## 如何开始

- 安装 [bee](https://github.com/beego/bee) 工具
- go get github.com/rubinliudongpo/airad （注意配置GOROOT，GOPATH，详情请参考 http://sourabhbajaj.com/mac-setup/Go/README.html）
- 在mysql数据库里创建数据库名字叫airad,创建（并授权给）用户（gouser）和密码（gopassword）

## 注意
