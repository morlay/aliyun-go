package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetSupportedOpsCommandRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64 `position:"Query" name:"ResourceOwnerId"`
}

func (req *GetSupportedOpsCommandRequest) Invoke(client *sdk.Client) (resp *GetSupportedOpsCommandResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "GetSupportedOpsCommand", "", "")
	resp = &GetSupportedOpsCommandResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetSupportedOpsCommandResponse struct {
	responses.BaseResponse
	RequestId string
	List      GetSupportedOpsCommandOpsCommandCategoryList
}

type GetSupportedOpsCommandOpsCommandCategory struct {
	Category    string
	CommandList GetSupportedOpsCommandOpsCommandInfoList
}

type GetSupportedOpsCommandOpsCommandInfo struct {
	Id          string
	Name        string
	Discription string
	TargetType  string
	Params      string
}

type GetSupportedOpsCommandOpsCommandCategoryList []GetSupportedOpsCommandOpsCommandCategory

func (list *GetSupportedOpsCommandOpsCommandCategoryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetSupportedOpsCommandOpsCommandCategory)
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

type GetSupportedOpsCommandOpsCommandInfoList []GetSupportedOpsCommandOpsCommandInfo

func (list *GetSupportedOpsCommandOpsCommandInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetSupportedOpsCommandOpsCommandInfo)
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
