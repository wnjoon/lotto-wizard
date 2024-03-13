package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
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

var lottoMap = make(map[int]int)

func main() {
	step1()
	fmt.Println("====")
	for k, v := range lottoMap {
		fmt.Println(k, ":", v)
	}
	fmt.Println("====")
}

func step1() error {
	r, err := getLatestRound()
	if err != nil {
		return err
	}

	for i := 1; i <= r; i++ {
		l, err := getLottoNumberByRound(i)
		if err != nil {
			return err
		}
		setLottoMap(l)
	}
	return nil
}

func setLottoMap(l Lotto) {
	lottoMap[l.DrwtNo1] = lottoMap[l.DrwtNo1] + 1
	lottoMap[l.DrwtNo2] = lottoMap[l.DrwtNo2] + 1
	lottoMap[l.DrwtNo3] = lottoMap[l.DrwtNo3] + 1
	lottoMap[l.DrwtNo4] = lottoMap[l.DrwtNo4] + 1
	lottoMap[l.DrwtNo5] = lottoMap[l.DrwtNo5] + 1
	lottoMap[l.DrwtNo6] = lottoMap[l.DrwtNo6] + 1
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

func getLatestRound() (int, error) {
	resp, err := http.Get(MainUrl)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	html, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	round, err := strconv.Atoi(html.Find("#lottoDrwNo").Text())
	if err != nil {
		return 0, err
	}

	return round, nil
}
