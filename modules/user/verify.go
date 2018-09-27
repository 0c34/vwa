package user

import (
	"fmt"
	"log"

	"net/http"
	"github.com/vwa/util"
	"github.com/vwa/util/render"
	"github.com/vwa/util/session"
	"github.com/vwa/util/database"

	"github.com/julienschmidt/httprouter"
)

type Response struct{
	Body string `json:"body"`
}
func LoginVerify(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	s := session.New()
	if s.IsLoggedIn(r){
		
		uid := s.GetSession(r, "id")

		data, err := GetProfile(uid)

		if err != nil{
			log.Println(err.Error())
		}

		res := Response{}

		profile := fmt.Sprintf(`
			<p><strong>Name :</strong> %s</p>
			<p><strong>Email :</strong> %s</p>
			<p><strong>Phone Number :</strong> %s</p>`,
			data.uname,data.email,data.msisdn)
		
		res.Body =  profile
		render.JSONRender(w, res)

	}else{
		loginForm := fmt.Sprintf(`
		<div class="alert alert-danger" id="msg" style="display:none"></div>	
		<form id="loginform" method="post" action="#" accept-charset="utf-8">
			<fieldset>
				<div class="form-group">
					<input type="text" name="email" value="" class="form-control" placeholder="email" />
				</div>
				<div class="form-group">
					<input type="password" name="password" value="" class="form-control" placeholder="Password" />
				</div>
				
			</fieldset>
		</form> 
		<button id="btnlogin" class="btn btn-success btn-small btn-block">Log in</button>
		<script>
			var loginurl = "%s/login"
            $("#btnlogin").on('click',function(){
              data = $('#loginform').serialize()
              $.post(loginurl, data)
              .done(function(res){
				if(res[0].success == false){
					$("#msg").text(res[0].body)
					$("#msg").show()
					$("#msg").delay(2000).fadeOut();
				}else{
					window.location.reload()
				}
            	}) 
          		}
			)
		</script>`,util.Fullurl)

		res := Response{}
		res.Body = loginForm
		render.JSONRender(w, res)
	}
}

func GetProfile(uid string)(*UserData, error){
	const (
		query = `SELECT username, email, phone_number FROM users where id=$1`
	)
	DB,_ := database.Connect()
	userdata := UserData{}

	stmt,err := DB.Prepare(query)
	if err != nil{
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(uid).Scan(&userdata.uname, &userdata.email, &userdata.msisdn)
	if err != nil{
		return nil, err
	}
	return &userdata, nil
}