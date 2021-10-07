package userhandler

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"time"
	"../../model/usermodel"
	"../../databaseConnect"

)


var Jwtkey = []byte("secret-key")



func GetUser(w http.ResponseWriter, r *http.Request){
	var user []usermodel.Users;
	databaseConnect.DB.Find(&user);
	json.NewEncoder(w).Encode(&user);
}

func CreateUser(w http.ResponseWriter, r *http.Request){
	var user usermodel.Users;
	json.NewDecoder(r.Body).Decode(&user);
	createdUser := databaseConnect.DB.Create(&user);
	err := createdUser.Error
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}else{
		json.NewEncoder(w).Encode(&user)
	}
}

func AuthenticateUser(w http.ResponseWriter, r *http.Request){
	var user usermodel.Users;
	json.NewDecoder(r.Body).Decode(&user);
	//checking for whether the data gets properly or not

	//condition when user make either username and password entry as empty
	if len(user.Username) == 0 || len(user.Password) == 0{
		w.WriteHeader(http.StatusUnauthorized)
	}else{
	//condition for checking whether user exist in the database(check by email id) and check whether the correspond password is equal to given password
			var newuser usermodel.Users;
			test := databaseConnect.DB.First(&newuser,"users.Username = ?" ,user.Username)//user.Username
			if test.Error != nil || newuser.Password != user.Password{
				w.WriteHeader(http.StatusUnauthorized)
			}else{
				expirationTime := time.Now().Add(time.Minute * 5) 

				claim := &usermodel.Claims{
					Username: newuser.Username,
					UsersID: newuser.ID,
					StandardClaims: jwt.StandardClaims{
						ExpiresAt : expirationTime.Unix(),
					}}
	
				token := jwt.NewWithClaims(jwt.SigningMethodHS256,claim)
				tokenString,err := token.SignedString(Jwtkey)
	
				if err != nil{
					w.WriteHeader(http.StatusInternalServerError)
				}else{
					http.SetCookie(w,
						&http.Cookie{
							Name : "token",
							Value : tokenString,
							Expires: expirationTime,
						})
				}
			}
	}
}


func UserDetails(w http.ResponseWriter, r *http.Request){
	cookie,err := r.Cookie("token")
	if err != nil{
		w.WriteHeader(http.StatusUnauthorized)
	}else{
		tokenStr := cookie.Value
		claim := &usermodel.Claims{}
		tkn,err := jwt.ParseWithClaims(tokenStr,claim,
			func(t *jwt.Token) (interface{},error){
				return Jwtkey,nil
		})
		if err != nil || !tkn.Valid{
			w.WriteHeader(http.StatusUnauthorized)
		}else{
			w.Write([]byte(fmt.Sprintf("hello %s", claim.Username)))
		}
	}
}


func RefreshToken(w http.ResponseWriter, r *http.Request){
	cookie,err := r.Cookie("token")
	if err != nil{
		w.WriteHeader(http.StatusUnauthorized)
	}else{
		tokenStr := cookie.Value
		claim := &usermodel.Claims{}
		tkn,err := jwt.ParseWithClaims(tokenStr,claim,
			func(t *jwt.Token) (interface{},error){
				return Jwtkey,nil
		})
		if err != nil || !tkn.Valid{
			w.WriteHeader(http.StatusUnauthorized)
		}else{
			if time.Unix(claim.ExpiresAt,0).Sub(time.Now()) > 30 * time.Second{
				w.WriteHeader(http.StatusBadRequest)
			} else{
				expirationTime := time.Now().Add(time.Minute * 5)
				claim.ExpiresAt = expirationTime.Unix() 
				token := jwt.NewWithClaims(jwt.SigningMethodHS256,claim)
				tokenString,err := token.SignedString(Jwtkey)
	
				if err != nil{
					w.WriteHeader(http.StatusInternalServerError)
				}else{
					http.SetCookie(w,
						&http.Cookie{
							Name : "refresh_token",
							Value : tokenString,
							Expires: expirationTime,
						})
				}
			}
		}
	}
}

