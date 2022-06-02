/*
 * Copyright 2022. The CICD-Tools Authors
 * This program is free software: you can redistribute it and/or
 * modify it under the terms of the GNU General Public License as
 * published by the Free Software Foundation, either version 3 of
 * the License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program. If not, see <https://www.gnu.org/licenses/>.
 */

package model

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strconv"
)

type User struct {
	gorm.Model
	Name        string        `gorm:"column:user_name;type:varchar(30);not null"`
	FullName    string        `gorm:"column:full_name;type:varchar(30)"`
	Gender      string        `gorm:"column:gender;type:varchar(10)"`
	Age         uint          `gorm:"column:age;type:integer"`
	Location    string        `gorm:"column:location;type:varchar(255)"`
	Job         string        `gorm:"column:job;type:varchar(30)"`
	Password    string        `gorm:"column:password;type:varchar(128)"`
	Email       string        `gorm:"column:email_address;type:varchar(90);unique;not null"`
	Mobile      string        `gorm:"column:mobile;type:varchar(30)"`
	DingTalkID  string        `gorm:"column:dingtalk_id;type:varchar(30)"`
	WXWorkID    string        `gorm:"column:wxwork_id;type:varchar(30)"`
	Groups      *[]Group      `gorm:"-"`
	Roles       *[]Role       `gorm:"-"`
	Permissions *[]Permission `gorm:"-"`
	UserGroup   *UserGroup    `gorm:"-"`
	UserRole    *UserRole     `gorm:"-"`
	Error       error         `gorm:"-"`
}

type Group struct {
	gorm.Model
	Name        string        `gorm:"column:group_name;type:varchar(60);unique;not null"`
	Intro       string        `gorm:"column:intro;type:varchar(128)"`
	Users       *[]User       `gorm:"-"`
	Roles       *[]Role       `gorm:"-"`
	Permissions *[]Permission `gorm:"-"`
	UserGroup   *UserGroup    `gorm:"-"`
	GroupRole   *GroupRole    `gorm:"-"`
	Error       error         `gorm:"-"`
}

type Role struct {
	gorm.Model
	Name        string        `gorm:"column:role;type:varchar(60);not null"`
	Intro       string        `gorm:"column:intro;type:varchar(128)"`
	Users       *[]User       `gorm:"-"`
	Groups      *[]Group      `gorm:"-"`
	Permissions *[]Permission `gorm:"-"`
	UserRole    *UserRole     `gorm:"-"`
	GroupRole   *GroupRole    `gorm:"-"`
	Error       error         `gorm:"-"`
}

type Permission struct {
	gorm.Model
	Name       string `gorm:"column:permission;type:varchar(60);not null"`
	ResourceID uint   `gorm:"column:resource_id;type:integer;not null;<-:create"`
	Category   string `gorm:"column:category;type:varchar(128);not null"`
	Action     string `gorm:"column:action;type:varchar(60);not null"`
	RoleID     uint   `gorm:"column:role_id;type:integer;not null;<-:create"`
	Error      error  `gorm:"-"`
}

type UserGroup struct {
	gorm.Model
	UserID  uint  `gorm:"column:user_id;type:integer;<-:create"`
	GroupID uint  `gorm:"column:group_id;type:integer;<-:create"`
	Error   error `gorm:"-"`
}

type UserRole struct {
	gorm.Model
	UserID uint  `gorm:"column:user_id;type:integer;<-:create"`
	RoleID uint  `gorm:"column:role_id;type:integer;<-:create"`
	Error  error `gorm:"-"`
}

type GroupRole struct {
	gorm.Model
	GroupID uint  `gorm:"column:group_id;type:integer;<-:create"`
	RoleID  uint  `gorm:"column:group_id;type:integer;<-:create"`
	Error   error `gorm:"-"`
}

func (u *User) Exists() bool {
	if result := db.Where(u).First(u); errors.Is(result.Error, gorm.ErrRecordNotFound) {
		u.Error = gorm.ErrRecordNotFound
		return false
	} else if result.RowsAffected > 1 {
		u.Error = errors.New("查询到存在重复用户, 需添加额外字段")
		return false
	}
	return true
}

func (u *User) Find() *User {
	if result := db.Where(u).First(u); errors.Is(result.Error, gorm.ErrRecordNotFound) {
		u.Error = gorm.ErrRecordNotFound
	}
	return u
}

func (u *User) Create() *User {
	if err := db.Where(u).FirstOrCreate(u).Error; err != nil {
		u.Error = fmt.Errorf("创建用户%s失败\n%w", u.Name, err)
	}
	return u
}

func (u *User) Update() *User {
	if err := db.Save(u).Error; err != nil {
		u.Error = fmt.Errorf("用户%v数据更新失败\n%w", u.Name, err)
	}
	return u
}

func (u *User) AddGroups(names ...string) {
	for _, value := range names {
		g := new(Group)
		g.Name = value
		if g.Exists() {
			err := u.UserGroup.AddRow(u.ID, g.ID)
			if err != nil {
				u.Error = fmt.Errorf("用户%s添加到组%s时发生错误\n%w", u.Name, g.Name, err)
				return
			}
		}
	}
}

func (u *User) GetGroups() *User {
	u.Groups = u.UserGroup.GetGroups(u.ID)
	if u.UserGroup.Error != nil {
		u.Error = fmt.Errorf("用户%s查询组时发生错误\n%w", u.Name, u.UserGroup.Error)
	}
	return u
}

func (u *User) GetRoles() *User {
	u.Roles = u.UserRole.GetRoles(u.ID)
	if u.UserRole.Error != nil {
		u.Error = fmt.Errorf("用户%s查询角色时发生错误\n%w", u.Name, u.UserRole.Error)
	}
	return u
}

func (u *User) Printf() {
	fmt.Printf(`用户:
	用户名: %v
	中文名: %v
	性别: %v
	年龄: %v
	地点: %v
	职位: %v
	邮箱: %v
	手机号: %v
	钉钉: %v
	企业微信: %v
`, u.Name, u.FullName, u.Gender, u.Age,
		u.Location, u.Job, u.Email, u.Mobile,
		u.DingTalkID, u.WXWorkID,
	)
}

func (u *User) Map() map[string]string {
	m := map[string]string{
		"user_name":   u.Name,
		"full_name":   u.FullName,
		"gender":      u.Gender,
		"age":         strconv.FormatUint(uint64(u.Age), 10),
		"location":    u.Location,
		"job":         u.Job,
		"email":       u.Email,
		"mobile":      u.Mobile,
		"dingtalk_id": u.DingTalkID,
		"wxwork_id":   u.WXWorkID,
	}
	return m
}

func (g *Group) Exists() bool {
	if result := db.Where(g).First(g); errors.Is(result.Error, gorm.ErrRecordNotFound) {
		g.Error = gorm.ErrRecordNotFound
		return false
	}
	return true
}

func (g *Group) Find() *Group {
	if result := db.Where(g).First(g); errors.Is(result.Error, gorm.ErrRecordNotFound) {
		g.Error = gorm.ErrRecordNotFound
	}
	return g
}

func (g *Group) Create() *Group {
	if err := db.Where(g).FirstOrCreate(g).Error; err != nil {
		g.Error = fmt.Errorf("创建组%s失败\n%w", g.Name, err)
	}
	return g
}

func (g *Group) Update() *Group {
	if err := db.Save(g).Error; err != nil {
		g.Error = fmt.Errorf("组%s数据更新失败\n%w", g.Name, err)
	}

	return g
}

func (g *Group) AddUsers(names ...string) {
	for _, value := range names {
		u := new(User)
		u.Name = value
		if u.Exists() {
			err := u.UserGroup.AddRow(u.ID, g.ID)
			if err != nil {
				g.Error = fmt.Errorf("组%s添加用户%s时发生错误\n%w", g.Name, u.Name, err)
				return
			}
		}
	}
}

func (g *Group) GetUsers() *Group {
	g.Users = g.UserGroup.GetUsers(g.ID)
	if g.UserGroup.Error != nil {
		g.Error = fmt.Errorf("组%s查询用户时发生错误\n%w", g.Name, g.UserGroup.Error)
	}
	return g
}

func (g *Group) GetRoles() *Group {
	g.Roles = g.GroupRole.GetRoles(g.ID)
	if g.GroupRole.Error != nil {
		g.Error = fmt.Errorf("组%s查询角色时发生错误\n%w", g.Name, g.GroupRole.Error)
	}
	return g
}

func (g *Group) Printf() {
	fmt.Printf(`组:
    组名: %v
	描述: %v
`, g.Name, g.Intro)
}

func (g *Group) Map() map[string]string {
	m := map[string]string{
		"group_name": g.Name,
		"intro":      g.Intro,
	}
	return m
}

func (r *Role) Exists() bool {
	if result := db.Where(r).First(r); errors.Is(result.Error, gorm.ErrRecordNotFound) {
		r.Error = gorm.ErrRecordNotFound
		return false
	}
	return true
}

func (r *Role) Find() *Role {
	if result := db.Where(r).First(r); errors.Is(result.Error, gorm.ErrRecordNotFound) {
		r.Error = gorm.ErrRecordNotFound
	}
	return r
}

func (r *Role) Create() *Role {
	if err := db.Where(r).FirstOrCreate(r).Error; err != nil {
		r.Error = fmt.Errorf("角色%v创建失败\n%w", r.Name, err)
	}
	return r
}

func (r *Role) Update() *Role {
	if r.Exists() {
		if err := db.Save(r).Error; err != nil {
			r.Error = fmt.Errorf("角色%v更新失败\n%w", r.Name, err)
		}
	} else {
		r.Error = fmt.Errorf("角色%v不存在\n%w", r.Name, r.Error)
	}
	return r
}

func (r *Role) GetUsers() *Role {
	r.Users = r.UserRole.GetUsers(r.ID)
	if r.UserRole.Error != nil {
		r.Error = fmt.Errorf("角色%s查询用户时发生错误\n%w", r.Name, r.UserRole.Error)
	}
	return r
}

func (r *Role) GetGroups() *Role {
	r.Groups = r.GroupRole.GetGroups(r.ID)
	if r.GroupRole.Error != nil {
		r.Error = fmt.Errorf("角色%s查询组时发生错误\n%w", r.Name, r.GroupRole.Error)
	}
	return r
}

func (r *Role) Printf() {
	fmt.Printf(`角色:
	角色名称: %s
	角色描述: %s
`, r.Name, r.Intro)
}

func (r *Role) Map() map[string]string {
	m := map[string]string{
		"role_name": r.Name,
		"intro":     r.Intro,
	}
	return m
}

func (ug *UserGroup) Exists(uid uint, gid uint) bool {
	ug.UserID = uid
	ug.GroupID = gid
	if result := db.Where(ug).First(ug); errors.Is(result.Error, gorm.ErrRecordNotFound) {
		ug.Error = gorm.ErrRecordNotFound
		return false
	}
	return true
}

func (ug *UserGroup) GetGroups(uid uint) (groups *[]Group) {
	gids := new([]uint)
	if err := db.Where("user_id = ?", uid).Find(ug).Limit(1000).Error; err != nil {
		ug.Error = fmt.Errorf("用户分组表中没有查询到user_id为%d的数据\n%w", uid, err)
		return nil
	}

	if err := db.Model(ug).Pluck("group_id", gids).Error; err != nil {
		ug.Error = fmt.Errorf("用户分组表中查询group_id集合失败\n%w", err)
		return nil
	} else if len(*gids) == 0 {
		ug.Error = errors.New(fmt.Sprintf("user_id为%d的用户没有添加任何组", uid))
		return nil
	}

	if err := db.Where(gids).Find(groups).Limit(1000).Error; err != nil {
		ug.Error = fmt.Errorf("通过group_id查询对应组时失败\n%w", err)
		return nil
	}

	return
}

func (ug *UserGroup) GetUsers(gid uint) (users *[]User) {
	uids := new([]uint)
	if err := db.Where("group_id = ?", gid).Find(ug).Limit(1000).Error; err != nil {
		ug.Error = fmt.Errorf("用户分组表中没有查询到group_id为%d的数据\n%w", gid, err)
		return nil
	}

	if err := db.Model(ug).Pluck("user_id", uids).Error; err != nil {
		ug.Error = fmt.Errorf("用户分组表中查询uid集合失败\n%w", err)
		return nil
	} else if len(*uids) == 0 {
		ug.Error = errors.New(fmt.Sprintf("group_id为%d的组中没有添加任何用户", gid))
		return nil
	}

	if err := db.Where(uids).Find(users).Limit(1000).Error; err != nil {
		ug.Error = fmt.Errorf("通过user_id查询对应用户时失败\n%w", err)
		return nil
	}

	return
}

func (ug *UserGroup) AddRow(uid uint, gid uint) error {
	session := db.Session(&gorm.Session{NewDB: true})
	if session.Error != nil {
		ug.Error = fmt.Errorf("添加用户分组时数据库连接失败\n%w", session.Error)
		return ug.Error
	}

	if err := session.Where(UserGroup{UserID: uid, GroupID: gid}).FirstOrCreate(ug).Error; err != nil {
		ug.Error = fmt.Errorf("用户分组表数据添加失败\n%w", err)
	}
	return ug.Error
}

func (ur *UserRole) GetRoles(uid uint) (roles *[]Role) {
	rids := new([]uint)
	if err := db.Where("user_id = ?", uid).Find(ur).Limit(1000).Error; err != nil {
		ur.Error = fmt.Errorf("用户角色表中没有查询到user_id为%d的数据\n%w", uid, err)
		return nil
	}

	if err := db.Model(ur).Pluck("role_id", rids).Error; err != nil {
		ur.Error = fmt.Errorf("用户角色表中查询role_id集合失败\n%w", err)
		return nil
	} else if len(*rids) == 0 {
		ur.Error = errors.New(fmt.Sprintf("user_id为%d的用户没有绑定任何角色", uid))
		return nil
	}

	if err := db.Where(rids).Find(roles).Limit(1000).Error; err != nil {
		ur.Error = fmt.Errorf("通过role_id查对应角色时失败\n%w", err)
		return nil
	}
	return
}

func (ur *UserRole) GetUsers(rid uint) (users *[]User) {
	uids := new([]uint)
	if err := db.Where("role_id = ?", rid).Find(ur).Limit(1000).Error; err != nil {
		ur.Error = fmt.Errorf("用户角色表中没有查询到role_id为%d的数据\n%w", rid, err)
		return nil
	}

	if err := db.Model(ur).Pluck("user_id", uids).Error; err != nil {
		ur.Error = fmt.Errorf("用户角色表中查询user_id集合失败\n%w", err)
		return nil
	} else if len(*uids) == 0 {
		ur.Error = errors.New(fmt.Sprintf("role_id为%d的角色没有绑定任何用户", rid))
		return nil
	}

	if err := db.Where(uids).Find(users).Limit(1000).Error; err != nil {
		ur.Error = fmt.Errorf("通过user_id查对应用户时失败\n%w", err)
		return nil
	}

	return
}

func (ur *UserRole) AddRow(uid uint, rid uint) error {
	session := db.Session(&gorm.Session{NewDB: true})
	if session.Error != nil {
		ur.Error = fmt.Errorf("添加用户角色时数据库连接失败\n%w", session.Error)
		return ur.Error
	}

	if err := session.Where(UserRole{UserID: uid, RoleID: rid}).FirstOrCreate(ur).Error; err != nil {
		ur.Error = fmt.Errorf("用户角色表数据添加失败\n%w", err)
	}
	return ur.Error
}

func (gr *GroupRole) GetRoles(gid uint) (roles *[]Role) {
	rids := new([]uint)
	if err := db.Where("group_id = ?", gid).Find(gr).Limit(1000).Error; err != nil {
		gr.Error = fmt.Errorf("组角色表中没有查询到group_id为%d的数据\n%w", gid, err)
		return nil
	}

	if err := db.Model(gr).Pluck("role_id", rids).Error; err != nil {
		gr.Error = fmt.Errorf("组角色表中查询role_id集合失败\n%w", err)
		return nil
	} else if len(*rids) == 0 {
		gr.Error = errors.New(fmt.Sprintf("group_id为%d的组没有绑定任何角色", gid))
		return nil
	}

	if err := db.Where(rids).Find(roles).Limit(1000).Error; err != nil {
		gr.Error = fmt.Errorf("通过role_id查对应角色时失败\n%w", err)
		return nil
	}
	return
}

func (gr *GroupRole) GetGroups(rid uint) (groups *[]Group) {
	gids := new([]uint)
	if err := db.Where("role_id = ?", rid).Find(gr).Limit(1000).Error; err != nil {
		gr.Error = fmt.Errorf("组角色表中没有查询到role_id为%d的数据\n%w", rid, err)
		return nil
	}

	if err := db.Model(gr).Pluck("group_id", gids).Error; err != nil {
		gr.Error = fmt.Errorf("组角色表中查询group_id集合失败\n%w", err)
		return nil
	} else if len(*gids) == 0 {
		gr.Error = errors.New(fmt.Sprintf("role_id为%d的角色没有绑定任何组", rid))
		return nil
	}

	if err := db.Where(gids).Find(groups).Limit(1000).Error; err != nil {
		gr.Error = fmt.Errorf("通过group_id查对应组时失败\n%w", err)
		return nil
	}
	return
}

func (gr *GroupRole) AddRow(gid uint, rid uint) error {
	session := db.Session(&gorm.Session{NewDB: true})
	if session.Error != nil {
		gr.Error = fmt.Errorf("添加组角色时数据库连接失败\n%w", session.Error)
		return gr.Error
	}
	if err := session.Where(GroupRole{GroupID: gid, RoleID: rid}).FirstOrCreate(gr).Error; err != nil {
		gr.Error = fmt.Errorf("组角色表数据添加失败\n%w", err)
	}
	return gr.Error
}
