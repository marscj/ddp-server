// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"ddp-server/internal/service/internal/dao/internal"
)

// pluginsManageDao is the data access object for table plugins_manage.
// You can define custom methods on it to extend its functionality as you wish.
type pluginsManageDao struct {
	*internal.PluginsManageDao
}

var (
	// PluginsManage is globally public accessible object for table plugins_manage operations.
	PluginsManage = pluginsManageDao{
		internal.NewPluginsManageDao(),
	}
)

// Fill with you ideas below.