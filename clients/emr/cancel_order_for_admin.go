package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CancelOrderForAdminRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	UserId          string `position:"Query" name:"UserId"`
}

func (req *CancelOrderForAdminRequest) Invoke(client *sdk.Client) (resp *CancelOrderForAdminResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "CancelOrderForAdmin", "", "")
	resp = &CancelOrderForAdminResponse{}
	err = client.DoAction(req, resp)
	return
}

type CancelOrderForAdminResponse struct {
	responses.BaseResponse
	RequestId string
	ClusterId string
}
