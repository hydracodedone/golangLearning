# golang 标准库解析
## MYSQL错误参考
https://github.com/VividCortex/mysqlerr
## 导入
```Go
import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)
```

## 注意
### 数据库连接
sql.Open() 不会建立与数据库的任何连接，也不会验证驱动程序的连接参数。而是它只准备数据库抽象以备后用。与基础数据存储区的第一个实际连接将延迟到第一次需要时建立。如果要立即检查数据库是否可用和可访问 (例如，检查是否可以建立网络连接并登录)，请使用 db.Ping() 进行操作
```Go
err = db.Ping()
if err != nil {
}
```
sql.DB 对象被设计为长期存在。不要经常使用 Open() 和 Close(),如果不将 sql.DB 视为长期对象，则可能会遇到诸如重用和连接共享不良
### Row的正确操作
```Go
var (
    id int
    name string
)
rows, err := db.Query("select id, name from users where id = ?", 1)
if err != nil {
    log.Fatal(err)
}
defer func(){
    err:=rows.Close() //rows.Close()在遍历异常退出时候不会归还连接,通过defer来保证归还(rows.Close()可以重复执行多次)
    if err != nil {
        log.Fatal(err)
    }
}()
for rows.Next() {
    err := rows.Scan(&id, &name) //应该始终检查 for rows.Next() 循环的末尾是否有错误
    if err != nil {
        log.Fatal(err)
    }
    log.Println(id, name)
}
err = rows.Err()
if err != nil {
    log.Fatal(err)
}

```
### queryRow 产生的sql.ErrNoRows
```Go
var name string
err = db.QueryRow("select name from users where id = ?", 1).Scan(&name)
if err != nil {
    if err == sql.ErrNoRows {
        log.Println("no rows found")
    } else {
        log.Fatal(err)
    }
}
```
### NULL的处理
从 SQL 语句出发,使用COALESCE避免NULL的处理
```SQL
select COALESCE(NULL,'') as result;
```
### 连接池

## 核心模块介绍
database/sql 关于数据库驱动模块下各核心 interface 主要包括：

    Connector：抽象的数据库连接器，需要具备创建数据库连接以及返回从属的数据库驱动的能力
    Driver：抽象的数据库驱动，具备创建数据库连接的能力
    Conn：抽象的数据库连接，具备预处理 sql 以及开启事务的能力
    Tx：抽象的事务，具备提交和回滚的能力
    Statement：抽象的请求预处理状态. 具备实际执行 sql 并返回执行结果的能力
    Result/Row：抽象的 sql 执行结果
![alt text](image.png)