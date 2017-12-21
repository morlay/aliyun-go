package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RepairGroupApRequest struct {
	Id int64 `position:"Query" name:"Id"`
}

func (r RepairGroupApRequest) Invoke(client *sdk.Client) (response *RepairGroupApResponse, err error) {
	req := struct {
		*requests.RpcRequest
		RepairGroupApRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "RepairGroupAp", "", "")

	resp := struct {
		*responses.BaseResponse
		RepairGroupApResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.RepairGroupApResponse

	err = client.DoAction(&req, &resp)
	return
}

type RepairGroupApResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
