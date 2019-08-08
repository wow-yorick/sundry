//默认模式
function inherit(C,P) {
	C.prototype = new P()
}

//借用构造函数
function Child(a,b,c,d) {
	Parent.apply(this,arguments);
}

//借用和设置原型
function Child(a,b,c,d) {
	Parent.apply(this,arguments);
}
Child.prototype = new Parent();

//共享原型
function inherit(C,P) {
	C.prototype = P.prototype;
}

//临时构造函数
function inherit(C,P) {
	var F = function() {}
	F.prototype = P.prototype;
	C.prototype = new F();
}

//存储超类
function inherit(C,P) {
	var F = function() {}
	F.prototype = P.prototype;
	C.prototype = new F();
	C.prototype = P.prototype;
}

//重置构造函数指针
function inherit(C,P) {
	var F = function() {}
	F.prototype = P.prototype;
	C.prototype = new F();
	C.prototype = P.prototype;
	C.prototype.constructor = C;
}

//存储临时构造函数
var inherit = (function(){
	var F = function(){}
	return function(C,P){
		F.prototype = P.prototype;
		C.prototype = new F();
		C.prototype = P.prototype;
		C.prototype.constructor = C;
	}
}())

//js中模拟类使用
/**
*@parm1:object 被继承的父类
*@parm2:object字面量  新类的实现
*/
var Man = klass(null, {
	__construct: function(what) {
		console.log("Man's constructor");
		this.name = what;
	},
	getName:function() {
		return this.name;
	}
})

/**
*klass 实现
*/
var klass = function(Parent,props) {
	var Child,F,i;
	//新构造函数
	Child = function() {
		if(Child.uber && Child.uber.hasOwnProperty("__construct")) {
			Child.uber.__construct.apply(this,arguments);
		}
		if(Child.prototype.hasOwnProperty("__construct")) {
			Child.prototype.__construct.apply(this.arguments);
		}
	}
	//继承
	Parent = Parent || Object;
	F = function() {};
	F.prototype = Parent.prototype;
	Child.prototype = new F();
	Child.uber = Parent.prototype;
	Child.prototype.constructor = Child;

	//添加新方法
	for (i in props) {
		if(props.hasOwnProperty(i)) {
			Child.prototype[i] = props[i];
		}
	}

	return Child;
}


//原型继承
function object(o) {
	function F(){};
	F.prototype = O;
	return new F();
}

//复制性继承
function extend(parent,child) {
	var i;
	child = child || {};
	for(i in parent) {
		if(parent.hasOwnProperty(i)) {
			child[i] = parent[i];
		}
	}
	return child;
}

//深复制
function extendDeep(parent,child) {
	var i,
	toStr = Object.prototype.toString,
	astr = "[object Array]";

	child = child || {};

	for(i in parent) {
		if(parent.hasOwnProperty(i)) {
			if(typeof parent[i] === "object") {
				child[i] = (toStr.call(parent[i]) === astr) ? [] : {};
				extendDeep(parent[i],child[i]);
			}else {
				child[i] = parent[i];
			}
		}
	}

	return child;
}

//混入
function mix() {
	var arg,prop,child = {};
	for (arg = 0; arg < arguments.length; arg++) {
		for(prop in arguments[arg]) {
			if(arguments.hasOwnProperty(prop)) {
				child[prop] = arguments[arg][prop];
			}
		}
	}

	return child;
}

//call 和 apply 的使用（方法借用）
//call
notmyobj.dostuff.call(myobj,param1,p2,p3);
//apply
notmyobj.dostuff.apply(myobj,[param1,p2,p3]);

//借用数组方法
function f(){
	var args = [].slice.call(arguments,1,3);
	return args;
}

//绑定
function bind(o,m) {
	return function() {
		return m.apply(o,[].slice.call(arguments));
	}
}

//ECMAScript 5 中bind()
var newFunc = obj.someFunc.bind(myobj,1,2,3);

//bind 修复版
if(typeof Function.prototype.bind ==='undefined') {
	function.prototype.bind = function(thisArg) {
		var fn = this,
			slice = Array.prototype.slice,
			args = slice.call(arguments,1);

			return function() {
				return fn.apply(thisArg, args.concat(slice.call(arguments)));
			}
	}
}
//use
var twosay2 = one.say.bind(two);
twosay2('bonjour'); //结果为“bonjour，another obj”