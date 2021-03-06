package main

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
	db.Register("Cangjie3", CangjieDataHandler{})
	db.Register("Cangjie5", CangjieDataHandler{})

	// parse these cj files
	parseCangjie3File(dir+"/cj3/FTCJ_UTF8.TXT", "Cangjie3", 8, "zh-hant", "FTCJ_B5.TXT", &db)
	parseCangjie3File(dir+"/cj3/JTCJ_UTF8.TXT", "Cangjie3", 5, "zh-hans", "JTCJ_GB.TXT", &db)
	parseCangjie5File(dir+"/cj5/cj5-21000", "Cangjie5", "zh-hant", "cj5-21000", &db)
	parseCangjie5File(dir+"/cj5/cj5-jt-7000", "Cangjie5", "zh-hans", "cj5-jt-7000", &db)
	parseLibcangjie1File(dir+"/cj3/libcangjie1-cj3-cc.txt", "Cangjie3", 7, "zh-hant", "libcangjie1-cj3-cc.txt", &db)
	parseLibcangjie1File(dir+"/cj3/libcangjie1-cj3-cjk.txt", "Cangjie3", 7, "zh-hant", "libcangjie1-cj3-cjk.txt", &db)
	parseLibcangjie1File(dir+"/cj3/libcangjie1-cj3-sc.txt", "Cangjie3", 6, "zh-hant", "libcangjie1-cj3-sc.txt", &db)
	parseLibcangjie1File(dir+"/cj3/libcangjie1-cj3-tc.txt", "Cangjie3", 7, "zh-hant", "libcangjie1-cj3-tc.txt", &db)
	parseLibcangjie1File(dir+"/cj5/libcangjie1-cj5-cc.txt", "Cangjie5", 7, "zh-hant", "libcangjie1-cj5-cc.txt", &db)
	parseLibcangjie1File(dir+"/cj5/libcangjie1-cj5-cjk.txt", "Cangjie5", 7, "zh-hant", "libcangjie1-cj5-cjk.txt", &db)
	parseLibcangjie1File(dir+"/cj5/libcangjie1-cj5-sc.txt", "Cangjie5", 6, "zh-hant", "libcangjie1-cj5-sc.txt", &db)
	parseLibcangjie1File(dir+"/cj5/libcangjie1-cj5-tc.txt", "Cangjie5", 7, "zh-hant", "libcangjie1-cj5-tc.txt", &db)
}
