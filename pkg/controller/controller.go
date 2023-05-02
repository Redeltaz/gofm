package controller

import (
	"io/fs"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

type Controller struct {
    CWD string
    CurrentContent []File
}

type File struct {
    Name string
    Size int64
    Permissions fs.FileMode
    lastModified string
    isDir bool
}

type ControllerActions interface {
    GetDirContent()
}

func InitController() *Controller {
    controller := &Controller{}

    cmd := exec.Command("pwd")
    out, err := cmd.Output()
    if err != nil {
        log.Fatal(err)
    }

    controller.CWD = strings.TrimSuffix(string(out), "\n")

    return controller
}

func (c *Controller) GetDirContent() []File {
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
            file.IsDir(),
        })
    }

    c.CurrentContent = dirContent
}
