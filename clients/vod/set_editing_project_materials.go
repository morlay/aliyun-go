package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SetEditingProjectMaterialsRequest struct {
	requests.RpcRequest
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	MaterialIds          string `position:"Query" name:"MaterialIds"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	ProjectId            string `position:"Query" name:"ProjectId"`
}

func (req *SetEditingProjectMaterialsRequest) Invoke(client *sdk.Client) (resp *SetEditingProjectMaterialsResponse, err error) {
	req.InitWithApiInfo("vod", "2017-03-21", "SetEditingProjectMaterials", "vod", "")
	resp = &SetEditingProjectMaterialsResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetEditingProjectMaterialsResponse struct {
	responses.BaseResponse
	RequestId common.String
}
