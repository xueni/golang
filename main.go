package main

import (
	"context"
	"fmt"
	"golang/gorm_use"
	"time"

	//"golang/gorm_use/models"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

type test struct {
	sli []string
}

func main() {
	db, err = gorm_use.ConnectDB()
	if err != nil {
		fmt.Println("connect db", err)
		return
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	db = db.WithContext(ctx)

	t := test{sli: []string{"y", "a", "n", "g"}}
	t1 := test{sli: t.sli}

	fmt.Println(t.sli, &t.sli[0])
	fmt.Println(t1.sli, &t1.sli[0])

	t1.sli[0] = "xiong"

	fmt.Println(t.sli, &t.sli[0])
	fmt.Println(t1.sli, &t1.sli[0])

	//var result []models.Result
	//data.JoinTableFind(db, &result)
	//fmt.Println(result)

	//db.Migrator().AutoMigrate(&models.Email{})

	//var users []models.User
	//err = data.DistinctFind(db, &users)
	//fmt.Println(users)

	//var users []models.User
	//user.Find(db, &users, []uint{1,2,3})
	//fmt.Println(users)

	//data := map[string]interface{}{}
	//user.FirstCond(db, user, 2)
	//fmt.Println(user)
	//
	//var users []models.User
	//user.FirstCond(db, users, []int{1,3})
	//fmt.Println(users)

	//user.First(db)
	//fmt.Println(*user)

	//if user.HasTable(db) {
	//	user.DropTable(db)
	//}
	//
	//err = user.CreateTable(db)
	//if err != nil {
	//	fmt.Println("Create Table",err)
	//	return
	//}

	//user.ModelBatchInsert(db)
	//db, err = user.OmitInsert(db)
	//if err != nil {
	//	fmt.Println("Omit Insert db",err)
	//	return
	//}
	//
	//db, err = user.Insert(db)
	//if err != nil {
	//	fmt.Println("Insert db",err)
	//	return
	//}
	//
	//db, err = user.Insert(db)
	//if err != nil {
	//	fmt.Println("Insert db",err)
	//	return
	//}
	//
	//db, err = user.SelectInsert(db)
	//if err != nil {
	//	fmt.Println("Select Insert db",err)
	//	return
	//}
}
