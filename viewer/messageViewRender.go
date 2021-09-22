package viewer

import (
	"fyne.io/fyne/v2"
	"image/color"
)

type messageViewRender struct {
	viewer *MessageViewWidget
}

func (r *messageViewRender) Layout(s fyne.Size) {
	r.viewer.content.Resize(s)
}

func (r *messageViewRender) MinSize() fyne.Size {
	return r.viewer.content.MinSize()
}

func (r *messageViewRender) Destroy() {
}

func (r *messageViewRender) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.viewer.content, r.viewer.cursor}
}

func (r *messageViewRender) Refresh() {
	r.viewer.cursor.Refresh()
	r.viewer.content.Refresh()
}

func (r *messageViewRender) BackgroundColor() color.Color {
	return color.Black
}
