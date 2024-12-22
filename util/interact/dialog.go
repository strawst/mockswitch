package interact

import "github.com/wailsapp/wails/v3/pkg/application"

func Error(title string, err error) {
	dialog := application.ErrorDialog()
	dialog.SetTitle(title)
	dialog.SetMessage(err.Error())
	dialog.Show()
}
