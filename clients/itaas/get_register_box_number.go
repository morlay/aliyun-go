package itaas

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetRegisterBoxNumberRequest struct {
	Sysfrom     string `position:"Query" name:"Sysfrom"`
	Operator    string `position:"Query" name:"Operator"`
	Clientappid string `position:"Query" name:"Clientappid"`
}

func (r GetRegisterBoxNumberRequest) Invoke(client *sdk.Client) (response *GetRegisterBoxNumberResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetRegisterBoxNumberRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("ITaaS", "2017-05-05", "GetRegisterBoxNumber", "", "")

	resp := struct {
		*responses.BaseResponse
		GetRegisterBoxNumberResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetRegisterBoxNumberResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetRegisterBoxNumberResponse struct {
	RequestId string
	ErrorCode int
	ErrorMsg  string
	Success   bool
	ErrorList GetRegisterBoxNumberErrorMessageList
	Data      GetRegisterBoxNumberData
}

type GetRegisterBoxNumberErrorMessage struct {
	ErrorMessage string
}

type GetRegisterBoxNumberData struct {
	ActivedNumber int
	BuyNumber     int
	BoxesList     GetRegisterBoxNumberBoxInfoList
}

type GetRegisterBoxNumberBoxInfo struct {
	CurVersion      string
	DrName          string
	DrSessionId     string
	DrStatus        string
	DrStatusTxt     string
	Ipaddress       string
	LastReportTimeL int64
	OnlineTimeL     int64
	Screencode      string
	SysVersion      string
}

type GetRegisterBoxNumberErrorMessageList []GetRegisterBoxNumberErrorMessage

func (list *GetRegisterBoxNumberErrorMessageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetRegisterBoxNumberErrorMessage)
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

type GetRegisterBoxNumberBoxInfoList []GetRegisterBoxNumberBoxInfo

func (list *GetRegisterBoxNumberBoxInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetRegisterBoxNumberBoxInfo)
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
