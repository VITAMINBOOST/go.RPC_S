package tasks

import (
	"runtime"
	"strconv"
)

type RsrcStruct struct {
	Title, Usage string
}

type Rsrc int

var rsrcSlice []RsrcStruct

func (r *Rsrc) GetSlice(title string, reply *[]RsrcStruct) error {
	*reply = rsrcSlice

	return nil
}

func (r *Rsrc) MakeRsrcUsage(rsrc RsrcStruct, reply *RsrcStruct) error {
	if rsrc.Title == "Mem" {
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		rsrc.Usage = strconv.Itoa(int(m.Sys/1024/1024)) + " MiB"
	}

	rsrcSlice = append(rsrcSlice, rsrc)
	*reply = rsrc
	return nil
}
