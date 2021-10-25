package models

import (
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Id         int32  `orm:"pk;auto" json:"id"`
	Name       string `orm:"size(30)" json:"name"`
	CertBase64 string `orm:"size(5000)" json:"cert"`
}

type Role struct {
	Id   int32  `orm:"pk;auto" json:"id"`
	Name string `orm:"size(30)" json:"name"`
}

type UserRole struct {
	Id     int32 `orm:"pk;auto" json:"id"`
	UserId int32 `json:"user_id"`
	RoleId int32 `json:"role_id"`
}

type Agent struct {
	Id   int32  `orm:"pk;auto" json:"id"`
	Name string `orm:"size(30)" json:"name"`
	//ip段.例:192.168.3.0
	IpRange string `orm:"size(39)" json:"ip_range"`
	//子网掩码.例 255.255.255.0
	//子网掩码和ip段共同标识此Agent可转发的网络
	Netmask string `orm:"size(39)" json:"netmask"`
	//此Agent的内部DNS地址,当填入后网关会将客户端的DNS请求转发到此Agent,并由此Agent完成DNS请求
	Dns        string `orm:"null;size(39)" json:"dns"`
	CertBase64 string `orm:"size(5000)" json:"cert"`
}

type Resource struct {
	Id    int32  `orm:"pk;auto" json:"id"`
	Name  string `orm:"size(30)" json:"name"`
	Ip    string `orm:"size(39)" json:"ip"`
	Port  int32  `json:"port"`
	Agent *Agent `orm:"rel(fk)" json:"agent"`
}

func init() {
	orm.RegisterDriver("sqlite3", orm.DRSqlite)

	orm.RegisterModel(new(User), new(Role), new(UserRole), new(Agent), new(Resource))

	orm.RegisterDataBase("default", "sqlite3", "./leonet.db")

	orm.RunSyncdb("default", true, true)
}
