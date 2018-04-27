package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RenewClusterForAdminRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64                               `position:"Query" name:"ResourceOwnerId"`
	RenewEcsDos     *RenewClusterForAdminRenewEcsDoList `position:"Query" type:"Repeated" name:"RenewEcsDo"`
	ClusterId       string                              `position:"Query" name:"ClusterId"`
	UserId          string                              `position:"Query" name:"UserId"`
}

func (req *RenewClusterForAdminRequest) Invoke(client *sdk.Client) (resp *RenewClusterForAdminResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "RenewClusterForAdmin", "", "")
	resp = &RenewClusterForAdminResponse{}
	err = client.DoAction(req, resp)
	return
}

type RenewClusterForAdminRenewEcsDo struct {
	EcsId     string `name:"EcsId"`
	EcsPeriod string `name:"EcsPeriod"`
	EmrPeriod string `name:"EmrPeriod"`
}

type RenewClusterForAdminResponse struct {
	responses.BaseResponse
	RequestId      string
	EcsOrderIdList string
	EmrOrderIdList string
}

type RenewClusterForAdminRenewEcsDoList []RenewClusterForAdminRenewEcsDo

func (list *RenewClusterForAdminRenewEcsDoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]RenewClusterForAdminRenewEcsDo)
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
