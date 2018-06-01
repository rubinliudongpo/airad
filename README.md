# AirAd
Beego based RestFul API service

which supports mobile app as cloud service
![](https://github.com/rubinliudongpo/airad/blob/master/pictures/airad.png)

## 特性

- RestFul API
- Access Token, User Auth

## 依赖

- [github.com/astaxie/beego](https://github.com/astaxie/beego)
- [github.com/astaxie/beego/context](https://github.com/astaxie/beego/context)
- [github.com/astaxie/beego/orm](https://github.com/astaxie/beego/orm)
- [github.com/dgrijalva/jwt-go](https://github.com/dgrijalva/jwt-go)
- [github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)


## 如何开始

- 安装 [bee](https://github.com/beego/bee) 工具
- go get github.com/rubinliudongpo/airad （注意配置GOROOT，GOPATH，详情请参考 http://sourabhbajaj.com/mac-setup/Go/README.html ）
- 在mysql数据库里创建数据库名字叫airad 
```
   mysql -uroot -pYOURROOTPASSWORD  -h 127.0.0.1 -e "CREATE DATABASE IF NOT EXISTS airad DEFAULT CHARSET utf8 COLLATE utf8_general_ci;"
```  
- 创建（并授权给）用户（gouser）和密码（gopassword）
```
    mysql -uroot -pYOURROOTPASSWORD  -h 127.0.0.1 -e "grant all privileges on airad.* to gouser@'%' identified by 'gopass';"
```
- 导入airad.sql
```
    mysql -ugouser -pgopass  airad < database/airad.sql)
```

## 查看和调试

 请通过 http://localhost:9080/swagger/ 试用API，界面如下
![](https://github.com/rubinliudongpo/airad/blob/master/pictures/airad_swagger.png)

## 注意
