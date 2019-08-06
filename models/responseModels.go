package models

type Result struct {
	Result bool `json:"result" xml:"result"`
}

type ResultWithNote struct {
	Result bool   `json:"result" xml:"result"`
	Note   string `json:"note" xml:"note"`
}

type Ext struct {
	Result bool   `json:"result" xml:"result"`
	Ext    string `json:"ext" xml:"ext"`
}

type Token struct {
	Result bool   `json:"result" xml:"result"`
	Token  string `json:"token" xml:"token"`
}
