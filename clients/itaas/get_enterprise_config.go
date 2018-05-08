package itaas

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetEnterpriseConfigRequest struct {
	requests.RpcRequest
	Sysfrom     string `position:"Query" name:"Sysfrom"`
	Operator    string `position:"Query" name:"Operator"`
	Clientappid string `position:"Query" name:"Clientappid"`
}

func (req *GetEnterpriseConfigRequest) Invoke(client *sdk.Client) (resp *GetEnterpriseConfigResponse, err error) {
	req.InitWithApiInfo("ITaaS", "2017-05-05", "GetEnterpriseConfig", "", "")
	resp = &GetEnterpriseConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetEnterpriseConfigResponse struct {
	responses.BaseResponse
	RequestId common.String
	ErrorCode common.Integer
	ErrorMsg  common.String
	Success   bool
	ErrorList GetEnterpriseConfigErrorMessageList
	Data      GetEnterpriseConfigData
}

type GetEnterpriseConfigErrorMessage struct {
	ErrorMessage common.String
}

type GetEnterpriseConfigData struct {
	AuthorizationNeedAccessToken bool
	DrMeetingQrUrl               common.String
	DrWelcomeUrl                 common.String
	ShareMboxNubmer              common.Integer
	ShareNeedInternet            bool
	ShareServiceFlag             bool
}

type GetEnterpriseConfigErrorMessageList []GetEnterpriseConfigErrorMessage

func (list *GetEnterpriseConfigErrorMessageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetEnterpriseConfigErrorMessage)
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
