package ModelImage

// User struct ...
type Image struct {
	ImageId   int    `json:"id"`
	ImageType string `json:"type"`
	Image     string `json:"image"`
	ImageName string `json:"name"`
	Txt       string `json:"txt"`
}
