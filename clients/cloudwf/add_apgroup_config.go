package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddApgroupConfigRequest struct {
	ParentApgroupId int64  `position:"Query" name:"ParentApgroupId"`
	Name            string `position:"Query" name:"Name"`
	Description     string `position:"Query" name:"Description"`
}

func (r AddApgroupConfigRequest) Invoke(client *sdk.Client) (response *AddApgroupConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AddApgroupConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "AddApgroupConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		AddApgroupConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AddApgroupConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type AddApgroupConfigResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
