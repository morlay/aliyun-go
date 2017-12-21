package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteApgroupConfigRequest struct {
	Id int64 `position:"Query" name:"Id"`
}

func (r DeleteApgroupConfigRequest) Invoke(client *sdk.Client) (response *DeleteApgroupConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteApgroupConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "DeleteApgroupConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		DeleteApgroupConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteApgroupConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteApgroupConfigResponse struct {
	RequestId string
	Success   bool
	Message   string
	ErrorCode int
	ErrorMsg  string
}
