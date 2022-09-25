package routes

import (
	"net/http"

	"github.com/abdrakhmanovzh/notemaker/pkg/controllers"
	"github.com/gin-gonic/gin"
)

var RegisterNoteRoutes = func(group *gin.RouterGroup) {
	group.Handle(http.MethodGet, "/", controllers.ShowAllNotes)
	group.Handle(http.MethodGet, "/:id", controllers.ShowNote)
	group.Handle(http.MethodPost, "/save_note", controllers.CreatedNote)
	group.Handle(http.MethodGet, "/create", controllers.CreateNote)
	group.Handle(http.MethodPost, "/:id", controllers.DeleteNote)
}
