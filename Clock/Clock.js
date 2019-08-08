function Clock(canvas) {
    var _that = this;
    this.HAND_TRUNCATION = canvas.width/25;
    this.HOUR_HAND_TRUNCATION = canvas.width/10;
    this.canvas = canvas;
    this.context = canvas.getContext('2d');
    this.RADIUS = canvas.width/2 - this.MARGIN;
    this.HAND_RADIUS = this.RADIUS + this.NUMERAL_SPACING;
    this.context.font = this.FONT_HEIGHT +'px Arial';
    //console.log(this.context);
};
Clock.prototype = {
    FONT_HEIGHT : 15,
    MARGIN : 35,
    HAND_TRUNCATION :0,
    HOUR_HAND_TRUNCATION :0,
    NUMERAL_SPACING : 20,
    RADIUS : 0,
    HAND_RADIUS :0,
    context : null,
    canvas : null,
    init : function(){
    },
    drawCircle : function(){
        this.context.beginPath();
        this.context.arc(this.canvas.width/2, this.canvas.height/2, this.RADIUS, 0, Math.PI*2, true);
        this.context.stroke();
    },
    drawNumerals : function(){
        var _that = this;
        var numerals = [1,2,3,4,5,6,7,8,9,10,11,12],
            angle = 0,
            numeralWidth = 0;
        numerals.forEach(function(numeral) {
        //console.log(_that);
            angle = Math.PI/6 * (numeral - 3);
            numeralWidth = _that.context.measureText(numeral).width;
            _that.context.fillText(numeral,
            _that.canvas.width/2 + Math.cos(angle) * _that.HAND_RADIUS - numeralWidth/2,
            _that.canvas.height/2 +Math.sin(angle) * _that.HAND_RADIUS + _that.FONT_HEIGHT/3);
        });
    },
    drawCenter : function(){
        this.context.beginPath();
        this.context.arc(this.canvas.width/2, this.canvas.height/2.5, 0, Math.PI*2, true);
        this.context.fill();
    },
    drawHand : function(loc, isHour){
        var angle = (Math.PI*2) * (loc/60) - Math.PI/2,
            handRadius = isHour ? this.RADIUS -this.HAND_TRUNCATION - this.HOUR_HAND_TRUNCATION : this.RADIUS - this.HAND_TRUNCATION;
        this.context.moveTo(this.canvas.width/2, this.canvas.height/2);
        this.context.lineTo(this.canvas.width/2 + Math.cos(angle) * handRadius, this.canvas.height/2 + Math.sin(angle) * handRadius);
        this.context.stroke();
    },
    drawHands : function() {
        var date = new Date,
            hour = date.getHours();
        hour = hour > 12 ? hour -12 : hour;

        this.drawHand(hour*5 + (date.getMinutes()/60)*5, true, 0.5);
        this.drawHand(date.getMinutes(), false, 0.5);
        this.drawHand(date.getSeconds(), false, 0.2);
    },
    drawClock : function() {
        //console.log(this.context);
        this.context.clearRect(0, 0, this.canvas.width, this.canvas.height);
        this.drawCircle();
        this.drawCenter();
        this.drawHands();
        this.drawNumerals();
    }
};
