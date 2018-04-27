package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId      string
	ClusterId      string
	RenewEcsDOList CheckRenewClusterRenewEcsDOList
}

type CheckRenewClusterRenewEcsDO struct {
	EcsId                string
	EcsExpiredTime       string
	EmrExpiredTime       string
	EcsAutoRenew         string
	EmrAutoRenew         string
	EcsAutoRenewDuration int
	EmrAutoRenewDuration int
	HostGroupId          string
	HostGroupType        string
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
