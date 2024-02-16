package executer

import (
	"bytes"
	"encoding/json"
	"firstgobot/byogram/errors"
	types "firstgobot/byogram/types"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

const None = -1

type Telegram struct {
	Url string
}

type TelegramTest struct {
	Url string
}

type Entry struct {
	UserId int
	Ch     chan *types.TelegramResponse
}

type RegTable struct {
	Reg []Entry
}

func GetpostRequest(url string, Buffer *bytes.Buffer, contenttype string) (err error) {
	var (
		request  *http.Request
		response *http.Response
		client   *http.Client
	)

	request, err = http.NewRequest("POST", url, Buffer)
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Set("Content-Type", contenttype)
	client = &http.Client{}
	response, err = client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))

	return err
}

func Send(buf *bytes.Buffer, function, contenttype string) {
	var (
		err error
		url string
	)
	url = fmt.Sprintf("%sbot%s/%s", types.HttpsRequest, types.TelebotToken, function)
	err = GetpostRequest(url, buf, contenttype)

	if err != nil {
		log.Fatal(err)
	}
}

func Updates(offset *int, telegramResponse *types.TelegramResponse) (err error) {
	var (
		response *http.Response
		body     []byte
	)

	url := fmt.Sprintf(types.HttpsRequest+"bot%s/getUpdates?limit=1&offset=%d", types.TelebotToken, *offset)
	response, err = http.Get(url)
	if err == nil {
		body, err = io.ReadAll(response.Body)
	}
	if err == nil {
		err = handlerTelegramResponse(body, telegramResponse)
	}
	response.Body.Close()
	return err
}

func handlerTelegramResponse(response []byte, telegramResponse *types.TelegramResponse) (err error) {
	err = json.Unmarshal(response, &telegramResponse)
	if err == nil {
		if !telegramResponse.Ok {
			err = fmt.Errorf(fmt.Sprintf("Telegram API вернул ошибку: %s", telegramResponse.Error.Message))
		}
	}

	return err
}

func (reg *RegTable) Seeker(chatID int) (index int) {
	var i int
	index = None
	if len(reg.Reg) != 0 {
		for i < len(reg.Reg) && reg.Reg[i].UserId != chatID {
			i++
		}
		if reg.Reg[i].UserId == chatID {
			index = i
		}
	}
	return index
}

func (reg *RegTable) NewIndex() (newIndex int) {
	newIndex = len(reg.Reg)
	reg.Reg = append(reg.Reg, Entry{})
	return newIndex
}

func handlerOffsetResponse(response []byte, offset *int) (err error) {
	var telegramResponse types.JustForUpdate

	err = json.Unmarshal(response, &telegramResponse)
	if err == nil {
		if len(telegramResponse.Result) > 0 {
			for _, storage := range telegramResponse.Result {
				*offset = storage.ID
			}
		} else {
			err = errors.UpdatesMisstakes("Telegram returned an empty data of telegramResponse")
		}
	}

	return err
}

func RequestOffset(TelebotToken string, offset *int) (err error) {
	var (
		response *http.Response
		body     []byte
	)

	Url := fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates?limit=1", url.PathEscape(TelebotToken))
	response, err = http.Get(Url)
	if err == nil {
		body, err = io.ReadAll(response.Body)
	}
	if err == nil {
		err = handlerOffsetResponse(body, offset)
	}
	response.Body.Close()

	return err
}

func (test *TelegramTest) RequestOffset(TelebotToken string, offset *int) error {
	return nil
}

func (test *TelegramTest) Updates(TelebotToken string, offset *int, telegramResponse *types.TelegramResponse) error {

	return nil
}

/*
type entry struct { user int; ch *chan<-Response }

var reg []entry

func this (user int) int //return index of user in the registry
{
	return none
}

funr free () int // return free index in the registry
{
	return none
}

func Writer ()
{
	var rs Response
	getRs(rs)
	n = this(rs.user)
	if n != none { reg[n].ch <- rs }
	else {
		n = free()
		reg[n].user = rs.user
		reg[n].ch = /* create a new channel
		reg[n].ch <- rs
		go Worker(reg[n].ch)
	}
}
*/
