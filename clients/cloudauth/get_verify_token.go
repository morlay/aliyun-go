package cloudauth

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetVerifyTokenRequest struct {
	requests.RpcRequest
	UserData        string `position:"Query" name:"UserData"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Biz             string `position:"Query" name:"Biz"`
	Binding         string `position:"Query" name:"Binding"`
	TicketId        string `position:"Query" name:"TicketId"`
}

func (req *GetVerifyTokenRequest) Invoke(client *sdk.Client) (resp *GetVerifyTokenResponse, err error) {
	req.InitWithApiInfo("Cloudauth", "2017-11-17", "GetVerifyToken", "cloudauth", "")
	resp = &GetVerifyTokenResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetVerifyTokenResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Code      string
	Message   string
	Data      GetVerifyTokenData
}

type GetVerifyTokenData struct {
	VerifyToken GetVerifyTokenVerifyToken
	StsToken    GetVerifyTokenStsToken
}

type GetVerifyTokenVerifyToken struct {
	Token           string
	DurationSeconds int
}

type GetVerifyTokenStsToken struct {
	AccessKeyId     string
	AccessKeySecret string
	Expiration      string
	EndPoint        string
	BucketName      string
	Path            string
	Token           string
}
