package csb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CheckServiceExistRequest struct {
	requests.RpcRequest
	CsbId       int64  `position:"Query" name:"CsbId"`
	ServiceName string `position:"Query" name:"ServiceName"`
}

func (req *CheckServiceExistRequest) Invoke(client *sdk.Client) (resp *CheckServiceExistResponse, err error) {
	req.InitWithApiInfo("CSB", "2017-11-18", "CheckServiceExist", "CSB", "")
	resp = &CheckServiceExistResponse{}
	err = client.DoAction(req, resp)
	return
}

type CheckServiceExistResponse struct {
	responses.BaseResponse
	Code      common.Integer
	Message   common.String
	RequestId common.String
	Data      CheckServiceExistData
}

type CheckServiceExistData struct {
	Exist bool
}
