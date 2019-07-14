package main

import (
	"fmt"
	"io"
	"os"
)

// duck typing
// 只关注行为：描述事物外部行为，而不是内部结构

// 接口把所有具有共性的方法定义在一起
// 任何其他类型只要实现了这些方法就是实现了这个接口

func main() {
	chris := Person{"chris"}
	// chris.Say("i can say!!!")
	SpeakAlphabet(&chris)

	console := new(SpeakerWriter)
	console.w = os.Stdout
	SpeakAlphabet(console)
}

type Speaker interface {
	Say(string)
}

type Person struct {
	name string
}

func (p *Person) Say(msg string) {
	fmt.Printf("%s:%s\n", p.name, msg)
}

func SpeakAlphabet(speaker Speaker) {
	speaker.Say("i can say!")
}

type SpeakerWriter struct {
	w io.Writer
}

func (sw *SpeakerWriter) Say(msg string) {
	// sw.Say(msg)
	io.WriteString(sw.w, msg)
}
