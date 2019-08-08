var MYAPP = MYAPP || {};


MYAPP.namespace = function(ns_string) {
	var parts = ns_string.split('.'),
	parent = MYAPP,
	i;

	if(parts.[0] === 'MYAPP'){
		parts = parts.slice(1);
	}

	for(i = 0;i < parts.length; i++){
		if(typeof parent[parts[i]] === 'undefined') {
			parent[parts[i]] = {};
		}
		parent = parent[parts[i]];
	}
	return parent;
}

MYAPP.initInfoPage = function(){}

//exp
var module2 = MYAPP.namespace('module1.module2');
module2 === module1.module2;  //true