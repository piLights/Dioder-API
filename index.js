var express    = require('express');
var bodyParser = require('body-parser');
var fs 		   = require('fs');
var app		   = express();
var jsonParser = bodyParser.json();

//Listen on the specified Port
var PORT = process.env.PORT || 8080;

var dioder = require('dioder')({
    color: '#000000',
    colorComponentPins: [18, 17, 4]
});

var animationDuration = 1500
	easing = 'easeInOutQuad';

app.get('/api/color/:color', function(req, res) {
	switch (req.params.color) {
		case 'red':
			dioder.animateTo('#f00', animationDuration, easing, function() {
				res.json({'Color': 'red'});
			});
			break;
		case 'green':
			dioder.animateTo('#0f0', animationDuration, easing, function() {
				res.json({'Color': 'green'});
			});			break;
		case 'blue':
			dioder.animateTo('#00f', animationDuration, easing, function() {
				res.json({'Color': 'blue'});
			});			break;
	}
});

app.get('/api/hex/:hexCode', function(req, res) {
	dioder.animateTo(req.params.hexCode, animationDuration, easing, function() {
		res.json({'ColorHex': req.params.hexCode});
	});
});

app.get('/api/rgb/:red/:green/:blue', function(req, res) {
	var color = [req.params.red, req.params.green, req.params.blue];
	dioder.animateTo(color, animationDuration, easing, function() {
		res.json({'Color': color})
	});
});

app.get('/api/off', function(req, res) {
	//See https://github.com/ffraenz/dioder/pull/1
	//dioder.turnOff();
	dioder.animateTo('#000', animationDuration, easing, function() {
		res.json({'Message': 'Turned off'});
	});
});

app.post('/api/setAnimationDuration', jsonParser, function (req, res) {
  if (!req.body) {
      return res.sendStatus(400);
  }
  animationDuration = req.body.AnimationDuration;
  res.json({'AnimationDuration': animationDuration});
})

app.get('/api/getAnimationDuration', function(req, res) {
	res.json({'AnimationDuration': animationDuration});
});

app.listen(PORT);
