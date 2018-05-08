package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeSslVpnClientCertRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	SslVpnClientCertId   string `position:"Query" name:"SslVpnClientCertId"`
}

func (req *DescribeSslVpnClientCertRequest) Invoke(client *sdk.Client) (resp *DescribeSslVpnClientCertResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "DescribeSslVpnClientCert", "vpc", "")
	resp = &DescribeSslVpnClientCertResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeSslVpnClientCertResponse struct {
	responses.BaseResponse
	RequestId          common.String
	RegionId           common.String
	SslVpnClientCertId common.String
	Name               common.String
	SslVpnServerId     common.String
	CaCert             common.String
	ClientCert         common.String
	ClientKey          common.String
	ClientConfig       common.String
	CreateTime         common.Long
	EndTime            common.Long
	Status             common.String
}
