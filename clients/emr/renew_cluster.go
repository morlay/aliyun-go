package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type RenewClusterRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64                       `position:"Query" name:"ResourceOwnerId"`
	RenewEcsDos     *RenewClusterRenewEcsDoList `position:"Query" type:"Repeated" name:"RenewEcsDo"`
	ClusterId       string                      `position:"Query" name:"ClusterId"`
}

func (req *RenewClusterRequest) Invoke(client *sdk.Client) (resp *RenewClusterResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "RenewCluster", "", "")
	resp = &RenewClusterResponse{}
	err = client.DoAction(req, resp)
	return
}

type RenewClusterRenewEcsDo struct {
	EcsId     string `name:"EcsId"`
	EcsPeriod string `name:"EcsPeriod"`
	EmrPeriod string `name:"EmrPeriod"`
}

type RenewClusterResponse struct {
	responses.BaseResponse
	RequestId      common.String
	EcsOrderIdList common.String
	EmrOrderIdList common.String
}

type RenewClusterRenewEcsDoList []RenewClusterRenewEcsDo

func (list *RenewClusterRenewEcsDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]RenewClusterRenewEcsDo)
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
