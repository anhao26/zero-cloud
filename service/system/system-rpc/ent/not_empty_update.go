// Code generated by ent, DO NOT EDIT.

package ent

// set field if value is not empty. e.g. string does not equal to ""
func (d *DepartmentUpdate) SetNotEmptyStatus(value uint8) *DepartmentUpdate {
	if value != 0 {
		return d.SetStatus(value)
	}
	return d
}

// set field if value is not empty. e.g. string does not equal to ""
func (d *DepartmentUpdateOne) SetNotEmptyStatus(value uint8) *DepartmentUpdateOne {
	if value != 0 {
		return d.SetStatus(value)
	}
	return d
}

// set field if value is not empty. e.g. string does not equal to ""
func (d *DepartmentUpdate) SetNotEmptySort(value uint32) *DepartmentUpdate {
	if value != 0 {
		return d.SetSort(value)
	}
	return d
}

// set field if value is not empty. e.g. string does not equal to ""
func (d *DepartmentUpdateOne) SetNotEmptySort(value uint32) *DepartmentUpdateOne {
	if value != 0 {
		return d.SetSort(value)
	}
	return d
}

// set field if value is not empty. e.g. string does not equal to ""
func (d *DepartmentUpdate) SetNotEmptyName(value string) *DepartmentUpdate {
	if value != "" {
		return d.SetName(value)
	}
	return d
}

// set field if value is not empty. e.g. string does not equal to ""
func (d *DepartmentUpdateOne) SetNotEmptyName(value string) *DepartmentUpdateOne {
	if value != "" {
		return d.SetName(value)
	}
	return d
}

// set field if value is not empty. e.g. string does not equal to ""
func (d *DepartmentUpdate) SetNotEmptyAncestors(value string) *DepartmentUpdate {
	if value != "" {
		return d.SetAncestors(value)
	}
	return d
}

// set field if value is not empty. e.g. string does not equal to ""
func (d *DepartmentUpdateOne) SetNotEmptyAncestors(value string) *DepartmentUpdateOne {
	if value != "" {
		return d.SetAncestors(value)
	}
	return d
}

// set field if value is not empty. e.g. string does not equal to ""
func (d *DepartmentUpdate) SetNotEmptyLeader(value string) *DepartmentUpdate {
	if value != "" {
		return d.SetLeader(value)
	}
	return d
}

// set field if value is not empty. e.g. string does not equal to ""
func (d *DepartmentUpdateOne) SetNotEmptyLeader(value string) *DepartmentUpdateOne {
	if value != "" {
		return d.SetLeader(value)
	}
	return d
}

// set field if value is not empty. e.g. string does not equal to ""
func (d *DepartmentUpdate) SetNotEmptyPhone(value string) *DepartmentUpdate {
	if value != "" {
		return d.SetPhone(value)
	}
	return d
}

// set field if value is not empty. e.g. string does not equal to ""
func (d *DepartmentUpdateOne) SetNotEmptyPhone(value string) *DepartmentUpdateOne {
	if value != "" {
		return d.SetPhone(value)
	}
	return d
}

// set field if value is not empty. e.g. string does not equal to ""
func (d *DepartmentUpdate) SetNotEmptyEmail(value string) *DepartmentUpdate {
	if value != "" {
		return d.SetEmail(value)
	}
	return d
}

// set field if value is not empty. e.g. string does not equal to ""
func (d *DepartmentUpdateOne) SetNotEmptyEmail(value string) *DepartmentUpdateOne {
	if value != "" {
		return d.SetEmail(value)
	}
	return d
}

// set field if value is not empty. e.g. string does not equal to ""
func (d *DepartmentUpdate) SetNotEmptyRemark(value string) *DepartmentUpdate {
	if value != "" {
		return d.SetRemark(value)
	}
	return d
}

// set field if value is not empty. e.g. string does not equal to ""
func (d *DepartmentUpdateOne) SetNotEmptyRemark(value string) *DepartmentUpdateOne {
	if value != "" {
		return d.SetRemark(value)
	}
	return d
}

// set field if value is not empty. e.g. string does not equal to ""
func (d *DepartmentUpdate) SetNotEmptyParentID(value uint64) *DepartmentUpdate {
	if value != 0 {
		return d.SetParentID(value)
	}
	return d
}

// set field if value is not empty. e.g. string does not equal to ""
func (d *DepartmentUpdateOne) SetNotEmptyParentID(value uint64) *DepartmentUpdateOne {
	if value != 0 {
		return d.SetParentID(value)
	}
	return d
}

// set field if value is not empty. e.g. string does not equal to ""
func (m *MenuUpdate) SetNotEmptySort(value uint32) *MenuUpdate {
	if value != 0 {
		return m.SetSort(value)
	}
	return m
}

// set field if value is not empty. e.g. string does not equal to ""
func (m *MenuUpdateOne) SetNotEmptySort(value uint32) *MenuUpdateOne {
	if value != 0 {
		return m.SetSort(value)
	}
	return m
}

// set field if value is not empty. e.g. string does not equal to ""
func (m *MenuUpdate) SetNotEmptyParentID(value uint64) *MenuUpdate {
	if value != 0 {
		return m.SetParentID(value)
	}
	return m
}

// set field if value is not empty. e.g. string does not equal to ""
func (m *MenuUpdateOne) SetNotEmptyParentID(value uint64) *MenuUpdateOne {
	if value != 0 {
		return m.SetParentID(value)
	}
	return m
}

// set field if value is not empty. e.g. string does not equal to ""
func (m *MenuUpdate) SetNotEmptyMenuLevel(value uint32) *MenuUpdate {
	if value != 0 {
		return m.SetMenuLevel(value)
	}
	return m
}

// set field if value is not empty. e.g. string does not equal to ""
func (m *MenuUpdateOne) SetNotEmptyMenuLevel(value uint32) *MenuUpdateOne {
	if value != 0 {
		return m.SetMenuLevel(value)
	}
	return m
}

// set field if value is not empty. e.g. string does not equal to ""
func (m *MenuUpdate) SetNotEmptyMenuType(value uint32) *MenuUpdate {
	if value != 0 {
		return m.SetMenuType(value)
	}
	return m
}

// set field if value is not empty. e.g. string does not equal to ""
func (m *MenuUpdateOne) SetNotEmptyMenuType(value uint32) *MenuUpdateOne {
	if value != 0 {
		return m.SetMenuType(value)
	}
	return m
}

// set field if value is not empty. e.g. string does not equal to ""
func (m *MenuUpdate) SetNotEmptyPath(value string) *MenuUpdate {
	if value != "" {
		return m.SetPath(value)
	}
	return m
}

// set field if value is not empty. e.g. string does not equal to ""
func (m *MenuUpdateOne) SetNotEmptyPath(value string) *MenuUpdateOne {
	if value != "" {
		return m.SetPath(value)
	}
	return m
}

// set field if value is not empty. e.g. string does not equal to ""
func (m *MenuUpdate) SetNotEmptyName(value string) *MenuUpdate {
	if value != "" {
		return m.SetName(value)
	}
	return m
}

// set field if value is not empty. e.g. string does not equal to ""
func (m *MenuUpdateOne) SetNotEmptyName(value string) *MenuUpdateOne {
	if value != "" {
		return m.SetName(value)
	}
	return m
}

// set field if value is not empty. e.g. string does not equal to ""
func (m *MenuUpdate) SetNotEmptyRedirect(value string) *MenuUpdate {
	if value != "" {
		return m.SetRedirect(value)
	}
	return m
}

// set field if value is not empty. e.g. string does not equal to ""
func (m *MenuUpdateOne) SetNotEmptyRedirect(value string) *MenuUpdateOne {
	if value != "" {
		return m.SetRedirect(value)
	}
	return m
}

// set field if value is not empty. e.g. string does not equal to ""
func (m *MenuUpdate) SetNotEmptyComponent(value string) *MenuUpdate {
	if value != "" {
		return m.SetComponent(value)
	}
	return m
}

// set field if value is not empty. e.g. string does not equal to ""
func (m *MenuUpdateOne) SetNotEmptyComponent(value string) *MenuUpdateOne {
	if value != "" {
		return m.SetComponent(value)
	}
	return m
}

// set field if value is not empty. e.g. string does not equal to ""
func (m *MenuUpdate) SetNotEmptyTitle(value string) *MenuUpdate {
	if value != "" {
		return m.SetTitle(value)
	}
	return m
}

// set field if value is not empty. e.g. string does not equal to ""
func (m *MenuUpdateOne) SetNotEmptyTitle(value string) *MenuUpdateOne {
	if value != "" {
		return m.SetTitle(value)
	}
	return m
}

// set field if value is not empty. e.g. string does not equal to ""
func (m *MenuUpdate) SetNotEmptyIcon(value string) *MenuUpdate {
	if value != "" {
		return m.SetIcon(value)
	}
	return m
}

// set field if value is not empty. e.g. string does not equal to ""
func (m *MenuUpdateOne) SetNotEmptyIcon(value string) *MenuUpdateOne {
	if value != "" {
		return m.SetIcon(value)
	}
	return m
}

// set field if value is not empty. e.g. string does not equal to ""
func (m *MenuUpdate) SetNotEmptyFrameSrc(value string) *MenuUpdate {
	if value != "" {
		return m.SetFrameSrc(value)
	}
	return m
}

// set field if value is not empty. e.g. string does not equal to ""
func (m *MenuUpdateOne) SetNotEmptyFrameSrc(value string) *MenuUpdateOne {
	if value != "" {
		return m.SetFrameSrc(value)
	}
	return m
}

// set field if value is not empty. e.g. string does not equal to ""
func (m *MenuUpdate) SetNotEmptyDynamicLevel(value uint32) *MenuUpdate {
	if value != 0 {
		return m.SetDynamicLevel(value)
	}
	return m
}

// set field if value is not empty. e.g. string does not equal to ""
func (m *MenuUpdateOne) SetNotEmptyDynamicLevel(value uint32) *MenuUpdateOne {
	if value != 0 {
		return m.SetDynamicLevel(value)
	}
	return m
}

// set field if value is not empty. e.g. string does not equal to ""
func (m *MenuUpdate) SetNotEmptyRealPath(value string) *MenuUpdate {
	if value != "" {
		return m.SetRealPath(value)
	}
	return m
}

// set field if value is not empty. e.g. string does not equal to ""
func (m *MenuUpdateOne) SetNotEmptyRealPath(value string) *MenuUpdateOne {
	if value != "" {
		return m.SetRealPath(value)
	}
	return m
}

// set field if value is not empty. e.g. string does not equal to ""
func (po *PositionUpdate) SetNotEmptyStatus(value uint8) *PositionUpdate {
	if value != 0 {
		return po.SetStatus(value)
	}
	return po
}

// set field if value is not empty. e.g. string does not equal to ""
func (po *PositionUpdateOne) SetNotEmptyStatus(value uint8) *PositionUpdateOne {
	if value != 0 {
		return po.SetStatus(value)
	}
	return po
}

// set field if value is not empty. e.g. string does not equal to ""
func (po *PositionUpdate) SetNotEmptySort(value uint32) *PositionUpdate {
	if value != 0 {
		return po.SetSort(value)
	}
	return po
}

// set field if value is not empty. e.g. string does not equal to ""
func (po *PositionUpdateOne) SetNotEmptySort(value uint32) *PositionUpdateOne {
	if value != 0 {
		return po.SetSort(value)
	}
	return po
}

// set field if value is not empty. e.g. string does not equal to ""
func (po *PositionUpdate) SetNotEmptyName(value string) *PositionUpdate {
	if value != "" {
		return po.SetName(value)
	}
	return po
}

// set field if value is not empty. e.g. string does not equal to ""
func (po *PositionUpdateOne) SetNotEmptyName(value string) *PositionUpdateOne {
	if value != "" {
		return po.SetName(value)
	}
	return po
}

// set field if value is not empty. e.g. string does not equal to ""
func (po *PositionUpdate) SetNotEmptyCode(value string) *PositionUpdate {
	if value != "" {
		return po.SetCode(value)
	}
	return po
}

// set field if value is not empty. e.g. string does not equal to ""
func (po *PositionUpdateOne) SetNotEmptyCode(value string) *PositionUpdateOne {
	if value != "" {
		return po.SetCode(value)
	}
	return po
}

// set field if value is not empty. e.g. string does not equal to ""
func (po *PositionUpdate) SetNotEmptyRemark(value string) *PositionUpdate {
	if value != "" {
		return po.SetRemark(value)
	}
	return po
}

// set field if value is not empty. e.g. string does not equal to ""
func (po *PositionUpdateOne) SetNotEmptyRemark(value string) *PositionUpdateOne {
	if value != "" {
		return po.SetRemark(value)
	}
	return po
}

// set field if value is not empty. e.g. string does not equal to ""
func (r *RoleUpdate) SetNotEmptyStatus(value uint8) *RoleUpdate {
	if value != 0 {
		return r.SetStatus(value)
	}
	return r
}

// set field if value is not empty. e.g. string does not equal to ""
func (r *RoleUpdateOne) SetNotEmptyStatus(value uint8) *RoleUpdateOne {
	if value != 0 {
		return r.SetStatus(value)
	}
	return r
}

// set field if value is not empty. e.g. string does not equal to ""
func (r *RoleUpdate) SetNotEmptyName(value string) *RoleUpdate {
	if value != "" {
		return r.SetName(value)
	}
	return r
}

// set field if value is not empty. e.g. string does not equal to ""
func (r *RoleUpdateOne) SetNotEmptyName(value string) *RoleUpdateOne {
	if value != "" {
		return r.SetName(value)
	}
	return r
}

// set field if value is not empty. e.g. string does not equal to ""
func (r *RoleUpdate) SetNotEmptyCode(value string) *RoleUpdate {
	if value != "" {
		return r.SetCode(value)
	}
	return r
}

// set field if value is not empty. e.g. string does not equal to ""
func (r *RoleUpdateOne) SetNotEmptyCode(value string) *RoleUpdateOne {
	if value != "" {
		return r.SetCode(value)
	}
	return r
}

// set field if value is not empty. e.g. string does not equal to ""
func (r *RoleUpdate) SetNotEmptyDefaultRouter(value string) *RoleUpdate {
	if value != "" {
		return r.SetDefaultRouter(value)
	}
	return r
}

// set field if value is not empty. e.g. string does not equal to ""
func (r *RoleUpdateOne) SetNotEmptyDefaultRouter(value string) *RoleUpdateOne {
	if value != "" {
		return r.SetDefaultRouter(value)
	}
	return r
}

// set field if value is not empty. e.g. string does not equal to ""
func (r *RoleUpdate) SetNotEmptyRemark(value string) *RoleUpdate {
	if value != "" {
		return r.SetRemark(value)
	}
	return r
}

// set field if value is not empty. e.g. string does not equal to ""
func (r *RoleUpdateOne) SetNotEmptyRemark(value string) *RoleUpdateOne {
	if value != "" {
		return r.SetRemark(value)
	}
	return r
}

// set field if value is not empty. e.g. string does not equal to ""
func (r *RoleUpdate) SetNotEmptySort(value uint32) *RoleUpdate {
	if value != 0 {
		return r.SetSort(value)
	}
	return r
}

// set field if value is not empty. e.g. string does not equal to ""
func (r *RoleUpdateOne) SetNotEmptySort(value uint32) *RoleUpdateOne {
	if value != 0 {
		return r.SetSort(value)
	}
	return r
}

// set field if value is not empty. e.g. string does not equal to ""
func (u *UserUpdate) SetNotEmptyStatus(value uint8) *UserUpdate {
	if value != 0 {
		return u.SetStatus(value)
	}
	return u
}

// set field if value is not empty. e.g. string does not equal to ""
func (u *UserUpdateOne) SetNotEmptyStatus(value uint8) *UserUpdateOne {
	if value != 0 {
		return u.SetStatus(value)
	}
	return u
}

// set field if value is not empty. e.g. string does not equal to ""
func (u *UserUpdate) SetNotEmptyUsername(value string) *UserUpdate {
	if value != "" {
		return u.SetUsername(value)
	}
	return u
}

// set field if value is not empty. e.g. string does not equal to ""
func (u *UserUpdateOne) SetNotEmptyUsername(value string) *UserUpdateOne {
	if value != "" {
		return u.SetUsername(value)
	}
	return u
}

// set field if value is not empty. e.g. string does not equal to ""
func (u *UserUpdate) SetNotEmptyPassword(value string) *UserUpdate {
	if value != "" {
		return u.SetPassword(value)
	}
	return u
}

// set field if value is not empty. e.g. string does not equal to ""
func (u *UserUpdateOne) SetNotEmptyPassword(value string) *UserUpdateOne {
	if value != "" {
		return u.SetPassword(value)
	}
	return u
}

// set field if value is not empty. e.g. string does not equal to ""
func (u *UserUpdate) SetNotEmptyNickname(value string) *UserUpdate {
	if value != "" {
		return u.SetNickname(value)
	}
	return u
}

// set field if value is not empty. e.g. string does not equal to ""
func (u *UserUpdateOne) SetNotEmptyNickname(value string) *UserUpdateOne {
	if value != "" {
		return u.SetNickname(value)
	}
	return u
}

// set field if value is not empty. e.g. string does not equal to ""
func (u *UserUpdate) SetNotEmptyDescription(value string) *UserUpdate {
	if value != "" {
		return u.SetDescription(value)
	}
	return u
}

// set field if value is not empty. e.g. string does not equal to ""
func (u *UserUpdateOne) SetNotEmptyDescription(value string) *UserUpdateOne {
	if value != "" {
		return u.SetDescription(value)
	}
	return u
}

// set field if value is not empty. e.g. string does not equal to ""
func (u *UserUpdate) SetNotEmptyHomePath(value string) *UserUpdate {
	if value != "" {
		return u.SetHomePath(value)
	}
	return u
}

// set field if value is not empty. e.g. string does not equal to ""
func (u *UserUpdateOne) SetNotEmptyHomePath(value string) *UserUpdateOne {
	if value != "" {
		return u.SetHomePath(value)
	}
	return u
}

// set field if value is not empty. e.g. string does not equal to ""
func (u *UserUpdate) SetNotEmptyMobile(value string) *UserUpdate {
	if value != "" {
		return u.SetMobile(value)
	}
	return u
}

// set field if value is not empty. e.g. string does not equal to ""
func (u *UserUpdateOne) SetNotEmptyMobile(value string) *UserUpdateOne {
	if value != "" {
		return u.SetMobile(value)
	}
	return u
}

// set field if value is not empty. e.g. string does not equal to ""
func (u *UserUpdate) SetNotEmptyEmail(value string) *UserUpdate {
	if value != "" {
		return u.SetEmail(value)
	}
	return u
}

// set field if value is not empty. e.g. string does not equal to ""
func (u *UserUpdateOne) SetNotEmptyEmail(value string) *UserUpdateOne {
	if value != "" {
		return u.SetEmail(value)
	}
	return u
}

// set field if value is not empty. e.g. string does not equal to ""
func (u *UserUpdate) SetNotEmptyAvatar(value string) *UserUpdate {
	if value != "" {
		return u.SetAvatar(value)
	}
	return u
}

// set field if value is not empty. e.g. string does not equal to ""
func (u *UserUpdateOne) SetNotEmptyAvatar(value string) *UserUpdateOne {
	if value != "" {
		return u.SetAvatar(value)
	}
	return u
}

// set field if value is not empty. e.g. string does not equal to ""
func (u *UserUpdate) SetNotEmptyDepartmentID(value uint64) *UserUpdate {
	if value != 0 {
		return u.SetDepartmentID(value)
	}
	return u
}

// set field if value is not empty. e.g. string does not equal to ""
func (u *UserUpdateOne) SetNotEmptyDepartmentID(value uint64) *UserUpdateOne {
	if value != 0 {
		return u.SetDepartmentID(value)
	}
	return u
}
