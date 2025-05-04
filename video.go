package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"os/exec"
)

func getVideoAspectRatio(filePath string) (string, error) {
	cmd := exec.Command("ffprobe", "-v", "error", "-print_format", "json", "-show_streams", filePath)
	buf := &bytes.Buffer{}
	cmd.Stdout = buf

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	data := buf.Bytes()
	ffProbe := FFProbeOutput{}

	err = json.Unmarshal(data, &ffProbe)
	if err != nil {
		return "", err
	}

	if len(ffProbe.Streams) == 0 {
		return "", errors.New("no streams found")
	}

	height := ffProbe.Streams[0].Height
	width := ffProbe.Streams[0].Width

	if width == 0 || height == 0 {
		return "", errors.New("invalid video dimensions: width or height is zero")
	}

	const tolerance = 10

	if approxEqual(width*9, height*16, tolerance) {
		return "16:9", nil
	}

	if approxEqual(width*16, height*9, tolerance) {
		return "9:16", nil
	}

	return "other", nil
}

func approxEqual(a, b, tolerance int) bool {
	diff := a - b
	if diff < 0 {
		diff = -diff
	}
	return diff <= tolerance
}

func processVideoForFastStart(filePath string) (string, error) {
	outputFilePath := filePath + ".processing"
	cmd := exec.Command("ffmpeg", "-i", filePath, "-c", "copy", "-movflags", "faststart", "-f", "mp4", outputFilePath)

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	return outputFilePath, nil
}
