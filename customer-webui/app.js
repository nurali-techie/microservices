var http = require('http');
var fs = require('fs');

function onRequest(request, response) {
    response.writeHead(200, { 'Content-Type': 'text/html' });
    fs.readFile('./search.html', null, function (error, data) {
        if (error) {
            response.writeHead(404);
            response.write('file not found');
        } else {
            response.write(data);
        }
        response.end();
    });
}

http.createServer(onRequest).listen(2000);

console.log(`Running on http://localhost:2000`);
