package csb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteCasServiceRequest struct {
	requests.RpcRequest
	LeafOnly     string `position:"Query" name:"LeafOnly"`
	CasCsbName   string `position:"Query" name:"CasCsbName"`
	SrcUserId    string `position:"Query" name:"SrcUserId"`
	CasServiceId string `position:"Query" name:"CasServiceId"`
}

func (req *DeleteCasServiceRequest) Invoke(client *sdk.Client) (resp *DeleteCasServiceResponse, err error) {
	req.InitWithApiInfo("CSB", "2017-11-18", "DeleteCasService", "CSB", "")
	resp = &DeleteCasServiceResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteCasServiceResponse struct {
	responses.BaseResponse
	Code      common.Integer
	Message   common.String
	RequestId common.String
}
