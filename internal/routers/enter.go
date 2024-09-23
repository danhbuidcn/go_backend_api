package routers

import (
	"github.com/danhbuidcn/go_backend_api/internal/routers/manager"
	"github.com/danhbuidcn/go_backend_api/internal/routers/user"
)

type RouterGroup struct {
	User    user.UserRouterGroup
	Manager manager.ManagerRouterGroup
}

var RouterGroupApp = new(RouterGroup)
