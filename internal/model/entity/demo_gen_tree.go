// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-02-17 18:44:17
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// DemoGenTree is the golang structure for table demo_gen_tree.
type DemoGenTree struct {
	Id         uint        `json:"id"         description:""`
	ParentId   uint        `json:"parentId"   description:"父级ID"`
	DemoName   string      `json:"demoName"   description:"姓名"`
	DemoAge    uint        `json:"demoAge"    description:"年龄"`
	Classes    string      `json:"classes"    description:"班级"`
	DemoBorn   *gtime.Time `json:"demoBorn"   description:"出生年月"`
	DemoGender uint        `json:"demoGender" description:"性别"`
	CreatedAt  *gtime.Time `json:"createdAt"  description:"创建日期"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  description:"修改日期"`
	DeletedAt  *gtime.Time `json:"deletedAt"  description:"删除日期"`
	CreatedBy  uint64      `json:"createdBy"  description:"创建人"`
	UpdatedBy  uint64      `json:"updatedBy"  description:"修改人"`
	DemoStatus int         `json:"demoStatus" description:"状态"`
	DemoCate   string      `json:"demoCate"   description:"分类"`
}
