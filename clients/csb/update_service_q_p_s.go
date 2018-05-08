package csb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type UpdateServiceQPSRequest struct {
	requests.RpcRequest
	Qps       string `position:"Query" name:"Qps"`
	ServiceId int64  `position:"Query" name:"ServiceId"`
}

func (req *UpdateServiceQPSRequest) Invoke(client *sdk.Client) (resp *UpdateServiceQPSResponse, err error) {
	req.InitWithApiInfo("CSB", "2017-11-18", "UpdateServiceQPS", "CSB", "")
	resp = &UpdateServiceQPSResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateServiceQPSResponse struct {
	responses.BaseResponse
	Code      common.Integer
	Message   common.String
	RequestId common.String
}
