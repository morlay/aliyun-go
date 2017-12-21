package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddApgroupConfigRequest struct {
	requests.RpcRequest
	ParentApgroupId int64  `position:"Query" name:"ParentApgroupId"`
	Name            string `position:"Query" name:"Name"`
	Description     string `position:"Query" name:"Description"`
}

func (req *AddApgroupConfigRequest) Invoke(client *sdk.Client) (resp *AddApgroupConfigResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "AddApgroupConfig", "", "")
	resp = &AddApgroupConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddApgroupConfigResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
