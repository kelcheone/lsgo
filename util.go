package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	color "github.com/fatih/color"
)

type Long struct {
	perm    string
	size    string
	sizeInt int
	time    string
	name    string
	file    fs.FileInfo
}

func ReadDir(dir string) ([]fs.FileInfo, error) {
	f, err := os.Open(dir)
	if err != nil {
		return nil, err
	}
	list, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return nil, err
	}
	return list, nil
}

func handleErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func createLong(path string) (Long, error) {
	f, err := os.Lstat(path)
	long := &Long{}
	if err != nil {
		return *long, err
	}
	fileName := filepath.Base(path)
	size := fmt.Sprintf("%v", f.Size())
	if int(f.Size()) > 1000 {
		size = fmt.Sprintf("%.1fK", float32(f.Size())/1000)
	}
	long = &Long{
		perm:    f.Mode().Perm().String(),
		size:    size,
		sizeInt: int(f.Size()),
		time:    f.ModTime().Format("01 Feb 15:04"),
		name:    fileName,
		file:    f,
	}

	return *long, nil
}

func hasFlag(flag string, prefix string) bool {
	return strings.HasPrefix(flag, prefix)
}

func listFlag(args []string) bool {
	return strings.HasPrefix(args[0], "-")
}

func AddToLong(dir string, files []fs.FileInfo, all []Long, ignore bool) []Long {
	for _, file := range files {
		if dir == "." {
			long, _ := createLong(file.Name())
			if !ignore {
				all = append(all, long)
			} else {
				if !hasFlag(long.name, ".") {
					all = append(all, long)
				}
			}
		} else {
			tPath := fmt.Sprintf("%v/%v", dir, file.Name())
			long, _ := createLong(tPath)
			if !ignore {
				all = append(all, long)
			} else {
				if !hasFlag(long.name, ".") {
					all = append(all, long)
				}
			}
		}
	}
	return all
}

func isDir(file string) {
	c := color.New(color.FgBlue)
	bold := c.Add(color.Bold)
	bold.Printf("%v\t", file)
}

func isExec(file string) {
	c := color.New(color.FgGreen)
	bold := c.Add(color.Bold)
	bold.Printf("%v\t", file)
}

func isExecutable(file fs.FileInfo) bool {
	return file.Mode()&0111 != 0
}

func lsPrint(file fs.FileInfo) {
	if file.IsDir() {
		isDir(file.Name())
	} else if isExecutable(file) {
		isExec(file.Name())
	} else {
		fmt.Printf("%v\t", file.Name())
	}
}

func lsPrintLong(file fs.FileInfo) {
	c := color.New(color.FgBlue)
	bold := c.Add(color.Bold)

	e := color.New(color.FgGreen)
	boldE := e.Add(color.Bold)
	if file.IsDir() {
		bold.Printf("%v \t %v \t%v\t %v \n", file.Mode().Perm().String(), file.Size(), file.ModTime().Format("01 Feb 15:04"), file.Name())
	} else if isExecutable(file) {
		boldE.Printf("%v \t %v \t%v\t %v \n", file.Mode().Perm().String(), file.Size(), file.ModTime().Format("01 Feb 15:04"), file.Name())
	} else {
		fmt.Printf("%v \t %v \t%v\t %v \n", file.Mode().Perm().String(), file.Size(), file.ModTime().Format("01 Feb 15:04"), file.Name())
	}

}
