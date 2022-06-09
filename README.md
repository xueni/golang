"# golang" 


https://gorm.io/zh_CN/docs/connecting_to_the_database.html  官方文档英文
https://learnku.com/docs/gorm/v2/query/9733 官方文档中文

链式方法
链式方法是将 Clauses 修改或添加到当前 Statement 的方法，例如：
Where, Select, Omit, Joins, Scopes, Preload, Raw (Raw can’t be used with other chainable methods to build SQL)…

Finisher 方法
Finishers 是会立即执行注册回调的方法，然后生成并执行 SQL，比如这些方法：
Create, First, Find, Take, Save, Update, Delete, Scan, Row, Rows…

新建会话模式
在初始化了 *gorm.DB 或 新建会话方法 后， 调用下面的方法会创建一个新的 Statement 实例而不是使用当前的
GROM 定义了 Session、WithContext、Debug 方法做为 新建会话方法，查看会话 获取详情


协程安全
db.Where("name = ?", "jinzhu").Where("age = ?", 18).Find(&users)
1.db.Where("name = ?", "jinzhu")
    调用getinstance生成一个新的db,判断db clone或者>0，则生成一个新的db,else则返回原有db，生成的新的db的clone为0
2.Where("age = ?", 18)
    调用getinstance发现新的db.clone == 0,返回
3.Find("age = ?", 18)
    调用getinstance发现新的db.clone == 0,返回
4.整个语句返回的db不要赋值给初始db,否则会发生多个协程共用一个db数据，会出现不可预知错误，像下边这种就是错误使用
    tx := db.Where("name = ?", "jinzhu") 生成一个新的db tx
    // 不安全的复用 Statement
    for i := 0; i < 100; i++ {
      go tx.Where(...).First(&user) 每一个单独的协程都共同使用了db tx 也就是说这100个协程会公用tx,也就是所谓的资源抢占
    }

