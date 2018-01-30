
var res = require('./lib/resp.js');
var cates = require("./models/cates.js");
var goods = require("./models/goods.js");
var dbUtils = require('./models/dbUtils.js');
var utils = require('./lib/utils.js');
var Page = require('./lib/Page.class.js');
var Common = require('./lib/Common.js');
var session = require('./lib/session.js');


if (request.getMethod() === 'GET') {
  var query = utils.getQuery();
  var catid = query.cat_id;
  var allCates = cates.find();
  var childs = cates.getChildCates(allCates, catid);

  childs.unshift(catid);//将当前栏目也包含在内
  var gsCond = { cat_id: { $in: childs } };//查询条件
  var num = goods.count(gsCond)

  var req = {
    originalUrl: request.getUri(),
    query: utils.getQuery(),
  }
  var page = new Page(req, num, 9);//取得分页参数
  var docs = goods.find(gsCond, { skip: page.firstRow, limit: page.listRows })
  var data = {
    gs: docs,
    user: session.getUser(),
    cart: session.getCart(),
    history: session.get('history') ? JSON.parse(session.get('history')) : [],
    page: page.show(),
    tree: cates.getTree(allCates),
    family: cates.getFamily(allCates, catid)
  };
  res.render('./views/category.ejs', data);
}

