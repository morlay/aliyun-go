package slb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddVServerGroupBackendServersRequest struct {
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	VServerGroupId       string `position:"Query" name:"VServerGroupId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	BackendServers       string `position:"Query" name:"BackendServers"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (r AddVServerGroupBackendServersRequest) Invoke(client *sdk.Client) (response *AddVServerGroupBackendServersResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AddVServerGroupBackendServersRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Slb", "2014-05-15", "AddVServerGroupBackendServers", "slb", "")

	resp := struct {
		*responses.BaseResponse
		AddVServerGroupBackendServersResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AddVServerGroupBackendServersResponse

	err = client.DoAction(&req, &resp)
	return
}

type AddVServerGroupBackendServersResponse struct {
	RequestId      string
	VServerGroupId string
	BackendServers AddVServerGroupBackendServersBackendServerList
}

type AddVServerGroupBackendServersBackendServer struct {
	ServerId string
	Port     int
	Weight   int
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
