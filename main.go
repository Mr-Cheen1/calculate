package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

// main - главная точка входа в приложение
func main() {
	// Создаём экземпляр структуры приложения
	app := NewApp()

	// Создаём приложение с настройками
	err := wails.Run(&options.App{
		Title:             "Калькулятор времени",
		Width:             350,
		Height:            550,
		MinWidth:          320,
		MinHeight:         400,
		MaxWidth:          400,
		MaxHeight:         700,
		DisableResize:     false,
		Fullscreen:        false,
		Frameless:         false,
		StartHidden:       false,
		HideWindowOnClose: false,
		BackgroundColour:  &options.RGBA{R: 240, G: 240, B: 240, A: 1},
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Menu:             nil,
		Logger:           nil,
		LogLevel:         0,
		OnStartup:        app.startup,
		OnDomReady:       nil,
		OnBeforeClose:    nil,
		OnShutdown:       nil,
		WindowStartState: options.Normal,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
