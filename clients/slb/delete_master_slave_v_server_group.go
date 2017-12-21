package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteMasterSlaveVServerGroupRequest struct {
	Access_key_id             string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId           int64  `position:"Query" name:"ResourceOwnerId"`
	MasterSlaveVServerGroupId string `position:"Query" name:"MasterSlaveVServerGroupId"`
	ResourceOwnerAccount      string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount              string `position:"Query" name:"OwnerAccount"`
	OwnerId                   int64  `position:"Query" name:"OwnerId"`
	Tags                      string `position:"Query" name:"Tags"`
}

func (r DeleteMasterSlaveVServerGroupRequest) Invoke(client *sdk.Client) (response *DeleteMasterSlaveVServerGroupResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteMasterSlaveVServerGroupRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Slb", "2014-05-15", "DeleteMasterSlaveVServerGroup", "slb", "")

	resp := struct {
		*responses.BaseResponse
		DeleteMasterSlaveVServerGroupResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteMasterSlaveVServerGroupResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteMasterSlaveVServerGroupResponse struct {
	RequestId string
}
