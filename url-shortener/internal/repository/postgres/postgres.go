package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rezaabaskhanian/go-url-shortener/internal/entity"
	"github.com/rezaabaskhanian/go-url-shortener/internal/repository"
)

type UrlMyPostgresRepo struct {
	DB *pgxpool.Pool
}

func NewMyPostgres(db *pgxpool.Pool) repository.UrlRepository {
	return &UrlMyPostgresRepo{DB: db}
}

// Create implements repository.UrlRepository.

func (u *UrlMyPostgresRepo) Create(original string, shortCode string, expireAt time.Time) (int64, error) {
	var id int64
	query := `INSERT INTO Urls (original,short_code,expire_at) VALUES ($1,$2,$3) RETURNING id`

	err := u.DB.QueryRow(context.Background(), query, original, shortCode, expireAt).Scan(&id)

	return id, err

}

func (r *UrlMyPostgresRepo) ShowAll() ([]entity.URL, error) {
	query := `SELECT id, original, short_code, created_at, expire_at FROM urls`

	rows, err := r.DB.Query(context.Background(), query)

	if err != nil {
		return nil, fmt.Errorf("khtata dra database %v", err)
	}

	defer rows.Close()

	var urls []entity.URL

	for rows.Next() {
		var item entity.URL
		if err := rows.Scan(
			&item.ID,
			&item.Original,
			&item.ShortCode,
			&item.CreatedAt,
			&item.ExpireAt,
		); err != nil {
			return nil, fmt.Errorf("scan error: %w", err)
		}
		urls = append(urls, item)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("rows iteration error: %w", rows.Err())
	}

	return urls, nil
}

// GetByShortCode implements repository.UrlRepository.
func (u *UrlMyPostgresRepo) GetByShortCode(shortCode string) (entity.URL, error) {
	query := `SELECT id, original, short_code, created_at, expire_at FROM urls WHERE short_code = $1`

	var res entity.URL

	row := u.DB.QueryRow(context.Background(), query, shortCode)

	err := row.Scan(&res.ID, &res.Original, &res.ShortCode, &res.CreatedAt, &res.ExpireAt)
	if err != nil {
		fmt.Println("Scan error:", err)
	}

	return res, nil

}
