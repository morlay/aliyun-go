package polardb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateDBClusterRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	DBClusterDescription string `position:"Query" name:"DBClusterDescription"`
	Period               string `position:"Query" name:"Period"`
	DBInstanceStorage    string `position:"Query" name:"DBInstanceStorage"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	UsedTime             string `position:"Query" name:"UsedTime"`
	DBInstanceClass      string `position:"Query" name:"DBInstanceClass"`
	ClusterNetworkType   string `position:"Query" name:"ClusterNetworkType"`
	SecurityIPList       string `position:"Query" name:"SecurityIPList"`
	VSwitchId            string `position:"Query" name:"VSwitchId"`
	PrivateIpAddress     string `position:"Query" name:"PrivateIpAddress"`
	Engine               string `position:"Query" name:"Engine"`
	VPCId                string `position:"Query" name:"VPCId"`
	DBType               string `position:"Query" name:"DBType"`
	ZoneId               string `position:"Query" name:"ZoneId"`
	DBVersion            string `position:"Query" name:"DBVersion"`
	PayType              string `position:"Query" name:"PayType"`
}

func (req *CreateDBClusterRequest) Invoke(client *sdk.Client) (resp *CreateDBClusterResponse, err error) {
	req.InitWithApiInfo("polardb", "2017-08-01", "CreateDBCluster", "polardb", "")
	resp = &CreateDBClusterResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateDBClusterResponse struct {
	responses.BaseResponse
	RequestId        common.String
	DBClusterId      common.String
	OrderId          common.String
	ConnectionString common.String
	Port             common.String
}
