package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CheckRenewClusterRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
}

func (req *CheckRenewClusterRequest) Invoke(client *sdk.Client) (resp *CheckRenewClusterResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "CheckRenewCluster", "", "")
	resp = &CheckRenewClusterResponse{}
	err = client.DoAction(req, resp)
	return
}

type CheckRenewClusterResponse struct {
	responses.BaseResponse
	RequestId      common.String
	ClusterId      common.String
	RenewEcsDOList CheckRenewClusterRenewEcsDOList
}

type CheckRenewClusterRenewEcsDO struct {
	EcsId                common.String
	EcsExpiredTime       common.String
	EmrExpiredTime       common.String
	EcsAutoRenew         common.String
	EmrAutoRenew         common.String
	EcsAutoRenewDuration common.Integer
	EmrAutoRenewDuration common.Integer
	HostGroupId          common.String
	HostGroupType        common.String
}

type CheckRenewClusterRenewEcsDOList []CheckRenewClusterRenewEcsDO

func (list *CheckRenewClusterRenewEcsDOList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CheckRenewClusterRenewEcsDO)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
