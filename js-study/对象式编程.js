var Class = {
  create: function() {
    return function() {
      this.initialize.apply(this, arguments);
    }
  },
  $:function (id) {
    return "string" == typeof id ? document.getElementById(id) : id;
  }
}

var Extend = function(destination, source) {
    for (var property in source) {
        destination[property] = source[property];
    }
    return destination;
}

var Roll = Class.create();
Roll.prototype = {
	initialize: function(container,options){
		this.Container = $(container);
		this.SetOptions(options);
	},
	SetOptions:function(options){
		this.options = {

		};
		Extend(this.options,options || {})
	}

}



var Class = {
	  create: function() {
	    return function() {
	      this.initialize.apply(this, arguments);
	    }
	  }
}

var Extend = function(destination, source) {
    for (var property in source) {
        destination[property] = source[property];
    }
    return destination;
}

var Roll = Class.create();
Roll.prototype = {
	initialize: function(container,options){
		this.Container = $(container);
		this.Roll = [];
		
		this.SetOptions(options);
		this.Num = this.options.Num || 3;
		this.Time = this.options.Time || 1000;
		this.Step = this.options.Step || 10;

		this.RollAnima();
	},
	SetOptions:function(options){
		this.options = {
			Num:3,
			Time:3000,
			IntLine : 1
		};
		Extend(this.options,options || {})
	},
	RollAnima:function(){
		
	}
}

var roll = new Roll('showarea');