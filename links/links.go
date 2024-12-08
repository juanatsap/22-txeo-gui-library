package links

import (
	"github.com/pkg/browser"
	log "github.com/sirupsen/logrus"
)

func OpenInBrowser(url string) {
	if err := browser.OpenURL(url); err != nil {
		log.Println(err)
	}
}
func OpenLaCaixa() {
	OpenInBrowser("https://www.caixabank.es/particular/home/particulares_es.html")
	return
}
func OpenExcel() {
	OpenInBrowser("https://obs-my.sharepoint.com/:x:/r/personal/juan_morenofl_olympicchannel_com/_layouts/15/doc2.aspx?sourcedoc=%7B074207B2-5AA2-4F4D-A0C1-7B6BB4DD2D5F%7D&file=Plantilla%20Juan%20Noviembre.xlsx&action=default&mobileredirect=true&ct=1731867510807&wdOrigin=OFFICECOM-WEB.START.UPLOAD&cid=efb2269b-7852-426a-9968-1236fe6e7242&wdPreviousSessionSrc=HarmonyWeb&wdPreviousSession=cfafc7a0-bcb6-449c-9c3c-bf7c6d967fed")
}
