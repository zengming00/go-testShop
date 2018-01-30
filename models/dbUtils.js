var utils = require('./lib/utils.js');
var mtypes = require('./lib/types.js');

function dml(db, sql, args) {
  var r = db.prepare(sql)
  if (r.err) {
    throw r.err;
  }

  var stmt = r.value;
  r = stmt.exec.apply(stmt, args)
  if (r.err) {
    throw r.err;
  }


  var result = r.value;
  var r = result.lastInsertId()
  if (r.err) {
    throw r.err;
  }
  var lastInsertId = r.id;


  var r = result.rowsAffected()
  if (r.err) {
    throw r.err;
  }
  var rowsAffected = r.value;

  return {
    lastInsertId: lastInsertId,
    rowsAffected: rowsAffected,
  };
}

// 生成len个?长度的?,?字符串
function makePlaceStr(len) {
  var arr = [];
  for (var i = 0; i < len; i++) {
    arr.push('?');
  }
  return arr.join(',');
}

// 生成xx=?,yy=?字符串
function makeFieldPlaceStr(fields) {
  var arr = []
  for (var i = 0; i < fields.length; i++) {
    arr.push(fields[i] + '=?');
  }
  return arr.join(',');
}

function makeInsertSql(table, fields) {
  var sql = "insert into " + table + "(" + fields.join(',') + ") values(" + makePlaceStr(fields.length) + ")";
  return sql;
}

function makeUpdateSql(table, fields, wheres) {
  var sql = 'update ' + table + ' set ' + makeFieldPlaceStr(fields) + ' where ' + makeFieldPlaceStr(wheres);
  return sql;
}

function makeOid() {
  return Date.now().toString(16);
}

function exportKeyValues(data) {
  var values = [];
  var keys = [];
  for (var k in data) {
    keys.push(k);
    values.push(data[k]);
  }
  return {
    values: values,
    keys: keys,
  }
}

function buildAnd(fields, values) {
  var args = [];
  var conds = [];
  for (var i = 0; i < fields.length; i++) {
    var k = fields[i];
    var v = values[i];
    if (mtypes.isObject(v)) {
      if (v.$in) {
        var ins = v.$in;
        conds.push(k + ' in (' + makePlaceStr(ins.length) + ')');
        args = args.concat(ins);
      }
      continue;
    }
    conds.push(k + '=?');
    args.push(v);
  }
  return {
    and: conds.join(' and '),
    args: args,
  }
}

function buildWhere(where) {
  if (where.$or) {
    var vs = where.$or;
    var args = [];
    var conds = [];
    for (var i = 0; i < vs.length; i++) {
      var r = exportKeyValues(vs[i]);
      r = buildAnd(r.keys, r.values);
      conds.push('(' + r.and + ')');
      args = args.concat(r.args);
    }
    return {
      where: ' where ' + conds.join(' or '),
      args: args,
    };
  }
  var r = exportKeyValues(where);
  r = buildAnd(r.keys, r.values);
  return {
    where: ' where ' + r.and,
    args: r.args,
  };
}

function find(sql, where, opt, queryFunc) {
  var params = [];
  if (!utils.isEmptyObject(where)) {
    var r = dbUtils.buildWhere(where);
    sql += r.where;
    params = params.concat(r.args);
  }

  if (opt.sort && !utils.isEmptyObject(opt.sort)) {
    var r = dbUtils.exportKeyValues(opt.sort)
    sql += ' order by ' + r.keys[0] + ' ' + r.values[0];
  }
  if (utils.isNotNil(opt.skip) && utils.isNotNil(opt.limit)) {
    sql += ' limit ?,?';
    params.push(opt.skip)
    params.push(opt.limit);
  }
  return queryFunc(sql, params);
}


function add(table, data) {
  var r = exportKeyValues(data);
  var sql = makeInsertSql(table, r.keys);
  return dml(db, sql, r.values);
}

exports.add = add;
exports.find = find;
exports.buildWhere = buildWhere;
exports.exportKeyValues = exportKeyValues;
exports.makeOid = makeOid;
exports.makePlaceStr = makePlaceStr;
exports.insert = dml;
exports.delete = dml;
exports.update = dml;
exports.makeInsertSql = makeInsertSql;
exports.makeUpdateSql = makeUpdateSql;