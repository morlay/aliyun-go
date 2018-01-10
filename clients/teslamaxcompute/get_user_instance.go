package teslamaxcompute

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetUserInstanceRequest struct {
	requests.RpcRequest
	PageSize int    `position:"Query" name:"PageSize"`
	PageNum  int    `position:"Query" name:"PageNum"`
	Region   string `position:"Query" name:"Region"`
	User     string `position:"Query" name:"User"`
	Status   string `position:"Query" name:"Status"`
}

func (req *GetUserInstanceRequest) Invoke(client *sdk.Client) (resp *GetUserInstanceResponse, err error) {
	req.InitWithApiInfo("TeslaMaxCompute", "2018-01-04", "GetUserInstance", "", "")
	resp = &GetUserInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetUserInstanceResponse struct {
	responses.BaseResponse
	Code      int
	Message   string
	RequestId string
	Data      GetUserInstanceData
}

type GetUserInstanceData struct {
	Total  int
	Detail GetUserInstanceInstanceList
}

type GetUserInstanceInstance struct {
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

type GetUserInstanceInstanceList []GetUserInstanceInstance

func (list *GetUserInstanceInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetUserInstanceInstance)
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
