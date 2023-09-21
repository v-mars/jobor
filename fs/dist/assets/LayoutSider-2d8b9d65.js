import{e as n,f as d,v as l,t as h,d as C,h as s,N as A,u as P,m as y,i as p,z as _,b7 as F,p as W,g as U,j as V,af as H,l as X,A as m}from"./index-969e67fe.js";import{C as q}from"./Dropdown-cd7deeac.js";import{a as D,p as K,c as G,l as J}from"./LayoutContent-b2bb953b.js";import{u as Q,f as x}from"./Icon-32b66fa9.js";const Z=n("layout-sider",`
 flex-shrink: 0;
 box-sizing: border-box;
 position: relative;
 z-index: 1;
 color: var(--n-text-color);
 transition:
 color .3s var(--n-bezier),
 border-color .3s var(--n-bezier),
 min-width .3s var(--n-bezier),
 max-width .3s var(--n-bezier),
 transform .3s var(--n-bezier),
 background-color .3s var(--n-bezier);
 background-color: var(--n-color);
 display: flex;
 justify-content: flex-end;
`,[d("bordered",[l("border",`
 content: "";
 position: absolute;
 top: 0;
 bottom: 0;
 width: 1px;
 background-color: var(--n-border-color);
 transition: background-color .3s var(--n-bezier);
 `)]),l("left-placement",[d("bordered",[l("border",`
 right: 0;
 `)])]),d("right-placement",`
 justify-content: flex-start;
 `,[d("bordered",[l("border",`
 left: 0;
 `)]),d("collapsed",[n("layout-toggle-button",[n("base-icon",`
 transform: rotate(180deg);
 `)]),n("layout-toggle-bar",[h("&:hover",[l("top",{transform:"rotate(-12deg) scale(1.15) translateY(-2px)"}),l("bottom",{transform:"rotate(12deg) scale(1.15) translateY(2px)"})])])]),n("layout-toggle-button",`
 left: 0;
 transform: translateX(-50%) translateY(-50%);
 `,[n("base-icon",`
 transform: rotate(0);
 `)]),n("layout-toggle-bar",`
 left: -28px;
 transform: rotate(180deg);
 `,[h("&:hover",[l("top",{transform:"rotate(12deg) scale(1.15) translateY(-2px)"}),l("bottom",{transform:"rotate(-12deg) scale(1.15) translateY(2px)"})])])]),d("collapsed",[n("layout-toggle-bar",[h("&:hover",[l("top",{transform:"rotate(-12deg) scale(1.15) translateY(-2px)"}),l("bottom",{transform:"rotate(12deg) scale(1.15) translateY(2px)"})])]),n("layout-toggle-button",[n("base-icon",`
 transform: rotate(0);
 `)])]),n("layout-toggle-button",`
 transition:
 color .3s var(--n-bezier),
 right .3s var(--n-bezier),
 left .3s var(--n-bezier),
 border-color .3s var(--n-bezier),
 background-color .3s var(--n-bezier);
 cursor: pointer;
 width: 24px;
 height: 24px;
 position: absolute;
 top: 50%;
 right: 0;
 border-radius: 50%;
 display: flex;
 align-items: center;
 justify-content: center;
 font-size: 18px;
 color: var(--n-toggle-button-icon-color);
 border: var(--n-toggle-button-border);
 background-color: var(--n-toggle-button-color);
 box-shadow: 0 2px 4px 0px rgba(0, 0, 0, .06);
 transform: translateX(50%) translateY(-50%);
 z-index: 1;
 `,[n("base-icon",`
 transition: transform .3s var(--n-bezier);
 transform: rotate(180deg);
 `)]),n("layout-toggle-bar",`
 cursor: pointer;
 height: 72px;
 width: 32px;
 position: absolute;
 top: calc(50% - 36px);
 right: -28px;
 `,[l("top, bottom",`
 position: absolute;
 width: 4px;
 border-radius: 2px;
 height: 38px;
 left: 14px;
 transition: 
 background-color .3s var(--n-bezier),
 transform .3s var(--n-bezier);
 `),l("bottom",`
 position: absolute;
 top: 34px;
 `),h("&:hover",[l("top",{transform:"rotate(12deg) scale(1.15) translateY(-2px)"}),l("bottom",{transform:"rotate(-12deg) scale(1.15) translateY(2px)"})]),l("top, bottom",{backgroundColor:"var(--n-toggle-bar-color)"}),h("&:hover",[l("top, bottom",{backgroundColor:"var(--n-toggle-bar-color-hover)"})])]),l("border",`
 position: absolute;
 top: 0;
 right: 0;
 bottom: 0;
 width: 1px;
 transition: background-color .3s var(--n-bezier);
 `),n("layout-sider-scroll-container",`
 flex-grow: 1;
 flex-shrink: 0;
 box-sizing: border-box;
 height: 100%;
 opacity: 0;
 transition: opacity .3s var(--n-bezier);
 max-width: 100%;
 `),d("show-content",[n("layout-sider-scroll-container",{opacity:1})]),d("absolute-positioned",`
 position: absolute;
 left: 0;
 top: 0;
 bottom: 0;
 `)]),ee=C({name:"LayoutToggleButton",props:{clsPrefix:{type:String,required:!0},onClick:Function},render(){const{clsPrefix:e}=this;return s("div",{class:`${e}-layout-toggle-button`,onClick:this.onClick},s(A,{clsPrefix:e},{default:()=>s(q,null)}))}}),oe=C({props:{clsPrefix:{type:String,required:!0},onClick:Function},render(){const{clsPrefix:e}=this;return s("div",{onClick:this.onClick,class:`${e}-layout-toggle-bar`},s("div",{class:`${e}-layout-toggle-bar__top`}),s("div",{class:`${e}-layout-toggle-bar__bottom`}))}}),te={position:K,bordered:Boolean,collapsedWidth:{type:Number,default:48},width:{type:[Number,String],default:272},contentStyle:{type:[String,Object],default:""},collapseMode:{type:String,default:"transform"},collapsed:{type:Boolean,default:void 0},defaultCollapsed:Boolean,showCollapsedContent:{type:Boolean,default:!0},showTrigger:{type:[Boolean,String],default:!1},nativeScrollbar:{type:Boolean,default:!0},inverted:Boolean,scrollbarProps:Object,triggerStyle:[String,Object],collapsedTriggerStyle:[String,Object],"onUpdate:collapsed":[Function,Array],onUpdateCollapsed:[Function,Array],onAfterEnter:Function,onAfterLeave:Function,onExpand:[Function,Array],onCollapse:[Function,Array],onScroll:Function},se=C({name:"LayoutSider",props:Object.assign(Object.assign({},P.props),te),setup(e){const a=X(G),c=y(null),b=y(null),$=p(()=>x(f.value?e.collapsedWidth:e.width)),j=p(()=>e.collapseMode!=="transform"?{}:{minWidth:x(e.width)}),I=p(()=>a?a.siderPlacement:"left"),S=y(e.defaultCollapsed),f=Q(_(e,"collapsed"),S);function O(r,o){if(e.nativeScrollbar){const{value:t}=c;t&&(o===void 0?t.scrollTo(r):t.scrollTo(r,o))}else{const{value:t}=b;t&&t.scrollTo(r,o)}}function E(){const{"onUpdate:collapsed":r,onUpdateCollapsed:o,onExpand:t,onCollapse:v}=e,{value:g}=f;o&&m(o,!g),r&&m(r,!g),S.value=!g,g?t&&m(t):v&&m(v)}let T=0,z=0;const Y=r=>{var o;const t=r.target;T=t.scrollLeft,z=t.scrollTop,(o=e.onScroll)===null||o===void 0||o.call(e,r)};F(()=>{if(e.nativeScrollbar){const r=c.value;r&&(r.scrollTop=z,r.scrollLeft=T)}}),W(D,{collapsedRef:f,collapseModeRef:_(e,"collapseMode")});const{mergedClsPrefixRef:k,inlineThemeDisabled:w}=U(e),B=P("Layout","-layout-sider",Z,J,e,k);function L(r){var o,t;r.propertyName==="max-width"&&(f.value?(o=e.onAfterLeave)===null||o===void 0||o.call(e):(t=e.onAfterEnter)===null||t===void 0||t.call(e))}const M={scrollTo:O},R=p(()=>{const{common:{cubicBezierEaseInOut:r},self:o}=B.value,{siderToggleButtonColor:t,siderToggleButtonBorder:v,siderToggleBarColor:g,siderToggleBarColorHover:N}=o,i={"--n-bezier":r,"--n-toggle-button-color":t,"--n-toggle-button-border":v,"--n-toggle-bar-color":g,"--n-toggle-bar-color-hover":N};return e.inverted?(i["--n-color"]=o.siderColorInverted,i["--n-text-color"]=o.textColorInverted,i["--n-border-color"]=o.siderBorderColorInverted,i["--n-toggle-button-icon-color"]=o.siderToggleButtonIconColorInverted,i.__invertScrollbar=o.__invertScrollbar):(i["--n-color"]=o.siderColor,i["--n-text-color"]=o.textColor,i["--n-border-color"]=o.siderBorderColor,i["--n-toggle-button-icon-color"]=o.siderToggleButtonIconColor),i}),u=w?V("layout-sider",p(()=>e.inverted?"a":"b"),R,e):void 0;return Object.assign({scrollableElRef:c,scrollbarInstRef:b,mergedClsPrefix:k,mergedTheme:B,styleMaxWidth:$,mergedCollapsed:f,scrollContainerStyle:j,siderPlacement:I,handleNativeElScroll:Y,handleTransitionend:L,handleTriggerClick:E,inlineThemeDisabled:w,cssVars:R,themeClass:u==null?void 0:u.themeClass,onRender:u==null?void 0:u.onRender},M)},render(){var e;const{mergedClsPrefix:a,mergedCollapsed:c,showTrigger:b}=this;return(e=this.onRender)===null||e===void 0||e.call(this),s("aside",{class:[`${a}-layout-sider`,this.themeClass,`${a}-layout-sider--${this.position}-positioned`,`${a}-layout-sider--${this.siderPlacement}-placement`,this.bordered&&`${a}-layout-sider--bordered`,c&&`${a}-layout-sider--collapsed`,(!c||this.showCollapsedContent)&&`${a}-layout-sider--show-content`],onTransitionend:this.handleTransitionend,style:[this.inlineThemeDisabled?void 0:this.cssVars,{maxWidth:this.styleMaxWidth,width:x(this.width)}]},this.nativeScrollbar?s("div",{class:`${a}-layout-sider-scroll-container`,onScroll:this.handleNativeElScroll,style:[this.scrollContainerStyle,{overflow:"auto"},this.contentStyle],ref:"scrollableElRef"},this.$slots):s(H,Object.assign({},this.scrollbarProps,{onScroll:this.onScroll,ref:"scrollbarInstRef",style:this.scrollContainerStyle,contentStyle:this.contentStyle,theme:this.mergedTheme.peers.Scrollbar,themeOverrides:this.mergedTheme.peerOverrides.Scrollbar,builtinThemeOverrides:this.inverted&&this.cssVars.__invertScrollbar==="true"?{colorHover:"rgba(255, 255, 255, .4)",color:"rgba(255, 255, 255, .3)"}:void 0}),this.$slots),b?b==="bar"?s(oe,{clsPrefix:a,style:c?this.collapsedTriggerStyle:this.triggerStyle,onClick:this.handleTriggerClick}):s(ee,{clsPrefix:a,style:c?this.collapsedTriggerStyle:this.triggerStyle,onClick:this.handleTriggerClick}):null,this.bordered?s("div",{class:`${a}-layout-sider__border`}):null)}});export{se as N};
