package emailmanager

import "github.com/labstack/echo/v4"

func (em *emailManager) Register(g *echo.Group) {
	g.POST("/", em.send)
}
