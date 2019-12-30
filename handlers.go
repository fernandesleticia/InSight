package main

import (
    "net/http"
    "encoding/json"
    "time"
    "strings"
    "strconv" 
    "fmt"  
)

// Recebe, aplica mudanças e retorna nova data
func ConvertDate(w http.ResponseWriter, r *http.Request) {
    
    w.Header().Set("Content-Type","application/json")

    decoder := json.NewDecoder(r.Body)
    var t DataToReceive
    err := decoder.Decode(&t)
    if err != nil {
        panic(err)
    }
    r.Body.Close()

    dateReceived := t.Today


    // getting values of date

    dateReceivedPieces := strings.Split(dateReceived, "-")
    yearString, monthString, dayString := dateReceivedPieces[0], dateReceivedPieces[1], dateReceivedPieces[2]

    year, err := strconv.Atoi(yearString)
    month, err := strconv.Atoi(monthString)
    day, err := strconv.Atoi(dayString)
    hour := 0
    minute := 0
    
    // getting location
    location, err := time.LoadLocation("America/Fortaleza")
    if err != nil {
        fmt.Printf("ERROR : %s", err)
    }

    // mounting date
    today := time.Date(year, time.Month(month), day, hour, minute, 0, 0, location)
    tomorrow := today.AddDate(0, 0, 1)

    data := DataToSend{Tomorrow: tomorrow.Format(DATE_FORMAT_DELFOS)}

    json.NewEncoder(w).Encode(data)
}