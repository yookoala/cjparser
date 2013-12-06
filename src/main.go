package cjparser

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	var dir, dbfn string

	flag.StringVar(&dir, "f", "./data", "Folder containing Cangjie mapping definition files")
	flag.StringVar(&dbfn, "d", "./data/cangjie.db", "Output database filename")
	flag.Parse()

	// check if Cangjie mapping definition folder is a folder
	finfo, err := os.Stat(dir)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	} else if !finfo.IsDir() {
		fmt.Println("The directory \"" + dir + "\" should be a directory")
		os.Exit(1)
	}

	// cangjie definition database object
	db := ConvertDB{
		Filename: dbfn,
	}

	// register handlers
	db.Register("Cangjie5", Cangjie5DataHandler{
		Type: "cj5-hant",
	})

	// parse these cj files
	parseIniFile(dir+"/cj5/cj5-21000", "Cangjie5", &db)
}
