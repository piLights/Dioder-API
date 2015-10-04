var express = require('express');
var fs 		= require('fs');
var app		= express();

//Listen on the specified Port
var PORT = process.env.PORT || 8080;

var dioder = require('dioder')({
    color: '#000000',
    colorComponentPins: [18, 17, 4]
});

app.get('/api/color/:color', function(req, res) {
	switch (req.params.color) {
		case 'red':
			dioder.animateTo('#f00');
			break;
		case 'green':
			dioder.animateTo('#0f0');
			break;
		case 'blue':
			dioder.animateTo('#00f');
			break;
	}
	res.json({'Color': req.params.color});
});

app.get('/api/hex/:hexCode', function(req, res) {
	dioder.animateTo(req.params.hexCode);
});

app.get('/api/rgb/:red/:green/:blue', function(req, res) {
	var color = [req.params.red, req.params.green, req.params.blue];
	dioder.animateTo(color);
});

app.get('/api/off', function(req, res) {
	//See https://github.com/ffraenz/dioder/pull/1
	//dioder.turnOff();
	dioder.animateTo('#000');
});

app.listen(PORT);
