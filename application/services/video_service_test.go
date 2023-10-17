package services_test

import (
	"converter/application/repositories"
	"converter/application/services"
	"converter/domain"
	"converter/infrastructure/database"
	"log"
	"testing"
	"time"

	"github.com/joho/godotenv"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Erro loading .env file, %v", err)
	}
}

func prepare() (*domain.Video, repositories.VideoRepository) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "convite.mp4"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}

	return video, &repo
}

func TestVideoServiceDownload(t *testing.T) {

	video, repo := prepare()

	videoService := services.NewVideoService()
	videoService.Video = video
	videoService.VideoRepository = &repo

	err := videoService.Download("convertr")

	require.Nil(t, err)
	err = videoService.Fragment()
	if err != nil {
		log.Printf("Error %v", err)
	}
	require.Nil(t, err)

	// err = videoService.Encode()
	// require.Nil(t, err)

	// err = videoService.Finish()
	// require.Nil(t, err)
}
