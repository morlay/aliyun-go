package csb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteUnionCasServiceRequest struct {
	requests.RpcRequest
	LeafOnly     string `position:"Query" name:"LeafOnly"`
	CasCsbName   string `position:"Query" name:"CasCsbName"`
	SrcUserId    string `position:"Query" name:"SrcUserId"`
	CasServiceId string `position:"Query" name:"CasServiceId"`
}

func (req *DeleteUnionCasServiceRequest) Invoke(client *sdk.Client) (resp *DeleteUnionCasServiceResponse, err error) {
	req.InitWithApiInfo("CSB", "2017-11-18", "DeleteUnionCasService", "CSB", "")
	resp = &DeleteUnionCasServiceResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteUnionCasServiceResponse struct {
	responses.BaseResponse
	Code      common.Integer
	Message   common.String
	RequestId common.String
}
