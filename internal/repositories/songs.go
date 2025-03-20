package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	"github.com/donskova1ex/effective_mobile/internal"
	"github.com/donskova1ex/effective_mobile/internal/domain"
)

func (r *Repository) CreateSong(ctx context.Context, song *domain.Song) (*domain.Song, error) {
	query := `
	INSERT INTO songs (group_name, song_name, release_date, text, link)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, group_name, song_name, release_date, text, link
	`
	var txCommited bool
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		if !txCommited {
			if err := tx.Rollback(); err != nil {
				r.logger.Error("failed to rollback transaction", slog.String("err", err.Error()))
				return
			}
		}
	}()

	newSong := &domain.Song{}
	err = tx.QueryRowContext(ctx, query,
		song.GroupName,
		song.SongName,
		song.ReleaseDate,
		song.Text,
		song.Link,
	).Scan(
		&newSong.ID,
		&newSong.GroupName,
		&newSong.SongName,
		&newSong.ReleaseDate,
		&newSong.Text,
		&newSong.Link,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create song: %w", err)
	}
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("error committing transaction: %w", err)
	}
	txCommited = true

	return newSong, nil
}

func (r *Repository) GetSong(ctx context.Context, groupName, songName string) (*domain.Song, error) {
	query := `
	SELECT id, group_name, song_name, release_date, text, link
	FROM songs
	WHERE group_name = $1 AND song_name = $2
	`

	song := &domain.Song{}
	err := r.db.QueryRowContext(ctx, query, groupName, songName).Scan(
		&song.ID,
		&song.GroupName,
		&song.SongName,
		&song.ReleaseDate,
		&song.Text,
		&song.Link,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("song not found: %w", internal.ErrNotFound)
		}
		return nil, fmt.Errorf("failed to get song: %w", err)
	}

	return song, nil
}

func (r *Repository) UpdateSong(ctx context.Context, song *domain.Song) (*domain.Song, error) {
	query := `
	UPDATE songs
	SET release_date = $1, text = $2, link = $3
	WHERE group_name = $4 AND song_name = $5
	RETURNING id, group_name, song_name, release_date, text, link
	`
	var txCommited bool
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		if !txCommited {
			if err := tx.Rollback(); err != nil {
				r.logger.Error("failed to rollback transaction", slog.String("err", err.Error()))
				return
			}
		}
	}()

	var updatedSong domain.Song
	err = tx.QueryRowContext(ctx, query,
		song.ReleaseDate,
		song.Text,
		song.Link,
		song.GroupName,
		song.SongName,
	).Scan(
		&updatedSong.ID,
		&updatedSong.GroupName,
		&updatedSong.SongName,
		&updatedSong.ReleaseDate,
		&updatedSong.Text,
		&updatedSong.Link,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("song not found")
		}
		return nil, fmt.Errorf("failed to update song: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return &updatedSong, nil
}

func (r *Repository) DeleteSong(ctx context.Context, groupName, songName string) error {
	query := `
	DELETE FROM songs
	WHERE group_name = $1 AND song_name = $2
	`

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	result, err := tx.ExecContext(ctx, query, groupName, songName)
	if err != nil {
		return fmt.Errorf("failed to delete song: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("song not found: %w", internal.ErrNotFound)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
