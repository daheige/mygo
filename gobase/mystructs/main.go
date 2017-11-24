package mystructs

type HgStruct struct {
	name string
	job  string
}

type File struct {
	Fd   int    "fp int" //结构体标签
	Name string "filename"
}

//通过工厂模式创建结构体实例 返回实例的指针
func GetFileInstance(fd int, name string) *File {
	if fd < 0 {
		return nil
	}

	return &File{Fd: fd, Name: name}
}
