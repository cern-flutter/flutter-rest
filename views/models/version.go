package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type SchemaVersion struct {
	Major int `orm:"pk"`
	Minor int
	Patch int
	Message string
}

func (sv *SchemaVersion) TableName() string {
	return "t_schema_version"
}
func (sv *SchemaVersion) TableCPK() []string {
	return []string{"Major", "Minor", "Patch"}
}

func (sv SchemaVersion) String() string {
   return fmt.Sprintf("Version <%d %d %d %s>", sv.Major, sv.Minor, sv.Patch, sv.Message)
}

func init() {
	orm.RegisterModel(new(SchemaVersion))
}
