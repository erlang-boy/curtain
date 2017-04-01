package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/throttle"
	"io/ioutil"
	"log"
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println(Logo)
	m := martini.Classic()
	// A Rate Limit Policy
	m.Use(throttle.Policy(&throttle.Quota{
		Limit:  1000,
		Within: time.Hour,
	}))

	// An Interval Policy
	m.Use(throttle.Policy(&throttle.Quota{
		Limit:  1,
		Within: time.Second,
	}))
	m.Action(router().Handle)
	martini.Env = martini.Dev
	m.Run()
}

func router() martini.Router {
	r := martini.NewRouter()
	r.Group("/v1/actions/public", func(r martini.Router) {
		r.Put("/:name", Public)
	})
	return r
}

func Public(params martini.Params, req *http.Request, log *log.Logger) (int, string) {
	body, err := ioutil.ReadAll(req.Body)
	log.Printf("read req:%+v\n", params)
	if err != nil {
		log.Printf("read body error:%s\n", err)
		return http.StatusUnprocessableEntity, err.Error()
	}
	defer req.Body.Close()
	filename, ok := params["name"]
	log.Printf("filename:%+v, is:%+v\n", filename, ok)
	if ok {
		err = ioutil.WriteFile(filename, body, 0644)
		if err != nil {
			log.Printf("write file error:%s\n", err)
			return http.StatusInternalServerError, err.Error()
		}
		return http.StatusOK, `ok`
	} else {
		return http.StatusBadRequest, "BadRequest"
	}
}
