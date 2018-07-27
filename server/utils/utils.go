package utils

/*Some helpers */
import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"

	"github.com/gorilla/context"
)

func ResponseNotFound(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(404)
	io.WriteString(w, "Not Found")
}
func ResponseBadRequest(w http.ResponseWriter) {
	responseWithJSON(w, 400, map[string]interface{}{"error": "Bad Request", "success": false})
}
func ResponseWithError(w http.ResponseWriter, code int, message string) {
	responseWithJSON(w, code, map[string]interface{}{"error": message, "success": false})
}
func ResponseOK(w http.ResponseWriter, returned interface{}) {

	responseWithJSON(w, 200, map[string]interface{}{"result": returned, "success": true})

}
func responseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
func Cors(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS,PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type")
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "OPTIONS" {
		return
	}
	next(w, r)
}
func ValidateMiddleware(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) {

	authorizationHeader := req.Header.Get("authorization")
	if authorizationHeader != "" {
		bearerToken := strings.Split(authorizationHeader, " ")
		if len(bearerToken) == 2 {
			token, error := jwt.ParseWithClaims(bearerToken[1], &AccessToken{}, func(token *jwt.Token) (verifykey interface{}, err error) {
				pubkey, e := decodeHeader(token.Raw)

				return pubkey, e
			})
			if error != nil {
				responseWithJSON(w, http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
				return
			}
			if token.Valid {
				context.Set(req, "decoded", token.Claims)
				next(w, req)
			} else {
				responseWithJSON(w, http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
			}
		}
	} else {
		responseWithJSON(w, http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
	}

}
func decodeHeader(token string) (*rsa.PublicKey, error) {

	/*AWS cognito send two keys and we never know what was used. So we have to check */
	kids := map[string]string{"gDYYfz/saCSIWCL7mgfjQh3pr7sXivkaDesj3X1sr/c=": "0NoiEACN-_WrgBm0FAo1PSk45W8kcjacHYhbkob3mL646govPD815SXqXPJ59UAD37KJifma6A5_NRlgs__kyjyunmehrgnFYvdfjnrAB8Wk6h29Qomlo2WlG3kYHvlBVJ2XRsi3zQiMqkR4WOnbTSdztYvn4mtN2uQOPu0ghRcIDs6xLuiDX-SVV2xSzk3a7jXa8jo9MN_UD1Mg4-9JqLfR3u-noxOofy-IHqHSuu08ahEehynBbGrux_8MVdf51to-24N3iVAOQh_K5yuNNDfzvcFQU6WSxe3c1sLwfSj2ii2YnHEM_j__ZFJBRvvb0v4EcSyawFrWEnrR-j5sqw",
		"ybDfwfeIuNtuhzo7fgqoER2prO0aUjpVRvNqzHWZY9U=": "q_jjGzfSFtSPgtjfosZmkHUu4ZkfvixBkKbDEJ70WQWnrDR03QqLzfoZ_PwFyks3JKCVqT0Tz3pqiAtIDATEaeO9zCWqjW4NpXSFYIrj-flryObx4Q3mBD3gdTrLy-cyqUod-bdh3n2GzR91UnMGleucrNlpACfTxFdYO_bkjorM0KefyQphq4NUwHYkxt3eBr5Rvv3TF0bFwD5UdJKOYfycEzdXYjODz6VnPnylR3dFdF5oVOBc1KfsHCwGMDCB7IN-0gn7O54pD-8ys2RJoTBGhpN_SDNrzMbEhPhOkHFFMdchJ5XKvuLTUNwP25BOr2FPMtZO35Tc6-FJk50WDw"}

	tokenin, err := jwt.DecodeSegment(strings.Split(token, ".")[0])

	if err != nil {
		return nil, fmt.Errorf("Invalid Header")
	}

	var dat map[string]interface{}
	if err := json.Unmarshal(tokenin, &dat); err != nil {
		return nil, fmt.Errorf("Invalid Header")
	}
	rawN := kids[dat["kid"].(string)]
	rawE := "AQAB"
	decodedE, err := base64.RawURLEncoding.DecodeString(rawE)
	if err != nil {
		return nil, fmt.Errorf("Invalid Header")
	}
	if len(decodedE) < 4 {
		ndata := make([]byte, 4)
		copy(ndata[4-len(decodedE):], decodedE)
		decodedE = ndata
	}
	pubKey := &rsa.PublicKey{
		N: &big.Int{},
		E: int(binary.BigEndian.Uint32(decodedE[:])),
	}
	decodedN, err := base64.RawURLEncoding.DecodeString(rawN)
	if err != nil {
		return nil, fmt.Errorf("Invalid Header")
	}
	pubKey.N.SetBytes(decodedN)
	return pubKey, nil
}

type AccessToken struct {
	Exp      uint32 `json:"exp"`
	TokenUse string `json:"token_use"`
	Iss      string `json:"iss"`
	ClientId string `json:"client_id"`
	Username string `json:"username"`
	Kid      string `json:"kid"`
	Alg      string `json:"alg"`
	jwt.StandardClaims
}
