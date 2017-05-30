package main

import (
	"crypto/rand"
	"fmt"
	"io"
	//"math/rand"

	"time"
)

type User struct {
	Id                int        `json:"id"`
	AgentToken_Id     int        `json:"agent_id"`
	AgentToken        AgentToken ` json:"agent"`
	Registration_Date time.Time  `json:"registration_date"`
	Birth_date        time.Time  `json:"birth_date"`
	Lat_Init          float32    `json:"lat_init"`
	Lng_Init          float32    `json:"lng_init"`
	Lat_Current       float32    `json:"lat_Current"`
	Lng_Current       float32    `json:"lng_Current"`
	// 0 female , 1 male
	Gender    int    `json:"gender"`
	Image_Url string `json:"image_url"`
	//Is_Active bool   `json:"is_active"`
	Device    Device `json:"device"`
	Device_Id int    `json:"device_id"`
}
type Device struct {
	Id       int    `json:"id"`
	Agent_Id int    `json:"agent_id"`
	Model    string `json:"model"`
	Platform string `json:"platform"`
	Uuid     string `json:"uuid"`
	Version  string `json:"version"`
	Ip       string `json:"ip"`
	CM_Id    string `json:"cm_id" `
}
type AgentToken struct {
	Id            int       `json:"id"`
	User_Id       int       `json:"user_id"`
	Token         string    `json:"token"`
	Is_Expire     bool      `json:"is_expire"`
	Creation_Date time.Time `json:"creation_date"`
	Type          string    `json:"type"`
}

type UserCollection struct {
}

func (col *UserCollection) NewUUID() (string, error) { //create token
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil

}
func (col *UserCollection) Create_User(user *User) (string, int, error) {
	var id int
	user.Registration_Date = time.Now().Local()
	err := db.QueryRow("insert into users(gender,birth_date,registration_date,lat_init,lng_init) values($1,$2,$3,$4,$5) returning id", user.Gender, user.Birth_date, user.Registration_Date, user.Lat_Init, user.Lng_Init).Scan(&id)
	fmt.Println("last inserted id", id)
	if err != nil {
		fmt.Println("error in Create_User", err)
		return "", 0, err
	}
	token, err := col.NewUUID()
	if err != nil {
		fmt.Printf("error in create token in login: %v\n", err)
		return "", 0, err
	}
	typetoken := "Token"
	Creation_Date := time.Now().Local()
	var token_id int
	err = db.QueryRow("insert into agent_tokens(token,user_id,type,creation_date) values($1,$2,$3,$4) returning id", token, id, typetoken, Creation_Date).Scan(&token_id)
	fmt.Println("last inserted token_id", token_id)
	if err != nil {
		fmt.Println("error insert user", err)
		return "", 0, err
	}

	_, err = db.Exec(fmt.Sprintf("insert into device(usertokens_id,model,platform,uuid,version,ip,cm_id) values (%d,'%s','%s','%s','%s','%s','%s')", token_id, user.Device.Model, user.Device.Platform, user.Device.Uuid, user.Device.Version, user.Device.Ip, user.Device.CM_Id))
	if err != nil {
		fmt.Println("error insert in device", err)
		return "", 0, err
	}
	return token, id, nil
}
func (col *UserCollection) Logout(user *User) (string, error) {
	var err error
	var is_expire bool
	var token string
	err = db.Get(&token, fmt.Sprintf("select token from agent_tokens where user_id=%d", user.Id))
	if err != nil {
		fmt.Println("error in get is_expire in updatelocation", err)
		return "", err
	}
	err = db.Get(&is_expire, fmt.Sprintf("select is_expire from agent_tokens where user_id=%d and token='%s'", user.Id, user.AgentToken.Token))
	if err != nil {
		fmt.Println("error in get is_expire in updatelocation", err)
		return "", err
	}
	if token == "" && token != user.AgentToken.Token {
		var err error
		err = ErrTokenInvalid
		text := "Can Not LogOut  BCZ Token Is Not Valid"
		return text, err
	} else if is_expire {
		var err error
		err = ErrNotFound
		text := "Before LogingOut  BCZ This User Is Expire "
		return text, err
	} else {
		_, err := db.Exec(fmt.Sprintf("update agent_tokens set is_expire=true where token= '%s'", user.AgentToken.Token))
		if err != nil {
			return "", err
		}
	}

	text := "LogOut Is Succefully"
	return text, nil

}
func (col *UserCollection) UpdateLocation(user *User) (string, error) {
	var err error
	var is_expire bool
	var token string
	// var lat_init, lng_init float32
	err = db.Get(&token, fmt.Sprintf("select token from agent_tokens where user_id=%d", user.Id))
	if err != nil {
		fmt.Println("error in get is_expire in updatelocation", err)
		return "", err
	}
	err = db.Get(&is_expire, fmt.Sprintf("select is_expire from agent_tokens where user_id=%d and token='%s'", user.Id, user.AgentToken.Token))
	if err != nil {
		fmt.Println("error in get is_expire in updatelocation", err)
		return "", err
	}
	if token == "" && token != user.AgentToken.Token {
		var err error
		err = ErrTokenInvalid
		text := "Can Not Update Location BCZ Token Is Not Valid"
		return text, err

	} else if is_expire {
		var err error
		err = ErrNotFound
		text := "Can Not Update Location BCZ This User Is Expire "
		return text, err
	} else {
		_, err := db.Exec(fmt.Sprintf("update users set  lat_current=%f where id= %d", user.Lat_Current, user.Id))
		if err != nil {
			fmt.Println("error in update lng", err)
			return "", err
		}
		_, err1 := db.Exec(fmt.Sprintf("update users set  lng_current=%f where id= %d", user.Lng_Current, user.Id))
		if err1 != nil {
			fmt.Println("error in update lng", err1)
			return "", err1
		}
	}

	text := "Update Location Is Succefully"
	return text, nil
	// err = db.Get(&lat_init, fmt.Sprintf("select  lat_init from users where id=%d", user.Id))
	// if err != nil {
	// 	fmt.Println("error in get  lat_init in updatelocation", err)
	// 	return "", err
	// }
	// err = db.Get(&lng_init, fmt.Sprintf("select  lng_init from users where id=%d", user.Id))
	// if err != nil {
	// 	fmt.Println("error in get  lng_init in updatelocation", err)
	// 	return "", err
	// }
	// if lat_init == 0 {
	// 	lat_init = user.Lat_Current
	// 	_, err := db.Exec(fmt.Sprintf("update users set lat_init=%f where id= %d", lat_init, user.Id))
	// 	if err != nil {
	// 		fmt.Println("error in update lat ", err)
	// 		return "", err
	// 	}
	// }
	// if lng_init == 0 {
	// 	lng_init = user.Lng_Current
	// 	_, err := db.Exec(fmt.Sprintf("update users set lng_init=%f where id= %d", lng_init, user.Id))
	// 	if err != nil {
	// 		fmt.Println("error in update lng", err)
	// 		return "", err
	// 	}

	// }

}

func (col *UserCollection) Get_User(i int) (User, error) {

	user := User{}
	err := db.Get(&user, fmt.Sprintf("select * from users where id=%d", i))
	if err != nil {
		return user, err
	}
	return user, nil
}
