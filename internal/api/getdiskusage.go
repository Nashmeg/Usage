package api

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func (r *queryResolver) Files(ctx context.Context, where *FileWhereInput) ([]*File, error) {
	var files []*File
	files, err := getFiles(where.MountPoint)
	return files, err
}

func getFiles(mountPoint string) ([]*File, error) {
	var results []*File
	//mountPoint = "/Users/Negar/Desktop"
	files, err := ioutil.ReadDir(mountPoint)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	for _, f := range files {
		var result File
		result.MountPoint = &mountPoint
		name := mountPoint + "/" + f.Name()
		result.Name = &name
		size := f.Size() / 1024
		result.Usage = &size
		results = append(results, &result)
	}
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)

	return results, nil

}
