package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func Init() {
	db_user := beego.AppConfig.String("DB_USER")
	db_host := beego.AppConfig.String("DB_HOST")
	db_port := beego.AppConfig.String("DB_PORT")
	db_password := beego.AppConfig.String("DB_PASSWORD")
	db_name := beego.AppConfig.String("DB_NAME")
	if db_port == "" {
		db_port = "3306"
	}
	dsn := db_user + ":" + db_password + "@tcp(" + db_host + ":" + db_port + ")/" + db_name + "?charset=utf8&loc=Asia%2FShanghai"
	orm.RegisterDataBase("default", "mysql", dsn)
	// orm.RegisterModel()
}

func TableName(str string) string {
	return beego.AppConfig.String("DB_PREFIX") + str
}
