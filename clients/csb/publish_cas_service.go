package csb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type PublishCasServiceRequest struct {
	requests.RpcRequest
	CasCsbName string `position:"Query" name:"CasCsbName"`
	Data       string `position:"Body" name:"Data"`
}

func (req *PublishCasServiceRequest) Invoke(client *sdk.Client) (resp *PublishCasServiceResponse, err error) {
	req.InitWithApiInfo("CSB", "2017-11-18", "PublishCasService", "CSB", "")
	resp = &PublishCasServiceResponse{}
	err = client.DoAction(req, resp)
	return
}

type PublishCasServiceResponse struct {
	responses.BaseResponse
	Code      common.Integer
	Message   common.String
	RequestId common.String
}
