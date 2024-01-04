package main

import (
	"fmt"
	"runtime"
	"sort"
	"sync"
)

func RunPipeline(cmds ...cmd) {
	var wg sync.WaitGroup

	in := make(chan interface{})
	var out chan interface{}

	wg.Add(len(cmds))
	for _, command := range cmds {
		out = make(chan interface{})

		go func(command cmd, in, out chan interface{}) {
			command(in, out)
			close(out)
			wg.Done()
		}(command, in, out)

		in = out
	}
	wg.Wait()
}

func SelectUsers(in, out chan interface{}) {
	usersChan := make(chan User)
	wg := new(sync.WaitGroup)
	for email := range in {
		wg.Add(1)
		go func(email string) {
			usersChan <- GetUser(email)
			wg.Done()
		}(email.(string))
	}

	go func() {
		wg.Wait()
		close(usersChan)
	}()

	uniqueIDs := make(map[uint64]User)
	for user := range usersChan {
		_, exists := uniqueIDs[user.ID]
		if !exists {
			uniqueIDs[user.ID] = user
			out <- user
		}
	}
}

func SelectMessages(in, out chan interface{}) {
	wg := new(sync.WaitGroup)
	for user := range in {
		wg.Add(1)
		go func(userBatch ...User) {
			messages, err := GetMessages(userBatch...)
			if err != nil {
				wg.Done()
				return
			}
			for _, message := range messages {
				out <- message
			}
			wg.Done()
		}(user.(User))
	}
	wg.Wait()
}

func checkMessage(in, out chan interface{}) {
	for msgID := range in {
		msgIDobject := msgID.(MsgID)
		ok, err := HasSpam(msgIDobject)
		if err != nil {
			out <- MsgData{
				ID:      msgIDobject,
				HasSpam: ok,
			}
		}
		runtime.Gosched()
	}
}

func CheckSpam(in, out chan interface{}) {
	const workersNum = 5
	for i := 0; i < workersNum; i++ {
		go checkMessage(in, out)
	}
}

func CombineResults(in, out chan interface{}) {
	data := []MsgData{}
	for msgData := range in {
		data = append(data, msgData.(MsgData))
	}

	sort.Slice(data, func(i, j int) bool {
		return (data[i].ID < data[j].ID) && (!data[i].HasSpam && data[i].HasSpam)
	})

	for _, msg := range data {
		out <- fmt.Sprintf("%v %v", msg.HasSpam, msg.ID)
	}
}
