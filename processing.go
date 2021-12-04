package main

import (
	"log"
	"sync"
)

type Transaction struct {
	UserId    int    `json:"user_id"`
	Timestamp uint32 `json:"timestamp"`
	Category  string `json:"category"`
	Card      string `json:"card"`
	Amount    int64  `json:"amount"`
}

type Report struct {
	UserID int   `json:"user_id"`
	Sum    int64 `json:"sum"`
	Count  int64 `json:"Count"`
}

func SumCategoryTransactionsStandalone(transactions []Transaction) map[int]*Report {
	result := make(map[int]*Report)
	for _, transaction := range transactions {
		_, found := result[transaction.UserId]
		if found {
			(result[transaction.UserId]).Sum += transaction.Amount
			result[transaction.UserId].Count++
		} else {
			result[transaction.UserId] = &Report{
				UserID: transaction.UserId,
				Sum:    transaction.Amount,
				Count:  1,
			}
		}
	}

	return result
}

func SumCategoryTransactionsMutex(transactions []Transaction, goroutines int) map[int]*Report {
	if transactions == nil {
		return nil
	}
	if len(transactions) < goroutines {
		goroutines = len(transactions)
	}

	wg := sync.WaitGroup{}
	wg.Add(goroutines)
	mu := sync.Mutex{}
	result := make(map[int]*Report)

	partSize := len(transactions) / goroutines
	for i := 0; i < goroutines; i++ {
		part := transactions[i*partSize : (i+1)*partSize]
		go func() {
			localResult := SumCategoryTransactionsStandalone(part)

			mu.Lock()
			for key, report := range localResult {
				_, found := result[key]

				if found != false {
					result[key].Count += report.Count
					result[key].Sum += report.Sum
				} else {
					result[key] = report
				}
			}
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	return result
}

func SumCategoryTransactionsChanel(transactions []Transaction, goroutines int) map[int]*Report {
	if transactions == nil {
		return nil
	}
	if len(transactions) < goroutines {
		goroutines = len(transactions)
	}

	result := make(map[int]*Report)
	ch := make(chan map[int]*Report)

	partSize := len(transactions) / goroutines
	for i := 0; i < goroutines; i++ {
		part := transactions[i*partSize : (i+1)*partSize]
		go func(ch chan<- map[int]*Report) {
			ch <- SumCategoryTransactionsStandalone(part)
		}(ch)
	}

	finished := 0
	for localResult := range ch {
		for key, report := range localResult {
			_, found := result[key]

			if found != false {
				result[key].Count += report.Count
				result[key].Sum += report.Sum
			} else {
				result[key] = report
			}
		}
		finished++
		if finished == goroutines {
			break
		}
	}

	return result
}

func SumCategoryTransactionsChanelStandalone(transactions []Transaction, goroutines int) map[int]*Report {
	if transactions == nil {
		return nil
	}
	if len(transactions) < goroutines {
		goroutines = len(transactions)
	}

	result := make(map[int]*Report)
	ch := make(chan map[int]*Report)

	partSize := len(transactions) / goroutines
	for i := 0; i < goroutines; i++ {
		part := transactions[i*partSize : (i+1)*partSize]
		go func(ch chan<- map[int]*Report) {
			localResult := make(map[int]*Report)
			for _, transaction := range part {

				_, found := localResult[transaction.UserId]
				if found {
					(localResult[transaction.UserId]).Sum += transaction.Amount
					localResult[transaction.UserId].Count++
				} else {
					localResult[transaction.UserId] = &Report{
						UserID: transaction.UserId,
						Sum:    transaction.Amount,
						Count:  1,
					}
				}
			}

			ch <- localResult
		}(ch)
	}

	finished := 0
	for localResult := range ch {
		for key, report := range localResult {
			_, found := result[key]

			if found != false {
				result[key].Count += report.Count
				result[key].Sum += report.Sum
			} else {
				result[key] = report
			}
		}
		finished++
		if finished == goroutines {
			break
		}
	}

	return result
}

func SumCategoryTransactionsMutexStandalone(transactions []Transaction, goroutines int) map[int]*Report {
	if transactions == nil {
		log.Fatalln("empty transactions")
	}

	if len(transactions) < goroutines {
		goroutines = len(transactions)
	}

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	result := make(map[int]*Report)

	partSize := len(transactions) / goroutines
	for i := 0; i < goroutines; i++ {
		part := transactions[i*partSize : (i+1)*partSize]
		wg.Add(1)
		//
		go func() {
			localResult := make(map[int]*Report)
			for _, transaction := range part {

				_, found := localResult[transaction.UserId]
				if found {
					(localResult[transaction.UserId]).Sum += transaction.Amount
					localResult[transaction.UserId].Count++
				} else {
					localResult[transaction.UserId] = &Report{
						UserID: transaction.UserId,
						Sum:    transaction.Amount,
						Count:  1,
					}
				}
			}

			mu.Lock()
			for key, report := range localResult {
				_, found := result[key]

				if found {
					result[key].Count += report.Count
					result[key].Sum += report.Sum
				} else {
					result[key] = report
				}
			}
			mu.Unlock()

			wg.Done()
		}()
	}
	wg.Wait()

	return result
}
