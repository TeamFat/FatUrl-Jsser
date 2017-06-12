package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type LoginLog_20170613_012614 struct {
	migration.Migration
}

type LoginLog struct {
}

// DO NOT MODIFY
func init() {
	m := &LoginLog_20170613_012614{}
	m.Created = "20170613_012614"
	migration.Register("LoginLog_20170613_012614", m)
}

// Run the migrations
func (m *LoginLog_20170613_012614) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	sql := "CREATE TABLE `login_log` ("
	sql = "`id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,"
	sql = "`user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户ID',"
	sql = "`login_ip` varchar(11) NOT NULL DEFAULT '' COMMENT '登陆IP',"
	sql = "`created_at` datetime NOT NULL COMMENT '登陆时间',"
	sql = "`register_time` datetime NOT NULL COMMENT '注册时间',"
	sql = "PRIMARY KEY (`id`),"
	sql = "KEY `created_at_user_id_index` (`created_at`,`user_id`),"
	sql = "KEY `register_time_index` (`register_time`)"
	sql = ") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;"
	m.SQL(sql)
}

// Reverse the migrations
func (m *LoginLog_20170613_012614) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `login_log`")
}
