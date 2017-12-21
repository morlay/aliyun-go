package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLiveStreamFrameInfoRequest struct {
	requests.RpcRequest
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
}

func (req *DescribeLiveStreamFrameInfoRequest) Invoke(client *sdk.Client) (resp *DescribeLiveStreamFrameInfoResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamFrameInfo", "", "")
	resp = &DescribeLiveStreamFrameInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLiveStreamFrameInfoResponse struct {
	responses.BaseResponse
	RequestId      string
	FrameDataInfos DescribeLiveStreamFrameInfoFrameDataModelList
}

type DescribeLiveStreamFrameInfoFrameDataModel struct {
	Time       string
	Stream     string
	ClientAddr string
	Server     string
	AudioRate  float32
	AudioByte  float32
	FrameRate  float32
	FrameByte  float32
}

type DescribeLiveStreamFrameInfoFrameDataModelList []DescribeLiveStreamFrameInfoFrameDataModel

func (list *DescribeLiveStreamFrameInfoFrameDataModelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamFrameInfoFrameDataModel)
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
