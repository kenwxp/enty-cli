// Copyright 2017 Vector Creations Ltd
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package auth implements authentication checks and storage.
package auth

import (
	"crypto/rand"
	"encoding/base64"
	"entysquare/filer-backend/jsonerror"
	"entysquare/filer-backend/storage"
	"entysquare/filer-backend/storage/types"
	"entysquare/filer-backend/util"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// VerifyUserFromRequest authenticates the HTTP request,
// on success returns Device of the requester.
// Finds local user or an application service user.
// Note: For an AS user, AS dummy device is returned.
// On failure returns an JSON error response which can be sent to the client.
func VerifyUserFromRequest(
	req *http.Request, accountDB *storage.Database,
) (*types.Account, *util.JSONResponse) {
	// Try to find the Application Service user
	token, err := ExtractAccessToken(req)
	fmt.Println("find token: ", token, err)
	if err != nil {
		return nil, &util.JSONResponse{
			Code: http.StatusUnauthorized,
			JSON: jsonerror.MissingToken(err.Error()),
		}
	}
	userPre, err := accountDB.SelectAccountByToken(req.Context(), nil, token)
	if err != nil || userPre == nil {
		return nil, &util.JSONResponse{
			Code: http.StatusUnauthorized,
			JSON: jsonerror.MissingToken("missing token"),
		}
	}
	mid := strings.Split(userPre.Token, ":")
	tokenPre := mid[0]
	timeStamp := mid[1]
	preTs, err := strconv.ParseInt(timeStamp, 10, 64)
	nowTs := time.Now().Unix()

	// if tokenTs is exacted 1 week it's broken
	diff := nowTs - preTs
	unify := 1 * 7 * 24 * 60 * 60 * time.Second
	if time.Duration(diff*1000*1000*1000) > unify {
		return nil, &util.JSONResponse{
			Code: http.StatusUnauthorized,
			JSON: jsonerror.MissingToken("token is out of date"),
		}
	}
	nowToken := tokenPre + ":" + strconv.FormatInt(nowTs, 10)
	userPre.Token = nowToken
	err = accountDB.UpdateAccountToken(req.Context(), nil, userPre.FilerID, nowToken)
	if err != nil {
		return nil, &util.JSONResponse{
			Code: http.StatusUnauthorized,
			JSON: jsonerror.MissingToken(err.Error()),
		}
	}

	return userPre, nil
}

// GenerateAccessToken creates a new access token. Returns an error if failed to generate
// random bytes.
func GenerateAccessToken() (string, error) {
	b := make([]byte, 10)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	// url-safe no padding
	return base64.RawURLEncoding.EncodeToString(b), nil
}

// ExtractAccessToken from a request, or return an error detailing what went wrong. The
// error message MUST be human-readable and comprehensible to the client.
func ExtractAccessToken(req *http.Request) (string, error) {
	fmt.Println("headers are ", req.Header)
	queryToken := req.Header.Get("access_token")
	fmt.Print("access_token is ", queryToken)
	if queryToken != "" && queryToken != "null" {
		return queryToken, nil
	}
	return "", fmt.Errorf("missing access token")
}
