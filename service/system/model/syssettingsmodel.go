package model

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	sysSettingsFieldNames          = builderx.RawFieldNames(&SysSettings{})
	sysSettingsRows                = strings.Join(sysSettingsFieldNames, ",")
	sysSettingsRowsExpectAutoSet   = strings.Join(stringx.Remove(sysSettingsFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	sysSettingsRowsWithPlaceHolder = strings.Join(stringx.Remove(sysSettingsFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheSysSettingsIdPrefix   = "cache:sysSettings:id:"
	cacheSysSettingsNamePrefix = "cache:sysSettings:name:"
)

type (
	SysSettingsModel interface {
		Insert(data SysSettings) (sql.Result, error)
		FindOne(id int64) (*SysSettings, error)
		FindOneByName(name sql.NullInt64) (*SysSettings, error)
		Update(data SysSettings) error
		Delete(id int64) error
	}

	defaultSysSettingsModel struct {
		sqlc.CachedConn
		table string
	}

	SysSettings struct {
		Id         int64         `db:"id"`       // 编号
		Name       sql.NullInt64 `db:"name"`     // 分类
		Classify   sql.NullInt64 `db:"classify"` // 分类
		Content    string        `db:"content"`  // 内容
		CreateTime time.Time     `db:"create_time"`
		UpdateTime time.Time     `db:"update_time"`
	}
)

func NewSysSettingsModel(conn sqlx.SqlConn, c cache.CacheConf) SysSettingsModel {
	return &defaultSysSettingsModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`sys_settings`",
	}
}

func (m *defaultSysSettingsModel) Insert(data SysSettings) (sql.Result, error) {
	sysSettingsNameKey := fmt.Sprintf("%s%v", cacheSysSettingsNamePrefix, data.Name)
	ret, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, sysSettingsRowsExpectAutoSet)
		return conn.Exec(query, data.Name, data.Classify, data.Content)
	}, sysSettingsNameKey)
	return ret, err
}

func (m *defaultSysSettingsModel) FindOne(id int64) (*SysSettings, error) {
	sysSettingsIdKey := fmt.Sprintf("%s%v", cacheSysSettingsIdPrefix, id)
	var resp SysSettings
	err := m.QueryRow(&resp, sysSettingsIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysSettingsRows, m.table)
		return conn.QueryRow(v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSysSettingsModel) FindOneByName(name sql.NullInt64) (*SysSettings, error) {
	sysSettingsNameKey := fmt.Sprintf("%s%v", cacheSysSettingsNamePrefix, name)
	var resp SysSettings
	err := m.QueryRowIndex(&resp, sysSettingsNameKey, m.formatPrimary, func(conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `name` = ? limit 1", sysSettingsRows, m.table)
		if err := conn.QueryRow(&resp, query, name); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSysSettingsModel) Update(data SysSettings) error {
	sysSettingsIdKey := fmt.Sprintf("%s%v", cacheSysSettingsIdPrefix, data.Id)
	sysSettingsNameKey := fmt.Sprintf("%s%v", cacheSysSettingsNamePrefix, data.Name)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysSettingsRowsWithPlaceHolder)
		return conn.Exec(query, data.Name, data.Classify, data.Content, data.Id)
	}, sysSettingsIdKey, sysSettingsNameKey)
	return err
}

func (m *defaultSysSettingsModel) Delete(id int64) error {
	data, err := m.FindOne(id)
	if err != nil {
		return err
	}

	sysSettingsIdKey := fmt.Sprintf("%s%v", cacheSysSettingsIdPrefix, id)
	sysSettingsNameKey := fmt.Sprintf("%s%v", cacheSysSettingsNamePrefix, data.Name)
	_, err = m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, sysSettingsIdKey, sysSettingsNameKey)
	return err
}

func (m *defaultSysSettingsModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheSysSettingsIdPrefix, primary)
}

func (m *defaultSysSettingsModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysSettingsRows, m.table)
	return conn.QueryRow(v, query, primary)
}
