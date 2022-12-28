# octaveserver
该案例用于演示上传参数文件返回Dat文件的例子

### 说明
该服务将Matlab算法封装为服务并提供Restful-HTTP接口
1. server-上传CSV文件，控制输出文件。
2. matlab-实际调用Matlab算法
3. matlab-返回结果
4. server-JSON组装Dat结果返回

![图片](images/octave.png)

### 编译

```bash
docker build -t octaveserver:v1.0  .
```
### 运行

```bash
docker run --restart=always -itd \
-p 8630:8630 --name octaveserver octaveserver:v1.0
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
    "version":"v1.0.1",
    "updateat":"2022-12-23"
}
```

（2）调用求和算法
```bash
cd tests

curl -X POST -v --form type=1 --form "matrix=@./matrix.csv" --form "effective=@./EffectiveT.csv" http://localhost:8630/convert

返回结果：
{
    "code":0,
    "msg":"",
    "datname":"JCNINP.DAT",
    "datcontent":"JCNOPT IS  MN                     C  NID             FLFL    SMPT    S     1.75\nRELIEF\nEND"
}
```

### 分析镜像

```bash 
dive octaveserver:v1.0
```

### 查看容器日志

```bash
docker logs -f octaveserver
```