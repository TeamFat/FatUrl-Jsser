package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Url_20170613_005323 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Url_20170613_005323{}
	m.Created = "20170613_005323"
	migration.Register("Url_20170613_005323", m)
}

// Run the migrations
func (m *Url_20170613_005323) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	sql := "CREATE TABLE `url` ("
	sql += "`id` int(11) unsigned NOT NULL AUTO_INCREMENT,"
	sql += "`name` varchar(128) NOT NULL DEFAULT '' COMMENT '名称',"
	sql += "`source_full_url` varchar(255) NOT NULL DEFAULT '' COMMENT '完整地址',"
	sql += "`source_main_domain` varchar(255) NOT NULL DEFAULT '' COMMENT '主域',"
	sql += "`user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户ID',"
	sql += "`status` tinyint(2) NOT NULL DEFAULT '1' COMMENT '状态',"
	sql += "`last_visit_time` datetime NOT NULL COMMENT '最后访问时间',"
	sql += "`created_at` datetime NOT NULL COMMENT '创建时间',"
	sql += "PRIMARY KEY (`id`),"
	sql += "KEY `created_at_user_id_index` (`created_at`,`user_id`)"
	sql += ") ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8mb4;"
	m.SQL(sql)
}

// Reverse the migrations
func (m *Url_20170613_005323) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `url`")

}
