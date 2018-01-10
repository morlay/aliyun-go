package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SaveApgroupScanConfigRequest struct {
	requests.RpcRequest
	JsonData  string `position:"Query" name:"JsonData"`
	ApgroupId int64  `position:"Query" name:"ApgroupId"`
}

func (req *SaveApgroupScanConfigRequest) Invoke(client *sdk.Client) (resp *SaveApgroupScanConfigResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "SaveApgroupScanConfig", "", "")
	resp = &SaveApgroupScanConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveApgroupScanConfigResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
