package cloudauth

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetVerifyTokenRequest struct {
	UserData        string `position:"Query" name:"UserData"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Biz             string `position:"Query" name:"Biz"`
	Binding         string `position:"Query" name:"Binding"`
	TicketId        string `position:"Query" name:"TicketId"`
}

func (r GetVerifyTokenRequest) Invoke(client *sdk.Client) (response *GetVerifyTokenResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetVerifyTokenRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cloudauth", "2017-11-17", "GetVerifyToken", "cloudauth", "")

	resp := struct {
		*responses.BaseResponse
		GetVerifyTokenResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetVerifyTokenResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetVerifyTokenResponse struct {
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
