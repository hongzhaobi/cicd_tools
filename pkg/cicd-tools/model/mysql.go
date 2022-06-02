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
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const (
	MySQLDSN    string = "root:123456@tcp(192.168.116.131:3306)/data100_rnd?charset=utf8mb4&parseTime=True&loc=Local"
	TablePrefix string = "rnd_"
)

var (
	db, _ = gorm.Open(mysql.Open(MySQLDSN), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   TablePrefix,
			SingularTable: true,
		},
	})
)

// func (m *MySQL) DB() *gorm.DB {
//     db, err := gorm.Open(mysql.Open(MySQLDSN), &gorm.Config{
//         NamingStrategy: schema.NamingStrategy{
//             TablePrefix:   TablePrefix,
//             SingularTable: true,
//         },
//     })
//
//     if err != nil {
//         fmt.Println(err)
//     }
//     return db
// }
