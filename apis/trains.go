package apis

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var msql *sql.DB

type paths struct {
	StaionName             string
	ArrivedT, StartT, RunT string
	dayDiff                uint
}
type TrainPath struct {
	StationCode, TainsCode string
	Paths                  []paths

	From, To string
}

func init() {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", os.Getenv("USERNAME"), os.Getenv("PASSWORD"), os.Getenv("NETWORK"), os.Getenv("SERVER"), os.Getenv("PORT"), os.Getenv("DATABASE"))

	msqlconn, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	msql = msqlconn
	if err = msql.Ping(); err != nil {
		panic(err)
	}
}

func TrainInfo(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.Write([]byte("param error"))
		return
	}

}
func GetTrains(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.Write([]byte("param error"))
		return
	}
	train := &TrainPath{}
	var m string
	row := msql.QueryRow("SELECT a.train_no FROM trains AS a INNER JOIN trains AS b ON a.train_no=b.train_no WHERE a.station_no>b.station_no AND a.name=? AND b.name=? ", r.Form.Get("from"), r.FormValue("to"))
	log.Println(row.Scan(&m))
	if m == "" {
		data, _ := json.Marshal(map[string]interface{}{"Code": 200, "Data": nil, "Msg": "No Result"})
		w.Write(data)
		return
	}
	if rowResult, err := msql.Query("SELECT * FROM trains WHERE train_no=?", m); err == nil {
		var s struct {
			id                                                                int
			TRAIN_NO, STATION_NO, ARRIVE_TIME, START_TIME, RUNNING_TIME, NAME string
			DAY_DIFF                                                          uint
		}
		for rowResult.Next() {
			if err := rowResult.Scan(&s.id, &s.TRAIN_NO, &s.STATION_NO, &s.ARRIVE_TIME, &s.START_TIME, &s.RUNNING_TIME, &s.DAY_DIFF, &s.NAME); err != nil {
				log.Println(err)
			}
			train.Paths = append(train.Paths, paths{s.NAME, s.ARRIVE_TIME, s.START_TIME, s.RUNNING_TIME, s.DAY_DIFF})
		}
	}
	train.StationCode = m
	train.From, train.To = r.FormValue("from"), r.FormValue("to")
	train.TainsCode = m
	data, _ := json.Marshal(&train)
	w.Write(data)

}
