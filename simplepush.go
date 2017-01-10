package simplepush

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"crypto/tls"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Message struct {
	SimplePushKey, Title, Message, Event string
	Password                             string
	Encrypt                              bool
}

var Salt = "1789F0B8C4A051E5"
var APIUrl = "https://api.simplepush.io/"

func paddingPKCS5(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func encrypt(key, iv, buf []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	buf = paddingPKCS5(buf, block.BlockSize())
	if len(buf)%block.BlockSize() != 0 {
		return "", errors.New("Plaintext is not a multiple of the block size")
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(buf, buf)

	return base64.URLEncoding.EncodeToString(buf), nil
}

func Send(m Message) error {
	title := m.Title
	message := m.Message
	var iv []byte

	if m.SimplePushKey == "" {
		return errors.New("simplepush.io key is not allowed to be empty")
	}

	if message == "" { // actually this is not true, but the upstream python lib handles it like that
		return errors.New("Message is not allowed to be empty")
	}

	if m.Encrypt {
		var err error
		sha := sha1.Sum([]byte(m.Password + Salt))
		key, _ := hex.DecodeString(fmt.Sprintf("%X", sha)[:32])
		iv = make([]byte, aes.BlockSize)

		if _, err := io.ReadFull(rand.Reader, iv); err != nil {
			return err
		}

		if title, err = encrypt(key, iv, []byte(title)); err != nil {
			return err
		}
		if message, err = encrypt(key, iv, []byte(message)); err != nil {
			return err
		}
	}

	data := url.Values{}
	data.Set("key", m.SimplePushKey)
	data.Add("title", title)
	data.Add("msg", message)
	data.Add("event", m.Event)
	if m.Encrypt {
		data.Add("encrypted", "true")
		data.Add("iv", fmt.Sprintf("%X", iv))
	}

	u, err := url.ParseRequestURI(APIUrl)
	if err != nil {
		return err
	}
	resource := "/send"
	u.Path = resource
	urlStr := fmt.Sprintf("%v", u)

	// actually the certificate looks OK, but it was not possible to send messages without this config :-/
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
	}
	client := &http.Client{Transport: transCfg}

	resp, err := client.PostForm(urlStr, data)
	if err != nil {
		return err
	}
	resp.Body.Close()

	return nil
}
