# GORM

## 参数说明
### DB
gorm.DB 实例在创建时,默认情况下会向数据库服务端发起一次连接，以保证 dsn 的正确性
### 配置
```Go
type Config struct {
    // gorm 中，针对单笔增、删、改操作默认会启用事务. 可以通过将该参数设置为 true，禁用此机制
    SkipDefaultTransaction bool
    // 表、列的命名策略
    NamingStrategy schema.Namer
    // 自定义日志模块
    Logger logger.Interface
    // 自定义获取当前时间的方法
    NowFunc func() time.Time
    // 是否启用 prepare sql 模板缓存模式
    PrepareStmt bool
    // 在 gorm 创建 db 实例时，会创建 conn 并通过 ping 方法确认 dsn 的正确性. 倘若设置此参数，则会禁用 db 初始化时的 ping 操作
    DisableAutomaticPing bool
    // 不启用迁移过程中的外联键限制
    DisableForeignKeyConstraintWhenMigrating bool
    // 是否禁用嵌套事务
    DisableNestedTransaction bool
    // 是否允许全局更新操作. 即未使用 where 条件的情况下，对整张表的字段进行更新
    AllowGlobalUpdate bool
    // 执行 sql 查询时使用全量字段
    QueryFields bool
    // 批量创建时，每个批次的数据量大小
    CreateBatchSize int
    // 条件创建器
    ClauseBuilders map[string]clause.ClauseBuilder
    // 数据库连接池
    ConnPool ConnPool
    // 数据库连接器
    Dialector
    // 插件集合
    Plugins map[string]Plugin
    // 回调钩子
    callbacks  *callbacks
    // 全局缓存数据，如 stmt、schema 等内容
    cacheStore *sync.Map
}
```
### session
```Go
// Session 配置
type Session struct {
  DryRun                   bool
  PrepareStmt              bool
  NewDB                    bool
  Initialized              bool
  SkipHooks                bool
  SkipDefaultTransaction   bool
  DisableNestedTransaction bool
  AllowGlobalUpdate        bool
  FullSaveAssociations     bool
  QueryFields              bool
  Context                  context.Context
  Logger                   logger.Interface
  NowFunc                  func() time.Time
  CreateBatchSize          int
}
```

### Statement
```Go
type Statement struct {
	*DB
	TableExpr            *clause.Expr
	Table                string
	Model                interface{}
	Unscoped             bool
	Dest                 interface{}
	ReflectValue         reflect.Value
	Clauses              map[string]clause.Clause
	BuildClauses         []string
	Distinct             bool
	Selects              []string // selected columns
	Omits                []string // omit columns
	Joins                []join
	Preloads             map[string][]interface{}
	Settings             sync.Map
	ConnPool             ConnPool
	Schema               *schema.Schema
	Context              context.Context
	RaiseErrorOnNotFound bool
	SkipHooks            bool
	SQL                  strings.Builder
	Vars                 []interface{}
	CurDestIndex         int
	attrs                []interface{}
	assigns              []interface{}
	scopes               []func(*DB) *DB
}
type Statement struct {
    // 数据库实例
    *DB
    // ...
    // 表名
    Table                string
    // 操作的模型
    Model                interface{}
    // ...
    // 处理结果反序列化到此处
    Dest                 interface{}
    // ...
    // 各种条件语句
    Clauses              map[string]clause.Clause
    
    // ...
    // 是否启用 distinct 模式
    Distinct             bool
    // select 语句
    Selects              []string // selected columns
    // omit 语句
    Omits                []string // omit columns
    // join 
    Joins                []join
    
    // ...
    // 连接池，通常情况下是 database/sql 库下的 *DB  类型.  在 prepare 模式为 gorm.PreparedStmtDB
    ConnPool             ConnPool
    // 操作表的概要信息
    Schema               *schema.Schema
    // 上下文，请求生命周期控制管理
    Context              context.Context
    // 在未查找到数据记录时，是否抛出 recordNotFound 错误
    RaiseErrorOnNotFound bool
    // ...
    // 执行的 sql，调用 state.Build 方法后，会将 sql 各部分文本依次追加到其中. 具体可见 2.5 小节
    SQL                  strings.Builder
    // 存储的变量
    Vars                 []interface{}
    // ...
}

```
### 执行器
```Go
type processor struct {
    // 从属的 DB 实例
    db        *DB
    // 拼接 sql 时的关键字顺序. 比如 query 类，固定为 SELECT,FROM,WHERE,GROUP BY, ORDER BY, LIMIT, FOR
    Clauses   []string
    // 对应于 crud 类型的执行函数链
    fns       []func(*DB)
    callbacks []*callback
}
```
### 条件clause
```GO
  	createClauses = []string{"INSERT", "VALUES", "ON CONFLICT"}
    queryClauses  = []string{"SELECT", "FROM", "WHERE", "GROUP BY", "ORDER BY", "LIMIT", "FOR"}
    updateClauses = []string{"UPDATE", "SET", "WHERE"}
    deleteClauses = []string{"DELETE", "FROM", "WHERE"}
```
一条执行 sql 中，各个部分都属于一个 clause，比如一条 SELECT * FROM reward WHERE id < 10 ORDER by id 的 SQL，其中就包含了 SELECT、FROM、WHERE 和 ORDER 四个 clause.

当使用方通过链式操作克隆 DB时，对应追加的状态信息就会生成一个新的 clause，追加到 statement 对应的 clauses 集合当中. 当请求实际执行时，会取出 clauses 集合，拼接生成完整的 sql 用于执行.

条件 clause 本身是个抽象的 interface，定义如下：


### 标签
tag 名大小写不敏感，但建议使用 camelCase 风格

|标签名	                      | 说明 |
|:---                        |:---|
|column                      |指定 db 列名|
|type	                     |列数据类型，推荐使用兼容性好的通用类型，例如：所有数据库都支持 bool、int、uint、float、string、time、bytes 并且可以和其他标签一起使用，例如：not null、size, autoIncrement… 像 varbinary(8) 这样指定数据库数据类型也是支持的。在使用指定数据库数据类型时，它需要是完整的数据库数据类型，如：MEDIUMINT UNSIGNED not NULL AUTO_INCREMENT|
|serializer                  |指定将数据序列化或反序列化到数据库中的序列化器, 例如: serializer:json/gob/unixtime|
|size	                     |定义列数据类型的大小或长度，例如 size: 256|
|primaryKey                  |将列定义为主键|
|unique	                     |将列定义为唯一键|
|default	                 |定义列的默认值|
|precision	                 |指定列的精度|
|scale	                     |指定列大小|
|not null                    |指定列为 NOT NULL|
|autoIncrement	             |指定列为自动增长|
|autoIncrementIncrement      |自动步长，控制连续记录之间的间隔
|embedded                    |嵌套字段|
|embeddedPrefix	             |嵌入字段的列名前缀|
|autoCreateTime	             |创建时追踪当前时间，对于 int 字段，它会追踪时间戳秒数，您可以使用 nano/milli 来追踪纳秒、毫秒时间戳，例如：autoCreateTime:nano|
autoUpdateTime	             |创建/更新时追踪当前时间，对于 int 字段，它会追踪时间戳秒数，您可以使用 nano/milli 来追踪纳秒、毫秒时间戳，例如：autoUpdateTime:milli|
|index	                     |根据参数创建索引，多个字段使用相同的名称则创建复合索引，查看 索引 获取详情|
|uniqueIndex	             |与 index 相同，但创建的是唯一索引|
|check	                     |创建检查约束，例如 check:age > 13，查看 约束 获取详情|
|<-	                         |设置字段写入的权限， <-:create 只创建、<-:update 只更新、<-:false 无写入权限、<- 创建和更新权限|
|->	                         |置字段读的权限，->:false 无读权限|
|-	                         |忽略该字段，- 表示无读写，-:migration 表示无迁移权限，-:all 表示无读写迁移权限|
|comment	                 |迁移时为字段添加注释|


### 时间
在设置 dsn 时，建议添加上 parseTime=true 的设置，这样能兼容支持将 mysql 中的时间解析到 golang 中的 time.Time 类型字段
但是需要注意,如果指定字段类型是time.Time,那么在创建时候,如果字段名称是CreatedAt/UpdatedAt/DeletedAt 这三个字段，那么会自动维护,要是把其他字段（如CreateTime）定义为time.Time不会填充为当前时间，而是填充为0000-00-00 00:00:00,而这个时间是不会被mysql的datetime参数接受的,同样,如果类型是int,只有CreatedAt/UpdatedAt/DeletedAt这三个字段才会自动维护

### error
```Go
var (
	// ErrRecordNotFound record not found error
	ErrRecordNotFound = logger.ErrRecordNotFound
	// ErrInvalidTransaction invalid transaction when you are trying to `Commit` or `Rollback`
	ErrInvalidTransaction = errors.New("invalid transaction")
	// ErrNotImplemented not implemented
	ErrNotImplemented = errors.New("not implemented")
	// ErrMissingWhereClause missing where clause
	ErrMissingWhereClause = errors.New("WHERE conditions required")
	// ErrUnsupportedRelation unsupported relations
	ErrUnsupportedRelation = errors.New("unsupported relations")
	// ErrPrimaryKeyRequired primary keys required
	ErrPrimaryKeyRequired = errors.New("primary key required")
	// ErrModelValueRequired model value required
	ErrModelValueRequired = errors.New("model value required")
	// ErrModelAccessibleFieldsRequired model accessible fields required
	ErrModelAccessibleFieldsRequired = errors.New("model accessible fields required")
	// ErrSubQueryRequired sub query required
	ErrSubQueryRequired = errors.New("sub query required")
	// ErrInvalidData unsupported data
	ErrInvalidData = errors.New("unsupported data")
	// ErrUnsupportedDriver unsupported driver
	ErrUnsupportedDriver = errors.New("unsupported driver")
	// ErrRegistered registered
	ErrRegistered = errors.New("registered")
	// ErrInvalidField invalid field
	ErrInvalidField = errors.New("invalid field")
	// ErrEmptySlice empty slice found
	ErrEmptySlice = errors.New("empty slice found")
	// ErrDryRunModeUnsupported dry run mode unsupported
	ErrDryRunModeUnsupported = errors.New("dry run mode unsupported")
	// ErrInvalidDB invalid db
	ErrInvalidDB = errors.New("invalid db")
	// ErrInvalidValue invalid value
	ErrInvalidValue = errors.New("invalid value, should be pointer to struct or slice")
	// ErrInvalidValueOfLength invalid values do not match length
	ErrInvalidValueOfLength = errors.New("invalid association values, length doesn't match")
	// ErrPreloadNotAllowed preload is not allowed when count is used
	ErrPreloadNotAllowed = errors.New("preload is not allowed when count is used")
)
```
## 子句构造器
gorm与子句生成器有关的类，按父级到子集排列为 
DB Statement Clause Expression
对应
数据库连接对象 语句 子句 表达式

GORM 内部使用 SQL builder 生成 SQL。对于每个操作，GORM 都会创建一个 *gorm.Statement 对象，所有的 GORM API 都是在为 statement 添加、修改 属性，执行过程中调用Expression接口的表达式生成器，生成最终的sql语句。

## 自定义数据类型
https://github.com/go-gorm/datatype

自定义的数据类型必须实现Scanner/Valuer接口,完成数据的存取和读出

	Scanner接口的Scan方法，是从数据库读取数据到Go变量时需要进行的解析处理，与解码的过程类型。

	valuer接口的Value方法，是将Go变量存到数据库时进行编码处理。
## NULL的处理
golang 中一些基础类型都存在对应的零值，即便用户未显式给字段赋值，字段在初始化时也会首先赋上零值(非NULL)
可以使用SQL.NULLXXX或者指针来处理NULL值

```Go
type User struct {
  gorm.Model
  Name string
  Age  *int           `gorm:"default:18"`
  Active sql.NullBool 
}
```	
其中的Null都为结构体,通过Valid字段进行标识持久化的字段是否为NULL,如果Valid true,则XXX字段是没有意义的,数据库中保存的字段是NULL

```Go
type NullXXX struct {
	XXX XXX
	Valid  bool // Valid is true if String is not NULL
}

```
## 创建

### 指定字段创建
可以通过DB.select()指定字段创建记录:
```Go
db.Select("Name", "Age", "CreatedAt").Create(&user)
// INSERT INTO `users` (`name`,`age`,`created_at`) VALUES ("jinzhu", 18, "2020-07-04 11:05:21.775")
```
没有指定的字段会被忽略
可以通过DB.omit()指定字段创建记录,omit指定的字段则会被忽略
### 批量创建
create传递一个slice
或者
采用CreateInBatches
### 通过map 或者slice map创建
GORM 支持根据 map[string]interface{} 和 []map[string]interface{}{} 创建
## 查询
###  gorm.ErrRecordNotFound
gorm 中，First、Last、Take、Find 方法都可以用于查询单条记录. 前三个方法的特点是，倘若未查询到指定记录，则会报错 gorm.ErrRecordNotFound；最后一个方法的语义更软一些，即便没有查到指定记录，也不会返回错误.
### 内联条件(建议不使用)
内联条件查询需要注意SQL注入问题
```Go
db.First(&user, "id = ?", "string_primary_key")
// SELECT * FROM users WHERE id = 'string_primary_key';

// Plain SQL
db.Find(&user, "name = ?", "jinzhu")
// SELECT * FROM users WHERE name = "jinzhu";

db.Find(&users, "name <> ? AND age > ?", "jinzhu", 20)
// SELECT * FROM users WHERE name <> "jinzhu" AND age > 20;

// Struct
db.Find(&users, User{Age: 20})
// SELECT * FROM users WHERE age = 20;

// Map
db.Find(&users, map[string]interface{}{"age": 20})
// SELECT * FROM users WHERE age = 20;
```
### 查询条件
使用where

    可以使用string条件,map条件,struct条件,以及主键切片条件

    当使用结构作为条件查询时，GORM 只会查询非零值字段。这意味着如果您的字段值为 0、''、false 或其他 零值，该字段不会被用于构建查询条件
### 查询结果的接收

对于单个查询结果,可以接收map[string]interface{}
对于多个查询结果,可以接收[]Model类型以及[]map[string]interface{}
### 查询特定的字段
查询特定字段时候,可以通过Select指定
```Go
db.Table("users").Select("COALESCE(age,?)", 42).Rows()
// SELECT COALESCE(age,'42') FROM users;
```
### order
默认是升序,降序:
```Go
db.Order("age desc").Order("name").Find(&users)
// SELECT * FROM users ORDER BY age desc, name;
```
### Limit 
通过Limit(-1)消除链式查询中的limit条件
### Offset
通过Offset(-1)消除链式查询中的limit条件
### 分页
```Go

func Paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
  return func (db *gorm.DB) *gorm.DB {
    page, _ := strconv.Atoi(r.Query("page"))
    if page == 0 {
      page = 1
    }

    pageSize, _ := strconv.Atoi(r.Query("page_size"))
    switch {
    case pageSize > 100:
      pageSize = 100
    case pageSize <= 0:
      pageSize = 10
    }

    offset := (page - 1) * pageSize
    return db.Offset(offset).Limit(pageSize)
  }
}

db.Scopes(Paginate(r)).Find(&users)
db.Scopes(Paginate(r)).Find(&articles)
```

### Group
注意:
  
    可以不用预先分配内存给slice或者map
## save
待更新的模型数据中只包含表的部分字段时，Save函数会把未指定的字段值更新成对应类型的默认值。

save的操作逻辑:
	
	1.如果对应的struct有主键,则根据主键条件进行更新
	1.1如果第一步失败,会继续执行
	create ... on duplicate key update ...来尝试创建记录
	
	2.如果对应的struct没有主键,则创建

	gorm默认的主键名称为ID,如果主键名称为非ID的形式,要通过对model主键对应的struct字段打primaryKey的tag指定

	如果存在ID,但是不是主键,并没有通过tag指定主键,Gorm仍然以ID作为主键的方式进行操作


	
## 关联关系
### belong to
A belong to B

	A中存在外键,与主表中的主键/唯一列形成一对一映射
	A内嵌了B的模型

 	外键默认的名称主表名称+主表主键
```Go
type NormalBelongToCompany struct {
	SID          int    `gorm:"primaryKey"`
	CompanyName string `gorm:"type:varchar(30);"`
}

type NormalBelongToUser struct {
	ID                      int    `gorm:"primaryKey"`
	UserName                string `gorm:"type:varchar(30)"`
	NormalBelongToCompanySID string `gorm:"type:varchar(30)"` //foreignKey
	NormalBelongToCompany   NormalBelongToCompany
}

func (n *NormalBelongToCompany) TableName() string {
	return "normal_belong_to_company"
}
func (n *NormalBelongToUser) TableName() string {
	return "normal_belong_to_user"
}

```
```SQL
CREATE TABLE `normal_belong_to_company` (`s_id` bigint AUTO_INCREMENT,`company_name` varchar(30),PRIMARY KEY (`id`))

CREATE TABLE `normal_belong_to_user` (`id` bigint AUTO_INCREMENT,`user_name` varchar(30),`normal_belong_to_company_s_id` bigint,PRIMARY KEY (`id`),CONSTRAINT `fk_normal_belong_to_user_normal_belong_to_company` FOREIGN KEY (`normal_belong_to_company_s_id`) REFERENCES `normal_belong_to_company`(`s_id`))
```
以上述的例子为例:

主表: NormalBelongToCompany
子表: NormalBelongToUser

	外键名称默认是主表名称+主键名称

	如果外键名词需要指定非默认的名称,则需要通过foreignKey来指定
	
	如果外键对应的不是主表的主键,则对应的一定是一个主表的唯一键,必须显式指定unique,并通过references指定


```Go
type ForeignKeyBelongToCompany struct {
	ID          int    `gorm:"primaryKey"`
	CompanyName string `gorm:"type:varchar(30);"`
}

type ForeignKeyBelongToUser struct {
	ID        int                       `gorm:"primaryKey"`
	UserName  string                    `gorm:"type:varchar(30)"`
	CompanyId string                    `gorm:"type:varchar(30)"` //foreignKey
	Company   ForeignKeyBelongToCompany `gorm:"foreignKey:CompanyId"`
}

func (n *ForeignKeyBelongToCompany) TableName() string {
	return "foreign_key_belong_to_company"
}
func (n *ForeignKeyBelongToUser) TableName() string {
	return "foreign_key_belong_to_user"
}
```
```SQL
CREATE TABLE `foreign_key_belong_to_company` (`id` bigint AUTO_INCREMENT,`company_name` varchar(30),PRIMARY KEY (`id`))

CREATE TABLE `foreign_key_belong_to_user` (`id` bigint AUTO_INCREMENT,`user_name` varchar(30),`company_id` bigint,PRIMARY KEY (`id`),CONSTRAINT `fk_foreign_key_belong_to_user_company` FOREIGN KEY (`company_id`) REFERENCES `foreign_key_belong_to_company`(`id`))
```

```Go
type ReferenceBelongToCompany struct {
	ID          int    `gorm:"primaryKey"`
	CompanyName string `gorm:"type:varchar(30);unique"`
}

type ReferenceBelongToUser struct {
	ID                         int                      `gorm:"primaryKey"`
	UserName                   string                   `gorm:"type:varchar(30)"`
	ReferenceBelongToCompanyID string                   `gorm:"type:varchar(30)"` //不使用foreignKey显式指定时候必须是这个ReferenceBelongToCompany名词,不然报错
	ReferenceBelongToCompany   ReferenceBelongToCompany `gorm:"references:CompanyName"`
}

func (n *ReferenceBelongToCompany) TableName() string {
	return "reference_key_belong_to_company"
}
func (n *ReferenceBelongToUser) TableName() string {
	return "reference_key_belong_to_user"
}

```
```SQL
CREATE TABLE `reference_key_belong_to_company` (`id` bigint AUTO_INCREMENT,`company_name` varchar(30) UNIQUE,PRIMARY KEY (`id`))

CREATE TABLE `reference_key_belong_to_user` (`id` bigint AUTO_INCREMENT,`user_name` varchar(30),`reference_belong_to_company_id` varchar(30),PRIMARY KEY (`id`),CONSTRAINT `fk_reference_key_belong_to_user_reference_belong_to_company` FOREIGN KEY (`reference_belong_to_company_id`) REFERENCES `reference_key_belong_to_company`(`company_name`))
```


当外键关联的是主表的非主键.要求显式的指定关联键是唯一键(unique)

否则会报
Error 1822 (HY000): Failed to add the foreign key constraint. Missing index for constraint 'fk_reference_key_belong_to_user_reference_belong_to_company' in the referenced table 'reference_key_belong_to_company'

外键的名称还是默认的主表名称+主表主键名称

```GO
type ReferenceForeignKeyBelongToCompany struct {
	ID          int    `gorm:"primaryKey"`
	CompanyName string `gorm:"type:varchar(30);unique"`
}

type ReferenceForeignKeyBelongToUser struct {
	ID         int                                `gorm:"primaryKey"`
	UserName   string                             `gorm:"type:varchar(30)"`
	//!!!如果采用CompanyName则会报错,并且创建逻辑会异常!!!
	CompanyKey string                             `gorm:"type:varchar(30)"` //foreignKey
	Company    ReferenceForeignKeyBelongToCompany `gorm:"foreignKey:CompanyKey;references:CompanyName;"`
}

func (n *ReferenceForeignKeyBelongToCompany) TableName() string {
	return "reference_belong_to_key_belong_to_company"
}
func (n *ReferenceForeignKeyBelongToUser) TableName() string {
	return "reference_belong_to_key_belong_to_user"
}
```
```SQL
CREATE TABLE `reference_belong_to_key_belong_to_company` (`id` bigint AUTO_INCREMENT,`company_name` varchar(30) UNIQUE,PRIMARY KEY (`id`))

CREATE TABLE `reference_belong_to_key_belong_to_user` (`id` bigint AUTO_INCREMENT,`user_name` varchar(30),`company_key` varchar(30),PRIMARY KEY (`id`),CONSTRAINT `fk_reference_belong_to_key_belong_to_user_company` FOREIGN KEY (`company_key`) REFERENCES `reference_belong_to_key_belong_to_company`(`company_name`))
```

注意:
	
	!!!BelongTo情况下在同时使用foreignKey和references时,外键名称和对应的主表的键名称不能一样,不然会导致异常!!!
	HasOne没有该问题
### has one
### has one 和 belong to总结
不论是BelongTo或者是HasOne,二者对应的数据库结构都是主表和从表的关系,即从表的外键和主表的关联关系
而所谓的BelongTo和HasOne只是针对于golang的结构体而言,通过从属的结构体内嵌表示从属关系

### manyTomany
需要注意的是:
对于多对多,存在反向引用的时候,如果要自定义外键时候,要同时指定foreignKey 和 references,如果只指定foreignKey则会出现异常
```Go
type ForeignKeyMany2ManyUser struct {
	ID        string `gorm:"primaryKey"`
	UserName  string `gorm:"unique"`
	UserInfo  string
	Languages []ForeignKeyMany2ManyLanguage `gorm:"many2many:foreign_key_many_2_many_mid_table;foreignKey:UserName;"`
}

func (n *ForeignKeyMany2ManyUser) TableName() string {
	return "foreign_key_many_to_many_user"
}

type ForeignKeyMany2ManyLanguage struct {
	ID           string `gorm:"primaryKey"`
	LanguageName string `gorm:"unique"`
	LanguageInfo string
	Users        []ForeignKeyMany2ManyUser `gorm:"many2many:foreign_key_many_2_many_mid_table;foreignKey:LanguageName;"`
}

func (n *ForeignKeyMany2ManyLanguage) TableName() string {
	return "foreign_key_many_to_many_language"
}

```
对应的migrate语句
```SQL
CREATE TABLE `foreign_key_many_to_many_language` (`id` varchar(191),`language_name` varchar(191) UNIQUE,`language_info` longtext,PRIMARY KEY (`id`))

CREATE TABLE `foreign_key_many_to_many_user` (`id` varchar(191),`user_name` varchar(191) UNIQUE,`user_info` longtext,PRIMARY KEY (`id`))

CREATE TABLE `foreign_key_many_2_many_mid_table` (`foreign_key_many2_many_user_user_name` varchar(191),`foreign_key_many2_many_language_id` varchar(191),PRIMARY KEY (`foreign_key_many2_many_user_user_name`,`foreign_key_many2_many_language_id`),CONSTRAINT `fk_foreign_key_many_2_many_mid_table_foreign_key_many2_m6fd85cb4` FOREIGN KEY (`foreign_key_many2_many_language_id`) REFERENCES `foreign_key_many_to_many_language`(`id`),CONSTRAINT `fk_foreign_key_many_2_many_mid_table_foreign_key_many2_many_user` FOREIGN KEY (`foreign_key_many2_many_user_user_name`) REFERENCES `foreign_key_many_to_many_user`(`user_name`))

ALTER TABLE `foreign_key_many_2_many_mid_table` ADD `foreign_key_many2_many_language_language_name` varchar(191)

ALTER TABLE `foreign_key_many_2_many_mid_table` ADD `foreign_key_many2_many_user_id` varchar(191)
```
可以发现反向引用定义的外键并没有生效

必须同时指定 foreignKey 和references
```Go
type ForeignKeyMany2ManyUser struct {
	ID        string `gorm:"primaryKey"`
	UserName  string `gorm:"unique"`
	UserInfo  string
	Languages []ForeignKeyMany2ManyLanguage `gorm:"many2many:foreign_key_many_2_many_mid_table;foreignKey:UserName;references:LanguageName;"`
}

func (n *ForeignKeyMany2ManyUser) TableName() string {
	return "foreign_key_many_to_many_user"
}

type ForeignKeyMany2ManyLanguage struct {
	ID           string `gorm:"primaryKey"`
	LanguageName string `gorm:"unique"`
	LanguageInfo string
	Users        []ForeignKeyMany2ManyUser `gorm:"many2many:foreign_key_many_2_many_mid_table;foreignKey:LanguageName;references:UserName;"`
}

func (n *ForeignKeyMany2ManyLanguage) TableName() string {
	return "foreign_key_many_to_many_language"
}
```
```SQL
CREATE TABLE `foreign_key_many_to_many_language` (`id` varchar(191),`language_name` varchar(191) UNIQUE,`language_info` longtext,PRIMARY KEY (`id`))

CREATE TABLE `foreign_key_many_to_many_user` (`id` varchar(191),`user_name` varchar(191) UNIQUE,`user_info` longtext,PRIMARY KEY (`id`))

CREATE TABLE `foreign_key_many_2_many_mid_table` (`foreign_key_many2_many_user_user_name` varchar(191),`foreign_key_many2_many_language_language_name` varchar(191),PRIMARY KEY (`foreign_key_many2_many_user_user_name`,`foreign_key_many2_many_language_language_name`),CONSTRAINT `fk_foreign_key_many_2_many_mid_table_foreign_key_many2_many_user` FOREIGN KEY (`foreign_key_many2_many_user_user_name`) REFERENCES `foreign_key_many_to_many_user`(`user_name`),CONSTRAINT `fk_foreign_key_many_2_many_mid_table_foreign_key_many2_m6fd85cb4` FOREIGN KEY (`foreign_key_many2_many_language_language_name`) REFERENCES `foreign_key_many_to_many_language`(`language_name`))
```
或者取消反向引用
```Go
type ForeignKeyMany2ManyUser struct {
	ID        string `gorm:"primaryKey"`
	UserName  string `gorm:"unique"`
	UserInfo  string
	Languages []ForeignKeyMany2ManyLanguage `gorm:"many2many:foreign_key_many_2_many_mid_table;foreignKey:UserName;references:LanguageName;"`
}

func (n *ForeignKeyMany2ManyUser) TableName() string {
	return "foreign_key_many_to_many_user"
}

type ForeignKeyMany2ManyLanguage struct {
	ID           string `gorm:"primaryKey"`
	LanguageName string `gorm:"unique"`
	LanguageInfo string
}

func (n *ForeignKeyMany2ManyLanguage) TableName() string {
	return "foreign_key_many_to_many_language"
}

```
```SQL
CREATE TABLE `foreign_key_many_to_many_language` (`id` varchar(191),`language_name` varchar(191) UNIQUE,`language_info` longtext,PRIMARY KEY (`id`))

CREATE TABLE `foreign_key_many_to_many_user` (`id` varchar(191),`user_name` varchar(191) UNIQUE,`user_info` longtext,PRIMARY KEY (`id`))

CREATE TABLE `foreign_key_many_2_many_mid_table` (`foreign_key_many2_many_user_user_name` varchar(191),`foreign_key_many2_many_language_language_name` varchar(191),PRIMARY KEY (`foreign_key_many2_many_user_user_name`,`foreign_key_many2_many_language_language_name`),CONSTRAINT `fk_foreign_key_many_2_many_mid_table_foreign_key_many2_many_user` FOREIGN KEY (`foreign_key_many2_many_user_user_name`) REFERENCES `foreign_key_many_to_many_user`(`user_name`),CONSTRAINT `fk_foreign_key_many_2_many_mid_table_foreign_key_many2_m6fd85cb4` FOREIGN KEY (`foreign_key_many2_many_language_language_name`) REFERENCES `foreign_key_many_to_many_language`(`language_name`))
```
自定义连接表的标准操作(不同于官方文档)
```Go
type CustomMany2ManyUser struct {
	ID        string `gorm:"primaryKey"`
	UserName  string `gorm:"unique"`
	UserInfo  string
	Languages []CustomMany2ManyLanguage `gorm:"many2many:custom_many_to_many_user_language_mid_table;foreignKey:UserName;joinForeignKey:UserName;references:LanguageName;joinReferences:LanguageName"`
}

func (n *CustomMany2ManyUser) TableName() string {
	return "custom_many_to_many_user"
}

type CustomMany2ManyLanguage struct {
	ID           string `gorm:"primaryKey"`
	LanguageName string `gorm:"unique"`
	LanguageInfo string
	Users        []CustomMany2ManyUser `gorm:"many2many:custom_many_to_many_user_language_mid_table;foreignKey:LanguageName;joinForeignKey:LanguageName;references:UserName;joinReferences:UserName"`
}

func (n *CustomMany2ManyLanguage) TableName() string {
	return "custom_many_to_many_language"
}

type CustomMany2ManyUserLanguageMidTable struct {
	UserName     string `gorm:"primaryKey"`
	LanguageName string `gorm:"primaryKey"`
	OtherInfo    string
}

func (n *CustomMany2ManyUserLanguageMidTable) TableName() string {
	return "custom_many_to_many_user_language_mid_table"
}

```

```SQL
CREATE TABLE `custom_many_to_many_user` (`id` varchar(191),`user_name` varchar(191) UNIQUE,`user_info` longtext,PRIMARY KEY (`id`))

CREATE TABLE `custom_many_to_many_language` (`id` varchar(191),`language_name` varchar(191) UNIQUE,`language_info` longtext,PRIMARY KEY (`id`))

CREATE TABLE `custom_many_to_many_user_language_mid_table` (`user_name` varchar(191),`language_name` varchar(191),PRIMARY KEY (`user_name`,`language_name`),CONSTRAINT `fk_custom_many_to_many_user_language_mid_table_custom_ma82cd0b74` FOREIGN KEY (`user_name`) REFERENCES `custom_many_to_many_user`(`user_name`),CONSTRAINT `fk_custom_many_to_many_user_language_mid_table_custom_maf64eeb08` FOREIGN KEY (`language_name`) REFERENCES `custom_many_to_many_language`(`language_name`))

ALTER TABLE `custom_many_to_many_user_language_mid_table` ADD `other_info` longtext
```
### 多对多关联关系的处理
需要注意,如果采用save保存多对多数据的变更,通过Debug发现操作复杂
建议直接使用Association进行操作
```Go
type ForeignKeyMany2ManyUser struct {
	ID        int    `gorm:"primaryKey"`
	UserName  string `gorm:"unique"`
	UserInfo  string
	Languages []ForeignKeyMany2ManyLanguage `gorm:"many2many:foreign_key_many_2_many_mid_table;foreignKey:UserName;joinForeignKey:UserName;references:LanguageName;joinReferences:LanguageName"`
}

func (n *ForeignKeyMany2ManyUser) TableName() string {
	return "foreign_key_many_to_many_user"
}

type ForeignKeyMany2ManyLanguage struct {
	ID           int    `gorm:"primaryKey"`
	LanguageName string `gorm:"unique"`
	LanguageInfo string
	Users        []ForeignKeyMany2ManyUser `gorm:"many2many:foreign_key_many_2_many_mid_table;foreignKey:LanguageName;joinForeignKey:LanguageName;references:UserName;joinReferences:UserName"`
}

func (n *ForeignKeyMany2ManyLanguage) TableName() string {
	return "foreign_key_many_to_many_language"
}

func Create3() {
	user1 := model.ForeignKeyMany2ManyUser{
		ID:        1,
		UserName:  "user1",
		UserInfo:  "this is user1",
		Languages: []model.ForeignKeyMany2ManyLanguage{},
	}
	user2 := model.ForeignKeyMany2ManyUser{
		ID:        2,
		UserName:  "user2",
		UserInfo:  "this is user2",
		Languages: []model.ForeignKeyMany2ManyLanguage{},
	}
	language1 := model.ForeignKeyMany2ManyLanguage{
		ID:           1,
		LanguageName: "language1",
		LanguageInfo: "this is language1",
		Users:        []model.ForeignKeyMany2ManyUser{},
	}
	language2 := model.ForeignKeyMany2ManyLanguage{
		ID:           2,
		LanguageName: "language2",
		LanguageInfo: "this is language2",
		Users:        []model.ForeignKeyMany2ManyUser{},
	}

	db := connection.GormDB.Debug().Create(user1)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", user1)
	}
	db = connection.GormDB.Debug().Create(user2)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", user2)
	}
	user1.Languages = []model.ForeignKeyMany2ManyLanguage{language1, language2}
	user2.Languages = []model.ForeignKeyMany2ManyLanguage{language2}
	fmt.Println(111)
	db = connection.GormDB.Debug().Save(&user1)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", user1)
	}
	fmt.Println(222)

	db = connection.GormDB.Debug().Save(&user2)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", user2)
	}
}
```
```SQL
INSERT INTO `foreign_key_many_to_many_user` (`user_name`,`user_info`,`id`) VALUES ('user1','this is user1',1)

INSERT INTO `foreign_key_many_to_many_user` (`user_name`,`user_info`,`id`) VALUES ('user2','this is user2',2)
//插入user1的关联关系前,创建languages数据
INSERT INTO `foreign_key_many_to_many_language` (`language_name`,`language_info`,`id`) VALUES ('language1','this is language1',1),('language2','this is language2',2) ON DUPLICATE KEY UPDATE `id`=`id`

//插入user1的关联关系
INSERT INTO `foreign_key_many_2_many_mid_table` (`user_name`,`language_name`) VALUES ('user1','language1'),('user1','language2') ON DUPLICATE KEY UPDATE `user_name`=`user_name`

//插入user2的关联关系前,创建languages数据
INSERT INTO `foreign_key_many_to_many_language` (`language_name`,`language_info`,`id`) VALUES ('language1','this is language1',1) ON DUPLICATE KEY UPDATE `id`=`id`
//插入user2的关联关系
INSERT INTO `foreign_key_many_2_many_mid_table` (`user_name`,`language_name`) VALUES ('user2','language1') ON DUPLICATE KEY UPDATE `user_name`=`user_name`
```

### 多对多关系中的软删除

如果存在软删除字段,则删除不会删除关联信息
## 注意事项

### 表名,字段名的命名规则
gorm的命令策略是表为蛇形复数,字段名是蛇形单数

  可以通过tag字段column指定列名
  可以通过TableName()指定表明




### gormDB.Table 和gorm.Model的区别
```Go
result = connection.GormDB.Debug().Table("base_model").First(&data)
	fmt.Printf("RowAffected %v\n", result.RowsAffected)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Println(result.Error.Error())
		} else {
			log.Fatalln(result.Error.Error())
		}
	} else {
		fmt.Println(data)
	}
```
First、Last 方法会根据主键查找到第一个、最后一个记录， 它仅在通过 struct 或提供 model 值进行查询时才起作用
因此,使用Table是不能够执行First或者Last操作的

### Take

take为全表扫描

### row/rows
https://juejin.cn/post/6844904038140477453#heading-5
rows.Next() 在获取到最后一条记录之后，会调用 rows.Close() 将连接放回连接池或交给其他等待的请求方，所以不需要手动调用 rows.Close()
如果在rows.Next()没有执行完就return,则连接不会归还,需要手动调用rows.Close()不然导致连接的占用
```Go
rows, err := db.Model(&Modle{}).Where("id = ?", "1").Rows()
defer rows.Close()
```

### distinct
distinct 表示后面的查询的所有字段都不重复

### 使用了Gorm.DeletedAt字段
考虑查询语句会带有where DeletedAt is NULL
### Find,Scan
查询不到不会报错,建议的处理逻辑如下:
```Go
	var basicModel model.Basic
	db := connection.GormDB.Debug().Model(&basicModel).Where("id=?", 10).Scan(&basicModel)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else if db.RowsAffected != 0 {
		fmt.Printf("%+v\n", basicModel)
	}
```

```Go
	var basicModel model.Basic
	var sliceResult []struct {
		Count     int
		EmbedInfo string
	}
	//自定义字段通过Find映射
	//Find查询不到不会报错
	db := connection.GormDB.Debug().Session(&gorm.Session{SkipHooks: true}).Model(&basicModel).Select("count(*) as count, embed_info").Where("id<?", 20).Group("embed_info").Having("count>?", 0).Find(&sliceResult)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else if db.RowsAffected != 0 {
		fmt.Printf("%+v\n", sliceResult)
	}
```
### 注意Assign和Attrs在FirstOrInit与FirstOrCreate中的区别

对于FirstOrInit,本身不会创建,只是将信息初始化到某个struct

	Attrs   找到不生效
			找不到生效

	Assgin	找不到/找到都会生效

对于FirstOrCreate

	Attrs   找到不生效
			找不到生效

	Assign  找不到/找到都会生效,并且都会更新数据库

### 零值的处理

	更新:
		Updates 提供的参数如果是struct,则只会更新非零值的字段
	创建:
		创建时候， 0、''、false 等零值，不会将这些字段定义的默认值保存到数据库
	查询:
		当使用结构作为条件查询时，GORM 只会查询非零值字段。这意味着如果您的字段值为 0、''、false 或其他 零值，该字段不会被用于构建查询条件


### gorm存在注入风险的地方
```Go
db.Select("name; drop table users;").First(&user)

db.Distinct("name; drop table users;").First(&user)

db.Model(&user).Pluck("name; drop table users;", &names)

db.Group("name; drop table users;").First(&user)

db.Group("name").Having("1 = 1;drop table users;").First(&user)

db.Raw("select name from users; drop table users;").First(&user)

db.Exec("select name from users; drop table users;")

```

### 优化器,索引提示

优化器提示用于控制查询优化器选择某个查询执行计划，GORM 通过 gorm.io/hints 提供支持
索引提示允许传递索引提示到数据库，以防查询计划器出现混乱

### 锁配置

