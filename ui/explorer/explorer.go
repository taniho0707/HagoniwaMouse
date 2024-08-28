package explorer

import (
	"fmt"
	"io"
	"os"

	"gioui.org/app"
	"gioui.org/x/explorer"
)

type Explorer struct {
	expl *explorer.Explorer
	w    *app.Window
}

type Result struct {
	Data     []byte
	Error    error
	Filepath string
}

func NewExplorer(w *app.Window) *Explorer {
	return &Explorer{
		expl: explorer.NewExplorer(w),
		w:    w,
	}
}

func (e *Explorer) ChoseFile(onResult func(r Result), extensions ...string) {
	go func(onResult func(r Result)) {
		defer func(e *Explorer) {
			e.w.Invalidate()
		}(e)

		file, err := e.expl.ChooseFile(extensions...)
		if err != nil {
			err = fmt.Errorf("failed to open file: %w", err)
			onResult(Result{Error: err})
			return
		}

		defer func(file io.ReadCloser) {
			err := file.Close()
			if err != nil {
				err = fmt.Errorf("failed to close file: %w", err)
				onResult(Result{Error: err})
			}
		}(file)

		filePath := ""
		if f, ok := file.(*os.File); ok {
			filePath = f.Name()
		}

		data, err := io.ReadAll(file)
		if err != nil {
			err = fmt.Errorf("failed to read file: %w", err)
			onResult(Result{Error: err})
			return
		}

		onResult(Result{Data: data, Filepath: filePath, Error: nil})
	}(onResult)
}
