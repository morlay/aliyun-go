package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type BusinessShowListRequest struct {
	Page int `position:"Query" name:"Page"`
	Per  int `position:"Query" name:"Per"`
}

func (r BusinessShowListRequest) Invoke(client *sdk.Client) (response *BusinessShowListResponse, err error) {
	req := struct {
		*requests.RpcRequest
		BusinessShowListRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "BusinessShowList", "", "")

	resp := struct {
		*responses.BaseResponse
		BusinessShowListResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.BusinessShowListResponse

	err = client.DoAction(&req, &resp)
	return
}

type BusinessShowListResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
