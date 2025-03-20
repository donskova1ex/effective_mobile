package processors

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/donskova1ex/effective_mobile/internal/domain"
)

type SongRepository interface {
	CreateSong(ctx context.Context, song *domain.Song) (*domain.Song, error)
	GetSong(ctx context.Context, groupName, songName string) (*domain.Song, error)
	UpdateSong(ctx context.Context, song *domain.Song) (*domain.Song, error)
	DeleteSong(ctx context.Context, groupName, songName string) error
}
type SongsLogger interface {
	Info(ctx context.Context, msg string, args ...any)
	Error(ctx context.Context, msg string, args ...any)
}

type song struct {
	repo   SongRepository
	logger SongsLogger
}

func NewSong(repo SongRepository, logger SongsLogger) *song {
	return &song{
		repo:   repo,
		logger: logger,
	}
}

func (s *song) CreateSong(ctx context.Context, song *domain.Song) (*domain.Song, error) {
	song, err := s.repo.CreateSong(ctx, song)
	if err != nil {
		s.logger.Error(ctx, "failed to create song", slog.String("err", err.Error()))
		return nil, fmt.Errorf("failed to create song: %w", err)
	}
	return song, nil
}

func (s *song) GetSong(ctx context.Context, groupName, songName string) (*domain.Song, error) {
	if groupName == "" || songName == "" {
		return nil, fmt.Errorf("group name and song name are required")
	}
	song, err := s.repo.GetSong(ctx, groupName, songName)
	if err != nil {
		s.logger.Error(ctx, "failed to get song", slog.String("err", err.Error()))
		return nil, fmt.Errorf("failed to get song: %w", err)
	}
	return song, nil
}

func (s *song) UpdateSong(ctx context.Context, song *domain.Song) (*domain.Song, error) {

	song, err := s.repo.UpdateSong(ctx, song)
	if err != nil {
		s.logger.Error(ctx, "failed to update song", slog.String("err", err.Error()))
		return nil, fmt.Errorf("failed to update song: %w", err)
	}
	return song, nil
}

func (s *song) DeleteSong(ctx context.Context, groupName, songName string) error {
	err := s.repo.DeleteSong(ctx, groupName, songName)
	if err != nil {
		s.logger.Error(ctx, "failed to delete song", slog.String("err", err.Error()))
		return fmt.Errorf("failed to delete song: %w", err)
	}
	return nil
}
