// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-02-17 18:44:17
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SysWebSet is the golang structure of table sys_web_set for DAO operations like Where/Data.
type SysWebSet struct {
	g.Meta     `orm:"table:sys_web_set, do:true"`
	WebId      interface{} // 主键
	WebContent interface{} // 站点信息
}