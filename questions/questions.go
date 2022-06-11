package questions

import (
	"log"

	"github.com/henriquetied472/tabshell-cli/api"
	"github.com/AlecAivazis/survey/v2"
)

func Init(debugger, logger *log.Logger, dgb *bool) {
	tc := api.Init(debugger, logger, dgb)
	allTitles := getAllTitles(*tc)
	var sel string

	qt := &survey.Select{
		Message: "Bem vindo ao TabShell, qual postagem deseja ver?",
		Options: allTitles,
		Default: allTitles[len(allTitles) - 1],
	}

	survey.AskOne(qt, &sel)

	if *dgb { debugger.Printf(sel) }
	logger.Printf("Selected: %v", sel)
}

func getBodyFromID(tc api.TabContents, id string) string {
	for _, tr := range tc.Resp {
		if tr.ID == id {
			return tr.Body
		}
	}

	return ""
}

func getAllTitles(tc api.TabContents) []string {
	var r []string
	for _, tr := range tc.Resp {
		r = append(r, tr.Title)
	}

	return r
}

func getAllIDs(tc api.TabContents) []string {
	var r []string
	for _, tr := range tc.Resp {
		r = append(r, tr.ID)
	}

	return r
}

func getIDFromTitle(tc api.TabContents, title string) string {
	for _, tr := range tc.Resp {
		if tr.Title == title {
			return tr.ID
		}
	}

	return ""
}
