package repository

type Repository interface {
	LoadMazeData() ([]byte, error)
}

type FilePath struct {
	Path    string
	NewName string
}
