package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type IrDataFile struct {
	irDataMap map[string]*IrData
}

func NewDataFile(filename string) (*IrDataFile, error) {

	txt, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var jsonData jsonIrDataMap
	err = json.Unmarshal(txt, &jsonData)
	if err != nil {
		return nil, err
	}

	irDataMap := map[string]*IrData{}
	for key, val := range jsonData.Ac {
		irDataMap[key] = &IrData{
			Pattern:  val.Pattern,
			Interval: val.Interval,
		}
	}

	return &IrDataFile{irDataMap}, nil
}

func (i *IrDataFile) Get(name string) (*IrData, error) {
	res, ok := i.irDataMap[name]
	if !ok {
		return nil, fmt.Errorf("Error: Command name not found: %s", name)
	}
	return res, nil
}

type jsonIrDataMap struct {
	Ac map[string]struct {
		Pattern  string `json:"pattern"`
		Interval int    `json:"interval"`
	} `json:"ac"`
}
