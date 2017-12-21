package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ProfileHistoryRequest struct {
	requests.RpcRequest
	Idtype     int64  `position:"Query" name:"Idtype"`
	EndMonth   string `position:"Query" name:"EndMonth"`
	BeginMonth string `position:"Query" name:"BeginMonth"`
	Agsid      int64  `position:"Query" name:"Agsid"`
}

func (req *ProfileHistoryRequest) Invoke(client *sdk.Client) (resp *ProfileHistoryResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ProfileHistory", "", "")
	resp = &ProfileHistoryResponse{}
	err = client.DoAction(req, resp)
	return
}

type ProfileHistoryResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
