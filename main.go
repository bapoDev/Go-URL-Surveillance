package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

type CheckResult struct {
	URL        string
	Status     string
	StatusCode int
	Latency    time.Duration
	Error      error
}

func CheckService(url string) CheckResult {
	var beginning = time.Now()

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	var response, err = client.Get(url)
	var latency = time.Since(beginning)
	if err != nil {
		status := "DOWN"
		if timeoutErr, ok := err.(interface{ Timeout() bool }); ok && timeoutErr.Timeout() {
			status = "TIMEOUT"
		}

		return CheckResult{
			URL:        url,
			Status:     status,
			StatusCode: 0,
			Latency:    latency,
			Error:      err,
		}
	}

	defer response.Body.Close()

	return CheckResult{
		URL:        url,
		Status:     "UP",
		StatusCode: response.StatusCode,
		Latency:    latency,
		Error:      nil,
	}
}

func CheckAndSend(url string, resultsCh chan CheckResult, wg *sync.WaitGroup) {
	defer wg.Done()

	result := CheckService(url)

	resultsCh <- result
}

func ShowResult(urlength int, finalResults []CheckResult) {
	const headerFormat = "| %-*s | %-10s | %-5s | %-10s | %s\n"
	const lineFormat = "| %-*s | %-10s | %-5d | %-10s | %s\n"

	fmt.Println("\n--- Surveillance results ---")
	fmt.Printf(headerFormat,
		urlength,
		"URL",
		"STATUS",
		"CODE",
		"LATENCY",
		"ERROR")
	fmt.Println("------------------------------------------------------------------------------------------")

	for _, result := range finalResults {
		errorMsg := "-"
		if result.Error != nil {
			errorMsg = result.Error.Error()
		}

		fmt.Printf(lineFormat,
			urlength,
			result.URL,
			result.Status,
			result.StatusCode,
			result.Latency.Round(time.Millisecond),
			errorMsg)
	}
}

func FillUrls() (resultsCh chan CheckResult, urls []string, urlength int) {
	resultsCh = make(chan CheckResult)

	file, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatal("Please create a urls.txt with an URL each line.")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		urls = append(urls, line)
		x := len(line)
		if x > urlength {
			urlength = x
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	return
}

func main() {
	var wg sync.WaitGroup
	var wgCollector sync.WaitGroup

	resultsCh, urls, urlength := FillUrls()

	for _, url := range urls {
		wg.Add(1)
		go CheckAndSend(url, resultsCh, &wg)
	}

	var finalResults []CheckResult

	wgCollector.Go(func() {
		for result := range resultsCh {
			finalResults = append(finalResults, result)
		}
	})

	wg.Wait()

	close(resultsCh)

	wgCollector.Wait()

	ShowResult(urlength, finalResults)
}
