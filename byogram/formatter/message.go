package formatter

func (fm *Formatter) SetIkbdDim(dim []int) {

	fm.Keyboard.Keyboard = make([][]btn, len(dim))
	for i := 0; i < len(dim); i++ {
		fm.Keyboard.Keyboard[i] = make([]btn, dim[i])
	}
}

func (fm *Formatter) doRutine() {
	if fm.Keyboard.x == len(fm.Keyboard.Keyboard[fm.Keyboard.y]) {
		fm.Keyboard.x = 0
		fm.Keyboard.y = fm.Keyboard.y + 1
	}
}

func (fm *Formatter) WriteInlineButtonCmd(label, cmd string) {
	fm.doRutine()
	fm.Keyboard.Keyboard[fm.Keyboard.y][fm.Keyboard.x].Label = label
	fm.Keyboard.Keyboard[fm.Keyboard.y][fm.Keyboard.x].what = bCmd
	fm.Keyboard.Keyboard[fm.Keyboard.y][fm.Keyboard.x].Cmd = cmd

	fm.Keyboard.x = fm.Keyboard.x + 1

}

func (fm *Formatter) WriteInlineButtonUrl(label, url string) {
	fm.doRutine()
	fm.Keyboard.Keyboard[fm.Keyboard.y][fm.Keyboard.x].Label = label
	fm.Keyboard.Keyboard[fm.Keyboard.y][fm.Keyboard.x].what = bUrl
	fm.Keyboard.Keyboard[fm.Keyboard.y][fm.Keyboard.x].Url = url

	fm.Keyboard.x = fm.Keyboard.x + 1

}
