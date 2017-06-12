package main

import (
	"github.com/astaxie/beego/migration"
)

const tableName = "user"

// DO NOT MODIFY
type User_20170612_233723 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &User_20170612_233723{}
	m.Created = "20170612_233723"
	migration.Register("User_20170612_233723", m)
}

// Up the migrations
func (m *User_20170612_233723) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	sql := "CREATE TABLE `user`"
	sql += " ("
	sql += " `id` int(11) unsigned NOT NULL AUTO_INCREMENT,"
	sql += " `nickname` varchar(64) NOT NULL DEFAULT '' COMMENT '昵称',"
	sql += " `username` varchar(64) NOT NULL DEFAULT '' COMMENT '用户名',"
	sql += " `password` varchar(64) NOT NULL DEFAULT '' COMMENT '密码',"
	sql += " `email` varchar(255) NOT NULL DEFAULT '' COMMENT '邮箱',"
	sql += " `mobile` char(20) NOT NULL DEFAULT '' COMMENT '手机号',"
	sql += " `avatar_url` varchar(255) NOT NULL DEFAULT '' COMMENT '头像',"
	sql += " `salt` varchar(64) NOT NULL DEFAULT '' COMMENT '盐值',"
	sql += " `password_reset_token` varchar(64) NOT NULL DEFAULT '' COMMENT '密码重置Token',"
	sql += " `register_ip` varchar(32) NOT NULL DEFAULT '' COMMENT '注册IP',"
	sql += " `created_at` datetime NOT NULL COMMENT '创建时间',"
	sql += " `last_login_time` datetime NOT NULL COMMENT '最后登录时间',"
	sql += " `last_login_ip` varchar(32) NOT NULL DEFAULT '' COMMENT '最后登录IP',"
	sql += " `status` tinyint(2) NOT NULL DEFAULT '1' COMMENT '用户状态',"
	sql += " PRIMARY KEY (`id`),"
	sql += " UNIQUE KEY `username_unique` (`username`),"
	sql += " UNIQUE KEY `email_unique` (`email`)"
	sql += " ) ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8mb4;"
	m.SQL(sql)
}

//Down Reverse the migrations
func (m *User_20170612_233723) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `user`")
}
