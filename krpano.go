package krptool

import "encoding/xml"

func NewKrpano() *Krpano {
	return &Krpano{}
}

func (krp *Krpano) AddScene() {

}

func (krp *Krpano) Dump() (string, error) {
	data, err := xml.Marshal(krp)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
