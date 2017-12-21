package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GroupIntimeRequest struct {
	Gsid int64 `position:"Query" name:"Gsid"`
}

func (r GroupIntimeRequest) Invoke(client *sdk.Client) (response *GroupIntimeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GroupIntimeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GroupIntime", "", "")

	resp := struct {
		*responses.BaseResponse
		GroupIntimeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GroupIntimeResponse

	err = client.DoAction(&req, &resp)
	return
}

type GroupIntimeResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
