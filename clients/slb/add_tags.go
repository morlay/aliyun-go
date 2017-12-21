package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddTagsRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (r AddTagsRequest) Invoke(client *sdk.Client) (response *AddTagsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AddTagsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Slb", "2014-05-15", "AddTags", "slb", "")

	resp := struct {
		*responses.BaseResponse
		AddTagsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AddTagsResponse

	err = client.DoAction(&req, &resp)
	return
}

type AddTagsResponse struct {
	RequestId string
}
