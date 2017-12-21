package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListApgroupConfigRequest struct {
	OrderCol      string `position:"Query" name:"OrderCol"`
	SearchName    string `position:"Query" name:"SearchName"`
	SearchCompany string `position:"Query" name:"SearchCompany"`
	Length        int    `position:"Query" name:"Length"`
	PageIndex     int    `position:"Query" name:"PageIndex"`
	OrderDir      string `position:"Query" name:"OrderDir"`
}

func (r ListApgroupConfigRequest) Invoke(client *sdk.Client) (response *ListApgroupConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ListApgroupConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ListApgroupConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		ListApgroupConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ListApgroupConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type ListApgroupConfigResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
