package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {

	a := app.New()
	a.Settings().SetTheme(theme.LightTheme())
	dark := false

	textoColor := color.RGBA{R: 35, G: 38, B: 52, A: 255}

	w := a.NewWindow("ASIR GARMIN")
	w.Resize(fyne.NewSize(800, 600))

	btn1 := widget.NewButton("Setup", func() { go Startup(a) })
	btn2 := widget.NewButton("OS Info", func() {
		osWin := a.NewWindow("OS Info")
		osWin.Resize(fyne.NewSize(600, 400))
		osWin.SetFixedSize(true)
		go OsInfo(osWin)
	})

	banner := canvas.NewText("ASIR", color.Black)
	banner.TextStyle = fyne.TextStyle{Bold: true}
	banner.TextSize = 30
	banner.Alignment = fyne.TextAlignCenter

	bannerMargen := widget.NewLabel("")

	garminImg := canvas.NewImageFromFile("D:\\Garmin\\assets\\garmin.jpg")
	garminImg.SetMinSize(fyne.NewSize(800, 600))
	garminImg.FillMode = canvas.ImageFillContain

	sideGap := canvas.NewRectangle(color.Transparent)
	sideGap.SetMinSize(fyne.NewSize(15, 0))
	lineaDeMargen := canvas.NewRectangle(color.Black)
	lineaDeMargen.SetMinSize(fyne.NewSize(5, 0))

	topGap := canvas.NewRectangle(color.Transparent)
	topGap.SetMinSize(fyne.NewSize(0, 15))

	garminText := canvas.NewText("Okay Garmin...", textoColor)
	garminText.TextSize = 35
	garminText.TextStyle = fyne.TextStyle{Bold: true, Italic: true}

	btnTheme := widget.NewButton("", nil)
	darkmodeIcon, darkmodeErr := fyne.LoadResourceFromPath("D:\\Garmin\\assets\\darkmode.png")
	if darkmodeErr != nil {
		btnTheme.SetText("Oscuro/Claro")
	}
	btnTheme.SetIcon(darkmodeIcon)

	btnTheme.OnTapped = func() {
		if !dark {
			a.Settings().SetTheme(theme.DarkTheme())
			banner.Color = color.White
			lineaDeMargen.FillColor = color.White
			garminText.Color = color.White
			dark = true
		} else {
			a.Settings().SetTheme(theme.LightTheme())
			banner.Color = color.Black
			garminText.Color = color.Black
			lineaDeMargen.FillColor = color.Black
			dark = false
		}
	}

	box1 := container.NewVBox(banner, bannerMargen, btn1, btn2, btnTheme)
	box2 := container.NewVBox(garminText, garminImg)
	content := container.NewHBox(sideGap, box1, sideGap, lineaDeMargen, sideGap, box2, sideGap)
	root := container.NewVBox(topGap, content)

	w.SetContent(root)
	w.ShowAndRun()

}
