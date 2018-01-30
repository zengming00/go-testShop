var types = require('types');
var sqlpkg = require('sql');
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
  return dbUtils.add('users', data);
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
  var sql = 'select * from users';
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
    var userName = types.newString();
    var phone = types.newString();
    var email = types.newString();
    var password = types.newString();
    var salt = types.newString();
    var address = sqlpkg.newNullString();
    var created_at = types.newString();

    var err = rows.scan(id, oid, userName, phone, email, password, salt, address, created_at);
    if (err) {
      rows.close();
      throw err;
    }

    address = sqlpkg.nullStringValue(address);
    ret.push({
      id: types.intValue(id),
      oid: types.stringValue(oid),
      userName: types.stringValue(userName),
      phone: types.stringValue(phone),
      email: types.stringValue(email),
      password: types.stringValue(password),
      salt: types.stringValue(salt),
      address: address.valid ? address.value : null,
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
  var sql = dbUtils.makeUpdateSql('users', r.keys, ['oid'])
  // 添加where条件值
  r.values.push(oid)
  return dbUtils.update(db, sql, r.values)
}

exports.find = find;
exports.updateByOid = updateByOid;
