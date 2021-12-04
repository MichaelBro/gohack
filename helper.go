package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"
)

func GenerateStaticTransactions(users, transactionsPerUser, transactionAmount int) []Transaction {
	var timestamp = 1577836800 // 01.01.2020
	transactions := make([]Transaction, users*transactionsPerUser)

	rand.Seed(time.Now().UnixNano())

	for index := range transactions {
		timestamp += 60 * 60 * 24

		switch index % 100 {
		case 0:
			transactions[index] = Transaction{
				UserId:    rand.Intn(10) + 1,
				Amount:    int64(transactionAmount),
				Timestamp: uint32(timestamp),
				Category:  "Развлечения",
				Card:      "7773737",
			}
		case 20:
			transactions[index] = Transaction{
				UserId:    rand.Intn(10) + 1,
				Amount:    int64(transactionAmount),
				Timestamp: uint32(timestamp),
				Category:  "Транспорт",
				Card:      "5553535",
			}
		default:
			transactions[index] = Transaction{
				UserId:    rand.Intn(10) + 1,
				Amount:    int64(transactionAmount),
				Timestamp: uint32(timestamp),
				Category:  "Супермаркеты",
				Card:      "6663636",
			}
		}
	}
	return transactions
}

func ExportToJson(items interface{}, toFileName string) {
	jsonString, err := json.Marshal(items)

	if err != nil {
		log.Fatalln(err)
	}

	err = ioutil.WriteFile(toFileName, jsonString, os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}

}

func ImportFromJson(filePath string) []Transaction {
	file, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Fatalln(err)
	}

	var decoded []Transaction

	err = json.Unmarshal(file, &decoded)

	if err != nil {
		log.Fatalln(err)
	}

	return decoded
}