package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type TabResponse struct {
	ID                string      `json:"id"`
	OwnerID           string      `json:"owner_id"`
	ParentID          interface{} `json:"parent_id"`
	Slug              string      `json:"slug"`
	Title             string      `json:"title"`
	Body              string      `json:"body"`
	Status            string      `json:"status"`
	SourceURL         interface{} `json:"source_url"`
	CreatedAt         time.Time   `json:"created_at"`
	UpdatedAt         time.Time   `json:"updated_at"`
	PublishedAt       time.Time   `json:"published_at"`
	Username          string      `json:"username"`
	ParentTitle       interface{} `json:"parent_title"`
	ParentSlug        interface{} `json:"parent_slug"`
	ParentUsername    interface{} `json:"parent_username"`
	ChildrenDeepCount int         `json:"children_deep_count"`
}

type TabContents struct {
	Resp []TabResponse
}

func readTCs(debugger *log.Logger, content []byte, dgb *bool) (*TabContents, string) {
	var trs []TabResponse
	var tcIdented bytes.Buffer

	json.Unmarshal(content, &trs)
	json.Indent(&tcIdented, content, "", "\t")

	if *dgb {
		debugger.Printf(tcIdented.String())
	}

	tc := TabContents{trs}

	return &tc, tcIdented.String()
}

var FPath string

func Init(debugger, logger *log.Logger, dgb *bool) *TabContents {
	res, _ := http.Get("https://tabnews.com.br/api/v1/contents")
	defer res.Body.Close()
	j, _ := ioutil.ReadAll(res.Body)
	tc, idented := readTCs(debugger, j, dgb)

	FPath = fmt.Sprint(time.Now().UnixMicro()) + ".json"

	f, _ := os.Create("./data", FPath)
	f.WriteString(idented)

	return tc
}