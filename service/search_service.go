package service

import (
	"lcb-go/cache"
	"github.com/gomodule/redigo/redis"
	"encoding/json"
	"fmt"
)

type SearchService struct {
	Phone string `form:"phone" binding:"required,min=11,max=11"`
	Name string `form:"name" binding:"required,min=2"`
}

func (service *SearchService) Search() map[string]string {
	// get redis conn
	conn := cache.GetConn(cache.RedisPool)
	defer conn.Close()
	
	key, _ := json.Marshal(service.Phone + ":" + service.Name)
	value, err := redis.Bytes(conn.Do("GET", key))
	fmt.Println(value)
	if err == nil {
		 var imap map[string]string
		 errShal := json.Unmarshal(value, &imap)
		 if errShal != nil {
		 	return nil
		 }
		return imap
	} else {
		return nil
	}
}
