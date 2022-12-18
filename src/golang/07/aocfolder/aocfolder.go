package aocfolders

type (
	File struct {
		Name string
		Size int64
	}
	Directory struct {
		Subdirectories []*Directory
		Files          []File
		Name           string
		Path           string
		Parent         *Directory
		FileSizes      int64
	}
)

func NewSystem() *Directory {
	return &Directory{nil, nil, "/", "/", nil, 0}
}

func (this *Directory) MkDir(name string) *Directory {
	newDirectory := Directory{nil, nil, name, this.Path + name + "/", this, 0}
	this.Subdirectories = append(this.Subdirectories, &newDirectory)
	return this
}

func (this *Directory) ChangeDir(name string) *Directory {
	switch name {
	case "..":
		return this.Parent
	case "/":
		if this.Parent == nil {
			return this
		} else {
			return this.Parent.ChangeDir("/")
		}
	default:
		for _, subDir := range this.Subdirectories {
			if subDir.Name == name {
				return subDir
			}
		}
	}
	return this
}

func (this *Directory) TouchFile(name string, size int64) {
	newFile := File{name, size}
	this.Files = append(this.Files, newFile)

	this.FileSizes += size

	//also increase file size of parent directories
	parent := this.Parent
	for ok := (parent != nil); ok == true; ok = (parent != nil) {
		parent.FileSizes += size
		parent = parent.Parent
	}

}
