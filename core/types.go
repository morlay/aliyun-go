package core

import (
	"time"
	"fmt"
	"github.com/morlay/aliyun-go/core/util"
)

var (
	SignatureVersion = "1.0"
	SignatureMethod  = "HMAC-SHA1"
)

type Request struct {
	Format               string
	Version              string
	RegionId             string
	AccessKeyId          string
	SecurityToken        string
	Signature            string
	SignatureMethod      string
	Timestamp            util.ISO6801Time
	SignatureVersion     string
	SignatureNonce       string
	ResourceOwnerAccount string
	Action               string
}

func (request *Request) init(version string, action string, AccessKeyId string, securityToken string, regionId string) {
	request.Format = "JSON"
	request.Timestamp = util.NewISO6801Time(time.Now().UTC())
	request.Version = version
	request.SignatureVersion = SignatureVersion
	request.SignatureMethod = SignatureMethod
	request.SignatureNonce = util.CreateRandomString()
	request.Action = action
	request.AccessKeyId = AccessKeyId
	request.SecurityToken = securityToken
	request.RegionId = regionId
}

type Response struct {
	RequestId string
}

type ErrorResponse struct {
	Response
	HostId  string
	Code    string
	Message string
}

type Error struct {
	ErrorResponse
	StatusCode int
}

func (e *Error) Error() string {
	return fmt.Sprintf(
		"Aliyun API Error: RequestId: %s Status Code: %d Code: %s Message: %s", e.RequestId, e.StatusCode, e.Code, e.Message,
	)
}
