package model

import (
	"fmt"
	"strconv"

	"github.com/jackc/pgx"
)

func WriteExampleGiftToDb(pool *pgx.Conn, eg ExampleGift) (ExampleGift, error) {
	giftIdStr := strconv.FormatInt(eg.GiftId, 10)

	err := pool.QueryRow(fmt.Sprintf("INSERT INTO examplegifts (gift_id, name, price) VALUES ('%s','%s', '%d') RETURNING gift_id;", giftIdStr, eg.Name, eg.Price)).Scan(&eg.GiftId)

	if err != nil {
		return ExampleGift{}, err
	}

	return eg, nil
}

func GetExampleGiftFromDB(pool *pgx.Conn, giftId int64) (ExampleGift, error) {
	eg := ExampleGift{
		GiftId: giftId,
	}

	var gid int
	err := pool.QueryRow(fmt.Sprintf("SELECT gift_id, name, price FROM examplegifts where gift_id = '%d';", giftId)).Scan(&gid, &eg.Name, &eg.Price)

	if err != nil {
		panic(err)
	}

	return eg, nil
}

func GetAllExampleGiftsFromDB(pool *pgx.Conn) ([]ExampleGift, error) {
	rows, err := pool.Query("SELECT  gift_id, name, price FROM examplegifts;")

	if err != nil {
		panic(err)
	}

	results := []ExampleGift{}

	defer rows.Close()

	for rows.Next() {
		eg := ExampleGift{}
		err := rows.Scan(&eg.GiftId, &eg.Name, &eg.Price)

		if err != nil {
			panic(err)
		}

		results = append(results, eg)
	}

	return results, nil
}
