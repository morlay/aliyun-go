package qualitycheck

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetScoreInfoRequest struct {
	requests.RpcRequest
	JsonStr string `position:"Query" name:"JsonStr"`
}

func (req *GetScoreInfoRequest) Invoke(client *sdk.Client) (resp *GetScoreInfoResponse, err error) {
	req.InitWithApiInfo("Qualitycheck", "2016-08-01", "GetScoreInfo", "", "")
	resp = &GetScoreInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetScoreInfoResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Code      common.String
	Message   common.String
	Data      GetScoreInfoScorePoList
}

type GetScoreInfoScorePo struct {
	ScoreId    common.Integer
	ScoreName  common.String
	ScoreInfos GetScoreInfoScoreParamList
}

type GetScoreInfoScoreParam struct {
	ScoreNum     common.Integer
	ScoreSubId   common.Integer
	ScoreSubName common.String
	ScoreType    common.Integer
}

type GetScoreInfoScorePoList []GetScoreInfoScorePo

func (list *GetScoreInfoScorePoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetScoreInfoScorePo)
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

type GetScoreInfoScoreParamList []GetScoreInfoScoreParam

func (list *GetScoreInfoScoreParamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetScoreInfoScoreParam)
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
