<script type="text/javascript">

function Sale(price) {
	this.price = price || 100;
}
Sale.prototype.getPrice = function() {
	return this.price;
}

Sale.decorators = {};

Sale.decorators.fedtax = {
	getPrice: function() {
		var price = this.uber.getPrice();
		price +=price *5 /100;
		return price;
	}
};

Sale.decorators.quebec = {
	getPrice:function() {
		var price = this.uber.getPrice();
		price +=price * 7.5/100;
		return price;
	}
};

Sale.decorators.cdn = {
	getPrice: function(){
		return "CND$" + this.uber.getPrice().toFixed(2);
	}
};

Sale.prototype.decorate = function(decorator) {
	var F = function(){},
		overrides = this.constructor.decorators[decorator],
		i,newobj;
	F.prototype = this;
	newobj = new F();
	newobj.uber = F.prototype;
	for(i in overrides) {
		console.log(i);
		if(overrides.hasOwnProperty(i)) {
			newobj[i] = overrides[i];
		}
	}
	return newobj;
};

var sale = new Sale(100);
sale = sale.decorate('fedtax');
console.log(sale.getPrice());
sale = sale.decorate('money');
console.log(sale.getPrice());



//装饰器列表实现

function Sal(price) {
	this.price = (price > 0) || 100;
	this.decorators_list = [];
}

Sale.decorators.fedtax = {
	getPrice: function(price) {
		return price + price*5/100;
	}
};

Sale.decorators.quebec = {
	getPrice: function(price) {
		return price + price*5/100;
	}
};
Sale.decorators.money = {
	getPrice: function (price) {
		return "$" +price.toFixed(2);
	}
};


Sale.prototype.decorate = function (decorator) {
	this.decorators_list.push(decorator);
}

Sale.prototype.getPrice = function () {
	var price = this.price,
		i,
		max = this.decorators_list.length,
		name;
	for (var i = 0; i < max; i++) {
		name = this.decorators_list[i];
		price = Sale.decorators[name].getPrice(price);
	}
	return price;
};
</script>