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

package logger

import (
	"log"
	"os"
)

var (
	logFlags = log.LstdFlags | log.Lmsgprefix
)

func Info(v ...interface{}) {
	l := log.New(os.Stdout, "[INFO]: ", logFlags)
	l.Println(v...)
}

func Warn(v ...interface{}) {
	l := log.New(os.Stdout, "[WARN]: ", logFlags)
	l.Println(v...)
}

func Error(v ...interface{}) {
	l := log.New(os.Stderr, "[ERROR]: ", logFlags)
	l.Println(v...)
}
