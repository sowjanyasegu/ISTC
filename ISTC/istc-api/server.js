var app = require('./app');
const path = require('path');
var port = process.env.PORT || 3000;
app.get('/tea', (req, res) => {
res.sendFile(path.join(__dirname,'istcapp','./Index.html'));
});
var server = app.listen(port, function() {
  console.log('Express server listening on port ' + port);
});
