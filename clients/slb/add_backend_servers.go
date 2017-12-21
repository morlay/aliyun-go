package slb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddBackendServersRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	BackendServers       string `position:"Query" name:"BackendServers"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (r AddBackendServersRequest) Invoke(client *sdk.Client) (response *AddBackendServersResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AddBackendServersRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Slb", "2014-05-15", "AddBackendServers", "slb", "")

	resp := struct {
		*responses.BaseResponse
		AddBackendServersResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AddBackendServersResponse

	err = client.DoAction(&req, &resp)
	return
}

type AddBackendServersResponse struct {
	RequestId      string
	LoadBalancerId string
	BackendServers AddBackendServersBackendServerList
}

type AddBackendServersBackendServer struct {
	ServerId string
	Weight   string
}

type AddBackendServersBackendServerList []AddBackendServersBackendServer

func (list *AddBackendServersBackendServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]AddBackendServersBackendServer)
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
