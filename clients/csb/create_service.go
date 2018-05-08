package csb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateServiceRequest struct {
	requests.RpcRequest
	Data  string `position:"Body" name:"Data"`
	CsbId int64  `position:"Query" name:"CsbId"`
}

func (req *CreateServiceRequest) Invoke(client *sdk.Client) (resp *CreateServiceResponse, err error) {
	req.InitWithApiInfo("CSB", "2017-11-18", "CreateService", "CSB", "")
	resp = &CreateServiceResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateServiceResponse struct {
	responses.BaseResponse
	Code      common.Integer
	Message   common.String
	RequestId common.String
	Data      CreateServiceData
}

type CreateServiceData struct {
	Id common.Long
}
