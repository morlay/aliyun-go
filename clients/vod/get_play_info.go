package vod

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetPlayInfoRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	StreamType           string `position:"Query" name:"StreamType"`
	Formats              string `position:"Query" name:"Formats"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Channel              string `position:"Query" name:"Channel"`
	VideoId              string `position:"Query" name:"VideoId"`
	PlayerVersion        string `position:"Query" name:"PlayerVersion"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Rand                 string `position:"Query" name:"Rand"`
	ReAuthInfo           string `position:"Query" name:"ReAuthInfo"`
	OutputType           string `position:"Query" name:"OutputType"`
	Definition           string `position:"Query" name:"Definition"`
	AuthTimeout          int64  `position:"Query" name:"AuthTimeout"`
	AuthInfo             string `position:"Query" name:"AuthInfo"`
}

func (req *GetPlayInfoRequest) Invoke(client *sdk.Client) (resp *GetPlayInfoResponse, err error) {
	req.InitWithApiInfo("vod", "2017-03-21", "GetPlayInfo", "vod", "")
	resp = &GetPlayInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetPlayInfoResponse struct {
	responses.BaseResponse
	RequestId    common.String
	PlayInfoList GetPlayInfoPlayInfoList
	VideoBase    GetPlayInfoVideoBase
}

type GetPlayInfoPlayInfo struct {
	Width            common.Long
	Height           common.Long
	Size             common.Long
	PlayURL          common.String
	Bitrate          common.String
	Definition       common.String
	Duration         common.String
	Format           common.String
	Fps              common.String
	Encrypt          common.Long
	Plaintext        common.String
	Complexity       common.String
	StreamType       common.String
	Rand             common.String
	JobId            common.String
	PreprocessStatus common.String
}

type GetPlayInfoVideoBase struct {
	CoverURL     common.String
	Duration     common.String
	Status       common.String
	Title        common.String
	VideoId      common.String
	MediaType    common.String
	CreationTime common.String
}

type GetPlayInfoPlayInfoList []GetPlayInfoPlayInfo

func (list *GetPlayInfoPlayInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetPlayInfoPlayInfo)
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
