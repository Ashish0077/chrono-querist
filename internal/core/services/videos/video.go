package videos

import "chrono-querist/internal/ports/secondary/repository"

type Service struct {
	videosRepo repository.VideosPort
}

func NewService(videosRepo repository.VideosPort) *Service {
	return &Service{
		videosRepo: videosRepo,
	}
}
