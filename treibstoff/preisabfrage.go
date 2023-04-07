package treibstoff

import (
	"log"
	"net/http"
)

func PreisabfrageStarten() {
	preisabfrage()
}

//API Key 01c92cee-9fdf-18e1-cdf9-1351e969fd73
/*	Hessol Bad Vilbel
	lat: 50.19396
	lng: 8.72511
*/
type parameterUmkreissuche struct {
	lat           string //Längengrad
	lng           string //Breitengrad
	rad           string //Radius
	treibstoffart string //e5 e10 diesel all
	sort          string //price dist
}

func preisabfrage() {
	response, err := http.Get("https://creativecommons.tankerkoenig.de/json/list.php?")
	if err != nil {
		log.Println("[{preisabfrage} Error HTTP Get]: ", err)
	}
}
