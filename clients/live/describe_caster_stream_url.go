package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeCasterStreamUrlRequest struct {
	requests.RpcRequest
	CasterId string `position:"Query" name:"CasterId"`
	OwnerId  int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeCasterStreamUrlRequest) Invoke(client *sdk.Client) (resp *DescribeCasterStreamUrlResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "DescribeCasterStreamUrl", "live", "")
	resp = &DescribeCasterStreamUrlResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeCasterStreamUrlResponse struct {
	responses.BaseResponse
	RequestId     common.String
	CasterId      common.String
	Total         common.Integer
	CasterStreams DescribeCasterStreamUrlCasterStreamList
}

type DescribeCasterStreamUrlCasterStream struct {
	SceneId     common.String
	StreamUrl   common.String
	OutputType  common.Integer
	StreamInfos DescribeCasterStreamUrlStreamInfoList
}

type DescribeCasterStreamUrlStreamInfo struct {
	TranscodeConfig common.String
	VideoFormat     common.String
	OutputStreamUrl common.String
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
