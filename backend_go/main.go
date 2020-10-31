package main

import(
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"
	"crypto/sha256"
	"encoding/hex"
)

type Out struct {
	Name    string
	Result	string
  }

func writehandling(w http.ResponseWriter, r *http.Request)  {
	//TODO check wether it is a get or not
	
}

func shahandling(w http.ResponseWriter, r *http.Request)  {
	//TODO check wether it is a post or not
	reqBody, err := ioutil.ReadAll(r.Body)
    if err != nil {
        log.Fatal(err)
    }
	fmt.Printf("%s\n", reqBody)
	var dat map[string]interface{}
	if err := json.Unmarshal(reqBody, &dat); err != nil {
        panic(err)
	}
	var t1 = dat["1"].(string)
	var t2 = dat["2"].(string)
	n1, err := strconv.ParseInt(t1, 10, 64)
	n2, err := strconv.ParseInt(t2, 10, 64)
	n1 = n1 + n2
	fmt.Println(n1)

	hash := sha256.New()
	s := strconv.FormatInt(n1, 10)
	hash.Write([]byte(s))
	md := hash.Sum(nil)
	mdStr := hex.EncodeToString(md)
	res := Out{"Ali", mdStr}

  	js, err := json.Marshal(res)
  	if err != nil {
    	http.Error(w, err.Error(), http.StatusInternalServerError)
    	return
  	}
  	w.Header().Set("Content-Type", "application/json")
  	w.Write(js)
}

func main()  {
	fmt.Println("Starting a web server...")
	http.HandleFunc("/go/sha256", shahandling)
	http.HandleFunc("/go/sha256", writehandling)
	http.ListenAndServe(":3000", nil)
}