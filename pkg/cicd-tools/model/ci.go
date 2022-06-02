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
	"gorm.io/gorm"
	"time"
)

type Project struct {
	ID          uint         `gorm:"column:id;primaryKey;autoIncrement"`
	Name        string       `gorm:"column:project;type:varchar(90);not null"`
	Intro       string       `gorm:"column:intro;type:varchar(256)"`
	Env         *[]Env       `gorm:"-"`
	Item        *[]Item      `gorm:"-"`
	ProjectEnv  *ProjectEnv  `gorm:"-"`
	ProjectItem *ProjectItem `gorm:"-"`
	Error       error        `gorm:"-"`
}

type Env struct {
	ID    uint   `gorm:"column:id;primaryKey;autoIncrement"`
	Name  string `gorm:"column:env;type:varchar(60);not null"`
	Intro string `gorm:"column:intro;type:varchar(128)"`
	Error error  `gorm:"-"`
}

type Item struct {
	ID       uint   `gorm:"column:id;primaryKey;autoIncrement"`
	Name     string `gorm:"column:item;type:varchar(90);not null"`
	Category string `gorm:"column:category;tape:varchar(60)"`
	Language string `gorm:"column:language;tape:varchar(30)"`
	Tier     string `gorm:"column:tier;type:varchar(20)"`
	Intro    string `gorm:"column:intro;type:varchar(128)"`
	Error    error  `gorm:"-"`
}

type ProjectEnv struct {
	ID        uint  `gorm:"column:id;primaryKey;autoIncrement"`
	ProjectID uint  `gorm:"column:project_id;type:integer;<-:create"`
	EnvID     uint  `gorm:"column:env_id;type:integer;<-:create"`
	Error     error `gorm:"-"`
}

type ProjectItem struct {
	ID        uint  `gorm:"column:id;primaryKey;autoIncrement"`
	ProjectID uint  `gorm:"column:project_id;type:integer;<-:create"`
	ItemID    uint  `gorm:"column:item_id;type:integer;<-:create"`
	Error     error `gorm:"-"`
}

type ProjectEnvItem struct {
	gorm.Model
	Project       string `gorm:"column:project;type:varchar(256)"`
	ProjectID     uint   `gorm:"column:project_id;type:integer;<-:create"`
	ProjectIntro  string `gorm:"column:project_intro;type:varchar(256)"`
	Env           string `gorm:"column:env;type:varchar(60)"`
	EnvID         uint   `gorm:"column:env_id;type:integer;<-:create"`
	EnvItro       string `gorm:"column:env_intro;type:varchar(256)"`
	Item          string `gorm:"column:item;type:varchar(256)"`
	ItemID        uint   `gorm:"column:item_id;type:integer;<-:create"`
	ItemIntro     string `gorm:"column:item_intro;type:varchar(256)"`
	GitRepoID     uint   `gorm:"column:git_repo_id;type:integer;<-:create"`
	GitConfigID   uint   `gorm:"column:git_config_id;type:integer;<-:create"`
	BuildConfigID uint   `gorm:"column:build_config_id;type:integer;<-:create"`
	Error         error  `gorm:"-"`
}

type GitRepo struct {
	gorm.Model
	Name       string `gorm:"column:name;type:varchar(60)"`
	RepoURL    string `gorm:"column:repo_url;type:varchar(512)"`
	RepoSSHURL string `gorm:"column:repo_ssh_url;type:varchar(512)"`
	Intro      string `gorm:"column:intro;type:varchar(256)"`
	Error      error  `gorm:"-"`
}

type GitConfig struct {
	gorm.Model
	GitRepoID  uint   `gorm:"column:git_repo_id;type:integer;<-:create"`
	Remote     string `gorm:"column:remote;type:varchar(90)"`
	GitBranch  string `gorm:"column:git_branch;type:varchar(90)"`
	UserName   string `gorm:"column:user_name;type:varchar(60)"`
	UserEmail  string `gorm:"column:user_email;type:varchar(90)"`
	Password   string `gorm:"column:password;type:varchar(128)"`
	Credential string `gorm:"column:credential;type:varchar(256)"`
	Error      error  `gorm:"-"`
}

type CommitInfo struct {
	gorm.Model
	GitRepoID       uint      `gorm:"column:git_repo_id;type:integer;<-:create"`
	GitBranch       string    `gorm:"column:git_branch;type:varchar(90)"`
	GitTag          string    `gorm:"column:git_tag;type:varchar(256)"`
	CommitHash      string    `gorm:"column:commit_hash;type:varchar(128)"`
	CommitDate      time.Time `gorm:"column:commit_date;type:datetime"`
	CommitUser      string    `gorm:"column:commit_user;type:varchar(90)"`
	CommitUserEmail string    `gorm:"column:commit_user_email;type:varchar(90)"`
	CommitMessage   string    `gorm:"column:commit_message;type:varchar(65531)"`
	ChangeLogs      string    `gorm:"column:change_logs;type:varchar(65531)"`
	Error           error     `gorm:"-"`
}

type Artifact struct {
	gorm.Model
	Name             string `gorm:"column:artifact_name;varchar(256)"`
	Release          string `gorm:"column:release;varchar(60)"`
	Version          string `gorm:"column:version;varchar(60)"`
	Md5              string `gorm:"column:md5_checksum;varchar(32);index:idx_atf_checksum"`
	SHA1             string `gorm:"column:sha1_checksum;varchar(40);index:idx_atf_checksum"`
	SHA256           string `gorm:"column:sha256_checksum;varchar(64);index:idx_atf_checksum"`
	SHA512           string `gorm:"column:sha512_checksum;varchar(128);index:idx_atf_checksum"`
	ProjectEnvItemID uint   `gorm:"column:project_env_item_id;type:integer;<-:create"`
	BuildInfoID      uint   `gorm:"column:build_info_id;type:integer;<-:create"`
	Error            error  `gorm:"-"`
}

type BuildConfig struct {
	gorm.Model
	BuildDir         string `gorm:"column:build_dir;varchar(256)"`
	BuildCmd         string `gorm:"column:build_cmd;varchar(65531)"`
	BuildEnv         string `gorm:"column:build_env;varchar(256)"`
	ProjectEnvItemID uint   `gorm:"column:project_env_item_id;type:integer;<-:create"`
	GitRepoID        uint   `gorm:"column:git_repo_id;type:integer;<-:create"`
	Error            error  `gorm:"-"`
}

type BuildInfo struct {
	gorm.Model
	BuildID          uint      `gorm:"column:build_id;type:integer;<-:create"`
	BuildName        string    `gorm:"column:build_name;varchar(90)"`
	BuildDate        time.Time `gorm:"column:build_date;type:datetime"`
	BuildUserID      uint      `gorm:"column:build_user_id;type:integer;<-:create"`
	BuildUserName    string    `gorm:"column:build_user_name;varchar(90)"`
	BuildEnv         string    `gorm:"column:build_env;varchar(256)"`
	ProjectEnvItemID uint      `gorm:"column:project_env_item_id;type:integer;<-:create"`
	GitRepoID        uint      `gorm:"column:git_repo_id;type:integer;<-:create"`
	GitBranch        string    `grom:"column:git_branch;varchar(90)"`
	CommitInfoID     uint      `gorm:"column:commit_info_id;type:integer;<-:create"`
	BuildConfigID    uint      `gorm:"column:build_config_id;type:integer;<-:create"`
	ArtifactID       uint      `gorm:"column:artifact_id;type:integer;<-:create"`
	BuildState       string    `gorm:"column:build_state;type:varchar(30)"`
}

func (Project) TableName() string {
	return "cicd_project"
}

func (Env) TableName() string {
	return "cicd_env"
}

func (Item) TableName() string {
	return "cicd_item"
}

func (ProjectEnv) TableName() string {
	return "cicd_project_env"
}

func (ProjectItem) TableName() string {
	return "cicd_project_item"
}

func (ProjectEnvItem) TableName() string {
	return "cicd_project_env_item"
}

func (GitRepo) TableName() string {
	return "cicd_git_repo"
}

func (GitConfig) TableName() string {
	return "cicd_git_config"
}

func (CommitInfo) TableName() string {
	return "cicd_commit_info"
}

func (Artifact) TableName() string {
	return "cicd_artifact"
}

func (BuildConfig) TableName() string {
	return "cicd_build_config"
}

func (BuildInfo) TableName() string {
	return "cicd_build_info"
}

func (p *Project) name() {

}
