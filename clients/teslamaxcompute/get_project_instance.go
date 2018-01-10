package teslamaxcompute

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetProjectInstanceRequest struct {
	requests.RpcRequest
	PageSize int    `position:"Query" name:"PageSize"`
	Project  string `position:"Query" name:"Project"`
	PageNum  int    `position:"Query" name:"PageNum"`
	Region   string `position:"Query" name:"Region"`
	Status   string `position:"Query" name:"Status"`
}

func (req *GetProjectInstanceRequest) Invoke(client *sdk.Client) (resp *GetProjectInstanceResponse, err error) {
	req.InitWithApiInfo("TeslaMaxCompute", "2018-01-04", "GetProjectInstance", "", "")
	resp = &GetProjectInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetProjectInstanceResponse struct {
	responses.BaseResponse
	Code      int
	Message   string
	RequestId string
	Data      GetProjectInstanceData
}

type GetProjectInstanceData struct {
	Total  int
	Detail GetProjectInstanceInstanceList
}

type GetProjectInstanceInstance struct {
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

type GetProjectInstanceInstanceList []GetProjectInstanceInstance

func (list *GetProjectInstanceInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetProjectInstanceInstance)
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
