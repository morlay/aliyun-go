package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type PlayInfoRequest struct {
	requests.RpcRequest
	PlayDomain           string `position:"Query" name:"PlayDomain"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	Formats              string `position:"Query" name:"Formats"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	HlsUriToken          string `position:"Query" name:"HlsUriToken"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	MediaId              string `position:"Query" name:"MediaId"`
	Rand                 string `position:"Query" name:"Rand"`
	AuthTimeout          int64  `position:"Query" name:"AuthTimeout"`
	AuthInfo             string `position:"Query" name:"AuthInfo"`
}

func (req *PlayInfoRequest) Invoke(client *sdk.Client) (resp *PlayInfoResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "PlayInfo", "mts", "")
	resp = &PlayInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type PlayInfoResponse struct {
	responses.BaseResponse
	RequestId         common.String
	PlayInfoList      PlayInfoPlayInfoList
	NotFoundCDNDomain PlayInfoNotFoundCDNDomainList
}

type PlayInfoPlayInfo struct {
	Url            common.String
	Duration       common.String
	Size           common.String
	Width          common.String
	Height         common.String
	Bitrate        common.String
	Fps            common.String
	Format         common.String
	Definition     common.String
	Encryption     common.String
	Rand           common.String
	Plaintext      common.String
	Complexity     common.String
	ActivityName   common.String
	EncryptionType common.String
	DownloadType   common.String
}

type PlayInfoPlayInfoList []PlayInfoPlayInfo

func (list *PlayInfoPlayInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]PlayInfoPlayInfo)
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

type PlayInfoNotFoundCDNDomainList []common.String

func (list *PlayInfoNotFoundCDNDomainList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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
