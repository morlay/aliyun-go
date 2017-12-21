package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeApiStageRequest struct {
	requests.RpcRequest
	GroupId string `position:"Query" name:"GroupId"`
	StageId string `position:"Query" name:"StageId"`
}

func (req *DescribeApiStageRequest) Invoke(client *sdk.Client) (resp *DescribeApiStageResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeApiStage", "apigateway", "")
	resp = &DescribeApiStageResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeApiStageResponse struct {
	responses.BaseResponse
	RequestId    string
	GroupId      string
	StageId      string
	StageName    string
	Description  string
	CreatedTime  string
	ModifiedTime string
	Variables    DescribeApiStageVariableItemList
}

type DescribeApiStageVariableItem struct {
	VariableName  string
	VariableValue string
}

type DescribeApiStageVariableItemList []DescribeApiStageVariableItem

func (list *DescribeApiStageVariableItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiStageVariableItem)
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
