package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetApgroupDetailedConfigRequest struct {
	Id int64 `position:"Query" name:"Id"`
}

func (r GetApgroupDetailedConfigRequest) Invoke(client *sdk.Client) (response *GetApgroupDetailedConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetApgroupDetailedConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetApgroupDetailedConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		GetApgroupDetailedConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetApgroupDetailedConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetApgroupDetailedConfigResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
