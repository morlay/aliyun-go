package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId common.String
	Success   bool
	Message   common.String
	Data      common.String
	ErrorCode common.Integer
	ErrorMsg  common.String
}
