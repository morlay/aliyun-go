package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type HeadquartersToolsCoincideRequest struct {
	Bid int64 `position:"Query" name:"Bid"`
}

func (r HeadquartersToolsCoincideRequest) Invoke(client *sdk.Client) (response *HeadquartersToolsCoincideResponse, err error) {
	req := struct {
		*requests.RpcRequest
		HeadquartersToolsCoincideRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "HeadquartersToolsCoincide", "", "")

	resp := struct {
		*responses.BaseResponse
		HeadquartersToolsCoincideResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.HeadquartersToolsCoincideResponse

	err = client.DoAction(&req, &resp)
	return
}

type HeadquartersToolsCoincideResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
