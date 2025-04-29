package main

type FFProbeOutput struct {
	Streams []StreamInfo `json:"streams"`
}

type StreamInfo struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}
