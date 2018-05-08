package r_kvstore

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateTempInstanceRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SnapshotId           string `position:"Domain" name:"SnapshotId"`
	InstanceId           string `position:"Domain" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *CreateTempInstanceRequest) Invoke(client *sdk.Client) (resp *CreateTempInstanceResponse, err error) {
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "CreateTempInstance", "redisa", "")
	resp = &CreateTempInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateTempInstanceResponse struct {
	responses.BaseResponse
	RequestId      common.String
	InstanceId     common.String
	SnapshotId     common.String
	TempInstanceId common.String
	Status         common.String
}
