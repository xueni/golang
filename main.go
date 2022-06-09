package main

import (
	"fmt"
	"golang/gorm_use"
	"golang/gorm_use/models"

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

	user := models.User{}
	stmt := db.Session(&gorm.Session{DryRun: true}).First(&user, 1).Statement
	fmt.Println(stmt.SQL.String())
	fmt.Println(stmt.Vars)

	db.Dialector.Explain(stmt.SQL.String(), stmt.Vars...)
	fmt.Println("test")
	//ctx, _ := context.WithTimeout(context.Background(), time.Second)
	//db = db.WithContext(ctx)

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
