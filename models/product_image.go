package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

//ProductImage 产品图片
type ProductImage struct {
	ID              int64            `orm:"column(id);pk;auto" json:"id"` //主键
	CreateUserID    int64            `orm:"column(create_user_id);null"`  //创建者
	UpdateUserID    int64            `orm:"column(update_user_id);null"`  //最后更新者
	CreateDate      time.Time        `orm:"auto_now_add;type(datetime)"`  //创建时间
	UpdateDate      time.Time        `orm:"auto_now;type(datetime)"`      //最后更新时间
	Name            string           `orm:"unique" form:"name"`           //图片名称
	ProductTemplate *ProductTemplate `orm:"rel(fk);null"`                 //款式图片
	ProductProduct  *ProductProduct  `orm:"rel(fk);null"`                 //规格图片

	FormAction   string   `orm:"-" json:"FormAction"`   //非数据库字段，用于表示记录的增加，修改
	ActionFields []string `orm:"-" json:"ActionFields"` //需要操作的字段,用于update时
}

func init() {
	orm.RegisterModel(new(ProductImage))
}
