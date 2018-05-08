package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListGroupApBriefConfigRequest struct {
	requests.RpcRequest
	OrderCol   string `position:"Query" name:"OrderCol"`
	SearchName string `position:"Query" name:"SearchName"`
	ApgroupId  int64  `position:"Query" name:"ApgroupId"`
	ColCnt     int    `position:"Query" name:"ColCnt"`
	Length     int    `position:"Query" name:"Length"`
	PageIndex  int    `position:"Query" name:"PageIndex"`
	SearchMac  string `position:"Query" name:"SearchMac"`
	OrderDir   string `position:"Query" name:"OrderDir"`
}

func (req *ListGroupApBriefConfigRequest) Invoke(client *sdk.Client) (resp *ListGroupApBriefConfigResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ListGroupApBriefConfig", "", "")
	resp = &ListGroupApBriefConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListGroupApBriefConfigResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Message   common.String
	Data      common.String
	ErrorCode common.Integer
	ErrorMsg  common.String
}
