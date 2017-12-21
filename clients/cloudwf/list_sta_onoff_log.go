package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListStaOnoffLogRequest struct {
	requests.RpcRequest
	OrderCol       string `position:"Query" name:"OrderCol"`
	SearchSsid     string `position:"Query" name:"SearchSsid"`
	SearchApName   string `position:"Query" name:"SearchApName"`
	Length         int    `position:"Query" name:"Length"`
	SearchUsername string `position:"Query" name:"SearchUsername"`
	PageIndex      int    `position:"Query" name:"PageIndex"`
	Id             int64  `position:"Query" name:"Id"`
	OrderDir       string `position:"Query" name:"OrderDir"`
}

func (req *ListStaOnoffLogRequest) Invoke(client *sdk.Client) (resp *ListStaOnoffLogResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ListStaOnoffLog", "", "")
	resp = &ListStaOnoffLogResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListStaOnoffLogResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
