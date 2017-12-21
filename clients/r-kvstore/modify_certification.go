package r_kvstore

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyCertificationRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	NoCertification      string `position:"Query" name:"NoCertification"`
}

func (req *ModifyCertificationRequest) Invoke(client *sdk.Client) (resp *ModifyCertificationResponse, err error) {
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "ModifyCertification", "redisa", "")
	resp = &ModifyCertificationResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyCertificationResponse struct {
	responses.BaseResponse
	RequestId string
}
