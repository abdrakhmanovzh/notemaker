package controllers

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/abdrakhmanovzh/notemaker/pkg/models"
	"github.com/gin-gonic/gin"
)

func ShowAllNotes(ctx *gin.Context) {
	var notes = models.ShowAllNotes()
	var toCheck bool

	if len(notes) == 0 {
		toCheck = false
	} else {
		toCheck = true
	}

	ctx.HTML(200, "notes.page.tmpl", gin.H{
		"ToCheck": toCheck,
		"Note":    notes,
	})
}

func ShowNote(ctx *gin.Context) {
	NoteId := ctx.Param("id")
	NoteDetails := models.ShowNote(NoteId)

	ctx.HTML(200, "show.page.tmpl", gin.H{
		"Name":    NoteDetails.Name,
		"Content": NoteDetails.Content,
		"ID":      NoteId,
	})
}

func CreateNote(ctx *gin.Context) {
	ctx.HTML(200, "create.page.tmpl", nil)
}

func CreatedNote(ctx *gin.Context) {
	CreateNote := new(models.Note)
	CreateNote.Name = ctx.Request.FormValue("title")
	CreateNote.Content = ctx.Request.FormValue("content")

	rand.Seed(time.Now().UnixNano())
	CreateNote.Id = fmt.Sprintf("%v", rand.Intn(1000))
	models.CreateNote(CreateNote)

	ctx.Redirect(301, "/notemaker")
}

func DeleteNote(ctx *gin.Context) {
	NoteId := ctx.Param("id")
	models.DeleteNote(NoteId)
	ctx.Redirect(301, "/notemaker")
}

func Home(ctx *gin.Context) {
	ctx.HTML(200, "home.page.tmpl", nil)
}
