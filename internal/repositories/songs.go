package repositories

import (
	"context"
	"fmt"

	"github.com/donskova1ex/effective_mobile/internal/domain"
)

func (r *Repository) CreateSong(ctx context.Context, song *domain.Song) (*domain.Song, error) {
	query := `
	INSERT INTO songs (group_name, song_name, release_date, text, link)
	VALUES (:group_name, :song_name, :release_date, :text, :link)
	`
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	// TODO: add rollback
	defer tx.Rollback()
	return nil, nil
}
