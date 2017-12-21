package ocs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type FlushInstanceRequest struct {
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OcsInstanceId        string `position:"Query" name:"OcsInstanceId"`
}

func (r FlushInstanceRequest) Invoke(client *sdk.Client) (response *FlushInstanceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		FlushInstanceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ocs", "2015-03-01", "FlushInstance", "", "")

	resp := struct {
		*responses.BaseResponse
		FlushInstanceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.FlushInstanceResponse

	err = client.DoAction(&req, &resp)
	return
}

type FlushInstanceResponse struct {
	RequestId string
}
