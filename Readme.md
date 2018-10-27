# 内师助手 Engine

![language-python](https://img.shields.io/badge/language-Go-blue.svg)
[![](https://img.shields.io/badge/license-MIT-red.svg)](https://github.com/strugglerx/gpaCalculator/blob/master/LICENSE)

*代码仅供参考*

## 基于go语言重构的后端
>  因为最近大四没事干,还不能离校,所以只能不断学习咯,这个还有一些就口没有写完,并不是最终版本...

>  之前小程序使用的是python做后端,后来改为node,当然对于内师助手这个小程序,使用node后端应该说绰绰有余了,但是人生么,不就是一直学习么,再说go有很好的高并发,关键词又少,学学又何妨呢?

## 如何启动?

##### 开发环境
```
> bee run
```
##### 生产环境
如果你嫌配置环境麻烦的话,我已经帮你打包好了,x64Linux的环境下亲测可以使用!
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
```

##### 怎么体验内师助手?


扫它!

![](./gg_20180603205458.jpg)

##### 联系方式

*我的公众号：[wx-struggler](https://mp.weixin.qq.com/s/KOydGJa7D3dJzl9fvOUTQg)*
*个人微信：（strongdreams）期待我们有共同语言！*



##### 更多待补充中.....


### 开源协议

[MIT Copyright (c) 2018 STRUGGLER](https://github.com/strugglerx/nsEngine/LICENSE)
