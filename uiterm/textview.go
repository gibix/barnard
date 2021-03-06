package uiterm

import (
	"strings"

	"github.com/nsf/termbox-go"
)

type Textview struct {
	Lines          []string
	CurrentLine    int
	Fg, Bg         Attribute
	showTimestamps bool
	parsedLines    []string

	ui             *Ui
	x0, y0, x1, y1 int
}

func (t *Textview) uiInitialize(ui *Ui) {
	t.ui = ui
	t.showTimestamps = true
}

func (t *Textview) ToggleTimestamps() {
	if t.showTimestamps == true {
		t.showTimestamps = false
	} else {
		t.showTimestamps = true
	}
	t.updateParsedLines()
	t.uiDraw()
}

func (t *Textview) uiSetActive(active bool) {
}

func (t *Textview) uiSetBounds(x0, y0, x1, y1 int) {
	t.x0 = x0
	t.y0 = y0
	t.x1 = x1
	t.y1 = y1
	t.updateParsedLines()
	t.uiDraw()
}

func (t *Textview) ScrollUp() {
	if newLine := t.CurrentLine + 1; newLine < len(t.parsedLines) {
		t.CurrentLine = newLine
	}
	t.uiDraw()
}

func (t *Textview) ScrollDown() {
	if newLine := t.CurrentLine - 1; newLine >= 0 {
		t.CurrentLine = newLine
	}
	t.uiDraw()
}

func (t *Textview) ScrollTop() {
	if newLine := len(t.parsedLines) - 1; newLine > 0 {
		t.CurrentLine = newLine
	} else {
		t.CurrentLine = 0
	}
	t.uiDraw()
}

func (t *Textview) ScrollBottom() {
	t.CurrentLine = 0
	t.uiDraw()
}

func (t *Textview) updateParsedLines() {
	width := t.x1 - t.x0

	if t.Lines == nil || width <= 0 {
		t.parsedLines = nil
		return
	}

	parsed := make([]string, 0, len(t.Lines))
	for _, line := range t.Lines {
		var l = line
		if t.showTimestamps == false {
			l = strings.TrimSpace(strings.Split(line, "]")[1])
		}
		current := ""
		chars := 0
		reader := strings.NewReader(l)
		for {
			if chars >= width {
				parsed = append(parsed, current)
				chars = 0
				current = ""
			}
			if reader.Len() <= 0 {
				if chars > 0 {
					parsed = append(parsed, current)
				}
				break
			}
			if ch, _, err := reader.ReadRune(); err == nil {
				current = current + string(ch)
				chars++
			}
		}
	}
	t.parsedLines = parsed
}

func (t *Textview) AddLine(line string) {
	t.Lines = append(t.Lines, line)
	t.updateParsedLines()
	t.uiDraw()
}

func (t *Textview) Clear() {
	t.Lines = nil
	t.CurrentLine = 0
	t.parsedLines = nil
	t.uiDraw()
}

func (t *Textview) uiDraw() {
	t.ui.beginDraw()
	defer t.ui.endDraw()

	var reader *strings.Reader
	writeableLines := t.y1 - t.y0
	lineNum := 0
	if writeableLines < len(t.parsedLines) {
		lineNum = len(t.parsedLines) - writeableLines
	}
	//Beep()
	for y := t.y0; y < writeableLines; y++ {
		if lineNum < len(t.parsedLines) {
			reader = strings.NewReader(t.parsedLines[lineNum])
		} else {
			reader = nil
		}
		for x := t.x0; x < t.x1; x++ {
			var chr rune = ' '
			if reader != nil {
				if ch, _, err := reader.ReadRune(); err == nil {
					chr = ch
				} //no err
			} //reader != nil
			termbox.SetCell(x, y, chr, termbox.Attribute(t.Fg), termbox.Attribute(t.Bg))
		} //each x
		lineNum++
	} //each y
} //func

func (t *Textview) uiKeyEvent(key Key) {
}

func (t *Textview) uiCharacterEvent(chr rune) {
}
