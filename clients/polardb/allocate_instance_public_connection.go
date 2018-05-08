package polardb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type AllocateInstancePublicConnectionRequest struct {
	requests.RpcRequest
	ResourceOwnerId        int64  `position:"Query" name:"ResourceOwnerId"`
	ConnectionStringPrefix string `position:"Query" name:"ConnectionStringPrefix"`
	ResourceOwnerAccount   string `position:"Query" name:"ResourceOwnerAccount"`
	Port                   string `position:"Query" name:"Port"`
	OwnerAccount           string `position:"Query" name:"OwnerAccount"`
	DBInstanceId           string `position:"Query" name:"DBInstanceId"`
	OwnerId                int64  `position:"Query" name:"OwnerId"`
}

func (req *AllocateInstancePublicConnectionRequest) Invoke(client *sdk.Client) (resp *AllocateInstancePublicConnectionResponse, err error) {
	req.InitWithApiInfo("polardb", "2017-08-01", "AllocateInstancePublicConnection", "polardb", "")
	resp = &AllocateInstancePublicConnectionResponse{}
	err = client.DoAction(req, resp)
	return
}

type AllocateInstancePublicConnectionResponse struct {
	responses.BaseResponse
	RequestId        common.String
	DBInstanceId     common.String
	ConnectionString common.String
	IPType           common.String
	Port             common.String
}
