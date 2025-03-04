package repository

import (
	"database/sql"
	"go-myClient/models"
	"time"
)

type ClientRepo struct {
	db *sql.DB
}

func NewClient(db *sql.DB) *ClientRepo {
	return &ClientRepo{db: db}
}

func (r *ClientRepo) Create(c *models.Client) error {
	query := `INSERT INTO my_client (name, slug, is_project, self_capture, client_prefix, client_logo, address, phone_number, city, created_at, updated_at)
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id`
	return r.db.QueryRow(query, c.Name, c.Slug, c.IsProject, c.SelfCapture, c.ClientPrefix, c.ClientLogo, c.Address, c.PhoneNumber, c.City, time.Now(), time.Now()).Scan(&c.ID)
}

func (r *ClientRepo) Update(c *models.Client) error {
	query := `UPDATE my_client SET name=$1, slug=$2, is_project=$3, self_capture=$4, client_prefix=$5, client_logo=$6, address=$7, phone_number=$8, city=$9, updated_at=$10 WHERE id=$11`
	_, err := r.db.Exec(query, c.Name, c.Slug, c.IsProject, c.SelfCapture, c.ClientPrefix, c.ClientLogo, c.Address, c.PhoneNumber, c.City, time.Now(), c.ID)
	return err
}

func (r *ClientRepo) Delete(id int) error {
	query := `UPDATE my_client SET deleted_at=$1 WHERE id=$2`
	_, err := r.db.Exec(query, time.Now(), id)
	return err
}

func (r *ClientRepo) GetByID(id int) (*models.Client, error) {
	var client models.Client
	query := `SELECT * FROM my_client WHERE id=$1 AND deleted_at IS NULL`
	err := r.db.QueryRow(query, id).Scan(&client.ID, &client.Name, &client.Slug, &client.IsProject, &client.SelfCapture, &client.ClientPrefix, &client.ClientLogo, &client.Address, &client.PhoneNumber, &client.City, &client.CreatedAt, &client.UpdatedAt, &client.DeletedAt)
	if err != nil {
		return nil, err
	}
	return &client, nil
}
