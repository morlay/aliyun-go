package slb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyVServerGroupBackendServersRequest struct {
	requests.RpcRequest
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	VServerGroupId       string `position:"Query" name:"VServerGroupId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	OldBackendServers    string `position:"Query" name:"OldBackendServers"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	NewBackendServers    string `position:"Query" name:"NewBackendServers"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (req *ModifyVServerGroupBackendServersRequest) Invoke(client *sdk.Client) (resp *ModifyVServerGroupBackendServersResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "ModifyVServerGroupBackendServers", "slb", "")
	resp = &ModifyVServerGroupBackendServersResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyVServerGroupBackendServersResponse struct {
	responses.BaseResponse
	RequestId      string
	VServerGroupId string
	BackendServers ModifyVServerGroupBackendServersBackendServerList
}

type ModifyVServerGroupBackendServersBackendServer struct {
	ServerId string
	Port     int
	Weight   int
	Type     string
	ServerIp string
	VpcId    string
}

type ModifyVServerGroupBackendServersBackendServerList []ModifyVServerGroupBackendServersBackendServer

func (list *ModifyVServerGroupBackendServersBackendServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ModifyVServerGroupBackendServersBackendServer)
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
