package crud

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"

	"stribog/internal/models"

	"stribog/internal/db"
)

var (
	ErrForgeNotFound = errors.New("forge not found")
	ErrForgeExists   = errors.New("forge already exists")
	ErrInvalidForge  = errors.New("invalid forge data")
)

type ForgeService struct {
	db *db.Pool
}

func NewForgeService(db *db.Pool) *ForgeService {
	return &ForgeService{db: db}
}

func (s *ForgeService) Create(ctx context.Context, forge *models.Forge) error {
	if forge.Title == "" || forge.OwnerID == "" {
		return ErrInvalidForge
	}

	tx, err := s.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	query := `
		INSERT INTO forges (id, title, description, owner_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, NOW(), NOW())
		RETURNING created_at, updated_at
	`
	err = tx.QueryRow(
		ctx,
		query,
		forge.ID,
		forge.Title,
		forge.Description,
		forge.OwnerID,
	).Scan(&forge.CreatedAt, &forge.UpdatedAt)

	if err != nil {
		// Check for unique constraint violations or other errors
		return err
	}

	return tx.Commit(ctx)
}

func (s *ForgeService) GetByID(ctx context.Context, id string) (*models.Forge, error) {
	query := `
		SELECT id, title, description, owner_id, created_at, updated_at
		FROM forges
		WHERE id = $1
	`

	forge := &models.Forge{}
	err := s.db.QueryRow(ctx, query, id).Scan(
		&forge.ID,
		&forge.Title,
		&forge.Description,
		&forge.OwnerID,
		&forge.CreatedAt,
		&forge.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrForgeNotFound
		}
		return nil, err
	}

	return forge, nil
}

func (s *ForgeService) GetByOwnerID(ctx context.Context, ownerID string) ([]*models.Forge, error) {
	query := `
		SELECT id, title, description, owner_id, created_at, updated_at
		FROM forges
		WHERE owner_id = $1
		ORDER BY updated_at DESC
	`

	rows, err := s.db.Query(ctx, query, ownerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var forges []*models.Forge
	for rows.Next() {
		forge := &models.Forge{}
		err := rows.Scan(
			&forge.ID,
			&forge.Title,
			&forge.Description,
			&forge.OwnerID,
			&forge.CreatedAt,
			&forge.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		forges = append(forges, forge)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return forges, nil
}

func (s *ForgeService) Update(ctx context.Context, forge *models.Forge) error {
	_, err := s.GetByID(ctx, forge.ID)
	if err != nil {
		return err
	}

	tx, err := s.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	query := `
		UPDATE forges
		SET title = $1, description = $2
		WHERE id = $3
		RETURNING updated_at
	`
	err = tx.QueryRow(
		ctx,
		query,
		forge.Title,
		forge.Description,
		forge.ID,
	).Scan(&forge.UpdatedAt)

	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (s *ForgeService) Delete(ctx context.Context, id string) error {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	query := `DELETE FROM forges WHERE id = $1`
	commandTag, err := tx.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return ErrForgeNotFound
	}

	return tx.Commit(ctx)
}

func (s *ForgeService) Count(ctx context.Context) (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM forges`

	err := s.db.QueryRow(ctx, query).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (s *ForgeService) Search(ctx context.Context, term string, limit, offset int) ([]*models.Forge, error) {
	query := `
		SELECT id, title, description, owner_id, created_at, updated_at
		FROM forges
		WHERE title ILIKE $1 OR description ILIKE $1
		ORDER BY updated_at DESC
		LIMIT $2 OFFSET $3
	`

	searchTerm := "%" + term + "%"
	rows, err := s.db.Query(ctx, query, searchTerm, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var forges []*models.Forge
	for rows.Next() {
		forge := &models.Forge{}
		err := rows.Scan(
			&forge.ID,
			&forge.Title,
			&forge.Description,
			&forge.OwnerID,
			&forge.CreatedAt,
			&forge.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		forges = append(forges, forge)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return forges, nil
}
