package teslamaxcompute

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetUserInfoRequest struct {
	PageSize int    `position:"Query" name:"PageSize"`
	PageNum  int    `position:"Query" name:"PageNum"`
	User     string `position:"Query" name:"User"`
	Status   string `position:"Query" name:"Status"`
}

func (r GetUserInfoRequest) Invoke(client *sdk.Client) (response *GetUserInfoResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetUserInfoRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("TeslaMaxCompute", "2017-11-30", "GetUserInfo", "", "")

	resp := struct {
		*responses.BaseResponse
		GetUserInfoResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetUserInfoResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetUserInfoResponse struct {
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
