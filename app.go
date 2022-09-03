package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/cavaliergopher/grab/v3"
	"github.com/jpastorm/gotube"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetVideoDetails(url string) (string, error) {
	yr, err := gotube.GetMetaData(url)
	if err != nil {
		return "", err
	}

	json_data, err := json.Marshal(yr)
	if err != nil {
		return "", err
	}

	return string(json_data), nil
}

func (a *App) Dialog(title string, message string, dialog_type string) (string, error) {
	selection, err := runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Title:   title,
		Message: message,
		Type:    runtime.DialogType(dialog_type),
	})

	if err != nil {
		return "", err
	}

	return selection, nil
}

func (a *App) GetDownloadSize(url string) (int64, error) {
	resp, err := http.Head(url)
	if err != nil {
		return 0, err
	}

	size, _ := strconv.Atoi(resp.Header.Get("Content-Length"))
	downloadSize := int64(size)

	return downloadSize, nil
}

var downloadProgress = 0

func (a *App) DownloadFile(name string, url string) (string, error) {

	downloadPath, err := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
		Title:                "Save File",
		DefaultFilename:      name,
		CanCreateDirectories: true,
	})

	if err != nil {
		return "", err
	}

	client := grab.NewClient()
	req, _ := grab.NewRequest(downloadPath, url)

	resp := client.Do(req)

	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()

Loop:
	for {
		select {
		case <-t.C:
			downloadProgress = int(resp.Progress() * 100)

		case <-resp.Done:
			break Loop
		}
	}

	// check for errors
	if err := resp.Err(); err != nil {
		return "", err
	}

	return downloadPath, nil
}

func (a *App) GetDownloadProgress() int {
	return downloadProgress
}
