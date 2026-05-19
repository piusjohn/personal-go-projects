package main

import "fmt"

func main() {
    var mediaPlayer Media = Audio{ duration: 120 }

    // Enter Code Here
    if t, ok := mediaPlayer.(Audio); ok {
      fmt.Println("correct", t.duration)
    }else{
      fmt.Println("incorrect")
    }

    switch m := mediaPlayer.(type){
      case Audio:
        fmt.Println("Handling an audio device", m.duration)
      case Video:
        fmt.Println("Handling a video device", m.frameRate)
      default:
        fmt.Println("device not available")
    }
}

type Media interface {
    Play()
}

type Audio struct {
    duration int
}

func (a Audio) Play() {
    fmt.Println("Playing audio for", a.duration, "seconds")
}

type Video struct {
    frameRate int
}

func (v Video) Play() {
    fmt.Println("Playing video at", v.frameRate, "fps")
}