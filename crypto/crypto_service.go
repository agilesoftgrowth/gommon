package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/agilesoftgrowth/gommon/logger"
)

type CryptoService interface {
	Encrypt(text string) (string, error)
	Decrypt(text string) (string, error)
	encode(b []byte) string
	decode(text string) ([]byte, error)
}

func NewCryptoService(logger logger.LoggerService, cipherKey string) CryptoService {
	return &cryptoService{
		logger:    logger,
		cipherKey: cipherKey,
	}
}

type cryptoService struct {
	logger    logger.LoggerService
	cipherKey string
}

func (s *cryptoService) Encrypt(text string) (string, error) {
	block, err := aes.NewCipher([]byte(s.cipherKey))
	if err != nil {
		s.logger.Error("Cannot create aes block cipher", "error", err.Error())
		return "", fmt.Errorf("cannot encrypt data")
	}
	plaintext := []byte(text)
	bytes := make([]byte, 16)
	rand.Read(bytes)
	iv := bytes
	cfb := cipher.NewCFBEncrypter(block, iv)
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)

	return s.encode(append(iv, ciphertext...)), nil
}

func (s *cryptoService) Decrypt(text string) (string, error) {
	block, err := aes.NewCipher([]byte(s.cipherKey))
	if err != nil {
		s.logger.Error("Cannot create aes block cipher", "error", err.Error())
		return "", errors.New("cannot decrypt data")
	}
	ciphertext, err := s.decode(text)
	if err != nil {
		s.logger.Error("Cannot decode text", "error", err.Error())
		return "", errors.New("text is not base64 encoded")
	}
	iv := ciphertext[:16]
	encryptedText := ciphertext[16:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	plaintext := make([]byte, len(encryptedText))
	cfb.XORKeyStream(plaintext, encryptedText)

	return string(plaintext), nil
}

func (s *cryptoService) encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func (s *cryptoService) decode(text string) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return nil, err
	}
	return data, nil
}
