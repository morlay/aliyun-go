package iot

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryApplyStatusRequest struct {
	ApplyId int64 `position:"Query" name:"ApplyId"`
}

func (r QueryApplyStatusRequest) Invoke(client *sdk.Client) (response *QueryApplyStatusResponse, err error) {
	req := struct {
		*requests.RpcRequest
		QueryApplyStatusRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Iot", "2017-04-20", "QueryApplyStatus", "", "")

	resp := struct {
		*responses.BaseResponse
		QueryApplyStatusResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.QueryApplyStatusResponse

	err = client.DoAction(&req, &resp)
	return
}

type QueryApplyStatusResponse struct {
	RequestId    string
	Success      bool
	ErrorMessage string
	Finish       bool
}
