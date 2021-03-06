package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId common.String
	Success   bool
	Message   common.String
	Data      common.String
	ErrorCode common.Integer
	ErrorMsg  common.String
}
