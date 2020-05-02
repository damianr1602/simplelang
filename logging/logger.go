package logging

import (
	"flag"
	"log"
	"os"

	"github.com/damianr1602/simplelang/config"
)

var (
	Log *log.Logger
)

func init() {
	flag.Parse()
	var file, err1 = os.OpenFile(config.LogPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err1 != nil {
		panic(err1)
	}
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
}
