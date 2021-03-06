package main

import (
	"fmt"
	"strings"
)

type BitFlag int

const (
	Active BitFlag = 1 << iota
	Send
	Receive
)

func main() {

	flag := Active | Send
	BitFlag.String(flag)
}

func (flag BitFlag) String() string {

	var flags []string
	if flag&Active == Active {
		flags = append(flags, "Active")
	}
	if flag&Send == Send {
		flags = append(flags, "Send")
	}
	if flag&Receive == Receive {
		flags = append(flags, "Receive")
	}
	if len(flags) > 0 {
		return fmt.Sprintf("%d(%s)", int(flag), strings.Join(flags, "|"))
	}
	return "0()"

}
