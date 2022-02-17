// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-02-17 19:13:28
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// File is the golang structure for table file.
type File struct {
	Id        uint        `json:"id"        description:"自增ID"`
	Name      string      `json:"name"      description:"文件名称"`
	Src       string      `json:"src"       description:"本地文件存储路径"`
	Url       string      `json:"url"       description:"URL地址，可能为空"`
	UserId    uint        `json:"userId"    description:"操作用户"`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
}
