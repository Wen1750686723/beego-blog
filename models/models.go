package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type BlogContent struct {
	Id      int64
	Content string `orm:"type(text)"`
}

func (*BlogContent) TableEngine() string {
	return engine()
}

func engine() string {
	return "INNODB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci"
}

func init() {
	orm.RegisterModelWithPrefix("bb_", new(BlogContent))
}
func Getone() string {
	return "fgoajfld"
}
func Insertblogconteng() {
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库

	user := new(BlogContent)
	user.Content = "slene"
	fmt.Println(o.Insert(user))
}
