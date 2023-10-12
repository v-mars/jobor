import{a as b,t as m,e as c,c0 as z,f as u,d as S,g as C,u as h,i as p,j as x,h as r,b1 as $,ay as k,a$ as T,a1 as w}from"./index-3a30b39c.js";import{a as R}from"./Icon-026626f3.js";const B=i=>{const{opacityDisabled:a,heightTiny:e,heightSmall:t,heightMedium:l,heightLarge:s,heightHuge:n,primaryColor:o,fontSize:d}=i;return{fontSize:d,textColor:o,sizeTiny:e,sizeSmall:t,sizeMedium:l,sizeLarge:s,sizeHuge:n,color:o,opacitySpinning:a}},L={name:"Spin",common:b,self:B},N=L,P=m([m("@keyframes spin-rotate",`
 from {
 transform: rotate(0);
 }
 to {
 transform: rotate(360deg);
 }
 `),c("spin-container",{position:"relative"},[c("spin-body",`
 position: absolute;
 top: 50%;
 left: 50%;
 transform: translateX(-50%) translateY(-50%);
 `,[z()])]),c("spin-body",`
 display: inline-flex;
 align-items: center;
 justify-content: center;
 flex-direction: column;
 `),c("spin",`
 display: inline-flex;
 height: var(--n-size);
 width: var(--n-size);
 font-size: var(--n-size);
 color: var(--n-color);
 `,[u("rotate",`
 animation: spin-rotate 2s linear infinite;
 `)]),c("spin-description",`
 display: inline-block;
 font-size: var(--n-font-size);
 color: var(--n-text-color);
 transition: color .3s var(--n-bezier);
 margin-top: 8px;
 `),c("spin-content",`
 opacity: 1;
 transition: opacity .3s var(--n-bezier);
 pointer-events: all;
 `,[u("spinning",`
 user-select: none;
 -webkit-user-select: none;
 pointer-events: none;
 opacity: var(--n-opacity-spinning);
 `)])]),V={small:20,medium:18,large:16},W=Object.assign(Object.assign({},h.props),{description:String,stroke:String,size:{type:[String,Number],default:"medium"},show:{type:Boolean,default:!0},strokeWidth:Number,rotate:{type:Boolean,default:!0},spinning:{type:Boolean,validator:()=>!0,default:void 0}}),O=S({name:"Spin",props:W,setup(i){const{mergedClsPrefixRef:a,inlineThemeDisabled:e}=C(i),t=h("Spin","-spin",P,N,i,a),l=p(()=>{const{size:n}=i,{common:{cubicBezierEaseInOut:o},self:d}=t.value,{opacitySpinning:f,color:g,textColor:y}=d,v=typeof n=="number"?T(n):d[w("size",n)];return{"--n-bezier":o,"--n-opacity-spinning":f,"--n-size":v,"--n-color":g,"--n-text-color":y}}),s=e?x("spin",p(()=>{const{size:n}=i;return typeof n=="number"?String(n):n[0]}),l,i):void 0;return{mergedClsPrefix:a,compitableShow:R(i,["spinning","show"]),mergedStrokeWidth:p(()=>{const{strokeWidth:n}=i;if(n!==void 0)return n;const{size:o}=i;return V[typeof o=="number"?"medium":o]}),cssVars:e?void 0:l,themeClass:s==null?void 0:s.themeClass,onRender:s==null?void 0:s.onRender}},render(){var i,a;const{$slots:e,mergedClsPrefix:t,description:l}=this,s=e.icon&&this.rotate,n=(l||e.description)&&r("div",{class:`${t}-spin-description`},l||((i=e.description)===null||i===void 0?void 0:i.call(e))),o=e.icon?r("div",{class:[`${t}-spin-body`,this.themeClass]},r("div",{class:[`${t}-spin`,s&&`${t}-spin--rotate`],style:e.default?"":this.cssVars},e.icon()),n):r("div",{class:[`${t}-spin-body`,this.themeClass]},r($,{clsPrefix:t,style:e.default?"":this.cssVars,stroke:this.stroke,"stroke-width":this.mergedStrokeWidth,class:`${t}-spin`}),n);return(a=this.onRender)===null||a===void 0||a.call(this),e.default?r("div",{class:[`${t}-spin-container`,this.themeClass],style:this.cssVars},r("div",{class:[`${t}-spin-content`,this.compitableShow&&`${t}-spin-content--spinning`]},e),r(k,{name:"fade-in-transition"},{default:()=>this.compitableShow?o:null})):o}});export{O as N};
