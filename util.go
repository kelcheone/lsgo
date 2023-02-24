package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

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

	sort.Slice(list, func(i, j int) bool { return list[i].Name() < list[j].Name() })

	return list, nil
}

func handleErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func createLong(path string) (string, error) {
	f, err := os.Lstat(path)
	if err != nil {
		return "", err
	}
	fileName := filepath.Base(path)

	nf := fmt.Sprintf("%v \t %v \t%v\t %v", f.Mode().Perm(), f.Size(), f.ModTime().Format("01 Feb 15:04"), fileName)

	return nf, nil
}

func hasFlag(flag string, prefix string) bool {
	return strings.HasPrefix(flag, prefix)
}

func listFlag(args []string) bool {
	return strings.HasPrefix(args[0], "-")
}
