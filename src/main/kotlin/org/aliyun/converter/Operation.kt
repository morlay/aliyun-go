package org.aliyun.converter


class Operation(val action: String) {
    private var parameters = Schema("")
    private var response = Schema("")

    fun setParam(s: Schema) {
        this.parameters = s
    }

    fun setResp(s: Schema) {
        this.response = s
    }

    fun goClientMethod(name: String): String {
        val subTypePrefix = this.action

        val sideCodes = mutableMapOf<String, String>()

        var others = ""

        this.parameters.getAllDefinitions().forEach { s ->
            others += """
type ${subTypePrefix}${s.key} ${s.value.goType(subTypePrefix, sideCodes)}
"""
        }

        this.response.getAllDefinitions().forEach { s ->
            others += """
type ${subTypePrefix}${s.key} ${s.value.goType(subTypePrefix, sideCodes)}
"""
        }

        others += "type ${this.action}Args ${this.parameters.goType(subTypePrefix, sideCodes)}\n"
        others += "type ${this.action}Response ${this.response.goType(subTypePrefix, sideCodes)}\n"


        sideCodes.forEach { s ->
            others += s.value
        }

        return """

        func (c *${toUpperCamelCase(name)}Client) ${this.action}(req *${this.action}Args) (resp *${this.action}Response, err error) {
            resp = &${this.action}Response{}
            err = c.Request("${this.action}", req, resp)
            return
        }

        ${others}
"""
    }
}