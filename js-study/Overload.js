Jyo.Overload = function () {
    "use strict";

    // 重载所记录的数组
    var list = [];

    var fun = function () {
        /// <summary>调用重载函数</summary>

        var types = [],
            isMatch = false,
            type = "",
            input = null;

        for (var i = 0; i < list.length; i += 2) {
            if (list[i] != null && list[i] != "") {
                types = list[i].split(",");
            }
            if (types.length == 0 && arguments.length == 0) {
                isMatch = true;
                break;
            }
            if (types.length != arguments.length && types.length != 0 && types[types.length - 1].indexOf("...") < 0) continue;
            for (var cm = 0; cm < types.length; cm++) {
                type = types[cm] != "*" ? eval(types[cm]) : "*";
                input = arguments[cm];

                var typeName = type.toString().match(/function\s((\w|\w)+)/) || "";
                if (typeName != null && typeName != "") typeName = typeName[1];

                if (type == "...") {
                    if (!types[cm + 1]) isMatch = true;
                    else break;
                }
                else if (type == "*" || input === type || input instanceof type || typeof input == typeName.toLowerCase()) {
                    if (cm + 1 == types.length) isMatch = true;
                    continue;
                }
                break;
            }
            if (isMatch) break;
        }

        if (isMatch) {
            // 如果匹配了参数则调用
            return list[i + 1].apply(this, arguments);
        }

        throw new TypeError("Invalid parameter");
    };

    fun.add = function (types, callback) {
        /// <summary>添加重载函数</summary>
        /// <param name="types" type="String">重载所需要的函数类型表</param>
        /// <param name="callback" type="Function">重载所触发的函数</param>

        var typeList = null;
        if (types == null || types == "") {
            typeList = [];
        } else {
            typeList = types.split(",");
        }
        for (var i = 0; i < typeList.length; i++) {
            typeList[i] = typeList[i].trim();
            if (typeList[i] == "*" || typeList[i] == "...") continue;
            else if (eval("typeof " + typeList[i] + " == 'undefined'")) throw new ReferenceError(typeList[i] + " is not defined");
        }
        list.push(types, callback);
        return this;
    };

    return fun;
};