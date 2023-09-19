import{m as E,S as Oe,i as F,bj as On,bo as Rn,d as ie,at as Mt,aF as dn,az as Vn,h as d,a0 as Ft,bv as co,as as Gn,bw as cn,aw as un,bx as At,by as _t,bz as Un,bA as Bt,ay as De,bB as Ke,bC as Bo,bD as Nn,bE as Nt,bF as Et,bG as uo,bH as Lt,bI as fo,bJ as ho,bK as on,bL as Dt,bM as vo,av as Ht,bN as Kt,bO as qn,bP as Wt,bQ as jt,bR as Vt,bS as No,bq as Gt,bT as Ut,ax as qt,bU as Zt,bV as Yt,a as ke,e as B,v as H,t as ee,g as We,u as se,ah as Eo,j as ze,N as Lo,l as be,bW as Jt,G as oe,c as qe,ad as Xt,n as Re,r as xe,T as fn,f as U,w as pe,aE as Zn,z as te,p as fe,aY as Le,D as Qt,aj as er,aH as nr,am as or,bd as kn,bX as tr,y as Yn,bY as rn,aq as Do,bi as tn,bZ as En,aJ as po,b_ as rr,an as hn,aZ as Jn,b$ as Xn,c0 as Ho,c1 as Ko,F as Qn,be as Wo,al as jo,A as ln,c2 as go,aU as ir,c3 as lr,q as Vo,a4 as ar,B as ce,b as X,aG as sr,ai as dr,k as vn,aK as bo,c4 as cr,c5 as ur,ak as fr,c6 as hr,ag as vr,bs as Go}from "./index-a4639770.js";import{c as pr,a as gr,i as eo,V as br,h as He,F as mr,d as yr,u as _e,e as no,f as oo,j as to,k as wr}from "./FocusDetector-f0e5f0c6.js";function xr(e){switch(typeof e){case"string":return e||void 0;case"number":return String(e);default:return}}function Cr(e){return n=>{n?e.value=n.$el:e.value=null}}function $n(e){const n=e.filter(o=>o!==void 0);if(n.length!==0)return n.length===1?n[0]: o=>{e.forEach(t=>{t&&t(o)})}}let Tn;function Sr(){return Tn===void 0&&(Tn=navigator.userAgent.includes("Node.js")||navigator.userAgent.includes("jsdom")),Tn}function Pr(e, n, o){if(!n)return e;const t=E(e.value);let r=null;return Oe(e, i=>{r!==null&&window.clearTimeout(r),i===!0?o&&!o.value?t.value=!0:r=window.setTimeout(()=>{t.value=!0},n):t.value=!1}),t}function Uo(e, n){return F(()=>{for(const o of n)if(e[o]!==void 0)return e[o];return e[n[n.length-1]]})}const Ee="@@mmoContext",Or={mounted(e, {value:n}){e[Ee]={handler:void 0},typeof n=="function"&&(e[Ee].handler=n,On("mousemoveoutside",e,n))},updated(e, {value:n}){const o=e[Ee];typeof n=="function"?o.handler?o.handler!==n&&(Rn("mousemoveoutside",e,o.handler),o.handler=n,On("mousemoveoutside",e,n)):(e[Ee].handler=n,On("mousemoveoutside",e,n)):o.handler&&(Rn("mousemoveoutside",e,o.handler),o.handler=void 0)},unmounted(e){const{handler:n}=e[Ee];n&&Rn("mousemoveoutside",e,n),e[Ee].handler=void 0}},Rr=Or,Fe="v-hidden",kr=gr("[v-hidden]",{display:"none!important"}),mo=ie({name:"Overflow",props:{getCounter:Function,getTail:Function,updateCounter:Function,onUpdateOverflow:Function},setup(e, {slots:n}){const o=E(null),t=E(null);function r(){const{value:l}=o,{getCounter:a,getTail:s}=e;let c;if(a!==void 0?c=a():c=t.value,!l||!c)return;c.hasAttribute(Fe)&&c.removeAttribute(Fe);const{children:u}=l,h=l.offsetWidth,g=[],b=n.tail?s==null?void 0:s():null;let v=b?b.offsetWidth:0,m=!1;const R=l.children.length-(n.tail?1:0);for(let O=0; O<R-1; ++O){if(O<0)continue;const D=u[O];if(m){D.hasAttribute(Fe)||D.setAttribute(Fe,"");continue}else D.hasAttribute(Fe)&&D.removeAttribute(Fe);const w=D.offsetWidth;if(v+=w,g[O]=w,v>h){const{updateCounter:S}=e;for(let _=O; _>=0; --_){const T=R-1-_;S!==void 0?S(T):c.textContent=`${T}`;const N=c.offsetWidth;if(v-=g[_],v+N<=h||_===0){m=!0,O=_-1,b&&(O===-1?(b.style.maxWidth=`${h-N}px`,b.style.boxSizing="border-box"):b.style.maxWidth="");break}}}}const{onUpdateOverflow:y}=e;m?y!==void 0&&y(!0):(y!==void 0&&y(!1),c.setAttribute(Fe,""))}const i=Mt();return kr.mount({id:"vueuc/overflow",head:!0,anchorMetaName:pr,ssr:i}),dn(r),{selfRef:o,counterRef:t,sync:r}},render(){const{$slots:e}=this;return Vn(this.sync),d("div",{class:"v-overflow",ref:"selfRef"},[Ft(e,"default"),e.counter?e.counter():d("span",{style:{display:"inline-block"},ref:"counterRef"}),e.tail?e.tail():null])}});function qo(e, n){n&&(dn(()=>{const{value:o}=e;o&&co.registerHandler(o,n)}),Gn(()=>{const{value:o}=e;o&&co.unregisterHandler(o)}))}var $r=cn(un,"WeakMap");const Ln=$r;var Tr=At(Object.keys,Object);const zr=Tr;var Ir=Object.prototype,Mr=Ir.hasOwnProperty;function Fr(e){if(!_t(e))return zr(e);var n=[];for(var o in Object(e))Mr.call(e,o)&&o!="constructor"&&n.push(o);return n}function ro(e){return Un(e)?Bt(e):Fr(e)}function Ar(e, n){for(var o=-1,t=n.length,r=e.length; ++o<t;)e[r+o]=n[o];return e}function _r(e, n){for(var o=-1,t=e==null?0:e.length,r=0,i=[]; ++o<t;){var l=e[o];n(l,o,e)&&(i[r++]=l)}return i}function Br(){return[]}var Nr=Object.prototype,Er=Nr.propertyIsEnumerable,yo=Object.getOwnPropertySymbols,Lr=yo?function(e){return e==null?[]:(e=Object(e),_r(yo(e),function(n){return Er.call(e,n)}))}:Br;const Dr=Lr;function Hr(e, n, o){var t=n(e);return De(e)?t:Ar(t,o(e))}function wo(e){return Hr(e,ro,Dr)}var Kr=cn(un,"DataView");const Dn=Kr;var Wr=cn(un,"Promise");const Hn=Wr;var jr=cn(un,"Set");const Kn=jr;var xo="[object Map]",Vr="[object Object]",Co="[object Promise]",So="[object Set]",Po="[object WeakMap]",Oo="[object DataView]",Gr=Ke(Dn),Ur=Ke(Nn),qr=Ke(Hn),Zr=Ke(Kn),Yr=Ke(Ln),Ae=Bo;(Dn&&Ae(new Dn(new ArrayBuffer(1)))!=Oo||Nn&&Ae(new Nn)!=xo||Hn&&Ae(Hn.resolve())!=Co||Kn&&Ae(new Kn)!=So||Ln&&Ae(new Ln)!=Po)&&(Ae=function(e){var n=Bo(e),o=n==Vr?e.constructor:void 0,t=o?Ke(o):"";if(t)switch(t){case Gr:return Oo;case Ur:return xo;case qr:return Co;case Zr:return So;case Yr:return Po}return n});const Ro=Ae;function Jr(e, n){for(var o=-1,t=e==null?0:e.length; ++o<t;)if(n(e[o],o,e))return!0;return!1}var Xr=1,Qr=2;function Zo(e, n, o, t, r, i){var l=o&Xr,a=e.length,s=n.length;if(a!=s&&!(l&&s>a))return!1;var c=i.get(e),u=i.get(n);if(c&&u)return c==n&&u==e;var h=-1,g=!0,b=o&Qr?new Nt:void 0;for(i.set(e,n),i.set(n,e); ++h<a;){var v=e[h],m=n[h];if(t)var R=l?t(m,v,h,n,e,i):t(v,m,h,e,n,i);if(R!==void 0){if(R)continue;g=!1;break}if(b){if(!Jr(n,function(y, O){if(!Et(b,O)&&(v===y||r(v,y,o,t,i)))return b.push(O)})){g=!1;break}}else if(!(v===m||r(v,m,o,t,i))){g=!1;break}}return i.delete(e),i.delete(n),g}function ei(e){var n=-1,o=Array(e.size);return e.forEach(function(t, r){o[++n]=[r,t]}),o}function ni(e){var n=-1,o=Array(e.size);return e.forEach(function(t){o[++n]=t}),o}var oi=1,ti=2,ri="[object Boolean]",ii="[object Date]",li="[object Error]",ai="[object Map]",si="[object Number]",di="[object RegExp]",ci="[object Set]",ui="[object String]",fi="[object Symbol]",hi="[object ArrayBuffer]",vi="[object DataView]",ko=uo?uo.prototype:void 0,zn=ko?ko.valueOf:void 0;function pi(e, n, o, t, r, i, l){switch(o){case vi:if(e.byteLength!=n.byteLength||e.byteOffset!=n.byteOffset)return!1;e=e.buffer,n=n.buffer;case hi:return!(e.byteLength!=n.byteLength||!i(new fo(e),new fo(n)));case ri:case ii:case si:return Lt(+e,+n);case li:return e.name==n.name&&e.message==n.message;case di:case ui:return e==n+"";case ai:var a=ei;case ci:var s=t&oi;if(a||(a=ni),e.size!=n.size&&!s)return!1;var c=l.get(e);if(c)return c==n;t|=ti,l.set(e,n);var u=Zo(a(e),a(n),t,r,i,l);return l.delete(e),u;case fi:if(zn)return zn.call(e)==zn.call(n)}return!1}var gi=1,bi=Object.prototype,mi=bi.hasOwnProperty;function yi(e, n, o, t, r, i){var l=o&gi,a=wo(e),s=a.length,c=wo(n),u=c.length;if(s!=u&&!l)return!1;for(var h=s; h--;){var g=a[h];if(!(l?g in n:mi.call(n,g)))return!1}var b=i.get(e),v=i.get(n);if(b&&v)return b==n&&v==e;var m=!0;i.set(e,n),i.set(n,e);for(var R=l; ++h<s;){g=a[h];var y=e[g],O=n[g];if(t)var D=l?t(O,y,g,n,e,i):t(y,O,g,e,n,i);if(!(D===void 0?y===O||r(y,O,o,t,i):D)){m=!1;break}R||(R=g=="constructor")}if(m&&!R){var w=e.constructor,S=n.constructor;w!=S&&"constructor"in e&&"constructor"in n&&!(typeof w=="function"&&w instanceof w&&typeof S=="function"&&S instanceof S)&&(m=!1)}return i.delete(e),i.delete(n),m}var wi=1,$o="[object Arguments]",To="[object Array]",nn="[object Object]",xi=Object.prototype,zo=xi.hasOwnProperty;function Ci(e, n, o, t, r, i){var l=De(e),a=De(n),s=l?To:Ro(e),c=a?To:Ro(n);s=s==$o?nn:s,c=c==$o?nn:c;var u=s==nn,h=c==nn,g=s==c;if(g&&ho(e)){if(!ho(n))return!1;l=!0,u=!1}if(g&&!u)return i||(i=new on),l||Dt(e)?Zo(e,n,o,t,r,i):pi(e,n,s,o,t,r,i);if(!(o&wi)){var b=u&&zo.call(e,"__wrapped__"),v=h&&zo.call(n,"__wrapped__");if(b||v){var m=b?e.value():e,R=v?n.value():n;return i||(i=new on),r(m,R,o,t,i)}}return g?(i||(i=new on),yi(e,n,o,t,r,i)):!1}function io(e, n, o, t, r){return e===n?!0:e==null||n==null||!vo(e)&&!vo(n)?e!==e&&n!==n:Ci(e,n,o,t,io,r)}var Si=1,Pi=2;function Oi(e, n, o, t){var r=o.length,i=r,l=!t;if(e==null)return!i;for(e=Object(e); r--;){var a=o[r];if(l&&a[2]?a[1]!==e[a[0]]:!(a[0]in e))return!1}for(; ++r<i;){a=o[r];var s=a[0],c=e[s],u=a[1];if(l&&a[2]){if(c===void 0&&!(s in e))return!1}else{var h=new on;if(t)var g=t(c,u,s,e,n,h);if(!(g===void 0?io(u,c,Si|Pi,t,h):g))return!1}}return!0}function Yo(e){return e===e&&!Ht(e)}function Ri(e){for(var n=ro(e),o=n.length; o--;){var t=n[o],r=e[t];n[o]=[t,r,Yo(r)]}return n}function Jo(e, n){return function(o){return o==null?!1:o[e]===n&&(n!==void 0||e in Object(o))}}function ki(e){var n=Ri(e);return n.length==1&&n[0][2]?Jo(n[0][0],n[0][1]):function(o){return o===e||Oi(o,e,n)}}function $i(e, n){return e!=null&&n in Object(e)}function Ti(e, n, o){n=Kt(n,e);for(var t=-1,r=n.length,i=!1; ++t<r;){var l=qn(n[t]);if(!(i=e!=null&&o(e,l)))break;e=e[l]}return i||++t!=r?i:(r=e==null?0:e.length,!!r&&Wt(r)&&jt(l,r)&&(De(e)||Vt(e)))}function zi(e, n){return e!=null&&Ti(e,n,$i)}var Ii=1,Mi=2;function Fi(e, n){return No(e)&&Yo(n)?Jo(qn(e),n):function(o){var t=Gt(o,e);return t===void 0&&t===n?zi(o,e):io(n,t,Ii|Mi)}}function Ai(e){return function(n){return n==null?void 0:n[e]}}function _i(e){return function(n){return Ut(n,e)}}function Bi(e){return No(e)?Ai(qn(e)):_i(e)}function Ni(e){return typeof e=="function"?e:e==null?qt:typeof e=="object"?De(e)?Fi(e[0],e[1]):ki(e):Bi(e)}function Ei(e, n){return e&&Zt(e,n,ro)}function Li(e, n){return function(o, t){if(o==null)return o;if(!Un(o))return e(o,t);for(var r=o.length,i=n?r:-1,l=Object(o); (n?i--:++i<r)&&t(l[i],i,l)!==!1;);return o}}var Di=Li(Ei);const Hi=Di;function Ki(e, n){var o=-1,t=Un(e)?Array(e.length):[];return Hi(e,function(r, i, l){t[++o]=n(r,i,l)}),t}function Wi(e, n){var o=De(e)?Yt:Ki;return o(e,Ni(n))}const ji=ie({name:"Checkmark",render(){return d("svg",{xmlns:"http://www.w3.org/2000/svg",viewBox:"0 0 16 16"},d("g",{fill:"none"},d("path",{d:"M14.046 3.486a.75.75 0 0 1-.032 1.06l-7.93 7.474a.85.85 0 0 1-1.188-.022l-2.68-2.72a.75.75 0 1 1 1.068-1.053l2.234 2.267l7.468-7.038a.75.75 0 0 1 1.06.032z",fill:"currentColor"})))}}),Vi=ie({name:"ChevronRight",render(){return d("svg",{viewBox:"0 0 16 16",fill:"none",xmlns:"http://www.w3.org/2000/svg"},d("path",{d:"M5.64645 3.14645C5.45118 3.34171 5.45118 3.65829 5.64645 3.85355L9.79289 8L5.64645 12.1464C5.45118 12.3417 5.45118 12.6583 5.64645 12.8536C5.84171 13.0488 6.15829 13.0488 6.35355 12.8536L10.8536 8.35355C11.0488 8.15829 11.0488 7.84171 10.8536 7.64645L6.35355 3.14645C6.15829 2.95118 5.84171 2.95118 5.64645 3.14645Z",fill:"currentColor"}))}}),Gi=ie({name:"Empty",render(){return d("svg",{viewBox:"0 0 28 28",fill:"none",xmlns:"http://www.w3.org/2000/svg"},d("path",{d:"M26 7.5C26 11.0899 23.0899 14 19.5 14C15.9101 14 13 11.0899 13 7.5C13 3.91015 15.9101 1 19.5 1C23.0899 1 26 3.91015 26 7.5ZM16.8536 4.14645C16.6583 3.95118 16.3417 3.95118 16.1464 4.14645C15.9512 4.34171 15.9512 4.65829 16.1464 4.85355L18.7929 7.5L16.1464 10.1464C15.9512 10.3417 15.9512 10.6583 16.1464 10.8536C16.3417 11.0488 16.6583 11.0488 16.8536 10.8536L19.5 8.20711L22.1464 10.8536C22.3417 11.0488 22.6583 11.0488 22.8536 10.8536C23.0488 10.6583 23.0488 10.3417 22.8536 10.1464L20.2071 7.5L22.8536 4.85355C23.0488 4.65829 23.0488 4.34171 22.8536 4.14645C22.6583 3.95118 22.3417 3.95118 22.1464 4.14645L19.5 6.79289L16.8536 4.14645Z",fill:"currentColor"}),d("path",{d:"M25 22.75V12.5991C24.5572 13.0765 24.053 13.4961 23.5 13.8454V16H17.5L17.3982 16.0068C17.0322 16.0565 16.75 16.3703 16.75 16.75C16.75 18.2688 15.5188 19.5 14 19.5C12.4812 19.5 11.25 18.2688 11.25 16.75L11.2432 16.6482C11.1935 16.2822 10.8797 16 10.5 16H4.5V7.25C4.5 6.2835 5.2835 5.5 6.25 5.5H12.2696C12.4146 4.97463 12.6153 4.47237 12.865 4H6.25C4.45507 4 3 5.45507 3 7.25V22.75C3 24.5449 4.45507 26 6.25 26H21.75C23.5449 26 25 24.5449 25 22.75ZM4.5 22.75V17.5H9.81597L9.85751 17.7041C10.2905 19.5919 11.9808 21 14 21L14.215 20.9947C16.2095 20.8953 17.842 19.4209 18.184 17.5H23.5V22.75C23.5 23.7165 22.7165 24.5 21.75 24.5H6.25C5.2835 24.5 4.5 23.7165 4.5 22.75Z",fill:"currentColor"}))}});function Io(e){return Array.isArray(e)?e:[e]}const Wn={STOP:"STOP"};function Xo(e, n){const o=n(e);e.children!==void 0&&o!==Wn.STOP&&e.children.forEach(t=>Xo(t,n))}function Ui(e, n={}){const{preserveGroup:o=!1}=n,t=[],r=o? l=>{l.isLeaf||(t.push(l.key),i(l.children))}: l=>{l.isLeaf||(l.isGroup||t.push(l.key),i(l.children))};function i(l){l.forEach(r)}return i(e),t}function qi(e, n){const{isLeaf:o}=e;return o!==void 0?o:!n(e)}function Zi(e){return e.children}function Yi(e){return e.key}function Ji(){return!1}function Xi(e, n){const{isLeaf:o}=e;return!(o===!1&&!Array.isArray(n(e)))}function Qi(e){return e.disabled===!0}function el(e, n){return e.isLeaf===!1&&!Array.isArray(n(e))}function In(e){var n;return e==null?[]:Array.isArray(e)?e:(n=e.checkedKeys)!==null&&n!==void 0?n:[]}function Mn(e){var n;return e==null||Array.isArray(e)?[]:(n=e.indeterminateKeys)!==null&&n!==void 0?n:[]}function nl(e, n){const o=new Set(e);return n.forEach(t=>{o.has(t)||o.add(t)}),Array.from(o)}function ol(e, n){const o=new Set(e);return n.forEach(t=>{o.has(t)&&o.delete(t)}),Array.from(o)}function tl(e){return(e==null?void 0:e.type)==="group"}function rl(e){const n=new Map;return e.forEach((o, t)=>{n.set(o.key,t)}), o=>{var t;return(t=n.get(o))!==null&&t!==void 0?t:null}}class il extends Error{constructor(){super(),this.message="SubtreeNotLoadedError: checking a subtree whose required nodes are not fully loaded."}}function ll(e, n, o, t){return an(n.concat(e),o,t,!1)}function al(e, n){const o=new Set;return e.forEach(t=>{const r=n.treeNodeMap.get(t);if(r!==void 0){let i=r.parent;for(; i!==null&&!(i.disabled||o.has(i.key));)o.add(i.key),i=i.parent}}),o}function sl(e, n, o, t){const r=an(n,o,t,!1),i=an(e,o,t,!0),l=al(e,o),a=[];return r.forEach(s=>{(i.has(s)||l.has(s))&&a.push(s)}),a.forEach(s=>r.delete(s)),r}function Fn(e, n){const{checkedKeys:o,keysToCheck:t,keysToUncheck:r,indeterminateKeys:i,cascade:l,leafOnly:a,checkStrategy:s,allowNotLoaded:c}=e;if(!l)return t!==void 0?{checkedKeys:nl(o,t),indeterminateKeys:Array.from(i)}:r!==void 0?{checkedKeys:ol(o,r),indeterminateKeys:Array.from(i)}:{checkedKeys:Array.from(o),indeterminateKeys:Array.from(i)};const{levelTreeNodeMap:u}=n;let h;r!==void 0?h=sl(r,o,n,c):t!==void 0?h=ll(t,o,n,c):h=an(o,n,c,!1);const g=s==="parent",b=s==="child"||a,v=h,m=new Set,R=Math.max.apply(null,Array.from(u.keys()));for(let y=R; y>=0; y-=1){const O=y===0,D=u.get(y);for(const w of D){if(w.isLeaf)continue;const{key:S,shallowLoaded:_}=w;if(b&&_&&w.children.forEach(k=>{!k.disabled&&!k.isLeaf&&k.shallowLoaded&&v.has(k.key)&&v.delete(k.key)}),w.disabled||!_)continue;let T=!0,N=!1,L=!0;for(const k of w.children){const G=k.key;if(!k.disabled){if(L&&(L=!1),v.has(G))N=!0;else if(m.has(G)){N=!0,T=!1;break}else if(T=!1,N)break}}T&&!L?(g&&w.children.forEach(k=>{!k.disabled&&v.has(k.key)&&v.delete(k.key)}),v.add(S)):N&&m.add(S),O&&b&&v.has(S)&&v.delete(S)}}return{checkedKeys:Array.from(v),indeterminateKeys:Array.from(m)}}function an(e, n, o, t){const{treeNodeMap:r,getChildren:i}=n,l=new Set,a=new Set(e);return e.forEach(s=>{const c=r.get(s);c!==void 0&&Xo(c, u=>{if(u.disabled)return Wn.STOP;const{key:h}=u;if(!l.has(h)&&(l.add(h),a.add(h),el(u.rawNode,i))){if(t)return Wn.STOP;if(!o)throw new il}})}),a}function dl(e, {includeGroup:n=!1,includeSelf:o=!0}, t){var r;const i=t.treeNodeMap;let l=e==null?null:(r=i.get(e))!==null&&r!==void 0?r:null;const a={keyPath:[],treeNodePath:[],treeNode:l};if(l!=null&&l.ignored)return a.treeNode=null,a;for(; l;)!l.ignored&&(n||!l.isGroup)&&a.treeNodePath.push(l),l=l.parent;return a.treeNodePath.reverse(),o||a.treeNodePath.pop(),a.keyPath=a.treeNodePath.map(s=>s.key),a}function cl(e){if(e.length===0)return null;const n=e[0];return n.isGroup||n.ignored||n.disabled?n.getNext():n}function ul(e, n){const o=e.siblings,t=o.length,{index:r}=e;return n?o[(r+1)%t]:r===o.length-1?null:o[r+1]}function Mo(e, n, {loop:o=!1,includeDisabled:t=!1}={}){const r=n==="prev"?fl:ul,i={reverse:n==="prev"};let l=!1,a=null;function s(c){if(c!==null){if(c===e){if(!l)l=!0;else if(!e.disabled&&!e.isGroup){a=e;return}}else if((!c.disabled||t)&&!c.ignored&&!c.isGroup){a=c;return}if(c.isGroup){const u=lo(c,i);u!==null?a=u:s(r(c,o))}else{const u=r(c,!1);if(u!==null)s(u);else{const h=hl(c);h!=null&&h.isGroup?s(r(h,o)):o&&s(r(c,!0))}}}}return s(e),a}function fl(e, n){const o=e.siblings,t=o.length,{index:r}=e;return n?o[(r-1+t)%t]:r===0?null:o[r-1]}function hl(e){return e.parent}function lo(e, n={}){const{reverse:o=!1}=n,{children:t}=e;if(t){const{length:r}=t,i=o?r-1:0,l=o?-1:r,a=o?-1:1;for(let s=i; s!==l; s+=a){const c=t[s];if(!c.disabled&&!c.ignored)if(c.isGroup){const u=lo(c,n);if(u!==null)return u}else return c}}return null}const vl={getChild(){return this.ignored?null:lo(this)},getParent(){const{parent:e}=this;return e!=null&&e.isGroup?e.getParent():e},getNext(e={}){return Mo(this,"next",e)},getPrev(e={}){return Mo(this,"prev",e)}};function pl(e, n){const o=n?new Set(n):void 0,t=[];function r(i){i.forEach(l=>{t.push(l),!(l.isLeaf||!l.children||l.ignored)&&(l.isGroup||o===void 0||o.has(l.key))&&r(l.children)})}return r(e),t}function gl(e, n){const o=e.key;for(; n;){if(n.key===o)return!0;n=n.parent}return!1}function Qo(e, n, o, t, r, i=null, l=0){const a=[];return e.forEach((s, c)=>{var u;const h=Object.create(t);if(h.rawNode=s,h.siblings=a,h.level=l,h.index=c,h.isFirstChild=c===0,h.isLastChild=c+1===e.length,h.parent=i,!h.ignored){const g=r(s);Array.isArray(g)&&(h.children=Qo(g,n,o,t,r,h,l+1))}a.push(h),n.set(h.key,h),o.has(l)||o.set(l,[]),(u=o.get(l))===null||u===void 0||u.push(h)}),a}function et(e, n={}){var o;const t=new Map,r=new Map,{getDisabled:i=Qi,getIgnored:l=Ji,getIsGroup:a=tl,getKey:s=Yi}=n,c=(o=n.getChildren)!==null&&o!==void 0?o:Zi,u=n.ignoreEmptyChildren? w=>{const S=c(w);return Array.isArray(S)?S.length?S:null:S}:c,h=Object.assign({get key(){return s(this.rawNode)},get disabled(){return i(this.rawNode)},get isGroup(){return a(this.rawNode)},get isLeaf(){return qi(this.rawNode,u)},get shallowLoaded(){return Xi(this.rawNode,u)},get ignored(){return l(this.rawNode)},contains(w){return gl(this,w)}},vl),g=Qo(e,t,r,h,u);function b(w){if(w==null)return null;const S=t.get(w);return S&&!S.isGroup&&!S.ignored?S:null}function v(w){if(w==null)return null;const S=t.get(w);return S&&!S.ignored?S:null}function m(w, S){const _=v(w);return _?_.getPrev(S):null}function R(w, S){const _=v(w);return _?_.getNext(S):null}function y(w){const S=v(w);return S?S.getParent():null}function O(w){const S=v(w);return S?S.getChild():null}const D={treeNodes:g,treeNodeMap:t,levelTreeNodeMap:r,maxLevel:Math.max(...r.keys()),getChildren:u,getFlattenedNodes(w){return pl(g,w)},getNode:b,getPrev:m,getNext:R,getParent:y,getChild:O,getFirstAvailableNode(){return cl(g)},getPath(w, S={}){return dl(w,S,D)},getCheckedKeys(w, S={}){const{cascade:_=!0,leafOnly:T=!1,checkStrategy:N="all",allowNotLoaded:L=!1}=S;return Fn({checkedKeys:In(w),indeterminateKeys:Mn(w),cascade:_,leafOnly:T,checkStrategy:N,allowNotLoaded:L},D)},check(w, S, _={}){const{cascade:T=!0,leafOnly:N=!1,checkStrategy:L="all",allowNotLoaded:k=!1}=_;return Fn({checkedKeys:In(S),indeterminateKeys:Mn(S),keysToCheck:w==null?[]:Io(w),cascade:T,leafOnly:N,checkStrategy:L,allowNotLoaded:k},D)},uncheck(w, S, _={}){const{cascade:T=!0,leafOnly:N=!1,checkStrategy:L="all",allowNotLoaded:k=!1}=_;return Fn({checkedKeys:In(S),indeterminateKeys:Mn(S),keysToUncheck:w==null?[]:Io(w),cascade:T,leafOnly:N,checkStrategy:L,allowNotLoaded:k},D)},getNonLeafKeys(w={}){return Ui(g,w)}};return D}const bl={iconSizeSmall:"34px",iconSizeMedium:"40px",iconSizeLarge:"46px",iconSizeHuge:"52px"},ml= e=>{const{textColorDisabled:n,iconColor:o,textColor2:t,fontSizeSmall:r,fontSizeMedium:i,fontSizeLarge:l,fontSizeHuge:a}=e;return Object.assign(Object.assign({},bl),{fontSizeSmall:r,fontSizeMedium:i,fontSizeLarge:l,fontSizeHuge:a,textColor:n,iconColor:o,extraTextColor:t})},yl={name:"Empty",common:ke,self:ml},nt=yl,wl=B("empty",`
 display: flex;
 flex-direction: column;
 align-items: center;
 font-size: var(--n-font-size);
`,[H("icon",`
 width: var(--n-icon-size);
 height: var(--n-icon-size);
 font-size: var(--n-icon-size);
 line-height: var(--n-icon-size);
 color: var(--n-icon-color);
 transition:
 color .3s var(--n-bezier);
 `,[ee("+",[H("description",`
 margin-top: 8px;
 `)])]),H("description",`
 transition: color .3s var(--n-bezier);
 color: var(--n-text-color);
 `),H("extra",`
 text-align: center;
 transition: color .3s var(--n-bezier);
 margin-top: 12px;
 color: var(--n-extra-text-color);
 `)]),xl=Object.assign(Object.assign({},se.props),{description:String,showDescription:{type:Boolean,default:!0},showIcon:{type:Boolean,default:!0},size:{type:String,default:"medium"},renderIcon:Function}),Cl=ie({name:"Empty",props:xl,setup(e){const{mergedClsPrefixRef:n,inlineThemeDisabled:o}=We(e),t=se("Empty","-empty",wl,nt,e,n),{localeRef:r}=Eo("Empty"),i=be(Jt,null),l=F(()=>{var u,h,g;return(u=e.description)!==null&&u!==void 0?u:(g=(h=i==null?void 0:i.mergedComponentPropsRef.value)===null||h===void 0?void 0:h.Empty)===null||g===void 0?void 0:g.description}),a=F(()=>{var u,h;return((h=(u=i==null?void 0:i.mergedComponentPropsRef.value)===null||u===void 0?void 0:u.Empty)===null||h===void 0?void 0:h.renderIcon)||(()=>d(Gi,null))}),s=F(()=>{const{size:u}=e,{common:{cubicBezierEaseInOut:h},self:{[oe("iconSize",u)]:g,[oe("fontSize",u)]:b,textColor:v,iconColor:m,extraTextColor:R}}=t.value;return{"--n-icon-size":g,"--n-font-size":b,"--n-bezier":h,"--n-text-color":v,"--n-icon-color":m,"--n-extra-text-color":R}}),c=o?ze("empty",F(()=>{let u="";const{size:h}=e;return u+=h[0],u}),s,e):void 0;return{mergedClsPrefix:n,mergedRenderIcon:a,localizedDescription:F(()=>l.value||r.value.description),cssVars:o?void 0:s,themeClass:c==null?void 0:c.themeClass,onRender:c==null?void 0:c.onRender}},render(){const{$slots:e,mergedClsPrefix:n,onRender:o}=this;return o==null||o(),d("div",{class:[`${n}-empty`,this.themeClass],style:this.cssVars},this.showIcon?d("div",{class:`${n}-empty__icon`},e.icon?e.icon():d(Lo,{clsPrefix:n},{default:this.mergedRenderIcon})):null,this.showDescription?d("div",{class:`${n}-empty__description`},e.default?e.default():this.localizedDescription):null,e.extra?d("div",{class:`${n}-empty__extra`},e.extra()):null)}}),Sl={height:"calc(var(--n-option-height) * 7.6)",paddingSmall:"4px 0",paddingMedium:"4px 0",paddingLarge:"4px 0",paddingHuge:"4px 0",optionPaddingSmall:"0 12px",optionPaddingMedium:"0 12px",optionPaddingLarge:"0 12px",optionPaddingHuge:"0 12px",loadingSize:"18px"},Pl=e=>{const{borderRadius:n,popoverColor:o,textColor3:t,dividerColor:r,textColor2:i,primaryColorPressed:l,textColorDisabled:a,primaryColor:s,opacityDisabled:c,hoverColor:u,fontSizeSmall:h,fontSizeMedium:g,fontSizeLarge:b,fontSizeHuge:v,heightSmall:m,heightMedium:R,heightLarge:y,heightHuge:O}=e;return Object.assign(Object.assign({},Sl),{optionFontSizeSmall:h,optionFontSizeMedium:g,optionFontSizeLarge:b,optionFontSizeHuge:v,optionHeightSmall:m,optionHeightMedium:R,optionHeightLarge:y,optionHeightHuge:O,borderRadius:n,color:o,groupHeaderTextColor:t,actionDividerColor:r,optionTextColor:i,optionTextColorPressed:l,optionTextColorDisabled:a,optionTextColorActive:s,optionOpacityDisabled:c,optionCheckColor:s,optionColorPending:u,optionColorActive:"rgba(0, 0, 0, 0)",optionColorActivePending:u,actionTextColor:i,loadingColor:s})},Ol=qe({name:"InternalSelectMenu",common:ke,peers:{Scrollbar:Xt,Empty:nt},self:Pl}),ot=Ol;function Rl(e,n){return d(fn,{name:"fade-in-scale-up-transition"},{default:()=>e?d(Lo,{clsPrefix:n,class:`${n}-base-select-option__check`},{default:()=>d(ji)}):null})}const Fo=ie({name:"NBaseSelectOption",props:{clsPrefix:{type:String,required:!0},tmNode:{type:Object,required:!0}},setup(e){const{valueRef:n,pendingTmNodeRef:o,multipleRef:t,valueSetRef:r,renderLabelRef:i,renderOptionRef:l,labelFieldRef:a,valueFieldRef:s,showCheckmarkRef:c,nodePropsRef:u,handleOptionClick:h,handleOptionMouseEnter:g}=be(eo),b=Re(()=>{const{value:y}=o;return y?e.tmNode.key===y.key:!1});function v(y){const{tmNode:O}=e;O.disabled||h(y,O)}function m(y){const{tmNode:O}=e;O.disabled||g(y,O)}function R(y){const{tmNode:O}=e,{value:D}=b;O.disabled||D||g(y,O)}return{multiple:t,isGrouped:Re(()=>{const{tmNode:y}=e,{parent:O}=y;return O&&O.rawNode.type==="group"}),showCheckmark:c,nodeProps:u,isPending:b,isSelected:Re(()=>{const{value:y}=n,{value:O}=t;if(y===null)return!1;const D=e.tmNode.rawNode[s.value];if(O){const{value:w}=r;return w.has(D)}else return y===D}),labelField:a,renderLabel:i,renderOption:l,handleMouseMove:R,handleMouseEnter:m,handleClick:v}},render(){const{clsPrefix:e,tmNode:{rawNode:n},isSelected:o,isPending:t,isGrouped:r,showCheckmark:i,nodeProps:l,renderOption:a,renderLabel:s,handleClick:c,handleMouseEnter:u,handleMouseMove:h}=this,g=Rl(o,e),b=s?[s(n,o),i&&g]:[xe(n[this.labelField],n,o),i&&g],v=l==null?void 0:l(n),m=d("div",Object.assign({},v,{class:[`${e}-base-select-option`,n.class,v==null?void 0:v.class,{[`${e}-base-select-option--disabled`]:n.disabled,[`${e}-base-select-option--selected`]:o,[`${e}-base-select-option--grouped`]:r,[`${e}-base-select-option--pending`]:t,[`${e}-base-select-option--show-checkmark`]:i}],style:[(v==null?void 0:v.style)||"",n.style||""],onClick:$n([c,v==null?void 0:v.onClick]),onMouseenter:$n([u,v==null?void 0:v.onMouseenter]),onMousemove:$n([h,v==null?void 0:v.onMousemove])}),d("div",{class:`${e}-base-select-option__content`},b));return n.render?n.render({node:m,option:n,selected:o}):a?a({node:m,option:n,selected:o}):m}}),Ao=ie({name:"NBaseSelectGroupHeader",props:{clsPrefix:{type:String,required:!0},tmNode:{type:Object,required:!0}},setup(){const{renderLabelRef:e,renderOptionRef:n,labelFieldRef:o,nodePropsRef:t}=be(eo);return{labelField:o,nodeProps:t,renderLabel:e,renderOption:n}},render(){const{clsPrefix:e,renderLabel:n,renderOption:o,nodeProps:t,tmNode:{rawNode:r}}=this,i=t==null?void 0:t(r),l=n?n(r,!1):xe(r[this.labelField],r,!1),a=d("div",Object.assign({},i,{class:[`${e}-base-select-group-header`,i==null?void 0:i.class]}),l);return r.render?r.render({node:a,option:r}):o?o({node:a,option:r,selected:!1}):a}}),kl=B("base-select-menu",`
 line-height: 1.5;
 outline: none;
 z-index: 0;
 position: relative;
 border-radius: var(--n-border-radius);
 transition:
 background-color .3s var(--n-bezier),
 box-shadow .3s var(--n-bezier);
 background-color: var(--n-color);
`,[B("scrollbar",`
 max-height: var(--n-height);
 `),B("virtual-list",`
 max-height: var(--n-height);
 `),B("base-select-option",`
 min-height: var(--n-option-height);
 font-size: var(--n-option-font-size);
 display: flex;
 align-items: center;
 `,[H("content",`
 z-index: 1;
 white-space: nowrap;
 text-overflow: ellipsis;
 overflow: hidden;
 `)]),B("base-select-group-header",`
 min-height: var(--n-option-height);
 font-size: .93em;
 display: flex;
 align-items: center;
 `),B("base-select-menu-option-wrapper",`
 position: relative;
 width: 100%;
 `),H("loading, empty",`
 display: flex;
 padding: 12px 32px;
 flex: 1;
 justify-content: center;
 `),H("loading",`
 color: var(--n-loading-color);
 font-size: var(--n-loading-size);
 `),H("action",`
 padding: 8px var(--n-option-padding-left);
 font-size: var(--n-option-font-size);
 transition: 
 color .3s var(--n-bezier),
 border-color .3s var(--n-bezier);
 border-top: 1px solid var(--n-action-divider-color);
 color: var(--n-action-text-color);
 `),B("base-select-group-header",`
 position: relative;
 cursor: default;
 padding: var(--n-option-padding);
 color: var(--n-group-header-text-color);
 `),B("base-select-option",`
 cursor: pointer;
 position: relative;
 padding: var(--n-option-padding);
 transition:
 color .3s var(--n-bezier),
 opacity .3s var(--n-bezier);
 box-sizing: border-box;
 color: var(--n-option-text-color);
 opacity: 1;
 `,[U("show-checkmark",`
 padding-right: calc(var(--n-option-padding-right) + 20px);
 `),ee("&::before",`
 content: "";
 position: absolute;
 left: 4px;
 right: 4px;
 top: 0;
 bottom: 0;
 border-radius: var(--n-border-radius);
 transition: background-color .3s var(--n-bezier);
 `),ee("&:active",`
 color: var(--n-option-text-color-pressed);
 `),U("grouped",`
 padding-left: calc(var(--n-option-padding-left) * 1.5);
 `),U("pending",[ee("&::before",`
 background-color: var(--n-option-color-pending);
 `)]),U("selected",`
 color: var(--n-option-text-color-active);
 `,[ee("&::before",`
 background-color: var(--n-option-color-active);
 `),U("pending",[ee("&::before",`
 background-color: var(--n-option-color-active-pending);
 `)])]),U("disabled",`
 cursor: not-allowed;
 `,[pe("selected",`
 color: var(--n-option-text-color-disabled);
 `),U("selected",`
 opacity: var(--n-option-opacity-disabled);
 `)]),H("check",`
 font-size: 16px;
 position: absolute;
 right: calc(var(--n-option-padding-right) - 4px);
 top: calc(50% - 7px);
 color: var(--n-option-check-color);
 transition: color .3s var(--n-bezier);
 `,[Zn({enterScale:"0.5"})])])]),$l=ie({name:"InternalSelectMenu",props:Object.assign(Object.assign({},se.props),{clsPrefix:{type:String,required:!0},scrollable:{type:Boolean,default:!0},treeMate:{type:Object,required:!0},multiple:Boolean,size:{type:String,default:"medium"},value:{type:[String,Number,Array],default:null},autoPending:Boolean,virtualScroll:{type:Boolean,default:!0},show:{type:Boolean,default:!0},labelField:{type:String,default:"label"},valueField:{type:String,default:"value"},loading:Boolean,focusable:Boolean,renderLabel:Function,renderOption:Function,nodeProps:Function,showCheckmark:{type:Boolean,default:!0},onMousedown:Function,onScroll:Function,onFocus:Function,onBlur:Function,onKeyup:Function,onKeydown:Function,onTabOut:Function,onMouseenter:Function,onMouseleave:Function,onResize:Function,resetMenuOnOptionsChange:{type:Boolean,default:!0},inlineThemeDisabled:Boolean,onToggle:Function}),setup(e){const n=se("InternalSelectMenu","-internal-select-menu",kl,ot,e,te(e,"clsPrefix")),o=E(null),t=E(null),r=E(null),i=F(()=>e.treeMate.getFlattenedNodes()),l=F(()=>rl(i.value)),a=E(null);function s(){const{treeMate:p}=e;let P=null;const{value:ne}=e;ne===null?P=p.getFirstAvailableNode():(e.multiple?P=p.getNode((ne||[])[(ne||[]).length-1]):P=p.getNode(ne),(!P||P.disabled)&&(P=p.getFirstAvailableNode())),G(P||null)}function c(){const{value:p}=a;p&&!e.treeMate.getNode(p.key)&&(a.value=null)}let u;Oe(()=>e.show,p=>{p?u=Oe(()=>e.treeMate,()=>{e.resetMenuOnOptionsChange?(e.autoPending?s():c(),Vn(z)):c()},{immediate:!0}):u==null||u()},{immediate:!0}),Gn(()=>{u==null||u()});const h=F(()=>or(n.value.self[oe("optionHeight",e.size)])),g=F(()=>kn(n.value.self[oe("padding",e.size)])),b=F(()=>e.multiple&&Array.isArray(e.value)?new Set(e.value):new Set),v=F(()=>{const p=i.value;return p&&p.length===0});function m(p){const{onToggle:P}=e;P&&P(p)}function R(p){const{onScroll:P}=e;P&&P(p)}function y(p){var P;(P=r.value)===null||P===void 0||P.sync(),R(p)}function O(){var p;(p=r.value)===null||p===void 0||p.sync()}function D(){const{value:p}=a;return p||null}function w(p,P){P.disabled||G(P,!1)}function S(p,P){P.disabled||m(P)}function _(p){var P;He(p,"action")||(P=e.onKeyup)===null||P===void 0||P.call(e,p)}function T(p){var P;He(p,"action")||(P=e.onKeydown)===null||P===void 0||P.call(e,p)}function N(p){var P;(P=e.onMousedown)===null||P===void 0||P.call(e,p),!e.focusable&&p.preventDefault()}function L(){const{value:p}=a;p&&G(p.getNext({loop:!0}),!0)}function k(){const{value:p}=a;p&&G(p.getPrev({loop:!0}),!0)}function G(p,P=!1){a.value=p,P&&z()}function z(){var p,P;const ne=a.value;if(!ne)return;const le=l.value(ne.key);le!==null&&(e.virtualScroll?(p=t.value)===null||p===void 0||p.scrollTo({index:le}):(P=r.value)===null||P===void 0||P.scrollTo({index:le,elSize:h.value}))}function W(p){var P,ne;!((P=o.value)===null||P===void 0)&&P.contains(p.target)&&((ne=e.onFocus)===null||ne===void 0||ne.call(e,p))}function x(p){var P,ne;!((P=o.value)===null||P===void 0)&&P.contains(p.relatedTarget)||(ne=e.onBlur)===null||ne===void 0||ne.call(e,p)}fe(eo,{handleOptionMouseEnter:w,handleOptionClick:S,valueSetRef:b,pendingTmNodeRef:a,nodePropsRef:te(e,"nodeProps"),showCheckmarkRef:te(e,"showCheckmark"),multipleRef:te(e,"multiple"),valueRef:te(e,"value"),renderLabelRef:te(e,"renderLabel"),renderOptionRef:te(e,"renderOption"),labelFieldRef:te(e,"labelField"),valueFieldRef:te(e,"valueField")}),fe(yr,o),dn(()=>{const{value:p}=r;p&&p.sync()});const I=F(()=>{const{size:p}=e,{common:{cubicBezierEaseInOut:P},self:{height:ne,borderRadius:le,color:me,groupHeaderTextColor:Ce,actionDividerColor:Se,optionTextColorPressed:ye,optionTextColor:Q,optionTextColorDisabled:ge,optionTextColorActive:ue,optionOpacityDisabled:$e,optionCheckColor:we,actionTextColor:je,optionColorPending:Ie,optionColorActive:Me,loadingColor:Ve,loadingSize:Ge,optionColorActivePending:Ue,[oe("optionFontSize",p)]:Be,[oe("optionHeight",p)]:Ne,[oe("optionPadding",p)]:he}}=n.value;return{"--n-height":ne,"--n-action-divider-color":Se,"--n-action-text-color":je,"--n-bezier":P,"--n-border-radius":le,"--n-color":me,"--n-option-font-size":Be,"--n-group-header-text-color":Ce,"--n-option-check-color":we,"--n-option-color-pending":Ie,"--n-option-color-active":Me,"--n-option-color-active-pending":Ue,"--n-option-height":Ne,"--n-option-opacity-disabled":$e,"--n-option-text-color":Q,"--n-option-text-color-active":ue,"--n-option-text-color-disabled":ge,"--n-option-text-color-pressed":ye,"--n-option-padding":he,"--n-option-padding-left":kn(he,"left"),"--n-option-padding-right":kn(he,"right"),"--n-loading-color":Ve,"--n-loading-size":Ge}}),{inlineThemeDisabled:j}=e,A=j?ze("internal-select-menu",F(()=>e.size[0]),I,e):void 0,V={selfRef:o,next:L,prev:k,getPendingTmNode:D};return qo(o,e.onResize),Object.assign({mergedTheme:n,virtualListRef:t,scrollbarRef:r,itemSize:h,padding:g,flattenedNodes:i,empty:v,virtualListContainer(){const{value:p}=t;return p==null?void 0:p.listElRef},virtualListContent(){const{value:p}=t;return p==null?void 0:p.itemsElRef},doScroll:R,handleFocusin:W,handleFocusout:x,handleKeyUp:_,handleKeyDown:T,handleMouseDown:N,handleVirtualListResize:O,handleVirtualListScroll:y,cssVars:j?void 0:I,themeClass:A==null?void 0:A.themeClass,onRender:A==null?void 0:A.onRender},V)},render(){const{$slots:e,virtualScroll:n,clsPrefix:o,mergedTheme:t,themeClass:r,onRender:i}=this;return i==null||i(),d("div",{ref:"selfRef",tabindex:this.focusable?0:-1,class:[`${o}-base-select-menu`,r,this.multiple&&`${o}-base-select-menu--multiple`],style:this.cssVars,onFocusin:this.handleFocusin,onFocusout:this.handleFocusout,onKeyup:this.handleKeyUp,onKeydown:this.handleKeyDown,onMousedown:this.handleMouseDown,onMouseenter:this.onMouseenter,onMouseleave:this.onMouseleave},this.loading?d("div",{class:`${o}-base-select-menu__loading`},d(Qt,{clsPrefix:o,strokeWidth:20})):this.empty?d("div",{class:`${o}-base-select-menu__empty`,"data-empty":!0},nr(e.empty,()=>[d(Cl,{theme:t.peers.Empty,themeOverrides:t.peerOverrides.Empty})])):d(er,{ref:"scrollbarRef",theme:t.peers.Scrollbar,themeOverrides:t.peerOverrides.Scrollbar,scrollable:this.scrollable,container:n?this.virtualListContainer:void 0,content:n?this.virtualListContent:void 0,onScroll:n?void 0:this.doScroll},{default:()=>n?d(br,{ref:"virtualListRef",class:`${o}-virtual-list`,items:this.flattenedNodes,itemSize:this.itemSize,showScrollbar:!1,paddingTop:this.padding.top,paddingBottom:this.padding.bottom,onResize:this.handleVirtualListResize,onScroll:this.handleVirtualListScroll,itemResizable:!0},{default:({item:l})=>l.isGroup?d(Ao,{key:l.key,clsPrefix:o,tmNode:l}):l.ignored?null:d(Fo,{clsPrefix:o,key:l.key,tmNode:l})}):d("div",{class:`${o}-base-select-menu-option-wrapper`,style:{paddingTop:this.padding.top,paddingBottom:this.padding.bottom}},this.flattenedNodes.map(l=>l.isGroup?d(Ao,{key:l.key,clsPrefix:o,tmNode:l}):d(Fo,{clsPrefix:o,key:l.key,tmNode:l})))}),Le(e.action,l=>l&&[d("div",{class:`${o}-base-select-menu__action`,"data-action":!0,key:"action"},l),d(mr,{onFocus:this.onTabOut,key:"focus-detector"})]))}}),Tl={space:"6px",spaceArrow:"10px",arrowOffset:"10px",arrowOffsetVertical:"10px",arrowHeight:"6px",padding:"8px 14px"},zl=e=>{const{boxShadow2:n,popoverColor:o,textColor2:t,borderRadius:r,fontSize:i,dividerColor:l}=e;return Object.assign(Object.assign({},Tl),{fontSize:i,borderRadius:r,color:o,dividerColor:l,textColor:t,boxShadow:n})},Il={name:"Popover",common:ke,self:zl},pn=Il,An={top:"bottom",bottom:"top",left:"right",right:"left"},ae="var(--n-arrow-height) * 1.414",Ml=ee([B("popover",`
 transition:
 box-shadow .3s var(--n-bezier),
 background-color .3s var(--n-bezier),
 color .3s var(--n-bezier);
 position: relative;
 font-size: var(--n-font-size);
 color: var(--n-text-color);
 box-shadow: var(--n-box-shadow);
 word-break: break-word;
 `,[ee(">",[B("scrollbar",`
 height: inherit;
 max-height: inherit;
 `)]),pe("raw",`
 background-color: var(--n-color);
 border-radius: var(--n-border-radius);
 `,[pe("scrollable",[pe("show-header-or-footer","padding: var(--n-padding);")])]),H("header",`
 padding: var(--n-padding);
 border-bottom: 1px solid var(--n-divider-color);
 transition: border-color .3s var(--n-bezier);
 `),H("footer",`
 padding: var(--n-padding);
 border-top: 1px solid var(--n-divider-color);
 transition: border-color .3s var(--n-bezier);
 `),U("scrollable, show-header-or-footer",[H("content",`
 padding: var(--n-padding);
 `)])]),B("popover-shared",`
 transform-origin: inherit;
 `,[B("popover-arrow-wrapper",`
 position: absolute;
 overflow: hidden;
 pointer-events: none;
 `,[B("popover-arrow",`
 transition: background-color .3s var(--n-bezier);
 position: absolute;
 display: block;
 width: calc(${ae});
 height: calc(${ae});
 box-shadow: 0 0 8px 0 rgba(0, 0, 0, .12);
 transform: rotate(45deg);
 background-color: var(--n-color);
 pointer-events: all;
 `)]),ee("&.popover-transition-enter-from, &.popover-transition-leave-to",`
 opacity: 0;
 transform: scale(.85);
 `),ee("&.popover-transition-enter-to, &.popover-transition-leave-from",`
 transform: scale(1);
 opacity: 1;
 `),ee("&.popover-transition-enter-active",`
 transition:
 box-shadow .3s var(--n-bezier),
 background-color .3s var(--n-bezier),
 color .3s var(--n-bezier),
 opacity .15s var(--n-bezier-ease-out),
 transform .15s var(--n-bezier-ease-out);
 `),ee("&.popover-transition-leave-active",`
 transition:
 box-shadow .3s var(--n-bezier),
 background-color .3s var(--n-bezier),
 color .3s var(--n-bezier),
 opacity .15s var(--n-bezier-ease-in),
 transform .15s var(--n-bezier-ease-in);
 `)]),ve("top-start",`
 top: calc(${ae} / -2);
 left: calc(${Pe("top-start")} - var(--v-offset-left));
 `),ve("top",`
 top: calc(${ae} / -2);
 transform: translateX(calc(${ae} / -2)) rotate(45deg);
 left: 50%;
 `),ve("top-end",`
 top: calc(${ae} / -2);
 right: calc(${Pe("top-end")} + var(--v-offset-left));
 `),ve("bottom-start",`
 bottom: calc(${ae} / -2);
 left: calc(${Pe("bottom-start")} - var(--v-offset-left));
 `),ve("bottom",`
 bottom: calc(${ae} / -2);
 transform: translateX(calc(${ae} / -2)) rotate(45deg);
 left: 50%;
 `),ve("bottom-end",`
 bottom: calc(${ae} / -2);
 right: calc(${Pe("bottom-end")} + var(--v-offset-left));
 `),ve("left-start",`
 left: calc(${ae} / -2);
 top: calc(${Pe("left-start")} - var(--v-offset-top));
 `),ve("left",`
 left: calc(${ae} / -2);
 transform: translateY(calc(${ae} / -2)) rotate(45deg);
 top: 50%;
 `),ve("left-end",`
 left: calc(${ae} / -2);
 bottom: calc(${Pe("left-end")} + var(--v-offset-top));
 `),ve("right-start",`
 right: calc(${ae} / -2);
 top: calc(${Pe("right-start")} - var(--v-offset-top));
 `),ve("right",`
 right: calc(${ae} / -2);
 transform: translateY(calc(${ae} / -2)) rotate(45deg);
 top: 50%;
 `),ve("right-end",`
 right: calc(${ae} / -2);
 bottom: calc(${Pe("right-end")} + var(--v-offset-top));
 `),...Wi({top:["right-start","left-start"],right:["top-end","bottom-end"],bottom:["right-end","left-end"],left:["top-start","bottom-start"]},(e,n)=>{const o=["right","left"].includes(n),t=o?"width":"height";return e.map(r=>{const i=r.split("-")[1]==="end",a=`calc((${`var(--v-target-${t}, 0px)`} - ${ae}) / 2)`,s=Pe(r);return ee(`[v-placement="${r}"] >`,[B("popover-shared",[U("center-arrow",[B("popover-arrow",`${n}: calc(max(${a}, ${s}) ${i?"+":"-"} var(--v-offset-${o?"left":"top"}));`)])])])})})]);function Pe(e){return["top","bottom"].includes(e.split("-")[0])?"var(--n-arrow-offset)":"var(--n-arrow-offset-vertical)"}function ve(e,n){const o=e.split("-")[0],t=["top","bottom"].includes(o)?"height: var(--n-space-arrow);":"width: var(--n-space-arrow);";return ee(`[v-placement="${e}"] >`,[B("popover-shared",`
 margin-${An[o]}: var(--n-space);
 `,[U("show-arrow",`
 margin-${An[o]}: var(--n-space-arrow);
 `),U("overlap",`
 margin: 0;
 `),tr("popover-arrow-wrapper",`
 right: 0;
 left: 0;
 top: 0;
 bottom: 0;
 ${o}: 100%;
 ${An[o]}: auto;
 ${t}
 `,[B("popover-arrow",n)])])])}const tt=Object.assign(Object.assign({},se.props),{to:_e.propTo,show:Boolean,trigger:String,showArrow:Boolean,delay:Number,duration:Number,raw:Boolean,arrowPointToCenter:Boolean,arrowStyle:[String,Object],displayDirective:String,x:Number,y:Number,flip:Boolean,overlap:Boolean,placement:String,width:[Number,String],keepAliveOnHover:Boolean,scrollable:Boolean,contentStyle:[Object,String],headerStyle:[Object,String],footerStyle:[Object,String],internalDeactivateImmediately:Boolean,animated:Boolean,onClickoutside:Function,internalTrapFocus:Boolean,internalOnAfterLeave:Function,minWidth:Number,maxWidth:Number}),rt=({arrowStyle:e,clsPrefix:n})=>d("div",{key:"__popover-arrow__",class:`${n}-popover-arrow-wrapper`},d("div",{class:`${n}-popover-arrow`,style:e})),Fl=ie({name:"PopoverBody",inheritAttrs:!1,props:tt,setup(e,{slots:n,attrs:o}){const{namespaceRef:t,mergedClsPrefixRef:r,inlineThemeDisabled:i}=We(e),l=se("Popover","-popover",Ml,pn,e,r),a=E(null),s=be("NPopover"),c=E(null),u=E(e.show),h=E(!1);Yn(()=>{const{show:T}=e;T&&!Sr()&&!e.internalDeactivateImmediately&&(h.value=!0)});const g=F(()=>{const{trigger:T,onClickoutside:N}=e,L=[],{positionManuallyRef:{value:k}}=s;return k||(T==="click"&&!N&&L.push([rn,w,void 0,{capture:!0}]),T==="hover"&&L.push([Rr,D])),N&&L.push([rn,w,void 0,{capture:!0}]),(e.displayDirective==="show"||e.animated&&h.value)&&L.push([Do,e.show]),L}),b=F(()=>{const T=e.width==="trigger"?void 0:tn(e.width),N=[];T&&N.push({width:T});const{maxWidth:L,minWidth:k}=e;return L&&N.push({maxWidth:tn(L)}),k&&N.push({maxWidth:tn(k)}),i||N.push(v.value),N}),v=F(()=>{const{common:{cubicBezierEaseInOut:T,cubicBezierEaseIn:N,cubicBezierEaseOut:L},self:{space:k,spaceArrow:G,padding:z,fontSize:W,textColor:x,dividerColor:I,color:j,boxShadow:A,borderRadius:V,arrowHeight:p,arrowOffset:P,arrowOffsetVertical:ne}}=l.value;return{"--n-box-shadow":A,"--n-bezier":T,"--n-bezier-ease-in":N,"--n-bezier-ease-out":L,"--n-font-size":W,"--n-text-color":x,"--n-color":j,"--n-divider-color":I,"--n-border-radius":V,"--n-arrow-height":p,"--n-arrow-offset":P,"--n-arrow-offset-vertical":ne,"--n-padding":z,"--n-space":k,"--n-space-arrow":G}}),m=i?ze("popover",void 0,v,e):void 0;s.setBodyInstance({syncPosition:R}),Gn(()=>{s.setBodyInstance(null)}),Oe(te(e,"show"),T=>{e.animated||(T?u.value=!0:u.value=!1)});function R(){var T;(T=a.value)===null||T===void 0||T.syncPosition()}function y(T){e.trigger==="hover"&&e.keepAliveOnHover&&e.show&&s.handleMouseEnter(T)}function O(T){e.trigger==="hover"&&e.keepAliveOnHover&&s.handleMouseLeave(T)}function D(T){e.trigger==="hover"&&!S().contains(En(T))&&s.handleMouseMoveOutside(T)}function w(T){(e.trigger==="click"&&!S().contains(En(T))||e.onClickoutside)&&s.handleClickOutside(T)}function S(){return s.getTriggerElement()}fe(Xn,c),fe(Ho,null),fe(Ko,null);function _(){if(m==null||m.onRender(),!(e.displayDirective==="show"||e.show||e.animated&&h.value))return null;let N;const L=s.internalRenderBodyRef.value,{value:k}=r;if(L)N=L([`${k}-popover-shared`,m==null?void 0:m.themeClass.value,e.overlap&&`${k}-popover-shared--overlap`,e.showArrow&&`${k}-popover-shared--show-arrow`,e.arrowPointToCenter&&`${k}-popover-shared--center-arrow`],c,b.value,y,O);else{const{value:G}=s.extraClassRef,{internalTrapFocus:z}=e,W=!po(n.header)||!po(n.footer),x=()=>{var I;const j=W?d(Qn,null,Le(n.header,p=>p?d("div",{class:`${k}-popover__header`,style:e.headerStyle},p):null),Le(n.default,p=>p?d("div",{class:`${k}-popover__content`,style:e.contentStyle},n):null),Le(n.footer,p=>p?d("div",{class:`${k}-popover__footer`,style:e.footerStyle},p):null)):e.scrollable?(I=n.default)===null||I===void 0?void 0:I.call(n):d("div",{class:`${k}-popover__content`,style:e.contentStyle},n),A=e.scrollable?d(Wo,{contentClass:W?void 0:`${k}-popover__content`,contentStyle:W?void 0:e.contentStyle},{default:()=>j}):j,V=e.showArrow?rt({arrowStyle:e.arrowStyle,clsPrefix:k}):null;return[A,V]};N=d("div",hn({class:[`${k}-popover`,`${k}-popover-shared`,m==null?void 0:m.themeClass.value,G.map(I=>`${k}-${I}`),{[`${k}-popover--scrollable`]:e.scrollable,[`${k}-popover--show-header-or-footer`]:W,[`${k}-popover--raw`]:e.raw,[`${k}-popover-shared--overlap`]:e.overlap,[`${k}-popover-shared--show-arrow`]:e.showArrow,[`${k}-popover-shared--center-arrow`]:e.arrowPointToCenter}],ref:c,style:b.value,onKeydown:s.handleKeydown,onMouseenter:y,onMouseleave:O},o),z?d(rr,{active:e.show,autoFocus:!0},{default:x}):x())}return Jn(N,g.value)}return{displayed:h,namespace:t,isMounted:s.isMountedRef,zIndex:s.zIndexRef,followerRef:a,adjustedTo:_e(e),followerEnabled:u,renderContentNode:_}},render(){return d(no,{ref:"followerRef",zIndex:this.zIndex,show:this.show,enabled:this.followerEnabled,to:this.adjustedTo,x:this.x,y:this.y,flip:this.flip,placement:this.placement,containerClass:this.namespace,overlap:this.overlap,width:this.width==="trigger"?"target":void 0,teleportDisabled:this.adjustedTo===_e.tdkey},{default:()=>this.animated?d(fn,{name:"popover-transition",appear:this.isMounted,onEnter:()=>{this.followerEnabled=!0},onAfterLeave:()=>{var e;(e=this.internalOnAfterLeave)===null||e===void 0||e.call(this),this.followerEnabled=!1,this.displayed=!1}},{default:this.renderContentNode}):this.renderContentNode()})}}),Al=Object.keys(tt),_l={focus:["onFocus","onBlur"],click:["onClick"],hover:["onMouseenter","onMouseleave"],manual:[],nested:["onFocus","onBlur","onMouseenter","onMouseleave","onClick"]};function Bl(e,n,o){_l[n].forEach(t=>{e.props?e.props=Object.assign({},e.props):e.props={};const r=e.props[t],i=o[t];r?e.props[t]=(...l)=>{r(...l),i(...l)}:e.props[t]=i})}const Nl=ar("").type,gn={show:{type:Boolean,default:void 0},defaultShow:Boolean,showArrow:{type:Boolean,default:!0},trigger:{type:String,default:"hover"},delay:{type:Number,default:100},duration:{type:Number,default:100},raw:Boolean,placement:{type:String,default:"top"},x:Number,y:Number,arrowPointToCenter:Boolean,disabled:Boolean,getDisabled:Function,displayDirective:{type:String,default:"if"},arrowStyle:[String,Object],flip:{type:Boolean,default:!0},animated:{type:Boolean,default:!0},width:{type:[Number,String],default:void 0},overlap:Boolean,keepAliveOnHover:{type:Boolean,default:!0},zIndex:Number,to:_e.propTo,scrollable:Boolean,contentStyle:[Object,String],headerStyle:[Object,String],footerStyle:[Object,String],onClickoutside:Function,"onUpdate:show":[Function,Array],onUpdateShow:[Function,Array],internalDeactivateImmediately:Boolean,internalSyncTargetWithParent:Boolean,internalInheritedEventHandlers:{type:Array,default:()=>[]},internalTrapFocus:Boolean,internalExtraClass:{type:Array,default:()=>[]},onShow:[Function,Array],onHide:[Function,Array],arrow:{type:Boolean,default:void 0},minWidth:Number,maxWidth:Number},El=Object.assign(Object.assign(Object.assign({},se.props),gn),{internalOnAfterLeave:Function,internalRenderBody:Function}),ao=ie({name:"Popover",inheritAttrs:!1,props:El,__popover__:!0,setup(e){const n=jo(),o=E(null),t=F(()=>e.show),r=E(e.defaultShow),i=ln(t,r),l=Re(()=>e.disabled?!1:i.value),a=()=>{if(e.disabled)return!0;const{getDisabled:x}=e;return!!(x!=null&&x())},s=()=>a()?!1:i.value,c=Uo(e,["arrow","showArrow"]),u=F(()=>e.overlap?!1:c.value);let h=null;const g=E(null),b=E(null),v=Re(()=>e.x!==void 0&&e.y!==void 0);function m(x){const{"onUpdate:show":I,onUpdateShow:j,onShow:A,onHide:V}=e;r.value=x,I&&ce(I,x),j&&ce(j,x),x&&A&&ce(A,!0),x&&V&&ce(V,!1)}function R(){h&&h.syncPosition()}function y(){const{value:x}=g;x&&(window.clearTimeout(x),g.value=null)}function O(){const{value:x}=b;x&&(window.clearTimeout(x),b.value=null)}function D(){const x=a();if(e.trigger==="focus"&&!x){if(s())return;m(!0)}}function w(){const x=a();if(e.trigger==="focus"&&!x){if(!s())return;m(!1)}}function S(){const x=a();if(e.trigger==="hover"&&!x){if(O(),g.value!==null||s())return;const I=()=>{m(!0),g.value=null},{delay:j}=e;j===0?I():g.value=window.setTimeout(I,j)}}function _(){const x=a();if(e.trigger==="hover"&&!x){if(y(),b.value!==null||!s())return;const I=()=>{m(!1),b.value=null},{duration:j}=e;j===0?I():b.value=window.setTimeout(I,j)}}function T(){_()}function N(x){var I;s()&&(e.trigger==="click"&&(y(),O(),m(!1)),(I=e.onClickoutside)===null||I===void 0||I.call(e,x))}function L(){if(e.trigger==="click"&&!a()){y(),O();const x=!s();m(x)}}function k(x){e.internalTrapFocus&&x.key==="Escape"&&(y(),O(),m(!1))}function G(x){r.value=x}function z(){var x;return(x=o.value)===null||x===void 0?void 0:x.targetRef}function W(x){h=x}return fe("NPopover",{getTriggerElement:z,handleKeydown:k,handleMouseEnter:S,handleMouseLeave:_,handleClickOutside:N,handleMouseMoveOutside:T,setBodyInstance:W,positionManuallyRef:v,isMountedRef:n,zIndexRef:te(e,"zIndex"),extraClassRef:te(e,"internalExtraClass"),internalRenderBodyRef:te(e,"internalRenderBody")}),Yn(()=>{i.value&&a()&&m(!1)}),{binderInstRef:o,positionManually:v,mergedShowConsideringDisabledProp:l,uncontrolledShow:r,mergedShowArrow:u,getMergedShow:s,setShow:G,handleClick:L,handleMouseEnter:S,handleMouseLeave:_,handleFocus:D,handleBlur:w,syncPosition:R}},render(){var e;const{positionManually:n,$slots:o}=this;let t,r=!1;if(!n&&(o.activator?t=go(o,"activator"):t=go(o,"trigger"),t)){t=ir(t),t=t.type===Nl?d("span",[t]):t;const i={onClick:this.handleClick,onMouseenter:this.handleMouseEnter,onMouseleave:this.handleMouseLeave,onFocus:this.handleFocus,onBlur:this.handleBlur};if(!((e=t.type)===null||e===void 0)&&e.__popover__)r=!0,t.props||(t.props={internalSyncTargetWithParent:!0,internalInheritedEventHandlers:[]}),t.props.internalSyncTargetWithParent=!0,t.props.internalInheritedEventHandlers?t.props.internalInheritedEventHandlers=[i,...t.props.internalInheritedEventHandlers]:t.props.internalInheritedEventHandlers=[i];else{const{internalInheritedEventHandlers:l}=this,a=[i,...l],s={onBlur:c=>{a.forEach(u=>{u.onBlur(c)})},onFocus:c=>{a.forEach(u=>{u.onFocus(c)})},onClick:c=>{a.forEach(u=>{u.onClick(c)})},onMouseenter:c=>{a.forEach(u=>{u.onMouseenter(c)})},onMouseleave:c=>{a.forEach(u=>{u.onMouseleave(c)})}};Bl(t,l?"nested":n?"manual":this.trigger,s)}}return d(to,{ref:"binderInstRef",syncTarget:!r,syncTargetWithParent:this.internalSyncTargetWithParent},{default:()=>{this.mergedShowConsideringDisabledProp;const i=this.getMergedShow();return[this.internalTrapFocus&&i?Jn(d("div",{style:{position:"fixed",inset:0}}),[[lr,{enabled:i,zIndex:this.zIndex}]]):null,n?null:d(oo,null,{default:()=>t}),d(Fl,Vo(this.$props,Al,Object.assign(Object.assign({},this.$attrs),{showArrow:this.mergedShowArrow,show:i})),{default:()=>{var l,a;return(a=(l=this.$slots).default)===null||a===void 0?void 0:a.call(l)},header:()=>{var l,a;return(a=(l=this.$slots).header)===null||a===void 0?void 0:a.call(l)},footer:()=>{var l,a;return(a=(l=this.$slots).footer)===null||a===void 0?void 0:a.call(l)}})]}})}}),Ll={closeIconSizeTiny:"12px",closeIconSizeSmall:"12px",closeIconSizeMedium:"14px",closeIconSizeLarge:"14px",closeSizeTiny:"16px",closeSizeSmall:"16px",closeSizeMedium:"18px",closeSizeLarge:"18px",padding:"0 7px",closeMargin:"0 0 0 4px",closeMarginRtl:"0 4px 0 0"},Dl=e=>{const{textColor2:n,primaryColorHover:o,primaryColorPressed:t,primaryColor:r,infoColor:i,successColor:l,warningColor:a,errorColor:s,baseColor:c,borderColor:u,opacityDisabled:h,tagColor:g,closeIconColor:b,closeIconColorHover:v,closeIconColorPressed:m,borderRadiusSmall:R,fontSizeMini:y,fontSizeTiny:O,fontSizeSmall:D,fontSizeMedium:w,heightMini:S,heightTiny:_,heightSmall:T,heightMedium:N,closeColorHover:L,closeColorPressed:k,buttonColor2Hover:G,buttonColor2Pressed:z,fontWeightStrong:W}=e;return Object.assign(Object.assign({},Ll),{closeBorderRadius:R,heightTiny:S,heightSmall:_,heightMedium:T,heightLarge:N,borderRadius:R,opacityDisabled:h,fontSizeTiny:y,fontSizeSmall:O,fontSizeMedium:D,fontSizeLarge:w,fontWeightStrong:W,textColorCheckable:n,textColorHoverCheckable:n,textColorPressedCheckable:n,textColorChecked:c,colorCheckable:"#0000",colorHoverCheckable:G,colorPressedCheckable:z,colorChecked:r,colorCheckedHover:o,colorCheckedPressed:t,border:`1px solid ${u}`,textColor:n,color:g,colorBordered:"rgb(250, 250, 252)",closeIconColor:b,closeIconColorHover:v,closeIconColorPressed:m,closeColorHover:L,closeColorPressed:k,borderPrimary:`1px solid ${X(r,{alpha:.3})}`,textColorPrimary:r,colorPrimary:X(r,{alpha:.12}),colorBorderedPrimary:X(r,{alpha:.1}),closeIconColorPrimary:r,closeIconColorHoverPrimary:r,closeIconColorPressedPrimary:r,closeColorHoverPrimary:X(r,{alpha:.12}),closeColorPressedPrimary:X(r,{alpha:.18}),borderInfo:`1px solid ${X(i,{alpha:.3})}`,textColorInfo:i,colorInfo:X(i,{alpha:.12}),colorBorderedInfo:X(i,{alpha:.1}),closeIconColorInfo:i,closeIconColorHoverInfo:i,closeIconColorPressedInfo:i,closeColorHoverInfo:X(i,{alpha:.12}),closeColorPressedInfo:X(i,{alpha:.18}),borderSuccess:`1px solid ${X(l,{alpha:.3})}`,textColorSuccess:l,colorSuccess:X(l,{alpha:.12}),colorBorderedSuccess:X(l,{alpha:.1}),closeIconColorSuccess:l,closeIconColorHoverSuccess:l,closeIconColorPressedSuccess:l,closeColorHoverSuccess:X(l,{alpha:.12}),closeColorPressedSuccess:X(l,{alpha:.18}),borderWarning:`1px solid ${X(a,{alpha:.35})}`,textColorWarning:a,colorWarning:X(a,{alpha:.15}),colorBorderedWarning:X(a,{alpha:.12}),closeIconColorWarning:a,closeIconColorHoverWarning:a,closeIconColorPressedWarning:a,closeColorHoverWarning:X(a,{alpha:.12}),closeColorPressedWarning:X(a,{alpha:.18}),borderError:`1px solid ${X(s,{alpha:.23})}`,textColorError:s,colorError:X(s,{alpha:.1}),colorBorderedError:X(s,{alpha:.08}),closeIconColorError:s,closeIconColorHoverError:s,closeIconColorPressedError:s,closeColorHoverError:X(s,{alpha:.12}),closeColorPressedError:X(s,{alpha:.18})})},Hl={name:"Tag",common:ke,self:Dl},Kl=Hl,Wl={color:Object,type:{type:String,default:"default"},round:Boolean,size:{type:String,default:"medium"},closable:Boolean,disabled:{type:Boolean,default:void 0}},jl=B("tag",`
 white-space: nowrap;
 position: relative;
 box-sizing: border-box;
 cursor: default;
 display: inline-flex;
 align-items: center;
 flex-wrap: nowrap;
 padding: var(--n-padding);
 border-radius: var(--n-border-radius);
 color: var(--n-text-color);
 background-color: var(--n-color);
 transition: 
 border-color .3s var(--n-bezier),
 background-color .3s var(--n-bezier),
 color .3s var(--n-bezier),
 box-shadow .3s var(--n-bezier),
 opacity .3s var(--n-bezier);
 line-height: 1;
 height: var(--n-height);
 font-size: var(--n-font-size);
`,[U("strong",`
 font-weight: var(--n-font-weight-strong);
 `),H("border",`
 pointer-events: none;
 position: absolute;
 left: 0;
 right: 0;
 top: 0;
 bottom: 0;
 border-radius: inherit;
 border: var(--n-border);
 transition: border-color .3s var(--n-bezier);
 `),H("icon",`
 display: flex;
 margin: 0 4px 0 0;
 color: var(--n-text-color);
 transition: color .3s var(--n-bezier);
 font-size: var(--n-avatar-size-override);
 `),H("avatar",`
 display: flex;
 margin: 0 6px 0 0;
 `),H("close",`
 margin: var(--n-close-margin);
 transition:
 background-color .3s var(--n-bezier),
 color .3s var(--n-bezier);
 `),U("round",`
 padding: 0 calc(var(--n-height) / 3);
 border-radius: calc(var(--n-height) / 2);
 `,[H("icon",`
 margin: 0 4px 0 calc((var(--n-height) - 8px) / -2);
 `),H("avatar",`
 margin: 0 6px 0 calc((var(--n-height) - 8px) / -2);
 `),U("closable",`
 padding: 0 calc(var(--n-height) / 4) 0 calc(var(--n-height) / 3);
 `)]),U("icon, avatar",[U("round",`
 padding: 0 calc(var(--n-height) / 3) 0 calc(var(--n-height) / 2);
 `)]),U("disabled",`
 cursor: not-allowed !important;
 opacity: var(--n-opacity-disabled);
 `),U("checkable",`
 cursor: pointer;
 box-shadow: none;
 color: var(--n-text-color-checkable);
 background-color: var(--n-color-checkable);
 `,[pe("disabled",[ee("&:hover","background-color: var(--n-color-hover-checkable);",[pe("checked","color: var(--n-text-color-hover-checkable);")]),ee("&:active","background-color: var(--n-color-pressed-checkable);",[pe("checked","color: var(--n-text-color-pressed-checkable);")])]),U("checked",`
 color: var(--n-text-color-checked);
 background-color: var(--n-color-checked);
 `,[pe("disabled",[ee("&:hover","background-color: var(--n-color-checked-hover);"),ee("&:active","background-color: var(--n-color-checked-pressed);")])])])]),Vl=Object.assign(Object.assign(Object.assign({},se.props),Wl),{bordered:{type:Boolean,default:void 0},checked:Boolean,checkable:Boolean,strong:Boolean,triggerClickOnClose:Boolean,onClose:[Array,Function],onMouseenter:Function,onMouseleave:Function,"onUpdate:checked":Function,onUpdateChecked:Function,internalCloseFocusable:{type:Boolean,default:!0},internalCloseIsButtonTag:{type:Boolean,default:!0},onCheckedChange:Function}),Gl=vn("n-tag"),_n=ie({name:"Tag",props:Vl,setup(e){const n=E(null),{mergedBorderedRef:o,mergedClsPrefixRef:t,inlineThemeDisabled:r,mergedRtlRef:i}=We(e),l=se("Tag","-tag",jl,Kl,e,t);fe(Gl,{roundRef:te(e,"round")});function a(b){if(!e.disabled&&e.checkable){const{checked:v,onCheckedChange:m,onUpdateChecked:R,"onUpdate:checked":y}=e;R&&R(!v),y&&y(!v),m&&m(!v)}}function s(b){if(e.triggerClickOnClose||b.stopPropagation(),!e.disabled){const{onClose:v}=e;v&&ce(v,b)}}const c={setTextContent(b){const{value:v}=n;v&&(v.textContent=b)}},u=sr("Tag",i,t),h=F(()=>{const{type:b,size:v,color:{color:m,textColor:R}={}}=e,{common:{cubicBezierEaseInOut:y},self:{padding:O,closeMargin:D,closeMarginRtl:w,borderRadius:S,opacityDisabled:_,textColorCheckable:T,textColorHoverCheckable:N,textColorPressedCheckable:L,textColorChecked:k,colorCheckable:G,colorHoverCheckable:z,colorPressedCheckable:W,colorChecked:x,colorCheckedHover:I,colorCheckedPressed:j,closeBorderRadius:A,fontWeightStrong:V,[oe("colorBordered",b)]:p,[oe("closeSize",v)]:P,[oe("closeIconSize",v)]:ne,[oe("fontSize",v)]:le,[oe("height",v)]:me,[oe("color",b)]:Ce,[oe("textColor",b)]:Se,[oe("border",b)]:ye,[oe("closeIconColor",b)]:Q,[oe("closeIconColorHover",b)]:ge,[oe("closeIconColorPressed",b)]:ue,[oe("closeColorHover",b)]:$e,[oe("closeColorPressed",b)]:we}}=l.value;return{"--n-font-weight-strong":V,"--n-avatar-size-override":`calc(${me} - 8px)`,"--n-bezier":y,"--n-border-radius":S,"--n-border":ye,"--n-close-icon-size":ne,"--n-close-color-pressed":we,"--n-close-color-hover":$e,"--n-close-border-radius":A,"--n-close-icon-color":Q,"--n-close-icon-color-hover":ge,"--n-close-icon-color-pressed":ue,"--n-close-icon-color-disabled":Q,"--n-close-margin":D,"--n-close-margin-rtl":w,"--n-close-size":P,"--n-color":m||(o.value?p:Ce),"--n-color-checkable":G,"--n-color-checked":x,"--n-color-checked-hover":I,"--n-color-checked-pressed":j,"--n-color-hover-checkable":z,"--n-color-pressed-checkable":W,"--n-font-size":le,"--n-height":me,"--n-opacity-disabled":_,"--n-padding":O,"--n-text-color":R||Se,"--n-text-color-checkable":T,"--n-text-color-checked":k,"--n-text-color-hover-checkable":N,"--n-text-color-pressed-checkable":L}}),g=r?ze("tag",F(()=>{let b="";const{type:v,size:m,color:{color:R,textColor:y}={}}=e;return b+=v[0],b+=m[0],R&&(b+=`a${bo(R)}`),y&&(b+=`b${bo(y)}`),o.value&&(b+="c"),b}),h,e):void 0;return Object.assign(Object.assign({},c),{rtlEnabled:u,mergedClsPrefix:t,contentRef:n,mergedBordered:o,handleClick:a,handleCloseClick:s,cssVars:r?void 0:h,themeClass:g==null?void 0:g.themeClass,onRender:g==null?void 0:g.onRender})},render(){var e,n;const{mergedClsPrefix:o,rtlEnabled:t,closable:r,color:{borderColor:i}={},round:l,onRender:a,$slots:s}=this;a==null||a();const c=Le(s.avatar,h=>h&&d("div",{class:`${o}-tag__avatar`},h)),u=Le(s.icon,h=>h&&d("div",{class:`${o}-tag__icon`},h));return d("div",{class:[`${o}-tag`,this.themeClass,{[`${o}-tag--rtl`]:t,[`${o}-tag--strong`]:this.strong,[`${o}-tag--disabled`]:this.disabled,[`${o}-tag--checkable`]:this.checkable,[`${o}-tag--checked`]:this.checkable&&this.checked,[`${o}-tag--round`]:l,[`${o}-tag--avatar`]:c,[`${o}-tag--icon`]:u,[`${o}-tag--closable`]:r}],style:this.cssVars,onClick:this.handleClick,onMouseenter:this.onMouseenter,onMouseleave:this.onMouseleave},u||c,d("span",{class:`${o}-tag__content`,ref:"contentRef"},(n=(e=this.$slots).default)===null||n===void 0?void 0:n.call(e)),!this.checkable&&r?d(dr,{clsPrefix:o,class:`${o}-tag__close`,disabled:this.disabled,onClick:this.handleCloseClick,focusable:this.internalCloseFocusable,round:l,isButtonTag:this.internalCloseIsButtonTag,absolute:!0}):null,!this.checkable&&this.mergedBordered?d("div",{class:`${o}-tag__border`,style:{borderColor:i}}):null)}}),Ul={paddingSingle:"0 26px 0 12px",paddingMultiple:"3px 26px 0 12px",clearSize:"16px",arrowSize:"16px"},ql=e=>{const{borderRadius:n,textColor2:o,textColorDisabled:t,inputColor:r,inputColorDisabled:i,primaryColor:l,primaryColorHover:a,warningColor:s,warningColorHover:c,errorColor:u,errorColorHover:h,borderColor:g,iconColor:b,iconColorDisabled:v,clearColor:m,clearColorHover:R,clearColorPressed:y,placeholderColor:O,placeholderColorDisabled:D,fontSizeTiny:w,fontSizeSmall:S,fontSizeMedium:_,fontSizeLarge:T,heightTiny:N,heightSmall:L,heightMedium:k,heightLarge:G}=e;return Object.assign(Object.assign({},Ul),{fontSizeTiny:w,fontSizeSmall:S,fontSizeMedium:_,fontSizeLarge:T,heightTiny:N,heightSmall:L,heightMedium:k,heightLarge:G,borderRadius:n,textColor:o,textColorDisabled:t,placeholderColor:O,placeholderColorDisabled:D,color:r,colorDisabled:i,colorActive:r,border:`1px solid ${g}`,borderHover:`1px solid ${a}`,borderActive:`1px solid ${l}`,borderFocus:`1px solid ${a}`,boxShadowHover:"none",boxShadowActive:`0 0 0 2px ${X(l,{alpha:.2})}`,boxShadowFocus:`0 0 0 2px ${X(l,{alpha:.2})}`,caretColor:l,arrowColor:b,arrowColorDisabled:v,loadingColor:l,borderWarning:`1px solid ${s}`,borderHoverWarning:`1px solid ${c}`,borderActiveWarning:`1px solid ${s}`,borderFocusWarning:`1px solid ${c}`,boxShadowHoverWarning:"none",boxShadowActiveWarning:`0 0 0 2px ${X(s,{alpha:.2})}`,boxShadowFocusWarning:`0 0 0 2px ${X(s,{alpha:.2})}`,colorActiveWarning:r,caretColorWarning:s,borderError:`1px solid ${u}`,borderHoverError:`1px solid ${h}`,borderActiveError:`1px solid ${u}`,borderFocusError:`1px solid ${h}`,boxShadowHoverError:"none",boxShadowActiveError:`0 0 0 2px ${X(u,{alpha:.2})}`,boxShadowFocusError:`0 0 0 2px ${X(u,{alpha:.2})}`,colorActiveError:r,caretColorError:u,clearColor:m,clearColorHover:R,clearColorPressed:y})},Zl=qe({name:"InternalSelection",common:ke,peers:{Popover:pn},self:ql}),it=Zl,Yl=ee([B("base-selection",`
 position: relative;
 z-index: auto;
 box-shadow: none;
 width: 100%;
 max-width: 100%;
 display: inline-block;
 vertical-align: bottom;
 border-radius: var(--n-border-radius);
 min-height: var(--n-height);
 line-height: 1.5;
 font-size: var(--n-font-size);
 `,[B("base-loading",`
 color: var(--n-loading-color);
 `),B("base-selection-tags","min-height: var(--n-height);"),H("border, state-border",`
 position: absolute;
 left: 0;
 right: 0;
 top: 0;
 bottom: 0;
 pointer-events: none;
 border: var(--n-border);
 border-radius: inherit;
 transition:
 box-shadow .3s var(--n-bezier),
 border-color .3s var(--n-bezier);
 `),H("state-border",`
 z-index: 1;
 border-color: #0000;
 `),B("base-suffix",`
 cursor: pointer;
 position: absolute;
 top: 50%;
 transform: translateY(-50%);
 right: 10px;
 `,[H("arrow",`
 font-size: var(--n-arrow-size);
 color: var(--n-arrow-color);
 transition: color .3s var(--n-bezier);
 `)]),B("base-selection-overlay",`
 display: flex;
 align-items: center;
 white-space: nowrap;
 pointer-events: none;
 position: absolute;
 top: 0;
 right: 0;
 bottom: 0;
 left: 0;
 padding: var(--n-padding-single);
 transition: color .3s var(--n-bezier);
 `,[H("wrapper",`
 flex-basis: 0;
 flex-grow: 1;
 overflow: hidden;
 text-overflow: ellipsis;
 `)]),B("base-selection-placeholder",`
 color: var(--n-placeholder-color);
 `,[H("inner",`
 max-width: 100%;
 overflow: hidden;
 `)]),B("base-selection-tags",`
 cursor: pointer;
 outline: none;
 box-sizing: border-box;
 position: relative;
 z-index: auto;
 display: flex;
 padding: var(--n-padding-multiple);
 flex-wrap: wrap;
 align-items: center;
 width: 100%;
 vertical-align: bottom;
 background-color: var(--n-color);
 border-radius: inherit;
 transition:
 color .3s var(--n-bezier),
 box-shadow .3s var(--n-bezier),
 background-color .3s var(--n-bezier);
 `),B("base-selection-label",`
 height: var(--n-height);
 display: inline-flex;
 width: 100%;
 vertical-align: bottom;
 cursor: pointer;
 outline: none;
 z-index: auto;
 box-sizing: border-box;
 position: relative;
 transition:
 color .3s var(--n-bezier),
 box-shadow .3s var(--n-bezier),
 background-color .3s var(--n-bezier);
 border-radius: inherit;
 background-color: var(--n-color);
 align-items: center;
 `,[B("base-selection-input",`
 font-size: inherit;
 line-height: inherit;
 outline: none;
 cursor: pointer;
 box-sizing: border-box;
 border:none;
 width: 100%;
 padding: var(--n-padding-single);
 background-color: #0000;
 color: var(--n-text-color);
 transition: color .3s var(--n-bezier);
 caret-color: var(--n-caret-color);
 `,[H("content",`
 text-overflow: ellipsis;
 overflow: hidden;
 white-space: nowrap; 
 `)]),H("render-label",`
 color: var(--n-text-color);
 `)]),pe("disabled",[ee("&:hover",[H("state-border",`
 box-shadow: var(--n-box-shadow-hover);
 border: var(--n-border-hover);
 `)]),U("focus",[H("state-border",`
 box-shadow: var(--n-box-shadow-focus);
 border: var(--n-border-focus);
 `)]),U("active",[H("state-border",`
 box-shadow: var(--n-box-shadow-active);
 border: var(--n-border-active);
 `),B("base-selection-label","background-color: var(--n-color-active);"),B("base-selection-tags","background-color: var(--n-color-active);")])]),U("disabled","cursor: not-allowed;",[H("arrow",`
 color: var(--n-arrow-color-disabled);
 `),B("base-selection-label",`
 cursor: not-allowed;
 background-color: var(--n-color-disabled);
 `,[B("base-selection-input",`
 cursor: not-allowed;
 color: var(--n-text-color-disabled);
 `),H("render-label",`
 color: var(--n-text-color-disabled);
 `)]),B("base-selection-tags",`
 cursor: not-allowed;
 background-color: var(--n-color-disabled);
 `),B("base-selection-placeholder",`
 cursor: not-allowed;
 color: var(--n-placeholder-color-disabled);
 `)]),B("base-selection-input-tag",`
 height: calc(var(--n-height) - 6px);
 line-height: calc(var(--n-height) - 6px);
 outline: none;
 display: none;
 position: relative;
 margin-bottom: 3px;
 max-width: 100%;
 vertical-align: bottom;
 `,[H("input",`
 font-size: inherit;
 font-family: inherit;
 min-width: 1px;
 padding: 0;
 background-color: #0000;
 outline: none;
 border: none;
 max-width: 100%;
 overflow: hidden;
 width: 1em;
 line-height: inherit;
 cursor: pointer;
 color: var(--n-text-color);
 caret-color: var(--n-caret-color);
 `),H("mirror",`
 position: absolute;
 left: 0;
 top: 0;
 white-space: pre;
 visibility: hidden;
 user-select: none;
 -webkit-user-select: none;
 opacity: 0;
 `)]),["warning","error"].map(e=>U(`${e}-status`,[H("state-border",`border: var(--n-border-${e});`),pe("disabled",[ee("&:hover",[H("state-border",`
 box-shadow: var(--n-box-shadow-hover-${e});
 border: var(--n-border-hover-${e});
 `)]),U("active",[H("state-border",`
 box-shadow: var(--n-box-shadow-active-${e});
 border: var(--n-border-active-${e});
 `),B("base-selection-label",`background-color: var(--n-color-active-${e});`),B("base-selection-tags",`background-color: var(--n-color-active-${e});`)]),U("focus",[H("state-border",`
 box-shadow: var(--n-box-shadow-focus-${e});
 border: var(--n-border-focus-${e});
 `)])])]))]),B("base-selection-popover",`
 margin-bottom: -3px;
 display: flex;
 flex-wrap: wrap;
 margin-right: -8px;
 `),B("base-selection-tag-wrapper",`
 max-width: 100%;
 display: inline-flex;
 padding: 0 7px 3px 0;
 `,[ee("&:last-child","padding-right: 0;"),B("tag",`
 font-size: 14px;
 max-width: 100%;
 `,[H("content",`
 line-height: 1.25;
 text-overflow: ellipsis;
 overflow: hidden;
 `)])])]),Jl=ie({name:"InternalSelection",props:Object.assign(Object.assign({},se.props),{clsPrefix:{type:String,required:!0},bordered:{type:Boolean,default:void 0},active:Boolean,pattern:{type:String,default:""},placeholder:String,selectedOption:{type:Object,default:null},selectedOptions:{type:Array,default:null},labelField:{type:String,default:"label"},valueField:{type:String,default:"value"},multiple:Boolean,filterable:Boolean,clearable:Boolean,disabled:Boolean,size:{type:String,default:"medium"},loading:Boolean,autofocus:Boolean,showArrow:{type:Boolean,default:!0},inputProps:Object,focused:Boolean,renderTag:Function,onKeydown:Function,onClick:Function,onBlur:Function,onFocus:Function,onDeleteOption:Function,maxTagCount:[String,Number],onClear:Function,onPatternInput:Function,onPatternFocus:Function,onPatternBlur:Function,renderLabel:Function,status:String,inlineThemeDisabled:Boolean,ignoreComposition:{type:Boolean,default:!0},onResize:Function}),setup(e){const n=E(null),o=E(null),t=E(null),r=E(null),i=E(null),l=E(null),a=E(null),s=E(null),c=E(null),u=E(null),h=E(!1),g=E(!1),b=E(!1),v=se("InternalSelection","-internal-selection",Yl,it,e,te(e,"clsPrefix")),m=F(()=>e.clearable&&!e.disabled&&(b.value||e.active)),R=F(()=>e.selectedOption?e.renderTag?e.renderTag({option:e.selectedOption,handleClose:()=>{}}):e.renderLabel?e.renderLabel(e.selectedOption,!0):xe(e.selectedOption[e.labelField],e.selectedOption,!0):e.placeholder),y=F(()=>{const C=e.selectedOption;if(C)return C[e.labelField]}),O=F(()=>e.multiple?!!(Array.isArray(e.selectedOptions)&&e.selectedOptions.length):e.selectedOption!==null);function D(){var C;const{value:M}=n;if(M){const{value:re}=o;re&&(re.style.width=`${M.offsetWidth}px`,e.maxTagCount!=="responsive"&&((C=c.value)===null||C===void 0||C.sync()))}}function w(){const{value:C}=u;C&&(C.style.display="none")}function S(){const{value:C}=u;C&&(C.style.display="inline-block")}Oe(te(e,"active"),C=>{C||w()}),Oe(te(e,"pattern"),()=>{e.multiple&&Vn(D)});function _(C){const{onFocus:M}=e;M&&M(C)}function T(C){const{onBlur:M}=e;M&&M(C)}function N(C){const{onDeleteOption:M}=e;M&&M(C)}function L(C){const{onClear:M}=e;M&&M(C)}function k(C){const{onPatternInput:M}=e;M&&M(C)}function G(C){var M;(!C.relatedTarget||!(!((M=t.value)===null||M===void 0)&&M.contains(C.relatedTarget)))&&_(C)}function z(C){var M;!((M=t.value)===null||M===void 0)&&M.contains(C.relatedTarget)||T(C)}function W(C){L(C)}function x(){b.value=!0}function I(){b.value=!1}function j(C){!e.active||!e.filterable||C.target!==o.value&&C.preventDefault()}function A(C){N(C)}function V(C){if(C.key==="Backspace"&&!p.value&&!e.pattern.length){const{selectedOptions:M}=e;M!=null&&M.length&&A(M[M.length-1])}}const p=E(!1);let P=null;function ne(C){const{value:M}=n;if(M){const re=C.target.value;M.textContent=re,D()}e.ignoreComposition&&p.value?P=C:k(C)}function le(){p.value=!0}function me(){p.value=!1,e.ignoreComposition&&k(P),P=null}function Ce(C){var M;g.value=!0,(M=e.onPatternFocus)===null||M===void 0||M.call(e,C)}function Se(C){var M;g.value=!1,(M=e.onPatternBlur)===null||M===void 0||M.call(e,C)}function ye(){var C,M;if(e.filterable)g.value=!1,(C=l.value)===null||C===void 0||C.blur(),(M=o.value)===null||M===void 0||M.blur();else if(e.multiple){const{value:re}=r;re==null||re.blur()}else{const{value:re}=i;re==null||re.blur()}}function Q(){var C,M,re;e.filterable?(g.value=!1,(C=l.value)===null||C===void 0||C.focus()):e.multiple?(M=r.value)===null||M===void 0||M.focus():(re=i.value)===null||re===void 0||re.focus()}function ge(){const{value:C}=o;C&&(S(),C.focus())}function ue(){const{value:C}=o;C&&C.blur()}function $e(C){const{value:M}=a;M&&M.setTextContent(`+${C}`)}function we(){const{value:C}=s;return C}function je(){return o.value}let Ie=null;function Me(){Ie!==null&&window.clearTimeout(Ie)}function Ve(){e.disabled||e.active||(Me(),Ie=window.setTimeout(()=>{O.value&&(h.value=!0)},100))}function Ge(){Me()}function Ue(C){C||(Me(),h.value=!1)}Oe(O,C=>{C||(h.value=!1)}),dn(()=>{Yn(()=>{const C=l.value;C&&(C.tabIndex=e.disabled||g.value?-1:0)})}),qo(t,e.onResize);const{inlineThemeDisabled:Be}=e,Ne=F(()=>{const{size:C}=e,{common:{cubicBezierEaseInOut:M},self:{borderRadius:re,color:Ze,placeholderColor:mn,textColor:yn,paddingSingle:wn,paddingMultiple:xn,caretColor:Ye,colorDisabled:Je,textColorDisabled:Xe,placeholderColorDisabled:Cn,colorActive:Sn,boxShadowFocus:Qe,boxShadowActive:Te,boxShadowHover:f,border:$,borderFocus:K,borderHover:J,borderActive:q,arrowColor:Y,arrowColorDisabled:Z,loadingColor:de,colorActiveWarning:en,boxShadowFocusWarning:Pn,boxShadowActiveWarning:ut,boxShadowHoverWarning:ft,borderWarning:ht,borderFocusWarning:vt,borderHoverWarning:pt,borderActiveWarning:gt,colorActiveError:bt,boxShadowFocusError:mt,boxShadowActiveError:yt,boxShadowHoverError:wt,borderError:xt,borderFocusError:Ct,borderHoverError:St,borderActiveError:Pt,clearColor:Ot,clearColorHover:Rt,clearColorPressed:kt,clearSize:$t,arrowSize:Tt,[oe("height",C)]:zt,[oe("fontSize",C)]:It}}=v.value;return{"--n-bezier":M,"--n-border":$,"--n-border-active":q,"--n-border-focus":K,"--n-border-hover":J,"--n-border-radius":re,"--n-box-shadow-active":Te,"--n-box-shadow-focus":Qe,"--n-box-shadow-hover":f,"--n-caret-color":Ye,"--n-color":Ze,"--n-color-active":Sn,"--n-color-disabled":Je,"--n-font-size":It,"--n-height":zt,"--n-padding-single":wn,"--n-padding-multiple":xn,"--n-placeholder-color":mn,"--n-placeholder-color-disabled":Cn,"--n-text-color":yn,"--n-text-color-disabled":Xe,"--n-arrow-color":Y,"--n-arrow-color-disabled":Z,"--n-loading-color":de,"--n-color-active-warning":en,"--n-box-shadow-focus-warning":Pn,"--n-box-shadow-active-warning":ut,"--n-box-shadow-hover-warning":ft,"--n-border-warning":ht,"--n-border-focus-warning":vt,"--n-border-hover-warning":pt,"--n-border-active-warning":gt,"--n-color-active-error":bt,"--n-box-shadow-focus-error":mt,"--n-box-shadow-active-error":yt,"--n-box-shadow-hover-error":wt,"--n-border-error":xt,"--n-border-focus-error":Ct,"--n-border-hover-error":St,"--n-border-active-error":Pt,"--n-clear-size":$t,"--n-clear-color":Ot,"--n-clear-color-hover":Rt,"--n-clear-color-pressed":kt,"--n-arrow-size":Tt}}),he=Be?ze("internal-selection",F(()=>e.size[0]),Ne,e):void 0;return{mergedTheme:v,mergedClearable:m,patternInputFocused:g,filterablePlaceholder:R,label:y,selected:O,showTagsPanel:h,isComposing:p,counterRef:a,counterWrapperRef:s,patternInputMirrorRef:n,patternInputRef:o,selfRef:t,multipleElRef:r,singleElRef:i,patternInputWrapperRef:l,overflowRef:c,inputTagElRef:u,handleMouseDown:j,handleFocusin:G,handleClear:W,handleMouseEnter:x,handleMouseLeave:I,handleDeleteOption:A,handlePatternKeyDown:V,handlePatternInputInput:ne,handlePatternInputBlur:Se,handlePatternInputFocus:Ce,handleMouseEnterCounter:Ve,handleMouseLeaveCounter:Ge,handleFocusout:z,handleCompositionEnd:me,handleCompositionStart:le,onPopoverUpdateShow:Ue,focus:Q,focusInput:ge,blur:ye,blurInput:ue,updateCounter:$e,getCounter:we,getTail:je,renderLabel:e.renderLabel,cssVars:Be?void 0:Ne,themeClass:he==null?void 0:he.themeClass,onRender:he==null?void 0:he.onRender}},render(){const{status:e,multiple:n,size:o,disabled:t,filterable:r,maxTagCount:i,bordered:l,clsPrefix:a,onRender:s,renderTag:c,renderLabel:u}=this;s==null||s();const h=i==="responsive",g=typeof i=="number",b=h||g,v=d(ur,null,{default:()=>d(cr,{clsPrefix:a,loading:this.loading,showArrow:this.showArrow,showClear:this.mergedClearable&&this.selected,onClear:this.handleClear},{default:()=>{var R,y;return(y=(R=this.$slots).arrow)===null||y===void 0?void 0:y.call(R)}})});let m;if(n){const{labelField:R}=this,y=z=>d("div",{class:`${a}-base-selection-tag-wrapper`,key:z.value},c?c({option:z,handleClose:()=>this.handleDeleteOption(z)}):d(_n,{size:o,closable:!z.disabled,disabled:t,onClose:()=>this.handleDeleteOption(z),internalCloseIsButtonTag:!1,internalCloseFocusable:!1},{default:()=>u?u(z,!0):xe(z[R],z,!0)})),O=()=>(g?this.selectedOptions.slice(0,i):this.selectedOptions).map(y),D=r?d("div",{class:`${a}-base-selection-input-tag`,ref:"inputTagElRef",key:"__input-tag__"},d("input",Object.assign({},this.inputProps,{ref:"patternInputRef",tabindex:-1,disabled:t,value:this.pattern,autofocus:this.autofocus,class:`${a}-base-selection-input-tag__input`,onBlur:this.handlePatternInputBlur,onFocus:this.handlePatternInputFocus,onKeydown:this.handlePatternKeyDown,onInput:this.handlePatternInputInput,onCompositionstart:this.handleCompositionStart,onCompositionend:this.handleCompositionEnd})),d("span",{ref:"patternInputMirrorRef",class:`${a}-base-selection-input-tag__mirror`},this.pattern)):null,w=h?()=>d("div",{class:`${a}-base-selection-tag-wrapper`,ref:"counterWrapperRef"},d(_n,{size:o,ref:"counterRef",onMouseenter:this.handleMouseEnterCounter,onMouseleave:this.handleMouseLeaveCounter,disabled:t})):void 0;let S;if(g){const z=this.selectedOptions.length-i;z>0&&(S=d("div",{class:`${a}-base-selection-tag-wrapper`,key:"__counter__"},d(_n,{size:o,ref:"counterRef",onMouseenter:this.handleMouseEnterCounter,disabled:t},{default:()=>`+${z}`})))}const _=h?r?d(mo,{ref:"overflowRef",updateCounter:this.updateCounter,getCounter:this.getCounter,getTail:this.getTail,style:{width:"100%",display:"flex",overflow:"hidden"}},{default:O,counter:w,tail:()=>D}):d(mo,{ref:"overflowRef",updateCounter:this.updateCounter,getCounter:this.getCounter,style:{width:"100%",display:"flex",overflow:"hidden"}},{default:O,counter:w}):g?O().concat(S):O(),T=b?()=>d("div",{class:`${a}-base-selection-popover`},h?O():this.selectedOptions.map(y)):void 0,N=b?{show:this.showTagsPanel,trigger:"hover",overlap:!0,placement:"top",width:"trigger",onUpdateShow:this.onPopoverUpdateShow,theme:this.mergedTheme.peers.Popover,themeOverrides:this.mergedTheme.peerOverrides.Popover}:null,k=(this.selected?!1:this.active?!this.pattern&&!this.isComposing:!0)?d("div",{class:`${a}-base-selection-placeholder ${a}-base-selection-overlay`},d("div",{class:`${a}-base-selection-placeholder__inner`},this.placeholder)):null,G=r?d("div",{ref:"patternInputWrapperRef",class:`${a}-base-selection-tags`},_,h?null:D,v):d("div",{ref:"multipleElRef",class:`${a}-base-selection-tags`,tabindex:t?void 0:0},_,v);m=d(Qn,null,b?d(ao,Object.assign({},N,{scrollable:!0,style:"max-height: calc(var(--v-target-height) * 6.6);"}),{trigger:()=>G,default:T}):G,k)}else if(r){const R=this.pattern||this.isComposing,y=this.active?!R:!this.selected,O=this.active?!1:this.selected;m=d("div",{ref:"patternInputWrapperRef",class:`${a}-base-selection-label`},d("input",Object.assign({},this.inputProps,{ref:"patternInputRef",class:`${a}-base-selection-input`,value:this.active?this.pattern:"",placeholder:"",readonly:t,disabled:t,tabindex:-1,autofocus:this.autofocus,onFocus:this.handlePatternInputFocus,onBlur:this.handlePatternInputBlur,onInput:this.handlePatternInputInput,onCompositionstart:this.handleCompositionStart,onCompositionend:this.handleCompositionEnd})),O?d("div",{class:`${a}-base-selection-label__render-label ${a}-base-selection-overlay`,key:"input"},d("div",{class:`${a}-base-selection-overlay__wrapper`},c?c({option:this.selectedOption,handleClose:()=>{}}):u?u(this.selectedOption,!0):xe(this.label,this.selectedOption,!0))):null,y?d("div",{class:`${a}-base-selection-placeholder ${a}-base-selection-overlay`,key:"placeholder"},d("div",{class:`${a}-base-selection-overlay__wrapper`},this.filterablePlaceholder)):null,v)}else m=d("div",{ref:"singleElRef",class:`${a}-base-selection-label`,tabindex:this.disabled?void 0:0},this.label!==void 0?d("div",{class:`${a}-base-selection-input`,title:xr(this.label),key:"input"},d("div",{class:`${a}-base-selection-input__content`},c?c({option:this.selectedOption,handleClose:()=>{}}):u?u(this.selectedOption,!0):xe(this.label,this.selectedOption,!0))):d("div",{class:`${a}-base-selection-placeholder ${a}-base-selection-overlay`,key:"placeholder"},d("div",{class:`${a}-base-selection-placeholder__inner`},this.placeholder)),v);return d("div",{ref:"selfRef",class:[`${a}-base-selection`,this.themeClass,e&&`${a}-base-selection--${e}-status`,{[`${a}-base-selection--active`]:this.active,[`${a}-base-selection--selected`]:this.selected||this.active&&this.pattern,[`${a}-base-selection--disabled`]:this.disabled,[`${a}-base-selection--multiple`]:this.multiple,[`${a}-base-selection--focus`]:this.focused}],style:this.cssVars,onClick:this.onClick,onMouseenter:this.handleMouseEnter,onMouseleave:this.handleMouseLeave,onKeydown:this.onKeydown,onFocusin:this.handleFocusin,onFocusout:this.handleFocusout,onMousedown:this.handleMouseDown},m,l?d("div",{class:`${a}-base-selection__border`}):null,l?d("div",{class:`${a}-base-selection__state-border`}):null)}});function sn(e){return e.type==="group"}function lt(e){return e.type==="ignored"}function Bn(e,n){try{return!!(1+n.toString().toLowerCase().indexOf(e.trim().toLowerCase()))}catch{return!1}}function Xl(e,n){return{getIsGroup:sn,getIgnored:lt,getKey(t){return sn(t)?t.name||t.key||"key-required":t[e]},getChildren(t){return t[n]}}}function Ql(e,n,o,t){if(!n)return e;function r(i){if(!Array.isArray(i))return[];const l=[];for(const a of i)if(sn(a)){const s=r(a[t]);s.length&&l.push(Object.assign({},a,{[t]:s}))}else{if(lt(a))continue;n(o,a)&&l.push(a)}return l}return r(e)}function ea(e,n,o){const t=new Map;return e.forEach(r=>{sn(r)?r[o].forEach(i=>{t.set(i[n],i)}):t.set(r[n],r)}),t}function na(e){const{boxShadow2:n}=e;return{menuBoxShadow:n}}const oa=qe({name:"Select",common:ke,peers:{InternalSelection:it,InternalSelectMenu:ot},self:na}),ta=oa,ra=ee([B("select",`
 z-index: auto;
 outline: none;
 width: 100%;
 position: relative;
 `),B("select-menu",`
 margin: 4px 0;
 box-shadow: var(--n-menu-box-shadow);
 `,[Zn({originalTransition:"background-color .3s var(--n-bezier), box-shadow .3s var(--n-bezier)"})])]),ia=Object.assign(Object.assign({},se.props),{to:_e.propTo,bordered:{type:Boolean,default:void 0},clearable:Boolean,clearFilterAfterSelect:{type:Boolean,default:!0},options:{type:Array,default:()=>[]},defaultValue:{type:[String,Number,Array],default:null},value:[String,Number,Array],placeholder:String,menuProps:Object,multiple:Boolean,size:String,filterable:Boolean,disabled:{type:Boolean,default:void 0},remote:Boolean,loading:Boolean,filter:Function,placement:{type:String,default:"bottom-start"},widthMode:{type:String,default:"trigger"},tag:Boolean,onCreate:Function,fallbackOption:{type:[Function,Boolean],default:void 0},show:{type:Boolean,default:void 0},showArrow:{type:Boolean,default:!0},maxTagCount:[Number,String],consistentMenuWidth:{type:Boolean,default:!0},virtualScroll:{type:Boolean,default:!0},labelField:{type:String,default:"label"},valueField:{type:String,default:"value"},childrenField:{type:String,default:"children"},renderLabel:Function,renderOption:Function,renderTag:Function,"onUpdate:value":[Function,Array],inputProps:Object,nodeProps:Function,ignoreComposition:{type:Boolean,default:!0},showOnFocus:Boolean,onUpdateValue:[Function,Array],onBlur:[Function,Array],onClear:[Function,Array],onFocus:[Function,Array],onScroll:[Function,Array],onSearch:[Function,Array],onUpdateShow:[Function,Array],"onUpdate:show":[Function,Array],displayDirective:{type:String,default:"show"},resetMenuOnOptionsChange:{type:Boolean,default:!0},status:String,showCheckmark:{type:Boolean,default:!0},onChange:[Function,Array],items:Array}),Ma=ie({name:"Select",props:ia,setup(e){const{mergedClsPrefixRef:n,mergedBorderedRef:o,namespaceRef:t,inlineThemeDisabled:r}=We(e),i=se("Select","-select",ra,ta,e,n),l=E(e.defaultValue),a=te(e,"value"),s=ln(a,l),c=E(!1),u=E(""),h=F(()=>{const{valueField:f,childrenField:$}=e,K=Xl(f,$);return et(z.value,K)}),g=F(()=>ea(k.value,e.valueField,e.childrenField)),b=E(!1),v=ln(te(e,"show"),b),m=E(null),R=E(null),y=E(null),{localeRef:O}=Eo("Select"),D=F(()=>{var f;return(f=e.placeholder)!==null&&f!==void 0?f:O.value.placeholder}),w=Uo(e,["items","options"]),S=[],_=E([]),T=E([]),N=E(new Map),L=F(()=>{const{fallbackOption:f}=e;if(f===void 0){const{labelField:$,valueField:K}=e;return J=>({[$]:String(J),[K]:J})}return f===!1?!1:$=>Object.assign(f($),{value:$})}),k=F(()=>T.value.concat(_.value).concat(w.value)),G=F(()=>{const{filter:f}=e;if(f)return f;const{labelField:$,valueField:K}=e;return(J,q)=>{if(!q)return!1;const Y=q[$];if(typeof Y=="string")return Bn(J,Y);const Z=q[K];return typeof Z=="string"?Bn(J,Z):typeof Z=="number"?Bn(J,String(Z)):!1}}),z=F(()=>{if(e.remote)return w.value;{const{value:f}=k,{value:$}=u;return!$.length||!e.filterable?f:Ql(f,G.value,$,e.childrenField)}});function W(f){const $=e.remote,{value:K}=N,{value:J}=g,{value:q}=L,Y=[];return f.forEach(Z=>{if(J.has(Z))Y.push(J.get(Z));else if($&&K.has(Z))Y.push(K.get(Z));else if(q){const de=q(Z);de&&Y.push(de)}}),Y}const x=F(()=>{if(e.multiple){const{value:f}=s;return Array.isArray(f)?W(f):[]}return null}),I=F(()=>{const{value:f}=s;return!e.multiple&&!Array.isArray(f)?f===null?null:W([f])[0]||null:null}),j=fr(e),{mergedSizeRef:A,mergedDisabledRef:V,mergedStatusRef:p}=j;function P(f,$){const{onChange:K,"onUpdate:value":J,onUpdateValue:q}=e,{nTriggerFormChange:Y,nTriggerFormInput:Z}=j;K&&ce(K,f,$),q&&ce(q,f,$),J&&ce(J,f,$),l.value=f,Y(),Z()}function ne(f){const{onBlur:$}=e,{nTriggerFormBlur:K}=j;$&&ce($,f),K()}function le(){const{onClear:f}=e;f&&ce(f)}function me(f){const{onFocus:$,showOnFocus:K}=e,{nTriggerFormFocus:J}=j;$&&ce($,f),J(),K&&ge()}function Ce(f){const{onSearch:$}=e;$&&ce($,f)}function Se(f){const{onScroll:$}=e;$&&ce($,f)}function ye(){var f;const{remote:$,multiple:K}=e;if($){const{value:J}=N;if(K){const{valueField:q}=e;(f=x.value)===null||f===void 0||f.forEach(Y=>{J.set(Y[q],Y)})}else{const q=I.value;q&&J.set(q[e.valueField],q)}}}function Q(f){const{onUpdateShow:$,"onUpdate:show":K}=e;$&&ce($,f),K&&ce(K,f),b.value=f}function ge(){V.value||(Q(!0),b.value=!0,e.filterable&&Xe())}function ue(){Q(!1)}function $e(){u.value="",T.value=S}const we=E(!1);function je(){e.filterable&&(we.value=!0)}function Ie(){e.filterable&&(we.value=!1,v.value||$e())}function Me(){V.value||(v.value?e.filterable?Xe():ue():ge())}function Ve(f){var $,K;!((K=($=y.value)===null||$===void 0?void 0:$.selfRef)===null||K===void 0)&&K.contains(f.relatedTarget)||(c.value=!1,ne(f),ue())}function Ge(f){me(f),c.value=!0}function Ue(f){c.value=!0}function Be(f){var $;!(($=m.value)===null||$===void 0)&&$.$el.contains(f.relatedTarget)||(c.value=!1,ne(f),ue())}function Ne(){var f;(f=m.value)===null||f===void 0||f.focus(),ue()}function he(f){var $;v.value&&(!(($=m.value)===null||$===void 0)&&$.$el.contains(En(f))||ue())}function C(f){if(!Array.isArray(f))return[];if(L.value)return Array.from(f);{const{remote:$}=e,{value:K}=g;if($){const{value:J}=N;return f.filter(q=>K.has(q)||J.has(q))}else return f.filter(J=>K.has(J))}}function M(f){re(f.rawNode)}function re(f){if(V.value)return;const{tag:$,remote:K,clearFilterAfterSelect:J,valueField:q}=e;if($&&!K){const{value:Y}=T,Z=Y[0]||null;if(Z){const de=_.value;de.length?de.push(Z):_.value=[Z],T.value=S}}if(K&&N.value.set(f[q],f),e.multiple){const Y=C(s.value),Z=Y.findIndex(de=>de===f[q]);if(~Z){if(Y.splice(Z,1),$&&!K){const de=Ze(f[q]);~de&&(_.value.splice(de,1),J&&(u.value=""))}}else Y.push(f[q]),J&&(u.value="");P(Y,W(Y))}else{if($&&!K){const Y=Ze(f[q]);~Y?_.value=[_.value[Y]]:_.value=S}Je(),ue(),P(f[q],f)}}function Ze(f){return _.value.findIndex(K=>K[e.valueField]===f)}function mn(f){v.value||ge();const{value:$}=f.target;u.value=$;const{tag:K,remote:J}=e;if(Ce($),K&&!J){if(!$){T.value=S;return}const{onCreate:q}=e,Y=q?q($):{[e.labelField]:$,[e.valueField]:$},{valueField:Z}=e;w.value.some(de=>de[Z]===Y[Z])||_.value.some(de=>de[Z]===Y[Z])?T.value=S:T.value=[Y]}}function yn(f){f.stopPropagation();const{multiple:$}=e;!$&&e.filterable&&ue(),le(),$?P([],[]):P(null,null)}function wn(f){!He(f,"action")&&!He(f,"empty")&&f.preventDefault()}function xn(f){Se(f)}function Ye(f){var $,K,J,q,Y;switch(f.key){case" ":if(e.filterable)break;f.preventDefault();case"Enter":if(!(!(($=m.value)===null||$===void 0)&&$.isComposing)){if(v.value){const Z=(K=y.value)===null||K===void 0?void 0:K.getPendingTmNode();Z?M(Z):e.filterable||(ue(),Je())}else if(ge(),e.tag&&we.value){const Z=T.value[0];if(Z){const de=Z[e.valueField],{value:en}=s;e.multiple&&Array.isArray(en)&&en.some(Pn=>Pn===de)||re(Z)}}}f.preventDefault();break;case"ArrowUp":if(f.preventDefault(),e.loading)return;v.value&&((J=y.value)===null||J===void 0||J.prev());break;case"ArrowDown":if(f.preventDefault(),e.loading)return;v.value?(q=y.value)===null||q===void 0||q.next():ge();break;case"Escape":v.value&&(hr(f),ue()),(Y=m.value)===null||Y===void 0||Y.focus();break}}function Je(){var f;(f=m.value)===null||f===void 0||f.focus()}function Xe(){var f;(f=m.value)===null||f===void 0||f.focusInput()}function Cn(){var f;v.value&&((f=R.value)===null||f===void 0||f.syncPosition())}ye(),Oe(te(e,"options"),ye);const Sn={focus:()=>{var f;(f=m.value)===null||f===void 0||f.focus()},blur:()=>{var f;(f=m.value)===null||f===void 0||f.blur()}},Qe=F(()=>{const{self:{menuBoxShadow:f}}=i.value;return{"--n-menu-box-shadow":f}}),Te=r?ze("select",void 0,Qe,e):void 0;return Object.assign(Object.assign({},Sn),{mergedStatus:p,mergedClsPrefix:n,mergedBordered:o,namespace:t,treeMate:h,isMounted:jo(),triggerRef:m,menuRef:y,pattern:u,uncontrolledShow:b,mergedShow:v,adjustedTo:_e(e),uncontrolledValue:l,mergedValue:s,followerRef:R,localizedPlaceholder:D,selectedOption:I,selectedOptions:x,mergedSize:A,mergedDisabled:V,focused:c,activeWithoutMenuOpen:we,inlineThemeDisabled:r,onTriggerInputFocus:je,onTriggerInputBlur:Ie,handleTriggerOrMenuResize:Cn,handleMenuFocus:Ue,handleMenuBlur:Be,handleMenuTabOut:Ne,handleTriggerClick:Me,handleToggle:M,handleDeleteOption:re,handlePatternInput:mn,handleClear:yn,handleTriggerBlur:Ve,handleTriggerFocus:Ge,handleKeydown:Ye,handleMenuAfterLeave:$e,handleMenuClickOutside:he,handleMenuScroll:xn,handleMenuKeydown:Ye,handleMenuMousedown:wn,mergedTheme:i,cssVars:r?void 0:Qe,themeClass:Te==null?void 0:Te.themeClass,onRender:Te==null?void 0:Te.onRender})},render(){return d("div",{class:`${this.mergedClsPrefix}-select`},d(to,null,{default:()=>[d(oo,null,{default:()=>d(Jl,{ref:"triggerRef",inlineThemeDisabled:this.inlineThemeDisabled,status:this.mergedStatus,inputProps:this.inputProps,clsPrefix:this.mergedClsPrefix,showArrow:this.showArrow,maxTagCount:this.maxTagCount,bordered:this.mergedBordered,active:this.activeWithoutMenuOpen||this.mergedShow,pattern:this.pattern,placeholder:this.localizedPlaceholder,selectedOption:this.selectedOption,selectedOptions:this.selectedOptions,multiple:this.multiple,renderTag:this.renderTag,renderLabel:this.renderLabel,filterable:this.filterable,clearable:this.clearable,disabled:this.mergedDisabled,size:this.mergedSize,theme:this.mergedTheme.peers.InternalSelection,labelField:this.labelField,valueField:this.valueField,themeOverrides:this.mergedTheme.peerOverrides.InternalSelection,loading:this.loading,focused:this.focused,onClick:this.handleTriggerClick,onDeleteOption:this.handleDeleteOption,onPatternInput:this.handlePatternInput,onClear:this.handleClear,onBlur:this.handleTriggerBlur,onFocus:this.handleTriggerFocus,onKeydown:this.handleKeydown,onPatternBlur:this.onTriggerInputBlur,onPatternFocus:this.onTriggerInputFocus,onResize:this.handleTriggerOrMenuResize,ignoreComposition:this.ignoreComposition},{arrow:()=>{var e,n;return[(n=(e=this.$slots).arrow)===null||n===void 0?void 0:n.call(e)]}})}),d(no,{ref:"followerRef",show:this.mergedShow,to:this.adjustedTo,teleportDisabled:this.adjustedTo===_e.tdkey,containerClass:this.namespace,width:this.consistentMenuWidth?"target":void 0,minWidth:"target",placement:this.placement},{default:()=>d(fn,{name:"fade-in-scale-up-transition",appear:this.isMounted,onAfterLeave:this.handleMenuAfterLeave},{default:()=>{var e,n,o;return this.mergedShow||this.displayDirective==="show"?((e=this.onRender)===null||e===void 0||e.call(this),Jn(d($l,Object.assign({},this.menuProps,{ref:"menuRef",onResize:this.handleTriggerOrMenuResize,inlineThemeDisabled:this.inlineThemeDisabled,virtualScroll:this.consistentMenuWidth&&this.virtualScroll,class:[`${this.mergedClsPrefix}-select-menu`,this.themeClass,(n=this.menuProps)===null||n===void 0?void 0:n.class],clsPrefix:this.mergedClsPrefix,focusable:!0,labelField:this.labelField,valueField:this.valueField,autoPending:!0,nodeProps:this.nodeProps,theme:this.mergedTheme.peers.InternalSelectMenu,themeOverrides:this.mergedTheme.peerOverrides.InternalSelectMenu,treeMate:this.treeMate,multiple:this.multiple,size:"medium",renderOption:this.renderOption,renderLabel:this.renderLabel,value:this.mergedValue,style:[(o=this.menuProps)===null||o===void 0?void 0:o.style,this.cssVars],onToggle:this.handleToggle,onScroll:this.handleMenuScroll,onFocus:this.handleMenuFocus,onBlur:this.handleMenuBlur,onKeydown:this.handleMenuKeydown,onTabOut:this.handleMenuTabOut,onMousedown:this.handleMenuMousedown,show:this.mergedShow,showCheckmark:this.showCheckmark,resetMenuOnOptionsChange:this.resetMenuOnOptionsChange}),{empty:()=>{var t,r;return[(r=(t=this.$slots).empty)===null||r===void 0?void 0:r.call(t)]},action:()=>{var t,r;return[(r=(t=this.$slots).action)===null||r===void 0?void 0:r.call(t)]}}),this.displayDirective==="show"?[[Do,this.mergedShow],[rn,this.handleMenuClickOutside,void 0,{capture:!0}]]:[[rn,this.handleMenuClickOutside,void 0,{capture:!0}]])):null}})})]}))}}),la={padding:"8px 14px"},aa=e=>{const{borderRadius:n,boxShadow2:o,baseColor:t}=e;return Object.assign(Object.assign({},la),{borderRadius:n,boxShadow:o,color:vr(t,"rgba(0, 0, 0, .85)"),textColor:t})},sa=qe({name:"Tooltip",common:ke,peers:{Popover:pn},self:aa}),da=sa,ca={padding:"4px 0",optionIconSizeSmall:"14px",optionIconSizeMedium:"16px",optionIconSizeLarge:"16px",optionIconSizeHuge:"18px",optionSuffixWidthSmall:"14px",optionSuffixWidthMedium:"14px",optionSuffixWidthLarge:"16px",optionSuffixWidthHuge:"16px",optionIconSuffixWidthSmall:"32px",optionIconSuffixWidthMedium:"32px",optionIconSuffixWidthLarge:"36px",optionIconSuffixWidthHuge:"36px",optionPrefixWidthSmall:"14px",optionPrefixWidthMedium:"14px",optionPrefixWidthLarge:"16px",optionPrefixWidthHuge:"16px",optionIconPrefixWidthSmall:"36px",optionIconPrefixWidthMedium:"36px",optionIconPrefixWidthLarge:"40px",optionIconPrefixWidthHuge:"40px"},ua=e=>{const{primaryColor:n,textColor2:o,dividerColor:t,hoverColor:r,popoverColor:i,invertedColor:l,borderRadius:a,fontSizeSmall:s,fontSizeMedium:c,fontSizeLarge:u,fontSizeHuge:h,heightSmall:g,heightMedium:b,heightLarge:v,heightHuge:m,textColor3:R,opacityDisabled:y}=e;return Object.assign(Object.assign({},ca),{optionHeightSmall:g,optionHeightMedium:b,optionHeightLarge:v,optionHeightHuge:m,borderRadius:a,fontSizeSmall:s,fontSizeMedium:c,fontSizeLarge:u,fontSizeHuge:h,optionTextColor:o,optionTextColorHover:o,optionTextColorActive:n,optionTextColorChildActive:n,color:i,dividerColor:t,suffixColor:o,prefixColor:o,optionColorHover:r,optionColorActive:X(n,{alpha:.1}),groupHeaderTextColor:R,optionTextColorInverted:"#BBB",optionTextColorHoverInverted:"#FFF",optionTextColorActiveInverted:"#FFF",optionTextColorChildActiveInverted:"#FFF",colorInverted:l,dividerColorInverted:"#BBB",suffixColorInverted:"#BBB",prefixColorInverted:"#BBB",optionColorHoverInverted:n,optionColorActiveInverted:n,groupHeaderTextColorInverted:"#AAA",optionOpacityDisabled:y})},fa=qe({name:"Dropdown",common:ke,peers:{Popover:pn},self:ua}),ha=fa,va=Object.assign(Object.assign({},gn),se.props),Fa=ie({name:"Tooltip",props:va,__popover__:!0,setup(e){const n=se("Tooltip","-tooltip",void 0,da,e),o=E(null);return Object.assign(Object.assign({},{syncPosition(){o.value.syncPosition()},setShow(r){o.value.setShow(r)}}),{popoverRef:o,mergedTheme:n,popoverThemeOverrides:F(()=>n.value.self)})},render(){const{mergedTheme:e,internalExtraClass:n}=this;return d(ao,Object.assign(Object.assign({},this.$props),{theme:e.peers.Popover,themeOverrides:e.peerOverrides.Popover,builtinThemeOverrides:this.popoverThemeOverrides,internalExtraClass:n.concat("tooltip"),ref:"popoverRef"}),this.$slots)}}),at=ie({name:"DropdownDivider",props:{clsPrefix:{type:String,required:!0}},render(){return d("div",{class:`${this.clsPrefix}-dropdown-divider`})}}),pa=e=>{const{textColorBase:n,opacity1:o,opacity2:t,opacity3:r,opacity4:i,opacity5:l}=e;return{color:n,opacity1Depth:o,opacity2Depth:t,opacity3Depth:r,opacity4Depth:i,opacity5Depth:l}},ga={name:"Icon",common:ke,self:pa},ba=ga,ma=B("icon",`
 height: 1em;
 width: 1em;
 line-height: 1em;
 text-align: center;
 display: inline-block;
 position: relative;
 fill: currentColor;
 transform: translateZ(0);
`,[U("color-transition",{transition:"color .3s var(--n-bezier)"}),U("depth",{color:"var(--n-color)"},[ee("svg",{opacity:"var(--n-opacity)",transition:"opacity .3s var(--n-bezier)"})]),ee("svg",{height:"1em",width:"1em"})]),ya=Object.assign(Object.assign({},se.props),{depth:[String,Number],size:[Number,String],color:String,component:Object}),wa=ie({_n_icon__:!0,name:"Icon",inheritAttrs:!1,props:ya,setup(e){const{mergedClsPrefixRef:n,inlineThemeDisabled:o}=We(e),t=se("Icon","-icon",ma,ba,e,n),r=F(()=>{const{depth:l}=e,{common:{cubicBezierEaseInOut:a},self:s}=t.value;if(l!==void 0){const{color:c,[`opacity${l}Depth`]:u}=s;return{"--n-bezier":a,"--n-color":c,"--n-opacity":u}}return{"--n-bezier":a,"--n-color":"","--n-opacity":""}}),i=o?ze("icon",F(()=>`${e.depth||"d"}`),r,e):void 0;return{mergedClsPrefix:n,mergedStyle:F(()=>{const{size:l,color:a}=e;return{fontSize:tn(l),color:a}}),cssVars:o?void 0:r,themeClass:i==null?void 0:i.themeClass,onRender:i==null?void 0:i.onRender}},render(){var e;const{$parent:n,depth:o,mergedClsPrefix:t,component:r,onRender:i,themeClass:l}=this;return!((e=n==null?void 0:n.$options)===null||e===void 0)&&e._n_icon__&&Go("icon","don't wrap `n-icon` inside `n-icon`"),i==null||i(),d("i",hn(this.$attrs,{role:"img",class:[`${t}-icon`,l,{[`${t}-icon--depth`]:o,[`${t}-icon--color-transition`]:o!==void 0}],style:[this.cssVars,this.mergedStyle]}),r?d(r):this.$slots)}}),so=vn("n-dropdown-menu"),bn=vn("n-dropdown"),_o=vn("n-dropdown-option");function jn(e,n){return e.type==="submenu"||e.type===void 0&&e[n]!==void 0}function xa(e){return e.type==="group"}function st(e){return e.type==="divider"}function Ca(e){return e.type==="render"}const dt=ie({name:"DropdownOption",props:{clsPrefix:{type:String,required:!0},tmNode:{type:Object,required:!0},parentKey:{type:[String,Number],default:null},placement:{type:String,default:"right-start"},props:Object,scrollable:Boolean},setup(e){const n=be(bn),{hoverKeyRef:o,keyboardKeyRef:t,lastToggledSubmenuKeyRef:r,pendingKeyPathRef:i,activeKeyPathRef:l,animatedRef:a,mergedShowRef:s,renderLabelRef:c,renderIconRef:u,labelFieldRef:h,childrenFieldRef:g,renderOptionRef:b,nodePropsRef:v,menuPropsRef:m}=n,R=be(_o,null),y=be(so),O=be(Xn),D=F(()=>e.tmNode.rawNode),w=F(()=>{const{value:A}=g;return jn(e.tmNode.rawNode,A)}),S=F(()=>{const{disabled:A}=e.tmNode;return A}),_=F(()=>{if(!w.value)return!1;const{key:A,disabled:V}=e.tmNode;if(V)return!1;const{value:p}=o,{value:P}=t,{value:ne}=r,{value:le}=i;return p!==null?le.includes(A):P!==null?le.includes(A)&&le[le.length-1]!==A:ne!==null?le.includes(A):!1}),T=F(()=>t.value===null&&!a.value),N=Pr(_,300,T),L=F(()=>!!(R!=null&&R.enteringSubmenuRef.value)),k=E(!1);fe(_o,{enteringSubmenuRef:k});function G(){k.value=!0}function z(){k.value=!1}function W(){const{parentKey:A,tmNode:V}=e;V.disabled||s.value&&(r.value=A,t.value=null,o.value=V.key)}function x(){const{tmNode:A}=e;A.disabled||s.value&&o.value!==A.key&&W()}function I(A){if(e.tmNode.disabled||!s.value)return;const{relatedTarget:V}=A;V&&!He({target:V},"dropdownOption")&&!He({target:V},"scrollbarRail")&&(o.value=null)}function j(){const{value:A}=w,{tmNode:V}=e;s.value&&!A&&!V.disabled&&(n.doSelect(V.key,V.rawNode),n.doUpdateShow(!1))}return{labelField:h,renderLabel:c,renderIcon:u,siblingHasIcon:y.showIconRef,siblingHasSubmenu:y.hasSubmenuRef,menuProps:m,popoverBody:O,animated:a,mergedShowSubmenu:F(()=>N.value&&!L.value),rawNode:D,hasSubmenu:w,pending:Re(()=>{const{value:A}=i,{key:V}=e.tmNode;return A.includes(V)}),childActive:Re(()=>{const{value:A}=l,{key:V}=e.tmNode,p=A.findIndex(P=>V===P);return p===-1?!1:p<A.length-1}),active:Re(()=>{const{value:A}=l,{key:V}=e.tmNode,p=A.findIndex(P=>V===P);return p===-1?!1:p===A.length-1}),mergedDisabled:S,renderOption:b,nodeProps:v,handleClick:j,handleMouseMove:x,handleMouseEnter:W,handleMouseLeave:I,handleSubmenuBeforeEnter:G,handleSubmenuAfterEnter:z}},render(){var e,n;const{animated:o,rawNode:t,mergedShowSubmenu:r,clsPrefix:i,siblingHasIcon:l,siblingHasSubmenu:a,renderLabel:s,renderIcon:c,renderOption:u,nodeProps:h,props:g,scrollable:b}=this;let v=null;if(r){const O=(e=this.menuProps)===null||e===void 0?void 0:e.call(this,t,t.children);v=d(ct,Object.assign({},O,{clsPrefix:i,scrollable:this.scrollable,tmNodes:this.tmNode.children,parentKey:this.tmNode.key}))}const m={class:[`${i}-dropdown-option-body`,this.pending&&`${i}-dropdown-option-body--pending`,this.active&&`${i}-dropdown-option-body--active`,this.childActive&&`${i}-dropdown-option-body--child-active`,this.mergedDisabled&&`${i}-dropdown-option-body--disabled`],onMousemove:this.handleMouseMove,onMouseenter:this.handleMouseEnter,onMouseleave:this.handleMouseLeave,onClick:this.handleClick},R=h==null?void 0:h(t),y=d("div",Object.assign({class:[`${i}-dropdown-option`,R==null?void 0:R.class],"data-dropdown-option":!0},R),d("div",hn(m,g),[d("div",{class:[`${i}-dropdown-option-body__prefix`,l&&`${i}-dropdown-option-body__prefix--show-icon`]},[c?c(t):xe(t.icon)]),d("div",{"data-dropdown-option":!0,class:`${i}-dropdown-option-body__label`},s?s(t):xe((n=t[this.labelField])!==null&&n!==void 0?n:t.title)),d("div",{"data-dropdown-option":!0,class:[`${i}-dropdown-option-body__suffix`,a&&`${i}-dropdown-option-body__suffix--has-submenu`]},this.hasSubmenu?d(wa,null,{default:()=>d(Vi,null)}):null)]),this.hasSubmenu?d(to,null,{default:()=>[d(oo,null,{default:()=>d("div",{class:`${i}-dropdown-offset-container`},d(no,{show:this.mergedShowSubmenu,placement:this.placement,to:b&&this.popoverBody||void 0,teleportDisabled:!b},{default:()=>d("div",{class:`${i}-dropdown-menu-wrapper`},o?d(fn,{onBeforeEnter:this.handleSubmenuBeforeEnter,onAfterEnter:this.handleSubmenuAfterEnter,name:"fade-in-scale-up-transition",appear:!0},{default:()=>v}):v)}))})]}):null);return u?u({node:y,option:t}):y}}),Sa=ie({name:"DropdownGroupHeader",props:{clsPrefix:{type:String,required:!0},tmNode:{type:Object,required:!0}},setup(){const{showIconRef:e,hasSubmenuRef:n}=be(so),{renderLabelRef:o,labelFieldRef:t,nodePropsRef:r,renderOptionRef:i}=be(bn);return{labelField:t,showIcon:e,hasSubmenu:n,renderLabel:o,nodeProps:r,renderOption:i}},render(){var e;const{clsPrefix:n,hasSubmenu:o,showIcon:t,nodeProps:r,renderLabel:i,renderOption:l}=this,{rawNode:a}=this.tmNode,s=d("div",Object.assign({class:`${n}-dropdown-option`},r==null?void 0:r(a)),d("div",{class:`${n}-dropdown-option-body ${n}-dropdown-option-body--group`},d("div",{"data-dropdown-option":!0,class:[`${n}-dropdown-option-body__prefix`,t&&`${n}-dropdown-option-body__prefix--show-icon`]},xe(a.icon)),d("div",{class:`${n}-dropdown-option-body__label`,"data-dropdown-option":!0},i?i(a):xe((e=a.title)!==null&&e!==void 0?e:a[this.labelField])),d("div",{class:[`${n}-dropdown-option-body__suffix`,o&&`${n}-dropdown-option-body__suffix--has-submenu`],"data-dropdown-option":!0})));return l?l({node:s,option:a}):s}}),Pa=ie({name:"NDropdownGroup",props:{clsPrefix:{type:String,required:!0},tmNode:{type:Object,required:!0},parentKey:{type:[String,Number],default:null}},render(){const{tmNode:e,parentKey:n,clsPrefix:o}=this,{children:t}=e;return d(Qn,null,d(Sa,{clsPrefix:o,tmNode:e,key:e.key}),t==null?void 0:t.map(r=>{const{rawNode:i}=r;return i.show===!1?null:st(i)?d(at,{clsPrefix:o,key:r.key}):r.isGroup?(Go("dropdown","`group` node is not allowed to be put in `group` node."),null):d(dt,{clsPrefix:o,tmNode:r,parentKey:n,key:r.key})}))}}),Oa=ie({name:"DropdownRenderOption",props:{tmNode:{type:Object,required:!0}},render(){const{rawNode:{render:e,props:n}}=this.tmNode;return d("div",n,[e==null?void 0:e()])}}),ct=ie({name:"DropdownMenu",props:{scrollable:Boolean,showArrow:Boolean,arrowStyle:[String,Object],clsPrefix:{type:String,required:!0},tmNodes:{type:Array,default:()=>[]},parentKey:{type:[String,Number],default:null}},setup(e){const{renderIconRef:n,childrenFieldRef:o}=be(bn);fe(so,{showIconRef:F(()=>{const r=n.value;return e.tmNodes.some(i=>{var l;if(i.isGroup)return(l=i.children)===null||l===void 0?void 0:l.some(({rawNode:s})=>r?r(s):s.icon);const{rawNode:a}=i;return r?r(a):a.icon})}),hasSubmenuRef:F(()=>{const{value:r}=o;return e.tmNodes.some(i=>{var l;if(i.isGroup)return(l=i.children)===null||l===void 0?void 0:l.some(({rawNode:s})=>jn(s,r));const{rawNode:a}=i;return jn(a,r)})})});const t=E(null);return fe(Ko,null),fe(Ho,null),fe(Xn,t),{bodyRef:t}},render(){const{parentKey:e,clsPrefix:n,scrollable:o}=this,t=this.tmNodes.map(r=>{const{rawNode:i}=r;return i.show===!1?null:Ca(i)?d(Oa,{tmNode:r,key:r.key}):st(i)?d(at,{clsPrefix:n,key:r.key}):xa(i)?d(Pa,{clsPrefix:n,tmNode:r,parentKey:e,key:r.key}):d(dt,{clsPrefix:n,tmNode:r,parentKey:e,key:r.key,props:i.props,scrollable:o})});return d("div",{class:[`${n}-dropdown-menu`,o&&`${n}-dropdown-menu--scrollable`],ref:"bodyRef"},o?d(Wo,{contentClass:`${n}-dropdown-menu__content`},{default:()=>t}):t,this.showArrow?rt({clsPrefix:n,arrowStyle:this.arrowStyle}):null)}}),Ra=B("dropdown-menu",`
 transform-origin: var(--v-transform-origin);
 background-color: var(--n-color);
 border-radius: var(--n-border-radius);
 box-shadow: var(--n-box-shadow);
 position: relative;
 transition:
 background-color .3s var(--n-bezier),
 box-shadow .3s var(--n-bezier);
`,[Zn(),B("dropdown-option",`
 position: relative;
 `,[ee("a",`
 text-decoration: none;
 color: inherit;
 outline: none;
 `,[ee("&::before",`
 content: "";
 position: absolute;
 left: 0;
 right: 0;
 top: 0;
 bottom: 0;
 `)]),B("dropdown-option-body",`
 display: flex;
 cursor: pointer;
 position: relative;
 height: var(--n-option-height);
 line-height: var(--n-option-height);
 font-size: var(--n-font-size);
 color: var(--n-option-text-color);
 transition: color .3s var(--n-bezier);
 `,[ee("&::before",`
 content: "";
 position: absolute;
 top: 0;
 bottom: 0;
 left: 4px;
 right: 4px;
 transition: background-color .3s var(--n-bezier);
 border-radius: var(--n-border-radius);
 `),pe("disabled",[U("pending",`
 color: var(--n-option-text-color-hover);
 `,[H("prefix, suffix",`
 color: var(--n-option-text-color-hover);
 `),ee("&::before","background-color: var(--n-option-color-hover);")]),U("active",`
 color: var(--n-option-text-color-active);
 `,[H("prefix, suffix",`
 color: var(--n-option-text-color-active);
 `),ee("&::before","background-color: var(--n-option-color-active);")]),U("child-active",`
 color: var(--n-option-text-color-child-active);
 `,[H("prefix, suffix",`
 color: var(--n-option-text-color-child-active);
 `)])]),U("disabled",`
 cursor: not-allowed;
 opacity: var(--n-option-opacity-disabled);
 `),U("group",`
 font-size: calc(var(--n-font-size) - 1px);
 color: var(--n-group-header-text-color);
 `,[H("prefix",`
 width: calc(var(--n-option-prefix-width) / 2);
 `,[U("show-icon",`
 width: calc(var(--n-option-icon-prefix-width) / 2);
 `)])]),H("prefix",`
 width: var(--n-option-prefix-width);
 display: flex;
 justify-content: center;
 align-items: center;
 color: var(--n-prefix-color);
 transition: color .3s var(--n-bezier);
 z-index: 1;
 `,[U("show-icon",`
 width: var(--n-option-icon-prefix-width);
 `),B("icon",`
 font-size: var(--n-option-icon-size);
 `)]),H("label",`
 white-space: nowrap;
 flex: 1;
 z-index: 1;
 `),H("suffix",`
 box-sizing: border-box;
 flex-grow: 0;
 flex-shrink: 0;
 display: flex;
 justify-content: flex-end;
 align-items: center;
 min-width: var(--n-option-suffix-width);
 padding: 0 8px;
 transition: color .3s var(--n-bezier);
 color: var(--n-suffix-color);
 z-index: 1;
 `,[U("has-submenu",`
 width: var(--n-option-icon-suffix-width);
 `),B("icon",`
 font-size: var(--n-option-icon-size);
 `)]),B("dropdown-menu","pointer-events: all;")]),B("dropdown-offset-container",`
 pointer-events: none;
 position: absolute;
 left: 0;
 right: 0;
 top: -4px;
 bottom: -4px;
 `)]),B("dropdown-divider",`
 transition: background-color .3s var(--n-bezier);
 background-color: var(--n-divider-color);
 height: 1px;
 margin: 4px 0;
 `),B("dropdown-menu-wrapper",`
 transform-origin: var(--v-transform-origin);
 width: fit-content;
 `),ee(">",[B("scrollbar",`
 height: inherit;
 max-height: inherit;
 `)]),pe("scrollable",`
 padding: var(--n-padding);
 `),U("scrollable",[H("content",`
 padding: var(--n-padding);
 `)])]),ka={animated:{type:Boolean,default:!0},keyboard:{type:Boolean,default:!0},size:{type:String,default:"medium"},inverted:Boolean,placement:{type:String,default:"bottom"},onSelect:[Function,Array],options:{type:Array,default:()=>[]},menuProps:Function,showArrow:Boolean,renderLabel:Function,renderIcon:Function,renderOption:Function,nodeProps:Function,labelField:{type:String,default:"label"},keyField:{type:String,default:"key"},childrenField:{type:String,default:"children"},value:[String,Number]},$a=Object.keys(gn),Ta=Object.assign(Object.assign(Object.assign({},gn),ka),se.props),Aa=ie({name:"Dropdown",inheritAttrs:!1,props:Ta,setup(e){const n=E(!1),o=ln(te(e,"show"),n),t=F(()=>{const{keyField:z,childrenField:W}=e;return et(e.options,{getKey(x){return x[z]},getDisabled(x){return x.disabled===!0},getIgnored(x){return x.type==="divider"||x.type==="render"},getChildren(x){return x[W]}})}),r=F(()=>t.value.treeNodes),i=E(null),l=E(null),a=E(null),s=F(()=>{var z,W,x;return(x=(W=(z=i.value)!==null&&z!==void 0?z:l.value)!==null&&W!==void 0?W:a.value)!==null&&x!==void 0?x:null}),c=F(()=>t.value.getPath(s.value).keyPath),u=F(()=>t.value.getPath(e.value).keyPath),h=Re(()=>e.keyboard&&o.value);wr({keydown:{ArrowUp:{prevent:!0,handler:S},ArrowRight:{prevent:!0,handler:w},ArrowDown:{prevent:!0,handler:_},ArrowLeft:{prevent:!0,handler:D},Enter:{prevent:!0,handler:T},Escape:O}},h);const{mergedClsPrefixRef:g,inlineThemeDisabled:b}=We(e),v=se("Dropdown","-dropdown",Ra,ha,e,g);fe(bn,{labelFieldRef:te(e,"labelField"),childrenFieldRef:te(e,"childrenField"),renderLabelRef:te(e,"renderLabel"),renderIconRef:te(e,"renderIcon"),hoverKeyRef:i,keyboardKeyRef:l,lastToggledSubmenuKeyRef:a,pendingKeyPathRef:c,activeKeyPathRef:u,animatedRef:te(e,"animated"),mergedShowRef:o,nodePropsRef:te(e,"nodeProps"),renderOptionRef:te(e,"renderOption"),menuPropsRef:te(e,"menuProps"),doSelect:m,doUpdateShow:R}),Oe(o,z=>{!e.animated&&!z&&y()});function m(z,W){const{onSelect:x}=e;x&&ce(x,z,W)}function R(z){const{"onUpdate:show":W,onUpdateShow:x}=e;W&&ce(W,z),x&&ce(x,z),n.value=z}function y(){i.value=null,l.value=null,a.value=null}function O(){R(!1)}function D(){L("left")}function w(){L("right")}function S(){L("up")}function _(){L("down")}function T(){const z=N();z!=null&&z.isLeaf&&o.value&&(m(z.key,z.rawNode),R(!1))}function N(){var z;const{value:W}=t,{value:x}=s;return!W||x===null?null:(z=W.getNode(x))!==null&&z!==void 0?z:null}function L(z){const{value:W}=s,{value:{getFirstAvailableNode:x}}=t;let I=null;if(W===null){const j=x();j!==null&&(I=j.key)}else{const j=N();if(j){let A;switch(z){case"down":A=j.getNext();break;case"up":A=j.getPrev();break;case"right":A=j.getChild();break;case"left":A=j.getParent();break}A&&(I=A.key)}}I!==null&&(i.value=null,l.value=I)}const k=F(()=>{const{size:z,inverted:W}=e,{common:{cubicBezierEaseInOut:x},self:I}=v.value,{padding:j,dividerColor:A,borderRadius:V,optionOpacityDisabled:p,[oe("optionIconSuffixWidth",z)]:P,[oe("optionSuffixWidth",z)]:ne,[oe("optionIconPrefixWidth",z)]:le,[oe("optionPrefixWidth",z)]:me,[oe("fontSize",z)]:Ce,[oe("optionHeight",z)]:Se,[oe("optionIconSize",z)]:ye}=I,Q={"--n-bezier":x,"--n-font-size":Ce,"--n-padding":j,"--n-border-radius":V,"--n-option-height":Se,"--n-option-prefix-width":me,"--n-option-icon-prefix-width":le,"--n-option-suffix-width":ne,"--n-option-icon-suffix-width":P,"--n-option-icon-size":ye,"--n-divider-color":A,"--n-option-opacity-disabled":p};return W?(Q["--n-color"]=I.colorInverted,Q["--n-option-color-hover"]=I.optionColorHoverInverted,Q["--n-option-color-active"]=I.optionColorActiveInverted,Q["--n-option-text-color"]=I.optionTextColorInverted,Q["--n-option-text-color-hover"]=I.optionTextColorHoverInverted,Q["--n-option-text-color-active"]=I.optionTextColorActiveInverted,Q["--n-option-text-color-child-active"]=I.optionTextColorChildActiveInverted,Q["--n-prefix-color"]=I.prefixColorInverted,Q["--n-suffix-color"]=I.suffixColorInverted,Q["--n-group-header-text-color"]=I.groupHeaderTextColorInverted):(Q["--n-color"]=I.color,Q["--n-option-color-hover"]=I.optionColorHover,Q["--n-option-color-active"]=I.optionColorActive,Q["--n-option-text-color"]=I.optionTextColor,Q["--n-option-text-color-hover"]=I.optionTextColorHover,Q["--n-option-text-color-active"]=I.optionTextColorActive,Q["--n-option-text-color-child-active"]=I.optionTextColorChildActive,Q["--n-prefix-color"]=I.prefixColor,Q["--n-suffix-color"]=I.suffixColor,Q["--n-group-header-text-color"]=I.groupHeaderTextColor),Q}),G=b?ze("dropdown",F(()=>`${e.size[0]}${e.inverted?"i":""}`),k,e):void 0;return{mergedClsPrefix:g,mergedTheme:v,tmNodes:r,mergedShow:o,handleAfterLeave:()=>{e.animated&&y()},doUpdateShow:R,cssVars:b?void 0:k,themeClass:G==null?void 0:G.themeClass,onRender:G==null?void 0:G.onRender}},render(){const e=(t,r,i,l,a)=>{var s;const{mergedClsPrefix:c,menuProps:u}=this;(s=this.onRender)===null||s===void 0||s.call(this);const h=(u==null?void 0:u(void 0,this.tmNodes.map(b=>b.rawNode)))||{},g={ref:Cr(r),class:[t,`${c}-dropdown`,this.themeClass],clsPrefix:c,tmNodes:this.tmNodes,style:[i,this.cssVars],showArrow:this.showArrow,arrowStyle:this.arrowStyle,scrollable:this.scrollable,onMouseenter:l,onMouseleave:a};return d(ct,hn(this.$attrs,g,h))},{mergedTheme:n}=this,o={show:this.mergedShow,theme:n.peers.Popover,themeOverrides:n.peerOverrides.Popover,internalOnAfterLeave:this.handleAfterLeave,internalRenderBody:e,onUpdateShow:this.doUpdateShow,"onUpdate:show":void 0};return d(ao,Object.assign({},Vo(this.$props,$a),o),{trigger:()=>{var t,r;return(r=(t=this.$slots).default)===null||r===void 0?void 0:r.call(t)}})}});export{Vi as C,Aa as N,Fa as a,wa as b,et as c,ha as d,_n as e,nt as f,xr as g,Cl as h,Ma as i,Hi as j,ao as k,rl as l,pl as m,gn as n,ot as o,pn as p,$l as q,Xl as r,Cr as s,da as t,Uo as u,$n as v,ta as w};
