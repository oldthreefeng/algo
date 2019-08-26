/*
@Time : 2019/8/24 22:55
@Author : Administrator
@File : main_v1
@Software: GoLand
*/

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type userinfo struct {
	uid int
	username,department,created string
}

/*
CREATE TABLE `userinfo` (
	`uid` INT(10) NOT NULL AUTO_INCREMENT,
	`username` VARCHAR(64) NULL DEFAULT NULL,
	`department` VARCHAR(64) NULL DEFAULT NULL,
	`created` DATE NULL DEFAULT NULL,
	PRIMARY KEY (`uid`)
);

CREATE TABLE `userdetail` (
	`uid` INT(10) NOT NULL DEFAULT '0',
	`intro` TEXT NULL,
	`profile` TEXT NULL,
	PRIMARY KEY (`uid`)
)
*/

func main() {
	//username:password@protocol(address)/dbname?param=value
	//protocol : tcp  address: ip:port
	db,err:= sql.Open("mysql","root:wangke123@tcp(101.132.33.146:3306)/testgo?charset=utf8")
	checkErr(err)
	defer db.Close()
	//插入数据
	//stms,err := db.Prepare("INSERT INTO userinfo SET username=?,department=?,created=?")
	res,err := db.Exec("INSERT INTO userinfo(username,department,created) VALUES (?,?,?)","louis","研发部","2019-08-24")
	checkErr(err)
	//res,err := stms.Exec("louis","研发部","2019-08-24")
	//checkErr(err)
	id,err:= res.LastInsertId()
	checkErr(err)
	fmt.Println(id)

	//更新数据
	stms,err := db.Prepare("UPDATE userinfo SET username=? WHERE uid=?")
	checkErr(err)
	res,err = stms.Exec("louisHon",id)
	checkErr(err)
	affect,err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)

	//查询数据

	rows,err := db.Query("SELECT * from userinfo")
	checkErr(err)
	for rows.Next() {
		user := userinfo{}
		err := rows.Scan(&user.uid,&user.username,&user.department,&user.created)
		checkErr(err)
		fmt.Println(user)
	}
}