# octaveserver
该案例用于演示上传参数文件返回Dat文件的例子

### 说明
该服务将Matlab算法封装为服务并提供Restful-HTTP接口
1. server-上传csv文件组
2. matlab-实际调用Matlab算法
3. matlab-返回结果
4. server-JSON组装多个DAT卡片文件结果返回


### 编译

```bash
docker build -t octaveserver:v1.1  .
```
### 运行

```bash
docker rm -f octaveserver
docker run --restart=always -itd \
-p 8630:8630 --name octaveserver octaveserver:v1.1
```
### 测试
（1）查看版本号
```bash
 curl -v http://localhost:8630/version

* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Content-Type: application/json; charset=utf-8
< Date: Fri, 23 Dec 2022 07:02:22 GMT
< Content-Length: 44
< 
* Connection #0 to host localhost left intact
{
    "version":"v1.1.1",
    "updateat":"2022-12-23"
}
```

（2）调用Dat转换算法
```bash
cd tests

curl -X POST -v --form "inputs=@./matrix.csv" --form "inputs=@./EffectiveT.csv" http://localhost:8630/convert

返回结果：
{   "code":0,
    "msg":"Success",
    "datfiles":
    [
        {"datname":"JCNINP.DAT",
         "datcontent":"JCNOPT IS  MN                     C  NID             FLFL    SMPT    S     22.01\nRELIEF\nEND"}
    ]
}

```

### 分析镜像

```bash 
dive octaveserver:v1.1
```

### 查看容器日志

```bash
docker logs -f octaveserver
```