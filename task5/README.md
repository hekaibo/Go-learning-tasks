### 1.目录结构
```

    ├─dao
    │  ├─ init.go
    │  │  
    │  ├─dbop
    │  │   ├─ dbop_notice.go
    │  │   └─ dbop_user.go
    │  │      
    │  └─model
    │      ├─ notice.go
    │      └─ user.go
    │          
    ├─grpc
    │  ├─ main.go
    │  │  
    │  └─service
    │      ├─ notice.go
    │      └─ user.go
    │          
    ├─hertz
    │  ├─ .gitignore
    │  ├─ .hz
    │  ├─ main.go
    │  ├─ router.go
    │  ├─ router_gen.go
    │  │  
    │  └─biz
    │      ├─handler
    │      │   ├─ notice_api.go
    │      │   └─ user_api.go
    │      │      
    │      └─router
    │          └─ register.go
    │              
    └─pb
    │  ├─ notice.pb.go
    │  ├─ notice.proto
    │  ├─ notice_grpc.pb.go
    │  ├─ user.pb.go
    │  ├─ user.proto
    │  └─ user_grpc.pb.go
    │
    ├─  .env
    └─  README.md
```

### 2.启动gRPC
``` 
    go run grpc/main.go
```

### 3.启动hertz
``` 
    go run ./hertz
```
### 4.数据库设计
|   字段   |    类型    |  说明  |
|:------:|:--------:|:----:|
|   id   |  bigint  |  主键  |
| created_at | datetime | 创建时间 |
| updated_at | datetime | 更新时间 |
| deleted_at | datetime | 删除时间 |
| title | longtext | 公告标题 |
| content | longtext | 公告内容 |
| publish_user | longtext | 发布者  |

### 5.接口
``` 
    对外提供5个接口
    (POST)createnotice      新增游戏公告
    (Get)shownotice/id      根据id查询公告详细内容
    (Get)showallnotice      查询所有游戏公告
    (POST)updatenotice      根据id更新公告
    (Get)deletenotice/id    根据id删除公告
```

### 5.1新增公告
#### 利用postman访问 http://localhost:8888/createnotice/ 创建两个公告
![task5_resultimg1.png](..%2Fimg%2Ftask5_resultimg1.png)
![task5_resultimg2.png](..%2Fimg%2Ftask5_resultimg2.png)
### 5.2查询所有公告
#### 访问 http://localhost:8888/showallnotice
![task5_resultimg3.png](..%2Fimg%2Ftask5_resultimg3.png)
### 5.3查看公告详情
#### 访问 http://localhost:8888/shownotice/1 
![task5_resultimg4.png](..%2Fimg%2Ftask5_resultimg4.png)
#### 访问 http://localhost:8888/shownotice/2
![task5_resultimg5.png](..%2Fimg%2Ftask5_resultimg5.png)
#### 访问 http://localhost:8888/shownotice/3
![task5_resultimg6.png](..%2Fimg%2Ftask5_resultimg6.png)
### 5.4更新公告
#### 访问 http://localhost:8888/updatenotice/
![task5_resultimg7.png](..%2Fimg%2Ftask5_resultimg7.png)
#### 查看修改详情，id为1的公告修改成功
![task5_resultimg8.png](..%2Fimg%2Ftask5_resultimg8.png)
### 5.5删除公告
#### 访问 http://localhost:8888/deletenotice/1
![task5_resultimg9.png](..%2Fimg%2Ftask5_resultimg9.png)
#### 访问所有公告，id为1的公告已删除
![task5_resultimg10.png](..%2Fimg%2Ftask5_resultimg10.png)






