var cates = require("./models/cates.js");
var res = require('./lib/resp.js');
var session = require('./lib/session.js');
var utils = require('./lib/utils.js');
var users = require("./models/users.js");



var user = session.getUser();
if (!user) {
  return res.redirect('/routes/user/login.js');
}

if (request.getMethod() === 'GET') {
  var allData = cates.find();
  var data = {
    tree: cates.getTree(allData),
    user: user,
    cart: session.getCart(),
  };
  return res.render('./views/user/repwd.ejs', data);
}

if (request.getMethod() === 'POST') {
  var $oldpassword = request.formValue('old_password');
  var $password = request.formValue('new_password');
  var $repassword = request.formValue('comfirm_password');

  console.log('o:', $oldpassword)
  console.log('p:', $password)
  console.log('r:', $repassword)

  if (!/^[a-zA-Z0-9]{6,16}$/.test($password)) {
    return res.json({ error: "密码只能是6-16位英文字母+数字" });
  } else if ($password !== $repassword) {
    return res.json({ error: "两次密码不一致" });
  }

  if (user.password === utils.encodePassword($oldpassword, user.salt)) {
    //旧密码验证成功
    $password = utils.encodePassword($password, user.salt);
    var r = users.updateByOid(user.oid, { password: $password });
    console.log('%j', r);
    user.password = $password;
    session.set('user', JSON.stringify(user));
    res.json({ success: '修改成功' });
  } else {
    return res.json({ error: "原密码错误" });
  }

}