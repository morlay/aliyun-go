package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type FindApRequest struct {
	Id int64 `position:"Query" name:"Id"`
}

func (r FindApRequest) Invoke(client *sdk.Client) (response *FindApResponse, err error) {
	req := struct {
		*requests.RpcRequest
		FindApRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "FindAp", "", "")

	resp := struct {
		*responses.BaseResponse
		FindApResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.FindApResponse

	err = client.DoAction(&req, &resp)
	return
}

type FindApResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
