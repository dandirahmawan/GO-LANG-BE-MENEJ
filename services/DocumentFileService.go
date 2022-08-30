package services

import (
	"fmt"
	"mime/multipart"
	"os"
	"strings"
	"time"

	"github.com/dandirahmawan/menej_api_go/commons"
	"github.com/dandirahmawan/menej_api_go/model"
	"github.com/gin-gonic/gin"
)

type FormUpload struct {
	ProjectId string                `form:"projectId"`
	ModuleId  string                `form:"moduleId"`
	File      *multipart.FileHeader `form:"file"`
	DescFile  string                `form:"descFile"`
	Base64    string                `form:"base64"`
	Ort       string                `form:"ort"`
	FileName  string                `form:"fileName"`
}

func Upload(ctx *gin.Context) interface{} {
	// file, _ := ctx.FormFile("file")
	// projectId := ctx.Bind("projectId")
	var dataForm FormUpload
	ctx.Bind(&dataForm)
	fmt.Println(dataForm)

	file := dataForm.File

	fileSize := file.Size
	filenameUpload := file.Filename
	idx := strings.Split(filenameUpload, ".")
	ext := "." + idx[len(idx)-1]
	now := time.Now()

	sessionid := ctx.Request.Header.Get("sessionid")
	sessionData := commons.GetDataSession(sessionid)
	userid := sessionData.AccountId
	root := os.Getenv("SYSTEMDRIVE")

	uploadDir := root + "/menej/upload"
	filename := userid + "_" + now.Format("20060102150405") + ext
	ctx.SaveUploadedFile(file, uploadDir+"/"+filename)

	/*save data to database*/
	id := commons.GeneratdUUID(20)
	var mDocFile model.DocumentFile
	mDocFile.Id = id
	mDocFile.UserId = userid
	mDocFile.DescriptionFile = dataForm.DescFile
	mDocFile.FileName = dataForm.FileName
	mDocFile.Extension = ext
	mDocFile.FileSize = fileSize
	mDocFile.ModulId = dataForm.ModuleId
	mDocFile.ProjectId = dataForm.ProjectId
	mDocFile.Path = uploadDir + "/" + filename
	mDocFile.UploadDate = now
	fmt.Println(mDocFile)
	mDocFile.Save()

	/*get data upload*/
	var data model.ViewDocumentFile
	data.Id = id
	data = data.FindById()
	return data
}

func DeleteDocumentFile(ctx *gin.Context) interface{} {
	id := ctx.Param("id")
	var mDocFile model.DocumentFile
	mDocFile.Id = id

	data := mDocFile.FindById()
	path := data.Path
	// fmt.Println(path)
	// fmt.Println("=======================")
	err := os.Remove(path)
	if err != nil {
		fmt.Println(nil)
		return 0
	}

	mDocFile.DeleteById()
	return 1
}
