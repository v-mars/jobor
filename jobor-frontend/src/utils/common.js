/*邮箱验证正则*/
export function EmailReCheck(email){
  let reg = new RegExp("^[a-z0-9]+([._\\-]*[a-z0-9])*@([a-z0-9]+[-a-z0-9]*[a-z0-9]+.){1,63}[a-z0-9]+$"); //正则表达式
  if(email === ""){ //输入不能为空
    //alert("输入不能为空!");
    return false;
  }else if(!reg.test(email)){ //正则验证不通过，格式不对
    //alert("验证不通过!");
    return false;
  }else{
    //alert("通过！");
    return true;
  }
}

/*url验证正则*/
export  function validateUrl(str) {
  let oRegUrl = new RegExp();
  //aa.bb.com
  oRegUrl.compile("^[A-Za-z]+://[A-Za-z0-9-_]+\\.[A-Za-z0-9-_%&\?\/.=]+$");
  return oRegUrl.test(str);

}

// js保存内容至本地文件
export function saveShareContent (content, fileName) {
  let downLink = document.createElement('a')
  downLink.download = fileName
  //字符内容转换为blod地址
  let blob = new Blob([content])
  downLink.href = URL.createObjectURL(blob)
  // 链接插入到页面
  document.body.appendChild(downLink)
  downLink.click()
  // 移除下载链接
  document.body.removeChild(downLink)
}

export function templateFunc(str,data){
  return str.replace(/{{[\s+]?(\w+)[\s+]?}}/g, function (match, key) {
    return data[key];
  });
}
// let templateStr = 'i am {{name}} name {{name }} {{ name}},age {{age}},job {{job}} job2 {{ job }} ';
// let data = {
//   name:'xbd',
//   age:18,
//   job:'CTO'
// }
// console.log(this.templateFunc(templateStr,data));

/* 获取url参数 */
export function getParams(key) {
  var reg = new RegExp("(^|&)" + key + "=([^&]*)(&|$)");
  var r = window.location.search.substr(1).match(reg);
  // console.log('getParms',window.location.search.substr(1))
  if (r != null) {
    return unescape(r[2]);
  }
  return null;
}


// 下载文件
export function downloadFile(data, filename) {
  // 文件导出
  if (!data) {
    return
  }
  let url = window.URL.createObjectURL(new Blob([data]));
  let link = document.createElement('a');
  link.style.display = 'none';
  link.href = url;
  link.setAttribute('download', filename);

  document.body.appendChild(link);
  link.click()
}


// 通过浏览器下载文件
export function OpenDownload(url){
  window.open(url, '_blank');
}


/* 获取url参数 */
export function getUrlKey(name) {
  return decodeURIComponent((new RegExp('[?|&]' + name + '=' + '([^&;]+?)(&|#|;|$)').exec(location.href) || [, ""])[1].replace(/\+/g, '%20')) || null
}
// 当前时间换日期字符串
export function getCurTimeStr() {
  let now = new Date();
  let yy = now.getFullYear();      //年
  let mm = now.getMonth() + 1;     //月
  let dd = now.getDate();          //日
  let hh = now.getHours();         //时
  let ii = now.getMinutes();       //分
  let ss = now.getSeconds();       //秒
  let clock = String(yy);
  if(mm < 10) clock += "0";
  clock += String(mm);
  if(dd < 10) clock += "0";
  clock += String(dd);
  if(hh < 10) clock += "0";
  clock += "_"+String(hh);
  if (ii < 10) clock += '0';
  clock += String(ii);
  if (ss < 10) clock += '0';
  clock += String(ss);
  return clock
}

// 时间戳转日期
export function formatDate(now) {
  const year = now.getFullYear();    //取得4位数的年份
  const month = now.getMonth() + 1;  //取得日期中的月份，其中0表示1月，11表示12月
  const date = now.getDate();        //返回日期月份中的天数（1到31）
  const hour = now.getHours();       //返回日期中的小时数（0到23）
  const minute = now.getMinutes();   //返回日期中的分钟数（0到59）
  const second = now.getSeconds();   //返回日期中的秒数（0到59）
  return year+"-"+month+"-"+date+" "+hour+":"+minute+":"+second;
}
// 1575 4629 93
// var g=1551 3342 52272; //定义一个时间戳变量
// var d=new Date(g);   //创建一个指定的日期对象
// console.log(d);
// console.log(formatDate(d))


/*
roles: 为用户所拥有的权限;
permissionRoles: 为设定的权限
*/
export function hasPermission(roles, permissionRoles) {
  // console.log('roles:', roles)
  let admin_list = ['admin', 'administrators', 'super_admin', 'root']
  if (roles.some(role => admin_list.indexOf(role) >= 0)) return true // admin permission passed directly
  if (!permissionRoles||permissionRoles.length===0) return true
  return roles.some(role => permissionRoles.indexOf(role) >= 0)
}


// 生成yaml编辑页面  npm install vue2-ace-editor yaml
export function editorInit(editor) {
  require('brace/ext/language_tools') //language extension prerequsite...
  require('brace/mode/yaml')
  // require('brace/mode/json')
  require('brace/theme/monokai')
  require('brace/theme/merbivore')
  require('brace/theme/ambiance')
  require('brace/snippets/yaml') //snippet
  require('brace/snippets/text') //snippet
  // editor.setReadOnly(true)  // 设置只读
  editor.setOption("wrap", true)    // 设置自动换行
  editor.setFontSize(12)
  editor.setOptions({
    // minLines: 20, // 最小行数
    // maxLines: Infinity, //高度自适应
    enableBasicAutocompletion: true, //开启基础自动补全
    enableLiveAutocompletion: true, // 开启实时自动补全
    enableSnippets: true, // 启用代码片段
  })
}

export default {
  EmailReCheck,
  // saveShareContent,
  getParams,
  // getUrlKey,
  hasPermission,
  // editorInit,
  templateFunc,
  validateUrl,
}

