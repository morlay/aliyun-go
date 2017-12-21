package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RemoveTagsRequest struct {
	Tag4Value            string `position:"Query" name:"Tag.4.Value"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceId           string `position:"Query" name:"ResourceId"`
	Tag2Key              string `position:"Query" name:"Tag.2.Key"`
	Tag5Key              string `position:"Query" name:"Tag.5.Key"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Tag3Key              string `position:"Query" name:"Tag.3.Key"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceType         string `position:"Query" name:"ResourceType"`
	Tag5Value            string `position:"Query" name:"Tag.5.Value"`
	Tag1Key              string `position:"Query" name:"Tag.1.Key"`
	Tag1Value            string `position:"Query" name:"Tag.1.Value"`
	Tag2Value            string `position:"Query" name:"Tag.2.Value"`
	Tag4Key              string `position:"Query" name:"Tag.4.Key"`
	Tag3Value            string `position:"Query" name:"Tag.3.Value"`
}

func (r RemoveTagsRequest) Invoke(client *sdk.Client) (response *RemoveTagsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		RemoveTagsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "RemoveTags", "ecs", "")

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
