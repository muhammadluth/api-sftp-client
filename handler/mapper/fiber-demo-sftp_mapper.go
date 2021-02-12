package mapper

import (
	"fiber-demo-sftp/handler"
	"fiber-demo-sftp/model"
	"math"
	"os"
	"strconv"
)

type FiberDemoSftpMapper struct {
}

func NewFiberDemoSftpMapper() handler.IFiberDemoSftpMapper {
	return &FiberDemoSftpMapper{}
}

func (m *FiberDemoSftpMapper) ToResponseGetFileFromSftp(pathDirectory string, files []os.FileInfo) (response []model.ResponseGetFileSFTP) {
	for i := range files {
		if !files[i].IsDir() {
			payload := model.ResponseGetFileSFTP{
				Directory:    pathDirectory,
				Filename:     files[i].Name(),
				Size:         m.doFormatSize(float64(files[i].Size())),
				FileModified: files[i].ModTime().Format("2006-01-02 15:04:05"),
			}
			response = append(response, payload)
		}
	}
	return response
}

func (m *FiberDemoSftpMapper) ToResponseGetDirectoryFromSftp(directory []os.FileInfo) (response []model.ResponseGetDirectorySFTP) {
	for i := range directory {
		if directory[i].IsDir() {
			payload := model.ResponseGetDirectorySFTP{
				Directory:    directory[i].Name(),
				Size:         m.doFormatSize(float64(directory[i].Size())),
				FileModified: directory[i].ModTime().Format("2006-01-02 15:04:05"),
			}
			response = append(response, payload)
		}
	}
	return response
}

func (m *FiberDemoSftpMapper) doFormatSize(size float64) string {
	var suffixes [5]string
	suffixes[0] = "B"
	suffixes[1] = "KB"
	suffixes[2] = "MB"
	suffixes[3] = "GB"
	suffixes[4] = "TB"

	if size == 0 {
		return strconv.FormatFloat(size, 'f', -1, 64) + " " + suffixes[0]
	}

	base := math.Log(size) / math.Log(1024)
	getSize := m.doSize(math.Pow(1024, base-math.Floor(base)), .5, 2)
	getSuffix := suffixes[int(math.Floor(base))]
	return strconv.FormatFloat(getSize, 'f', -1, 64) + " " + string(getSuffix)
}

func (m *FiberDemoSftpMapper) doSize(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}
