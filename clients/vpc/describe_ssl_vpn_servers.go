package vpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId     string
	TotalCount    int
	PageNumber    int
	PageSize      int
	SslVpnServers DescribeSslVpnServersSslVpnServerList
}

type DescribeSslVpnServersSslVpnServer struct {
	RegionId       string
	SslVpnServerId string
	VpnGatewayId   string
	Name           string
	LocalSubnet    string
	ClientIpPool   string
	CreateTime     int64
	Cipher         string
	Proto          string
	Port           int
	Compress       bool
	Connections    int
	MaxConnections int
	InternetIp     string
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
