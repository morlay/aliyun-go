package vpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeSslVpnClientCertsRequest struct {
	requests.RpcRequest
	SslVpnServerId       string `position:"Query" name:"SslVpnServerId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	PageSize             int    `position:"Query" name:"PageSize"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	SslVpnClientCertId   string `position:"Query" name:"SslVpnClientCertId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeSslVpnClientCertsRequest) Invoke(client *sdk.Client) (resp *DescribeSslVpnClientCertsResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "DescribeSslVpnClientCerts", "vpc", "")
	resp = &DescribeSslVpnClientCertsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeSslVpnClientCertsResponse struct {
	responses.BaseResponse
	RequestId            common.String
	TotalCount           common.Integer
	PageNumber           common.Integer
	PageSize             common.Integer
	SslVpnClientCertKeys DescribeSslVpnClientCertsSslVpnClientCertKeyList
}

type DescribeSslVpnClientCertsSslVpnClientCertKey struct {
	RegionId           common.String
	SslVpnClientCertId common.String
	Name               common.String
	SslVpnServerId     common.String
	CreateTime         common.Long
	EndTime            common.Long
	Status             common.String
}

type DescribeSslVpnClientCertsSslVpnClientCertKeyList []DescribeSslVpnClientCertsSslVpnClientCertKey

func (list *DescribeSslVpnClientCertsSslVpnClientCertKeyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSslVpnClientCertsSslVpnClientCertKey)
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
