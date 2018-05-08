package polardb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifySecurityIpsRequest struct {
	requests.RpcRequest
	ResourceOwnerId           int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount      string `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId               string `position:"Query" name:"DBClusterId"`
	OwnerAccount              string `position:"Query" name:"OwnerAccount"`
	SecurityIps               string `position:"Query" name:"SecurityIps"`
	DBClusterIPArrayName      string `position:"Query" name:"DBClusterIPArrayName"`
	OwnerId                   int64  `position:"Query" name:"OwnerId"`
	DBClusterIPArrayAttribute string `position:"Query" name:"DBClusterIPArrayAttribute"`
}

func (req *ModifySecurityIpsRequest) Invoke(client *sdk.Client) (resp *ModifySecurityIpsResponse, err error) {
	req.InitWithApiInfo("polardb", "2017-08-01", "ModifySecurityIps", "polardb", "")
	resp = &ModifySecurityIpsResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifySecurityIpsResponse struct {
	responses.BaseResponse
	RequestId common.String
}
