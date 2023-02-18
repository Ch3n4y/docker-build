package main

import (
	"C"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/tendermint/tendermint/crypto/secp256k1"
	"math/rand"
	"net/http"
	"time"
)

func randomString(l int) []byte {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		rand.NewSource(time.Now().UnixNano())
		bytes[i] = byte(randInt(1, 2^256-1))
	}
	return bytes
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

type Res struct {
	Signature string `json:"signature"`
	PublicKey string `json:"publicKey"`
}

type Req struct {
	AppId    string `json:"appId"`
	DeviceId string `json:"deviceId"`
	UserId   string `json:"userId"`
	Nonce    int    `json:"nonce"`
}

func Sign(appId string, deviceId string, userId string, nonce int) Res {
	max := 32
	key := randomString(max)
	data := fmt.Sprintf("%s:%s:%s:%d", appId, deviceId, userId, nonce)
	privKey := secp256k1.PrivKey(key)
	pubKey := privKey.PubKey()
	publicKey := "04" + hex.EncodeToString(pubKey.Bytes())
	signature, err := privKey.Sign([]byte(data))
	if err != nil {
		return Res{}
	}
	return Res{Signature: hex.EncodeToString(signature) + "01", PublicKey: publicKey}
}

func main() {
	//res := Sign("appId", "deviceId", "userId", 1)
	//fmt.Println(res)
	r := gin.Default()
	r.GET("/sign", func(c *gin.Context) {
		req := Req{}
		err := binding.JSON.Bind(c.Request, &req)
		if err != nil {
			return
		}
		res := Sign(req.AppId, req.DeviceId, req.UserId, req.Nonce)
		c.JSON(http.StatusOK, res)
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
