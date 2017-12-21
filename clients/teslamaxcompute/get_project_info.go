package teslamaxcompute

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetProjectInfoRequest struct {
	requests.RpcRequest
	PageSize int    `position:"Query" name:"PageSize"`
	Project  string `position:"Query" name:"Project"`
	PageNum  int    `position:"Query" name:"PageNum"`
	Status   string `position:"Query" name:"Status"`
}

func (req *GetProjectInfoRequest) Invoke(client *sdk.Client) (resp *GetProjectInfoResponse, err error) {
	req.InitWithApiInfo("TeslaMaxCompute", "2017-11-30", "GetProjectInfo", "", "")
	resp = &GetProjectInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetProjectInfoResponse struct {
	responses.BaseResponse
	Code      int
	Message   string
	RequestId string
	Data      GetProjectInfoData
}

type GetProjectInfoData struct {
	Total  int
	Detail GetProjectInfoInstanceList
}

type GetProjectInfoInstance struct {
	Project     string
	InstanceId  string
	Status      string
	UserAccount string
	NickName    string
	Cluster     string
	RunTime     string
}

type GetProjectInfoInstanceList []GetProjectInfoInstance

func (list *GetProjectInfoInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetProjectInfoInstance)
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
