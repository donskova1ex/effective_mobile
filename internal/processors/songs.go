package processors

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/donskova1ex/effective_mobile/internal"
	"github.com/donskova1ex/effective_mobile/internal/domain"
)

type SongRepository interface {
	CreateSong(ctx context.Context, song *domain.Song) (*domain.Song, error)
	GetSong(ctx context.Context, groupName, songName string) (*domain.Song, error)
	UpdateSong(ctx context.Context, song *domain.Song) (*domain.Song, error)
	DeleteSong(ctx context.Context, groupName, songName string) error
}
type SongsLogger interface {
	Info(msg string, args ...any)
	Error(msg string, args ...any)
}

type song struct {
	repo   SongRepository
	logger SongsLogger
}

func NewSongProcessor(repo SongRepository, logger SongsLogger) *song {
	return &song{
		repo:   repo,
		logger: logger,
	}
}

func (s *song) CreateSong(ctx context.Context, song *domain.Song) (*domain.Song, error) {
	if song.GroupName == "" || song.SongName == "" {
		return nil, fmt.Errorf("group name and song name are required: %w", internal.ErrEmptyParams)
	}
	if song.ReleaseDate.IsZero() {
		return nil, fmt.Errorf("release date is required: %w", internal.ErrEmptyParams)
	}
	if song.Text == "" {
		return nil, fmt.Errorf("text is required: %w", internal.ErrEmptyParams)
	}
	if song.Link == "" {
		return nil, fmt.Errorf("link is required: %w", internal.ErrEmptyParams)
	}

	song, err := s.repo.CreateSong(ctx, song)
	if err != nil {
		s.logger.Error("failed to create song", slog.String("err", err.Error()))
		return nil, fmt.Errorf("failed to create song: %w", internal.ErrBadRequest)
	}
	return song, nil
}

func (s *song) GetSong(ctx context.Context, groupName, songName string) (*domain.Song, error) {
	if groupName == "" || songName == "" {
		return nil, fmt.Errorf("group name and song name are required: %w", internal.ErrEmptyParams)
	}
	song, err := s.repo.GetSong(ctx, groupName, songName)
	if err != nil {
		s.logger.Error("failed to get song", slog.String("err", err.Error()))
		return nil, fmt.Errorf("failed to get song: %w", internal.ErrBadRequest)
	}
	return song, nil
}

func (s *song) UpdateSong(ctx context.Context, song *domain.Song) (*domain.Song, error) {
	if song.GroupName == "" || song.SongName == "" {
		return nil, fmt.Errorf("group name and song name are required: %w", internal.ErrEmptyParams)
	}
	if song.ReleaseDate.IsZero() {
		return nil, fmt.Errorf("release date is required: %w", internal.ErrEmptyParams)
	}
	if song.Text == "" {
		return nil, fmt.Errorf("text is required: %w", internal.ErrEmptyParams)
	}
	if song.Link == "" {
		return nil, fmt.Errorf("link is required: %w", internal.ErrEmptyParams)
	}

	song, err := s.repo.UpdateSong(ctx, song)
	if err != nil {
		s.logger.Error("failed to update song", slog.String("err", err.Error()))
		return nil, fmt.Errorf("failed to update song: %w", internal.ErrBadRequest)
	}
	return song, nil
}

func (s *song) DeleteSong(ctx context.Context, groupName, songName string) error {
	if groupName == "" || songName == "" {
		return fmt.Errorf("group name and song name are required: %w", internal.ErrEmptyParams)
	}
	err := s.repo.DeleteSong(ctx, groupName, songName)
	if err != nil {
		s.logger.Error("failed to delete song", slog.String("err", err.Error()))
		return fmt.Errorf("failed to delete song: %w", internal.ErrBadRequest)
	}
	return nil
}
