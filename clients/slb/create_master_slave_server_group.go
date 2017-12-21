package slb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateMasterSlaveServerGroupRequest struct {
	requests.RpcRequest
	Access_key_id              string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId            int64  `position:"Query" name:"ResourceOwnerId"`
	MasterSlaveBackendServers  string `position:"Query" name:"MasterSlaveBackendServers"`
	LoadBalancerId             string `position:"Query" name:"LoadBalancerId"`
	ResourceOwnerAccount       string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount               string `position:"Query" name:"OwnerAccount"`
	MasterSlaveServerGroupName string `position:"Query" name:"MasterSlaveServerGroupName"`
	OwnerId                    int64  `position:"Query" name:"OwnerId"`
	Tags                       string `position:"Query" name:"Tags"`
}

func (req *CreateMasterSlaveServerGroupRequest) Invoke(client *sdk.Client) (resp *CreateMasterSlaveServerGroupResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "CreateMasterSlaveServerGroup", "slb", "")
	resp = &CreateMasterSlaveServerGroupResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateMasterSlaveServerGroupResponse struct {
	responses.BaseResponse
	RequestId                 string
	MasterSlaveServerGroupId  string
	MasterSlaveBackendServers CreateMasterSlaveServerGroupMasterSlaveBackendServerList
}

type CreateMasterSlaveServerGroupMasterSlaveBackendServer struct {
	ServerId   string
	Port       int
	Weight     int
	ServerType string
}

type CreateMasterSlaveServerGroupMasterSlaveBackendServerList []CreateMasterSlaveServerGroupMasterSlaveBackendServer

func (list *CreateMasterSlaveServerGroupMasterSlaveBackendServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateMasterSlaveServerGroupMasterSlaveBackendServer)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
