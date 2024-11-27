package tokensService

import (
	"fmt"
	"log"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type TokensService struct {
	db  IDBManager
	jwt IJWTManager
}

type IDBManager interface {
	PutRefreshToken(guid uuid.UUID, token string) error
	GetRefreshToken(guid uuid.UUID) (string, error)
	GetUserEmail(guid uuid.UUID) (string, error)
}

type IJWTManager interface {
	CreateAccessToken(id uuid.UUID, ip string) (string, error)
	CreateRefreshToken() string
	GetSubjectFromToken(token string) (string, error)
}

func New(db IDBManager, jwt IJWTManager) (*TokensService, error) {
	return &TokensService{db: db, jwt: jwt}, nil
}

func (svc *TokensService) GetTokens(id uuid.UUID, ip string) (string, string, error) {
	accessToken, err := svc.jwt.CreateAccessToken(id, ip)
	if err != nil {
		log.Println("Failed to create access token:", err)
		return "", "", err
	}
	refreshToken := svc.jwt.CreateRefreshToken()
	err = svc.db.PutRefreshToken(id, refreshToken)
	if err != nil {
		log.Println("Failed to put refresh token in db:", err)
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

func (svc *TokensService) RefreshTokens(accessToken, refreshToken string, ip string) (string, string, error) {
	subject, err := svc.jwt.GetSubjectFromToken(accessToken)
	if err != nil {
		log.Println("")
	}
	id, err := uuid.Parse(strings.Split(subject, "@")[0])
	if err != nil {
		log.Println("Failed to get id from subject:", err)
		return "", "", err
	}
	tokenIP := strings.Split(subject, "@")[1]
	if ip != tokenIP {
		email, err := svc.db.GetUserEmail(id)
		if err != nil {
			log.Println("Failed to get email:", err)
			return "", "", err
		}
		svc.SendMessage(email)
	}
	dbRefreshToken, err := svc.db.GetRefreshToken(id)
	if err != nil {
		log.Println("Failed to get refresh token:", err)
		return "", "", err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(dbRefreshToken), []byte(refreshToken)); err != nil {
		log.Println("Tokens are not equal")
		return "", "", fmt.Errorf("refresh tokens are not equal")
	}
	return svc.GetTokens(id, ip)
}

func (svc *TokensService) SendMessage(email string) {
	log.Println("Sent message to the", email)
}
