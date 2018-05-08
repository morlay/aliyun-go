package slb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type AddVServerGroupBackendServersRequest struct {
	requests.RpcRequest
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	VServerGroupId       string `position:"Query" name:"VServerGroupId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	BackendServers       string `position:"Query" name:"BackendServers"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (req *AddVServerGroupBackendServersRequest) Invoke(client *sdk.Client) (resp *AddVServerGroupBackendServersResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "AddVServerGroupBackendServers", "slb", "")
	resp = &AddVServerGroupBackendServersResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddVServerGroupBackendServersResponse struct {
	responses.BaseResponse
	RequestId      common.String
	VServerGroupId common.String
	BackendServers AddVServerGroupBackendServersBackendServerList
}

type AddVServerGroupBackendServersBackendServer struct {
	ServerId common.String
	Port     common.Integer
	Weight   common.Integer
	Type     common.String
	ServerIp common.String
	VpcId    common.String
}

type AddVServerGroupBackendServersBackendServerList []AddVServerGroupBackendServersBackendServer

func (list *AddVServerGroupBackendServersBackendServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]AddVServerGroupBackendServersBackendServer)
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
