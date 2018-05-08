package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SaveApRadioConfigRequest struct {
	requests.RpcRequest
	RequireMode        string `position:"Query" name:"RequireMode"`
	Htmode             string `position:"Query" name:"Htmode"`
	Frag               int    `position:"Query" name:"Frag"`
	Minrate            int    `position:"Query" name:"Minrate"`
	McastRate          int    `position:"Query" name:"McastRate"`
	Probereq           int    `position:"Query" name:"Probereq"`
	Channel            int    `position:"Query" name:"Channel"`
	Shortgi            int    `position:"Query" name:"Shortgi"`
	Hwmode             string `position:"Query" name:"Hwmode"`
	Uapsd              int    `position:"Query" name:"Uapsd"`
	BeaconInt          int    `position:"Query" name:"BeaconInt"`
	Mac                string `position:"Query" name:"Mac"`
	Rts                int    `position:"Query" name:"Rts"`
	Txpower            int    `position:"Query" name:"Txpower"`
	Noscan             int    `position:"Query" name:"Noscan"`
	BcastRate          int    `position:"Query" name:"BcastRate"`
	Disabled           int    `position:"Query" name:"Disabled"`
	InstantlyEffective int    `position:"Query" name:"InstantlyEffective"`
	Id                 int64  `position:"Query" name:"Id"`
	RadioIndex         int    `position:"Query" name:"RadioIndex"`
}

func (req *SaveApRadioConfigRequest) Invoke(client *sdk.Client) (resp *SaveApRadioConfigResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "SaveApRadioConfig", "", "")
	resp = &SaveApRadioConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveApRadioConfigResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Message   common.String
	Data      common.String
	ErrorCode common.Integer
	ErrorMsg  common.String
}
