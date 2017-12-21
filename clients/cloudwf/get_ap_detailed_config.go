package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetApDetailedConfigRequest struct {
	Id int64 `position:"Query" name:"Id"`
}

func (r GetApDetailedConfigRequest) Invoke(client *sdk.Client) (response *GetApDetailedConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetApDetailedConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetApDetailedConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		GetApDetailedConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetApDetailedConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetApDetailedConfigResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
