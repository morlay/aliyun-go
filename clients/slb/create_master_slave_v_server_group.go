package slb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateMasterSlaveVServerGroupRequest struct {
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

func (r CreateMasterSlaveVServerGroupRequest) Invoke(client *sdk.Client) (response *CreateMasterSlaveVServerGroupResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateMasterSlaveVServerGroupRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Slb", "2014-05-15", "CreateMasterSlaveVServerGroup", "slb", "")

	resp := struct {
		*responses.BaseResponse
		CreateMasterSlaveVServerGroupResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateMasterSlaveVServerGroupResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateMasterSlaveVServerGroupResponse struct {
	RequestId                 string
	MasterSlaveVServerGroupId string
	MasterSlaveBackendServers CreateMasterSlaveVServerGroupMasterSlaveBackendServerList
}

type CreateMasterSlaveVServerGroupMasterSlaveBackendServer struct {
	ServerId string
	Port     int
	Weight   int
	IsBackup int
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
