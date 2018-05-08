package csb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteServiceRequest struct {
	requests.RpcRequest
	ServiceName string `position:"Query" name:"ServiceName"`
	ServiceId   int64  `position:"Query" name:"ServiceId"`
}

func (req *DeleteServiceRequest) Invoke(client *sdk.Client) (resp *DeleteServiceResponse, err error) {
	req.InitWithApiInfo("CSB", "2017-11-18", "DeleteService", "CSB", "")
	resp = &DeleteServiceResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteServiceResponse struct {
	responses.BaseResponse
	Code      common.Integer
	Message   common.String
	RequestId common.String
}
