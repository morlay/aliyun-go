package slb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RemoveVServerGroupBackendServersRequest struct {
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

func (req *RemoveVServerGroupBackendServersRequest) Invoke(client *sdk.Client) (resp *RemoveVServerGroupBackendServersResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "RemoveVServerGroupBackendServers", "slb", "")
	resp = &RemoveVServerGroupBackendServersResponse{}
	err = client.DoAction(req, resp)
	return
}

type RemoveVServerGroupBackendServersResponse struct {
	responses.BaseResponse
	RequestId      string
	VServerGroupId string
	BackendServers RemoveVServerGroupBackendServersBackendServerList
}

type RemoveVServerGroupBackendServersBackendServer struct {
	ServerId string
	Port     int
	Weight   int
}

type RemoveVServerGroupBackendServersBackendServerList []RemoveVServerGroupBackendServersBackendServer

func (list *RemoveVServerGroupBackendServersBackendServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]RemoveVServerGroupBackendServersBackendServer)
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
