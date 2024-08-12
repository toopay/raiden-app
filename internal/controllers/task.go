package controllers

import (
	"app/internal/models"
	"github.com/sev-2/raiden"
)

type BooksController struct {
	raiden.ControllerBase
	Http  string `path:"/tasks" type:"rest"`
	Model models.Tasks
}