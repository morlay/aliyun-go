package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListStaOnoffLogRequest struct {
	OrderCol       string `position:"Query" name:"OrderCol"`
	SearchSsid     string `position:"Query" name:"SearchSsid"`
	SearchApName   string `position:"Query" name:"SearchApName"`
	Length         int    `position:"Query" name:"Length"`
	SearchUsername string `position:"Query" name:"SearchUsername"`
	PageIndex      int    `position:"Query" name:"PageIndex"`
	Id             int64  `position:"Query" name:"Id"`
	OrderDir       string `position:"Query" name:"OrderDir"`
}

func (r ListStaOnoffLogRequest) Invoke(client *sdk.Client) (response *ListStaOnoffLogResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ListStaOnoffLogRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ListStaOnoffLog", "", "")

	resp := struct {
		*responses.BaseResponse
		ListStaOnoffLogResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ListStaOnoffLogResponse

	err = client.DoAction(&req, &resp)
	return
}

type ListStaOnoffLogResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
