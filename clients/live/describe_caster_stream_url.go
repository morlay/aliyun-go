package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeCasterStreamUrlRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	CasterId      string `position:"Query" name:"CasterId"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Version       string `position:"Query" name:"Version"`
}

func (r DescribeCasterStreamUrlRequest) Invoke(client *sdk.Client) (response *DescribeCasterStreamUrlResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeCasterStreamUrlRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("live", "2016-11-01", "DescribeCasterStreamUrl", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeCasterStreamUrlResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeCasterStreamUrlResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeCasterStreamUrlResponse struct {
	RequestId     string
	CasterId      string
	Total         int
	CasterStreams DescribeCasterStreamUrlCasterStreamList
}

type DescribeCasterStreamUrlCasterStream struct {
	SceneId     string
	StreamUrl   string
	OutputType  int
	StreamInfos DescribeCasterStreamUrlStreamInfoList
}

type DescribeCasterStreamUrlStreamInfo struct {
	TranscodeConfig string
	VideoFormat     string
	OutputStreamUrl string
}

type DescribeCasterStreamUrlCasterStreamList []DescribeCasterStreamUrlCasterStream

func (list *DescribeCasterStreamUrlCasterStreamList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCasterStreamUrlCasterStream)
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

type DescribeCasterStreamUrlStreamInfoList []DescribeCasterStreamUrlStreamInfo

func (list *DescribeCasterStreamUrlStreamInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCasterStreamUrlStreamInfo)
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
