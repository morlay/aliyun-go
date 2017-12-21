package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListGroupApBriefConfigRequest struct {
	OrderCol   string `position:"Query" name:"OrderCol"`
	SearchName string `position:"Query" name:"SearchName"`
	ApgroupId  int64  `position:"Query" name:"ApgroupId"`
	ColCnt     int    `position:"Query" name:"ColCnt"`
	Length     int    `position:"Query" name:"Length"`
	PageIndex  int    `position:"Query" name:"PageIndex"`
	SearchMac  string `position:"Query" name:"SearchMac"`
	OrderDir   string `position:"Query" name:"OrderDir"`
}

func (r ListGroupApBriefConfigRequest) Invoke(client *sdk.Client) (response *ListGroupApBriefConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ListGroupApBriefConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ListGroupApBriefConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		ListGroupApBriefConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ListGroupApBriefConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type ListGroupApBriefConfigResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
