import{N as ue}from"./GradientText-869ca789.js";import{d as N,h as v,c as ke,a as Le,b as Y,e as h,f as A,u as Q,g as Ce,i as I,j as be,k as re,l as V,p as G,r as U,F as Ie,N as Fe,m as F,n as oe,o as Oe,q as ee,s as ne,t as C,v as m,w as J,x as Ee,y as ve,z as me,A as B,B as M,C as X,D as E,E as R,G as te,H as w,I as O,J as Ke,K as Be,L as ze,M as ye,O as je,P as Ve,Q as De,R as Ue,S as qe,T as H,U as Ge,V as We,W as Je,X as Ze,Y as Qe,Z as we,_ as Xe,$ as Ye}from"./index-04cf60af.js";import{u as he,a as eo,N as oo}from"./Icon-b6ee2b01.js";import{p as to,l as ro,a as no,N as io,b as pe}from"./LayoutContent-bab11fa4.js";import{t as lo,d as ao,N as _e,a as co,c as so}from"./Dropdown-ff39b5a5.js";import{N as uo}from"./LayoutSider-161f480f.js";import{_ as Ae}from"./_plugin-vue_export-helper-c27b6911.js";import{u as vo}from"./useRequest-4a6f3b67.js";import{_ as mo}from"./UserSelect.vue_vue_type_script_setup_true_lang-a504e86d.js";import{N as ho,a as po}from"./FormItem-960404fe.js";import{N as fe}from"./Space-d4aba141.js";import{N as fo}from"./Spin-da978b45.js";import"./FocusDetector-40d296f1.js";import"./get-a19d687e.js";const go=N({name:"ChevronDownFilled",render(){return v("svg",{viewBox:"0 0 16 16",fill:"none",xmlns:"http://www.w3.org/2000/svg"},v("path",{d:"M3.20041 5.73966C3.48226 5.43613 3.95681 5.41856 4.26034 5.70041L8 9.22652L11.7397 5.70041C12.0432 5.41856 12.5177 5.43613 12.7996 5.73966C13.0815 6.0432 13.0639 6.51775 12.7603 6.7996L8.51034 10.7996C8.22258 11.0668 7.77743 11.0668 7.48967 10.7996L3.23966 6.7996C2.93613 6.51775 2.91856 6.0432 3.20041 5.73966Z",fill:"currentColor"}))}});function xo(e,r,o,n){return{itemColorHoverInverted:"#0000",itemColorActiveInverted:r,itemColorActiveHoverInverted:r,itemColorActiveCollapsedInverted:r,itemTextColorInverted:e,itemTextColorHoverInverted:o,itemTextColorChildActiveInverted:o,itemTextColorChildActiveHoverInverted:o,itemTextColorActiveInverted:o,itemTextColorActiveHoverInverted:o,itemTextColorHorizontalInverted:e,itemTextColorHoverHorizontalInverted:o,itemTextColorChildActiveHorizontalInverted:o,itemTextColorChildActiveHoverHorizontalInverted:o,itemTextColorActiveHorizontalInverted:o,itemTextColorActiveHoverHorizontalInverted:o,itemIconColorInverted:e,itemIconColorHoverInverted:o,itemIconColorActiveInverted:o,itemIconColorActiveHoverInverted:o,itemIconColorChildActiveInverted:o,itemIconColorChildActiveHoverInverted:o,itemIconColorCollapsedInverted:e,itemIconColorHorizontalInverted:e,itemIconColorHoverHorizontalInverted:o,itemIconColorActiveHorizontalInverted:o,itemIconColorActiveHoverHorizontalInverted:o,itemIconColorChildActiveHorizontalInverted:o,itemIconColorChildActiveHoverHorizontalInverted:o,arrowColorInverted:e,arrowColorHoverInverted:o,arrowColorActiveInverted:o,arrowColorActiveHoverInverted:o,arrowColorChildActiveInverted:o,arrowColorChildActiveHoverInverted:o,groupTextColorInverted:n}}const Co=e=>{const{borderRadius:r,textColor3:o,primaryColor:n,textColor2:l,textColor1:a,fontSize:s,dividerColor:d,hoverColor:c,primaryColorHover:x}=e;return Object.assign({borderRadius:r,color:"#0000",groupTextColor:o,itemColorHover:c,itemColorActive:Y(n,{alpha:.1}),itemColorActiveHover:Y(n,{alpha:.1}),itemColorActiveCollapsed:Y(n,{alpha:.1}),itemTextColor:l,itemTextColorHover:l,itemTextColorActive:n,itemTextColorActiveHover:n,itemTextColorChildActive:n,itemTextColorChildActiveHover:n,itemTextColorHorizontal:l,itemTextColorHoverHorizontal:x,itemTextColorActiveHorizontal:n,itemTextColorActiveHoverHorizontal:n,itemTextColorChildActiveHorizontal:n,itemTextColorChildActiveHoverHorizontal:n,itemIconColor:a,itemIconColorHover:a,itemIconColorActive:n,itemIconColorActiveHover:n,itemIconColorChildActive:n,itemIconColorChildActiveHover:n,itemIconColorCollapsed:a,itemIconColorHorizontal:a,itemIconColorHoverHorizontal:x,itemIconColorActiveHorizontal:n,itemIconColorActiveHoverHorizontal:n,itemIconColorChildActiveHorizontal:n,itemIconColorChildActiveHoverHorizontal:n,itemHeight:"42px",arrowColor:l,arrowColorHover:l,arrowColorActive:n,arrowColorActiveHover:n,arrowColorChildActive:n,arrowColorChildActiveHover:n,colorInverted:"#0000",borderColorHorizontal:"#0000",fontSize:s,dividerColor:d},xo("#BBB",n,"#FFF","#AAA"))},bo=ke({name:"Menu",common:Le,peers:{Tooltip:lo,Dropdown:ao},self:Co}),Io=bo,zo=h("layout-header",`
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
 `)]),yo={position:to,inverted:Boolean,bordered:{type:Boolean,default:!1}},wo=N({name:"LayoutHeader",props:Object.assign(Object.assign({},Q.props),yo),setup(e){const{mergedClsPrefixRef:r,inlineThemeDisabled:o}=Ce(e),n=Q("Layout","-layout-header",zo,ro,e,r),l=I(()=>{const{common:{cubicBezierEaseInOut:s},self:d}=n.value,c={"--n-bezier":s};return e.inverted?(c["--n-color"]=d.headerColorInverted,c["--n-text-color"]=d.textColorInverted,c["--n-border-color"]=d.headerBorderColorInverted):(c["--n-color"]=d.headerColor,c["--n-text-color"]=d.textColor,c["--n-border-color"]=d.headerBorderColor),c}),a=o?be("layout-header",I(()=>e.inverted?"a":"b"),l,e):void 0;return{mergedClsPrefix:r,cssVars:o?void 0:l,themeClass:a==null?void 0:a.themeClass,onRender:a==null?void 0:a.onRender}},render(){var e;const{mergedClsPrefix:r}=this;return(e=this.onRender)===null||e===void 0||e.call(this),v("div",{class:[`${r}-layout-header`,this.themeClass,this.position&&`${r}-layout-header--${this.position}-positioned`,this.bordered&&`${r}-layout-header--bordered`],style:this.cssVars},this.$slots)}}),W=re("n-menu"),ie=re("n-submenu"),le=re("n-menu-item-group"),Z=8;function ae(e){const r=V(W),{props:o,mergedCollapsedRef:n}=r,l=V(ie,null),a=V(le,null),s=I(()=>o.mode==="horizontal"),d=I(()=>s.value?o.dropdownPlacement:"tmNodes"in e?"right-start":"right"),c=I(()=>{var u;return Math.max((u=o.collapsedIconSize)!==null&&u!==void 0?u:o.iconSize,o.iconSize)}),x=I(()=>{var u;return!s.value&&e.root&&n.value&&(u=o.collapsedIconSize)!==null&&u!==void 0?u:o.iconSize}),_=I(()=>{if(s.value)return;const{collapsedWidth:u,indent:y,rootIndent:z}=o,{root:T,isGroup:P}=e,$=z===void 0?y:z;if(T)return n.value?u/2-c.value/2:$;if(a)return y/2+a.paddingLeftRef.value;if(l)return(P?y/2:y)+l.paddingLeftRef.value}),g=I(()=>{const{collapsedWidth:u,indent:y,rootIndent:z}=o,{value:T}=c,{root:P}=e;return s.value||!P||!n.value?Z:(z===void 0?y:z)+T+Z-(u+T)/2});return{dropdownPlacement:d,activeIconSize:x,maxIconSize:c,paddingLeft:_,iconMarginRight:g,NMenu:r,NSubmenu:l}}const ce={internalKey:{type:[String,Number],required:!0},root:Boolean,isGroup:Boolean,level:{type:Number,required:!0},title:[String,Function],extra:[String,Function]},He=Object.assign(Object.assign({},ce),{tmNode:{type:Object,required:!0},tmNodes:{type:Array,required:!0}}),_o=N({name:"MenuOptionGroup",props:He,setup(e){G(ie,null);const r=ae(e);G(le,{paddingLeftRef:r.paddingLeft});const{mergedClsPrefixRef:o,props:n}=V(W);return function(){const{value:l}=o,a=r.paddingLeft.value,{nodeProps:s}=n,d=s==null?void 0:s(e.tmNode.rawNode);return v("div",{class:`${l}-menu-item-group`,role:"group"},v("div",Object.assign({},d,{class:[`${l}-menu-item-group-title`,d==null?void 0:d.class],style:[(d==null?void 0:d.style)||"",a!==void 0?`padding-left: ${a}px;`:""]}),U(e.title),e.extra?v(Ie,null," ",U(e.extra)):null),v("div",null,e.tmNodes.map(c=>de(c,n))))}}}),Se=N({name:"MenuOptionContent",props:{collapsed:Boolean,disabled:Boolean,title:[String,Function],icon:Function,extra:[String,Function],showArrow:Boolean,childActive:Boolean,hover:Boolean,paddingLeft:Number,selected:Boolean,maxIconSize:{type:Number,required:!0},activeIconSize:{type:Number,required:!0},iconMarginRight:{type:Number,required:!0},clsPrefix:{type:String,required:!0},onClick:Function,tmNode:{type:Object,required:!0}},setup(e){const{props:r}=V(W);return{menuProps:r,style:I(()=>{const{paddingLeft:o}=e;return{paddingLeft:o&&`${o}px`}}),iconStyle:I(()=>{const{maxIconSize:o,activeIconSize:n,iconMarginRight:l}=e;return{width:`${o}px`,height:`${o}px`,fontSize:`${n}px`,marginRight:`${l}px`}})}},render(){const{clsPrefix:e,tmNode:r,menuProps:{renderIcon:o,renderLabel:n,renderExtra:l,expandIcon:a}}=this,s=o?o(r.rawNode):U(this.icon);return v("div",{onClick:d=>{var c;(c=this.onClick)===null||c===void 0||c.call(this,d)},role:"none",class:[`${e}-menu-item-content`,{[`${e}-menu-item-content--selected`]:this.selected,[`${e}-menu-item-content--collapsed`]:this.collapsed,[`${e}-menu-item-content--child-active`]:this.childActive,[`${e}-menu-item-content--disabled`]:this.disabled,[`${e}-menu-item-content--hover`]:this.hover}],style:this.style},s&&v("div",{class:`${e}-menu-item-content__icon`,style:this.iconStyle,role:"none"},[s]),v("div",{class:`${e}-menu-item-content-header`,role:"none"},n?n(r.rawNode):U(this.title),this.extra||l?v("span",{class:`${e}-menu-item-content-header__extra`}," ",l?l(r.rawNode):U(this.extra)):null),this.showArrow?v(Fe,{ariaHidden:!0,class:`${e}-menu-item-content__arrow`,clsPrefix:e},{default:()=>a?a(r.rawNode):v(go,null)}):null)}}),Re=Object.assign(Object.assign({},ce),{rawNodes:{type:Array,default:()=>[]},tmNodes:{type:Array,default:()=>[]},tmNode:{type:Object,required:!0},disabled:{type:Boolean,default:!1},icon:Function,onClick:Function}),Ao=N({name:"Submenu",props:Re,setup(e){const r=ae(e),{NMenu:o,NSubmenu:n}=r,{props:l,mergedCollapsedRef:a,mergedThemeRef:s}=o,d=I(()=>{const{disabled:u}=e;return n!=null&&n.mergedDisabledRef.value||l.disabled?!0:u}),c=F(!1);G(ie,{paddingLeftRef:r.paddingLeft,mergedDisabledRef:d}),G(le,null);function x(){const{onClick:u}=e;u&&u()}function _(){d.value||(a.value||o.toggleExpand(e.internalKey),x())}function g(u){c.value=u}return{menuProps:l,mergedTheme:s,doSelect:o.doSelect,inverted:o.invertedRef,isHorizontal:o.isHorizontalRef,mergedClsPrefix:o.mergedClsPrefixRef,maxIconSize:r.maxIconSize,activeIconSize:r.activeIconSize,iconMarginRight:r.iconMarginRight,dropdownPlacement:r.dropdownPlacement,dropdownShow:c,paddingLeft:r.paddingLeft,mergedDisabled:d,mergedValue:o.mergedValueRef,childActive:oe(()=>o.activePathRef.value.includes(e.internalKey)),collapsed:I(()=>l.mode==="horizontal"?!1:a.value?!0:!o.mergedExpandedKeysRef.value.includes(e.internalKey)),dropdownEnabled:I(()=>!d.value&&(l.mode==="horizontal"||a.value)),handlePopoverShowChange:g,handleClick:_}},render(){var e;const{mergedClsPrefix:r,menuProps:{renderIcon:o,renderLabel:n}}=this,l=()=>{const{isHorizontal:s,paddingLeft:d,collapsed:c,mergedDisabled:x,maxIconSize:_,activeIconSize:g,title:u,childActive:y,icon:z,handleClick:T,menuProps:{nodeProps:P},dropdownShow:$,iconMarginRight:q,tmNode:k,mergedClsPrefix:K}=this,L=P==null?void 0:P(k.rawNode);return v("div",Object.assign({},L,{class:[`${K}-menu-item`,L==null?void 0:L.class],role:"menuitem"}),v(Se,{tmNode:k,paddingLeft:d,collapsed:c,disabled:x,iconMarginRight:q,maxIconSize:_,activeIconSize:g,title:u,extra:this.extra,showArrow:!s,childActive:y,clsPrefix:K,icon:z,hover:$,onClick:T}))},a=()=>v(Oe,null,{default:()=>{const{tmNodes:s,collapsed:d}=this;return d?null:v("div",{class:`${r}-submenu-children`,role:"menu"},s.map(c=>de(c,this.menuProps)))}});return this.root?v(_e,Object.assign({size:"large",trigger:"hover"},(e=this.menuProps)===null||e===void 0?void 0:e.dropdownProps,{themeOverrides:this.mergedTheme.peerOverrides.Dropdown,theme:this.mergedTheme.peers.Dropdown,builtinThemeOverrides:{fontSizeLarge:"14px",optionIconSizeLarge:"18px"},value:this.mergedValue,disabled:!this.dropdownEnabled,placement:this.dropdownPlacement,keyField:this.menuProps.keyField,labelField:this.menuProps.labelField,childrenField:this.menuProps.childrenField,onUpdateShow:this.handlePopoverShowChange,options:this.rawNodes,onSelect:this.doSelect,inverted:this.inverted,renderIcon:o,renderLabel:n}),{default:()=>v("div",{class:`${r}-submenu`,role:"menuitem","aria-expanded":!this.collapsed},l(),this.isHorizontal?null:a())}):v("div",{class:`${r}-submenu`,role:"menuitem","aria-expanded":!this.collapsed},l(),a())}}),Ne=Object.assign(Object.assign({},ce),{tmNode:{type:Object,required:!0},disabled:Boolean,icon:Function,onClick:Function}),Ho=N({name:"MenuOption",props:Ne,setup(e){const r=ae(e),{NSubmenu:o,NMenu:n}=r,{props:l,mergedClsPrefixRef:a,mergedCollapsedRef:s}=n,d=o?o.mergedDisabledRef:{value:!1},c=I(()=>d.value||e.disabled);function x(g){const{onClick:u}=e;u&&u(g)}function _(g){c.value||(n.doSelect(e.internalKey,e.tmNode.rawNode),x(g))}return{mergedClsPrefix:a,dropdownPlacement:r.dropdownPlacement,paddingLeft:r.paddingLeft,iconMarginRight:r.iconMarginRight,maxIconSize:r.maxIconSize,activeIconSize:r.activeIconSize,mergedTheme:n.mergedThemeRef,menuProps:l,dropdownEnabled:oe(()=>e.root&&s.value&&l.mode!=="horizontal"&&!c.value),selected:oe(()=>n.mergedValueRef.value===e.internalKey),mergedDisabled:c,handleClick:_}},render(){const{mergedClsPrefix:e,mergedTheme:r,tmNode:o,menuProps:{renderLabel:n,nodeProps:l}}=this,a=l==null?void 0:l(o.rawNode);return v("div",Object.assign({},a,{role:"menuitem",class:[`${e}-menu-item`,a==null?void 0:a.class]}),v(co,{theme:r.peers.Tooltip,themeOverrides:r.peerOverrides.Tooltip,trigger:"hover",placement:this.dropdownPlacement,disabled:!this.dropdownEnabled||this.title===void 0,internalExtraClass:["menu-tooltip"]},{default:()=>n?n(o.rawNode):U(this.title),trigger:()=>v(Se,{tmNode:o,clsPrefix:e,paddingLeft:this.paddingLeft,iconMarginRight:this.iconMarginRight,maxIconSize:this.maxIconSize,activeIconSize:this.activeIconSize,selected:this.selected,title:this.title,extra:this.extra,disabled:this.mergedDisabled,icon:this.icon,onClick:this.handleClick})}))}}),So=N({name:"MenuDivider",setup(){const e=V(W),{mergedClsPrefixRef:r,isHorizontalRef:o}=e;return()=>o.value?null:v("div",{class:`${r.value}-menu-divider`})}}),Ro=ne(He),No=ne(Ne),Po=ne(Re);function Pe(e){return e.type==="divider"||e.type==="render"}function To(e){return e.type==="divider"}function de(e,r){const{rawNode:o}=e,{show:n}=o;if(n===!1)return null;if(Pe(o))return To(o)?v(So,Object.assign({key:e.key},o.props)):null;const{labelField:l}=r,{key:a,level:s,isGroup:d}=e,c=Object.assign(Object.assign({},o),{title:o.title||o[l],extra:o.titleExtra||o.extra,key:a,internalKey:a,level:s,root:s===0,isGroup:d});return e.children?e.isGroup?v(_o,ee(c,Ro,{tmNode:e,tmNodes:e.children,key:a})):v(Ao,ee(c,Po,{key:a,rawNodes:o[r.childrenField],tmNodes:e.children,tmNode:e})):v(Ho,ee(c,No,{key:a,tmNode:e}))}const ge=[C("&::before","background-color: var(--n-item-color-hover);"),m("arrow",`
 color: var(--n-arrow-color-hover);
 `),m("icon",`
 color: var(--n-item-icon-color-hover);
 `),h("menu-item-content-header",`
 color: var(--n-item-text-color-hover);
 `,[C("a",`
 color: var(--n-item-text-color-hover);
 `),m("extra",`
 color: var(--n-item-text-color-hover);
 `)])],xe=[m("icon",`
 color: var(--n-item-icon-color-hover-horizontal);
 `),h("menu-item-content-header",`
 color: var(--n-item-text-color-hover-horizontal);
 `,[C("a",`
 color: var(--n-item-text-color-hover-horizontal);
 `),m("extra",`
 color: var(--n-item-text-color-hover-horizontal);
 `)])],Mo=C([h("menu",`
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
 `)]),J("disabled",[J("selected, child-active",[C("&:focus-within",xe)]),A("selected",[j(null,[m("icon","color: var(--n-item-icon-color-active-hover-horizontal);"),h("menu-item-content-header",`
 color: var(--n-item-text-color-active-hover-horizontal);
 `,[C("a","color: var(--n-item-text-color-active-hover-horizontal);"),m("extra","color: var(--n-item-text-color-active-hover-horizontal);")])])]),A("child-active",[j(null,[m("icon","color: var(--n-item-icon-color-child-active-hover-horizontal);"),h("menu-item-content-header",`
 color: var(--n-item-text-color-child-active-hover-horizontal);
 `,[C("a","color: var(--n-item-text-color-child-active-hover-horizontal);"),m("extra","color: var(--n-item-text-color-child-active-hover-horizontal);")])])]),j("border-bottom: 2px solid var(--n-border-color-horizontal);",xe)]),h("menu-item-content-header",[C("a","color: var(--n-item-text-color-horizontal);")])])]),A("collapsed",[h("menu-item-content",[A("selected",[C("&::before",`
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
 `)]),J("disabled",[J("selected, child-active",[C("&:focus-within",ge)]),A("selected",[j(null,[m("arrow","color: var(--n-arrow-color-active-hover);"),m("icon","color: var(--n-item-icon-color-active-hover);"),h("menu-item-content-header",`
 color: var(--n-item-text-color-active-hover);
 `,[C("a","color: var(--n-item-text-color-active-hover);"),m("extra","color: var(--n-item-text-color-active-hover);")])])]),A("child-active",[j(null,[m("arrow","color: var(--n-arrow-color-child-active-hover);"),m("icon","color: var(--n-item-icon-color-child-active-hover);"),h("menu-item-content-header",`
 color: var(--n-item-text-color-child-active-hover);
 `,[C("a","color: var(--n-item-text-color-child-active-hover);"),m("extra","color: var(--n-item-text-color-child-active-hover);")])])]),A("selected",[j(null,[C("&::before","background-color: var(--n-item-color-active-hover);")])]),j(null,ge)]),m("icon",`
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
 `,[Ee({duration:".2s"})])]),h("menu-item-group",[h("menu-item-group-title",`
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
 `)]);function j(e,r){return[A("hover",e,r),C("&:hover",e,r)]}const $o=Object.assign(Object.assign({},Q.props),{options:{type:Array,default:()=>[]},collapsed:{type:Boolean,default:void 0},collapsedWidth:{type:Number,default:48},iconSize:{type:Number,default:20},collapsedIconSize:{type:Number,default:24},rootIndent:Number,indent:{type:Number,default:32},labelField:{type:String,default:"label"},keyField:{type:String,default:"key"},childrenField:{type:String,default:"children"},disabledField:{type:String,default:"disabled"},defaultExpandAll:Boolean,defaultExpandedKeys:Array,expandedKeys:Array,value:[String,Number],defaultValue:{type:[String,Number],default:null},mode:{type:String,default:"vertical"},watchProps:{type:Array,default:void 0},disabled:Boolean,show:{type:Boolean,defalut:!0},inverted:Boolean,"onUpdate:expandedKeys":[Function,Array],onUpdateExpandedKeys:[Function,Array],onUpdateValue:[Function,Array],"onUpdate:value":[Function,Array],expandIcon:Function,renderIcon:Function,renderLabel:Function,renderExtra:Function,dropdownProps:Object,accordion:Boolean,nodeProps:Function,items:Array,onOpenNamesChange:[Function,Array],onSelect:[Function,Array],onExpandedNamesChange:[Function,Array],expandedNames:Array,defaultExpandedNames:Array,dropdownPlacement:{type:String,default:"bottom"}}),ko=N({name:"Menu",props:$o,setup(e){const{mergedClsPrefixRef:r,inlineThemeDisabled:o}=Ce(e),n=Q("Menu","-menu",Mo,Io,e,r),l=V(no,null),a=I(()=>{var p;const{collapsed:b}=e;if(b!==void 0)return b;if(l){const{collapseModeRef:t,collapsedRef:f}=l;if(t.value==="width")return(p=f.value)!==null&&p!==void 0?p:!1}return!1}),s=I(()=>{const{keyField:p,childrenField:b,disabledField:t}=e;return so(e.items||e.options,{getIgnored(f){return Pe(f)},getChildren(f){return f[b]},getDisabled(f){return f[t]},getKey(f){var S;return(S=f[p])!==null&&S!==void 0?S:f.name}})}),d=I(()=>new Set(s.value.treeNodes.map(p=>p.key))),{watchProps:c}=e,x=F(null);c!=null&&c.includes("defaultValue")?ve(()=>{x.value=e.defaultValue}):x.value=e.defaultValue;const _=me(e,"value"),g=he(_,x),u=F([]),y=()=>{u.value=e.defaultExpandAll?s.value.getNonLeafKeys():e.defaultExpandedNames||e.defaultExpandedKeys||s.value.getPath(g.value,{includeSelf:!1}).keyPath};c!=null&&c.includes("defaultExpandedKeys")?ve(y):y();const z=eo(e,["expandedNames","expandedKeys"]),T=he(z,u),P=I(()=>s.value.treeNodes),$=I(()=>s.value.getPath(g.value).keyPath);G(W,{props:e,mergedCollapsedRef:a,mergedThemeRef:n,mergedValueRef:g,mergedExpandedKeysRef:T,activePathRef:$,mergedClsPrefixRef:r,isHorizontalRef:I(()=>e.mode==="horizontal"),invertedRef:me(e,"inverted"),doSelect:q,toggleExpand:K});function q(p,b){const{"onUpdate:value":t,onUpdateValue:f,onSelect:S}=e;f&&B(f,p,b),t&&B(t,p,b),S&&B(S,p,b),x.value=p}function k(p){const{"onUpdate:expandedKeys":b,onUpdateExpandedKeys:t,onExpandedNamesChange:f,onOpenNamesChange:S}=e;b&&B(b,p),t&&B(t,p),f&&B(f,p),S&&B(S,p),u.value=p}function K(p){const b=Array.from(T.value),t=b.findIndex(f=>f===p);if(~t)b.splice(t,1);else{if(e.accordion&&d.value.has(p)){const f=b.findIndex(S=>d.value.has(S));f>-1&&b.splice(f,1)}b.push(p)}k(b)}const L=p=>{const b=s.value.getPath(p??g.value,{includeSelf:!1}).keyPath;if(!b.length)return;const t=Array.from(T.value),f=new Set([...t,...b]);e.accordion&&d.value.forEach(S=>{f.has(S)&&!b.includes(S)&&f.delete(S)}),k(Array.from(f))},se=I(()=>{const{inverted:p}=e,{common:{cubicBezierEaseInOut:b},self:t}=n.value,{borderRadius:f,borderColorHorizontal:S,fontSize:Te,itemHeight:Me,dividerColor:$e}=t,i={"--n-divider-color":$e,"--n-bezier":b,"--n-font-size":Te,"--n-border-color-horizontal":S,"--n-border-radius":f,"--n-item-height":Me};return p?(i["--n-group-text-color"]=t.groupTextColorInverted,i["--n-color"]=t.colorInverted,i["--n-item-text-color"]=t.itemTextColorInverted,i["--n-item-text-color-hover"]=t.itemTextColorHoverInverted,i["--n-item-text-color-active"]=t.itemTextColorActiveInverted,i["--n-item-text-color-child-active"]=t.itemTextColorChildActiveInverted,i["--n-item-text-color-child-active-hover"]=t.itemTextColorChildActiveInverted,i["--n-item-text-color-active-hover"]=t.itemTextColorActiveHoverInverted,i["--n-item-icon-color"]=t.itemIconColorInverted,i["--n-item-icon-color-hover"]=t.itemIconColorHoverInverted,i["--n-item-icon-color-active"]=t.itemIconColorActiveInverted,i["--n-item-icon-color-active-hover"]=t.itemIconColorActiveHoverInverted,i["--n-item-icon-color-child-active"]=t.itemIconColorChildActiveInverted,i["--n-item-icon-color-child-active-hover"]=t.itemIconColorChildActiveHoverInverted,i["--n-item-icon-color-collapsed"]=t.itemIconColorCollapsedInverted,i["--n-item-text-color-horizontal"]=t.itemTextColorHorizontalInverted,i["--n-item-text-color-hover-horizontal"]=t.itemTextColorHoverHorizontalInverted,i["--n-item-text-color-active-horizontal"]=t.itemTextColorActiveHorizontalInverted,i["--n-item-text-color-child-active-horizontal"]=t.itemTextColorChildActiveHorizontalInverted,i["--n-item-text-color-child-active-hover-horizontal"]=t.itemTextColorChildActiveHoverHorizontalInverted,i["--n-item-text-color-active-hover-horizontal"]=t.itemTextColorActiveHoverHorizontalInverted,i["--n-item-icon-color-horizontal"]=t.itemIconColorHorizontalInverted,i["--n-item-icon-color-hover-horizontal"]=t.itemIconColorHoverHorizontalInverted,i["--n-item-icon-color-active-horizontal"]=t.itemIconColorActiveHorizontalInverted,i["--n-item-icon-color-active-hover-horizontal"]=t.itemIconColorActiveHoverHorizontalInverted,i["--n-item-icon-color-child-active-horizontal"]=t.itemIconColorChildActiveHorizontalInverted,i["--n-item-icon-color-child-active-hover-horizontal"]=t.itemIconColorChildActiveHoverHorizontalInverted,i["--n-arrow-color"]=t.arrowColorInverted,i["--n-arrow-color-hover"]=t.arrowColorHoverInverted,i["--n-arrow-color-active"]=t.arrowColorActiveInverted,i["--n-arrow-color-active-hover"]=t.arrowColorActiveHoverInverted,i["--n-arrow-color-child-active"]=t.arrowColorChildActiveInverted,i["--n-arrow-color-child-active-hover"]=t.arrowColorChildActiveHoverInverted,i["--n-item-color-hover"]=t.itemColorHoverInverted,i["--n-item-color-active"]=t.itemColorActiveInverted,i["--n-item-color-active-hover"]=t.itemColorActiveHoverInverted,i["--n-item-color-active-collapsed"]=t.itemColorActiveCollapsedInverted):(i["--n-group-text-color"]=t.groupTextColor,i["--n-color"]=t.color,i["--n-item-text-color"]=t.itemTextColor,i["--n-item-text-color-hover"]=t.itemTextColorHover,i["--n-item-text-color-active"]=t.itemTextColorActive,i["--n-item-text-color-child-active"]=t.itemTextColorChildActive,i["--n-item-text-color-child-active-hover"]=t.itemTextColorChildActiveHover,i["--n-item-text-color-active-hover"]=t.itemTextColorActiveHover,i["--n-item-icon-color"]=t.itemIconColor,i["--n-item-icon-color-hover"]=t.itemIconColorHover,i["--n-item-icon-color-active"]=t.itemIconColorActive,i["--n-item-icon-color-active-hover"]=t.itemIconColorActiveHover,i["--n-item-icon-color-child-active"]=t.itemIconColorChildActive,i["--n-item-icon-color-child-active-hover"]=t.itemIconColorChildActiveHover,i["--n-item-icon-color-collapsed"]=t.itemIconColorCollapsed,i["--n-item-text-color-horizontal"]=t.itemTextColorHorizontal,i["--n-item-text-color-hover-horizontal"]=t.itemTextColorHoverHorizontal,i["--n-item-text-color-active-horizontal"]=t.itemTextColorActiveHorizontal,i["--n-item-text-color-child-active-horizontal"]=t.itemTextColorChildActiveHorizontal,i["--n-item-text-color-child-active-hover-horizontal"]=t.itemTextColorChildActiveHoverHorizontal,i["--n-item-text-color-active-hover-horizontal"]=t.itemTextColorActiveHoverHorizontal,i["--n-item-icon-color-horizontal"]=t.itemIconColorHorizontal,i["--n-item-icon-color-hover-horizontal"]=t.itemIconColorHoverHorizontal,i["--n-item-icon-color-active-horizontal"]=t.itemIconColorActiveHorizontal,i["--n-item-icon-color-active-hover-horizontal"]=t.itemIconColorActiveHoverHorizontal,i["--n-item-icon-color-child-active-horizontal"]=t.itemIconColorChildActiveHorizontal,i["--n-item-icon-color-child-active-hover-horizontal"]=t.itemIconColorChildActiveHoverHorizontal,i["--n-arrow-color"]=t.arrowColor,i["--n-arrow-color-hover"]=t.arrowColorHover,i["--n-arrow-color-active"]=t.arrowColorActive,i["--n-arrow-color-active-hover"]=t.arrowColorActiveHover,i["--n-arrow-color-child-active"]=t.arrowColorChildActive,i["--n-arrow-color-child-active-hover"]=t.arrowColorChildActiveHover,i["--n-item-color-hover"]=t.itemColorHover,i["--n-item-color-active"]=t.itemColorActive,i["--n-item-color-active-hover"]=t.itemColorActiveHover,i["--n-item-color-active-collapsed"]=t.itemColorActiveCollapsed),i}),D=o?be("menu",I(()=>e.inverted?"a":"b"),se,e):void 0;return{mergedClsPrefix:r,controlledExpandedKeys:z,uncontrolledExpanededKeys:u,mergedExpandedKeys:T,uncontrolledValue:x,mergedValue:g,activePath:$,tmNodes:P,mergedTheme:n,mergedCollapsed:a,cssVars:o?void 0:se,themeClass:D==null?void 0:D.themeClass,onRender:D==null?void 0:D.onRender,showOption:L}},render(){const{mergedClsPrefix:e,mode:r,themeClass:o,onRender:n}=this;return n==null||n(),v("div",{role:r==="horizontal"?"menubar":"menu",class:[`${e}-menu`,o,`${e}-menu--${r}`,this.mergedCollapsed&&`${e}-menu--collapsed`],style:this.cssVars},this.tmNodes.map(l=>de(l,this.$props)))}}),Lo={class:"logo-container",flex:"","items-center":"","justify-center":"","h-12":""},Fo=N({__name:"BaseLogo",props:{isMini:{type:Boolean}},setup(e){const r=e;return(o,n)=>(M(),X("div",Lo,[r.isMini?(M(),E(w(ue),{key:0,class:"h-80%",style:{"font-size":"28px"},gradient:{deg:90,from:"#4372d2",to:"rgb(163,184,225)"}},{default:R(()=>[te(" Jobor ")]),_:1},8,["gradient"])):(M(),E(w(ue),{key:1,class:"h-80%",style:{"font-size":"28px"},gradient:{deg:80,from:"#4372d2",to:"rgb(163,184,225)"}},{default:R(()=>[te(" Jobor ")]),_:1},8,["gradient"]))]))}}),Oo={class:"icon","aria-hidden":"true"},Eo=["xlink:href"],Ko=N({__name:"SvgIcon",props:{path:{default:""}},setup(e){const r=e;return(o,n)=>(M(),E(w(oo),Ke(Be(o.$attrs)),{default:R(()=>[(M(),X("svg",Oo,[O("use",{"xlink:href":r.path},null,8,Eo)]))]),_:1},16))}}),Bo=N({__name:"MenuGroup",setup(e){const r=Ve(),o=De(),n=ze(),{roles:l}=ye(n),a=F(String(r.name));je(()=>r.fullPath,()=>{a.value=String(r.name)});const s=I(()=>o.getRoutes().find(x=>x.name==="root"));function d(x){const _=v(Ko,{path:`#icon-${x}`});return()=>v(_,{})}const c=I(()=>{const x=[...s.value.children];function _(g,u){return g.map(y=>{var K;const{meta:z}=y;y.children&&(y.children=_(y.children));const P=Ue(l.value,z==null?void 0:z.role)&&!(z!=null&&z.hideMenu),$=!!(z!=null&&z.icon),q=z==null?void 0:z.icon;let k;return P&&(y.children?k=(K=y.meta)==null?void 0:K.title:k=v(qe,{to:y.path},{default:()=>{var L;return((L=y.meta)==null?void 0:L.title)??"--"}})),{...y,label:()=>P&&k,show:P,key:y.name,icon:$?d(q):null}})}return _(x)});return(x,_)=>(M(),E(w(ko),{mode:"vertical",value:a.value,indent:20,accordion:!1,"collapsed-width":64,"collapsed-icon-size":22,options:w(c)},null,8,["value","options"]))}}),jo=N({__name:"BaseSider",setup(e){const r=F(!1);function o(n){r.value=n}return(n,l)=>(M(),E(w(uo),{width:"200",bordered:"","native-scrollbar":!1,"show-trigger":"","collapse-mode":"width","collapsed-width":64,"onUpdate:collapsed":o},{default:R(()=>[H(Fo,{"is-mini":r.value},null,8,["is-mini"]),H(Bo)]),_:1}))}}),Vo={class:"avator"},Do=["alt"],Uo={class:"avator-container"},qo=N({__name:"AutoTextAvatar",props:{text:{default:""}},setup(e){const r=e,{text:o}=Ge(r);return(n,l)=>(M(),X("div",Vo,[O("div",{class:"avator-inner",alt:w(o)},[O("div",Uo,[O("span",null,We(w(o)),1)])],8,Do)]))}});const Go=Ae(qo,[["__scopeId","data-v-e59811c3"]]),Wo={flex:"","justify-end":""},Jo=N({__name:"UserSwitchModal",emits:["refresh"],setup(e,{emit:r}){const o=F(!1),{run:n,loading:l}=vo(we.userSwitchById),a=F(null),s=F({userId:void 0}),d={userId:{required:!0,message:"请选择用户",trigger:"blur",type:"number"}};function c(){o.value=!0}async function x(){var _;(_=a.value)==null||_.validate(async g=>{if(g)return;const{result:u}=await n(Number(s.value.userId));u&&(o.value=!1,r("refresh"))})}return(_,g)=>(M(),X(Ie,null,[O("span",{onClick:c},[Je(_.$slots,"default")]),H(w(Qe),{show:o.value,"onUpdate:show":g[2]||(g[2]=u=>o.value=u),title:"切换用户",preset:"card","auto-focus":!1,"w-100":""},{default:R(()=>[H(w(ho),{ref_key:"formRef",ref:a,model:s.value,rules:d},{default:R(()=>[H(w(po),{path:"userId","show-label":!1},{default:R(()=>[H(mo,{value:s.value.userId,"onUpdate:value":g[0]||(g[0]=u=>s.value.userId=u)},null,8,["value"])]),_:1}),O("div",Wo,[H(w(Ze),{type:"primary",loading:w(l),onClick:g[1]||(g[1]=u=>x())},{default:R(()=>[te(" 确定 ")]),_:1},8,["loading"])])]),_:1},8,["model"])]),_:1},8,["show"])],64))}}),Zo={"h-12":"","w-auto":"",flex:"","px-5":"","items-center":"","justify-between":""},Qo=O("span",null,null,-1),Xo=N({__name:"BaseHeader",setup(e){const r=ze(),{nickname:o,isAdmin:n}=ye(r);function l(){window.location.reload()}async function a(){const{result:d}=await we.toLogout();l()}const s=[{label:()=>v(Jo,{onRefresh:l},{default:()=>v("div",{},"切换用户")}),show:n.value},{label:"退出登录",props:{onClick:a}}];return(d,c)=>(M(),E(w(wo),{bordered:""},{default:R(()=>[O("div",Zo,[H(w(fe),{align:"center",justify:"center"},{default:R(()=>[Qo]),_:1}),H(w(fe),{align:"center",justify:"center"},{default:R(()=>[H(w(_e),{"display-directive":"show",options:s},{default:R(()=>[H(Go,{text:w(o)},null,8,["text"])]),_:1})]),_:1})])]),_:1}))}}),Yo=N({__name:"AppMain",setup(e){return(r,o)=>{const n=Ye("RouterView");return M(),E(Xe,null,{fallback:R(()=>[H(w(fo),{size:"large"})]),default:R(()=>[H(n)]),_:1})}}});const et=Ae(Yo,[["__scopeId","data-v-91d1c4f9"]]),pt=N({__name:"PageLayout",setup(e){return(r,o)=>(M(),E(w(pe),{"has-sider":"",position:"absolute"},{default:R(()=>[H(jo),H(w(pe),{"native-scrollbar":!0,style:{overflow:"initial"}},{default:R(()=>[H(Xo),H(w(io),{position:"absolute",style:{"margin-top":"50px"},"content-style":"padding:16px"},{default:R(()=>[H(et)]),_:1})]),_:1})]),_:1}))}});export{pt as default};
