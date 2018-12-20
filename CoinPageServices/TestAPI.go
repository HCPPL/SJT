package main

import (
  "net/http"
  //"encoding/json"
  //"time"
  
  "github.com/dgrijalva/jwt-go"
  "github.com/auth0/go-jwt-middleware"
  //"github.com/gorilla/mux"
)

var GetTokenInfoSecurity = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	  w.Write([]byte("secret"))
})

var mySigningKey = []byte("ABHGCF^&676VGVF&H(0&8NgGDY")

var GetTokenHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
  
    token := jwt.New(jwt.SigningMethodHS256)
  
    Claims := make(jwt.MapClaims)
    Claims["admin"] = true
    Claims["name"] = "Ado Kukic"

    token.Claims = Claims
  
    tokenString, _ := token.SignedString(mySigningKey)
  
    w.Write([]byte(tokenString))
})

var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
    ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
	      return mySigningKey, nil
	  },
	  SigningMethod: jwt.SigningMethodHS256,
})

/*
func main() {
    router := mux.NewRouter()
    router.Handle("/getTokenInfoSecurity", jwtMiddleware.Handler(GetTokenInfoSecurity)).Methods("GET")
    router.Handle("/get-token", GetTokenHandler).Methods("GET")
    log.Fatal(http.ListenAndServe(":8089", router))
}
*/