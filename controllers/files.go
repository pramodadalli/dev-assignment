package controllers

import (
	"dev-assignment/models"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/beego/beego/v2/server/web"
)

type FileController struct {
	web.Controller
}

func (c *FileController) Upload() {
	username := c.Ctx.Input.GetData("username").(string)
	user := models.GetUser(username)

	file, header, err := c.GetFile("file")
	if err != nil {
		c.CustomAbort(400, "file error")
		return
	}
	defer file.Close()

	if user.UsedStorage+header.Size > user.TotalStorage {
		c.CustomAbort(400, "storage quota exceeded")
		return
	}

	dir := filepath.Join("storage", username)
	os.MkdirAll(dir, os.ModePerm)

	dstPath := filepath.Join(dir, header.Filename)
	out, err := os.Create(dstPath)
	if err != nil {
		c.CustomAbort(500, "file save error")
		return
	}
	defer out.Close()
	io.Copy(out, file)

	user.UsedStorage += header.Size
	models.FileMetadata[username] = append(models.FileMetadata[username], models.FileMeta{
		Filename:     header.Filename,
		OriginalName: header.Filename,
		Size:         header.Size,
		UploadTime:   time.Now(),
	})

	c.Data["json"] = map[string]string{"message": "upload successful"}
	c.ServeJSON()
}

func (c *FileController) RemainingStorage() {
	username := c.Ctx.Input.GetData("username").(string)
	user := models.GetUser(username)
	c.Data["json"] = map[string]interface{}{
		"total":     user.TotalStorage,
		"used":      user.UsedStorage,
		"remaining": user.TotalStorage - user.UsedStorage,
	}
	c.ServeJSON()
}

func (c *FileController) ListFiles() {
	username := c.Ctx.Input.GetData("username").(string)
	files := models.FileMetadata[username]
	c.Data["json"] = files
	c.ServeJSON()
}
