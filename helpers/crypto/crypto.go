package crypto

import (
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
