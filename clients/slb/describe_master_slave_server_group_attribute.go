package slb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeMasterSlaveServerGroupAttributeRequest struct {
	Access_key_id            string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId          int64  `position:"Query" name:"ResourceOwnerId"`
	MasterSlaveServerGroupId string `position:"Query" name:"MasterSlaveServerGroupId"`
	ResourceOwnerAccount     string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount             string `position:"Query" name:"OwnerAccount"`
	OwnerId                  int64  `position:"Query" name:"OwnerId"`
	Tags                     string `position:"Query" name:"Tags"`
}

func (r DescribeMasterSlaveServerGroupAttributeRequest) Invoke(client *sdk.Client) (response *DescribeMasterSlaveServerGroupAttributeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeMasterSlaveServerGroupAttributeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Slb", "2014-05-15", "DescribeMasterSlaveServerGroupAttribute", "slb", "")

	resp := struct {
		*responses.BaseResponse
		DescribeMasterSlaveServerGroupAttributeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeMasterSlaveServerGroupAttributeResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeMasterSlaveServerGroupAttributeResponse struct {
	RequestId                  string
	MasterSlaveServerGroupId   string
	MasterSlaveServerGroupName string
	MasterSlaveBackendServers  DescribeMasterSlaveServerGroupAttributeMasterSlaveBackendServerList
}

type DescribeMasterSlaveServerGroupAttributeMasterSlaveBackendServer struct {
	ServerId   string
	Port       int
	Weight     int
	ServerType string
}

type DescribeMasterSlaveServerGroupAttributeMasterSlaveBackendServerList []DescribeMasterSlaveServerGroupAttributeMasterSlaveBackendServer

func (list *DescribeMasterSlaveServerGroupAttributeMasterSlaveBackendServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeMasterSlaveServerGroupAttributeMasterSlaveBackendServer)
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
