package slb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SetVServerGroupAttributeRequest struct {
	requests.RpcRequest
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	VServerGroupId       string `position:"Query" name:"VServerGroupId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	BackendServers       string `position:"Query" name:"BackendServers"`
	Tags                 string `position:"Query" name:"Tags"`
	VServerGroupName     string `position:"Query" name:"VServerGroupName"`
}

func (req *SetVServerGroupAttributeRequest) Invoke(client *sdk.Client) (resp *SetVServerGroupAttributeResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "SetVServerGroupAttribute", "slb", "")
	resp = &SetVServerGroupAttributeResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetVServerGroupAttributeResponse struct {
	responses.BaseResponse
	RequestId        common.String
	VServerGroupId   common.String
	VServerGroupName common.String
	BackendServers   SetVServerGroupAttributeBackendServerList
}

type SetVServerGroupAttributeBackendServer struct {
	ServerId common.String
	Port     common.Integer
	Weight   common.Integer
	Type     common.String
	ServerIp common.String
	VpcId    common.String
}

type SetVServerGroupAttributeBackendServerList []SetVServerGroupAttributeBackendServer

func (list *SetVServerGroupAttributeBackendServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SetVServerGroupAttributeBackendServer)
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
