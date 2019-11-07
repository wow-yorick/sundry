package other

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func echo(wr http.ResponseWriter, r *http.Request) {
	msg, err := ioutil.ReadAll(r.Body)
	if err != nil {
		wr.Write([]byte("echo error"))
		return
	}

	writeLen, err := wr.Write(msg)
	if err != nil || writeLen != len(msg) {
		log.Println(err, "write len:", writeLen)
	}

}

func tserver() {
	// http.HandleFunc("/", echo)
	// err := http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	r := httprouter.New()
	r.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("oh no, not found"))
	})

	r.PanicHandler = func(w http.ResponseWriter, r *http.Request, c interface{}) {
		log.Printf("Recovering from panic, Reason: %#v", c.(error))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(c.(error).Error()))
	}
}
