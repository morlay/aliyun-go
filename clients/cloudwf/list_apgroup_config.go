package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListApgroupConfigRequest struct {
	requests.RpcRequest
	OrderCol      string `position:"Query" name:"OrderCol"`
	SearchName    string `position:"Query" name:"SearchName"`
	SearchCompany string `position:"Query" name:"SearchCompany"`
	Length        int    `position:"Query" name:"Length"`
	PageIndex     int    `position:"Query" name:"PageIndex"`
	OrderDir      string `position:"Query" name:"OrderDir"`
}

func (req *ListApgroupConfigRequest) Invoke(client *sdk.Client) (resp *ListApgroupConfigResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ListApgroupConfig", "", "")
	resp = &ListApgroupConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListApgroupConfigResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Message   common.String
	Data      common.String
	ErrorCode common.Integer
	ErrorMsg  common.String
}
