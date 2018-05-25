package cloudauth

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetVerifyTokenRequest struct {
	requests.RpcRequest
	UserData        string `position:"Query" name:"UserData"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Biz             string `position:"Query" name:"Biz"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	Binding         string `position:"Query" name:"Binding"`
	TicketId        string `position:"Query" name:"TicketId"`
}

func (req *GetVerifyTokenRequest) Invoke(client *sdk.Client) (resp *GetVerifyTokenResponse, err error) {
	req.InitWithApiInfo("Cloudauth", "2018-05-04", "GetVerifyToken", "cloudauth", "")
	resp = &GetVerifyTokenResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetVerifyTokenResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Code      common.String
	Message   common.String
	Data      GetVerifyTokenData
}

type GetVerifyTokenData struct {
	VerifyToken GetVerifyTokenVerifyToken
	StsToken    GetVerifyTokenStsToken
}

type GetVerifyTokenVerifyToken struct {
	Token           common.String
	DurationSeconds common.Integer
}

type GetVerifyTokenStsToken struct {
	AccessKeyId     common.String
	AccessKeySecret common.String
	Expiration      common.String
	EndPoint        common.String
	BucketName      common.String
	Path            common.String
	Token           common.String
}
