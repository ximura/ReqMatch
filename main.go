package main

import (
	"errors"
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ximura/ReqMatch/internal"
)

type pairInfo struct {
	id        uint32
	timeStart time.Time
}

type statInfo struct {
	id    uint32
	delay int64
}

func main() {
	mux := http.NewServeMux()
	var pairCounter atomic.Uint32
	newReqChan := make(chan pairInfo)

	//stats := make(map[uint32]int64, 50000)
	var m sync.RWMutex
	stats := internal.NewSortedArray(50000, func(a, b statInfo) int {
		if a.delay == b.delay {
			return 0
		}
		if a.delay > b.delay {
			return 1
		}

		return -1
	}, func(a statInfo) string {
		return fmt.Sprintf("\"%d\":%d", a.id, a.delay)
	})

	mux.HandleFunc("POST /join", func(w http.ResponseWriter, r *http.Request) {
		id := pairCounter.Add(1)
		select {
		case newReqChan <- pairInfo{id: id, timeStart: time.Now()}:
			response := fmt.Sprintf("%d First\n", id)
			w.Write([]byte(response))

		case pairInfo := <-newReqChan:
			d := time.Now().Sub(pairInfo.timeStart)
			m.Lock()
			stats.Insert(statInfo{
				id:    pairInfo.id,
				delay: d.Nanoseconds(),
			})
			//stats[pairInfo.id] = d.Nanoseconds()
			m.Unlock()

			response := fmt.Sprintf("%d Second\n", pairInfo.id)
			w.Write([]byte(response))

		case <-time.After(10 * time.Second):
			w.Write([]byte("Timeout. No more connected clients.\n"))
		}
	})

	mux.HandleFunc("GET /stats", func(w http.ResponseWriter, r *http.Request) {
		m.RLock()
		jsonString, err := stats.Marshal()
		m.RUnlock()

		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		w.Write(jsonString)
	})

	fmt.Println("Starting server on 3000")
	err := http.ListenAndServe(":3000", mux)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("server closed")
		return
	}
	if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		return
	}
}
