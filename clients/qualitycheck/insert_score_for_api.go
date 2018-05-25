package qualitycheck

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type InsertScoreForApiRequest struct {
	requests.RpcRequest
	JsonStr string `position:"Query" name:"JsonStr"`
}

func (req *InsertScoreForApiRequest) Invoke(client *sdk.Client) (resp *InsertScoreForApiResponse, err error) {
	req.InitWithApiInfo("Qualitycheck", "2016-08-01", "InsertScoreForApi", "", "")
	resp = &InsertScoreForApiResponse{}
	err = client.DoAction(req, resp)
	return
}

type InsertScoreForApiResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Code      common.String
	Message   common.String
	Data      InsertScoreForApiScorePoList
}

type InsertScoreForApiScorePo struct {
	ScoreId   common.Long
	ScoreName common.String
}

type InsertScoreForApiScorePoList []InsertScoreForApiScorePo

func (list *InsertScoreForApiScorePoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]InsertScoreForApiScorePo)
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
