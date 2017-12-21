package slb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteMasterSlaveServerGroupRequest struct {
	Access_key_id            string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId          int64  `position:"Query" name:"ResourceOwnerId"`
	MasterSlaveServerGroupId string `position:"Query" name:"MasterSlaveServerGroupId"`
	ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount             string `position:"Query" name:"OwnerAccount"`
	OwnerId                  int64  `position:"Query" name:"OwnerId"`
	Tags                     string `position:"Query" name:"Tags"`
}

func (r DeleteMasterSlaveServerGroupRequest) Invoke(client *sdk.Client) (response *DeleteMasterSlaveServerGroupResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteMasterSlaveServerGroupRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Slb", "2014-05-15", "DeleteMasterSlaveServerGroup", "slb", "")

	resp := struct {
		*responses.BaseResponse
		DeleteMasterSlaveServerGroupResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteMasterSlaveServerGroupResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteMasterSlaveServerGroupResponse struct {
	RequestId string
}
