# 内师助手 Engine


![language-go](https://img.shields.io/badge/language-Go-blue.svg)
[![](https://img.shields.io/badge/license-MIT-red.svg)](https://github.com/strugglerx/gpaCalculator/blob/master/LICENSE)

*代码仅供参考(2018/10/30)*

## 基于go语言重构的后端
>  因为最近大四没事干,还不能离校,所以只能不断学习咯,目前所有接口已经移植过来了,小程序的开发板也全部用到了go的接口

>  之前小程序使用的是python做后端,后来改为node,当然对于内师助手这个小程序,使用node后端应该说绰绰有余了,但是人生么,不就是一直学习么,再说go有很好的高并发,关键词又少,学学又何妨呢?

> 开源的目的是,我即将步入社会大学,人的精力总是有限的,如果有一天我没有精力维护这个项目,我期待这个时候有一个人可以挺身而出,继续维护这个项目,造福更多19级,20级,甚至xx级的学弟学妹们

## 后端技术栈

```
nginx
go
vue+vue-router
mongodb
```

## 如何启动?

##### 开发环境
```
> bee run
```
##### 生产环境
```
> bee pack
> mkdir myserver
> tar -zxvf server.tar.gz  -C myserver
> nohup ./server&
```
##### 配置说明

./conf/app.conf

```
appname = server
httpport = 8888 #端口
runmode = dev #生产环境改prod
sessionon = true #开启session 后台管理部分还没下完orz
EnableGzip = true #开启gzip

appid =             #小程序的appid 接受客服消息,解密微信运动步数的时候有用
secret =            #小程序的secret 同上

[mongodb]   
url = mongodb://user:pwd@127.0.0.1:27017/struggler  #数据库地址

[proxy]
url = http://127.0.0.1:80  #代理地址有个接口不用代理很难访问

[sport]
url = http://127.0.0.1:8080/SportScore/default.aspx 体测页面是内网,需要反向代理
```
##### 页面&接口

|主路由|二级路由|method
| :------| -----
|/|--|get
|admin|--|get post
|logout|--|get
|api|msg|get post
|--|point|get post
|--|indexswiper|get
|--|indexconfig|get
|--|jobsdetail|get
|--|jobslist|get
|--|stepUpdate|post
|--|steplist|get
|--|library|get
|--|libraryDetail|get
|--|sports|post
|--|eip|post
|--|artlist|get
|--|artdetail|get
|--|artlike|post
|manager|info|get
|--|feedback|get post
|--|artinsert|post
|--|delete|post
|--|option|post
|--|changePwd|post

##### 后台管理截图

![模块开关](./google-chrome_20181030111528.png)

![模块开关](./google-chrome_20181030104626.png)


![坐标管理](./google-chrome_20181030104635.png)

##### 怎么体验内师助手?


扫它!

![](./gg_20180603205458.jpg)

##### 联系方式

*我的公众号：[wx-struggler](https://mp.weixin.qq.com/s/KOydGJa7D3dJzl9fvOUTQg)*
*个人微信：（strongdreams）期待我们有共同语言！*



##### 更多待补充中.....


### 开源协议

[MIT Copyright (c) 2018 STRUGGLER](https://github.com/strugglerx/nsEngine/blob/master/LICENSE)
