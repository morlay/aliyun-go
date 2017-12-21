package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyReplicaDescriptionRequest struct {
	ReplicaDescription   string `position:"Query" name:"ReplicaDescription"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ReplicaId            string `position:"Query" name:"ReplicaId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r ModifyReplicaDescriptionRequest) Invoke(client *sdk.Client) (response *ModifyReplicaDescriptionResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyReplicaDescriptionRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "ModifyReplicaDescription", "rds", "")

	resp := struct {
		*responses.BaseResponse
		ModifyReplicaDescriptionResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyReplicaDescriptionResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyReplicaDescriptionResponse struct {
	RequestId string
}
