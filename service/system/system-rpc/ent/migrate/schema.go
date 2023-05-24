// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// SysDepartmentsColumns holds the columns for the "sys_departments" table.
	SysDepartmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeUint8, Nullable: true, Default: 1},
		{Name: "sort", Type: field.TypeUint32, Default: 1},
		{Name: "name", Type: field.TypeString, Comment: "Department name | 部门名称"},
		{Name: "ancestors", Type: field.TypeString, Comment: "Parents' IDs | 父级列表"},
		{Name: "leader", Type: field.TypeString, Comment: "Department leader | 部门负责人"},
		{Name: "phone", Type: field.TypeString, Comment: "Leader's phone number | 负责人电话"},
		{Name: "email", Type: field.TypeString, Comment: "Leader's email | 部门负责人电子邮箱"},
		{Name: "remark", Type: field.TypeString, Comment: "Remark | 备注"},
		{Name: "parent_id", Type: field.TypeUint64, Nullable: true, Comment: "Parent department ID | 父级部门ID", Default: 0},
	}
	// SysDepartmentsTable holds the schema information for the "sys_departments" table.
	SysDepartmentsTable = &schema.Table{
		Name:       "sys_departments",
		Columns:    SysDepartmentsColumns,
		PrimaryKey: []*schema.Column{SysDepartmentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sys_departments_sys_departments_children",
				Columns:    []*schema.Column{SysDepartmentsColumns[11]},
				RefColumns: []*schema.Column{SysDepartmentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SysMenusColumns holds the columns for the "sys_menus" table.
	SysMenusColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "sort", Type: field.TypeUint32, Default: 1},
		{Name: "menu_level", Type: field.TypeUint32, Comment: "Menu level | 菜单层级"},
		{Name: "menu_type", Type: field.TypeUint32, Comment: "Menu type | 菜单类型 （菜单或目录）0 目录 1 菜单"},
		{Name: "path", Type: field.TypeString, Nullable: true, Comment: "Index path | 菜单路由路径", Default: ""},
		{Name: "name", Type: field.TypeString, Comment: "Index name | 菜单名称"},
		{Name: "redirect", Type: field.TypeString, Nullable: true, Comment: "Redirect path | 跳转路径 （外链）", Default: ""},
		{Name: "component", Type: field.TypeString, Nullable: true, Comment: "The path of vue file | 组件路径", Default: ""},
		{Name: "disabled", Type: field.TypeBool, Nullable: true, Comment: "Disable status | 是否停用", Default: false},
		{Name: "title", Type: field.TypeString, Comment: "Menu name | 菜单显示标题"},
		{Name: "icon", Type: field.TypeString, Comment: "Menu icon | 菜单图标"},
		{Name: "hide_menu", Type: field.TypeBool, Nullable: true, Comment: "Hide menu | 是否隐藏菜单", Default: false},
		{Name: "hide_breadcrumb", Type: field.TypeBool, Nullable: true, Comment: "Hide the breadcrumb | 隐藏面包屑", Default: false},
		{Name: "ignore_keep_alive", Type: field.TypeBool, Nullable: true, Comment: "Do not keep alive the tab | 取消页面缓存", Default: false},
		{Name: "hide_tab", Type: field.TypeBool, Nullable: true, Comment: "Hide the tab header | 隐藏页头", Default: false},
		{Name: "frame_src", Type: field.TypeString, Nullable: true, Comment: "Show iframe | 内嵌 iframe", Default: ""},
		{Name: "carry_param", Type: field.TypeBool, Nullable: true, Comment: "The route carries parameters or not | 携带参数", Default: false},
		{Name: "hide_children_in_menu", Type: field.TypeBool, Nullable: true, Comment: "Hide children menu or not | 隐藏所有子菜单", Default: false},
		{Name: "affix", Type: field.TypeBool, Nullable: true, Comment: "Affix tab | Tab 固定", Default: false},
		{Name: "dynamic_level", Type: field.TypeUint32, Nullable: true, Comment: "The maximum number of pages the router can open | 能打开的子TAB数", Default: 20},
		{Name: "real_path", Type: field.TypeString, Nullable: true, Comment: "The real path of the route without dynamic part | 菜单路由不包含参数部分", Default: ""},
		{Name: "parent_id", Type: field.TypeUint64, Nullable: true, Comment: "Parent menu ID | 父菜单ID", Default: 100000},
	}
	// SysMenusTable holds the schema information for the "sys_menus" table.
	SysMenusTable = &schema.Table{
		Name:       "sys_menus",
		Columns:    SysMenusColumns,
		PrimaryKey: []*schema.Column{SysMenusColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sys_menus_sys_menus_children",
				Columns:    []*schema.Column{SysMenusColumns[23]},
				RefColumns: []*schema.Column{SysMenusColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SysPositionsColumns holds the columns for the "sys_positions" table.
	SysPositionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeUint8, Nullable: true, Default: 1},
		{Name: "sort", Type: field.TypeUint32, Default: 1},
		{Name: "name", Type: field.TypeString, Comment: "Position Name | 职位名称"},
		{Name: "code", Type: field.TypeString, Comment: "The code of position | 职位编码"},
		{Name: "remark", Type: field.TypeString, Comment: "Remark | 备注"},
	}
	// SysPositionsTable holds the schema information for the "sys_positions" table.
	SysPositionsTable = &schema.Table{
		Name:       "sys_positions",
		Columns:    SysPositionsColumns,
		PrimaryKey: []*schema.Column{SysPositionsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "position_code",
				Unique:  true,
				Columns: []*schema.Column{SysPositionsColumns[6]},
			},
		},
	}
	// SysRolesColumns holds the columns for the "sys_roles" table.
	SysRolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeUint8, Nullable: true, Default: 1},
		{Name: "name", Type: field.TypeString, Comment: "Role name | 角色名"},
		{Name: "code", Type: field.TypeString, Unique: true, Comment: "Role code for permission control in front end | 角色码，用于前端权限控制"},
		{Name: "default_router", Type: field.TypeString, Comment: "Default menu : dashboard | 默认登录页面", Default: "dashboard"},
		{Name: "remark", Type: field.TypeString, Comment: "Remark | 备注", Default: ""},
		{Name: "sort", Type: field.TypeUint32, Comment: "Order number | 排序编号", Default: 0},
	}
	// SysRolesTable holds the schema information for the "sys_roles" table.
	SysRolesTable = &schema.Table{
		Name:       "sys_roles",
		Columns:    SysRolesColumns,
		PrimaryKey: []*schema.Column{SysRolesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "role_code",
				Unique:  true,
				Columns: []*schema.Column{SysRolesColumns[5]},
			},
		},
	}
	// SysUsersColumns holds the columns for the "sys_users" table.
	SysUsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeUint8, Nullable: true, Default: 1},
		{Name: "username", Type: field.TypeString, Unique: true, Comment: "User's login name | 登录名"},
		{Name: "password", Type: field.TypeString, Comment: "Password | 密码"},
		{Name: "nickname", Type: field.TypeString, Unique: true, Comment: "Nickname | 昵称"},
		{Name: "description", Type: field.TypeString, Nullable: true, Comment: "The description of user | 用户的描述信息"},
		{Name: "home_path", Type: field.TypeString, Comment: "The home page that the user enters after logging in | 用户登陆后进入的首页", Default: "/dashboard"},
		{Name: "mobile", Type: field.TypeString, Nullable: true, Comment: "Mobile number | 手机号"},
		{Name: "email", Type: field.TypeString, Nullable: true, Comment: "Email | 邮箱号"},
		{Name: "avatar", Type: field.TypeString, Nullable: true, Comment: "Avatar | 头像路径", Default: "", SchemaType: map[string]string{"mysql": "varchar(512)"}},
		{Name: "department_id", Type: field.TypeUint64, Nullable: true, Comment: "Department ID | 部门ID", Default: 1},
	}
	// SysUsersTable holds the schema information for the "sys_users" table.
	SysUsersTable = &schema.Table{
		Name:       "sys_users",
		Columns:    SysUsersColumns,
		PrimaryKey: []*schema.Column{SysUsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "sys_users_sys_departments_departments",
				Columns:    []*schema.Column{SysUsersColumns[12]},
				RefColumns: []*schema.Column{SysDepartmentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "user_username_email",
				Unique:  true,
				Columns: []*schema.Column{SysUsersColumns[4], SysUsersColumns[10]},
			},
		},
	}
	// RoleMenusColumns holds the columns for the "role_menus" table.
	RoleMenusColumns = []*schema.Column{
		{Name: "role_id", Type: field.TypeUint64},
		{Name: "menu_id", Type: field.TypeUint64},
	}
	// RoleMenusTable holds the schema information for the "role_menus" table.
	RoleMenusTable = &schema.Table{
		Name:       "role_menus",
		Columns:    RoleMenusColumns,
		PrimaryKey: []*schema.Column{RoleMenusColumns[0], RoleMenusColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "role_menus_role_id",
				Columns:    []*schema.Column{RoleMenusColumns[0]},
				RefColumns: []*schema.Column{SysRolesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "role_menus_menu_id",
				Columns:    []*schema.Column{RoleMenusColumns[1]},
				RefColumns: []*schema.Column{SysMenusColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UserPositionsColumns holds the columns for the "user_positions" table.
	UserPositionsColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "position_id", Type: field.TypeUint64},
	}
	// UserPositionsTable holds the schema information for the "user_positions" table.
	UserPositionsTable = &schema.Table{
		Name:       "user_positions",
		Columns:    UserPositionsColumns,
		PrimaryKey: []*schema.Column{UserPositionsColumns[0], UserPositionsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_positions_user_id",
				Columns:    []*schema.Column{UserPositionsColumns[0]},
				RefColumns: []*schema.Column{SysUsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_positions_position_id",
				Columns:    []*schema.Column{UserPositionsColumns[1]},
				RefColumns: []*schema.Column{SysPositionsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UserRolesColumns holds the columns for the "user_roles" table.
	UserRolesColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "role_id", Type: field.TypeUint64},
	}
	// UserRolesTable holds the schema information for the "user_roles" table.
	UserRolesTable = &schema.Table{
		Name:       "user_roles",
		Columns:    UserRolesColumns,
		PrimaryKey: []*schema.Column{UserRolesColumns[0], UserRolesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_roles_user_id",
				Columns:    []*schema.Column{UserRolesColumns[0]},
				RefColumns: []*schema.Column{SysUsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_roles_role_id",
				Columns:    []*schema.Column{UserRolesColumns[1]},
				RefColumns: []*schema.Column{SysRolesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		SysDepartmentsTable,
		SysMenusTable,
		SysPositionsTable,
		SysRolesTable,
		SysUsersTable,
		RoleMenusTable,
		UserPositionsTable,
		UserRolesTable,
	}
)

func init() {
	SysDepartmentsTable.ForeignKeys[0].RefTable = SysDepartmentsTable
	SysDepartmentsTable.Annotation = &entsql.Annotation{
		Table: "sys_departments",
	}
	SysMenusTable.ForeignKeys[0].RefTable = SysMenusTable
	SysMenusTable.Annotation = &entsql.Annotation{
		Table: "sys_menus",
	}
	SysPositionsTable.Annotation = &entsql.Annotation{
		Table: "sys_positions",
	}
	SysRolesTable.Annotation = &entsql.Annotation{
		Table: "sys_roles",
	}
	SysUsersTable.ForeignKeys[0].RefTable = SysDepartmentsTable
	SysUsersTable.Annotation = &entsql.Annotation{
		Table: "sys_users",
	}
	RoleMenusTable.ForeignKeys[0].RefTable = SysRolesTable
	RoleMenusTable.ForeignKeys[1].RefTable = SysMenusTable
	UserPositionsTable.ForeignKeys[0].RefTable = SysUsersTable
	UserPositionsTable.ForeignKeys[1].RefTable = SysPositionsTable
	UserRolesTable.ForeignKeys[0].RefTable = SysUsersTable
	UserRolesTable.ForeignKeys[1].RefTable = SysRolesTable
}
