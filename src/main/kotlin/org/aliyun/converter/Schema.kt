package org.aliyun.converter

import com.github.javaparser.ast.body.TypeDeclaration
import com.github.javaparser.ast.type.Type
import converter.Inflection

class Schema(val type: String) {
    companion object {
        var fieldWithNumber = Regex("[0-9]")

        fun fromTypeDecl(typeDeclaration: TypeDeclaration<*>, imports: Map<String, List<String>>): Schema {
            val schema = Schema("object").named(typeDeclaration.name.asString())
            typeDeclaration.fields.forEach { field ->
                val name = field.asFieldDeclaration().getVariable(0).nameAsString
                schema.propOf(Schema.fromType(field.commonType, name, imports))
            }
            typeDeclaration.members.forEach { member ->
                if (member.isTypeDeclaration) {
                    schema.def(Schema.fromTypeDecl(member.asTypeDeclaration(), imports))
                }
            }
            return patchSchema(schema)
        }

        fun getRefName(importPaths: List<String>): String {
            val action = importPaths[importPaths.size - 2].replace(Regex("Request|Response"), "")
            return action + importPaths[importPaths.size - 1]
        }

        fun fromType(type: Type, name: String, imports: Map<String, List<String>>): Schema {
            if (type.childNodes.size == 2) {
                val generalTypeName = type.childNodes.get(0).toString()
                val itemName = type.childNodes.get(1).toString()
                val itemSchema = Schema(itemName).named(itemName)

                if (imports[itemName] != null) {
                    itemSchema.withRef(getRefName(imports[itemName]!!))
                }

                return Schema(generalTypeName)
                        .named(name)
                        .itemOf(itemSchema)
            }

            val t = type.asString()
            if (imports[t] != null) {
                return Schema("").named(name).withRef(getRefName(imports[t]!!))
            }
            return Schema(type.asString()).named(name)
        }

        fun assignDefinitions(definitions: MutableMap<String, Schema>, schema: Schema) {
            schema.definitions.forEach { def ->
                definitions[def.key] = def.value
                assignDefinitions(definitions, def.value)
            }

            if (schema.isArray()) {
                assignDefinitions(definitions, schema.items!!)
            }

            if (schema.isObject()) {
                if (schema.props != null) {
                    schema.props!!.forEach { s ->
                        assignDefinitions(definitions, s.value)
                    }
                }
            }
        }

        fun patchSchema(s: Schema): Schema {
            if (!s.isObject()) {
                return s
            }

            if (s.props == null) {
                return s
            }

            val finalSchema = Schema("object").named(s.name)
            finalSchema.props = mutableMapOf()

            s.definitions.forEach { d ->
                finalSchema.definitions[d.key] = d.value
            }

            s.props!!.forEach { p ->
                val name = p.key
                val currentSchema = p.value.clone()
                if (fieldWithNumber.matches(name)) {
                    val args = name.split(fieldWithNumber)
                    if (finalSchema.props!!.get(args[0]) == null) {
                        val schema = Schema("array")
                        if (args[1] != "") {
                            schema.itemOf(Schema("object"))
                        }
                        finalSchema.propOf(schema.named(args[0]))
                    }
                    if (args[1] != "") {
                        finalSchema.props!![args[0]]!!.items!!.propOf(currentSchema.named(args[1]))
                    } else {
                        finalSchema.props!![args[0]]!!.itemOf(currentSchema)
                    }
                } else {
                    finalSchema.propOf(p.value)
                }
            }

            return finalSchema
        }
    }

    var name = ""
    var ref = ""
    var items: Schema? = null
    var props: MutableMap<String, Schema>? = null
    val definitions = mutableMapOf<String, Schema>()

    fun clone(): Schema {
        val s = Schema(this.type)
        s.named(this.name)
        if (this.isArray()) {
            s.itemOf(this.items!!)
        }
        if (this.isObject()) {
            this.props!!.forEach { p ->
                s.propOf(p.value)
            }
        }
        this.definitions.forEach { d ->
            s.definitions[d.key] = d.value
        }
        return s
    }

    fun getAllDefinitions(): Map<String, Schema> {
        val definitions = mutableMapOf<String, Schema>()

        Schema.assignDefinitions(definitions, this)

        return definitions
    }

    fun isArray(): Boolean {
        return this.items != null
    }

    fun isObject(): Boolean {
        return this.type == "object" || this.props != null
    }

    fun withRef(ref: String): Schema {
        this.ref = ref
        return this
    }

    fun named(name: String): Schema {
        this.name = name
        return this
    }

    fun itemOf(s: Schema): Schema {
        this.items = s
        return this
    }

    fun propOf(s: Schema): Schema {
        if (this.props == null) {
            this.props = mutableMapOf()
        }
        this.props!![s.name] = s
        return this
    }

    fun def(s: Schema): Schema {
        this.definitions[s.name] = s
        return this
    }

    fun goType(prefix: String, sideCodes: MutableMap<String, String>): String {
        if (this.ref != "") {
            return this.ref
        }

        if (this.isObject()) {
            var t = """struct {
"""
            if (this.props != null) {
                this.props!!.forEach { prop ->
                    t += """${toUpperCamelCase(prop.key)} ${prop.value.goType(prefix, sideCodes)}
            """
                }
            }
            t += "}"
            return t
        }

        if (this.isArray()) {
            when (this.type) {
                "List" -> {

                    var itemName = this.items!!.name
                    if (itemName == "" || listOf("String", "Long", "Integer", "Float", "Boolean", "boolean").contains(itemName)) {
                        itemName = Inflection.singularize(this.name)
                    }

                    val listTypeName = prefix + toUpperCamelCase(itemName) + "List"
                    val subType = Schema("array").named(itemName).itemOf(this.items!!).goType(prefix, sideCodes)

                    sideCodes.set(listTypeName, """
                    type ${listTypeName} ${subType}

                    func (list *${listTypeName}) UnmarshalJSON(data []byte) error {
                        m := make(map[string]${subType})
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
                    """)

                    return listTypeName
                }
                "array" -> {
                    return "[]${this.items!!.goType(prefix, sideCodes)}"
                }
                else -> {
                    println("unsupported type ${this.type} ${this.items!!.type}")
                    return ""
                }
            }
        }

        when (this.type) {
            "String" -> {
                return "string"
            }
            "Long" -> {
                return "int64"
            }
            "Float" -> {
                return "float32"
            }
            "Integer" -> {
                return "int"
            }
            "boolean" -> {
                return "bool"
            }
            "Boolean" -> {
                return "bool"
            }
            else -> {
                return prefix + this.type
            }
        }
    }
}


