package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

func main() {
	dataCh := make(chan int)
	evenCh := make(chan int)
	oddCh := make(chan int)
	evenSumCh := make(chan int)
	oddSumCh := make(chan int)

	processWg := &sync.WaitGroup{}
	fileWg := &sync.WaitGroup{}
	fileWg.Add(1)
	go source("data1.dat", dataCh, fileWg)
	fileWg.Add(1)
	go source("data2.dat", dataCh, fileWg)

	processWg.Add(4)
	go splitter(dataCh, evenCh, oddCh, processWg)
	go sum(evenCh, evenSumCh, processWg)
	go sum(oddCh, oddSumCh, processWg)
	go merger("result.dat", evenSumCh, oddSumCh, processWg)

	fileWg.Wait()
	close(dataCh)

	processWg.Wait()
	fmt.Println("Done")
}

func source(fileName string, ch chan int, wg *sync.WaitGroup) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := scanner.Text()
		no, err := strconv.Atoi(data)
		if err != nil {
			continue
		}
		ch <- no
	}
	wg.Done()
}

func splitter(dataCh chan int, evenCh chan int, oddCh chan int, wg *sync.WaitGroup) {
	defer func() {
		close(evenCh)
		close(oddCh)
		wg.Done()
	}()
	for no := range dataCh {
		if no%2 == 0 {
			evenCh <- no
		} else {
			oddCh <- no
		}
	}
}

func sum(noCh chan int, sumCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	result := 0
	for no := range noCh {
		result += no
	}
	sumCh <- result
}

func merger(resultFile string, evenSumCh chan int, oddSumCh chan int, wg *sync.WaitGroup) {
	file, err := os.Create(resultFile)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		file.Close()
		close(evenSumCh)
		close(oddSumCh)
		wg.Done()
	}()
	for i := 0; i < 2; i++ {
		select {
		case evenSum := <-evenSumCh:
			file.WriteString(fmt.Sprintf("Even Total = %d\n", evenSum))
		case oddSum := <-oddSumCh:
			file.WriteString(fmt.Sprintf("Odd Total = %d\n", oddSum))
		}
	}
}
