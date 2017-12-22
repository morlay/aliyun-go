package teslamaxcompute

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetClusterInstanceRequest struct {
	requests.RpcRequest
	Cluster  string `position:"Query" name:"Cluster"`
	PageSize int    `position:"Query" name:"PageSize"`
	PageNum  int    `position:"Query" name:"PageNum"`
	Status   string `position:"Query" name:"Status"`
}

func (req *GetClusterInstanceRequest) Invoke(client *sdk.Client) (resp *GetClusterInstanceResponse, err error) {
	req.InitWithApiInfo("TeslaMaxCompute", "2017-11-30", "GetClusterInstance", "", "")
	resp = &GetClusterInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetClusterInstanceResponse struct {
	responses.BaseResponse
	Code      int
	Message   string
	RequestId string
	Data      GetClusterInstanceData
}

type GetClusterInstanceData struct {
	Total  int
	Detail GetClusterInstanceInstanceList
}

type GetClusterInstanceInstance struct {
	Project         string
	InstanceId      string
	Status          string
	UserAccount     string
	NickName        string
	Cluster         string
	RunTime         string
	CpuUsed         int
	CpuRequest      int
	CpuUsedTotal    int
	CpuUsedRatioMax float32
	CpuUsedRatioMin float32
	MemUsed         int
	MemRequest      int
	MemUsedTotal    int
	MemUsedRatioMax float32
	MemUsedRatioMin float32
	TaskType        string
	SkynetId        string
}

type GetClusterInstanceInstanceList []GetClusterInstanceInstance

func (list *GetClusterInstanceInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetClusterInstanceInstance)
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
