package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	data := map[string]any{
		"method": "HelloService.Hi",
		"params": []string{
			"张三",
		},
		"id": 0,
	}
	buf, err := json.Marshal(data)
	if err != nil {
		log.Fatal("marshal failed: ", err)
	}

	res, err := http.Post("http://127.0.0.1:11223", "application/json;charset=utf-8", bytes.NewReader(buf))
	if err != nil {
		log.Fatal("post failed: ", err)
	}
	defer res.Body.Close()

	buf, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("read all failed: ", err)
	}
	log.Printf("%s\n", buf)
}
