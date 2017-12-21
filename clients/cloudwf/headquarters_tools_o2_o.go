package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type HeadquartersToolsO2ORequest struct {
	Bid int64 `position:"Query" name:"Bid"`
}

func (r HeadquartersToolsO2ORequest) Invoke(client *sdk.Client) (response *HeadquartersToolsO2OResponse, err error) {
	req := struct {
		*requests.RpcRequest
		HeadquartersToolsO2ORequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "HeadquartersToolsO2O", "", "")

	resp := struct {
		*responses.BaseResponse
		HeadquartersToolsO2OResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.HeadquartersToolsO2OResponse

	err = client.DoAction(&req, &resp)
	return
}

type HeadquartersToolsO2OResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
