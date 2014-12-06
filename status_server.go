package main

import(
    "net/http"
    "fmt"
    "time"
    "strings"
)

var St []string // Array storing Site Status
var checkInterval time.Duration = time.Second * 2 // Time to check TODO: Change this value
var historyCount int = 10

func main(){
    // Statistic Site Content Server, for Javascript Timer Code
    http.Handle("/", http.FileServer(http.Dir("./")))

    // Server latest queried Site status
    http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, strings.Join(St, "<br/>"))
    })

    // Ticker and go routine handling the Tick channel call
    ticker := time.NewTicker(checkInterval)
    go func() {
        for t := range ticker.C {
            fmt.Printf("Checking site status at %v.\n", t)
            resp, err := http.Head("http://localhost:4567/")
            if err != nil {
                if len(St) == historyCount {
                   St = append(St[:0],St[1:]...) // Remove first one
                }
                St = append(St, fmt.Sprintf("%d : %s", 0, time.Now().Format(time.Stamp)))
                //TODO: Do something when the site is down.
            }else{
                if len(St) == historyCount {
                   St = append(St[:0],St[1:]...) // Remove fist one
                }
                St = append(St, fmt.Sprintf("%d : %s", resp.StatusCode, time.Now().Format(time.Stamp)))
            }
        }
    }()

    fmt.Printf("Sitecheker started with %v interval.\n", checkInterval)
    fmt.Printf("Server starting at :8080")
    http.ListenAndServe(":8080", nil)
    ticker.Stop()
}
