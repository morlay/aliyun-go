package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeletePositionMapRequest struct {
	MapId int64 `position:"Query" name:"MapId"`
}

func (r DeletePositionMapRequest) Invoke(client *sdk.Client) (response *DeletePositionMapResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeletePositionMapRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "DeletePositionMap", "", "")

	resp := struct {
		*responses.BaseResponse
		DeletePositionMapResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeletePositionMapResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeletePositionMapResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
