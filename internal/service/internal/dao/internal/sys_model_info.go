// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-02-17 18:44:17
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysModelInfoDao is the data access object for table sys_model_info.
type SysModelInfoDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns SysModelInfoColumns // columns contains all the column names of Table for convenient usage.
}

// SysModelInfoColumns defines and stores column names for table sys_model_info.
type SysModelInfoColumns struct {
	ModelId         string // 模型ID
	ModelCategoryId string // 模板分类id
	ModelName       string // 模型标识
	ModelTitle      string // 模型名称
	ModelPk         string // 主键字段
	ModelOrder      string // 默认排序字段
	ModelSort       string // 表单字段排序
	ModelList       string // 列表显示字段，为空显示全部
	ModelEdit       string // 可编辑字段，为空则除主键外均可以编辑
	ModelIndexes    string // 索引字段
	SearchList      string // 高级搜索的字段
	CreateTime      string // 创建时间
	UpdateTime      string // 更新时间
	ModelStatus     string // 状态
	ModelEngine     string // 数据库引擎
	CreateBy        string // 创建人
	UpdateBy        string // 修改人
}

//  sysModelInfoColumns holds the columns for table sys_model_info.
var sysModelInfoColumns = SysModelInfoColumns{
	ModelId:         "model_id",
	ModelCategoryId: "model_category_id",
	ModelName:       "model_name",
	ModelTitle:      "model_title",
	ModelPk:         "model_pk",
	ModelOrder:      "model_order",
	ModelSort:       "model_sort",
	ModelList:       "model_list",
	ModelEdit:       "model_edit",
	ModelIndexes:    "model_indexes",
	SearchList:      "search_list",
	CreateTime:      "create_time",
	UpdateTime:      "update_time",
	ModelStatus:     "model_status",
	ModelEngine:     "model_engine",
	CreateBy:        "create_by",
	UpdateBy:        "update_by",
}

// NewSysModelInfoDao creates and returns a new DAO object for table data access.
func NewSysModelInfoDao() *SysModelInfoDao {
	return &SysModelInfoDao{
		group:   "default",
		table:   "sys_model_info",
		columns: sysModelInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysModelInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysModelInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysModelInfoDao) Columns() SysModelInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysModelInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysModelInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysModelInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}