package main

import (
	"net/http"
	"time"
	"github.com/wcharczuk/go-chart"
	"github.com/araddon/dateparse"
	"encoding/json"
	"fmt"
)

type Payload struct {
	AnomalyValue    []float64    `json:"anomalyValue"`
	TimeStampValues []string `json:"timeStampValues"`
	Title           string   `json:"title"`
	XLabel          string   `json:"xLabel"`
	YLabel          string   `json:"yLabel"`
}


func drawChart(res http.ResponseWriter, req *http.Request) {

	
	defer req.Body.Close()	
	decoder:=json.NewDecoder(req.Body)
        var t Payload;
  	err:=decoder.Decode(&t)
  	if err==nil {
		/* throw an error */
	}
	datapoints:=t.AnomalyValue	
	var timepoints[] time.Time	
	for timepoint:= range t.TimeStampValues {
		
		ts,err:=dateparse.ParseAny(t.TimeStampValues[timepoint])
		if err!=nil {
			fmt.Println("Error Occured",err)
		}
		timepoints=append(timepoints,ts)
	}
	graph := chart.Chart{
		XAxis: chart.XAxis{
			Name:      "The XAxis",
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
		},

		YAxis: chart.YAxis{
			Name:      "The YAxis",
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
		},
		Series: []chart.Series{
			chart.TimeSeries{
				XValues: timepoints,
				YValues: datapoints,
			},
		},
	}

	res.Header().Set("Content-Type", "image/svg+xml")
	graph.Render(chart.SVG, res)
}


func main() {
	http.HandleFunc("/", drawChart)
	http.HandleFunc("/favico.ico", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte{})
	})
	http.ListenAndServe(":8080", nil)
}
