package viewer

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"image/color"
	"fyne.io/fyne/v2/driver/desktop"
)

type MessageViewWidget struct {
	widget.BaseWidget
	fyne.ShortcutHandler
	desktop.Keyable
//	fyne.Focusable
	content *widget.TextGrid
	cursorRow, cursorCol int
	cursor *canvas.Rectangle
	currentFG, currentBG color.Color
}

func (self *MessageViewWidget) Resize(s fyne.Size) {
	if s.Width == self.Size().Width && s.Height == self.Size().Height {
		return
	}
	if s.Width < 20 { // not sure why we get tiny sizes
		return
	}
	self.BaseWidget.Resize(s)
	self.content.Resize(s)
}

func New() *MessageViewWidget {
	v := &MessageViewWidget{}
	v.ExtendBaseWidget(v)
	v.content = widget.NewTextGrid()

	return v
}

func (self *MessageViewWidget) CreateRenderer() fyne.WidgetRenderer {
	self.cursor = canvas.NewRectangle(theme.PrimaryColor())
	self.cursor.Resize(fyne.NewSize(8, 16))

	r := &messageViewRender{viewer: self}
//	self.cursorMoved = r.moveCursor
	return r
}

// AcceptsTab indicates that this widget will use the Tab key (avoids loss of focus).
func (self *MessageViewWidget) AcceptsTab() bool {
	return true
}

func (self *MessageViewWidget) TypedRune(r rune) {
}

func (self *MessageViewWidget) TypedKey(key *fyne.KeyEvent) {
}

func (self *MessageViewWidget) KeyDown(k *fyne.KeyEvent) {
}

func (self *MessageViewWidget) KeyUp(k *fyne.KeyEvent) {

	//fmt.Printf("key = %+v\n", k.Name)

	//cell := self.content.Rows[self.cursorRow].Cells[self.cursorCol]

	switch k.Name {

		case fyne.KeyUp:
			if self.cursorRow > 0 {
				self.cursorRow --
			}
			self.cursor.Move(fyne.NewPos(float32(8 * self.cursorCol), float32(16 * self.cursorRow)))
			self.cursor.Refresh()

		case fyne.KeyDown:
			if true {
				self.cursorRow++
			}
			self.cursor.Move(fyne.NewPos(float32(8 * self.cursorCol), float32(16 * self.cursorRow)))
			self.cursor.Refresh()

		case fyne.KeyLeft:
			if self.cursorCol > 0 {
				self.cursorCol--
			}
			self.cursor.Move(fyne.NewPos(float32(8 * self.cursorCol), float32(16 * self.cursorRow)))
			self.cursor.Refresh()

		case fyne.KeyRight:
			if true {
				self.cursorCol++
			}
			self.cursor.Move(fyne.NewPos(float32(8 * self.cursorCol), float32(16 * self.cursorRow)))
			self.cursor.Refresh()

	}

}

func (self *MessageViewWidget) handleLine(row string) {
	for _, char := range row {
		self.handleChar(char)
	}
	self.cursorRow++
	self.cursorCol = 0
}

func (self *MessageViewWidget) handleChar(r rune) {
//	if t.cursorCol >= int(t.config.Columns) || t.cursorRow >= int(t.config.Rows) {
//		return // TODO handle wrap?
//	}
	for len(self.content.Rows)-1 < self.cursorRow {
		self.content.Rows = append(self.content.Rows, widget.TextGridRow{})
	}
	cellStyle := &widget.CustomTextGridStyle{FGColor: self.currentFG, BGColor: self.currentBG}
	for len(self.content.Rows[self.cursorRow].Cells)-1 < self.cursorCol {
		newCell := widget.TextGridCell{
		    Rune:  ' ',
		    Style: cellStyle,
		}
		self.content.Rows[self.cursorRow].Cells = append(self.content.Rows[self.cursorRow].Cells, newCell)
	}
	cell := self.content.Rows[self.cursorRow].Cells[self.cursorCol]
	if cell.Rune != r || cell.Style.TextColor() != cellStyle.FGColor || cell.Style.BackgroundColor() != cellStyle.BGColor {
		cell.Rune = r
		cell.Style = cellStyle
		self.content.SetCell(self.cursorRow, self.cursorCol, cell)
	}
	self.cursorCol++
}

func (self *MessageViewWidget) SetForegroundColor(c color.Color) {
	self.currentFG = c
}

func (self *MessageViewWidget) SetBackgroundColor(c color.Color) {
	self.currentBG = c
}

func (self *MessageViewWidget) Writeln(s string) {
	self.handleLine(s)
}

func (self *MessageViewWidget) FocusGained() {
	fmt.Println("FocusGained is called when the Check has been given focus.")
}

func (self *MessageViewWidget) FocusLost() {
	fmt.Println("FocusLost is called when the Check has had focus removed.")
}

func (self *MessageViewWidget) Focused() bool {
	return true
}
