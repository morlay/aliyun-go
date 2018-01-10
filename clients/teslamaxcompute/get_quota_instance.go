package teslamaxcompute

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetQuotaInstanceRequest struct {
	requests.RpcRequest
	Cluster  string `position:"Query" name:"Cluster"`
	PageSize int    `position:"Query" name:"PageSize"`
	QuotaId  string `position:"Query" name:"QuotaId"`
	PageNum  int    `position:"Query" name:"PageNum"`
	Status   string `position:"Query" name:"Status"`
}

func (req *GetQuotaInstanceRequest) Invoke(client *sdk.Client) (resp *GetQuotaInstanceResponse, err error) {
	req.InitWithApiInfo("TeslaMaxCompute", "2018-01-04", "GetQuotaInstance", "", "")
	resp = &GetQuotaInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetQuotaInstanceResponse struct {
	responses.BaseResponse
	Code      int
	Message   string
	RequestId string
	Data      GetQuotaInstanceData
}

type GetQuotaInstanceData struct {
	Total  int
	Detail GetQuotaInstanceInstanceList
}

type GetQuotaInstanceInstance struct {
	Project         string
	InstanceId      string
	Status          string
	UserAccount     string
	NickName        string
	Cluster         string
	RunTime         string
	CpuUsed         int64
	CpuRequest      int64
	CpuUsedTotal    int64
	CpuUsedRatioMax float32
	CpuUsedRatioMin float32
	MemUsed         int64
	MemRequest      int64
	MemUsedTotal    int64
	MemUsedRatioMax float32
	MemUsedRatioMin float32
	TaskType        string
	SkynetId        string
	QuotaName       string
	QuotaId         int
}

type GetQuotaInstanceInstanceList []GetQuotaInstanceInstance

func (list *GetQuotaInstanceInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetQuotaInstanceInstance)
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
