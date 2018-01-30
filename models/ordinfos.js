var types = require('types');
var mtypes = require('./lib/types.js');
var db = require('./models/db.js');
var dbUtils = require('./models/dbUtils.js');
var utils = require('./lib/utils.js');

var db = db.getConn();


exports.count = function count(where) {
  var ret = null;
  if (where) {
    var r = dbUtils.exportKeyValues(where);
    var fields = r.keys;
    var values = r.values;
    var args = [];
    var arr = [];
    for (var i = 0; i < fields.length; i++) {
      var k = fields[i];
      var v = values[i];
      if (mtypes.isObject(v)) {
        if (v.$in) {
          var ins = v.$in;
          arr.push(k + ' in (' + dbUtils.makePlaceStr(ins.length) + ')');
          args = args.concat(ins);
        }
        continue;
      }
      arr.push(k + '=?');
      args.push(v);
    }
    var sql = 'select count(*) from goods where ' + arr.join(' and ');
    args.unshift(sql); // 将sql插入到最前面
    console.log(JSON.stringify(args));
    ret = db.query.apply(db, args);
  } else {
    ret = db.query('select count(*) from goods');
  }
  if (ret.err) {
    throw ret.err;
  }
  var rows = ret.value;
  while (rows.next()) {
    var n = types.newInt();
    var err = rows.scan(n);
    if (err) {
      throw err;
    }
    ret = types.intValue(n);
    break;
  }
  var err = rows.err();
  if (err) {
    throw err;
  }
  err = rows.close();
  if (err) {
    throw err;
  }
  return ret;
}

exports.add = function add(data) {
  var r = dbUtils.exportKeyValues(data);
  var sql = dbUtils.makeInsertSql('ordinfos', r.keys);
  return dbUtils.insert(db, sql, r.values);
}

exports.delByOid = function delByOid(oid) {
  return dbUtils.delete(db, 'delete from goods where oid=?', [oid]);
}

exports.getByOid = function getByOid(oid) {
  var gs = find({ oid: oid }, {});
  if (gs.length > 0) {
    return gs[0];
  }
  return null;
}

function find(where, opt) {
  var sql = 'select * from ordinfos';
  return dbUtils.find(sql, where, opt, query);
}

function query(sql, params) {
  var ret = [];
  params.unshift(sql);
  console.log('%j', params)
  var r = db.query.apply(db, params);
  if (r.err) {
    throw r.err;
  }
  var rows = r.value;
  while (rows.next()) {
    var id = types.newInt();
    var oid = types.newString();
    var ordId = types.newString();
    var userId = types.newString();
    var userName = types.newString();
    var address = types.newString();
    var payType = types.newString();
    var payState = types.newInt();
    var money = types.newString();
    var fuyan = types.newString();
    var created_at = types.newString();

    var err = rows.scan(id, oid, ordId, userId, userName, address, payType, payState, money, fuyan, created_at);
    if (err) {
      rows.close();
      throw err;
    }

    ret.push({
      id: types.intValue(id),
      oid: types.stringValue(oid),
      ordId: types.stringValue(ordId),
      userId: types.stringValue(userId),
      userName: types.stringValue(userName),
      address: types.stringValue(address),
      payType: types.stringValue(payType),
      payState: types.intValue(payState),
      money: types.stringValue(money),
      fuyan: types.stringValue(fuyan),
      created_at: types.stringValue(created_at),
    });
  }
  var err = rows.err();
  if (err) {
    rows.close();
    throw err;
  }
  err = rows.close();
  if (err) {
    throw err;
  }
  return ret;
}

function updateByOid(oid, data) {
  var r = dbUtils.exportKeyValues(data)
  var sql = dbUtils.makeUpdateSql('goods', r.keys, ['oid'])
  // 添加where条件值
  r.values.push(oid)
  return dbUtils.update(db, sql, r.values)
}

function decrGoodsNum(oid, num) {
  var sql = 'update goods set goods_number=goods_number-? where oid=? and goods_number>=?';
  var args = [num, oid, num];
  return dbUtils.update(db, sql, args);
}

exports.decrGoodsNum = decrGoodsNum;
exports.find = find;
exports.updateByOid = updateByOid;
