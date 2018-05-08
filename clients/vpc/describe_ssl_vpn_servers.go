package vpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeSslVpnServersRequest struct {
	requests.RpcRequest
	SslVpnServerId       string `position:"Query" name:"SslVpnServerId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	PageSize             int    `position:"Query" name:"PageSize"`
	VpnGatewayId         string `position:"Query" name:"VpnGatewayId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeSslVpnServersRequest) Invoke(client *sdk.Client) (resp *DescribeSslVpnServersResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "DescribeSslVpnServers", "vpc", "")
	resp = &DescribeSslVpnServersResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeSslVpnServersResponse struct {
	responses.BaseResponse
	RequestId     common.String
	TotalCount    common.Integer
	PageNumber    common.Integer
	PageSize      common.Integer
	SslVpnServers DescribeSslVpnServersSslVpnServerList
}

type DescribeSslVpnServersSslVpnServer struct {
	RegionId       common.String
	SslVpnServerId common.String
	VpnGatewayId   common.String
	Name           common.String
	LocalSubnet    common.String
	ClientIpPool   common.String
	CreateTime     common.Long
	Cipher         common.String
	Proto          common.String
	Port           common.Integer
	Compress       bool
	Connections    common.Integer
	MaxConnections common.Integer
	InternetIp     common.String
}

type DescribeSslVpnServersSslVpnServerList []DescribeSslVpnServersSslVpnServer

func (list *DescribeSslVpnServersSslVpnServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSslVpnServersSslVpnServer)
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
