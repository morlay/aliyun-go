package vpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId      common.String
	TotalCount     common.Integer
	PageNumber     common.Integer
	PageSize       common.Integer
	VpnConnections DescribeVpnConnectionsVpnConnectionList
}

type DescribeVpnConnectionsVpnConnection struct {
	VpnConnectionId   common.String
	CustomerGatewayId common.String
	VpnGatewayId      common.String
	Name              common.String
	LocalSubnet       common.String
	RemoteSubnet      common.String
	CreateTime        common.Long
	EffectImmediately bool
	Status            common.String
	IkeConfig         DescribeVpnConnectionsIkeConfig
	IpsecConfig       DescribeVpnConnectionsIpsecConfig
}

type DescribeVpnConnectionsIkeConfig struct {
	Psk         common.String
	IkeVersion  common.String
	IkeMode     common.String
	IkeEncAlg   common.String
	IkeAuthAlg  common.String
	IkePfs      common.String
	IkeLifetime common.Long
	LocalId     common.String
	RemoteId    common.String
}

type DescribeVpnConnectionsIpsecConfig struct {
	IpsecEncAlg   common.String
	IpsecAuthAlg  common.String
	IpsecPfs      common.String
	IpsecLifetime common.Long
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
