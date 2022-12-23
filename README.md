# octaveserver

### 说明
该服务将Matlab算法封装为服务并提供Restful-HTTP接口

1. server-上传CSV文件，文件中包括N行M列个元素，同时传入控制参数，控制调用算法的类型。
2. matlab-实际调用Matlab算法(type=1 求和  type=2 求平方和)
3. matlab-返回结果
4. server-JSON组装结果返回

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
（1） 查看版本号

（2） 调用求和算法

（3） 调用求平方和算法

