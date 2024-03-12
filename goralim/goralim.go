package main

import(
    "fmt"
    "time"
)

type RateLimiter struct{
    maxReqPerSec int
    currReq int
}

func NewRateLimiter(maxReqPerSec int) *RateLimiter {
    return &RateLimiter{
        maxReqPerSec: maxReqPerSec,
        currReq: 0,
    }
}

func (rl *RateLimiter) isAllowed() bool{
    if rl.currReq < rl.maxReqPerSec{
        rl.currReq++
        return true
    }

    return false
}

func main(){
    rl := NewRateLimiter(100)

    for i:=0; i<= 120; i++ {
        if rl.isAllowed(){
            fmt.Printf("Request is Accepted: %d\n", i)
           }else {
             fmt.Printf("Request is Denied: %d\n",i)
           }
        time.Sleep(1*time.Millisecond)
    }

    time.Sleep(200*time.Millisecond)
}


