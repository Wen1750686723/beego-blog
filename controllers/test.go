package controllers

import (
	"fmt"
	//	"myproject/g"
	//	"myproject/models"
	//	"time"

	"github.com/astaxie/beego"
)

type TestsController struct {
	beego.Controller
}

func (c *TestsController) Get() {
	fmt.Print("——ss")
	v := c.GetSession("asta")

	if v == nil {
		c.SetSession("asta", "sss")
		c.Ctx.WriteString("setchenggong")
	} else {
		fmt.Println(v)
		c.Ctx.WriteString(v.(string))
	}
	//	insertblogconteng()
	//	models.Insertblogconteng()
	//	g.Cache.Put("astaxie", 1, 10*time.Second)
	//	fmt.Println(g.Cache.Get("astaxie"))
	//	c.Ctx.WriteString("sss")
	//	c.Data["Website"] = "beego.me"
	//	c.Data["Email"] = "astaxie@gmail.com"
	//	c.TplName = "index.tpl"
}
