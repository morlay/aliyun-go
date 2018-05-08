package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SaveApScanConfigRequest struct {
	requests.RpcRequest
	JsonData   string `position:"Query" name:"JsonData"`
	ApConfigId int64  `position:"Query" name:"ApConfigId"`
}

func (req *SaveApScanConfigRequest) Invoke(client *sdk.Client) (resp *SaveApScanConfigResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "SaveApScanConfig", "", "")
	resp = &SaveApScanConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveApScanConfigResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Message   common.String
	Data      common.String
	ErrorCode common.Integer
	ErrorMsg  common.String
}
