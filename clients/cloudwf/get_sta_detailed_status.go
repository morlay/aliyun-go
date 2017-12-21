package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetStaDetailedStatusRequest struct {
	Id int64 `position:"Query" name:"Id"`
}

func (r GetStaDetailedStatusRequest) Invoke(client *sdk.Client) (response *GetStaDetailedStatusResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetStaDetailedStatusRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetStaDetailedStatus", "", "")

	resp := struct {
		*responses.BaseResponse
		GetStaDetailedStatusResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetStaDetailedStatusResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetStaDetailedStatusResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
