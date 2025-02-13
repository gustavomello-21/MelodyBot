package entities

// Video represents a video
type Video struct {
	Title     string  `json:"title"`
	Url       string  `json:"url"`
	Duration  float64 `json:"duration"`
	Thumbnail string  `json:"thumbnail"`
}
