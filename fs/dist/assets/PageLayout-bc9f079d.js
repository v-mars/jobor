import{d as R,h as v,c as $e,a as Me,b as Y,e as h,f as A,u as X,g as ge,i as I,j as xe,k as te,l as j,p as W,r as D,F as Ce,N as ke,m as F,n as oe,o as Le,q as ee,s as re,t as C,v as m,w as J,x as Fe,y as se,z as ue,A as K,B as $,C as U,D as q,E as N,G as O,H as Oe,I as Ee,J as _,K as be,L as Ie,M as Ke,O as Be,P as je,Q as Ve,R as De,S as H,T as Ue,U as qe,V as Ge,W as We,X as Ze,Y as Je,Z as ze,_ as Qe,$ as Xe}from"./index-969e67fe.js";import{u as ve,a as Ye,N as eo}from"./Icon-32b66fa9.js";import{p as oo,l as to,a as ro,N as no,b as me}from"./LayoutContent-b2bb953b.js";import{t as io,d as lo,N as ye,a as ao,c as co}from"./Dropdown-cd7deeac.js";import{N as so}from"./LayoutSider-2d8b9d65.js";import{_ as we}from"./_plugin-vue_export-helper-c27b6911.js";import{u as uo}from"./useRequest-c7c4a46d.js";import{_ as vo}from"./UserSelect.vue_vue_type_script_setup_true_lang-cec78c14.js";import{N as mo,a as ho}from"./FormItem-6aa25f6b.js";import{N as he}from"./Space-e91f3d3c.js";import{N as po}from"./Spin-2a22eadc.js";import"./FocusDetector-1e4512d4.js";import"./get-a0781824.js";const fo=R({name:"ChevronDownFilled",render(){return v("svg",{viewBox:"0 0 16 16",fill:"none",xmlns:"http://www.w3.org/2000/svg"},v("path",{d:"M3.20041 5.73966C3.48226 5.43613 3.95681 5.41856 4.26034 5.70041L8 9.22652L11.7397 5.70041C12.0432 5.41856 12.5177 5.43613 12.7996 5.73966C13.0815 6.0432 13.0639 6.51775 12.7603 6.7996L8.51034 10.7996C8.22258 11.0668 7.77743 11.0668 7.48967 10.7996L3.23966 6.7996C2.93613 6.51775 2.91856 6.0432 3.20041 5.73966Z",fill:"currentColor"}))}});function go(e,r,o,n){return{itemColorHoverInverted:"#0000",itemColorActiveInverted:r,itemColorActiveHoverInverted:r,itemColorActiveCollapsedInverted:r,itemTextColorInverted:e,itemTextColorHoverInverted:o,itemTextColorChildActiveInverted:o,itemTextColorChildActiveHoverInverted:o,itemTextColorActiveInverted:o,itemTextColorActiveHoverInverted:o,itemTextColorHorizontalInverted:e,itemTextColorHoverHorizontalInverted:o,itemTextColorChildActiveHorizontalInverted:o,itemTextColorChildActiveHoverHorizontalInverted:o,itemTextColorActiveHorizontalInverted:o,itemTextColorActiveHoverHorizontalInverted:o,itemIconColorInverted:e,itemIconColorHoverInverted:o,itemIconColorActiveInverted:o,itemIconColorActiveHoverInverted:o,itemIconColorChildActiveInverted:o,itemIconColorChildActiveHoverInverted:o,itemIconColorCollapsedInverted:e,itemIconColorHorizontalInverted:e,itemIconColorHoverHorizontalInverted:o,itemIconColorActiveHorizontalInverted:o,itemIconColorActiveHoverHorizontalInverted:o,itemIconColorChildActiveHorizontalInverted:o,itemIconColorChildActiveHoverHorizontalInverted:o,arrowColorInverted:e,arrowColorHoverInverted:o,arrowColorActiveInverted:o,arrowColorActiveHoverInverted:o,arrowColorChildActiveInverted:o,arrowColorChildActiveHoverInverted:o,groupTextColorInverted:n}}const xo=e=>{const{borderRadius:r,textColor3:o,primaryColor:n,textColor2:l,textColor1:a,fontSize:s,dividerColor:d,hoverColor:c,primaryColorHover:x}=e;return Object.assign({borderRadius:r,color:"#0000",groupTextColor:o,itemColorHover:c,itemColorActive:Y(n,{alpha:.1}),itemColorActiveHover:Y(n,{alpha:.1}),itemColorActiveCollapsed:Y(n,{alpha:.1}),itemTextColor:l,itemTextColorHover:l,itemTextColorActive:n,itemTextColorActiveHover:n,itemTextColorChildActive:n,itemTextColorChildActiveHover:n,itemTextColorHorizontal:l,itemTextColorHoverHorizontal:x,itemTextColorActiveHorizontal:n,itemTextColorActiveHoverHorizontal:n,itemTextColorChildActiveHorizontal:n,itemTextColorChildActiveHoverHorizontal:n,itemIconColor:a,itemIconColorHover:a,itemIconColorActive:n,itemIconColorActiveHover:n,itemIconColorChildActive:n,itemIconColorChildActiveHover:n,itemIconColorCollapsed:a,itemIconColorHorizontal:a,itemIconColorHoverHorizontal:x,itemIconColorActiveHorizontal:n,itemIconColorActiveHoverHorizontal:n,itemIconColorChildActiveHorizontal:n,itemIconColorChildActiveHoverHorizontal:n,itemHeight:"42px",arrowColor:l,arrowColorHover:l,arrowColorActive:n,arrowColorActiveHover:n,arrowColorChildActive:n,arrowColorChildActiveHover:n,colorInverted:"#0000",borderColorHorizontal:"#0000",fontSize:s,dividerColor:d},go("#BBB",n,"#FFF","#AAA"))},Co=$e({name:"Menu",common:Me,peers:{Tooltip:io,Dropdown:lo},self:xo}),bo=Co,Io=h("layout-header",`
 transition:
 color .3s var(--n-bezier),
 background-color .3s var(--n-bezier),
 box-shadow .3s var(--n-bezier),
 border-color .3s var(--n-bezier);
 box-sizing: border-box;
 width: 100%;
 background-color: var(--n-color);
 color: var(--n-text-color);
`,[A("absolute-positioned",`
 position: absolute;
 left: 0;
 right: 0;
 top: 0;
 `),A("bordered",`
 border-bottom: solid 1px var(--n-border-color);
 `)]),zo={position:oo,inverted:Boolean,bordered:{type:Boolean,default:!1}},yo=R({name:"LayoutHeader",props:Object.assign(Object.assign({},X.props),zo),setup(e){const{mergedClsPrefixRef:r,inlineThemeDisabled:o}=ge(e),n=X("Layout","-layout-header",Io,to,e,r),l=I(()=>{const{common:{cubicBezierEaseInOut:s},self:d}=n.value,c={"--n-bezier":s};return e.inverted?(c["--n-color"]=d.headerColorInverted,c["--n-text-color"]=d.textColorInverted,c["--n-border-color"]=d.headerBorderColorInverted):(c["--n-color"]=d.headerColor,c["--n-text-color"]=d.textColor,c["--n-border-color"]=d.headerBorderColor),c}),a=o?xe("layout-header",I(()=>e.inverted?"a":"b"),l,e):void 0;return{mergedClsPrefix:r,cssVars:o?void 0:l,themeClass:a==null?void 0:a.themeClass,onRender:a==null?void 0:a.onRender}},render(){var e;const{mergedClsPrefix:r}=this;return(e=this.onRender)===null||e===void 0||e.call(this),v("div",{class:[`${r}-layout-header`,this.themeClass,this.position&&`${r}-layout-header--${this.position}-positioned`,this.bordered&&`${r}-layout-header--bordered`],style:this.cssVars},this.$slots)}}),Z=te("n-menu"),ne=te("n-submenu"),ie=te("n-menu-item-group"),Q=8;function le(e){const r=j(Z),{props:o,mergedCollapsedRef:n}=r,l=j(ne,null),a=j(ie,null),s=I(()=>o.mode==="horizontal"),d=I(()=>s.value?o.dropdownPlacement:"tmNodes"in e?"right-start":"right"),c=I(()=>{var u;return Math.max((u=o.collapsedIconSize)!==null&&u!==void 0?u:o.iconSize,o.iconSize)}),x=I(()=>{var u;return!s.value&&e.root&&n.value&&(u=o.collapsedIconSize)!==null&&u!==void 0?u:o.iconSize}),w=I(()=>{if(s.value)return;const{collapsedWidth:u,indent:y,rootIndent:z}=o,{root:T,isGroup:P}=e,M=z===void 0?y:z;if(T)return n.value?u/2-c.value/2:M;if(a)return y/2+a.paddingLeftRef.value;if(l)return(P?y/2:y)+l.paddingLeftRef.value}),g=I(()=>{const{collapsedWidth:u,indent:y,rootIndent:z}=o,{value:T}=c,{root:P}=e;return s.value||!P||!n.value?Q:(z===void 0?y:z)+T+Q-(u+T)/2});return{dropdownPlacement:d,activeIconSize:x,maxIconSize:c,paddingLeft:w,iconMarginRight:g,NMenu:r,NSubmenu:l}}const ae={internalKey:{type:[String,Number],required:!0},root:Boolean,isGroup:Boolean,level:{type:Number,required:!0},title:[String,Function],extra:[String,Function]},_e=Object.assign(Object.assign({},ae),{tmNode:{type:Object,required:!0},tmNodes:{type:Array,required:!0}}),wo=R({name:"MenuOptionGroup",props:_e,setup(e){W(ne,null);const r=le(e);W(ie,{paddingLeftRef:r.paddingLeft});const{mergedClsPrefixRef:o,props:n}=j(Z);return function(){const{value:l}=o,a=r.paddingLeft.value,{nodeProps:s}=n,d=s==null?void 0:s(e.tmNode.rawNode);return v("div",{class:`${l}-menu-item-group`,role:"group"},v("div",Object.assign({},d,{class:[`${l}-menu-item-group-title`,d==null?void 0:d.class],style:[(d==null?void 0:d.style)||"",a!==void 0?`padding-left: ${a}px;`:""]}),D(e.title),e.extra?v(Ce,null," ",D(e.extra)):null),v("div",null,e.tmNodes.map(c=>ce(c,n))))}}}),Ae=R({name:"MenuOptionContent",props:{collapsed:Boolean,disabled:Boolean,title:[String,Function],icon:Function,extra:[String,Function],showArrow:Boolean,childActive:Boolean,hover:Boolean,paddingLeft:Number,selected:Boolean,maxIconSize:{type:Number,required:!0},activeIconSize:{type:Number,required:!0},iconMarginRight:{type:Number,required:!0},clsPrefix:{type:String,required:!0},onClick:Function,tmNode:{type:Object,required:!0}},setup(e){const{props:r}=j(Z);return{menuProps:r,style:I(()=>{const{paddingLeft:o}=e;return{paddingLeft:o&&`${o}px`}}),iconStyle:I(()=>{const{maxIconSize:o,activeIconSize:n,iconMarginRight:l}=e;return{width:`${o}px`,height:`${o}px`,fontSize:`${n}px`,marginRight:`${l}px`}})}},render(){const{clsPrefix:e,tmNode:r,menuProps:{renderIcon:o,renderLabel:n,renderExtra:l,expandIcon:a}}=this,s=o?o(r.rawNode):D(this.icon);return v("div",{onClick:d=>{var c;(c=this.onClick)===null||c===void 0||c.call(this,d)},role:"none",class:[`${e}-menu-item-content`,{[`${e}-menu-item-content--selected`]:this.selected,[`${e}-menu-item-content--collapsed`]:this.collapsed,[`${e}-menu-item-content--child-active`]:this.childActive,[`${e}-menu-item-content--disabled`]:this.disabled,[`${e}-menu-item-content--hover`]:this.hover}],style:this.style},s&&v("div",{class:`${e}-menu-item-content__icon`,style:this.iconStyle,role:"none"},[s]),v("div",{class:`${e}-menu-item-content-header`,role:"none"},n?n(r.rawNode):D(this.title),this.extra||l?v("span",{class:`${e}-menu-item-content-header__extra`}," ",l?l(r.rawNode):D(this.extra)):null),this.showArrow?v(ke,{ariaHidden:!0,class:`${e}-menu-item-content__arrow`,clsPrefix:e},{default:()=>a?a(r.rawNode):v(fo,null)}):null)}}),He=Object.assign(Object.assign({},ae),{rawNodes:{type:Array,default:()=>[]},tmNodes:{type:Array,default:()=>[]},tmNode:{type:Object,required:!0},disabled:{type:Boolean,default:!1},icon:Function,onClick:Function}),_o=R({name:"Submenu",props:He,setup(e){const r=le(e),{NMenu:o,NSubmenu:n}=r,{props:l,mergedCollapsedRef:a,mergedThemeRef:s}=o,d=I(()=>{const{disabled:u}=e;return n!=null&&n.mergedDisabledRef.value||l.disabled?!0:u}),c=F(!1);W(ne,{paddingLeftRef:r.paddingLeft,mergedDisabledRef:d}),W(ie,null);function x(){const{onClick:u}=e;u&&u()}function w(){d.value||(a.value||o.toggleExpand(e.internalKey),x())}function g(u){c.value=u}return{menuProps:l,mergedTheme:s,doSelect:o.doSelect,inverted:o.invertedRef,isHorizontal:o.isHorizontalRef,mergedClsPrefix:o.mergedClsPrefixRef,maxIconSize:r.maxIconSize,activeIconSize:r.activeIconSize,iconMarginRight:r.iconMarginRight,dropdownPlacement:r.dropdownPlacement,dropdownShow:c,paddingLeft:r.paddingLeft,mergedDisabled:d,mergedValue:o.mergedValueRef,childActive:oe(()=>o.activePathRef.value.includes(e.internalKey)),collapsed:I(()=>l.mode==="horizontal"?!1:a.value?!0:!o.mergedExpandedKeysRef.value.includes(e.internalKey)),dropdownEnabled:I(()=>!d.value&&(l.mode==="horizontal"||a.value)),handlePopoverShowChange:g,handleClick:w}},render(){var e;const{mergedClsPrefix:r,menuProps:{renderIcon:o,renderLabel:n}}=this,l=()=>{const{isHorizontal:s,paddingLeft:d,collapsed:c,mergedDisabled:x,maxIconSize:w,activeIconSize:g,title:u,childActive:y,icon:z,handleClick:T,menuProps:{nodeProps:P},dropdownShow:M,iconMarginRight:G,tmNode:k,mergedClsPrefix:E}=this,L=P==null?void 0:P(k.rawNode);return v("div",Object.assign({},L,{class:[`${E}-menu-item`,L==null?void 0:L.class],role:"menuitem"}),v(Ae,{tmNode:k,paddingLeft:d,collapsed:c,disabled:x,iconMarginRight:G,maxIconSize:w,activeIconSize:g,title:u,extra:this.extra,showArrow:!s,childActive:y,clsPrefix:E,icon:z,hover:M,onClick:T}))},a=()=>v(Le,null,{default:()=>{const{tmNodes:s,collapsed:d}=this;return d?null:v("div",{class:`${r}-submenu-children`,role:"menu"},s.map(c=>ce(c,this.menuProps)))}});return this.root?v(ye,Object.assign({size:"large",trigger:"hover"},(e=this.menuProps)===null||e===void 0?void 0:e.dropdownProps,{themeOverrides:this.mergedTheme.peerOverrides.Dropdown,theme:this.mergedTheme.peers.Dropdown,builtinThemeOverrides:{fontSizeLarge:"14px",optionIconSizeLarge:"18px"},value:this.mergedValue,disabled:!this.dropdownEnabled,placement:this.dropdownPlacement,keyField:this.menuProps.keyField,labelField:this.menuProps.labelField,childrenField:this.menuProps.childrenField,onUpdateShow:this.handlePopoverShowChange,options:this.rawNodes,onSelect:this.doSelect,inverted:this.inverted,renderIcon:o,renderLabel:n}),{default:()=>v("div",{class:`${r}-submenu`,role:"menuitem","aria-expanded":!this.collapsed},l(),this.isHorizontal?null:a())}):v("div",{class:`${r}-submenu`,role:"menuitem","aria-expanded":!this.collapsed},l(),a())}}),Se=Object.assign(Object.assign({},ae),{tmNode:{type:Object,required:!0},disabled:Boolean,icon:Function,onClick:Function}),Ao=R({name:"MenuOption",props:Se,setup(e){const r=le(e),{NSubmenu:o,NMenu:n}=r,{props:l,mergedClsPrefixRef:a,mergedCollapsedRef:s}=n,d=o?o.mergedDisabledRef:{value:!1},c=I(()=>d.value||e.disabled);function x(g){const{onClick:u}=e;u&&u(g)}function w(g){c.value||(n.doSelect(e.internalKey,e.tmNode.rawNode),x(g))}return{mergedClsPrefix:a,dropdownPlacement:r.dropdownPlacement,paddingLeft:r.paddingLeft,iconMarginRight:r.iconMarginRight,maxIconSize:r.maxIconSize,activeIconSize:r.activeIconSize,mergedTheme:n.mergedThemeRef,menuProps:l,dropdownEnabled:oe(()=>e.root&&s.value&&l.mode!=="horizontal"&&!c.value),selected:oe(()=>n.mergedValueRef.value===e.internalKey),mergedDisabled:c,handleClick:w}},render(){const{mergedClsPrefix:e,mergedTheme:r,tmNode:o,menuProps:{renderLabel:n,nodeProps:l}}=this,a=l==null?void 0:l(o.rawNode);return v("div",Object.assign({},a,{role:"menuitem",class:[`${e}-menu-item`,a==null?void 0:a.class]}),v(ao,{theme:r.peers.Tooltip,themeOverrides:r.peerOverrides.Tooltip,trigger:"hover",placement:this.dropdownPlacement,disabled:!this.dropdownEnabled||this.title===void 0,internalExtraClass:["menu-tooltip"]},{default:()=>n?n(o.rawNode):D(this.title),trigger:()=>v(Ae,{tmNode:o,clsPrefix:e,paddingLeft:this.paddingLeft,iconMarginRight:this.iconMarginRight,maxIconSize:this.maxIconSize,activeIconSize:this.activeIconSize,selected:this.selected,title:this.title,extra:this.extra,disabled:this.mergedDisabled,icon:this.icon,onClick:this.handleClick})}))}}),Ho=R({name:"MenuDivider",setup(){const e=j(Z),{mergedClsPrefixRef:r,isHorizontalRef:o}=e;return()=>o.value?null:v("div",{class:`${r.value}-menu-divider`})}}),So=re(_e),Ro=re(Se),No=re(He);function Re(e){return e.type==="divider"||e.type==="render"}function Po(e){return e.type==="divider"}function ce(e,r){const{rawNode:o}=e,{show:n}=o;if(n===!1)return null;if(Re(o))return Po(o)?v(Ho,Object.assign({key:e.key},o.props)):null;const{labelField:l}=r,{key:a,level:s,isGroup:d}=e,c=Object.assign(Object.assign({},o),{title:o.title||o[l],extra:o.titleExtra||o.extra,key:a,internalKey:a,level:s,root:s===0,isGroup:d});return e.children?e.isGroup?v(wo,ee(c,So,{tmNode:e,tmNodes:e.children,key:a})):v(_o,ee(c,No,{key:a,rawNodes:o[r.childrenField],tmNodes:e.children,tmNode:e})):v(Ao,ee(c,Ro,{key:a,tmNode:e}))}const pe=[C("&::before","background-color: var(--n-item-color-hover);"),m("arrow",`
 color: var(--n-arrow-color-hover);
 `),m("icon",`
 color: var(--n-item-icon-color-hover);
 `),h("menu-item-content-header",`
 color: var(--n-item-text-color-hover);
 `,[C("a",`
 color: var(--n-item-text-color-hover);
 `),m("extra",`
 color: var(--n-item-text-color-hover);
 `)])],fe=[m("icon",`
 color: var(--n-item-icon-color-hover-horizontal);
 `),h("menu-item-content-header",`
 color: var(--n-item-text-color-hover-horizontal);
 `,[C("a",`
 color: var(--n-item-text-color-hover-horizontal);
 `),m("extra",`
 color: var(--n-item-text-color-hover-horizontal);
 `)])],To=C([h("menu",`
 background-color: var(--n-color);
 color: var(--n-item-text-color);
 overflow: hidden;
 transition: background-color .3s var(--n-bezier);
 box-sizing: border-box;
 font-size: var(--n-font-size);
 padding-bottom: 6px;
 `,[A("horizontal",`
 display: inline-flex;
 padding-bottom: 0;
 `,[h("submenu","margin: 0;"),h("menu-item","margin: 0;"),h("menu-item-content",`
 padding: 0 20px;
 border-bottom: 2px solid #0000;
 `,[C("&::before","display: none;"),A("selected","border-bottom: 2px solid var(--n-border-color-horizontal)")]),h("menu-item-content",[A("selected",[m("icon","color: var(--n-item-icon-color-active-horizontal);"),h("menu-item-content-header",`
 color: var(--n-item-text-color-active-horizontal);
 `,[C("a","color: var(--n-item-text-color-active-horizontal);"),m("extra","color: var(--n-item-text-color-active-horizontal);")])]),A("child-active",`
 border-bottom: 2px solid var(--n-border-color-horizontal);
 `,[h("menu-item-content-header",`
 color: var(--n-item-text-color-child-active-horizontal);
 `,[C("a",`
 color: var(--n-item-text-color-child-active-horizontal);
 `),m("extra",`
 color: var(--n-item-text-color-child-active-horizontal);
 `)]),m("icon",`
 color: var(--n-item-icon-color-child-active-horizontal);
 `)]),J("disabled",[J("selected, child-active",[C("&:focus-within",fe)]),A("selected",[B(null,[m("icon","color: var(--n-item-icon-color-active-hover-horizontal);"),h("menu-item-content-header",`
 color: var(--n-item-text-color-active-hover-horizontal);
 `,[C("a","color: var(--n-item-text-color-active-hover-horizontal);"),m("extra","color: var(--n-item-text-color-active-hover-horizontal);")])])]),A("child-active",[B(null,[m("icon","color: var(--n-item-icon-color-child-active-hover-horizontal);"),h("menu-item-content-header",`
 color: var(--n-item-text-color-child-active-hover-horizontal);
 `,[C("a","color: var(--n-item-text-color-child-active-hover-horizontal);"),m("extra","color: var(--n-item-text-color-child-active-hover-horizontal);")])])]),B("border-bottom: 2px solid var(--n-border-color-horizontal);",fe)]),h("menu-item-content-header",[C("a","color: var(--n-item-text-color-horizontal);")])])]),A("collapsed",[h("menu-item-content",[A("selected",[C("&::before",`
 background-color: var(--n-item-color-active-collapsed) !important;
 `)]),h("menu-item-content-header","opacity: 0;"),m("arrow","opacity: 0;"),m("icon","color: var(--n-item-icon-color-collapsed);")])]),h("menu-item",`
 height: var(--n-item-height);
 margin-top: 6px;
 position: relative;
 `),h("menu-item-content",`
 box-sizing: border-box;
 line-height: 1.75;
 height: 100%;
 display: grid;
 grid-template-areas: "icon content arrow";
 grid-template-columns: auto 1fr auto;
 align-items: center;
 cursor: pointer;
 position: relative;
 padding-right: 18px;
 transition:
 background-color .3s var(--n-bezier),
 padding-left .3s var(--n-bezier),
 border-color .3s var(--n-bezier);
 `,[C("> *","z-index: 1;"),C("&::before",`
 z-index: auto;
 content: "";
 background-color: #0000;
 position: absolute;
 left: 8px;
 right: 8px;
 top: 0;
 bottom: 0;
 pointer-events: none;
 border-radius: var(--n-border-radius);
 transition: background-color .3s var(--n-bezier);
 `),A("disabled",`
 opacity: .45;
 cursor: not-allowed;
 `),A("collapsed",[m("arrow","transform: rotate(0);")]),A("selected",[C("&::before","background-color: var(--n-item-color-active);"),m("arrow","color: var(--n-arrow-color-active);"),m("icon","color: var(--n-item-icon-color-active);"),h("menu-item-content-header",`
 color: var(--n-item-text-color-active);
 `,[C("a","color: var(--n-item-text-color-active);"),m("extra","color: var(--n-item-text-color-active);")])]),A("child-active",[h("menu-item-content-header",`
 color: var(--n-item-text-color-child-active);
 `,[C("a",`
 color: var(--n-item-text-color-child-active);
 `),m("extra",`
 color: var(--n-item-text-color-child-active);
 `)]),m("arrow",`
 color: var(--n-arrow-color-child-active);
 `),m("icon",`
 color: var(--n-item-icon-color-child-active);
 `)]),J("disabled",[J("selected, child-active",[C("&:focus-within",pe)]),A("selected",[B(null,[m("arrow","color: var(--n-arrow-color-active-hover);"),m("icon","color: var(--n-item-icon-color-active-hover);"),h("menu-item-content-header",`
 color: var(--n-item-text-color-active-hover);
 `,[C("a","color: var(--n-item-text-color-active-hover);"),m("extra","color: var(--n-item-text-color-active-hover);")])])]),A("child-active",[B(null,[m("arrow","color: var(--n-arrow-color-child-active-hover);"),m("icon","color: var(--n-item-icon-color-child-active-hover);"),h("menu-item-content-header",`
 color: var(--n-item-text-color-child-active-hover);
 `,[C("a","color: var(--n-item-text-color-child-active-hover);"),m("extra","color: var(--n-item-text-color-child-active-hover);")])])]),A("selected",[B(null,[C("&::before","background-color: var(--n-item-color-active-hover);")])]),B(null,pe)]),m("icon",`
 grid-area: icon;
 color: var(--n-item-icon-color);
 transition:
 color .3s var(--n-bezier),
 font-size .3s var(--n-bezier),
 margin-right .3s var(--n-bezier);
 box-sizing: content-box;
 display: inline-flex;
 align-items: center;
 justify-content: center;
 `),m("arrow",`
 grid-area: arrow;
 font-size: 16px;
 color: var(--n-arrow-color);
 transform: rotate(180deg);
 opacity: 1;
 transition:
 color .3s var(--n-bezier),
 transform 0.2s var(--n-bezier),
 opacity 0.2s var(--n-bezier);
 `),h("menu-item-content-header",`
 grid-area: content;
 transition:
 color .3s var(--n-bezier),
 opacity .3s var(--n-bezier);
 opacity: 1;
 white-space: nowrap;
 overflow: hidden;
 text-overflow: ellipsis;
 color: var(--n-item-text-color);
 `,[C("a",`
 outline: none;
 text-decoration: none;
 transition: color .3s var(--n-bezier);
 color: var(--n-item-text-color);
 `,[C("&::before",`
 content: "";
 position: absolute;
 left: 0;
 right: 0;
 top: 0;
 bottom: 0;
 `)]),m("extra",`
 font-size: .93em;
 color: var(--n-group-text-color);
 transition: color .3s var(--n-bezier);
 `)])]),h("submenu",`
 cursor: pointer;
 position: relative;
 margin-top: 6px;
 `,[h("menu-item-content",`
 height: var(--n-item-height);
 `),h("submenu-children",`
 overflow: hidden;
 padding: 0;
 `,[Fe({duration:".2s"})])]),h("menu-item-group",[h("menu-item-group-title",`
 margin-top: 6px;
 color: var(--n-group-text-color);
 cursor: default;
 font-size: .93em;
 height: 36px;
 display: flex;
 align-items: center;
 transition:
 padding-left .3s var(--n-bezier),
 color .3s var(--n-bezier);
 `)])]),h("menu-tooltip",[C("a",`
 color: inherit;
 text-decoration: none;
 `)]),h("menu-divider",`
 transition: background-color .3s var(--n-bezier);
 background-color: var(--n-divider-color);
 height: 1px;
 margin: 6px 18px;
 `)]);function B(e,r){return[A("hover",e,r),C("&:hover",e,r)]}const $o=Object.assign(Object.assign({},X.props),{options:{type:Array,default:()=>[]},collapsed:{type:Boolean,default:void 0},collapsedWidth:{type:Number,default:48},iconSize:{type:Number,default:20},collapsedIconSize:{type:Number,default:24},rootIndent:Number,indent:{type:Number,default:32},labelField:{type:String,default:"label"},keyField:{type:String,default:"key"},childrenField:{type:String,default:"children"},disabledField:{type:String,default:"disabled"},defaultExpandAll:Boolean,defaultExpandedKeys:Array,expandedKeys:Array,value:[String,Number],defaultValue:{type:[String,Number],default:null},mode:{type:String,default:"vertical"},watchProps:{type:Array,default:void 0},disabled:Boolean,show:{type:Boolean,defalut:!0},inverted:Boolean,"onUpdate:expandedKeys":[Function,Array],onUpdateExpandedKeys:[Function,Array],onUpdateValue:[Function,Array],"onUpdate:value":[Function,Array],expandIcon:Function,renderIcon:Function,renderLabel:Function,renderExtra:Function,dropdownProps:Object,accordion:Boolean,nodeProps:Function,items:Array,onOpenNamesChange:[Function,Array],onSelect:[Function,Array],onExpandedNamesChange:[Function,Array],expandedNames:Array,defaultExpandedNames:Array,dropdownPlacement:{type:String,default:"bottom"}}),Mo=R({name:"Menu",props:$o,setup(e){const{mergedClsPrefixRef:r,inlineThemeDisabled:o}=ge(e),n=X("Menu","-menu",To,bo,e,r),l=j(ro,null),a=I(()=>{var p;const{collapsed:b}=e;if(b!==void 0)return b;if(l){const{collapseModeRef:t,collapsedRef:f}=l;if(t.value==="width")return(p=f.value)!==null&&p!==void 0?p:!1}return!1}),s=I(()=>{const{keyField:p,childrenField:b,disabledField:t}=e;return co(e.items||e.options,{getIgnored(f){return Re(f)},getChildren(f){return f[b]},getDisabled(f){return f[t]},getKey(f){var S;return(S=f[p])!==null&&S!==void 0?S:f.name}})}),d=I(()=>new Set(s.value.treeNodes.map(p=>p.key))),{watchProps:c}=e,x=F(null);c!=null&&c.includes("defaultValue")?se(()=>{x.value=e.defaultValue}):x.value=e.defaultValue;const w=ue(e,"value"),g=ve(w,x),u=F([]),y=()=>{u.value=e.defaultExpandAll?s.value.getNonLeafKeys():e.defaultExpandedNames||e.defaultExpandedKeys||s.value.getPath(g.value,{includeSelf:!1}).keyPath};c!=null&&c.includes("defaultExpandedKeys")?se(y):y();const z=Ye(e,["expandedNames","expandedKeys"]),T=ve(z,u),P=I(()=>s.value.treeNodes),M=I(()=>s.value.getPath(g.value).keyPath);W(Z,{props:e,mergedCollapsedRef:a,mergedThemeRef:n,mergedValueRef:g,mergedExpandedKeysRef:T,activePathRef:M,mergedClsPrefixRef:r,isHorizontalRef:I(()=>e.mode==="horizontal"),invertedRef:ue(e,"inverted"),doSelect:G,toggleExpand:E});function G(p,b){const{"onUpdate:value":t,onUpdateValue:f,onSelect:S}=e;f&&K(f,p,b),t&&K(t,p,b),S&&K(S,p,b),x.value=p}function k(p){const{"onUpdate:expandedKeys":b,onUpdateExpandedKeys:t,onExpandedNamesChange:f,onOpenNamesChange:S}=e;b&&K(b,p),t&&K(t,p),f&&K(f,p),S&&K(S,p),u.value=p}function E(p){const b=Array.from(T.value),t=b.findIndex(f=>f===p);if(~t)b.splice(t,1);else{if(e.accordion&&d.value.has(p)){const f=b.findIndex(S=>d.value.has(S));f>-1&&b.splice(f,1)}b.push(p)}k(b)}const L=p=>{const b=s.value.getPath(p??g.value,{includeSelf:!1}).keyPath;if(!b.length)return;const t=Array.from(T.value),f=new Set([...t,...b]);e.accordion&&d.value.forEach(S=>{f.has(S)&&!b.includes(S)&&f.delete(S)}),k(Array.from(f))},de=I(()=>{const{inverted:p}=e,{common:{cubicBezierEaseInOut:b},self:t}=n.value,{borderRadius:f,borderColorHorizontal:S,fontSize:Ne,itemHeight:Pe,dividerColor:Te}=t,i={"--n-divider-color":Te,"--n-bezier":b,"--n-font-size":Ne,"--n-border-color-horizontal":S,"--n-border-radius":f,"--n-item-height":Pe};return p?(i["--n-group-text-color"]=t.groupTextColorInverted,i["--n-color"]=t.colorInverted,i["--n-item-text-color"]=t.itemTextColorInverted,i["--n-item-text-color-hover"]=t.itemTextColorHoverInverted,i["--n-item-text-color-active"]=t.itemTextColorActiveInverted,i["--n-item-text-color-child-active"]=t.itemTextColorChildActiveInverted,i["--n-item-text-color-child-active-hover"]=t.itemTextColorChildActiveInverted,i["--n-item-text-color-active-hover"]=t.itemTextColorActiveHoverInverted,i["--n-item-icon-color"]=t.itemIconColorInverted,i["--n-item-icon-color-hover"]=t.itemIconColorHoverInverted,i["--n-item-icon-color-active"]=t.itemIconColorActiveInverted,i["--n-item-icon-color-active-hover"]=t.itemIconColorActiveHoverInverted,i["--n-item-icon-color-child-active"]=t.itemIconColorChildActiveInverted,i["--n-item-icon-color-child-active-hover"]=t.itemIconColorChildActiveHoverInverted,i["--n-item-icon-color-collapsed"]=t.itemIconColorCollapsedInverted,i["--n-item-text-color-horizontal"]=t.itemTextColorHorizontalInverted,i["--n-item-text-color-hover-horizontal"]=t.itemTextColorHoverHorizontalInverted,i["--n-item-text-color-active-horizontal"]=t.itemTextColorActiveHorizontalInverted,i["--n-item-text-color-child-active-horizontal"]=t.itemTextColorChildActiveHorizontalInverted,i["--n-item-text-color-child-active-hover-horizontal"]=t.itemTextColorChildActiveHoverHorizontalInverted,i["--n-item-text-color-active-hover-horizontal"]=t.itemTextColorActiveHoverHorizontalInverted,i["--n-item-icon-color-horizontal"]=t.itemIconColorHorizontalInverted,i["--n-item-icon-color-hover-horizontal"]=t.itemIconColorHoverHorizontalInverted,i["--n-item-icon-color-active-horizontal"]=t.itemIconColorActiveHorizontalInverted,i["--n-item-icon-color-active-hover-horizontal"]=t.itemIconColorActiveHoverHorizontalInverted,i["--n-item-icon-color-child-active-horizontal"]=t.itemIconColorChildActiveHorizontalInverted,i["--n-item-icon-color-child-active-hover-horizontal"]=t.itemIconColorChildActiveHoverHorizontalInverted,i["--n-arrow-color"]=t.arrowColorInverted,i["--n-arrow-color-hover"]=t.arrowColorHoverInverted,i["--n-arrow-color-active"]=t.arrowColorActiveInverted,i["--n-arrow-color-active-hover"]=t.arrowColorActiveHoverInverted,i["--n-arrow-color-child-active"]=t.arrowColorChildActiveInverted,i["--n-arrow-color-child-active-hover"]=t.arrowColorChildActiveHoverInverted,i["--n-item-color-hover"]=t.itemColorHoverInverted,i["--n-item-color-active"]=t.itemColorActiveInverted,i["--n-item-color-active-hover"]=t.itemColorActiveHoverInverted,i["--n-item-color-active-collapsed"]=t.itemColorActiveCollapsedInverted):(i["--n-group-text-color"]=t.groupTextColor,i["--n-color"]=t.color,i["--n-item-text-color"]=t.itemTextColor,i["--n-item-text-color-hover"]=t.itemTextColorHover,i["--n-item-text-color-active"]=t.itemTextColorActive,i["--n-item-text-color-child-active"]=t.itemTextColorChildActive,i["--n-item-text-color-child-active-hover"]=t.itemTextColorChildActiveHover,i["--n-item-text-color-active-hover"]=t.itemTextColorActiveHover,i["--n-item-icon-color"]=t.itemIconColor,i["--n-item-icon-color-hover"]=t.itemIconColorHover,i["--n-item-icon-color-active"]=t.itemIconColorActive,i["--n-item-icon-color-active-hover"]=t.itemIconColorActiveHover,i["--n-item-icon-color-child-active"]=t.itemIconColorChildActive,i["--n-item-icon-color-child-active-hover"]=t.itemIconColorChildActiveHover,i["--n-item-icon-color-collapsed"]=t.itemIconColorCollapsed,i["--n-item-text-color-horizontal"]=t.itemTextColorHorizontal,i["--n-item-text-color-hover-horizontal"]=t.itemTextColorHoverHorizontal,i["--n-item-text-color-active-horizontal"]=t.itemTextColorActiveHorizontal,i["--n-item-text-color-child-active-horizontal"]=t.itemTextColorChildActiveHorizontal,i["--n-item-text-color-child-active-hover-horizontal"]=t.itemTextColorChildActiveHoverHorizontal,i["--n-item-text-color-active-hover-horizontal"]=t.itemTextColorActiveHoverHorizontal,i["--n-item-icon-color-horizontal"]=t.itemIconColorHorizontal,i["--n-item-icon-color-hover-horizontal"]=t.itemIconColorHoverHorizontal,i["--n-item-icon-color-active-horizontal"]=t.itemIconColorActiveHorizontal,i["--n-item-icon-color-active-hover-horizontal"]=t.itemIconColorActiveHoverHorizontal,i["--n-item-icon-color-child-active-horizontal"]=t.itemIconColorChildActiveHorizontal,i["--n-item-icon-color-child-active-hover-horizontal"]=t.itemIconColorChildActiveHoverHorizontal,i["--n-arrow-color"]=t.arrowColor,i["--n-arrow-color-hover"]=t.arrowColorHover,i["--n-arrow-color-active"]=t.arrowColorActive,i["--n-arrow-color-active-hover"]=t.arrowColorActiveHover,i["--n-arrow-color-child-active"]=t.arrowColorChildActive,i["--n-arrow-color-child-active-hover"]=t.arrowColorChildActiveHover,i["--n-item-color-hover"]=t.itemColorHover,i["--n-item-color-active"]=t.itemColorActive,i["--n-item-color-active-hover"]=t.itemColorActiveHover,i["--n-item-color-active-collapsed"]=t.itemColorActiveCollapsed),i}),V=o?xe("menu",I(()=>e.inverted?"a":"b"),de,e):void 0;return{mergedClsPrefix:r,controlledExpandedKeys:z,uncontrolledExpanededKeys:u,mergedExpandedKeys:T,uncontrolledValue:x,mergedValue:g,activePath:M,tmNodes:P,mergedTheme:n,mergedCollapsed:a,cssVars:o?void 0:de,themeClass:V==null?void 0:V.themeClass,onRender:V==null?void 0:V.onRender,showOption:L}},render(){const{mergedClsPrefix:e,mode:r,themeClass:o,onRender:n}=this;return n==null||n(),v("div",{role:r==="horizontal"?"menubar":"menu",class:[`${e}-menu`,o,`${e}-menu--${r}`,this.mergedCollapsed&&`${e}-menu--collapsed`],style:this.cssVars},this.tmNodes.map(l=>ce(l,this.$props)))}}),ko="/assets/jobor_ico-0674759c.png",Lo="/assets/jobor_logo-6a5466fd.png",Fo={class:"logo-container",flex:"","items-center":"","justify-center":"","h-12":""},Oo={key:0,src:ko,class:"h-80%",alt:""},Eo={key:1,src:Lo,class:"h-80%",alt:""},Ko=R({__name:"BaseLogo",props:{isMini:{type:Boolean}},setup(e){const r=e;return(o,n)=>($(),U("div",Fo,[r.isMini?($(),U("img",Oo)):($(),U("img",Eo))]))}}),Bo={class:"icon","aria-hidden":"true"},jo=["xlink:href"],Vo=R({__name:"SvgIcon",props:{path:{default:""}},setup(e){const r=e;return(o,n)=>($(),q(_(eo),Oe(Ee(o.$attrs)),{default:N(()=>[($(),U("svg",Bo,[O("use",{"xlink:href":r.path},null,8,jo)]))]),_:1},16))}}),Do=R({__name:"MenuGroup",setup(e){const r=Be(),o=je(),n=be(),{roles:l}=Ie(n),a=F(String(r.name));Ke(()=>r.fullPath,()=>{a.value=String(r.name)});const s=I(()=>o.getRoutes().find(x=>x.name==="root"));function d(x){const w=v(Vo,{path:`#icon-${x}`});return()=>v(w,{})}const c=I(()=>{const x=[...s.value.children];function w(g,u){return g.map(y=>{var E;const{meta:z}=y;y.children&&(y.children=w(y.children));const P=Ve(l.value,z==null?void 0:z.role)&&!(z!=null&&z.hideMenu),M=!!(z!=null&&z.icon),G=z==null?void 0:z.icon;let k;return P&&(y.children?k=(E=y.meta)==null?void 0:E.title:k=v(De,{to:y.path},{default:()=>{var L;return((L=y.meta)==null?void 0:L.title)??"--"}})),{...y,label:()=>P&&k,show:P,key:y.name,icon:M?d(G):null}})}return w(x)});return(x,w)=>($(),q(_(Mo),{mode:"vertical",value:a.value,indent:20,accordion:!1,"collapsed-width":64,"collapsed-icon-size":22,options:_(c)},null,8,["value","options"]))}}),Uo=R({__name:"BaseSider",setup(e){const r=F(!1);function o(n){r.value=n}return(n,l)=>($(),q(_(so),{width:"200",bordered:"","native-scrollbar":!1,"show-trigger":"","collapse-mode":"width","collapsed-width":64,"onUpdate:collapsed":o},{default:N(()=>[H(Ko,{"is-mini":r.value},null,8,["is-mini"]),H(Do)]),_:1}))}}),qo={class:"avator"},Go=["alt"],Wo={class:"avator-container"},Zo=R({__name:"AutoTextAvatar",props:{text:{default:""}},setup(e){const r=e,{text:o}=Ue(r);return(n,l)=>($(),U("div",qo,[O("div",{class:"avator-inner",alt:_(o)},[O("div",Wo,[O("span",null,qe(_(o)),1)])],8,Go)]))}});const Jo=we(Zo,[["__scopeId","data-v-e59811c3"]]),Qo={flex:"","justify-end":""},Xo=R({__name:"UserSwitchModal",emits:["refresh"],setup(e,{emit:r}){const o=F(!1),{run:n,loading:l}=uo(ze.userSwitchById),a=F(null),s=F({userId:void 0}),d={userId:{required:!0,message:"请选择用户",trigger:"blur",type:"number"}};function c(){o.value=!0}async function x(){var w;(w=a.value)==null||w.validate(async g=>{if(g)return;const{result:u}=await n(Number(s.value.userId));u&&(o.value=!1,r("refresh"))})}return(w,g)=>($(),U(Ce,null,[O("span",{onClick:c},[Ge(w.$slots,"default")]),H(_(Je),{show:o.value,"onUpdate:show":g[2]||(g[2]=u=>o.value=u),title:"切换用户",preset:"card","auto-focus":!1,"w-100":""},{default:N(()=>[H(_(mo),{ref_key:"formRef",ref:a,model:s.value,rules:d},{default:N(()=>[H(_(ho),{path:"userId","show-label":!1},{default:N(()=>[H(vo,{value:s.value.userId,"onUpdate:value":g[0]||(g[0]=u=>s.value.userId=u)},null,8,["value"])]),_:1}),O("div",Qo,[H(_(We),{type:"primary",loading:_(l),onClick:g[1]||(g[1]=u=>x())},{default:N(()=>[Ze(" 确定 ")]),_:1},8,["loading"])])]),_:1},8,["model"])]),_:1},8,["show"])],64))}}),Yo={"h-12":"","w-auto":"",flex:"","px-5":"","items-center":"","justify-between":""},et=O("span",null,null,-1),ot=R({__name:"BaseHeader",setup(e){const r=be(),{nickname:o,isAdmin:n}=Ie(r);function l(){window.location.reload()}async function a(){const{result:d}=await ze.toLogout();l()}const s=[{label:()=>v(Xo,{onRefresh:l},{default:()=>v("div",{},"切换用户")}),show:n.value},{label:"退出登录",props:{onClick:a}}];return(d,c)=>($(),q(_(yo),{bordered:""},{default:N(()=>[O("div",Yo,[H(_(he),{align:"center",justify:"center"},{default:N(()=>[et]),_:1}),H(_(he),{align:"center",justify:"center"},{default:N(()=>[H(_(ye),{"display-directive":"show",options:s},{default:N(()=>[H(Jo,{text:_(o)},null,8,["text"])]),_:1})]),_:1})])]),_:1}))}}),tt=R({__name:"AppMain",setup(e){return(r,o)=>{const n=Xe("RouterView");return $(),q(Qe,null,{fallback:N(()=>[H(_(po),{size:"large"})]),default:N(()=>[H(n)]),_:1})}}});const rt=we(tt,[["__scopeId","data-v-91d1c4f9"]]),gt=R({__name:"PageLayout",setup(e){return(r,o)=>($(),q(_(me),{"has-sider":"",position:"absolute"},{default:N(()=>[H(Uo),H(_(me),{"native-scrollbar":!0,style:{overflow:"initial"}},{default:N(()=>[H(ot),H(_(no),{position:"absolute",style:{"margin-top":"50px"},"content-style":"padding:16px"},{default:N(()=>[H(rt)]),_:1})]),_:1})]),_:1}))}});export{gt as default};
