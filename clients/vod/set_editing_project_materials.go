package vod

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetEditingProjectMaterialsRequest struct {
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	MaterialIds          string `position:"Query" name:"MaterialIds"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	ProjectId            string `position:"Query" name:"ProjectId"`
}

func (r SetEditingProjectMaterialsRequest) Invoke(client *sdk.Client) (response *SetEditingProjectMaterialsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetEditingProjectMaterialsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("vod", "2017-03-21", "SetEditingProjectMaterials", "", "")

	resp := struct {
		*responses.BaseResponse
		SetEditingProjectMaterialsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetEditingProjectMaterialsResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetEditingProjectMaterialsResponse struct {
	RequestId string
}
