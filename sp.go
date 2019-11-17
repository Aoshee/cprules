package main

import (
	"fmt"
	"time"

	toml "github.com/pelletier/go-toml"
)

func main() {
	config, _ := toml.LoadFile("./cprules.toml")
	rulesConfig := config.Get("rules_path").(*toml.Tree)
	rule_url := rulesConfig.Get("rule_url").(string)
	rule_exists := rulesConfig.Get("rule_exists").(string)
	fmt.Println(rule_url)
	current_time := time.Now().Local()
	fmt.Println("The Current time is ", current_time.Format("2006-01-02"))

	current_time = time.Now().UTC()
	fmt.Println("The Current time is ", current_time.Format("2006-01-02 MST"))
	//	err := CopyFile("/home/aishee/Documents/vpn.zip", "/tmp/vpn-"+current_time.Format("2006-01-02")+".zip")
	//	if err != nil {
	//		fmt.Println("F")
	//	} else {
	//		fmt.Println("X")
	//	}
	//	f, err := os.Open("/home/aishee/Documents/testtules/etpro.rules.tar.gz")
	//	if err != nil {
	//		panic(err)
	//	}
	//	defer f.Close()
	//ExtractCompress(f)
}
