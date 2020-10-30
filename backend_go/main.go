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

func helloworld(w http.ResponseWriter, r *http.Request)  {
	if r.URL.Path != "/go/sha256" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
	}
	reqBody, err := ioutil.ReadAll(r.Body)
    if err != nil {
        log.Fatal(err)
    }
	fmt.Printf("%s\n", reqBody)
	//w.Write([]byte("Received a POST request\n"))
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


	// sum := sha256.Sum256([]byte())
	// fmt.Printf("%x", sum)
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
	http.HandleFunc("/go/sha256", helloworld)
	http.ListenAndServe(":3000", nil)
}