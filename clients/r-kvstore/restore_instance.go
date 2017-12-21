package r_kvstore

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RestoreInstanceRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	BackupId             string `position:"Query" name:"BackupId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *RestoreInstanceRequest) Invoke(client *sdk.Client) (resp *RestoreInstanceResponse, err error) {
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "RestoreInstance", "redisa", "")
	resp = &RestoreInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type RestoreInstanceResponse struct {
	responses.BaseResponse
	RequestId string
}
