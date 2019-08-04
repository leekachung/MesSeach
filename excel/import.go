package excel

import (
	"lcb-go/cache"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)

func RowToArr(rows [][]string, c chan []string)  {
	// push content to array
	for i, row := range rows {
		if i != 0 {
			temp := make([]string, 0)
			for ii, colCell := range row {
				if ii != 0 {
					temp = append(temp, colCell)
				}	
			}
			c <- temp
		}
	}
	defer close(c)
}

func Import() {
	// open file
	f, err := excelize.OpenFile("./excel/1.xlsx")
	if err != nil {
		fmt.Println(err)
	}
	// get file content
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
	}
	// make a channel
	c := make(chan []string)
	go RowToArr(rows, c)

	/**
	* Test Excel Data
	* for i := range c {
	* 	fmt.Println(i)
	* }
	*/

	// get redis connect
	conn := cache.GetConn(cache.RedisPool)
	defer conn.Close()

	// flush redis data
	conn.Do("FLUSHALL")

	// add data to redis
	for i := range c {
		// var key string
		// var value string
		// conn.Do("SETNX", key, value)
		fmt.Println(i)
	}
}
