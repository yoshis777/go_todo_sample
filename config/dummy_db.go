package config

import (
	"errors"
	"github.com/go_todo_sample/helpers/crypto"
)

// 構造体の定義
type DummyUserModel struct {
	Username      string
	Password      string
	Email         string
	authenticated bool
}

// ポインタ変数により、メモリを確保する（値で指定するより、使用するメモリ量が小さくなる）
func NewDummyUser(username, email string) *DummyUserModel {
	return &DummyUserModel{
		Username: username,
		Email:    email,
	}
}

// レシーバ変数（を使用した関数をメソッドという）
// 構造体のDummyUserModelにメソッドを付与してオブジェクトクラスとすることができる（オブジェクト志向らしく定義できる）
func (u *DummyUserModel) SetPassword(password string) error {
	hash, err := crypto.PasswordEncrypt(password)
	if err != nil {
		return err
	}

	u.Password = hash
	return nil
}

func (u *DummyUserModel) Authenticate() {
	u.authenticated = true
}

type DummyDatabase struct {
	// map[string]: マップ（連想配列）
	// interface{}: どんな型でも格納可
	database map[string]interface{}
}

var store DummyDatabase

func init() {
	store.database = map[string]interface{}{}
}

func DummyDB() *DummyDatabase {
	return &store
}

func (db *DummyDatabase) Exists(username string) bool {
	_, exist := db.database[username]

	return exist
}

func (db *DummyDatabase) SaveUser(username, email, password string) error {
	if db.Exists(username) {
		return errors.New("user \"" + username + "\" already exists")
	}

	user := NewDummyUser(username, email)
	if err := user.SetPassword(password); err != nil {
		return err
	}
	db.database[username] = user

	return nil
}

func (db *DummyDatabase) GetUser(username, password string) (*DummyUserModel, error) {
	buffer, exists := db.database[username]

	if !exists {
		return nil, errors.New("user \"" + username + "\" doesn't exists")
	}

	user := buffer.(*DummyUserModel) //ポインタの中身を取得（userがポインタ変数のため）
	if err := crypto.CompareHashAndPassword(user.Password, password); err != nil {
		return nil, errors.New("user \"" + username + "\" doesn't exists")
	}
	return user, nil
}
