package main

import (
	"errors"
	"github.com/bootdotdev/learn-file-storage-s3-golang-starter/internal/database"
	"strings"
	"time"
)

func (cfg *apiConfig) dbVideoToSignedVideo(video database.Video) (database.Video, error) {
	if video.VideoURL == nil {
		return database.Video{}, errors.New("VideoURL is nil")
	}
	
	parts := strings.Split(*video.VideoURL, ",")

	if len(parts) != 2 {
		return database.Video{}, errors.New("invalid video url")
	}

	bucket := parts[0]
	key := parts[1]

	url, err := generatePresignedURL(cfg.s3Client, bucket, key, 5*time.Minute)

	if err != nil {
		return database.Video{}, err
	}

	video.VideoURL = &url
	return video, nil
}
