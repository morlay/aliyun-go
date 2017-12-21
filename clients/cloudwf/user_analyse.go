package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UserAnalyseRequest struct {
	Gsid int64 `position:"Query" name:"Gsid"`
}

func (r UserAnalyseRequest) Invoke(client *sdk.Client) (response *UserAnalyseResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UserAnalyseRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "UserAnalyse", "", "")

	resp := struct {
		*responses.BaseResponse
		UserAnalyseResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UserAnalyseResponse

	err = client.DoAction(&req, &resp)
	return
}

type UserAnalyseResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
