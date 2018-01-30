var cates = require("./models/cates.js");
var users = require("./models/users.js");
var res = require('./lib/resp.js');
var Verify = require('./lib/Verify.js');
var session = require('./lib/session.js');
var utils = require('./lib/utils.js');

if (request.getMethod() === 'GET') {
  var allData = cates.find();
  var data = {
    tree: cates.getTree(allData),
    user: session.getUser(),
    cart: session.getCart(),
  };
  return res.render('./views/user/login.ejs', data);
}

if (request.getMethod() === 'POST') {
  var yzm = request.formValue('yzm');
  if (!Verify.verify(yzm)) {
    console.log('验证码错误');
    return res.json({ error: '验证码错误' });
  }

  var $username = request.formValue('username');
  var $password = request.formValue('password');

  //允许用户名、手机号、邮箱登录
  var cond = { $or: [{ userName: $username }, { phone: $username }, { email: $username }] };
  var us = users.find(cond, {});
  if (us.length > 0) {
    var doc = us[0];
    if (doc.password === utils.encodePassword($password, doc.salt)) {
      session.set('user', JSON.stringify(doc));
      res.json({ success: '登录成功！' });
    } else {
      res.json({ error: '密码错误！' });
    }
  } else {
    res.json({ error: '不存在的用户' });
  }
}