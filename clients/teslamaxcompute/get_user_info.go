package teslamaxcompute

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetUserInfoRequest struct {
	requests.RpcRequest
	PageSize int    `position:"Query" name:"PageSize"`
	PageNum  int    `position:"Query" name:"PageNum"`
	User     string `position:"Query" name:"User"`
	Status   string `position:"Query" name:"Status"`
}

func (req *GetUserInfoRequest) Invoke(client *sdk.Client) (resp *GetUserInfoResponse, err error) {
	req.InitWithApiInfo("TeslaMaxCompute", "2017-11-30", "GetUserInfo", "", "")
	resp = &GetUserInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetUserInfoResponse struct {
	responses.BaseResponse
	Code      int
	Message   string
	RequestId string
	Data      GetUserInfoData
}

type GetUserInfoData struct {
	Total  int
	Detail GetUserInfoInstanceList
}

type GetUserInfoInstance struct {
	Project     string
	InstanceId  string
	Status      string
	UserAccount string
	NickName    string
	Cluster     string
	RunTime     string
}

type GetUserInfoInstanceList []GetUserInfoInstance

func (list *GetUserInfoInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetUserInfoInstance)
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
