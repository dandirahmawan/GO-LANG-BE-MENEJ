package services

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"

	"github.com/dandirahmawan/menej_api_go/commons"
	"github.com/dandirahmawan/menej_api_go/model"
	"github.com/gin-gonic/gin"
)

type FormUpload struct {
	ProjectId   string                `form:"projectId"`
	ModuleId    string                `form:"moduleId"`
	DescFile    string                `form:"descFile"`
	Basesixfour string                `form:"basesixfour"`
	Ort         string                `form:"ort"`
	FileName    string                `form:"fileName"`
	File        *multipart.FileHeader `form:"file"`
}

func Upload(ctx *gin.Context) interface{} {
	// file, _ := ctx.FormFile("file")
	// projectId := ctx.Bind("projectId")
	var dataForm FormUpload
	ctx.Bind(&dataForm)

	sessionid := ctx.Request.Header.Get("sessionid")
	sessionData := commons.GetDataSession(sessionid)
	userid := sessionData.AccountId

	var root string = os.Getenv("SYSTEMDRIVE")
	var fileSize int64 = 0
	var now time.Time = time.Now()
	var img image.Image
	var fpath string
	var ext string

	base64Data := dataForm.Basesixfour
	// fmt.Println(base64)
	if base64Data != "" {
		// bs := strings.Split(base64, ";")[1]
		coI := strings.Index(base64Data, ",")
		rawImage := string(base64Data)[coI+1:]
		unbased, _ := base64.StdEncoding.DecodeString(rawImage)
		jpgI, errJpg := jpeg.Decode(bytes.NewReader([]byte(unbased)))

		if errJpg != nil {
			fmt.Println("Error decode base64 ", errJpg)
		}

		img = jpgI

		file := img
		// fileSize := file.Size
		// filenameUpload := file.Filename
		// idx := strings.Split(filenameUpload, ".")
		// ext := "." + idx[len(idx)-1]

		sessionid := ctx.Request.Header.Get("sessionid")
		sessionData := commons.GetDataSession(sessionid)
		userid := sessionData.AccountId
		root := os.Getenv("SYSTEMDRIVE")

		uploadDir := root + "/menej/upload"
		filename := userid + "_" + now.Format("20060102150405") + ".jpeg"

		f, err := os.Create(uploadDir + "/" + filename)
		if err != nil {
			panic(err)
		}

		defer f.Close()

		if err = jpeg.Encode(f, file, nil); err != nil {
			fmt.Printf("failed to encode: %v", err)
		}

		/*get file size image*/
		fi, err := os.Stat(uploadDir + "/" + filename)
		if err != nil {
			// Could not obtain stat, handle error
			fmt.Println(err)
		}

		fileSize = fi.Size()
		fpath = uploadDir + "/" + filename
		ext = ".jpeg"

	} else {
		file := dataForm.File
		fileSize = file.Size
		filenameUpload := file.Filename
		idx := strings.Split(filenameUpload, ".")
		ext = "." + idx[len(idx)-1]
		now := time.Now()

		uploadDir := root + "/menej/upload"
		filename := userid + "_" + now.Format("20060102150405") + ext
		fpath = uploadDir + "/" + filename
		ctx.SaveUploadedFile(file, uploadDir+"/"+filename)
	}

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
	mDocFile.Path = fpath
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

func GetDocumentFileByProjectId(ctx *gin.Context) interface{} {
	projectId := ctx.Param("id")
	var m model.ViewDocumentFile
	m.ProjectId = projectId
	data := m.FindByProjectId()
	return data
}

func DownloadFile(ctx *gin.Context) {
	id := ctx.Param("id")
	// fmt.Println(id)

	var vdf model.ViewDocumentFile
	vdf.Id = id
	data := vdf.FindById()

	fileContentType := path.Ext(data.Path)
	fmt.Println(fileContentType)
	ctx.Header("Content-Type", "application/octet-stream")
	//Force browser download
	// ctx.Header("Content-Disposition", "attachment; filename="+data.FileName)
	//Browser download or preview
	ctx.Header("Content-Disposition", "inline;filename="+data.FileName)
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.Header("Cache-Control", "no-cache")

	ctx.File(data.Path)
}
