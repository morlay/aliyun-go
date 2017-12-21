package dds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ReleaseReplicaRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ReplicaId            string `position:"Query" name:"ReplicaId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r ReleaseReplicaRequest) Invoke(client *sdk.Client) (response *ReleaseReplicaResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ReleaseReplicaRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Dds", "2015-12-01", "ReleaseReplica", "dds", "")

	resp := struct {
		*responses.BaseResponse
		ReleaseReplicaResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ReleaseReplicaResponse

	err = client.DoAction(&req, &resp)
	return
}

type ReleaseReplicaResponse struct {
	RequestId string
}
