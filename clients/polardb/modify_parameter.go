package polardb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyParameterRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	Forcerestart         string `position:"Query" name:"Forcerestart"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Parameters           string `position:"Query" name:"Parameters"`
}

func (req *ModifyParameterRequest) Invoke(client *sdk.Client) (resp *ModifyParameterResponse, err error) {
	req.InitWithApiInfo("polardb", "2017-08-01", "ModifyParameter", "polardb", "")
	resp = &ModifyParameterResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyParameterResponse struct {
	responses.BaseResponse
	RequestId string
}
