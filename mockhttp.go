package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	yaml "gopkg.in/yaml.v2"
)

type (
	// Reply stores reply for entrypoint
	Reply struct {
		Status int
		Body   string
		Delay  time.Duration
	}
)

func main() {
	addr := flag.String("url", ":8080", "URL to listen to")
	config := flag.String("config", "./mockhttp.yml", "Config file")
	flag.Parse()
	content, e := ioutil.ReadFile(*config)
	if e != nil {
		log.Fatal(e)
	}
	ymlCfg := make(map[string]Reply, 0)
	e = yaml.Unmarshal(content, &ymlCfg)
	if e != nil {
		log.Fatal(e)
	}
	for k, route := range ymlCfg {
		log.Println(k, route)
		http.HandleFunc(k, func(w http.ResponseWriter, r *http.Request) {

			v, ok := ymlCfg[r.URL.Path]
			if !ok {
				log.Println("No route found", r.URL.Path)
				w.WriteHeader(404)
				return
			}
			log.Println("Requested URL:", r.URL.String(), v)
			if v.Delay > 0 {
				<-time.After(v.Delay)
			}
			if v.Status != 0 {
				w.WriteHeader(v.Status)
			}
			if v.Body != "" {
				w.Write([]byte(v.Body))
			}
		})
	}
	http.ListenAndServe(*addr, nil)
}
