import{d as S,h as u,c as Me,a as Ce,b as Q,e as m,f as _,u as q,g as te,i as b,j as ne,k as re,l as j,p as X,r as D,F as be,N as Le,m as O,n as oe,o as Oe,q as ee,s as ie,t as C,v as h,w as Z,x as Fe,y as ve,z as me,A as he,B,C as Ee,D as Be,T as Ke,E as je,G as Ve,H as k,I as U,J as G,K as N,L as F,M as De,O as Ue,P as H,Q as ze,R as ye,S as qe,U as Ge,V as We,W as Xe,X as Ye,Y as A,Z as Ze,_ as Je,$ as Qe,a0 as eo,a1 as oo,a2 as to,a3 as no,a4 as ro,a5 as io,a6 as Ie,a7 as lo,a8 as ao}from "./index-a4639770.js";import{t as co,d as so,N as we,a as uo,u as _e,c as vo,b as mo}from "./Dropdown-c4cd24d0.js";import{p as ho,l as po,a as fo,N as go,b as xo,c as pe}from "./LayoutSider-33f9a29a.js";import{_ as He}from "./_plugin-vue_export-helper-c27b6911.js";import{_ as Co}from "./UserSelect.vue_vue_type_script_setup_true_lang-14681a6f.js";import{N as fe}from "./Space-1c3e7888.js";import"./FocusDetector-f0e5f0c6.js";const bo=S({name:"ChevronDownFilled",render(){return u("svg",{viewBox:"0 0 16 16",fill:"none",xmlns:"http://www.w3.org/2000/svg"},u("path",{d:"M3.20041 5.73966C3.48226 5.43613 3.95681 5.41856 4.26034 5.70041L8 9.22652L11.7397 5.70041C12.0432 5.41856 12.5177 5.43613 12.7996 5.73966C13.0815 6.0432 13.0639 6.51775 12.7603 6.7996L8.51034 10.7996C8.22258 11.0668 7.77743 11.0668 7.48967 10.7996L3.23966 6.7996C2.93613 6.51775 2.91856 6.0432 3.20041 5.73966Z",fill:"currentColor"}))}});function zo(e, n, o, t){return{itemColorHoverInverted:"#0000",itemColorActiveInverted:n,itemColorActiveHoverInverted:n,itemColorActiveCollapsedInverted:n,itemTextColorInverted:e,itemTextColorHoverInverted:o,itemTextColorChildActiveInverted:o,itemTextColorChildActiveHoverInverted:o,itemTextColorActiveInverted:o,itemTextColorActiveHoverInverted:o,itemTextColorHorizontalInverted:e,itemTextColorHoverHorizontalInverted:o,itemTextColorChildActiveHorizontalInverted:o,itemTextColorChildActiveHoverHorizontalInverted:o,itemTextColorActiveHorizontalInverted:o,itemTextColorActiveHoverHorizontalInverted:o,itemIconColorInverted:e,itemIconColorHoverInverted:o,itemIconColorActiveInverted:o,itemIconColorActiveHoverInverted:o,itemIconColorChildActiveInverted:o,itemIconColorChildActiveHoverInverted:o,itemIconColorCollapsedInverted:e,itemIconColorHorizontalInverted:e,itemIconColorHoverHorizontalInverted:o,itemIconColorActiveHorizontalInverted:o,itemIconColorActiveHoverHorizontalInverted:o,itemIconColorChildActiveHorizontalInverted:o,itemIconColorChildActiveHoverHorizontalInverted:o,arrowColorInverted:e,arrowColorHoverInverted:o,arrowColorActiveInverted:o,arrowColorActiveHoverInverted:o,arrowColorChildActiveInverted:o,arrowColorChildActiveHoverInverted:o,groupTextColorInverted:t}}const yo= e=>{const{borderRadius:n,textColor3:o,primaryColor:t,textColor2:a,textColor1:l,fontSize:c,dividerColor:s,hoverColor:d,primaryColorHover:g}=e;return Object.assign({borderRadius:n,color:"#0000",groupTextColor:o,itemColorHover:d,itemColorActive:Q(t,{alpha:.1}),itemColorActiveHover:Q(t,{alpha:.1}),itemColorActiveCollapsed:Q(t,{alpha:.1}),itemTextColor:a,itemTextColorHover:a,itemTextColorActive:t,itemTextColorActiveHover:t,itemTextColorChildActive:t,itemTextColorChildActiveHover:t,itemTextColorHorizontal:a,itemTextColorHoverHorizontal:g,itemTextColorActiveHorizontal:t,itemTextColorActiveHoverHorizontal:t,itemTextColorChildActiveHorizontal:t,itemTextColorChildActiveHoverHorizontal:t,itemIconColor:l,itemIconColorHover:l,itemIconColorActive:t,itemIconColorActiveHover:t,itemIconColorChildActive:t,itemIconColorChildActiveHover:t,itemIconColorCollapsed:l,itemIconColorHorizontal:l,itemIconColorHoverHorizontal:g,itemIconColorActiveHorizontal:t,itemIconColorActiveHoverHorizontal:t,itemIconColorChildActiveHorizontal:t,itemIconColorChildActiveHoverHorizontal:t,itemHeight:"42px",arrowColor:a,arrowColorHover:a,arrowColorActive:t,arrowColorActiveHover:t,arrowColorChildActive:t,arrowColorChildActiveHover:t,colorInverted:"#0000",borderColorHorizontal:"#0000",fontSize:c,dividerColor:s},zo("#BBB",t,"#FFF","#AAA"))},Io=Me({name:"Menu",common:Ce,peers:{Tooltip:co,Dropdown:so},self:yo}),wo=Io,_o= e=>{const{opacityDisabled:n,heightTiny:o,heightSmall:t,heightMedium:a,heightLarge:l,heightHuge:c,primaryColor:s,fontSize:d}=e;return{fontSize:d,textColor:s,sizeTiny:o,sizeSmall:t,sizeMedium:a,sizeLarge:l,sizeHuge:c,color:s,opacitySpinning:n}},Ho={name:"Spin",common:Ce,self:_o},Ao=Ho,So=m("layout-header",`
 transition:
 color .3s var(--n-bezier),
 background-color .3s var(--n-bezier),
 box-shadow .3s var(--n-bezier),
 border-color .3s var(--n-bezier);
 box-sizing: border-box;
 width: 100%;
 background-color: var(--n-color);
 color: var(--n-text-color);
`,[_("absolute-positioned",`
 position: absolute;
 left: 0;
 right: 0;
 top: 0;
 `),_("bordered",`
 border-bottom: solid 1px var(--n-border-color);
 `)]),Ro={position:ho,inverted:Boolean,bordered:{type:Boolean,default:!1}},No=S({name:"LayoutHeader",props:Object.assign(Object.assign({},q.props),Ro),setup(e){const{mergedClsPrefixRef:n,inlineThemeDisabled:o}=te(e),t=q("Layout","-layout-header",So,po,e,n),a=b(()=>{const{common:{cubicBezierEaseInOut:c},self:s}=t.value,d={"--n-bezier":c};return e.inverted?(d["--n-color"]=s.headerColorInverted,d["--n-text-color"]=s.textColorInverted,d["--n-border-color"]=s.headerBorderColorInverted):(d["--n-color"]=s.headerColor,d["--n-text-color"]=s.textColor,d["--n-border-color"]=s.headerBorderColor),d}),l=o?ne("layout-header",b(()=>e.inverted?"a":"b"),a,e):void 0;return{mergedClsPrefix:n,cssVars:o?void 0:a,themeClass:l==null?void 0:l.themeClass,onRender:l==null?void 0:l.onRender}},render(){var e;const{mergedClsPrefix:n}=this;return(e=this.onRender)===null||e===void 0||e.call(this),u("div",{class:[`${n}-layout-header`,this.themeClass,this.position&&`${n}-layout-header--${this.position}-positioned`,this.bordered&&`${n}-layout-header--bordered`],style:this.cssVars},this.$slots)}}),Y=re("n-menu"),le=re("n-submenu"),ae=re("n-menu-item-group"),J=8;function ce(e){const n=j(Y),{props:o,mergedCollapsedRef:t}=n,a=j(le,null),l=j(ae,null),c=b(()=>o.mode==="horizontal"),s=b(()=>c.value?o.dropdownPlacement:"tmNodes"in e?"right-start":"right"),d=b(()=>{var v;return Math.max((v=o.collapsedIconSize)!==null&&v!==void 0?v:o.iconSize,o.iconSize)}),g=b(()=>{var v;return!c.value&&e.root&&t.value&&(v=o.collapsedIconSize)!==null&&v!==void 0?v:o.iconSize}),I=b(()=>{if(c.value)return;const{collapsedWidth:v,indent:w,rootIndent:y}=o,{root:$,isGroup:P}=e,T=y===void 0?w:y;if($)return t.value?v/2-d.value/2:T;if(l)return w/2+l.paddingLeftRef.value;if(a)return(P?w/2:w)+a.paddingLeftRef.value}),p=b(()=>{const{collapsedWidth:v,indent:w,rootIndent:y}=o,{value:$}=d,{root:P}=e;return c.value||!P||!t.value?J:(y===void 0?w:y)+$+J-(v+$)/2});return{dropdownPlacement:s,activeIconSize:g,maxIconSize:d,paddingLeft:I,iconMarginRight:p,NMenu:n,NSubmenu:a}}const se={internalKey:{type:[String,Number],required:!0},root:Boolean,isGroup:Boolean,level:{type:Number,required:!0},title:[String,Function],extra:[String,Function]},Ae=Object.assign(Object.assign({},se),{tmNode:{type:Object,required:!0},tmNodes:{type:Array,required:!0}}),Po=S({name:"MenuOptionGroup",props:Ae,setup(e){X(le,null);const n=ce(e);X(ae,{paddingLeftRef:n.paddingLeft});const{mergedClsPrefixRef:o,props:t}=j(Y);return function(){const{value:a}=o,l=n.paddingLeft.value,{nodeProps:c}=t,s=c==null?void 0:c(e.tmNode.rawNode);return u("div",{class:`${a}-menu-item-group`,role:"group"},u("div",Object.assign({},s,{class:[`${a}-menu-item-group-title`,s==null?void 0:s.class],style:[(s==null?void 0:s.style)||"",l!==void 0?`padding-left: ${l}px;`:""]}),D(e.title),e.extra?u(be,null," ",D(e.extra)):null),u("div",null,e.tmNodes.map(d=>de(d,t))))}}}),Se=S({name:"MenuOptionContent",props:{collapsed:Boolean,disabled:Boolean,title:[String,Function],icon:Function,extra:[String,Function],showArrow:Boolean,childActive:Boolean,hover:Boolean,paddingLeft:Number,selected:Boolean,maxIconSize:{type:Number,required:!0},activeIconSize:{type:Number,required:!0},iconMarginRight:{type:Number,required:!0},clsPrefix:{type:String,required:!0},onClick:Function,tmNode:{type:Object,required:!0}},setup(e){const{props:n}=j(Y);return{menuProps:n,style:b(()=>{const{paddingLeft:o}=e;return{paddingLeft:o&&`${o}px`}}),iconStyle:b(()=>{const{maxIconSize:o,activeIconSize:t,iconMarginRight:a}=e;return{width:`${o}px`,height:`${o}px`,fontSize:`${t}px`,marginRight:`${a}px`}})}},render(){const{clsPrefix:e,tmNode:n,menuProps:{renderIcon:o,renderLabel:t,renderExtra:a,expandIcon:l}}=this,c=o?o(n.rawNode):D(this.icon);return u("div",{onClick:s=>{var d;(d=this.onClick)===null||d===void 0||d.call(this,s)},role:"none",class:[`${e}-menu-item-content`,{[`${e}-menu-item-content--selected`]:this.selected,[`${e}-menu-item-content--collapsed`]:this.collapsed,[`${e}-menu-item-content--child-active`]:this.childActive,[`${e}-menu-item-content--disabled`]:this.disabled,[`${e}-menu-item-content--hover`]:this.hover}],style:this.style},c&&u("div",{class:`${e}-menu-item-content__icon`,style:this.iconStyle,role:"none"},[c]),u("div",{class:`${e}-menu-item-content-header`,role:"none"},t?t(n.rawNode):D(this.title),this.extra||a?u("span",{class:`${e}-menu-item-content-header__extra`}," ",a?a(n.rawNode):D(this.extra)):null),this.showArrow?u(Le,{ariaHidden:!0,class:`${e}-menu-item-content__arrow`,clsPrefix:e},{default:()=>l?l(n.rawNode):u(bo,null)}):null)}}),Re=Object.assign(Object.assign({},se),{rawNodes:{type:Array,default:()=>[]},tmNodes:{type:Array,default:()=>[]},tmNode:{type:Object,required:!0},disabled:{type:Boolean,default:!1},icon:Function,onClick:Function}),$o=S({name:"Submenu",props:Re,setup(e){const n=ce(e),{NMenu:o,NSubmenu:t}=n,{props:a,mergedCollapsedRef:l,mergedThemeRef:c}=o,s=b(()=>{const{disabled:v}=e;return t!=null&&t.mergedDisabledRef.value||a.disabled?!0:v}),d=O(!1);X(le,{paddingLeftRef:n.paddingLeft,mergedDisabledRef:s}),X(ae,null);function g(){const{onClick:v}=e;v&&v()}function I(){s.value||(l.value||o.toggleExpand(e.internalKey),g())}function p(v){d.value=v}return{menuProps:a,mergedTheme:c,doSelect:o.doSelect,inverted:o.invertedRef,isHorizontal:o.isHorizontalRef,mergedClsPrefix:o.mergedClsPrefixRef,maxIconSize:n.maxIconSize,activeIconSize:n.activeIconSize,iconMarginRight:n.iconMarginRight,dropdownPlacement:n.dropdownPlacement,dropdownShow:d,paddingLeft:n.paddingLeft,mergedDisabled:s,mergedValue:o.mergedValueRef,childActive:oe(()=>o.activePathRef.value.includes(e.internalKey)),collapsed:b(()=>a.mode==="horizontal"?!1:l.value?!0:!o.mergedExpandedKeysRef.value.includes(e.internalKey)),dropdownEnabled:b(()=>!s.value&&(a.mode==="horizontal"||l.value)),handlePopoverShowChange:p,handleClick:I}},render(){var e;const{mergedClsPrefix:n,menuProps:{renderIcon:o,renderLabel:t}}=this,a=()=>{const{isHorizontal:c,paddingLeft:s,collapsed:d,mergedDisabled:g,maxIconSize:I,activeIconSize:p,title:v,childActive:w,icon:y,handleClick:$,menuProps:{nodeProps:P},dropdownShow:T,iconMarginRight:W,tmNode:M,mergedClsPrefix:E}=this,L=P==null?void 0:P(M.rawNode);return u("div",Object.assign({},L,{class:[`${E}-menu-item`,L==null?void 0:L.class],role:"menuitem"}),u(Se,{tmNode:M,paddingLeft:s,collapsed:d,disabled:g,iconMarginRight:W,maxIconSize:I,activeIconSize:p,title:v,extra:this.extra,showArrow:!c,childActive:w,clsPrefix:E,icon:y,hover:T,onClick:$}))},l=()=>u(Oe,null,{default:()=>{const{tmNodes:c,collapsed:s}=this;return s?null:u("div",{class:`${n}-submenu-children`,role:"menu"},c.map(d=>de(d,this.menuProps)))}});return this.root?u(we,Object.assign({size:"large",trigger:"hover"},(e=this.menuProps)===null||e===void 0?void 0:e.dropdownProps,{themeOverrides:this.mergedTheme.peerOverrides.Dropdown,theme:this.mergedTheme.peers.Dropdown,builtinThemeOverrides:{fontSizeLarge:"14px",optionIconSizeLarge:"18px"},value:this.mergedValue,disabled:!this.dropdownEnabled,placement:this.dropdownPlacement,keyField:this.menuProps.keyField,labelField:this.menuProps.labelField,childrenField:this.menuProps.childrenField,onUpdateShow:this.handlePopoverShowChange,options:this.rawNodes,onSelect:this.doSelect,inverted:this.inverted,renderIcon:o,renderLabel:t}),{default:()=>u("div",{class:`${n}-submenu`,role:"menuitem","aria-expanded":!this.collapsed},a(),this.isHorizontal?null:l())}):u("div",{class:`${n}-submenu`,role:"menuitem","aria-expanded":!this.collapsed},a(),l())}}),Ne=Object.assign(Object.assign({},se),{tmNode:{type:Object,required:!0},disabled:Boolean,icon:Function,onClick:Function}),ko=S({name:"MenuOption",props:Ne,setup(e){const n=ce(e),{NSubmenu:o,NMenu:t}=n,{props:a,mergedClsPrefixRef:l,mergedCollapsedRef:c}=t,s=o?o.mergedDisabledRef:{value:!1},d=b(()=>s.value||e.disabled);function g(p){const{onClick:v}=e;v&&v(p)}function I(p){d.value||(t.doSelect(e.internalKey,e.tmNode.rawNode),g(p))}return{mergedClsPrefix:l,dropdownPlacement:n.dropdownPlacement,paddingLeft:n.paddingLeft,iconMarginRight:n.iconMarginRight,maxIconSize:n.maxIconSize,activeIconSize:n.activeIconSize,mergedTheme:t.mergedThemeRef,menuProps:a,dropdownEnabled:oe(()=>e.root&&c.value&&a.mode!=="horizontal"&&!d.value),selected:oe(()=>t.mergedValueRef.value===e.internalKey),mergedDisabled:d,handleClick:I}},render(){const{mergedClsPrefix:e,mergedTheme:n,tmNode:o,menuProps:{renderLabel:t,nodeProps:a}}=this,l=a==null?void 0:a(o.rawNode);return u("div",Object.assign({},l,{role:"menuitem",class:[`${e}-menu-item`,l==null?void 0:l.class]}),u(uo,{theme:n.peers.Tooltip,themeOverrides:n.peerOverrides.Tooltip,trigger:"hover",placement:this.dropdownPlacement,disabled:!this.dropdownEnabled||this.title===void 0,internalExtraClass:["menu-tooltip"]},{default:()=>t?t(o.rawNode):D(this.title),trigger:()=>u(Se,{tmNode:o,clsPrefix:e,paddingLeft:this.paddingLeft,iconMarginRight:this.iconMarginRight,maxIconSize:this.maxIconSize,activeIconSize:this.activeIconSize,selected:this.selected,title:this.title,extra:this.extra,disabled:this.mergedDisabled,icon:this.icon,onClick:this.handleClick})}))}}),To=S({name:"MenuDivider",setup(){const e=j(Y),{mergedClsPrefixRef:n,isHorizontalRef:o}=e;return()=>o.value?null:u("div",{class:`${n.value}-menu-divider`})}}),Mo=ie(Ae),Lo=ie(Ne),Oo=ie(Re);function Pe(e){return e.type==="divider"||e.type==="render"}function Fo(e){return e.type==="divider"}function de(e,n){const{rawNode:o}=e,{show:t}=o;if(t===!1)return null;if(Pe(o))return Fo(o)?u(To,Object.assign({key:e.key},o.props)):null;const{labelField:a}=n,{key:l,level:c,isGroup:s}=e,d=Object.assign(Object.assign({},o),{title:o.title||o[a],extra:o.titleExtra||o.extra,key:l,internalKey:l,level:c,root:c===0,isGroup:s});return e.children?e.isGroup?u(Po,ee(d,Mo,{tmNode:e,tmNodes:e.children,key:l})):u($o,ee(d,Oo,{key:l,rawNodes:o[n.childrenField],tmNodes:e.children,tmNode:e})):u(ko,ee(d,Lo,{key:l,tmNode:e}))}const ge=[C("&::before","background-color: var(--n-item-color-hover);"),h("arrow",`
 color: var(--n-arrow-color-hover);
 `),h("icon",`
 color: var(--n-item-icon-color-hover);
 `),m("menu-item-content-header",`
 color: var(--n-item-text-color-hover);
 `,[C("a",`
 color: var(--n-item-text-color-hover);
 `),h("extra",`
 color: var(--n-item-text-color-hover);
 `)])],xe=[h("icon",`
 color: var(--n-item-icon-color-hover-horizontal);
 `),m("menu-item-content-header",`
 color: var(--n-item-text-color-hover-horizontal);
 `,[C("a",`
 color: var(--n-item-text-color-hover-horizontal);
 `),h("extra",`
 color: var(--n-item-text-color-hover-horizontal);
 `)])],Eo=C([m("menu",`
 background-color: var(--n-color);
 color: var(--n-item-text-color);
 overflow: hidden;
 transition: background-color .3s var(--n-bezier);
 box-sizing: border-box;
 font-size: var(--n-font-size);
 padding-bottom: 6px;
 `,[_("horizontal",`
 display: inline-flex;
 padding-bottom: 0;
 `,[m("submenu","margin: 0;"),m("menu-item","margin: 0;"),m("menu-item-content",`
 padding: 0 20px;
 border-bottom: 2px solid #0000;
 `,[C("&::before","display: none;"),_("selected","border-bottom: 2px solid var(--n-border-color-horizontal)")]),m("menu-item-content",[_("selected",[h("icon","color: var(--n-item-icon-color-active-horizontal);"),m("menu-item-content-header",`
 color: var(--n-item-text-color-active-horizontal);
 `,[C("a","color: var(--n-item-text-color-active-horizontal);"),h("extra","color: var(--n-item-text-color-active-horizontal);")])]),_("child-active",`
 border-bottom: 2px solid var(--n-border-color-horizontal);
 `,[m("menu-item-content-header",`
 color: var(--n-item-text-color-child-active-horizontal);
 `,[C("a",`
 color: var(--n-item-text-color-child-active-horizontal);
 `),h("extra",`
 color: var(--n-item-text-color-child-active-horizontal);
 `)]),h("icon",`
 color: var(--n-item-icon-color-child-active-horizontal);
 `)]),Z("disabled",[Z("selected, child-active",[C("&:focus-within",xe)]),_("selected",[K(null,[h("icon","color: var(--n-item-icon-color-active-hover-horizontal);"),m("menu-item-content-header",`
 color: var(--n-item-text-color-active-hover-horizontal);
 `,[C("a","color: var(--n-item-text-color-active-hover-horizontal);"),h("extra","color: var(--n-item-text-color-active-hover-horizontal);")])])]),_("child-active",[K(null,[h("icon","color: var(--n-item-icon-color-child-active-hover-horizontal);"),m("menu-item-content-header",`
 color: var(--n-item-text-color-child-active-hover-horizontal);
 `,[C("a","color: var(--n-item-text-color-child-active-hover-horizontal);"),h("extra","color: var(--n-item-text-color-child-active-hover-horizontal);")])])]),K("border-bottom: 2px solid var(--n-border-color-horizontal);",xe)]),m("menu-item-content-header",[C("a","color: var(--n-item-text-color-horizontal);")])])]),_("collapsed",[m("menu-item-content",[_("selected",[C("&::before",`
 background-color: var(--n-item-color-active-collapsed) !important;
 `)]),m("menu-item-content-header","opacity: 0;"),h("arrow","opacity: 0;"),h("icon","color: var(--n-item-icon-color-collapsed);")])]),m("menu-item",`
 height: var(--n-item-height);
 margin-top: 6px;
 position: relative;
 `),m("menu-item-content",`
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
 `),_("disabled",`
 opacity: .45;
 cursor: not-allowed;
 `),_("collapsed",[h("arrow","transform: rotate(0);")]),_("selected",[C("&::before","background-color: var(--n-item-color-active);"),h("arrow","color: var(--n-arrow-color-active);"),h("icon","color: var(--n-item-icon-color-active);"),m("menu-item-content-header",`
 color: var(--n-item-text-color-active);
 `,[C("a","color: var(--n-item-text-color-active);"),h("extra","color: var(--n-item-text-color-active);")])]),_("child-active",[m("menu-item-content-header",`
 color: var(--n-item-text-color-child-active);
 `,[C("a",`
 color: var(--n-item-text-color-child-active);
 `),h("extra",`
 color: var(--n-item-text-color-child-active);
 `)]),h("arrow",`
 color: var(--n-arrow-color-child-active);
 `),h("icon",`
 color: var(--n-item-icon-color-child-active);
 `)]),Z("disabled",[Z("selected, child-active",[C("&:focus-within",ge)]),_("selected",[K(null,[h("arrow","color: var(--n-arrow-color-active-hover);"),h("icon","color: var(--n-item-icon-color-active-hover);"),m("menu-item-content-header",`
 color: var(--n-item-text-color-active-hover);
 `,[C("a","color: var(--n-item-text-color-active-hover);"),h("extra","color: var(--n-item-text-color-active-hover);")])])]),_("child-active",[K(null,[h("arrow","color: var(--n-arrow-color-child-active-hover);"),h("icon","color: var(--n-item-icon-color-child-active-hover);"),m("menu-item-content-header",`
 color: var(--n-item-text-color-child-active-hover);
 `,[C("a","color: var(--n-item-text-color-child-active-hover);"),h("extra","color: var(--n-item-text-color-child-active-hover);")])])]),_("selected",[K(null,[C("&::before","background-color: var(--n-item-color-active-hover);")])]),K(null,ge)]),h("icon",`
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
 `),h("arrow",`
 grid-area: arrow;
 font-size: 16px;
 color: var(--n-arrow-color);
 transform: rotate(180deg);
 opacity: 1;
 transition:
 color .3s var(--n-bezier),
 transform 0.2s var(--n-bezier),
 opacity 0.2s var(--n-bezier);
 `),m("menu-item-content-header",`
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
 `)]),h("extra",`
 font-size: .93em;
 color: var(--n-group-text-color);
 transition: color .3s var(--n-bezier);
 `)])]),m("submenu",`
 cursor: pointer;
 position: relative;
 margin-top: 6px;
 `,[m("menu-item-content",`
 height: var(--n-item-height);
 `),m("submenu-children",`
 overflow: hidden;
 padding: 0;
 `,[Fe({duration:".2s"})])]),m("menu-item-group",[m("menu-item-group-title",`
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
 `)])]),m("menu-tooltip",[C("a",`
 color: inherit;
 text-decoration: none;
 `)]),m("menu-divider",`
 transition: background-color .3s var(--n-bezier);
 background-color: var(--n-divider-color);
 height: 1px;
 margin: 6px 18px;
 `)]);function K(e,n){return[_("hover",e,n),C("&:hover",e,n)]}const Bo=Object.assign(Object.assign({},q.props),{options:{type:Array,default:()=>[]},collapsed:{type:Boolean,default:void 0},collapsedWidth:{type:Number,default:48},iconSize:{type:Number,default:20},collapsedIconSize:{type:Number,default:24},rootIndent:Number,indent:{type:Number,default:32},labelField:{type:String,default:"label"},keyField:{type:String,default:"key"},childrenField:{type:String,default:"children"},disabledField:{type:String,default:"disabled"},defaultExpandAll:Boolean,defaultExpandedKeys:Array,expandedKeys:Array,value:[String,Number],defaultValue:{type:[String,Number],default:null},mode:{type:String,default:"vertical"},watchProps:{type:Array,default:void 0},disabled:Boolean,show:{type:Boolean,defalut:!0},inverted:Boolean,"onUpdate:expandedKeys":[Function,Array],onUpdateExpandedKeys:[Function,Array],onUpdateValue:[Function,Array],"onUpdate:value":[Function,Array],expandIcon:Function,renderIcon:Function,renderLabel:Function,renderExtra:Function,dropdownProps:Object,accordion:Boolean,nodeProps:Function,items:Array,onOpenNamesChange:[Function,Array],onSelect:[Function,Array],onExpandedNamesChange:[Function,Array],expandedNames:Array,defaultExpandedNames:Array,dropdownPlacement:{type:String,default:"bottom"}}),Ko=S({name:"Menu",props:Bo,setup(e){const{mergedClsPrefixRef:n,inlineThemeDisabled:o}=te(e),t=q("Menu","-menu",Eo,wo,e,n),a=j(fo,null),l=b(()=>{var f;const{collapsed:z}=e;if(z!==void 0)return z;if(a){const{collapseModeRef:r,collapsedRef:x}=a;if(r.value==="width")return(f=x.value)!==null&&f!==void 0?f:!1}return!1}),c=b(()=>{const{keyField:f,childrenField:z,disabledField:r}=e;return vo(e.items||e.options,{getIgnored(x){return Pe(x)},getChildren(x){return x[z]},getDisabled(x){return x[r]},getKey(x){var R;return(R=x[f])!==null&&R!==void 0?R:x.name}})}),s=b(()=>new Set(c.value.treeNodes.map(f=>f.key))),{watchProps:d}=e,g=O(null);d!=null&&d.includes("defaultValue")?ve(()=>{g.value=e.defaultValue}):g.value=e.defaultValue;const I=me(e,"value"),p=he(I,g),v=O([]),w=()=>{v.value=e.defaultExpandAll?c.value.getNonLeafKeys():e.defaultExpandedNames||e.defaultExpandedKeys||c.value.getPath(p.value,{includeSelf:!1}).keyPath};d!=null&&d.includes("defaultExpandedKeys")?ve(w):w();const y=_e(e,["expandedNames","expandedKeys"]),$=he(y,v),P=b(()=>c.value.treeNodes),T=b(()=>c.value.getPath(p.value).keyPath);X(Y,{props:e,mergedCollapsedRef:l,mergedThemeRef:t,mergedValueRef:p,mergedExpandedKeysRef:$,activePathRef:T,mergedClsPrefixRef:n,isHorizontalRef:b(()=>e.mode==="horizontal"),invertedRef:me(e,"inverted"),doSelect:W,toggleExpand:E});function W(f,z){const{"onUpdate:value":r,onUpdateValue:x,onSelect:R}=e;x&&B(x,f,z),r&&B(r,f,z),R&&B(R,f,z),g.value=f}function M(f){const{"onUpdate:expandedKeys":z,onUpdateExpandedKeys:r,onExpandedNamesChange:x,onOpenNamesChange:R}=e;z&&B(z,f),r&&B(r,f),x&&B(x,f),R&&B(R,f),v.value=f}function E(f){const z=Array.from($.value),r=z.findIndex(x=>x===f);if(~r)z.splice(r,1);else{if(e.accordion&&s.value.has(f)){const x=z.findIndex(R=>s.value.has(R));x>-1&&z.splice(x,1)}z.push(f)}M(z)}const L=f=>{const z=c.value.getPath(f??p.value,{includeSelf:!1}).keyPath;if(!z.length)return;const r=Array.from($.value),x=new Set([...r,...z]);e.accordion&&s.value.forEach(R=>{x.has(R)&&!z.includes(R)&&x.delete(R)}),M(Array.from(x))},ue=b(()=>{const{inverted:f}=e,{common:{cubicBezierEaseInOut:z},self:r}=t.value,{borderRadius:x,borderColorHorizontal:R,fontSize:$e,itemHeight:ke,dividerColor:Te}=r,i={"--n-divider-color":Te,"--n-bezier":z,"--n-font-size":$e,"--n-border-color-horizontal":R,"--n-border-radius":x,"--n-item-height":ke};return f?(i["--n-group-text-color"]=r.groupTextColorInverted,i["--n-color"]=r.colorInverted,i["--n-item-text-color"]=r.itemTextColorInverted,i["--n-item-text-color-hover"]=r.itemTextColorHoverInverted,i["--n-item-text-color-active"]=r.itemTextColorActiveInverted,i["--n-item-text-color-child-active"]=r.itemTextColorChildActiveInverted,i["--n-item-text-color-child-active-hover"]=r.itemTextColorChildActiveInverted,i["--n-item-text-color-active-hover"]=r.itemTextColorActiveHoverInverted,i["--n-item-icon-color"]=r.itemIconColorInverted,i["--n-item-icon-color-hover"]=r.itemIconColorHoverInverted,i["--n-item-icon-color-active"]=r.itemIconColorActiveInverted,i["--n-item-icon-color-active-hover"]=r.itemIconColorActiveHoverInverted,i["--n-item-icon-color-child-active"]=r.itemIconColorChildActiveInverted,i["--n-item-icon-color-child-active-hover"]=r.itemIconColorChildActiveHoverInverted,i["--n-item-icon-color-collapsed"]=r.itemIconColorCollapsedInverted,i["--n-item-text-color-horizontal"]=r.itemTextColorHorizontalInverted,i["--n-item-text-color-hover-horizontal"]=r.itemTextColorHoverHorizontalInverted,i["--n-item-text-color-active-horizontal"]=r.itemTextColorActiveHorizontalInverted,i["--n-item-text-color-child-active-horizontal"]=r.itemTextColorChildActiveHorizontalInverted,i["--n-item-text-color-child-active-hover-horizontal"]=r.itemTextColorChildActiveHoverHorizontalInverted,i["--n-item-text-color-active-hover-horizontal"]=r.itemTextColorActiveHoverHorizontalInverted,i["--n-item-icon-color-horizontal"]=r.itemIconColorHorizontalInverted,i["--n-item-icon-color-hover-horizontal"]=r.itemIconColorHoverHorizontalInverted,i["--n-item-icon-color-active-horizontal"]=r.itemIconColorActiveHorizontalInverted,i["--n-item-icon-color-active-hover-horizontal"]=r.itemIconColorActiveHoverHorizontalInverted,i["--n-item-icon-color-child-active-horizontal"]=r.itemIconColorChildActiveHorizontalInverted,i["--n-item-icon-color-child-active-hover-horizontal"]=r.itemIconColorChildActiveHoverHorizontalInverted,i["--n-arrow-color"]=r.arrowColorInverted,i["--n-arrow-color-hover"]=r.arrowColorHoverInverted,i["--n-arrow-color-active"]=r.arrowColorActiveInverted,i["--n-arrow-color-active-hover"]=r.arrowColorActiveHoverInverted,i["--n-arrow-color-child-active"]=r.arrowColorChildActiveInverted,i["--n-arrow-color-child-active-hover"]=r.arrowColorChildActiveHoverInverted,i["--n-item-color-hover"]=r.itemColorHoverInverted,i["--n-item-color-active"]=r.itemColorActiveInverted,i["--n-item-color-active-hover"]=r.itemColorActiveHoverInverted,i["--n-item-color-active-collapsed"]=r.itemColorActiveCollapsedInverted):(i["--n-group-text-color"]=r.groupTextColor,i["--n-color"]=r.color,i["--n-item-text-color"]=r.itemTextColor,i["--n-item-text-color-hover"]=r.itemTextColorHover,i["--n-item-text-color-active"]=r.itemTextColorActive,i["--n-item-text-color-child-active"]=r.itemTextColorChildActive,i["--n-item-text-color-child-active-hover"]=r.itemTextColorChildActiveHover,i["--n-item-text-color-active-hover"]=r.itemTextColorActiveHover,i["--n-item-icon-color"]=r.itemIconColor,i["--n-item-icon-color-hover"]=r.itemIconColorHover,i["--n-item-icon-color-active"]=r.itemIconColorActive,i["--n-item-icon-color-active-hover"]=r.itemIconColorActiveHover,i["--n-item-icon-color-child-active"]=r.itemIconColorChildActive,i["--n-item-icon-color-child-active-hover"]=r.itemIconColorChildActiveHover,i["--n-item-icon-color-collapsed"]=r.itemIconColorCollapsed,i["--n-item-text-color-horizontal"]=r.itemTextColorHorizontal,i["--n-item-text-color-hover-horizontal"]=r.itemTextColorHoverHorizontal,i["--n-item-text-color-active-horizontal"]=r.itemTextColorActiveHorizontal,i["--n-item-text-color-child-active-horizontal"]=r.itemTextColorChildActiveHorizontal,i["--n-item-text-color-child-active-hover-horizontal"]=r.itemTextColorChildActiveHoverHorizontal,i["--n-item-text-color-active-hover-horizontal"]=r.itemTextColorActiveHoverHorizontal,i["--n-item-icon-color-horizontal"]=r.itemIconColorHorizontal,i["--n-item-icon-color-hover-horizontal"]=r.itemIconColorHoverHorizontal,i["--n-item-icon-color-active-horizontal"]=r.itemIconColorActiveHorizontal,i["--n-item-icon-color-active-hover-horizontal"]=r.itemIconColorActiveHoverHorizontal,i["--n-item-icon-color-child-active-horizontal"]=r.itemIconColorChildActiveHorizontal,i["--n-item-icon-color-child-active-hover-horizontal"]=r.itemIconColorChildActiveHoverHorizontal,i["--n-arrow-color"]=r.arrowColor,i["--n-arrow-color-hover"]=r.arrowColorHover,i["--n-arrow-color-active"]=r.arrowColorActive,i["--n-arrow-color-active-hover"]=r.arrowColorActiveHover,i["--n-arrow-color-child-active"]=r.arrowColorChildActive,i["--n-arrow-color-child-active-hover"]=r.arrowColorChildActiveHover,i["--n-item-color-hover"]=r.itemColorHover,i["--n-item-color-active"]=r.itemColorActive,i["--n-item-color-active-hover"]=r.itemColorActiveHover,i["--n-item-color-active-collapsed"]=r.itemColorActiveCollapsed),i}),V=o?ne("menu",b(()=>e.inverted?"a":"b"),ue,e):void 0;return{mergedClsPrefix:n,controlledExpandedKeys:y,uncontrolledExpanededKeys:v,mergedExpandedKeys:$,uncontrolledValue:g,mergedValue:p,activePath:T,tmNodes:P,mergedTheme:t,mergedCollapsed:l,cssVars:o?void 0:ue,themeClass:V==null?void 0:V.themeClass,onRender:V==null?void 0:V.onRender,showOption:L}},render(){const{mergedClsPrefix:e,mode:n,themeClass:o,onRender:t}=this;return t==null||t(),u("div",{role:n==="horizontal"?"menubar":"menu",class:[`${e}-menu`,o,`${e}-menu--${n}`,this.mergedCollapsed&&`${e}-menu--collapsed`],style:this.cssVars},this.tmNodes.map(a=>de(a,this.$props)))}}),jo=C([C("@keyframes spin-rotate",`
 from {
 transform: rotate(0);
 }
 to {
 transform: rotate(360deg);
 }
 `),m("spin-container",{position:"relative"},[m("spin-body",`
 position: absolute;
 top: 50%;
 left: 50%;
 transform: translateX(-50%) translateY(-50%);
 `,[Ee()])]),m("spin-body",`
 display: inline-flex;
 align-items: center;
 justify-content: center;
 flex-direction: column;
 `),m("spin",`
 display: inline-flex;
 height: var(--n-size);
 width: var(--n-size);
 font-size: var(--n-size);
 color: var(--n-color);
 `,[_("rotate",`
 animation: spin-rotate 2s linear infinite;
 `)]),m("spin-description",`
 display: inline-block;
 font-size: var(--n-font-size);
 color: var(--n-text-color);
 transition: color .3s var(--n-bezier);
 margin-top: 8px;
 `),m("spin-content",`
 opacity: 1;
 transition: opacity .3s var(--n-bezier);
 pointer-events: all;
 `,[_("spinning",`
 user-select: none;
 -webkit-user-select: none;
 pointer-events: none;
 opacity: var(--n-opacity-spinning);
 `)])]),Vo={small:20,medium:18,large:16},Do=Object.assign(Object.assign({},q.props),{description:String,stroke:String,size:{type:[String,Number],default:"medium"},show:{type:Boolean,default:!0},strokeWidth:Number,rotate:{type:Boolean,default:!0},spinning:{type:Boolean,validator:()=>!0,default:void 0}}),Uo=S({name:"Spin",props:Do,setup(e){const{mergedClsPrefixRef:n,inlineThemeDisabled:o}=te(e),t=q("Spin","-spin",jo,Ao,e,n),a=b(()=>{const{size:c}=e,{common:{cubicBezierEaseInOut:s},self:d}=t.value,{opacitySpinning:g,color:I,textColor:p}=d,v=typeof c=="number"?je(c):d[Ve("size",c)];return{"--n-bezier":s,"--n-opacity-spinning":g,"--n-size":v,"--n-color":I,"--n-text-color":p}}),l=o?ne("spin",b(()=>{const{size:c}=e;return typeof c=="number"?String(c):c[0]}),a,e):void 0;return{mergedClsPrefix:n,compitableShow:_e(e,["spinning","show"]),mergedStrokeWidth:b(()=>{const{strokeWidth:c}=e;if(c!==void 0)return c;const{size:s}=e;return Vo[typeof s=="number"?"medium":s]}),cssVars:o?void 0:a,themeClass:l==null?void 0:l.themeClass,onRender:l==null?void 0:l.onRender}},render(){var e,n;const{$slots:o,mergedClsPrefix:t,description:a}=this,l=o.icon&&this.rotate,c=(a||o.description)&&u("div",{class:`${t}-spin-description`},a||((e=o.description)===null||e===void 0?void 0:e.call(o))),s=o.icon?u("div",{class:[`${t}-spin-body`,this.themeClass]},u("div",{class:[`${t}-spin`,l&&`${t}-spin--rotate`],style:o.default?"":this.cssVars},o.icon()),c):u("div",{class:[`${t}-spin-body`,this.themeClass]},u(Be,{clsPrefix:t,style:o.default?"":this.cssVars,stroke:this.stroke,"stroke-width":this.mergedStrokeWidth,class:`${t}-spin`}),c);return(n=this.onRender)===null||n===void 0||n.call(this),o.default?u("div",{class:[`${t}-spin-container`,this.themeClass],style:this.cssVars},u("div",{class:[`${t}-spin-content`,this.compitableShow&&`${t}-spin-content--spinning`]},o),u(Ke,{name:"fade-in-transition"},{default:()=>this.compitableShow?s:null})):s}}),qo="/assets/cmdb_icon-a206a804.jpg",Go="/assets/cmdb_logo-3b556475.jpg",Wo={class:"logo-container",flex:"","items-center":"","justify-center":"","h-12":""},Xo={key:0,src:qo,class:"h-80%",alt:""},Yo={key:1,src:Go,class:"h-80%",alt:""},Zo=S({__name:"BaseLogo",props:{isMini:{type:Boolean}},setup(e){const n=e;return(o,t)=>(k(),U("div",Wo,[n.isMini?(k(),U("img",Xo)):(k(),U("img",Yo))]))}}),Jo={class:"icon","aria-hidden":"true"},Qo=["xlink:href"],et=S({__name:"SvgIcon",props:{path:{default:""}},setup(e){const n=e;return(o,t)=>(k(),G(H(mo),De(Ue(o.$attrs)),{default:N(()=>[(k(),U("svg",Jo,[F("use",{"xlink:href":n.path},null,8,Qo)]))]),_:1},16))}}),ot=S({__name:"MenuGroup",setup(e){const n=Ge(),o=We(),t=ze(),{roles:a}=ye(t),l=O(String(n.name));qe(()=>n.fullPath,()=>{l.value=String(n.name)});const c=b(()=>o.getRoutes().find(g=>g.name==="root"));function s(g){const I=u(et,{path:`#icon-${g}`});return()=>u(I,{})}const d=b(()=>{const g=[...c.value.children];function I(p,v){return p.map(w=>{var E;const{meta:y}=w;w.children&&(w.children=I(w.children));const P=Xe(a.value,y==null?void 0:y.role)&&!(y!=null&&y.hideMenu),T=!!(y!=null&&y.icon),W=y==null?void 0:y.icon;let M;return P&&(w.children?M=(E=w.meta)==null?void 0:E.title:M=u(Ye,{to:w.path},{default:()=>{var L;return((L=w.meta)==null?void 0:L.title)??"--"}})),{...w,label:()=>P&&M,show:P,key:w.name,icon:T?s(W):null}})}return I(g)});return(g,I)=>(k(),G(H(Ko),{mode:"vertical",value:l.value,indent:20,accordion:!1,"collapsed-width":64,"collapsed-icon-size":22,options:H(d)},null,8,["value","options"]))}}),tt=S({__name:"BaseSider",setup(e){const n=O(!1);function o(t){n.value=t}return(t,a)=>(k(),G(H(go),{width:"200",bordered:"","native-scrollbar":!1,"show-trigger":"","collapse-mode":"width","collapsed-width":64,"onUpdate:collapsed":o},{default:N(()=>[A(Zo,{"is-mini":n.value},null,8,["is-mini"]),A(ot)]),_:1}))}}),nt={class:"avator"},rt=["alt"],it={class:"avator-container"},lt=S({__name:"AutoTextAvatar",props:{text:{default:""}},setup(e){const n=e,{text:o}=Ze(n);return(t,a)=>(k(),U("div",nt,[F("div",{class:"avator-inner",alt:H(o)},[F("div",it,[F("span",null,Je(H(o)),1)])],8,rt)]))}});const at=He(lt,[["__scopeId","data-v-e59811c3"]]),ct={flex:"","justify-end":""},st=S({__name:"UserSwitchModal",emits:["refresh"],setup(e,{emit:n}){const o=O(!1),{run:t,loading:a}=Qe(Ie.userSwitchById),l=O(null),c=O({userId:void 0}),s={userId:{required:!0,message:"请选择用户",trigger:"blur",type:"number"}};function d(){o.value=!0}async function g(){var I;(I=l.value)==null||I.validate(async p=>{if(p)return;const{result:v}=await t(Number(c.value.userId));v&&(o.value=!1,n("refresh"))})}return(I,p)=>(k(),U(be,null,[F("span",{onClick:d},[eo(I.$slots,"default")]),A(H(io),{show:o.value,"onUpdate:show":p[2]||(p[2]=v=>o.value=v),title:"切换用户",preset:"card","auto-focus":!1,"w-100":""},{default:N(()=>[A(H(oo),{ref_key:"formRef",ref:l,model:c.value,rules:s},{default:N(()=>[A(H(to),{path:"userId","show-label":!1},{default:N(()=>[A(Co,{value:c.value.userId,"onUpdate:value":p[0]||(p[0]=v=>c.value.userId=v)},null,8,["value"])]),_:1}),F("div",ct,[A(H(no),{type:"primary",loading:H(a),onClick:p[1]||(p[1]=v=>g())},{default:N(()=>[ro(" 确定 ")]),_:1},8,["loading"])])]),_:1},8,["model"])]),_:1},8,["show"])],64))}}),dt={"h-12":"","w-auto":"",flex:"","px-5":"","items-center":"","justify-between":""},ut=F("span",null,null,-1),vt=S({__name:"BaseHeader",setup(e){const n=ze(),{nickname:o,isAdmin:t}=ye(n);function a(){window.location.reload()}async function l(){const{result:s}=await Ie.toLogout();a()}const c=[{label:()=>u(st,{onRefresh:a},{default:()=>u("div",{},"切换用户")}),show:t.value},{label:"退出登录",props:{onClick:l}}];return(s,d)=>(k(),G(H(No),{bordered:""},{default:N(()=>[F("div",dt,[A(H(fe),{align:"center",justify:"center"},{default:N(()=>[ut]),_:1}),A(H(fe),{align:"center",justify:"center"},{default:N(()=>[A(H(we),{"display-directive":"show",options:c},{default:N(()=>[A(at,{text:H(o)},null,8,["text"])]),_:1})]),_:1})])]),_:1}))}}),mt=S({__name:"AppMain",setup(e){return(n,o)=>{const t=ao("RouterView");return k(),G(lo,null,{fallback:N(()=>[A(H(Uo),{size:"large"})]),default:N(()=>[A(t)]),_:1})}}});const ht=He(mt,[["__scopeId","data-v-91d1c4f9"]]),yt=S({__name:"PageLayout",setup(e){return(n,o)=>(k(),G(H(pe),{"has-sider":"",position:"absolute"},{default:N(()=>[A(tt),A(H(pe),{"native-scrollbar":!0,style:{overflow:"initial"}},{default:N(()=>[A(vt),A(H(xo),{position:"absolute",style:{"margin-top":"50px"},"content-style":"padding:16px"},{default:N(()=>[A(ht)]),_:1})]),_:1})]),_:1}))}});export{yt as default};
