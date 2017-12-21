package slb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetBackendServersRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	BackendServers       string `position:"Query" name:"BackendServers"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (r SetBackendServersRequest) Invoke(client *sdk.Client) (response *SetBackendServersResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetBackendServersRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Slb", "2014-05-15", "SetBackendServers", "slb", "")

	resp := struct {
		*responses.BaseResponse
		SetBackendServersResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetBackendServersResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetBackendServersResponse struct {
	RequestId      string
	LoadBalancerId string
	BackendServers SetBackendServersBackendServerList
}

type SetBackendServersBackendServer struct {
	ServerId string
	Weight   string
}

type SetBackendServersBackendServerList []SetBackendServersBackendServer

func (list *SetBackendServersBackendServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SetBackendServersBackendServer)
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
