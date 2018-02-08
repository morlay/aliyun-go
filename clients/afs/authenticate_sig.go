package afs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AuthenticateSigRequest struct {
	requests.RpcRequest
	Sig             string `position:"Query" name:"Sig"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	RemoteIp        string `position:"Query" name:"RemoteIp"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	AppKey          string `position:"Query" name:"AppKey"`
	SessionId       string `position:"Query" name:"SessionId"`
	Token           string `position:"Query" name:"Token"`
	Scene           string `position:"Query" name:"Scene"`
}

func (req *AuthenticateSigRequest) Invoke(client *sdk.Client) (resp *AuthenticateSigResponse, err error) {
	req.InitWithApiInfo("afs", "2018-01-12", "AuthenticateSig", "", "")
	resp = &AuthenticateSigResponse{}
	err = client.DoAction(req, resp)
	return
}

type AuthenticateSigResponse struct {
	responses.BaseResponse
	RequestId string
	Code      int
	Msg       string
	RiskLevel string
	Detail    string
}
