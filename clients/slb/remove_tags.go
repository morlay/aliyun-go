package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RemoveTagsRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (r RemoveTagsRequest) Invoke(client *sdk.Client) (response *RemoveTagsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		RemoveTagsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Slb", "2014-05-15", "RemoveTags", "slb", "")

	resp := struct {
		*responses.BaseResponse
		RemoveTagsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.RemoveTagsResponse

	err = client.DoAction(&req, &resp)
	return
}

type RemoveTagsResponse struct {
	RequestId string
}
