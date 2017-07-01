package utils

import "github.com/speps/go-hashids"
import "github.com/astaxie/beego"
import "fmt"

func HashIdEncode(id int) string {
	hd := hashids.NewData()
	hd.Salt = beego.AppConfig.String("UrlHashSalt")
	hd.MinLength, _ = beego.AppConfig.Int("UrlMinLength")
	hash, _ := hashids.NewWithData(hd)
	e, _ := hash.Encode([]int{id})
	fmt.Println(e)
	return e
}

func HashIdDecode(s string) int {
	hd := hashids.NewData()
	hd.Salt = beego.AppConfig.String("UrlHashSalt")
	hd.MinLength, _ = beego.AppConfig.Int("UrlMinLength")
	hash, _ := hashids.NewWithData(hd)
	i, _ := hash.DecodeWithError(s)
	fmt.Println(i)
	return i[0]
}
