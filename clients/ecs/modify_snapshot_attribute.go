package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifySnapshotAttributeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SnapshotId           string `position:"Query" name:"SnapshotId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Description          string `position:"Query" name:"Description"`
	SnapshotName         string `position:"Query" name:"SnapshotName"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r ModifySnapshotAttributeRequest) Invoke(client *sdk.Client) (response *ModifySnapshotAttributeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifySnapshotAttributeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "ModifySnapshotAttribute", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		ModifySnapshotAttributeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifySnapshotAttributeResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifySnapshotAttributeResponse struct {
	RequestId string
}
