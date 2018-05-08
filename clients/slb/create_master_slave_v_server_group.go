package slb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateMasterSlaveVServerGroupRequest struct {
	requests.RpcRequest
	Access_key_id               string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId             int64  `position:"Query" name:"ResourceOwnerId"`
	MasterSlaveBackendServers   string `position:"Query" name:"MasterSlaveBackendServers"`
	LoadBalancerId              string `position:"Query" name:"LoadBalancerId"`
	ResourceOwnerAccount        string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                string `position:"Query" name:"OwnerAccount"`
	MasterSlaveVServerGroupName string `position:"Query" name:"MasterSlaveVServerGroupName"`
	OwnerId                     int64  `position:"Query" name:"OwnerId"`
	Tags                        string `position:"Query" name:"Tags"`
}

func (req *CreateMasterSlaveVServerGroupRequest) Invoke(client *sdk.Client) (resp *CreateMasterSlaveVServerGroupResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "CreateMasterSlaveVServerGroup", "slb", "")
	resp = &CreateMasterSlaveVServerGroupResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateMasterSlaveVServerGroupResponse struct {
	responses.BaseResponse
	RequestId                 common.String
	MasterSlaveVServerGroupId common.String
	MasterSlaveBackendServers CreateMasterSlaveVServerGroupMasterSlaveBackendServerList
}

type CreateMasterSlaveVServerGroupMasterSlaveBackendServer struct {
	ServerId common.String
	Port     common.Integer
	Weight   common.Integer
	IsBackup common.Integer
	Type     common.String
	ServerIp common.String
	VpcId    common.String
}

type CreateMasterSlaveVServerGroupMasterSlaveBackendServerList []CreateMasterSlaveVServerGroupMasterSlaveBackendServer

func (list *CreateMasterSlaveVServerGroupMasterSlaveBackendServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateMasterSlaveVServerGroupMasterSlaveBackendServer)
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
