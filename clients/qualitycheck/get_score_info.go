package qualitycheck

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetScoreInfoRequest struct {
	JsonStr string `position:"Query" name:"JsonStr"`
}

func (r GetScoreInfoRequest) Invoke(client *sdk.Client) (response *GetScoreInfoResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetScoreInfoRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Qualitycheck", "2016-08-01", "GetScoreInfo", "", "")

	resp := struct {
		*responses.BaseResponse
		GetScoreInfoResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetScoreInfoResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetScoreInfoResponse struct {
	RequestId string
	Success   bool
	Code      string
	Message   string
	Data      GetScoreInfoScorePoList
}

type GetScoreInfoScorePo struct {
	ScoreId    int
	ScoreName  string
	ScoreInfos GetScoreInfoScoreParamList
}

type GetScoreInfoScoreParam struct {
	ScoreNum     int
	ScoreSubId   int
	ScoreSubName string
	ScoreType    int
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
