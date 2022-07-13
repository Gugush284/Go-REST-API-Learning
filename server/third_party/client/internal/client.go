package ServerClient

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"

	ModelImage "github.com/Gugush284/Go-server.git/internal/apiserver/model/image"
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

func Upload(cookie []*http.Cookie) int {
	var b bytes.Buffer
	var fw io.Writer

	w := multipart.NewWriter(&b)

	file, err := os.Open("./third_party/client/assets/testimage.png")
	if err != nil {
		log.Fatal(err)
	}

	if fw, err = w.CreateFormFile("image", file.Name()); err != nil {
		log.Fatal(err)
	}

	if _, err = io.Copy(fw, file); err != nil {
		log.Fatal(err)
	}

	w.Close()

	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatalf("Got error while creating cookie jar %s", err.Error())
		return 0
	}

	client := http.Client{
		Jar: jar,
	}

	urlObj, _ := url.Parse("http://localhost:8080/private/upload/image")

	client.Jar.SetCookies(urlObj, cookie)

	req, err := http.NewRequest("POST", "http://localhost:8080/private/upload/image", &b)
	if err != nil {
		log.Fatalf("Got error %s", err.Error())
		return 0
	}
	req.Header.Set("Content-Type", w.FormDataContentType())

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	log.Println(resp.Header.Get("X-Request-ID"))

	result := &ModelImage.Image{}

	json.Unmarshal(body, result)
	log.Println(result)
	return result.ImageId
}
