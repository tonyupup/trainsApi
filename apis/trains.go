package apis

import (
	"encoding/json"
	"log"
	"net/http"
)

var tains *Trains

func init() {
	if T, err := NewTrains(); err != nil {
		log.Fatal(err)
	} else {
		tains = T
	}

}
func Shutdown() {
	log.Fatalln(tains.Close())
}
func GetPathFromStationCode(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.Write([]byte("param error"))
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if p, err := tains.GetTrainsFromStationCode(r.FormValue("stcode")); err != nil {
		data, _ := json.Marshal(map[string]interface{}{"Code": -1, "Msg": err.Error()})
		w.Write(data)
	} else {
		w.Write(Trains2AmapPathSimplifier(p))
	}

}
func GetTrains(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.Write([]byte("param error"))
		return
	}
	from, to := r.FormValue("from"), r.FormValue("to")
	if from == "" || to == "" {
		w.Write([]byte("bad argument"))
	}
	if ep, err := tains.GetTrainsFromAddress(from, to); err != nil {

		w.Write([]byte(err.Error()))
	} else {
		data, _ := json.Marshal(ep)
		w.Write(data)
	}
}
