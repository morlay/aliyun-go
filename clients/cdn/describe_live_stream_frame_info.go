package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId      common.String
	FrameDataInfos DescribeLiveStreamFrameInfoFrameDataModelList
}

type DescribeLiveStreamFrameInfoFrameDataModel struct {
	Time       common.String
	Stream     common.String
	ClientAddr common.String
	Server     common.String
	AudioRate  common.Float
	AudioByte  common.Float
	FrameRate  common.Float
	FrameByte  common.Float
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
