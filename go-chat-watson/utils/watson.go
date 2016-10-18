package utils

type Watson_images struct {
	Classifiers classifiers `json:"classifiers"`
}

type classifiers struct {
	Classes []class `json:"classes"`
}

type class struct {
	Class string  `json:"class"`
	Score float32 `json:"score"`
}
