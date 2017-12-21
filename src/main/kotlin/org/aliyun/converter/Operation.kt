package org.aliyun.converter


class Operation(
        var requestType: String,
        val product: String,
        val version: String,
        val action: String,
        val serviceCode: String,
        var method: String,
        val uriPattern: String
) {
    companion object {
        fun withAnonymous(structStruct: String, s: String): String {
            return structStruct.replace(Regex("^struct \\{"), "struct {\n${s}")
        }
    }

    private var parameters = Schema("")
    private var response = Schema("")

    fun setParam(s: Schema) {
        this.parameters = s
    }

    fun setResp(s: Schema) {
        this.response = s
    }

    fun goClientMethod(): String {
        val subTypePrefix = this.action
        val sideCodes = mutableMapOf<String, String>()

        var codes = ""

        when (requestType) {
            "roa" -> {
                codes += """
type ${this.action}Request ${withAnonymous(this.parameters.goType(subTypePrefix, true, sideCodes), "requests.RoaRequest")}

func (req *${this.action}Request) Invoke(client *sdk.Client) (resp *${this.action}Response, err error) {
	req.InitWithApiInfo("${this.product}", "${this.version}", "${this.action}", "${this.uriPattern}","${this.serviceCode}", "")
    req.Method = "${this.method}"

    resp = &${this.action}Response{}
	err = client.DoAction(req, resp)
	return
}
"""
            }
            "rpc" -> {
                codes += """
type ${this.action}Request ${withAnonymous(this.parameters.goType(subTypePrefix, true, sideCodes), "requests.RpcRequest")}

func (req *${this.action}Request) Invoke(client *sdk.Client) (resp *${this.action}Response, err error) {
	req.InitWithApiInfo("${this.product}", "${this.version}", "${this.action}", "${this.serviceCode}", "")
	resp = &${this.action}Response{}
	err = client.DoAction(req, resp)
	return
}
"""
            }
        }

        this.parameters.getAllDefinitions().forEach { s ->
            codes += """
type ${subTypePrefix}${s.key} ${s.value.goType(subTypePrefix, true, sideCodes)}
"""
        }

        codes += """
type ${this.action}Response ${withAnonymous(this.response.goType(subTypePrefix, false, sideCodes), "responses.BaseResponse")}
"""
        this.response.getAllDefinitions().forEach { s ->
            codes += """
type ${subTypePrefix}${s.key} ${s.value.goType(subTypePrefix, false, sideCodes)}
"""
        }

        sideCodes.forEach { s ->
            codes += s.value
        }

        return codes
    }
}