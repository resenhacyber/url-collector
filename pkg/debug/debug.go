package debug

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/ShangRui-hash/url-collector/config"
	"github.com/olekukonko/tablewriter"
)

func Println(a ...interface{}) (n int, err error) {
	if !config.CurrentConf.Debug {
		return 0, nil
	}
	return fmt.Fprintln(os.Stderr, a...)
}

func WriteFile(filename string, a ...interface{}) error {
	if !config.CurrentConf.Debug {
		return nil
	}
	return ioutil.WriteFile(filename, []byte(fmt.Sprintf("%s", a...)), 0644)
}

func ShowConfig() {
	if !config.CurrentConf.Debug {
		return
	}
	table := tablewriter.NewWriter(os.Stderr)
	table.SetHeader([]string{"SearchEngine", "BaseURL", "RoutineCount", "Keyword", "InputFile", "OutputFile", "Format"})
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.SetBorder(true)
	table.SetRowLine(true)
	table.SetAutoMergeCells(true)
	data := [][]string{
		{
			config.CurrentConf.SearchEngine,
			config.CurrentConf.GetBaseURL(),
			fmt.Sprintf("%d", config.CurrentConf.RoutineCount),
			config.CurrentConf.Keyword,
			config.CurrentConf.InputFilePath,
			config.CurrentConf.OutputFilePath,
			config.CurrentConf.Format,
		},
	}
	table.AppendBulk(data)
	table.SetCaption(true, "Current Config")
	table.Render()
	fmt.Fprintln(os.Stderr, "[*] black list:", strings.Join(config.CurrentConf.BlackList, ","))

}
