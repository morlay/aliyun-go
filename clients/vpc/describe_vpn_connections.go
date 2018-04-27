package vpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeVpnConnectionsRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	VpnConnectionId      string `position:"Query" name:"VpnConnectionId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int    `position:"Query" name:"PageSize"`
	VpnGatewayId         string `position:"Query" name:"VpnGatewayId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	CustomerGatewayId    string `position:"Query" name:"CustomerGatewayId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeVpnConnectionsRequest) Invoke(client *sdk.Client) (resp *DescribeVpnConnectionsResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "DescribeVpnConnections", "vpc", "")
	resp = &DescribeVpnConnectionsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeVpnConnectionsResponse struct {
	responses.BaseResponse
	RequestId      string
	TotalCount     int
	PageNumber     int
	PageSize       int
	VpnConnections DescribeVpnConnectionsVpnConnectionList
}

type DescribeVpnConnectionsVpnConnection struct {
	VpnConnectionId   string
	CustomerGatewayId string
	VpnGatewayId      string
	Name              string
	LocalSubnet       string
	RemoteSubnet      string
	CreateTime        int64
	EffectImmediately bool
	Status            string
	IkeConfig         DescribeVpnConnectionsIkeConfig
	IpsecConfig       DescribeVpnConnectionsIpsecConfig
}

type DescribeVpnConnectionsIkeConfig struct {
	Psk         string
	IkeVersion  string
	IkeMode     string
	IkeEncAlg   string
	IkeAuthAlg  string
	IkePfs      string
	IkeLifetime int64
	LocalId     string
	RemoteId    string
}

type DescribeVpnConnectionsIpsecConfig struct {
	IpsecEncAlg   string
	IpsecAuthAlg  string
	IpsecPfs      string
	IpsecLifetime int64
}

type DescribeVpnConnectionsVpnConnectionList []DescribeVpnConnectionsVpnConnection

func (list *DescribeVpnConnectionsVpnConnectionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeVpnConnectionsVpnConnection)
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
