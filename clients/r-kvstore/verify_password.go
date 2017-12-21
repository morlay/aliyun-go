package r_kvstore

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type VerifyPasswordRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Password             string `position:"Query" name:"Password"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *VerifyPasswordRequest) Invoke(client *sdk.Client) (resp *VerifyPasswordResponse, err error) {
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "VerifyPassword", "redisa", "")
	resp = &VerifyPasswordResponse{}
	err = client.DoAction(req, resp)
	return
}

type VerifyPasswordResponse struct {
	responses.BaseResponse
	RequestId string
}
