//err := db.Exec("insert into users(phone_number) values ($1)", user.Phone_number)
//err := db.Exec(&userid, fmt.Sprintf("select id from users where phone_number=%s", phone))
//err := db.select(&userid, fmt.Sprintf("select id from users where phone_number=%s", phone))
//fmt.Println("userid:", userid)
//err = db.QueryRow("INSERT INTO users(phone_number) VALUES($1) returning id;", user.Phone_number).Scan(&lastInsertId)
package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"math/big"
	"strings"
	//"math/rand"
	"strconv"
	"time"

	"github.com/negah/kavenegar"
)

type Users struct {
	Id                int        `json:"id"`
	AgentToken_Id     int        `json:"agent_id"`
	AgentToken        AgentToken ` json:"agent"`
	Phone_number      string     `json:"phone_number" validate:"required"`
	Firstname         string     `json:"firstname"`
	Lastname          string     `json:"lastname"`
	Registration_Date time.Time  `json:"registration_date"`
	Lng               float32    `json:"lng"`
	Lat               float32    `json:"lat"`
	Email             string     `json:"email"`
	Image_Url         string     `json:"image_url"`
	Is_Spammed        bool       `json:"is_spammed"`
	Is_Active         bool       `json:"is_active"`
}
type AgentToken struct {
	Id            int       `json:"id"`
	User_Id       int       `json:"user_id"`
	Token         string    `json:"token"`
	Is_Expire     bool      `json:"is_expire"`
	Creation_Date time.Time `json:"creation_date"`
	Type          string    `json:"type"`
	Device_Id     int       `json:"device_id"`
	Device        Device    `db:"-" json:"device"`
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

type UserCollection struct {
}

func (col *UserCollection) IsExists(phone string) (bool, error) {

	var intCount int
	err := db.Get(&intCount, fmt.Sprintf("select count(*) from users where phone_number='%s'", phone))
	if err != nil {
		return false, err
	}
	if intCount > 0 {
		return true, nil
	}

	return false, nil
}
func (col *UserCollection) Create_User(phone_number string) (error, int) {
	// in login we should insert phone_number
	var userid int
	user := Users{}
	user.Phone_number = phone_number
	user.Registration_Date = time.Now().Local()

	_, err := db.Exec(fmt.Sprintf("insert into users(phone_number,registration_date) values ('%s','%s')", user.Phone_number, user.Registration_Date))

	if err != nil {
		fmt.Println("error in Create_User", err)
		return err, 0
	}

	err = db.Get(&userid, fmt.Sprintf("select id from users where phone_number='%s'", user.Phone_number))
	if err != nil {
		fmt.Println("error in id in creat_user", err)
		return err, 0
	}
	return nil, userid
}
func (col *UserCollection) SixDigits() (int64, error) {
	// create random digit
	max := big.NewInt(999999)
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return 0, err
	}
	return n.Int64(), nil
}
func (col *UserCollection) Get_Userid(phone_number string) (int, error) { // id from users

	var id int
	err := db.Get(&id, fmt.Sprintf("select id from users where phone_number='%s'", phone_number))

	if err != nil {
		fmt.Println("error from Get_Id for getting id", err)
		return 0, err
	}
	fmt.Println("GET_ID:", id)
	return id, nil
}
func (col *UserCollection) Create(Phone_number string) (error, int) {
	apikey := "6B4655787862666169592B50514C704379354D5264773D3D"
	isExist, err := col.IsExists(Phone_number)
	if err != nil {
		fmt.Println("error in exit phone number for new user", err)
		return err, 0
	}
	if !isExist {
		err, userid := col.Create_User(Phone_number)
		if err != nil {
			fmt.Println("error in create user by phone_number for new user", err)
			return err, 0
		}
		code, err := col.SixDigits()
		if err != nil {
			fmt.Println("error to create random digit in new user", err)
			return err, 0
		}
		active_code := strconv.FormatInt(code, 10)
		sms := kavenegar.NewSMS(apikey, "")
		status, err := sms.Send(Phone_number, active_code)

		if err != nil { //can not send sms
			fmt.Println("status sms", status)
			fmt.Println("error in send sms for new user", err)
			return err, 0
		}
		typecode := "Active_Code"
		creation_date := time.Now().Local()

		_, err = db.Exec(fmt.Sprintf("insert into agent_tokens(token,user_id,type,creation_date) values ('%s',%d,'%s','%s')", active_code, userid, typecode, creation_date))
		if err != nil {
			fmt.Println("error in insert", err)
			return err, 0
		}
		id, err := col.Get_Userid(Phone_number)
		return nil, id
	} else { // old user

		code, err := col.SixDigits()
		if err != nil {
			fmt.Println("error in create random digit in old user", err)
			return err, 0
		}
		sms := kavenegar.NewSMS(apikey, "")

		active_code := strconv.FormatInt(code, 10)

		status, err := sms.Send(Phone_number, active_code)

		if err != nil {
			fmt.Println("status sms", status)
			fmt.Println("err in send sms in old user", err)
			return err, 0
		}

		id, err := col.Get_Userid(Phone_number)
		if err != nil {
			fmt.Println("error in get id from users in old user", err)
			fmt.Println("id in old user", id)
			return err, 0
		}
		_, err = db.Exec(fmt.Sprintf("update agent_tokens set  is_expire= %t  where user_id=%d", true, id))
		if err != nil {
			fmt.Println("error in update is_expire in create user", err)
			return err, 0
		}
		typecode := "Active_Code"
		creation_date := time.Now().Local()
		_, err = db.Exec(fmt.Sprintf("insert into agent_tokens(token,user_id,type,creation_date) values ('%s',%d,'%s','%s')", active_code, id, typecode, creation_date))
		if err != nil {
			fmt.Println("error in insert active code in old user", err)
			return err, 0
		}
		return nil, id
	}

}
func (col *UserCollection) Login(user *Users) (Users, error) {
	var code string
	var err error
	var usertokens_id int
	typecode := "Active_Code"

	err = db.Get(&code, fmt.Sprintf("select token from agent_tokens where user_id=%d and is_expire=false and type='%s'", user.Id, typecode))
	if err != nil {
		fmt.Println("error in get active code in login", err)
		return *user, err
	}

	if user.AgentToken.Token == code {

		token, err := col.NewUUID()

		if err != nil {
			fmt.Printf("error in create token in login: %v\n", err)
			return *user, err
		}
		typetoken := "Token"
		Creation_Date := time.Now().Local()
		_, err = db.Exec(fmt.Sprintf("insert into agent_tokens(token,user_id,type,creation_date) values ('%s',%d,'%s','%s')", token, user.Id, typetoken, Creation_Date))
		if err != nil {
			fmt.Println("error insert user", err)
			return *user, err
		}

		_, err = db.Exec(fmt.Sprintf("update users  set is_active=true  where id=%d ", user.Id))
		if err != nil {
			fmt.Println("error in update is_active in login in login", err)
			return *user, err
		}

		usertokens_id, err = col.Get_Tokenid(token)
		if err != nil {
			fmt.Println("error in get token id in device", err)
			return *user, err
		}

		_, err = db.Exec(fmt.Sprintf("insert into device(usertokens_id,model,platform,uuid,version,ip,cm_id) values (%d,'%s','%s','%s','%s','%s','%s')", usertokens_id, user.AgentToken.Device.Model, user.AgentToken.Device.Platform, user.AgentToken.Device.Uuid, user.AgentToken.Device.Version, user.AgentToken.Device.Ip, user.AgentToken.Device.CM_Id))
		if err != nil {
			fmt.Println("error insert in device", err)
			return *user, err
		}
		var device_id int
		err = db.Get(&device_id, fmt.Sprintf("select id from device where usertokens_id=%d", usertokens_id))
		if err != nil {
			fmt.Println("error in get active code in login", err)
			return *user, err
		}

		user.AgentToken.Token = token
		user.AgentToken.User_Id = user.Id
		user.AgentToken.Creation_Date = time.Now().Local()
		user.AgentToken.Type = typetoken
		user.AgentToken.Device.Agent_Id = usertokens_id
		user.Is_Active = true
		user.AgentToken.Device.Id = device_id

		return *user, nil
	}
	return *user, err
}
func (col *UserCollection) Exists(token string) (bool, error) {
	result := AgentToken{}
	typetoken := "Token"
	err := db.Get(&result, fmt.Sprintf("select * from agent_tokens where is_expire=false and type='%s' and token= '%s' ", token, typetoken))

	if err != nil {
		return false, err
	}

	return true, err
}
func (col *UserCollection) Get_UserIdbytoken(token string) (int, error) {
	var err error
	var userid int
	err = db.Get(&userid, fmt.Sprintf("select id from agent_tokens where is_expire=false and token= '%s' ", token))
	if err != nil {
		return 0, err
	}
	return userid, nil
}
func (col *UserCollection) Get_Tokenid(token string) (int, error) {
	var token_id int
	var typetoken string
	typetoken = "Token"

	err := db.Get(&token_id, fmt.Sprintf("select id from agent_tokens where token='%s' and is_expire=false and type='%s'", token, typetoken))

	if err != nil {
		fmt.Println("error from Get_Id for getting id", err)
		return 0, err
	}
	fmt.Println("GET_ID:", token_id)
	return token_id, nil
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

func (col *UserCollection) Logout(user *Users) (Users, error) {
	typetoken := "Token"
	_, err := db.Exec(fmt.Sprintf("update agent_tokens set is_expired = true where token= '%s' and type='%s' ", user.AgentToken.Token, typetoken))
	if err != nil {
		return *user, err
	}
	typecode := "Active_Code"
	_, err = db.Exec(fmt.Sprintf("update agent_tokens set is_expired = true where token= '%s' and type='%s' ", user.AgentToken.Token, typecode))
	if err != nil {
		return *user, err
	}
	_, err = db.Exec(fmt.Sprintf("update users set is_active =false where phone= '%s'", user.Phone_number))
	if err != nil {
		return *user, err
	}
	return *user, nil
}

func (col *UserCollection) EditProfile(user *Users) (Users, error) {
	var err error

	if user.AgentToken.Token != "" {

		_, err = db.Exec(fmt.Sprintf("update  users  set firstname= '%s',lastname='%s',email='%s',phone_number='%s' where id=%d and is_active='%t'", user.Firstname, user.Lastname, user.Email, user.Phone_number, user.Id, true))
		fmt.Println("error in edit profile in update", err)
		if err != nil {
			fmt.Println("error in edit profile in update", err)
			return *user, err
		}
		return *user, nil
	}
	return *user, err

}

func (col *UserCollection) List(page, pageSize int) ([]SimpleSkill, error) {
	var sUsers []SimpleSkill
	// var agentoken AgentToken
	// var id int
	var err error
	var condition, fullQuery, pagingQuery, selectQ, query string
	// agentoken, err = col.Get_Token(token)
	// fmt.Println("eeror in get token", err)
	// fmt.Println("agentoken in list function", agentoken)
	// if err != nil {

	// 	return nil, err
	// }
	// id, err = col.Get_UserIdbytoken(token)
	// fmt.Println("error in get token id in list function", err)
	// fmt.Println("id in list function", id)
	// if err != nil {

	// 	return nil, err
	// }
	condition = " where users.is_spammed = false  and users.is_active=true "
	if page > 0 {
		offset := (page - 1) * pageSize
		pagingQuery = " LIMIT " + strconv.Itoa(pageSize) + " OFFSET " + strconv.Itoa(offset)
	}
	selectQ = " select users.registration_date,users.firstname,users.lastname,users.email,users.lng,users.lat,agent_tokens.user_id"
	query = " from users  inner join agent_tokens on users.id = agent_tokens.user_id " + condition
	fullQuery = selectQ + query
	if fullQuery != "" {
		fullQuery += pagingQuery
	}

	err = db.Select(&sUsers, fullQuery)
	return sUsers, err
}
func (col *UserCollection) EscapeChars(str string) string {
	return strings.Replace(str, "'", "''", -1)
}

// func (col *UserCollection) Get_Token(token string) (AgentToken, error) {
// 	var err error
// 	typetoken := "Token"
// 	result := AgentToken{}
// 	err = db.Get(&result, fmt.Sprintf("select * from agent_tokens where is_expire=false and type='%s' and token= '%s' ", false, token, typetoken))
// 	if err != nil {
// 		return AgentToken{}, err
// 	}
// 	return AgentToken{}, nil
// }
