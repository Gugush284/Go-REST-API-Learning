package ModelImage

// User struct ...
type Image struct {
	ImageId   int    `json:"id"`
	Image     string `json:"imagepath"`
	ImageName string `json:"name"`
	Txt       string `json:"txt"`
}
