package routes

import (
	"net/http"

	"github.com/abdrakhmanovzh/notemaker/pkg/controllers"
	"github.com/gin-gonic/gin"
)

var RegisterNoteRoutes = func(group *gin.RouterGroup) {
	group.Handle(http.MethodGet, "/", controllers.Home)
	group.Handle(http.MethodGet, "/notemaker", controllers.ShowAllNotes)
	group.Handle(http.MethodGet, "/notemaker/:id", controllers.ShowNote)
	group.Handle(http.MethodPost, "/notemaker/save_note", controllers.CreatedNote)
	group.Handle(http.MethodGet, "/notemaker/create", controllers.CreateNote)
	group.Handle(http.MethodPost, "/notemaker/:id", controllers.DeleteNote)
}
