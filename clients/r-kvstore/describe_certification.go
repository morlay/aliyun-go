package r_kvstore

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeCertificationRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Parameters           string `position:"Query" name:"Parameters"`
}

func (req *DescribeCertificationRequest) Invoke(client *sdk.Client) (resp *DescribeCertificationResponse, err error) {
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "DescribeCertification", "redisa", "")
	resp = &DescribeCertificationResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeCertificationResponse struct {
	responses.BaseResponse
	RequestId       common.String
	NoCertification bool
}
