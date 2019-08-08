//exm
var corolla = CarMaker.factory('Compact');
var solstice = CarMaker.factory('Convertible');
var cherokee = CarMaker.factory('SUV');
corolla.drive(); //结果为vroom，i have 4 doors

//实现
//父构造函数
function CarMaker() {}
CarMaker.prototype.driver = function() {
	return "vroom,i have "+this.doors+"doors";
}

//静态工厂方法
CarMaker.factory = function(type) {
	var constr = type,
		newcar;

	if(typeof CarMaker[constr] !== "function") {
		throw {
			name:"Error",
			message:constr + "doesn't exist"
		};
	}

	if(typeof CarMaker[constr].prototype.driver !== "function") {
		CarMaker[constr].prototype = new CarMaker();
	}

	newcar = new CarMaker[constr]();

	return newcar;
}

CarMaker.Compact = function() {
	this.doors = 4;
}

CarMaker.Convertible = function() {
	this.doors = 2
}

CarMaker.SUV = function() {
	this.doors = 24;
}