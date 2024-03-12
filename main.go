package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type Lotto struct {
	TotSellamnt    int64  `json:"totSellamnt"`
	ReturnValue    string `json:"returnValue"`
	DrwNoDate      string `json:"drwNoDate"`
	FirstWinamnt   int    `json:"firstWinamnt"`
	DrwtNo6        int    `json:"drwtNo6"`
	DrwtNo4        int    `json:"drwtNo4"`
	FirstPrzwnerCo int    `json:"firstPrzwnerCo"`
	DrwtNo5        int    `json:"drwtNo5"`
	BnusNo         int    `json:"bnusNo"`
	FirstAccumamnt int64  `json:"firstAccumamnt"`
	DrwNo          int    `json:"drwNo"`
	DrwtNo2        int    `json:"drwtNo2"`
	DrwtNo3        int    `json:"drwtNo3"`
	DrwtNo1        int    `json:"drwtNo1"`
}

const GetLottoNumberUrl string = "https://www.dhlottery.co.kr/common.do?method=getLottoNumber&drwNo="
const MainUrl string = "https://dhlottery.co.kr/common.do?method=main"

func main() {
	l, err := getLottoNumberByRound(100)
	if err != nil {
		panic(err)
	}
	fmt.Println(l)
}

func getLottoNumberByRound(round int) (Lotto, error) {

	lotto := Lotto{}

	resp, err := http.Get(fmt.Sprint(GetLottoNumberUrl, strconv.FormatInt(int64(round), 10)))
	if err != nil {
		return Lotto{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Lotto{}, err
	}

	err = json.Unmarshal(data, &lotto)
	if err != nil {
		return Lotto{}, err
	}
	return lotto, nil
}
