package webcontrollers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"resttest/db"
)

func HandleHealthCheck(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Serivce is healthy")
}

func HandleGetItems(res http.ResponseWriter, req *http.Request) {
	buffer, err := json.Marshal(db.Data)
	if err != nil {
		fmt.Printf("error")
	}
	res.Write(buffer)
}
