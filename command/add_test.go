package command

import (
	"os"
	"testing"

	"github.com/wlbr/shorty/base"
	"github.com/wlbr/shorty/gotils"
	"github.com/wlbr/shorty/store"
)

var config *gotils.Config
var sink base.Store

func TestMain(m *testing.M) {

	config = &gotils.Config{}

	gotils.ReadConfig(config, "../shorty.ini")
	config.Log.LogLevel = "All"
	config.Log.ActiveLogLevel, _ = gotils.LogLevelString(config.Log.LogLevel)
	gotils.AddAdditionalExpVars(config)

	sink = store.NewPostgressStore()
	gotils.LogInfo("Connecting to store %v", sink)
	sink.Connect(config)
	defer sink.Disconnect()

	code := m.Run()
	os.Exit(code)
}

func TestBasicAdd(t *testing.T) {
	sink.Add("wlbr", "https://www.wlbr.de", "mwolber")
	sink.AddShortURL(&base.ShortURL{LocalPart: "takeawaysm", LongURL: "https://appear.in/takeawaysm", UserName: "jbroos"})
	sink.Add("inoergo", "https://hangouts.google.com/hangouts/_/inovex.de/inoergo", "mwolber")
	sink.AddShortURL(&base.ShortURL{LocalPart: "inoergo2", LongURL: "https://hangouts.google.com/hangouts/_/inovex.de/inoergo2", UserName: "mwolber"})
	sink.AddShortURL(&base.ShortURL{LocalPart: "ergo", LongURL: "https://hangouts.google.com/hangouts/_/inovex.de/inoergo", UserName: "mwolber"})

}
