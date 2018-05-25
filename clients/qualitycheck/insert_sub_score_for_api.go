package qualitycheck

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type InsertSubScoreForApiRequest struct {
	requests.RpcRequest
	JsonStr string `position:"Query" name:"JsonStr"`
}

func (req *InsertSubScoreForApiRequest) Invoke(client *sdk.Client) (resp *InsertSubScoreForApiResponse, err error) {
	req.InitWithApiInfo("Qualitycheck", "2016-08-01", "InsertSubScoreForApi", "", "")
	resp = &InsertSubScoreForApiResponse{}
	err = client.DoAction(req, resp)
	return
}

type InsertSubScoreForApiResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Code      common.String
	Message   common.String
	Data      InsertSubScoreForApiScorePoList
}

type InsertSubScoreForApiScorePo struct {
	ScoreSubId   common.Long
	ScoreSubName common.String
}

type InsertSubScoreForApiScorePoList []InsertSubScoreForApiScorePo

func (list *InsertSubScoreForApiScorePoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]InsertSubScoreForApiScorePo)
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
