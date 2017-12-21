package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ProfileHistoryRequest struct {
	Idtype     int64  `position:"Query" name:"Idtype"`
	EndMonth   string `position:"Query" name:"EndMonth"`
	BeginMonth string `position:"Query" name:"BeginMonth"`
	Agsid      int64  `position:"Query" name:"Agsid"`
}

func (r ProfileHistoryRequest) Invoke(client *sdk.Client) (response *ProfileHistoryResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ProfileHistoryRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ProfileHistory", "", "")

	resp := struct {
		*responses.BaseResponse
		ProfileHistoryResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ProfileHistoryResponse

	err = client.DoAction(&req, &resp)
	return
}

type ProfileHistoryResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
