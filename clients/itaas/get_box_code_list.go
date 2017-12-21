package itaas

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetBoxCodeListRequest struct {
	Sysfrom     string `position:"Query" name:"Sysfrom"`
	Operator    string `position:"Query" name:"Operator"`
	Clientappid string `position:"Query" name:"Clientappid"`
}

func (r GetBoxCodeListRequest) Invoke(client *sdk.Client) (response *GetBoxCodeListResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetBoxCodeListRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("ITaaS", "2017-05-05", "GetBoxCodeList", "", "")

	resp := struct {
		*responses.BaseResponse
		GetBoxCodeListResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetBoxCodeListResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetBoxCodeListResponse struct {
	RequestId string
	ErrorCode int
	ErrorMsg  string
	Success   bool
	Data      GetBoxCodeListBoxCodeInfoList
	ErrorList GetBoxCodeListErrorMessageList
}

type GetBoxCodeListBoxCodeInfo struct {
	BeginTime  int64
	BoxInfo    string
	Code       string
	EndTime    int64
	ModifyTime int64
	Operator   string
	Screencode string
	Status     int
	StatusTxt  string
}

type GetBoxCodeListErrorMessage struct {
	ErrorMessage string
}

type GetBoxCodeListBoxCodeInfoList []GetBoxCodeListBoxCodeInfo

func (list *GetBoxCodeListBoxCodeInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetBoxCodeListBoxCodeInfo)
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

type GetBoxCodeListErrorMessageList []GetBoxCodeListErrorMessage

func (list *GetBoxCodeListErrorMessageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetBoxCodeListErrorMessage)
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
