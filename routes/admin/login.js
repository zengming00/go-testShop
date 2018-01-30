var res = require('./lib/resp.js');
var Verify = require('./lib/Verify.js');


if (request.getMethod() === 'GET') {
  res.render('./views/admin/login.ejs');
  return;
}

if (request.getMethod() === 'POST') {
  var username = request.formValue('username')
  var password = request.formValue('password')
  var yzm = request.formValue('yzm')

  var verifyOk = Verify.verify(yzm);

  if (username === 'admin' && password === 'admin123' && verifyOk) {
    session.set('isAdmin', true);
  }
  res.redirect('./index.js')
}
