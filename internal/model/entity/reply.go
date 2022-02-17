// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-02-17 19:13:28
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Reply is the golang structure for table reply.
type Reply struct {
	Id         uint        `json:"id"         description:"回复ID"`
	ParentId   uint        `json:"parentId"   description:"回复对应的上一级回复ID(没有的话默认为0)"`
	Title      string      `json:"title"      description:"回复标题"`
	Content    string      `json:"content"    description:"回复内容"`
	TargetType string      `json:"targetType" description:"评论类型: content, reply"`
	TargetId   uint        `json:"targetId"   description:"对应内容ID，可能回复的是另一个回复，所以这里没有使用content_id"`
	UserId     uint        `json:"userId"     description:"网站用户ID"`
	ZanCount   uint        `json:"zanCount"   description:"赞"`
	CaiCount   uint        `json:"caiCount"   description:"踩"`
	CreatedAt  *gtime.Time `json:"createdAt"  description:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  description:""`
}
