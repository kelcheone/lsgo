package main

import "fmt"

func All(dir string) {
	files, err := ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
