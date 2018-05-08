package itaas

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetRegisterBoxNumberRequest struct {
	requests.RpcRequest
	Sysfrom     string `position:"Query" name:"Sysfrom"`
	Operator    string `position:"Query" name:"Operator"`
	Clientappid string `position:"Query" name:"Clientappid"`
}

func (req *GetRegisterBoxNumberRequest) Invoke(client *sdk.Client) (resp *GetRegisterBoxNumberResponse, err error) {
	req.InitWithApiInfo("ITaaS", "2017-05-05", "GetRegisterBoxNumber", "", "")
	resp = &GetRegisterBoxNumberResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetRegisterBoxNumberResponse struct {
	responses.BaseResponse
	RequestId common.String
	ErrorCode common.Integer
	ErrorMsg  common.String
	Success   bool
	ErrorList GetRegisterBoxNumberErrorMessageList
	Data      GetRegisterBoxNumberData
}

type GetRegisterBoxNumberErrorMessage struct {
	ErrorMessage common.String
}

type GetRegisterBoxNumberData struct {
	ActivedNumber common.Integer
	BuyNumber     common.Integer
	BoxesList     GetRegisterBoxNumberBoxInfoList
}

type GetRegisterBoxNumberBoxInfo struct {
	CurVersion      common.String
	DrName          common.String
	DrSessionId     common.String
	DrStatus        common.String
	DrStatusTxt     common.String
	Ipaddress       common.String
	LastReportTimeL common.Long
	OnlineTimeL     common.Long
	Screencode      common.String
	SysVersion      common.String
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
