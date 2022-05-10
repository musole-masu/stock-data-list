package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rodaine/table"
)
type Quote struct {
	SYMBOL string `json:"01. symbol"`
	OPEN string `json:"02. open"`
	HIGH string `json:"03. high"`
	LOW string `json:"04. low"`
	PRICE string `json:"05. price"`
	VOLUME  string `json:"06. volume"`
	LATEST_TRADING_DAY string `json:"07. latest trading day"`
	PREVIOUS_CLOSE string `json:"08. previous close"`
	CHANGE string `json:"09. change"`
	CHANGE_PERCENT string `json:"10. change percent"`
}

type Stock struct {
	DETAIL Quote `json:"Global Quote"`
}
type WatchList struct {
	stock []Quote
}

func getStock(s string, ch chan Stock) Stock{
	var stock Stock

	
	api := "https://alpha-vantage.p.rapidapi.com/query?function=GLOBAL_QUOTE&symbol="+s+"&datatype=json"

	req, _ := http.NewRequest("GET", api, nil);

	req.Header.Add("X-RapidAPI-Host", "alpha-vantage.p.rapidapi.com")
	req.Header.Add("X-RapidAPI-Key", "HERE PUT YOUR RAPIDAPI KEY")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println("No response from request")
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	
	if err := json.Unmarshal(body, &stock); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}
	ch<- stock
	return stock
}

func genTable(watchList []Quote){

	tbl := table.New("SYMBOL", "OPEN", "HIGH", "LOW", "PRICE", "VOLUME", "PREVIOUS-CLOSE", "CHANGE")

	for _, stock := range watchList {
		tbl.AddRow(stock.SYMBOL, stock.OPEN, stock.HIGH, stock.LOW, stock.PRICE, stock.VOLUME, stock.PREVIOUS_CLOSE, stock.CHANGE)
	}

	tbl.Print()

}
