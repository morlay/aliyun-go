package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ProfileHistoryListRequest struct {
	Idtype int64 `position:"Query" name:"Idtype"`
	Page   int   `position:"Query" name:"Page"`
	Per    int   `position:"Query" name:"Per"`
	Agsid  int64 `position:"Query" name:"Agsid"`
}

func (r ProfileHistoryListRequest) Invoke(client *sdk.Client) (response *ProfileHistoryListResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ProfileHistoryListRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ProfileHistoryList", "", "")

	resp := struct {
		*responses.BaseResponse
		ProfileHistoryListResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ProfileHistoryListResponse

	err = client.DoAction(&req, &resp)
	return
}

type ProfileHistoryListResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
