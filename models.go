package main

type Message struct {
	Content string `json:"content"`
}

type Note struct {
	Index   int    `json:"index"`
	Content string `json:"content"`
}

type Document struct {
	DocId   string
	Content string
}

type Coordinate struct {
	Lattitude     float32
	Longitude     float32
	NS_Hemisphere string
	EW_Hemisphere string
}

type DocumentInfo struct {
	DocId string
}
