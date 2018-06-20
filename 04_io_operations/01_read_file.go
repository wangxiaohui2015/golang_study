package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

var fileName = "/etc/hosts"

func main() {
	// readFileToString()
	// showFileStat()
	// copyFile(fileName, "/tmp/hosts")
	listFiles()
}

func showFileStat() {
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		return
	}
	fmt.Println(fileInfo.IsDir())
	fmt.Println(fileInfo.Mode())
	fmt.Println(fileInfo.ModTime())
	fmt.Println(fileInfo.Name())
	fmt.Println(fileInfo.Size())
	fmt.Println(fileInfo.Sys())
}

func readFileToString() {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	str := string(bytes)
	fmt.Println(str)
}

func copyFile(src string, dest string) error {
	sfi, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !sfi.Mode().IsRegular() {
		// cannot copy non-regular files (e.g., directories,
		// symlinks, devices, etc.)
		return fmt.Errorf("CopyFile: non-regular source file %s (%q)", sfi.Name(), sfi.Mode().String())
	}
	dfi, err := os.Stat(dest)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	} else {
		if !(dfi.Mode().IsRegular()) {
			return fmt.Errorf("CopyFile: non-regular destination file %s (%q)", dfi.Name(), dfi.Mode().String())
		}
		if os.SameFile(sfi, dfi) {
			return err
		}
	}
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer func() {
		cerror := out.Close()
		if err == nil {
			err = cerror
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return err
	}
	err = out.Sync()
	if err != nil {
		return err
	}
	return nil
}

func listFiles() {
	files, err := ioutil.ReadDir("/home/sunny/")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, f := range files {
		fmt.Println(f.Name(), f.IsDir())
	}
}
