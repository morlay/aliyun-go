package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SaveApRadioSsidConfigRequest struct {
	requests.RpcRequest
	Nasid               string `position:"Query" name:"Nasid"`
	AuthPort            int    `position:"Query" name:"AuthPort"`
	Hidden              int    `position:"Query" name:"Hidden"`
	DynamicVlan         int    `position:"Query" name:"DynamicVlan"`
	AuthServer          string `position:"Query" name:"AuthServer"`
	SecondaryAcctServer string `position:"Query" name:"SecondaryAcctServer"`
	Ssid                string `position:"Query" name:"Ssid"`
	Cir                 int    `position:"Query" name:"Cir"`
	Mac                 string `position:"Query" name:"Mac"`
	SecondaryAcctSecret string `position:"Query" name:"SecondaryAcctSecret"`
	Ieee80211w          int    `position:"Query" name:"Ieee.80211.w"`
	Network             int    `position:"Query" name:"Network"`
	Isolate             int    `position:"Query" name:"Isolate"`
	ApAssetId           int64  `position:"Query" name:"ApAssetId"`
	EncKey              string `position:"Query" name:"EncKey"`
	MulticastForward    int    `position:"Query" name:"MulticastForward"`
	Encryption          string `position:"Query" name:"Encryption"`
	Wmm                 int    `position:"Query" name:"Wmm"`
	AuthCache           int    `position:"Query" name:"AuthCache"`
	Disabled            int    `position:"Query" name:"Disabled"`
	Id                  int64  `position:"Query" name:"Id"`
	RadioIndex          int    `position:"Query" name:"RadioIndex"`
	IgnoreWeakProbe     int    `position:"Query" name:"IgnoreWeakProbe"`
	Maxassoc            int    `position:"Query" name:"Maxassoc"`
	AcctServer          string `position:"Query" name:"AcctServer"`
	SecondaryAuthServer string `position:"Query" name:"SecondaryAuthServer"`
	DaeClient           string `position:"Query" name:"DaeClient"`
	DaeSecret           string `position:"Query" name:"DaeSecret"`
	DisassocLowAck      int    `position:"Query" name:"DisassocLowAck"`
	SecondaryAuthPort   int    `position:"Query" name:"SecondaryAuthPort"`
	AcctSecret          string `position:"Query" name:"AcctSecret"`
	DisassocWeakRssi    int    `position:"Query" name:"DisassocWeakRssi"`
	SecondaryAcctPort   int    `position:"Query" name:"SecondaryAcctPort"`
	DaePort             int    `position:"Query" name:"DaePort"`
	SsidLb              int    `position:"Query" name:"SsidLb"`
	AcctPort            int    `position:"Query" name:"AcctPort"`
	MaxInactivity       int    `position:"Query" name:"MaxInactivity"`
	VlanDhcp            int    `position:"Query" name:"VlanDhcp"`
	InstantlyEffective  int    `position:"Query" name:"InstantlyEffective"`
	ShortPreamble       int    `position:"Query" name:"ShortPreamble"`
	AuthSecret          string `position:"Query" name:"AuthSecret"`
	SecondaryAuthSecret string `position:"Query" name:"SecondaryAuthSecret"`
	Ownip               string `position:"Query" name:"Ownip"`
}

func (req *SaveApRadioSsidConfigRequest) Invoke(client *sdk.Client) (resp *SaveApRadioSsidConfigResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "SaveApRadioSsidConfig", "", "")
	resp = &SaveApRadioSsidConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveApRadioSsidConfigResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Message   common.String
	Data      common.String
	ErrorCode common.Integer
	ErrorMsg  common.String
}
