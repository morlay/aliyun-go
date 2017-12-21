package teslamaxcompute

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetProjectInfoRequest struct {
	PageSize int    `position:"Query" name:"PageSize"`
	Project  string `position:"Query" name:"Project"`
	PageNum  int    `position:"Query" name:"PageNum"`
	Status   string `position:"Query" name:"Status"`
}

func (r GetProjectInfoRequest) Invoke(client *sdk.Client) (response *GetProjectInfoResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetProjectInfoRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("TeslaMaxCompute", "2017-11-30", "GetProjectInfo", "", "")

	resp := struct {
		*responses.BaseResponse
		GetProjectInfoResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetProjectInfoResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetProjectInfoResponse struct {
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
