### 1.目录结构
```
    ├─grpc  （gRPC代码）
    │      service.go
    │      
    ├─hertz （hertz代码）
    │      hertz.go
    │      
    ├─pb    （protobuf文件）
    │      service.pb.go
    │      service.proto
    │      service_grpc.pb.go
    │  README.md
```

### 1.启动gRPC

``` 
    go run grpc/service.go
```

### 2.启动hertz
``` 
    go run hertz/hertz.go
```

### 3.浏览器访问 
#### http://localhost:8888/userinfo/1
![task2_resultimg1.png](..%2Fimg%2Ftask2_resultimg1.png)
#### http://localhost:8888/userinfo/2
![task2_resultimg2.png](..%2Fimg%2Ftask2_resultimg2.png)
#### http://localhost:8888/userinfo/3
![task2_resultimg3.png](..%2Fimg%2Ftask2_resultimg3.png)
#### http://localhost:8888/userinfo/4
![task2_resultimg4.png](..%2Fimg%2Ftask2_resultimg4.png)

