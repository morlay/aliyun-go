package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AreaShowListRequest struct {
	Page int   `position:"Query" name:"Page"`
	Per  int   `position:"Query" name:"Per"`
	Sid  int64 `position:"Query" name:"Sid"`
}

func (r AreaShowListRequest) Invoke(client *sdk.Client) (response *AreaShowListResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AreaShowListRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "AreaShowList", "", "")

	resp := struct {
		*responses.BaseResponse
		AreaShowListResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AreaShowListResponse

	err = client.DoAction(&req, &resp)
	return
}

type AreaShowListResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
