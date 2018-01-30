var res = require('./lib/resp.js');

session.set('user', null);
res.redirect('./login.js')