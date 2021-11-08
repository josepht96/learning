var request = require('request');
var url = process.argv.slice(2);

console.log('url: ', url);
const interval = setInterval(function() {
    request(url[0], function (error, response, body) {
        if (!error && response.statusCode == 200) {
            console.log(body) // Print the google web page.
        }
    })
}, 10000);
