package services

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"

	"stribog/internal/db"
	"stribog/internal/models"
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

func (fs *ForgeService) Create(ctx context.Context, forge *models.Forge) error {
	if forge.Title == "" || forge.OwnerID == "" {
		return ErrInvalidForge
	}

	tx, err := fs.db.Begin(ctx)
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
		return err
	}

	return tx.Commit(ctx)
}

func (fs *ForgeService) GetByID(ctx context.Context, id string) (*models.Forge, error) {
	query := `
		SELECT id, title, description, owner_id, created_at, updated_at
		FROM forges
		WHERE id = $1
	`

	forge := &models.Forge{}
	err := fs.db.QueryRow(ctx, query, id).Scan(
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

func (fs *ForgeService) GetByOwnerID(ctx context.Context, ownerID string) ([]models.Forge, error) {
	query := `
		SELECT id, title, description, owner_id, created_at, updated_at
		FROM forges
		WHERE owner_id = $1
		ORDER BY updated_at DESC
	`

	rows, err := fs.db.Query(ctx, query, ownerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var forges []models.Forge
	for rows.Next() {
		forge := models.Forge{}
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

func (fs *ForgeService) Update(ctx context.Context, forge models.Forge) error {
	query := `
		UPDATE forges
		SET title = COALESCE($1, title),
		    description = COALESCE($2, description)
		WHERE id = $3
		RETURNING updated_at
`

	err := fs.db.QueryRow(
		ctx,
		query,
		forge.Title,
		forge.Description,
		forge.ID,
	).Scan(&forge.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (fs *ForgeService) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM forges WHERE id = $1`
	_, err := fs.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (fs *ForgeService) Count(ctx context.Context) (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM forges`

	err := fs.db.QueryRow(ctx, query).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (fs *ForgeService) Search(ctx context.Context, term string, limit, offset int) ([]*models.Forge, error) {
	query := `
		SELECT id, title, description, owner_id, created_at, updated_at
		FROM forges
		WHERE title ILIKE $1 OR description ILIKE $1
		ORDER BY updated_at DESC
		LIMIT $2 OFFSET $3
	`

	searchTerm := "%" + term + "%"
	rows, err := fs.db.Query(ctx, query, searchTerm, limit, offset)
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
