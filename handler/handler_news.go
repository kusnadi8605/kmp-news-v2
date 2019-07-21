package handler

import (
	"encoding/json"
	conf "kmp-news-v2/config"
	dts "kmp-news-v2/datastruct"
	log "kmp-news-v2/logging"
	mdl "kmp-news-v2/models"
	"net/http"

	"github.com/gomodule/redigo/redis"
	"gopkg.in/olivere/elastic.v7"
)

//GetNewsHandler ..
func GetNewsHandler(conn *conf.Connection, client *elastic.Client, index string, perpage int) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var NewsResponse dts.NewsResponse
		arrNews := []dts.NewsDetail{}

		page := req.URL.Query().Get("page")

		//redis key
		redisKEY := conf.Param.RedisKEY + page
		cacheNews, err := mdl.Get(redisKEY)

		if err == redis.ErrNil {
			res, _ := mdl.GetNewsElastic(conn, client, index, page, perpage)
			if len(res) < 1 {
				NewsResponse.ResponseCode = "301"
				NewsResponse.ResponseDesc = "data not found"
				json.NewEncoder(w).Encode(NewsResponse)

				log.Logf("Response News : %v", NewsResponse)

				return
			}

			//save data to redis
			jsonData, err := json.Marshal(res)

			err = mdl.SetTex(redisKEY, conf.Param.RedisEXP, jsonData)
			if err != nil {
				NewsResponse.ResponseCode = "401"
				NewsResponse.ResponseDesc = err.Error()
				json.NewEncoder(w).Encode(NewsResponse)

				log.Logf("Response News : %v", NewsResponse)

				return
			}

			log.Logf("Save data to redis : %v", res)

			NewsResponse.ResponseCode = "000"
			NewsResponse.ResponseDesc = "Success"
			NewsResponse.Payload = res
			json.NewEncoder(w).Encode(NewsResponse)

			log.Logf("Response News : %v", NewsResponse)

			return
		} else if err != nil {
			NewsResponse.ResponseCode = "401"
			NewsResponse.ResponseDesc = err.Error()
			json.NewEncoder(w).Encode(NewsResponse)

			log.Logf("Response News : %v", NewsResponse)

			return
		} else {

			err := json.Unmarshal([]byte(cacheNews), &arrNews)

			log.Logf("Get data redis : %v", arrNews)

			if err != nil {
				NewsResponse.ResponseCode = "401"
				NewsResponse.ResponseDesc = err.Error()
				json.NewEncoder(w).Encode(NewsResponse)

				log.Logf("Response News : %v", NewsResponse)

				return
			}

			NewsResponse.ResponseCode = "000"
			NewsResponse.ResponseDesc = "Success"
			NewsResponse.Payload = arrNews
			json.NewEncoder(w).Encode(NewsResponse)

			log.Logf("Response News : %v", NewsResponse)
			return
		}
	}

}
