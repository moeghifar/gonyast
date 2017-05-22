package sklh

import (
	"fmt"
	"log"
	"net/http"

	"encoding/json"

	"time"

	"database/sql"

	"github.com/julienschmidt/httprouter"
)

// GetSklh ...
func GetSklh(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// begin time, marking performance
	timeStart := time.Now()

	// define tempResult as DataSklh type struct
	var tempResult DataSklh
	// define ds as array of DataSklh type struct
	var ds []DataSklh

	// Call query for `StmtGetSklh` with Queryx
	rows, err := stmtGetSklh.Queryx()

	// check if err exist and error not empty row result
	if err != nil && err != sql.ErrNoRows {
		//  return err
		log.Println("[ERROR] Failed querying StmtGetSklh ->", err)
	}

	// close sql connection
	defer rows.Close()

	// loop query result
	for rows.Next() {
		// loop rows and perform structscan to tempResult
		rows.StructScan(&tempResult)
		// appending result from `tempResult` to `ds` as array type
		ds = append(ds, tempResult)
	}

	// finish time
	timeFinish := fmt.Sprintf("%f", time.Since(timeStart).Seconds())

	// logging execution
	log.Printf("Done execution in %s", timeFinish)

	// generate response time
	rt := ResponseTime{
		Duration: timeFinish,
		Unit:     "sec",
	}

	// build response with OutputResponse struct type
	or := OutputResponse{
		Data:         ds,
		ResponseTime: rt,
	}

	// marshal or struct to json format
	orJSON, _ := json.Marshal(or)

	// build response output
	BuildResponse(w, orJSON)
}

// BuildResponse ...
func BuildResponse(w http.ResponseWriter, or []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", or)
}
