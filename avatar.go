package main

import (
	"errors"
	"io/ioutil"
	"path/filepath"
)

//ErrNoAvatarはAvatar instanceが AvatarのURLを返すことができない
var ErrNoAvatarURL = errors.New("chat: アバターのURLを取得できません.")

//Avatarはユーザのプロフィール画像を表す型
type Avatar interface {
	//GetAvatarURLは指定されたクライアントのアバターのurlを返す
	//問題が発生したばあい, エラーを返す。 URLを取得できなかったらErrNoAvatarURLを返す
	GetAvatarURL(u ChatUser) (string, error)
}
type AuthAvatar struct{}

var UseAuthAvatar AuthAvatar

func (_ AuthAvatar) GetAvatarURL(u ChatUser) (string, error) {
	url := u.AvatarURL()
	if url != "" {
		return url, nil
	}
	return "", ErrNoAvatarURL
}

type GravatarAvatar struct{}

var UseGravatar GravatarAvatar

func (_ GravatarAvatar) GetAvatarURL(u ChatUser) (string, error) {
	return "//www.gravatar.com/avatar/" + u.UniqueID(), nil
}


type FileSystemAvatar struct{}

var UseFileSystemAvatar FileSystemAvatar

func (_ FileSystemAvatar) GetAvatarURL(u ChatUser) (string, error) {
	if files, err := ioutil.ReadDir("avatars"); err == nil {
		for _, file := range files {
			if file.IsDir() {
				continue
			}
			if match, _ := filepath.Match(u.UniqueID()+"*", file.Name()); match {
				return "/avatars/" + file.Name(), nil
			}
		}
	}
	return "", ErrNoAvatarURL
}