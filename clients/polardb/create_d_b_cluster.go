package polardb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateDBClusterRequest struct {
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

func (r CreateDBClusterRequest) Invoke(client *sdk.Client) (response *CreateDBClusterResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateDBClusterRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("polardb", "2017-08-01", "CreateDBCluster", "polardb", "")

	resp := struct {
		*responses.BaseResponse
		CreateDBClusterResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateDBClusterResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateDBClusterResponse struct {
	RequestId        string
	DBClusterId      string
	OrderId          string
	ConnectionString string
	Port             string
}
