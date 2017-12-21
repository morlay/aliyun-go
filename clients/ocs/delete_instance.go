package ocs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteInstanceRequest struct {
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OcsInstanceId        string `position:"Query" name:"OcsInstanceId"`
}

func (r DeleteInstanceRequest) Invoke(client *sdk.Client) (response *DeleteInstanceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteInstanceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ocs", "2015-03-01", "DeleteInstance", "", "")

	resp := struct {
		*responses.BaseResponse
		DeleteInstanceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteInstanceResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteInstanceResponse struct {
	RequestId string
}
