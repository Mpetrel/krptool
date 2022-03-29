package krptool

func NewScene(name, thumburl, title string) *Scene {
	return &Scene{
		Name:     name,
		Thumburl: thumburl,
		Title:    title,
	}
}
