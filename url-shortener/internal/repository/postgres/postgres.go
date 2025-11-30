package postgres

import (
	"context"
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
	query := `INSERT INTO Urls (orginal,shortcode,expire_at) VALUES ($1,$2,$3) RETURNING id`

	err := u.DB.QueryRow(context.Background(), query, original, shortCode, expireAt).Scan(&id)

	return id, err

}

// GetByShortCode implements repository.UrlRepository.
func (u *UrlMyPostgresRepo) GetByShortCode(shortCode string) (string, error) {
	panic("yesterday")
}

func (u *UrlMyPostgresRepo) ShowAll() ([]entity.URL, error) {

	query := `SELECT id, original, short_code, created_at, expire_at FROM urls`

	rows, err := u.DB.Query(context.Background(), query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var urls []entity.URL
	for rows.Next() {
		var u entity.URL

		err := rows.Scan(&u.ID, &u.Original, &u.ShortCode, &u.CreatedAt, &u.ExpireAt)

		if err != nil {
			return nil, err
		}
		urls = append(urls, u)

	}

	return urls, err
}
