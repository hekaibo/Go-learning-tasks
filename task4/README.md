### 1.目录结构
``` 
    ├─gorm  
    │  │  init.go
    │  │  
    │  ├─dbop
    │  │      db.go
    │  │      
    │  └─model
    │          user.go
    │          
    ├─grpc
    │      service.go
    │      
    ├─hertz
    │  │  .gitignore
    │  │  .hz
    │  │  main.go
    │  │  router.go
    │  │  router_gen.go
    │  │  
    │  └─biz
    │      ├─handler
    │      │      user_api.go
    │      │      
    │      └─router
    │              register.go
    │              
    ├─pb
    │       service.pb.go
    │       service.proto
    │       service_grpc.pb.go
    │  
    |  .env
    │  README.md
    │ 
```

### 2.启动gRPC
``` 
    go run grpc/service.go
```

### 3.启动hertz
``` 
    go run ./hertz
```

### 接口
``` 
    对外提供5个接口
    (Get)createuser/username/email       新增用户接口
    (Get)getuserbyname/username          根据姓名查询用户信息
    (Get)getalluser                      查询所有用户信息
    (Get)updateuserbyname/username/email 更新用户邮箱
    (Get)deleteuserbyname/username       根据用户名删除用户
```
### 4.1新增用户
#### 访问 http://localhost:8888/createuser/zhangsan/zhangsan@email.com
![task4_resultimg1.png](..%2Fimg%2Ftask4_resultimg1.png)
#### 访问 http://localhost:8888/createuser/lisi/lisi@email.com
![task4_resultimg2.png](..%2Fimg%2Ftask4_resultimg2.png)

![task4_resultimg3.png](..%2Fimg%2Ftask4_resultimg3.png)
### 4.2查询用户
#### 查询单个用户 http://localhost:8888/getuserbyname/zhangsan
![task4_resultimg4.png](..%2Fimg%2Ftask4_resultimg4.png)
#### 查询所有用户 http://localhost:8888/getalluser
![task4_resultimg5.png](..%2Fimg%2Ftask4_resultimg5.png)
### 4.3更新用户
#### 访问 http://localhost:8888/updateuserbyname/zhangsan/zs@email.com
![task4_resultimg6.png](..%2Fimg%2Ftask4_resultimg6.png)
![task4_resultimg7.png](..%2Fimg%2Ftask4_resultimg7.png)
### 4.4删除用户
#### 访问 http://localhost:8888/deleteuserbyname/zhangsan
![task4_resultimg8.png](..%2Fimg%2Ftask4_resultimg8.png)
![task4_resultimg9.png](..%2Fimg%2Ftask4_resultimg9.png)


