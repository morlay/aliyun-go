package teslamaxcompute

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetQuotaInfoRequest struct {
	requests.RpcRequest
	Cluster  string `position:"Query" name:"Cluster"`
	PageSize int    `position:"Query" name:"PageSize"`
	QuotaId  string `position:"Query" name:"QuotaId"`
	PageNum  int    `position:"Query" name:"PageNum"`
	Status   string `position:"Query" name:"Status"`
}

func (req *GetQuotaInfoRequest) Invoke(client *sdk.Client) (resp *GetQuotaInfoResponse, err error) {
	req.InitWithApiInfo("TeslaMaxCompute", "2017-11-30", "GetQuotaInfo", "", "")
	resp = &GetQuotaInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetQuotaInfoResponse struct {
	responses.BaseResponse
	Code      int
	Message   string
	RequestId string
	Data      GetQuotaInfoData
}

type GetQuotaInfoData struct {
	Total  int
	Detail GetQuotaInfoInstanceList
}

type GetQuotaInfoInstance struct {
	Project     string
	InstanceId  string
	Status      string
	UserAccount string
	NickName    string
	Cluster     string
	RunTime     string
}

type GetQuotaInfoInstanceList []GetQuotaInfoInstance

func (list *GetQuotaInfoInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetQuotaInfoInstance)
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
