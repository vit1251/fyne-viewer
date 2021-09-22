package main

import (
//	"fmt"
//	"fyne.io/fyne/v2"
//	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"image/color"
	"github.com/vit1251/fyne-viewer/viewer"
)

func main() {

	a := app.New()
	w := a.NewWindow("Fyne Text Viewer")

	colorBlue := color.NRGBA{R: 0, G: 0, B: 180, A: 255}
	colorYellow := color.NRGBA{R: 180, G: 180, B: 0, A: 255}
	colorRed := color.NRGBA{R: 180, G: 0, B: 0, A: 255}
	colorGreen := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	colorGray := color.NRGBA{R: 180, G: 180, B: 180, A: 255}

	msgViewer := viewer.New()
	msgViewer.SetBackgroundColor(color.Black)


	msgViewer.SetForegroundColor(colorBlue)
	msgViewer.Writeln(" FROM: 'Vitold S.' <vit1251@gmail.com>")
	msgViewer.Writeln("   TO: 'Alexanr B.' <alexb@gmail.com>")

	msgViewer.SetForegroundColor(colorYellow)
	msgViewer.Writeln("-------------------------------------------------------")

	msgViewer.SetForegroundColor(colorGreen)
	msgViewer.Writeln(" VS>> Добрый день!")
	msgViewer.SetForegroundColor(colorRed)
	msgViewer.Writeln(" SV> Ага! Конечно")
	msgViewer.SetForegroundColor(colorGray)
	msgViewer.Writeln(" А чего?")

	w.SetContent(container.NewVBox(
	    msgViewer,
	))

	w.ShowAndRun()
}
