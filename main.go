package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	g "github.com/AllenDang/giu"
)

var text = cmd("echo", "Hello, World!", "a")

type Config struct {
	Entries [][][]string
}

var conf Config

func init() {
	b, err := ioutil.ReadFile("./config.json")
	if err != nil {
		b, err = ioutil.ReadFile(os.Getenv("HOME") + "/.config/servermon.json")
	}
	if err != nil {
		log.Fatal("No config found at ./config.json and", os.Getenv("HOME")+"/.config/servermon.json")
	}
	err = json.Unmarshal(b, &conf)
	if err != nil {
		log.Fatal(err)
	}
}

var mono *g.FontInfo

func init() {
	mono = g.AddFont("mono", 14)
}
func loop() {
	var table = g.Table()
	var cols []*g.TableColumnWidget
	var r []*g.TableRowWidget

	for i := range conf.Entries {
		var row []g.Widget
		for j := range conf.Entries[i] {
			if i == 0 {
				cols = append(cols, g.TableColumn(cmd(conf.Entries[i][j]...)))
			} else {
				if len(conf.Entries[i][j]) == 1 {
					row = append(row, g.Label(conf.Entries[i][j][0]).Font(mono))
				} else {
					c := conf.Entries[i][j]
					conf.Entries[i][j] = []string{"Loading..."}
					go func(c []string, i int, j int) {
						log.Println("loading...", c)
						conf.Entries[i][j] = []string{cmd(c...)}
						g.Update()
					}(c, i, j)
					row = append(row, g.Label("Loading..."))
				}
			}
		}
		if i == 0 {
			table.Columns(cols...)
		} else {
			r = append(r, g.TableRow(row...))
		}
	}
	table.Rows(r...)
	g.SingleWindow().Layout(
		table,
	)
}

func cmd(cm ...string) string {
	c := exec.Command(cm[0])
	c.Args = cm
	c.Dir = os.Getenv("HOME")
	b, err := c.CombinedOutput()
	if err != nil {
		return err.Error() + "|" + string(b)
	}
	return string(b)
}

func main() {
	b, _ := json.MarshalIndent(conf, "", "    ")
	fmt.Println(string(b))
	wnd := g.NewMasterWindow("ServerMon", 700, 400, 0)
	wnd.Run(loop)
}
