package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RepairApRequest struct {
	Id int64 `position:"Query" name:"Id"`
}

func (r RepairApRequest) Invoke(client *sdk.Client) (response *RepairApResponse, err error) {
	req := struct {
		*requests.RpcRequest
		RepairApRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "RepairAp", "", "")

	resp := struct {
		*responses.BaseResponse
		RepairApResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.RepairApResponse

	err = client.DoAction(&req, &resp)
	return
}

type RepairApResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
