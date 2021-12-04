package main

import (
	"log"
	"time"
)

func main() {
	//var staticTransactions = GenerateStaticTransactions(10, 1_000_00, 100)
	//ExportToJson(staticTransactions, "1M_transactions.json")

	transaction := ImportFromJson("10M_transactions.json")

	//start := time.Now()
	//reports := SumCategoryTransactionsMutexStandalone(transaction, 4)
	//log.Println(time.Since(start))

	//start := time.Now()
	//reports := SumCategoryTransactionsMutex(transaction, 4)
	//log.Println(time.Since(start))

	//start := time.Now()
	//reports := SumCategoryTransactionsChanel(transaction, 4)
	//log.Println(time.Since(start))

	//start := time.Now()
	//reports := SumCategoryTransactionsChanelStandalone(transaction, 4)
	//log.Println(time.Since(start))

	start := time.Now()
	reports := SumCategoryTransactionsStandalone(transaction)
	log.Println(time.Since(start))

	ExportToJson(reports, "result.json")
}
