package site

import (
	"log"
	"net/http"

	"github.com/VoltealProductions/TheArchiveReader/utils"
	"github.com/VoltealProductions/TheArchiveReader/views/pages"
)

func GetHomePage(w http.ResponseWriter, r *http.Request) {
	if err := utils.RenderView(w, r, pages.Index()); err != nil {
		log.Fatal(err)
	}
}
