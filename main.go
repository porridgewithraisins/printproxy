package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			http.ServeFile(w, r, "index.html")
			return
		}

		file, header, err := r.FormFile("file")
		if err != nil {
			w.Header().Set("Content-Type", "text/html")
			w.Write([]byte("Unable to get file from form"))
			return
		}
		defer file.Close()

		lpOptions := []string{}

		pages := r.FormValue("pages")
		pages = strings.ReplaceAll(pages, " ", "")
		if pages != "all" {
			lpOptions = append(lpOptions, "-P", pages)
		}

		orientation := r.FormValue("orientation")
		if orientation == "landscape" {
			lpOptions = append(lpOptions, "-o", "orientation-requested=4")
		}

		bothSides := r.FormValue("both-sides")
		lpOptions = append(lpOptions, "-o", fmt.Sprintf("sides=%s", bothSides))

		pagesPerSheet := r.FormValue("pages-per-sheet")
		if pagesPerSheet != "1" {
			lpOptions = append(lpOptions, "-o", fmt.Sprintf("number-up=%s", pagesPerSheet))
		}

		lpOptions = append(lpOptions, "-")

		lp := exec.Command("lp", lpOptions...)
		lp.Stdin = file
		res, err := lp.CombinedOutput()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "text/html")
		w.Write(res)

		status := "success"
		if err != nil {
			status = "failure"
		}
		fmt.Printf("%s\t%s\t%s\t%s\t%s\n", time.Now().Format(time.RFC3339), r.RemoteAddr, header.Filename, lp, status)
	})

	port, present := os.LookupEnv("PORT")
	if !present {
		port = "3000"
	}
	addr := fmt.Sprintf(":%s", port)

	fmt.Printf("Listening on address %s\n", addr)
	http.ListenAndServe(addr, nil)
}
