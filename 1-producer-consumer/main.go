//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer scenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"time"
)

const buf = 1024

func producer(stream Stream, ch chan *Tweet) (tweets []*Tweet) {
	defer close(ch)
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			return tweets
		}

		//tweets = append(tweets, tweet)
		ch <- tweet
	}

}

func consumer(tweets []*Tweet, ch chan *Tweet) {
	for {
		t, ok := <-ch
		if !ok {
			return
		}
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}
}

func main() {
	start := time.Now()
	stream := GetMockStream()

	ch := make(chan *Tweet, buf)

	// Consumer
	go consumer(nil, ch)

	// Producer
	producer(stream, ch)

	fmt.Printf("Process took %s\n", time.Since(start))
}
