package teslamaxcompute

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetClusterInfoRequest struct {
	requests.RpcRequest
	Cluster  string `position:"Query" name:"Cluster"`
	PageSize int    `position:"Query" name:"PageSize"`
	PageNum  int    `position:"Query" name:"PageNum"`
	Status   string `position:"Query" name:"Status"`
}

func (req *GetClusterInfoRequest) Invoke(client *sdk.Client) (resp *GetClusterInfoResponse, err error) {
	req.InitWithApiInfo("TeslaMaxCompute", "2017-11-30", "GetClusterInfo", "", "")
	resp = &GetClusterInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetClusterInfoResponse struct {
	responses.BaseResponse
	Code      int
	Message   string
	RequestId string
	Data      GetClusterInfoData
}

type GetClusterInfoData struct {
	Total  int
	Detail GetClusterInfoInstanceList
}

type GetClusterInfoInstance struct {
	Project     string
	InstanceId  string
	Status      string
	UserAccount string
	NickName    string
	Cluster     string
	RunTime     string
}

type GetClusterInfoInstanceList []GetClusterInfoInstance

func (list *GetClusterInfoInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetClusterInfoInstance)
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
