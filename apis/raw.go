package apis

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type paths struct {
	StaionName, ArrivedT, StartT, RunT string
	dayDiff                            int
	Point
}
type TrainPath struct {
	StationCode, TainsCode string
	Paths                  []paths

	From, To string
}
type Trains struct {
	msql *sql.DB
}
type Point struct {
	X float32
	Y float32
}

func (t *Trains) Close() error {
	return t.msql.Close()
}
func NewTrains() (*Trains, error) {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", os.Getenv("USERNAME"), os.Getenv("PASSWORD"), os.Getenv("NETWORK"), os.Getenv("SERVER"), os.Getenv("PORT"), os.Getenv("DATABASE"))

	msqlconn, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = msqlconn.Ping(); err != nil {
		return nil, err
	}
	return &Trains{msqlconn}, nil
}
func (t *Trains) GetTrainsFromTrainCode(m string) (*TrainPath, error) {
	train := &TrainPath{}
	if rowResult, err := t.msql.Query("SELECT a.*,x(b.lat) as x,y(b.lat) as y FROM trains AS a INNER JOIN station_info AS b ON a.name=b.name WHERE train_no=?", m); err == nil {
		defer rowResult.Close()
		var s struct {
			id, DAY_DIFF                                                      int
			TRAIN_NO, STATION_NO, ARRIVE_TIME, START_TIME, RUNNING_TIME, NAME string
			Point
		}

		for rowResult.Next() {
			if err := rowResult.Scan(&s.id, &s.TRAIN_NO, &s.STATION_NO, &s.ARRIVE_TIME, &s.START_TIME, &s.RUNNING_TIME, &s.DAY_DIFF, &s.NAME, &s.X, &s.Y); err != nil {
				log.Println(err)
			}
			train.Paths = append(train.Paths, paths{s.NAME, s.ARRIVE_TIME, s.START_TIME, s.RUNNING_TIME, s.DAY_DIFF, Point{s.X, s.Y}})
		}
		if err := rowResult.Err(); err != nil {
			return nil, err
		}
		if len(train.Paths) == 0 {
			return train, nil
		}
		train.StationCode = m
		train.TainsCode = m
		return train, nil
	} else {
		return nil, err
	}

}
func (t *Trains) GetTrainsFromStationCode(station_train_code string) (*TrainPath, error) {
	train := &TrainPath{}
	var from, to string
	if rowResult, err := t.msql.Query("SELECT b.*,a.from_station,a.to_station y FROM train_info AS a LEFT JOIN trains AS b  ON a.train_no=b.train_no  WHERE a.station_train_code=? ORDER BY b.station_no", station_train_code); err != nil {
		// data, _ := json.Marshal(map[string]interface{}{"code": 200, "data": nil, "msg": err.Error()})
		return nil, err
	} else {
		defer rowResult.Close()
		var s struct {
			id, DAY_DIFF                                                      int
			TRAIN_NO, STATION_NO, ARRIVE_TIME, START_TIME, RUNNING_TIME, NAME string
			Point
		}
		for rowResult.Next() {
			if err := rowResult.Scan(&s.id, &s.TRAIN_NO, &s.STATION_NO, &s.ARRIVE_TIME, &s.START_TIME, &s.RUNNING_TIME, &s.DAY_DIFF, &s.NAME, &from, &to); err != nil {
				log.Println(err)
			}
			train.Paths = append(train.Paths, paths{s.NAME, s.ARRIVE_TIME, s.START_TIME, s.RUNNING_TIME, s.DAY_DIFF, s.Point})
		}
	}
	if len(train.Paths) == 0 {
		return train, nil
	}
	for i, x := range train.Paths {

		resultRow := t.msql.QueryRow("SELECT x(lat) as x,y(lat) as y FROM station_info WHERE name=?", x.StaionName)
		if err := resultRow.Scan(&train.Paths[i].X, &train.Paths[i].Y); err != nil {
			log.Println(err.Error())
		}
	}

	train.StationCode = station_train_code
	// train.TainsCode =
	train.From, train.To = from, to
	return train, nil

}
func (t *Trains) GetTrainsFromAddress(from, to string) (*TrainPath, error) {

	var m string
	row := t.msql.QueryRow("SELECT a.train_no FROM trains AS a INNER JOIN trains AS b ON a.train_no=b.train_no WHERE a.station_no>b.station_no AND a.name=? AND b.name=? ", from, to)
	log.Println(row.Scan(&m))
	if m == "" {
		// data, _ := json.Marshal(map[string]interface{}{"Code": 200, "Data": nil, "Msg": "No Result"})
		return nil, fmt.Errorf("Resutl is null")
	}
	train, _ := t.GetTrainsFromTrainCode(m)
	train.From, train.To = from, to
	return train, nil
}

func Trains2AmapPathSimplifier(paths *TrainPath) []byte {
	data := make(map[string]interface{})
	data["train_code"] = paths.StationCode
	m := make([][]float32, len(paths.Paths))
	var stationInfo []interface{}
	for i, p := range paths.Paths {
		m[i] = []float32{p.X, p.Y}
		stationInfo = append(stationInfo, map[string]interface{}{"StationName": p.StaionName, "AT": p.ArrivedT, "ST": p.StartT, "RT": p.RunT, "DayDiff": p.dayDiff})
	}
	data["paths"] = m
	data["stationInnfo"] = stationInfo
	nd, _ := json.Marshal(data)
	return nd

}
