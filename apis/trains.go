package apis

import (
	rpc1 "api/rpc"
	"context"
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
	code := r.FormValue("stcode")
	if len(code) == 0 {
		data, _ := json.Marshal(map[string]interface{}{"Code": -1, "Msg": "use stcode."})
		w.Write(data)
		return
	}
	exists, err := rpc1.TrainsClient.Exits(context.Background(), &rpc1.TrainCode{Code: code})
	if err != nil {
		log.Println("some rpc proc has error,", err.Error())
	} else {
		if !exists.GetExists() {
			data, _ := json.Marshal(map[string]interface{}{"Code": -1, "Msg": "Not this train " + code})
			w.Write(data)
			return
		}
	}
	if p, err := tains.GetTrainsFromStationCode(code); err != nil {
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if ep, err := tains.GetTrainsFromAddress(from, to); err != nil {
		w.Write([]byte(err.Error()))
	} else {
		data, _ := json.Marshal(ep)
		w.Write(data)
	}
}
