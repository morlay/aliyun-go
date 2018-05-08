package polardb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type AllocateClusterPublicConnectionRequest struct {
	requests.RpcRequest
	ResourceOwnerId        int64  `position:"Query" name:"ResourceOwnerId"`
	ConnectionStringPrefix string `position:"Query" name:"ConnectionStringPrefix"`
	ResourceOwnerAccount   string `position:"Query" name:"ResourceOwnerAccount"`
	Port                   string `position:"Query" name:"Port"`
	DBClusterId            string `position:"Query" name:"DBClusterId"`
	OwnerAccount           string `position:"Query" name:"OwnerAccount"`
	OwnerId                int64  `position:"Query" name:"OwnerId"`
}

func (req *AllocateClusterPublicConnectionRequest) Invoke(client *sdk.Client) (resp *AllocateClusterPublicConnectionResponse, err error) {
	req.InitWithApiInfo("polardb", "2017-08-01", "AllocateClusterPublicConnection", "polardb", "")
	resp = &AllocateClusterPublicConnectionResponse{}
	err = client.DoAction(req, resp)
	return
}

type AllocateClusterPublicConnectionResponse struct {
	responses.BaseResponse
	RequestId        common.String
	DBInstanceId     common.String
	ConnectionString common.String
	IPType           common.String
	Port             common.String
}
