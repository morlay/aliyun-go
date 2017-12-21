package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyImageAttributeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ImageId              string `position:"Query" name:"ImageId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ImageName            string `position:"Query" name:"ImageName"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r ModifyImageAttributeRequest) Invoke(client *sdk.Client) (response *ModifyImageAttributeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyImageAttributeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "ModifyImageAttribute", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		ModifyImageAttributeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyImageAttributeResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyImageAttributeResponse struct {
	RequestId string
}
