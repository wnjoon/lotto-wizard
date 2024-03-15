package lottowizard

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

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

var lottoCountMap map[int]int

// 가장 최근에 진행된 로또 회차번호 조회
func getLatestRoundNumber() (int, error) {
	resp, err := http.Get("https://dhlottery.co.kr/common.do?method=main")
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

// 해당 회차의 로또 당첨 내역 조회
func getLottoNumberByRound(round int) (Lotto, error) {

	resp, err := http.Get(fmt.Sprint("https://www.dhlottery.co.kr/common.do?method=getLottoNumber&drwNo=", strconv.FormatInt(int64(round), 10)))
	if err != nil {
		return Lotto{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Lotto{}, err
	}

	lotto := Lotto{}

	err = json.Unmarshal(data, &lotto)
	if err != nil {
		return Lotto{}, err
	}
	return lotto, nil
}

func GetLottoList() (map[int]int, error) {

	latestRound, err := getLatestRoundNumber()
	if err != nil {
		return nil, err
	}

	lottoCountMap = make(map[int]int)

	var wg sync.WaitGroup
	wg.Add(latestRound)

	for i := 1; i <= latestRound; i++ {
		defer wg.Done()
		go func(i int) {
			l, err := getLottoNumberByRound(i)
			if err != nil {
				log.Panic(i, "th round : failed to get from url")

			}
			fmt.Print(i, "th round : success to get from url")
			lottoCountMap[l.DrwtNo1] = lottoCountMap[l.DrwtNo1] + 1
			lottoCountMap[l.DrwtNo2] = lottoCountMap[l.DrwtNo2] + 1
			lottoCountMap[l.DrwtNo3] = lottoCountMap[l.DrwtNo3] + 1
			lottoCountMap[l.DrwtNo4] = lottoCountMap[l.DrwtNo4] + 1
			lottoCountMap[l.DrwtNo5] = lottoCountMap[l.DrwtNo5] + 1
			lottoCountMap[l.DrwtNo6] = lottoCountMap[l.DrwtNo6] + 1
		}(i)
	}
	wg.Wait()

	for {
		if len(lottoCountMap) == latestRound {
			break
		}
		fmt.Println(lottoCountMap)
		time.Sleep(2000)

	}
	return lottoCountMap, nil
}
