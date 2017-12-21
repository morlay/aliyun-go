package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SaveAccountConfigRequest struct {
	JsonData string `position:"Query" name:"JsonData"`
}

func (r SaveAccountConfigRequest) Invoke(client *sdk.Client) (response *SaveAccountConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SaveAccountConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "SaveAccountConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		SaveAccountConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SaveAccountConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type SaveAccountConfigResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
