
function Universe() {
	var instance = this;

	this.start_time = 0;
	this.bang = 'big';

	Universe = function() {
		return instance;
	};
}

var uni1 = new Universe();
var uni2 = new Universe();
console.log(uni1 === uni2, 'Universe');



function Universe() {
	var instance;

	Universe = function() {
		return instance;
	};

	Universe.prototype = this;

	instance = new Universe();

	instance.constructor = Universe;

	instance.start_time = 0;
	instance.bang = 'big';

	return instance;
}

var Universe;
(function(){
	var instance;

	Universe = function Universe() {
		if(instance) {
			return instance;
		}

		instance = this;

		this.start_time = 0;
		this.bang = "big";
	};
	
}())

