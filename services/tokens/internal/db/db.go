package db

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type Config struct {
	Host     string `yaml:"host" env:"HOST" env-default:"localhost"`
	Port     string `yaml:"port" env:"PORT" env-default:"5432"`
	User     string `yaml:"user" env:"USER" env-default:"postgres"`
	Password string `yaml:"password" env:"password" env-default:"postgres"`
	DBName   string `yaml:"dbname" env:"DBNAME" env-default:"chat"`
}

type DB struct {
	config *Config
	db     *pgx.Conn
}

// Создает соединение с существующей БД
func New(cfg *Config) (*DB, error) {
	d := &DB{config: cfg}
	connection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	db, err := pgx.Connect(context.Background(), connection)
	log.Println("Connecting to: " + connection)
	if err != nil {
		return nil, err
	}
	d.db = db
	return d, nil
}

func (d *DB) PutRefreshToken(guid uuid.UUID, token string) error {
	_, err := d.db.Exec(context.Background(), `update public.users set refresh_token=$1 where id=$2`, token, guid)
	if err != nil {
		log.Println("Failed to put refresh token:", err)
		return err
	}
	return nil
}

func (d *DB) GetRefreshToken(guid uuid.UUID) (string, error) {
	var token pgtype.Text
	err := d.db.QueryRow(context.Background(), `select refresh_token from public.users where id=$1`, guid).Scan(&token)
	if err != nil {
		log.Println("Failed to get refresh token:", err)
		return "", err
	}
	return token.String, nil
}

func (d *DB) GetUserEmail(guid uuid.UUID) (string, error) {
	var email pgtype.Text
	err := d.db.QueryRow(context.Background(), `select email from public.users where id=$1`, guid).Scan(&email)
	if err != nil {
		log.Println("Failed to get email:", err)
		return "", err
	}
	return email.String, nil
}
