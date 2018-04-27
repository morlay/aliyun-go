package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId          string
	RegionId           string
	SslVpnClientCertId string
	Name               string
	SslVpnServerId     string
	CaCert             string
	ClientCert         string
	ClientKey          string
	ClientConfig       string
	CreateTime         int64
	EndTime            int64
	Status             string
}
