package tools

import (
	"fmt"

	"gorm.io/gorm"
)

func MysqlTestdb(db *gorm.DB) {

	insertProduct := &Product{Code: "D42", Price: 100}

	db.Create(insertProduct)
	fmt.Printf("insert ID: %d, Code: %s, Price: %d\n",
		insertProduct.ID, insertProduct.Code, insertProduct.Price)

	readProduct := &Product{}
	db.First(&readProduct, "code = ?", "D42") // find product with code D42

	fmt.Printf("read ID: %d, Code: %s, Price: %d\n",
		readProduct.ID, readProduct.Code, readProduct.Price)

	// 插入数据
	user := User{Name: "Alice"}
	db.Create(&user)
	// 查找数据 id=1 1条数据
	var foundUser User
	db.First(&foundUser, 1) // 查找 ID 为 1 的用户
	fmt.Println("查找的用户:", foundUser)

	//更新 age字段
	db.Model(&foundUser).Update("Age", 30)
	// 删除数据 根据数据来的
	db.Delete(&foundUser)
	// 查询 Name 为 'sxx' 的记录
	var username User
	db.Where("Name = ?", "sxx").First(&username)

	//查询 Name 为 'sxx' 输出查询结果
	if username.ID != 0 {
		fmt.Printf("找到用户: %v\n", username)
	} else {
		fmt.Println("没有找到 Name 为 'sxx' 的用户")
	}
	//查询到第一条数据
	var user1 User
	if err := db.First(&user1).Error; err != nil {
		fmt.Println("查询失败:", err)
	} else {
		fmt.Printf("找到的用户: %+v\n", user1)
	}

	// 查询所有用户数据
	var users []User
	if err := db.Find(&users).Error; err != nil {
		fmt.Println("查询失败:", err)
	} else {
		// 输出所有查询到的用户
		fmt.Printf("所有用户: %+v\n", users)
	}
	//删除所有用户数据

	for _, v := range users {

		db.Delete(&v)
	}
}
