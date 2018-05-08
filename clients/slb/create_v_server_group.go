package slb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateVServerGroupRequest struct {
	requests.RpcRequest
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	BackendServers       string `position:"Query" name:"BackendServers"`
	Tags                 string `position:"Query" name:"Tags"`
	VServerGroupName     string `position:"Query" name:"VServerGroupName"`
}

func (req *CreateVServerGroupRequest) Invoke(client *sdk.Client) (resp *CreateVServerGroupResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "CreateVServerGroup", "slb", "")
	resp = &CreateVServerGroupResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateVServerGroupResponse struct {
	responses.BaseResponse
	RequestId      common.String
	VServerGroupId common.String
	BackendServers CreateVServerGroupBackendServerList
}

type CreateVServerGroupBackendServer struct {
	ServerId common.String
	Port     common.Integer
	Weight   common.Integer
	Type     common.String
	ServerIp common.String
	VpcId    common.String
}

type CreateVServerGroupBackendServerList []CreateVServerGroupBackendServer

func (list *CreateVServerGroupBackendServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateVServerGroupBackendServer)
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
