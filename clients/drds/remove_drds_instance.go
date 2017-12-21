package drds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RemoveDrdsInstanceRequest struct {
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
}

func (r RemoveDrdsInstanceRequest) Invoke(client *sdk.Client) (response *RemoveDrdsInstanceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		RemoveDrdsInstanceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Drds", "2017-10-16", "RemoveDrdsInstance", "", "")

	resp := struct {
		*responses.BaseResponse
		RemoveDrdsInstanceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.RemoveDrdsInstanceResponse

	err = client.DoAction(&req, &resp)
	return
}

type RemoveDrdsInstanceResponse struct {
	RequestId string
	Success   bool
}
