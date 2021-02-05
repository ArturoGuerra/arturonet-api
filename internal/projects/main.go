package projects

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type (
	projects struct {
		Logger *logrus.Logger
	}
)

// New returns an instance of projects
func New(logger *logrus.Logger, g *echo.Group) error {
	p := &projects{
		Logger: logger,
	}

	g.GET("", p.getProjects())

	return nil
}
