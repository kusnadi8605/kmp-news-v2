package models

import (
	conf "kmp-news-v2/config"
	dts "kmp-news-v2/datastruct"
)

//GetNewsDetail ..
func GetNewsDetail(conn *conf.Connection, ID int64) ([]dts.NewsDetail, error) {
	arrNewsDetail := []dts.NewsDetail{}
	strNewsDetail := dts.NewsDetail{}

	strQuery := `select * from news where id=?`

	rows, err := conn.Query(strQuery, ID)

	defer rows.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {

		err := rows.Scan(&strNewsDetail.ID,
			&strNewsDetail.Author,
			&strNewsDetail.Body,
			&strNewsDetail.Created,
		)

		if err != nil {
			return nil, err
		}

		if err = rows.Err(); err != nil {
			return nil, err
		}

		arrNewsDetail = append(arrNewsDetail, strNewsDetail)
	}

	return arrNewsDetail, nil
}
