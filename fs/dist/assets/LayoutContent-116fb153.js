import{c as L,a as R,ab as j,ad as m,k as x,e as C,f as O,d as $,u as y,m as g,g as z,p as E,b7 as P,i as S,j as w,h as f,af as N}from"./index-1e5c8491.js";const F=u=>{const{baseColor:e,textColor2:o,bodyColor:a,cardColor:i,dividerColor:t,actionColor:d,scrollbarColor:b,scrollbarColorHover:h,invertedColor:s}=u;return{textColor:o,textColorInverted:"#FFF",color:a,colorEmbedded:d,headerColor:i,headerColorInverted:s,footerColor:d,footerColorInverted:s,headerBorderColor:t,headerBorderColorInverted:s,footerBorderColor:t,footerBorderColorInverted:s,siderBorderColor:t,siderBorderColorInverted:s,siderColor:i,siderColorInverted:s,siderToggleButtonBorder:`1px solid ${t}`,siderToggleButtonColor:e,siderToggleButtonIconColor:o,siderToggleButtonIconColorInverted:o,siderToggleBarColor:m(a,b),siderToggleBarColorHover:m(a,h),__invertScrollbar:"true"}},_=L({name:"Layout",common:R,peers:{Scrollbar:j},self:F}),k=_,W=x("n-layout-sider"),H={type:String,default:"static"},K=C("layout",`
 color: var(--n-text-color);
 background-color: var(--n-color);
 box-sizing: border-box;
 position: relative;
 z-index: auto;
 flex: auto;
 overflow: hidden;
 transition:
 box-shadow .3s var(--n-bezier),
 background-color .3s var(--n-bezier),
 color .3s var(--n-bezier);
`,[C("layout-scroll-container",`
 overflow-x: hidden;
 box-sizing: border-box;
 height: 100%;
 `),O("absolute-positioned",`
 position: absolute;
 left: 0;
 right: 0;
 top: 0;
 bottom: 0;
 `)]),V={embedded:Boolean,position:H,nativeScrollbar:{type:Boolean,default:!0},scrollbarProps:Object,onScroll:Function,contentStyle:{type:[String,Object],default:""},hasSider:Boolean,siderPlacement:{type:String,default:"left"}},D=x("n-layout");function p(u){return $({name:u?"LayoutContent":"Layout",props:Object.assign(Object.assign({},y.props),V),setup(e){const o=g(null),a=g(null),{mergedClsPrefixRef:i,inlineThemeDisabled:t}=z(e),d=y("Layout","-layout",K,k,e,i);function b(r,l){if(e.nativeScrollbar){const{value:n}=o;n&&(l===void 0?n.scrollTo(r):n.scrollTo(r,l))}else{const{value:n}=a;n&&n.scrollTo(r,l)}}E(D,e);let h=0,s=0;const T=r=>{var l;const n=r.target;h=n.scrollLeft,s=n.scrollTop,(l=e.onScroll)===null||l===void 0||l.call(e,r)};P(()=>{if(e.nativeScrollbar){const r=o.value;r&&(r.scrollTop=s,r.scrollLeft=h)}});const B={display:"flex",flexWrap:"nowrap",width:"100%",flexDirection:"row"},I={scrollTo:b},v=S(()=>{const{common:{cubicBezierEaseInOut:r},self:l}=d.value;return{"--n-bezier":r,"--n-color":e.embedded?l.colorEmbedded:l.color,"--n-text-color":l.textColor}}),c=t?w("layout",S(()=>e.embedded?"e":""),v,e):void 0;return Object.assign({mergedClsPrefix:i,scrollableElRef:o,scrollbarInstRef:a,hasSiderStyle:B,mergedTheme:d,handleNativeElScroll:T,cssVars:t?void 0:v,themeClass:c==null?void 0:c.themeClass,onRender:c==null?void 0:c.onRender},I)},render(){var e;const{mergedClsPrefix:o,hasSider:a}=this;(e=this.onRender)===null||e===void 0||e.call(this);const i=a?this.hasSiderStyle:void 0,t=[this.themeClass,u&&`${o}-layout-content`,`${o}-layout`,`${o}-layout--${this.position}-positioned`];return f("div",{class:t,style:this.cssVars},this.nativeScrollbar?f("div",{ref:"scrollableElRef",class:`${o}-layout-scroll-container`,style:[this.contentStyle,i],onScroll:this.handleNativeElScroll},this.$slots):f(N,Object.assign({},this.scrollbarProps,{onScroll:this.onScroll,ref:"scrollbarInstRef",theme:this.mergedTheme.peers.Scrollbar,themeOverrides:this.mergedTheme.peerOverrides.Scrollbar,contentStyle:[this.contentStyle,i]}),this.$slots))}})}const X=p(!1),Y=p(!0);export{Y as N,W as a,X as b,D as c,k as l,H as p};
