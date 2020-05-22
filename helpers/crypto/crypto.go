package crypto

import (
	"encoding/base64"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// 関数（引数：型）（戻り値 タプル）
func PasswordEncrypt(password string) (string, error) {
	// 複数変数代入、型変換
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hash), err
}

func CompareHashAndPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func SecureRandom() string {
	return uuid.New().String()
}

func SecureRandomBase64() string {
	return base64.StdEncoding.EncodeToString(uuid.New().NodeID())
}

func LongSecureRandomBase64() string {
	return SecureRandomBase64() + SecureRandomBase64()
}

func MultipleSecureRandomBase64(n int) string {
	if n <= 1 {
		return SecureRandomBase64()
	}
	return SecureRandomBase64() + MultipleSecureRandomBase64(n-1)
}
