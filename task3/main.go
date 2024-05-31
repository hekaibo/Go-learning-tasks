package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"mission/task3/dbop"
	"mission/task3/model"
	"os"
)

func printGamer(gamers []*model.Gamer) {
	for _, g := range gamers {
		fmt.Println("name:", g.Name, "email", g.Email)
	}
}

func main() {
	err := godotenv.Load()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_DATABASE"),
	)
	DB, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	// 迁移
	//DB.AutoMigrate(&model.Gamer{})

	fmt.Println("****************************")
	// 创建用户
	dbop.CreateGamer(DB, &model.Gamer{Name: "zhangsan", Email: "zhangsan@gmail.com"})
	dbop.CreateGamer(DB, &model.Gamer{Name: "lisi", Email: "lisi@gmail.com"})
	fmt.Println("新增用户zhangsan、lisi")
	// 读取所有用户
	fmt.Println("查询所有用户")
	gamer, err := dbop.GetAllGamer(DB)
	if err != nil {
		panic(err)
	}
	printGamer(gamer)
	fmt.Println("****************************")

	// 删除用户
	fmt.Println("删除用户lisi")
	dbop.DeleteGamer(DB, "lisi")

	fmt.Println("查询所有用户")
	gamer, err = dbop.GetAllGamer(DB)
	if err != nil {
		panic(err)
	}
	printGamer(gamer)
	fmt.Println("****************************")

	// 新增用户
	fmt.Println("新增用户wangwu")
	dbop.CreateGamer(DB, &model.Gamer{Name: "wangwu", Email: "wangwu@gmail.com"})
	fmt.Println("查询所有用户")
	gamer, err = dbop.GetAllGamer(DB)
	if err != nil {
		panic(err)
	}
	printGamer(gamer)
	fmt.Println("****************************")

	//更新用户
	fmt.Println("修改zhangsan用户邮箱为zs@gmail.com")
	dbop.UpdateGamer(DB, "zhangsan", "zs@gmail.com")

	// 查询zhangsan信息
	fmt.Println("查询zhangsan用户信息")
	game, _ := dbop.GetGamerByName(DB, "zhangsan")
	fmt.Println("name:", game.Name, "email", game.Email)
	fmt.Println("****************************")

}
