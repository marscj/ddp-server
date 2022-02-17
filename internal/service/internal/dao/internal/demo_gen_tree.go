// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-02-17 18:44:17
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DemoGenTreeDao is the data access object for table demo_gen_tree.
type DemoGenTreeDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns DemoGenTreeColumns // columns contains all the column names of Table for convenient usage.
}

// DemoGenTreeColumns defines and stores column names for table demo_gen_tree.
type DemoGenTreeColumns struct {
	Id         string //
	ParentId   string // 父级ID
	DemoName   string // 姓名
	DemoAge    string // 年龄
	Classes    string // 班级
	DemoBorn   string // 出生年月
	DemoGender string // 性别
	CreatedAt  string // 创建日期
	UpdatedAt  string // 修改日期
	DeletedAt  string // 删除日期
	CreatedBy  string // 创建人
	UpdatedBy  string // 修改人
	DemoStatus string // 状态
	DemoCate   string // 分类
}

//  demoGenTreeColumns holds the columns for table demo_gen_tree.
var demoGenTreeColumns = DemoGenTreeColumns{
	Id:         "id",
	ParentId:   "parent_id",
	DemoName:   "demo_name",
	DemoAge:    "demo_age",
	Classes:    "classes",
	DemoBorn:   "demo_born",
	DemoGender: "demo_gender",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	DeletedAt:  "deleted_at",
	CreatedBy:  "created_by",
	UpdatedBy:  "updated_by",
	DemoStatus: "demo_status",
	DemoCate:   "demo_cate",
}

// NewDemoGenTreeDao creates and returns a new DAO object for table data access.
func NewDemoGenTreeDao() *DemoGenTreeDao {
	return &DemoGenTreeDao{
		group:   "default",
		table:   "demo_gen_tree",
		columns: demoGenTreeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *DemoGenTreeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *DemoGenTreeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *DemoGenTreeDao) Columns() DemoGenTreeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *DemoGenTreeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *DemoGenTreeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *DemoGenTreeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}