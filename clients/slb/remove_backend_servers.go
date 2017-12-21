package slb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RemoveBackendServersRequest struct {
	requests.RpcRequest
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	BackendServers       string `position:"Query" name:"BackendServers"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (req *RemoveBackendServersRequest) Invoke(client *sdk.Client) (resp *RemoveBackendServersResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "RemoveBackendServers", "slb", "")
	resp = &RemoveBackendServersResponse{}
	err = client.DoAction(req, resp)
	return
}

type RemoveBackendServersResponse struct {
	responses.BaseResponse
	RequestId      string
	LoadBalancerId string
	BackendServers RemoveBackendServersBackendServerList
}

type RemoveBackendServersBackendServer struct {
	ServerId string
	Weight   int
}

type RemoveBackendServersBackendServerList []RemoveBackendServersBackendServer

func (list *RemoveBackendServersBackendServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]RemoveBackendServersBackendServer)
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
