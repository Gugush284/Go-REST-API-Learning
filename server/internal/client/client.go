package ServerClient

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

func SessionRequest() []*http.Cookie {

	message := map[string]interface{}{
		"login":    "Slava8924",
		"password": "JGjkgkggd9",
	}

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	resp, err := http.Post("http://localhost:8080/sessions", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	defer resp.Body.Close()

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	log.Println(result)
	log.Println(resp.Header.Get("X-Request-ID"))

	cookie := resp.Cookies()
	log.Println(cookie)

	return cookie
}

func CreateRequest() {

	message := map[string]interface{}{
		"login":    "Slava8924",
		"password": "JGjkgkggd9",
	}

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
		return
	}

	resp, err := http.Post("http://localhost:8080/users", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer resp.Body.Close()

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	log.Println(result)
	log.Println(resp.Header.Get("X-Request-ID"))
}

func WhoamiRequest(cookie []*http.Cookie) {
	req, err := http.NewRequest("GET", "http://localhost:8080/private/whoami", nil)
	if err != nil {
		log.Fatalf("Got error %s", err.Error())
		return
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatalf("Got error while creating cookie jar %s", err.Error())
		return
	}

	client := http.Client{
		Jar: jar,
	}

	urlObj, _ := url.Parse("http://localhost:8080/private/whoami")

	client.Jar.SetCookies(urlObj, cookie)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer resp.Body.Close()

	log.Println(string(body))
	log.Println(resp.Header.Get("X-Request-ID"))
}
