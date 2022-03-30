package krptool

import "encoding/xml"

func NewKrpano() *Krpano {
	return &Krpano{}
}

func (krp *Krpano) AddScene() {
}

func (krp *Krpano) ToXML() (string, error) {
	data, err := xml.Marshal(krp)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (krp *Krpano) AddPlugin(pluginStr string) error {
	var plugin Plugin
	err := xml.Unmarshal([]byte(pluginStr), &plugin)
	if err != nil {
		return err
	}
	krp.Plugin = append(krp.Plugin, &plugin)
	return nil
}
