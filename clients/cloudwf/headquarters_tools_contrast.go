package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type HeadquartersToolsContrastRequest struct {
	Bid int64 `position:"Query" name:"Bid"`
}

func (r HeadquartersToolsContrastRequest) Invoke(client *sdk.Client) (response *HeadquartersToolsContrastResponse, err error) {
	req := struct {
		*requests.RpcRequest
		HeadquartersToolsContrastRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "HeadquartersToolsContrast", "", "")

	resp := struct {
		*responses.BaseResponse
		HeadquartersToolsContrastResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.HeadquartersToolsContrastResponse

	err = client.DoAction(&req, &resp)
	return
}

type HeadquartersToolsContrastResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
