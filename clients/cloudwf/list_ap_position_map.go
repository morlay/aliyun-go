package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListApPositionMapRequest struct {
	requests.RpcRequest
	OrderCol          string `position:"Query" name:"OrderCol"`
	SearchName        string `position:"Query" name:"SearchName"`
	TotalItem         int    `position:"Query" name:"TotalItem"`
	Length            int    `position:"Query" name:"Length"`
	MapType           int    `position:"Query" name:"MapType"`
	PageIndex         int    `position:"Query" name:"PageIndex"`
	SearchApgroupName string `position:"Query" name:"SearchApgroupName"`
	OrderDir          string `position:"Query" name:"OrderDir"`
}

func (req *ListApPositionMapRequest) Invoke(client *sdk.Client) (resp *ListApPositionMapResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ListApPositionMap", "", "")
	resp = &ListApPositionMapResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListApPositionMapResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
