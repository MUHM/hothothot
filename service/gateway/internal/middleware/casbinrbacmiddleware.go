package middleware

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/casbin/casbin/v2"
	xormadapter "github.com/casbin/xorm-adapter/v2"
)

type CasbinRbacMiddleware struct {
	dataSourceName string
}

func NewCasbinRbacMiddleware(dataSourceName string) *CasbinRbacMiddleware {
	return &CasbinRbacMiddleware{
		dataSourceName: dataSourceName,
	}
}

func (m *CasbinRbacMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementatio
		userId := r.Context().Value("userId")
		// 获取请求的URI
		obj := r.URL.RequestURI()
		// 获取请求方法
		act := r.Method
		// 获取用户的角色
		sub := userId
		a, _ := xormadapter.NewAdapter("mysql", m.dataSourceName, true)
		e, _ := casbin.NewEnforcer("internal/config/rbac_model.conf", a)
		e.LoadPolicy()
		// 判断策略中是否存在
		success, _ := e.Enforce(sub, obj, act)
		if success {
			next(w, r)
		} else {
			w.WriteHeader(http.StatusForbidden)
			return
		}
	}
}
