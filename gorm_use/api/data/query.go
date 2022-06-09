package data

import (
	"golang/gorm_use"
	"golang/gorm_use/models"
	"gorm.io/gorm"
)

func First(db *gorm.DB, data interface{}) error {
	return db.First(data).Error
}

func Last(db *gorm.DB, data interface{}) (newDB *gorm.DB, err error) {
	newDB = db.Last(data)
	err = newDB.Error
	return
}

func Take(db *gorm.DB, data interface{}) error {
	return db.Take(data).Error
}

func Find(db *gorm.DB, users, conditions interface{}) error {
	return db.Find(users, conditions).Error
}

func ModelFirst(db *gorm.DB, data interface{}) error {
	return db.Model(data).First(data).Error
}

func ModelFirstMap(db *gorm.DB, model interface{}, data map[string]interface{}) error {
	return db.Model(model).First(&data).Error
}

//例子1 拿到值 user会将表结构带入
/*user := &models.User{}
err = user.TableFirst(db, &user)
fmt.Println(user)*/

//例子2  拿不到值 在根据逐渐id升序排序时拿不到表结构所以报错
//user := &models.User{}
//data := map[string]interface{}{}
//err = user.TableFirst(db, &data)
func TableFirst(db *gorm.DB, data interface{}) error {
	return db.Table(gorm_use.TBALE_USER).First(data).Error
}

func TableTake(db *gorm.DB, data interface{}) error {
	return db.Table(gorm_use.TBALE_USER).Take(data).Error
}

//例1 拿不到值 find函数里的conds参数时多参数不定参，传nil进去后会变成[nil]，数组长度为1，找不到id为nil的数据
//user := &models.User{}
//var users []models.User
//err = user.TableFind(db, &users, nil)

//例2 ok可以拿到值 数组长度为0，拿全部
//user := &models.User{}
//var users []models.User
//err = user.TableFind(db, &users, []interface{}{})

//例子3 拿到1条数据，多条嵌套map报错，不确定是不是传入方式不对
//user := &models.User{}
//data := map[string]interface{}{}
//err = user.TableFind(db, &data, []int{})
func TableFind(db *gorm.DB, users, conditions interface{}) error {
	return db.Table(gorm_use.TBALE_USER).Find(users, conditions).Error
}

//条件查询
func WhereFirst(db *gorm.DB, id uint, data interface{}, cond func(*gorm.DB, interface{}) *gorm.DB) error {
	return cond(db, id).First(data).Error
}

func CondWhereValue(db *gorm.DB, name interface{}) *gorm.DB {
	return db.Where("name=?", name)
}

func CondWhereNotValue(db *gorm.DB, name interface{}) *gorm.DB {
	return db.Where("name <> ?", name)
}

func CondWhereIn(db *gorm.DB, ins interface{}) *gorm.DB {
	return db.Where("name in?", ins)
}

func CondWhereLike(db *gorm.DB, like interface{}) *gorm.DB {
	return db.Where("name like?", like)
}

func CondWhereAnd(db *gorm.DB, values ...interface{}) *gorm.DB {
	return db.Where("name=? and value=?", values)
}

func CondWhereBetween(db *gorm.DB, values ...interface{}) *gorm.DB {
	return db.Where("id between ? and ?", values)
}

//选定查询 未返回数量
func SelectValueFind(db *gorm.DB, data interface{}, selectFunc func(*gorm.DB, interface{}) *gorm.DB, values ...interface{}) error {
	return selectFunc(db, values).Find(data).Error
}

func SelectValue(db *gorm.DB, values interface{}) *gorm.DB {
	return db.Select(values)
}

func SelectCondFind(db *gorm.DB, data interface{}, selectFunc func(*gorm.DB, string, interface{}) *gorm.DB, selectCond string, values ...interface{}) error {
	return selectFunc(db, selectCond, values).Find(data).Error
}

func SelectCond(db *gorm.DB, cond string, values interface{}) *gorm.DB {
	return db.Select(cond, values)
}

//指定从数据库检索时的排序方式
func OrderFind(db *gorm.DB, data interface{}, cond string, order func(db *gorm.DB, cond string)) error {
	return Order(db, cond).Find(data).Error
}

func Order(db *gorm.DB, cond string) *gorm.DB {
	return db.Order(cond)
}

func OrderMulti(db *gorm.DB, cond []string) (newDB *gorm.DB) {
	newDB = db
	for index := 0; index < len(cond); index++ {
		newDB = newDB.Order(cond)
	}

	return
}

//limit & offset 一般用来分页
//limit 指定获取记录的最大数量 offset 指定记录开始之前的要跳过的记录数
func LimitFind(db *gorm.DB, num int, data interface{}) error {
	return Limit(db, num).Find(data).Error
}

func Limit(db *gorm.DB, num int) *gorm.DB {
	return db.Limit(num)
}

//limit(-1)消除limit条件
//
func LimitMultiFind(db *gorm.DB, num int, data1, data2 interface{}) error {
	return Limit(db, num).Find(data1).Limit(-1).Find(data2).Error
}

//Count获取查询的数据条数
func OffsetFind(db *gorm.DB, offset int, data interface{}) (count int64, err error) {
	err = Offset(db, offset).Find(data).Count(&count).Error
	return
}

func Offset(db *gorm.DB, offset int) *gorm.DB {
	return db.Offset(offset)
}

func PageFind(db *gorm.DB, page *models.Page, data interface{}) error {
	return db.Limit(page.Limit).Offset(page.Offset).Find(data).Error
}

// SELECT * FROM users OFFSET 10; (users1)
// SELECT * FROM users; (users2)
func PageFindMulti(db *gorm.DB, page *models.Page, data1, data2 interface{}) error {
	return db.Limit(page.Limit).Offset(page.Offset).Find(data1).Limit(-1).Offset(-1).Find(data2).Error
}

//例子
//group having

//SELECT name, sum(id) as total FROM `users` WHERE name LIKE "y%" GROUP BY `name`
func GroupFind(db *gorm.DB, result *models.Result) error {
	err := db.Model(&models.User{}).Select("name, sum(id) as total").Where("name like ?", "y%").Group("name").Find(result).Error
	return err
}

func DistinctFind(db *gorm.DB, data interface{}) error {
	return db.Distinct("name", "value").Order("name desc").Find(data).Error
}

//type result struct {
//  Name  string
//  Email string
//}
//data为result类型指针
func JoinModelFind(db *gorm.DB, data interface{}) error {
	return db.Model(&models.User{}).Select("user.name, emails.email").Joins("left join emails on emails.user_id=user.id").Scan(data).Error
}

func JoinTableFind(db *gorm.DB, data interface{}) error {
	return db.Table("user").Select("user.name, emails.email").Joins("join emails on emails.user_id=user.id").Scan(data).Error
}
