package db

import (
	"fmt"
	_ "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {

	// PostgreSQL 配置
	orm.RegisterDriver("postgres", orm.DRPostgres) // 注册驱动
	orm.RegisterDataBase("default", "postgres", "user=postgres password=123456 dbname=mlt host=localhost port=5432 sslmode=disable")

	/**
	 * MySQL 配置
	 * 注册驱动
	 * orm.RegisterDriver("mysql", orm.DR_MySQL)
	 * mysql用户：root ，root的秘密：tom ， 数据库名称：test ， 数据库别名：default
	 * orm.RegisterDataBase("default", "mysql", "root:tom@/test?charset=utf8")
	 */
	/**
	 * Sqlite 配置
	 * 注册驱动
	 * orm.RegisterDriver("sqlite", orm.DR_Sqlite)
	 * 数据库存放位置：./datas/test.db ， 数据库别名：default
	 * orm.RegisterDataBase("default", "sqlite3", "./datas/test.db")
	 */
	// 自动建表

	//orm.RunSyncdb("default", false, true)
	orm.Debug = true
	//o := orm.NewOrm()
	//o.Using("default")
	fmt.Println("数据库连接成功")
}


