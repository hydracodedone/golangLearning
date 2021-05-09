package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

/*
sql.Open() 不会建立与数据库的任何连接，也不会验证驱动程序的连接参数。而是它只准备数据库抽象以备后用。
与基础数据存储区的第一个实际连接将延迟到第一次需要时建立。
如果要立即检查数据库是否可用和可访问 (例如，检查是否可以建立网络连接并登录) 请使用 db.Ping() 进行操作，并记录检查错误
*/
/*
但是 sql.DB 对象被设计为长期存在。
不要经常使用 Open() 和 Close()。
而是为每个需要访问的不同数据存储创建 一个 sql.DB 对象，并保留该对象直到程序完成对该数据存储的访问为止。
根据需要传递它，或以某种方式使其在全局范围内可用，但保持打开状态。不要通过短期函数来 Open() 和 Close()。
而是将 sql.DB 传递给该短期函数作为参数。
*/
type user struct {
	id   int
	name string
	age  int
	sex  string
}
type userList []*user

func databaseInit() *sql.DB {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/user")
	if err != nil {
		fmt.Println("OPEN FAIL")
		fmt.Println(err)
		return nil
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("PING FAIL")
		fmt.Println(err)
		return nil
	}
	return db
}
func databaseClose(db *sql.DB) {
	if db == nil {
		fmt.Println("NO DATABASE SKIP")
	} else {
		err := db.Close()
		if err != nil {
			panic(fmt.Errorf("ERROR: DATABASE CLOSE FAIL->%s", err))
		}
	}
}

/*"insert into `USER` (`name`,`age`,`sex`) values(?,?,?)"
如果一个函数名字包含 Query, 那么该函数旨在向数据库发出查询问题，并且即使它为空，也将返回一组行。
不返回行的语句不应该使用 Query 函数；而应使用 Exec()
*/
func databaseQuery(db *sql.DB) userList {
	var queryUser userList
	query, err := db.Prepare("select * from `USER` ")
	if err != nil {
		fmt.Printf("ERROR: PREPARE QUERY FAIL->%s", err)
		return nil
	}
	defer query.Close()
	rows, err := query.Query()
	if err != nil {
		fmt.Printf("ERROR: QUERY FAIL->%s", err)
		return nil
	}
	fmt.Println("QUERY SUCCESS")
	//释放连接
	defer func() {
		err = rows.Close()
		if err != nil {
			panic(fmt.Errorf("ERROR: ROWS CLOSE FAIL->%s", err))
		} else {
			fmt.Println("CLOSE ROWS SUCCESS")
		}
	}()
	for rows.Next() {
		singleUSer := user{}
		err = rows.Scan(&singleUSer.id, &singleUSer.name, &singleUSer.age, &singleUSer.sex)
		if err != nil {
			return queryUser
		}
		queryUser = append(queryUser, &singleUSer)
	}
	err = rows.Err()
	if err != nil {
		fmt.Printf("ERROR: ROWS ERROR-> %s\n", err)
		return queryUser
	}
	return queryUser
}

func databaseQuerySingle(db *sql.DB) *user {
	var userSingle = user{}
	if db == nil {
		return nil
	}
	singleQuery, err := db.Prepare("select * from `USER` where id =?")
	if err != nil {
		fmt.Printf("ERROR: PREPARE QUERY FAIL->%s", err)
		return nil
	}
	defer singleQuery.Close()
	//单行查询QueryRow调用的Scan语句包含了释放链接的方法,但是多行查询并没有包含释放连接的方法,需要手动执行
	err = singleQuery.QueryRow(1).Scan(&userSingle.id, &userSingle.name, &userSingle.age, &userSingle.sex)
	defer func() {
		err = singleQuery.Close()
		if err != nil {
			panic(fmt.Errorf("ERROR: ROW CLOSE FAIL->%s", err))
		} else {
			fmt.Println("CLOSE ROWS SUCCESS")
		}
	}()
	if err == sql.ErrNoRows {
		fmt.Println("NOT FOUND MATCHED ROW")
		return nil
	} else if err != nil {
		fmt.Println(err)
		return nil
	} else {
		return &userSingle
	}
}
func databaseExec(db *sql.DB) bool {
	if db == nil {
		fmt.Println("ERROR: DB IS NIL")
		return false
	}
	exec, err := db.Prepare("insert into `USER` (`name`,`age`,`sex`) values(?,?,?)")
	if err != nil {
		fmt.Printf("ERROR: PREPARE FAIL->%s", err)
		return false
	}
	defer exec.Close()
	res, err := exec.Exec("SomeoneElse", 24, "female")
	if err != nil {
		fmt.Printf("ERROR: EXEC FAIL->%s", err)
		return false
	}
	_, err = res.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return false
	}
	num, err := res.RowsAffected()
	if num != 1 || err != nil {
		if num != 1 {
			fmt.Println("ERROR AFFECT ROW NOT EQUAL ONE")
		}
		if err != nil {
			fmt.Println(err)
		}
		return false
	}
	return true
}

func databaseExecWithTransaction(tx *sql.Tx) bool {
	rollbackFunction := func() {
		err := tx.Rollback()
		if err != nil {
			fmt.Println("ROOLBACK FAIL")
			panic(err)
		}
	}
	if tx == nil {
		fmt.Println("ERROR: DB IS NIL")
		return false
	}
	exec, err := tx.Prepare("insert into `USER` (`name`,`age`,`sex`) values(?,?,?)")
	if err != nil {
		fmt.Printf("ERROR: PREPARE FAIL->%s", err)
		rollbackFunction()
		return false
	}
	defer exec.Close()
	res, err := exec.Exec("SomeoneElse", 24, "female")
	if err != nil {
		fmt.Printf("ERROR: EXEC FAIL->%s", err)
		rollbackFunction()
		return false
	}
	_, err = res.LastInsertId()
	if err != nil {
		fmt.Println(err)
		rollbackFunction()
		return false
	}
	num, err := res.RowsAffected()
	if num != 1 || err != nil {
		if num != 1 {
			fmt.Println("ERROR AFFECT ROW NOT EQUAL ONE")
		}
		if err != nil {
			rollbackFunction()
			fmt.Println(err)
		}
		return false
	}
	return true
}
func databaseTransaction(db *sql.DB, execFunc func(*sql.Tx) bool) bool {
	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("ERROR BEGIN TRANSACTION FAIl->%s", err)
		return false
	}
	execFunc(tx)
	err = tx.Commit()
	if err != nil {
		fmt.Printf("ERROR COMMIT TRANSACTION FAIl->%s", err)
		err = tx.Rollback()
		if err != nil {
			fmt.Printf("ERROR ROLLBACK TRANSACTION FAIl->%s", err)
			return false
		}
	}
	return false

}

func main() {
	db := databaseInit()
	if db == nil {
		return
	}
	defer databaseClose(db)
	res := databaseTransaction(db, databaseExecWithTransaction)
	fmt.Println(res)

}
