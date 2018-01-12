package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "user=manga password=manga host=127.0.0.1 port=5432 dbname=manga sslmode=disable")
}
