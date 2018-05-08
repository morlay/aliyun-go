package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteApRadioSsidConfigRequest struct {
	requests.RpcRequest
	InstantlyEffective int   `position:"Query" name:"InstantlyEffective"`
	Id                 int64 `position:"Query" name:"Id"`
}

func (req *DeleteApRadioSsidConfigRequest) Invoke(client *sdk.Client) (resp *DeleteApRadioSsidConfigResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "DeleteApRadioSsidConfig", "", "")
	resp = &DeleteApRadioSsidConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteApRadioSsidConfigResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Message   common.String
	Data      common.String
	ErrorCode common.Integer
	ErrorMsg  common.String
}
