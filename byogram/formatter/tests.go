package formatter

import (
	"firstgobot/byogram/errors"
	"fmt"
	"log"
)

func (tfm *Formatter) AssertPhoto(path string, condition bool) (err error) {
	var function string
	if tfm.Message.Photo != path {
		if tfm.kindofmedia == fromStorage {
			function = "AddPhotoFromStorage"
		} else if tfm.kindofmedia == fromInternet {
			function = "AddPhotoFromInternet"
		} else if tfm.kindofmedia == fromTelegram {
			function = "AddPhotoFromTG"
		}
		err = errors.AssertTest(tfm.Message.Photo, function, path, "AssertPhoto")
	}
	if condition {
		if err != nil {
			log.Fatal(err)
		}
	}
	return err
}

func (tfm *Formatter) AssertVideo(path string, condition bool) (err error) {
	var function string
	if tfm.Message.Video != path {
		if tfm.kindofmedia == fromStorage {
			function = "AddVideoFromStorage"
		} else if tfm.kindofmedia == fromInternet {
			function = "AddVideoFromInternet"
		} else if tfm.kindofmedia == fromTelegram {
			function = "AddVideoFromTG"
		}
		err = errors.AssertTest(tfm.Message.Video, function, path, "AssertVideo")
	}
	if condition {
		if err != nil {
			log.Fatal(err)
		}
	}
	return err
}

func (tfm *Formatter) AssertInlineKeyboard(testdim []int, kbNames, kbDatas, typeofbuttons []string, condition bool) (err error) {
	var (
		dim []int
	)

	for i := 0; i < len(tfm.Keyboard.Keyboard); i++ {
		dim = append(dim, len(tfm.Keyboard.Keyboard[i]))
	}

	if len(testdim) == len(dim) {
		for i := 0; i < len(dim); i++ {
			if testdim[i] != dim[i] {
				err = errors.AssertTest(fmt.Sprint(dim), "SetIkbdDim", fmt.Sprint(testdim), "AssertInlineKeyboard")
				if condition {
					log.Fatal(err)
				}
			}
		}
		if len(kbNames) == len(kbDatas) && len(kbNames) == len(typeofbuttons) && err == nil {
			for i := 0; i < len(testdim); i++ {
				for j := 0; j < testdim[i]; j++ {
					if tfm.Keyboard.Keyboard[i][j].Label != kbNames[i+j] {
						err = errors.AssertTest(fmt.Sprint("name of buttons is ", tfm.Keyboard.Keyboard[i][j].Label), "WriteInlineButtonUrl/WriteInlineButtonCmd", fmt.Sprint("name of buttons is ", kbNames[i]), "AssertInlineKeyboard")
						if condition {
							log.Fatal(err)
						}
					} else if typeofbuttons[i] == "url" && tfm.Keyboard.Keyboard[i][j].Url != kbDatas[i+j] {
						err = errors.AssertTest(fmt.Sprint("url of button is ", tfm.Keyboard.Keyboard[i][j].Url), "WriteInlineButtonUrl", fmt.Sprint("url of button is ", kbDatas[i]), "AssertInlineKeyboard")
						if condition {
							log.Fatal(err)
						}
					} else if typeofbuttons[i] == "cmd" && tfm.Keyboard.Keyboard[i][j].Cmd != kbDatas[i+j] {
						err = errors.AssertTest(fmt.Sprint("cmd of button is ", tfm.Keyboard.Keyboard[i][j].Cmd), "WriteInlineButtonCmd", fmt.Sprint("cmd of button is ", kbDatas[i]), "AssertInlineKeyboard")
						if condition {
							log.Fatal(err)
						}
					}
				}
			}
		} else if err == nil {
			err = errors.JustError()
			if condition {
				log.Fatal(err)
			}
		}
	} else {
		err = errors.AssertTest(fmt.Sprint("length of slice is ", len(testdim)), "SetIkbdDim", fmt.Sprint("length of slice is ", len(dim)), "AssertInlineKeyboard")
		if condition {
			log.Fatal(err)
		}
	}

	return err
}

func (tfm *Formatter) AssertString(lineoftext string, condition bool) (err error) {
	if tfm.Message.Text != lineoftext {
		err = errors.AssertTest(fmt.Sprint(tfm.Message.Text), "WriteString", fmt.Sprint(lineoftext), "AssertString")
	}
	if condition {
		if err != nil {
			log.Fatal(err)
		}
	}

	return err
}

func (tfm *Formatter) AssertChatId(chatID int, condition bool) (err error) {
	if tfm.Message.ChatID != chatID {
		err = errors.AssertTest(fmt.Sprint(tfm.Message.ChatID), "WriteChatId", fmt.Sprint(chatID), "AssertChatId")
	}
	if condition {
		if err != nil {
			log.Fatal(err)
		}
	}
	return err
}
