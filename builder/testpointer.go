package builder

import (
	"fmt"
	"strings"
)

type talker interface {
	talk() string
	sleep() string
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

type admin struct {
}

func (a admin) talk() string {
	return "ớ ớ"
}

// nếu define func (a *admin) thì admin{} sẽ không có method set là sleep này, vì receiver là pointer

func (a admin) sleep() string {
	return "hi hi"
}
