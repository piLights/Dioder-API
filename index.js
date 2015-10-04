var express = require('express');
var fs 		= require('fs');
var app		= express();

//Listen on the specified Port
var PORT = process.env.PORT || 8080;

var dioder = require('dioder')({
    color: '#000000',
    colorComponentPins: [18, 17, 4]
});

app.get('/api/:color', function(req, res){
	switch (req.params.color) {
		case 'red':
			dioder.animateTo('#ff0000');
			break;
		case 'green':
			dioder.animateTo('#ff0000');
			break;
		case 'blue':
			dioder.animateTo('#ff0000');
			break;
	}
	res.json({'Color': req.params.color});
});

app.listen(PORT);
