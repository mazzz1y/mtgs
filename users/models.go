package users

import (
	"crypto/rand"
	"encoding/hex"
	"mtg/config"
	"strconv"

	consul "github.com/hashicorp/consul/api"
)

// User defines struct for KV
type User struct {
	Name   string `json:"name"`
	Secret []byte `json:"secret"`
}

var kv *consul.KV

// InitKV initialize consul client
func InitKV(conf *config.Config) {
	client, err := consul.NewClient(&consul.Config{
		Address: conf.ConsulHost + ":" + strconv.Itoa(int(conf.ConsulPort)),
		Scheme:  "http",
	})
	if err != nil {
		panic(err)
	}
	kv = client.KV()
}

// Create method generate token and put user it into KV storage
func (u User) Create() (User, error) {
	u.Secret = generateSecret()
	p := &consul.KVPair{Key: u.Name, Value: u.Secret}
	_, err := kv.Put(p, nil)
	return u, err
}

// Exist method check if user exist in KV storage
func (u User) Exist() bool {
	res, _, _ := kv.Get(u.Name, nil)
	if res != nil {
		return true
	}
	return false
}

// Delete method remove user from KV storage
func (u User) Delete() error {
	_, err := kv.Delete(u.Name, nil)
	return err
}

// GetAll method return list of users from KV storage
func (u User) GetAll() ([]User, error) {
	var users []User

	list, _, err := kv.List("", nil)

	if err != nil {
		return users, err
	}

	for _, u := range list {
		users = append(users, User{u.Key, u.Value})
	}

	return users, err

}

func generateSecret() []byte {
	const lenght = 16
	src := make([]byte, lenght)
	rand.Read(src)

	hexString := hex.EncodeToString(src)
	h, _ := hex.DecodeString(hexString)
	return h
}
