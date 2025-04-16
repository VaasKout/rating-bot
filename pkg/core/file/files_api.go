package file

type FilesApi[T any] interface {
	ReadJsonFile(path string, model *T) error
	WriteJsonFile(path string, model *T) error
	ReadFolders(path string) ([]string, error)
}
