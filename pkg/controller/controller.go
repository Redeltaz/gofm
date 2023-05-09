package controller

import (
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"sort"
)

type Controller struct {
	CWD            string
	CurrentContent []File
}

type File struct {
	Name         string
	Size         int64
	Permissions  fs.FileMode
	lastModified string
}

func InitController() *Controller {
	controller := &Controller{}

	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	controller.CWD = path

	return controller
}

func (c *Controller) SetCurrentContent() {
	files, err := ioutil.ReadDir(c.CWD)
	if err != nil {
		log.Fatal(err)
	}

	dirContent := []File{}

	for _, file := range files {
		dirContent = append(dirContent, File{
			file.Name(),
			file.Size(),
			file.Mode(),
			file.ModTime().Format("1 jan 06 03:04PM"),
		})
	}

	sort.SliceStable(dirContent, func(i, j int) bool {
		return dirContent[i].Permissions.IsDir() && !dirContent[j].Permissions.IsDir()
	})

	c.CurrentContent = dirContent
}
