package model

import (
	"crypto/rand"
	"encoding/hex"
	"time"

	"github.com/google/uuid"
)

type RefreshToken struct {
	Id        uuid.UUID
	UserId    uuid.UUID
	Token     string `gorm:"uniqueIndex"`
	Revolked  bool
	CreatedAt time.Time
	ExpiredAt time.Time
	User      User `gorm:"constraint:OnDelete:CASCADE"`
}

func (r RefreshToken) isExpired() bool {
	return time.Now().Before(r.ExpiredAt)
}

func generateRefreshToken() (string, error) {
	// Creamos un slice de 32 bytes (equivalente a 256 bits de entropía)
	bytes := make([]byte, 32)

	// Llenamos el slice con bytes aleatorios seguros de forma nativa
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	// Lo convertimos a una cadena Hexadecimal (quedará de 64 caracteres)
	return hex.EncodeToString(bytes), nil
}
func NewRefreshToken(userId uuid.UUID, hours int64) RefreshToken {
	token, err := generateRefreshToken()
	if err != nil {
		panic(err)
	}

	uuid, err := uuid.NewV7()
	if err != nil {
		panic(err)
	}

	return RefreshToken{
		Id:        uuid,
		UserId:    userId,
		Token:     token,
		ExpiredAt: time.Now().Add(time.Hour * time.Duration(hours)),
		Revolked:  false,
	}
}
