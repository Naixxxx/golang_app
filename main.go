package main

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gogo/protobuf/proto"
	"github.com/gorilla/websocket"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func main() {
	http.HandleFunc("/", helloHandler)

	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}

func JWTlib() string {
	token := jwt.New(jwt.SigningMethodHS256)
	ss, _ := token.SignedString([]byte("secret"))

	return ss
}

func fakeGIN() {
	r := gin.Default()
	_ = r
}

func mockWebsocket() {
	_ = websocket.IsCloseError
}

func mockProtoBuff() {
	_ = proto.ErrInternalBadWireType
}
