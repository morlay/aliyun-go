package slb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeHealthStatusRequest struct {
	requests.RpcRequest
	Access_key_id        string `position:"Query" name:"Access_key_id"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ListenerPort         int    `position:"Query" name:"ListenerPort"`
	LoadBalancerId       string `position:"Query" name:"LoadBalancerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Tags                 string `position:"Query" name:"Tags"`
}

func (req *DescribeHealthStatusRequest) Invoke(client *sdk.Client) (resp *DescribeHealthStatusResponse, err error) {
	req.InitWithApiInfo("Slb", "2014-05-15", "DescribeHealthStatus", "slb", "")
	resp = &DescribeHealthStatusResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeHealthStatusResponse struct {
	responses.BaseResponse
	RequestId      common.String
	BackendServers DescribeHealthStatusBackendServerList
}

type DescribeHealthStatusBackendServer struct {
	ListenerPort       common.Integer
	ServerId           common.String
	Port               common.Integer
	ServerHealthStatus common.String
	ServerIp           common.String
	Type               common.String
}

type DescribeHealthStatusBackendServerList []DescribeHealthStatusBackendServer

func (list *DescribeHealthStatusBackendServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeHealthStatusBackendServer)
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
