package app

import (
	"container/types"
	"log"
	"net/http"

	"github.com/BurntSushi/toml"
	"github.com/gorilla/mux"
)

func Run(env string) {
	var c types.Config

	config := "local.toml"
	if env != "" {
		config = env + ".toml"
	}

	if _, err := toml.DecodeFile(config, &c); err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.PathPrefix("/").HandlerFunc(HandleRequest)

	err := http.ListenAndServe(c.Port, r)
	log.Fatal(err)

}
