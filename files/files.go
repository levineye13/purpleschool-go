package files

import "fmt"

type File struct {
}

func (file *File) ReadFile() error {
	fmt.Println("read")

	return nil
}

func (file *File) WriteFile() error {
	fmt.Println("write")

	return nil
}