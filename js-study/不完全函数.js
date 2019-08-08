/**
*使用场景 ：正在调用同一个函数，并且传递对的参数绝大多数都是相同的时候
*/
function schonfinkelize(fn) {
	var slice = Array.prototype.slice,
	stored_args = slice.call(arguments,1);
	//console.log('stored_args',stored_args);
	return function () {
		var new_args = slice.call(arguments),
			args = stored_args.concat(new_args);
		return fn.apply(null,args);
	}

}


//普通 exp
function add(x,y) {
	return x+y;
}

//使用一
var newadd = schonfinkelize(add,5);
newadd(4);

//使用二
schonfinkelize(add,6)(7);

