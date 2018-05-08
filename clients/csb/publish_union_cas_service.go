package csb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type PublishUnionCasServiceRequest struct {
	requests.RpcRequest
	CasCsbName string `position:"Query" name:"CasCsbName"`
	Data       string `position:"Body" name:"Data"`
}

func (req *PublishUnionCasServiceRequest) Invoke(client *sdk.Client) (resp *PublishUnionCasServiceResponse, err error) {
	req.InitWithApiInfo("CSB", "2017-11-18", "PublishUnionCasService", "CSB", "")
	resp = &PublishUnionCasServiceResponse{}
	err = client.DoAction(req, resp)
	return
}

type PublishUnionCasServiceResponse struct {
	responses.BaseResponse
	Code      common.Integer
	Message   common.String
	RequestId common.String
}
