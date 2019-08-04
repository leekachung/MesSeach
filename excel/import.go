package excel

import (
	"lcb-go/cache"
	"fmt"
	"encoding/json"
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
		key, _ := json.Marshal(i[0] + ":" + i[1])
		value := make(map[string]string)
		value["no"] = i[2]
		value["m_seat"] = i[3]
		value["l_seat"] = i[4]
		value["b_seat"] = i[5]
		/*
		value["a_seat"] = i[6]
		value["e_seat"] = i[7]
		value["hotel"] = i[8]
		value["room"] = i[9]
		value["car"] = i[10]
		*/
		data, _ := json.Marshal(value)
		conn.Do("SETNX", key, data)
	}
}
