// user.go
package ssui

import (
	"encoding/base64"
	"fmt"
	"sort"
	"strings"
	"sync"
)

type User struct {
	Name   string
	Pwd    string
	Auth   string
	Phone  string
	Remark string
}

func GetUser(a *HApp, name string) User {
	name = base64.URLEncoding.EncodeToString([]byte(name))
	userLock.Lock()
	defer userLock.Unlock()
	if u, ok := ssuiUsers[name]; ok {
		return decodeUser(a, u)
	}
	return User{}
}

func AddUser(a *HApp, u User) {
	u = encodeUser(a, u)
	userLock.Lock()
	defer userLock.Unlock()
	ssuiUsers[u.Name] = u
	var buf strings.Builder
	fmt.Fprintf(&buf, "%s\t%s\t%s\t%s\t%s\r\n", u.Name, u.Pwd, u.Auth, u.Phone, u.Remark)
	FileWriteAndAppend("ssui.users", buf.String())
}

func AllUser() []string {
	userLock.Lock()
	us := make([]string, 0, len(ssuiUsers))
	for s, _ := range ssuiUsers {
		name, _ := base64.URLEncoding.DecodeString(s)
		us = append(us, string(name))
	}
	userLock.Unlock()
	sort.Strings(us)
	return us
}

var (
	ssuiUsers = make(map[string]User, 0)
	userLock  sync.Mutex
)

func init() {
	cnt := 0
	FileIterateLine("ssui.users", func(i int, l string) bool {
		cnt++
		ss := strings.Split(l, "\t")
		if len(ss) == 5 {
			ssuiUsers[ss[0]] = User{ss[0], ss[1], ss[2], ss[3], ss[4]}
		}
		return true
	})
	if cnt == len(ssuiUsers) {
		return
	}
	var buf strings.Builder
	for _, u := range ssuiUsers {
		fmt.Fprintf(&buf, "%s\t%s\t%s\t%s\t%s\r\n", u.Name, u.Pwd, u.Auth, u.Phone, u.Remark)
	}
	FileCreateAndWrite("ssui.users", buf.String())
}

func encodeUser(a *HApp, ou User) User {
	var nu User
	p, e := Encrypt(ou.Pwd, a.Key)
	if e != nil {
		return nu
	}
	nu.Pwd = p
	nu.Name = base64.URLEncoding.EncodeToString([]byte(ou.Name))
	nu.Auth = base64.URLEncoding.EncodeToString([]byte(ou.Auth))
	nu.Phone = base64.URLEncoding.EncodeToString([]byte(ou.Phone))
	nu.Remark = base64.URLEncoding.EncodeToString([]byte(ou.Remark))
	return nu
}

func decodeUser(a *HApp, ou User) User {
	var nu User
	p, e := Decrypt(ou.Pwd, a.Key)
	if e != nil {
		return nu
	}
	nu.Pwd = p
	name, _ := base64.URLEncoding.DecodeString(ou.Name)
	nu.Name = string(name)
	auth, _ := base64.URLEncoding.DecodeString(ou.Auth)
	nu.Auth = string(auth)
	phone, _ := base64.URLEncoding.DecodeString(ou.Phone)
	nu.Phone = string(phone)
	remark, _ := base64.URLEncoding.DecodeString(ou.Remark)
	nu.Remark = string(remark)
	return nu
}
