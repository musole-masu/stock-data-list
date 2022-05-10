package main

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {

	var watchList []Quote
	ch := make(chan Stock)
	
	router := gin.Default()

	router.GET("/stocks/", func (c *gin.Context){
		stockSymbol := c.Query("symbol")
		listOfSymbol := strings.Split(stockSymbol, ",")

		for _, symbol := range listOfSymbol{
			go getStock(symbol, ch)
		}

		for i := 0; i < len(listOfSymbol); i++ {
			result := <-ch

			watchList = append(watchList, Quote{
				SYMBOL: result.DETAIL.SYMBOL,
				OPEN: result.DETAIL.OPEN, 
				HIGH: result.DETAIL.HIGH, 
				LOW: result.DETAIL.LOW, 
				PRICE: result.DETAIL.PRICE, 
				VOLUME: result.DETAIL.VOLUME, 
				LATEST_TRADING_DAY: result.DETAIL.LATEST_TRADING_DAY, 
				PREVIOUS_CLOSE: result.DETAIL.PREVIOUS_CLOSE, 
				CHANGE: result.DETAIL.CHANGE, 
				CHANGE_PERCENT: result.DETAIL.CHANGE_PERCENT,
			})

		}

		genTable(watchList)
	})

	router.Run()

}

