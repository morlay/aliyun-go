package bss

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetResourceBusinessStatusRequest struct {
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceType         string `position:"Query" name:"ResourceType"`
	ResourceId           string `position:"Query" name:"ResourceId"`
	BusinessStatus       string `position:"Query" name:"BusinessStatus"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
}

func (r SetResourceBusinessStatusRequest) Invoke(client *sdk.Client) (response *SetResourceBusinessStatusResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetResourceBusinessStatusRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Bss", "2014-07-14", "SetResourceBusinessStatus", "", "")

	resp := struct {
		*responses.BaseResponse
		SetResourceBusinessStatusResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetResourceBusinessStatusResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetResourceBusinessStatusResponse struct {
	RequestId string
}
