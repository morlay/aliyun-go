package polardb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyDBClusterDescriptionRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	DBClusterDescription string `position:"Query" name:"DBClusterDescription"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string `position:"Query" name:"DBClusterId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ModifyDBClusterDescriptionRequest) Invoke(client *sdk.Client) (resp *ModifyDBClusterDescriptionResponse, err error) {
	req.InitWithApiInfo("polardb", "2017-08-01", "ModifyDBClusterDescription", "polardb", "")
	resp = &ModifyDBClusterDescriptionResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyDBClusterDescriptionResponse struct {
	responses.BaseResponse
	RequestId string
}
