package polardb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyDBClusterConnectionStringRequest struct {
	requests.RpcRequest
	ResourceOwnerId         int64  `position:"Query" name:"ResourceOwnerId"`
	ConnectionStringPrefix  string `position:"Query" name:"ConnectionStringPrefix"`
	ResourceOwnerAccount    string `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId             string `position:"Query" name:"DBClusterId"`
	OwnerAccount            string `position:"Query" name:"OwnerAccount"`
	OwnerId                 int64  `position:"Query" name:"OwnerId"`
	CurrentConnectionString string `position:"Query" name:"CurrentConnectionString"`
}

func (req *ModifyDBClusterConnectionStringRequest) Invoke(client *sdk.Client) (resp *ModifyDBClusterConnectionStringResponse, err error) {
	req.InitWithApiInfo("polardb", "2017-08-01", "ModifyDBClusterConnectionString", "polardb", "")
	resp = &ModifyDBClusterConnectionStringResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyDBClusterConnectionStringResponse struct {
	responses.BaseResponse
	RequestId           string
	DBClusterId         string
	OldConnectionString string
	OldPort             string
	NewConnectionString string
	NewPort             string
	IPType              string
}
