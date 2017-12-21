package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteApgroupSsidConfigRequest struct {
	ApgroupId int64 `position:"Query" name:"ApgroupId"`
	Id        int64 `position:"Query" name:"Id"`
}

func (r DeleteApgroupSsidConfigRequest) Invoke(client *sdk.Client) (response *DeleteApgroupSsidConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteApgroupSsidConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "DeleteApgroupSsidConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		DeleteApgroupSsidConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteApgroupSsidConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteApgroupSsidConfigResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
