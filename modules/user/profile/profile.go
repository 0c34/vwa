package profile

import (

	"log"
	"fmt"
	"net/http"

	"github.com/vwa/util/render"
	"github.com/vwa/util/session"
	"github.com/vwa/util/database"
	"github.com/vwa/helper/middleware"

	"github.com/julienschmidt/httprouter"
)


type Self struct{}

func New() *Self {
	return &Self{}
}

func (self *Self) SetRouter(r *httprouter.Router) {
	/* register all router */

	mw := middleware.New() //implement middleware

	r.GET("/user", mw.LoggingMiddleware(mw.CapturePanic(UserHandler)))
	r.POST("/user", mw.LoggingMiddleware(mw.CapturePanic(GetUserHandler)))
	r.GET("/profile", mw.LoggingMiddleware(mw.CapturePanic(ProfileHandler)))
	r.POST("/profile", mw.LoggingMiddleware(mw.CapturePanic(UpdateProfileHandler)))

}

var DB,_ = database.Connect()

type UserData struct{
	UserID 		string `json:"uid"`
	UserName 	string `json:"username"`
	Email 		string `json:"email"`
	MSISDN		string `json:"msisdn"`
}

type Jsonresp struct{
		Success string `json:"success"`
		Data *UserData `json:"data"`
		Message string `json:"message"`
}

func UserHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	
	data := make(map[string]interface{})
	data["title"] = "User"

	render.HTMLRender(w,r, "template.user", data)

}

func GetUserHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	if r.Method == "POST"{
		uid := r.FormValue("uid")
		respdata, err := GetUserData(uid)
		if err != nil{
			resp := Jsonresp{}
			resp.Success = "0"
			resp.Data = respdata
			resp.Message = err.Error()
			render.JSONRender(w, resp)
		}else{
			resp := Jsonresp{}
			resp.Success = "1"
			resp.Data = respdata
			resp.Message = ""
			render.JSONRender(w, resp)
		}
	}
}

func ProfileHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	
	sess := session.New()
	data := make(map[string]interface{})

	if sess.IsLoggedIn(r){
		uid := sess.GetSession(r,"id")
		userdata, err := GetProfile(uid)

		if err != nil{
			log.Println(err.Error())
		}

		data["title"] = "Profile"
		data["uid"] = userdata.UserID
		data["email"] = userdata.Email
		data["name"] = userdata.UserName
		data["msisdn"] = userdata.MSISDN
	}else{
		data["title"] = "Profile"
	}
	

	render.HTMLRender(w,r, "template.profile", data)

}

func UpdateProfileHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	sess := session.New()
	resp := Jsonresp{}

	if sess.IsLoggedIn(r){
		if r.Method == "POST"{
			uid := r.FormValue("uid")
			name := r.FormValue("name")
			email := r.FormValue("email")
			msisdn := r.FormValue("msisdn")
			
			ok := updateProfile(uid, name, email, msisdn)
			if !ok{
				resp.Success = "0"
				resp.Message = "Gagal menperbaharui data"
			}else{
				resp.Success = "1"
				resp.Message = "Data berhasil diperbaharui"
			}
		}
	}else{
		resp.Message = "0"
		resp.Message = "Login untuk dapat memperbaharui data"
	}
	render.JSONRender(w, resp)
}

func GetUserData(uid string)(*UserData, error){
	
	query := fmt.Sprintf("SELECT username, email, phone_number FROM users where id=%s", uid)
	userdata := UserData{} 
	stmt := DB.QueryRow(query)

	err := stmt.Scan(&userdata.UserName, &userdata.Email, &userdata.MSISDN)
	if err != nil{
		return nil, err
	}
	return &userdata, nil
}

func GetProfile(uid string)(*UserData, error){
	const (
		query = `SELECT username, email, phone_number FROM users where id=$1`
	)
	userdata := UserData{}

	stmt,err := DB.Prepare(query)
	if err != nil{
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(uid).Scan(&userdata.UserName, &userdata.Email, &userdata.MSISDN)
	if err != nil{
		return nil, err
	}
	return &userdata, nil
}

func updateProfile(uid string, name string, email string, phone_number string)bool{
	const (
		query = `UPDATE users SET username=$1, email=$2, phone_number=$3 where id = $4`
	)
	_, err := DB.Exec(query, name, email, phone_number, uid)
	if err != nil{
		log.Println(err.Error())
		return false
	}
	return true
}