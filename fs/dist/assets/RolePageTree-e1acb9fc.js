import{d as L,h as l,c as Ae,a as Ve,ad as Be,ae as Ie,af as Ne,ag as Ue,k as Me,ah as He,a3 as j,l as K,n as De,ai as Ee,m as $,aj as J,a9 as D,N as je,A as qe,z as E,i as x,e as k,f as U,v as y,t as ee,w as Ke,g as We,u as ne,ak as Ye,p as Ge,al as Je,G as A,am as Xe,B as G,$ as X,H as Z,J as se,an as Ze,P as g,Z as Qe,I as et,L as te,a0 as tt,Y as S,K as T,a1 as rt,a2 as V,aa as lt,a4 as de,a5 as at,F as ot,ab as it,ac as nt}from "./index-a4639770.js";import{_ as st,r as q}from "./RoleSelect.vue_vue_type_script_setup_true_lang-0256ce2c.js";import{s as dt}from "./sysApi-004b66c0.js";import{f as ct,g as ut,h as ft,a as ht}from "./Dropdown-c4cd24d0.js";import{c as mt,N as gt,u as vt,_ as pt}from "./useTableData-1ff1bd86.js";import{V as bt}from "./FocusDetector-f0e5f0c6.js";import{N as xt}from "./InputNumber-4bde0193.js";import{_ as St,a as Ct}from "./EditButton.vue_vue_type_script_setup_true_lang-19ba1fe2.js";import"./Forward-11ddfa2e.js";import"./Popconfirm-faf230d4.js";const Rt=L({name:"Search",render(){return l("svg",{version:"1.1",xmlns:"http://www.w3.org/2000/svg",viewBox:"0 0 512 512",style:"enable-background: new 0 0 512 512"},l("path",{d:`M443.5,420.2L336.7,312.4c20.9-26.2,33.5-59.4,33.5-95.5c0-84.5-68.5-153-153.1-153S64,132.5,64,217s68.5,153,153.1,153
  c36.6,0,70.1-12.8,96.5-34.2l106.1,107.1c3.2,3.4,7.6,5.1,11.9,5.1c4.1,0,8.2-1.5,11.3-4.5C449.5,437.2,449.7,426.8,443.5,420.2z
   M217.1,337.1c-32.1,0-62.3-12.5-85-35.2c-22.7-22.7-35.2-52.9-35.2-84.9c0-32.1,12.5-62.3,35.2-84.9c22.7-22.7,52.9-35.2,85-35.2
  c32.1,0,62.3,12.5,85,35.2c22.7,22.7,35.2,52.9,35.2,84.9c0,32.1-12.5,62.3-35.2,84.9C279.4,324.6,249.2,337.1,217.1,337.1z`}))}}),yt={extraFontSizeSmall:"12px",extraFontSizeMedium:"12px",extraFontSizeLarge:"14px",titleFontSizeSmall:"14px",titleFontSizeMedium:"16px",titleFontSizeLarge:"16px",closeSize:"20px",closeIconSize:"16px",headerHeightSmall:"44px",headerHeightMedium:"44px",headerHeightLarge:"50px"},zt=e=>{const{fontWeight:i,fontSizeLarge:o,fontSizeMedium:t,fontSizeSmall:n,heightLarge:s,heightMedium:c,borderRadius:r,cardColor:h,tableHeaderColor:u,textColor1:d,textColorDisabled:p,textColor2:F,textColor3:z,borderColor:R,hoverColor:b,closeColorHover:f,closeColorPressed:v,closeIconColor:B,closeIconColorHover:I,closeIconColorPressed:a}=e;return Object.assign(Object.assign({},yt),{itemHeightSmall:c,itemHeightMedium:c,itemHeightLarge:s,fontSizeSmall:n,fontSizeMedium:t,fontSizeLarge:o,borderRadius:r,dividerColor:R,borderColor:R,listColor:h,headerColor:Ue(h,u),titleTextColor:d,titleTextColorDisabled:p,extraTextColor:z,extraTextColorDisabled:p,itemTextColor:F,itemTextColorDisabled:p,itemColorPending:b,titleFontWeight:i,closeColorHover:f,closeColorPressed:v,closeIconColor:B,closeIconColorHover:I,closeIconColorPressed:a})},kt=Ae({name:"Transfer",common:Ve,peers:{Checkbox:mt,Scrollbar:Be,Input:Ie,Empty:ct,Button:Ne},self:zt}),Tt=kt,M=Me("n-transfer"),re=L({name:"TransferHeader",props:{size:{type:String,required:!0},source:Boolean,onCheckedAll:Function,onClearAll:Function,title:String},setup(e){const{targetOptionsRef:i,canNotSelectAnythingRef:o,canBeClearedRef:t,allCheckedRef:n,mergedThemeRef:s,disabledRef:c,mergedClsPrefixRef:r,srcOptionsLengthRef:h}=K(M),{localeRef:u}=He("Transfer");return()=>{const{source:d,onClearAll:p,onCheckedAll:F}=e,{value:z}=s,{value:R}=r,{value:b}=u,f=e.size==="large"?"small":"tiny",{title:v}=e;return l("div",{class:`${R}-transfer-list-header`},v&&l("div",{class:`${R}-transfer-list-header__title`},v),d&&l(j,{class:`${R}-transfer-list-header__button`,theme:z.peers.Button,themeOverrides:z.peerOverrides.Button,size:f,tertiary:!0,onClick:n.value?p:F,disabled:o.value||c.value},{default:()=>n.value?b.unselectAll:b.selectAll}),!d&&t.value&&l(j,{class:`${R}-transfer-list-header__button`,theme:z.peers.Button,themeOverrides:z.peerOverrides.Button,size:f,tertiary:!0,onClick:p,disabled:c.value},{default:()=>b.clearAll}),l("div",{class:`${R}-transfer-list-header__extra`},d?b.total(h.value):b.selected(i.value.length)))}}}),le=L({name:"NTransferListItem",props:{source:Boolean,label:{type:String,required:!0},value:{type:[String,Number],required:!0},disabled:Boolean,option:{type:Object,required:!0}},setup(e){const{targetValueSetRef:i,mergedClsPrefixRef:o,mergedThemeRef:t,handleItemCheck:n,renderSourceLabelRef:s,renderTargetLabelRef:c,showSelectedRef:r}=K(M),h=De(()=>i.value.has(e.value));function u(){e.disabled||n(!h.value,e.value)}return{mergedClsPrefix:o,mergedTheme:t,checked:h,showSelected:r,renderSourceLabel:s,renderTargetLabel:c,handleClick:u}},render(){const{disabled:e,mergedTheme:i,mergedClsPrefix:o,label:t,checked:n,source:s,renderSourceLabel:c,renderTargetLabel:r}=this;return l("div",{class:[`${o}-transfer-list-item`,e&&`${o}-transfer-list-item--disabled`,s?`${o}-transfer-list-item--source`:`${o}-transfer-list-item--target`],onClick:s?this.handleClick:void 0},l("div",{class:`${o}-transfer-list-item__background`}),s&&this.showSelected&&l("div",{class:`${o}-transfer-list-item__checkbox`},l(gt,{theme:i.peers.Checkbox,themeOverrides:i.peerOverrides.Checkbox,disabled:e,checked:n})),l("div",{class:`${o}-transfer-list-item__label`,title:ut(t)},s?c?c({option:this.option}):t:r?r({option:this.option}):t),!s&&!e&&l(Ee,{focusable:!1,class:`${o}-transfer-list-item__close`,clsPrefix:o,onClick:this.handleClick}))}}),ae=L({name:"TransferList",props:{virtualScroll:{type:Boolean,required:!0},itemSize:{type:Number,required:!0},options:{type:Array,required:!0},disabled:{type:Boolean,required:!0},source:Boolean},setup(){const{mergedThemeRef:e,mergedClsPrefixRef:i}=K(M),o=$(null),t=$(null);function n(){var r;(r=o.value)===null||r===void 0||r.sync()}function s(){const{value:r}=t;if(!r)return null;const{listElRef:h}=r;return h}function c(){const{value:r}=t;if(!r)return null;const{itemsElRef:h}=r;return h}return{mergedTheme:e,mergedClsPrefix:i,scrollerInstRef:o,vlInstRef:t,syncVLScroller:n,scrollContainer:s,scrollContent:c}},render(){const{mergedTheme:e,options:i}=this;if(i.length===0)return l(ft,{theme:e.peers.Empty,themeOverrides:e.peerOverrides.Empty});const{mergedClsPrefix:o,virtualScroll:t,source:n,disabled:s,syncVLScroller:c}=this;return l(J,{ref:"scrollerInstRef",theme:e.peers.Scrollbar,themeOverrides:e.peerOverrides.Scrollbar,container:t?this.scrollContainer:void 0,content:t?this.scrollContent:void 0},{default:()=>t?l(bt,{ref:"vlInstRef",style:{height:"100%"},class:`${o}-transfer-list-content`,items:this.options,itemSize:this.itemSize,showScrollbar:!1,onResize:c,onScroll:c,keyField:"value"},{default:({item:r})=>{const{source:h,disabled:u}=this;return l(le,{source:h,key:r.value,value:r.value,disabled:r.disabled||u,label:r.label,option:r})}}):l("div",{class:`${o}-transfer-list-content`},i.map(r=>l(le,{source:n,key:r.value,value:r.value,disabled:r.disabled||s,label:r.label,option:r})))})}}),oe=L({name:"TransferFilter",props:{value:String,placeholder:String,disabled:Boolean,onUpdateValue:{type:Function,required:!0}},setup(){const{mergedThemeRef:e,mergedClsPrefixRef:i}=K(M);return{mergedClsPrefix:i,mergedTheme:e}},render(){const{mergedTheme:e,mergedClsPrefix:i}=this;return l("div",{class:`${i}-transfer-filter`},l(D,{value:this.value,onUpdateValue:this.onUpdateValue,disabled:this.disabled,placeholder:this.placeholder,theme:e.peers.Input,themeOverrides:e.peerOverrides.Input,clearable:!0,size:"small"},{"clear-icon-placeholder":()=>l(je,{clsPrefix:i},{default:()=>l(Rt,null)})}))}});function Ft(e){const i=$(e.defaultValue),o=qe(E(e,"value"),i),t=x(()=>{const a=new Map;return(e.options||[]).forEach(m=>a.set(m.value,m)),a}),n=x(()=>new Set(o.value||[])),s=x(()=>{const a=t.value,m=[];return(o.value||[]).forEach(N=>{const _=a.get(N);_&&m.push(_)}),m}),c=$(""),r=$(""),h=x(()=>e.sourceFilterable||!!e.filterable),u=x(()=>{const{showSelected:a,options:m,filter:N}=e;return h.value?m.filter(_=>N(c.value,_,"source")&&(a||!n.value.has(_.value))):a?m:m.filter(_=>!n.value.has(_.value))}),d=x(()=>{if(!e.targetFilterable)return s.value;const{filter:a}=e;return s.value.filter(m=>a(r.value,m,"target"))}),p=x(()=>{const{value:a}=o;return a===null?new Set:new Set(a)}),F=x(()=>{const a=new Set(p.value);return u.value.forEach(m=>{!m.disabled&&!a.has(m.value)&&a.add(m.value)}),a}),z=x(()=>{const a=new Set(p.value);return u.value.forEach(m=>{!m.disabled&&a.has(m.value)&&a.delete(m.value)}),a}),R=x(()=>{const a=new Set(p.value);return d.value.forEach(m=>{m.disabled||a.delete(m.value)}),a}),b=x(()=>u.value.every(a=>a.disabled)),f=x(()=>{if(!u.value.length)return!1;const a=p.value;return u.value.every(m=>m.disabled||a.has(m.value))}),v=x(()=>d.value.some(a=>!a.disabled));function B(a){c.value=a??""}function I(a){r.value=a??""}return{uncontrolledValueRef:i,mergedValueRef:o,targetValueSetRef:n,valueSetForCheckAllRef:F,valueSetForUncheckAllRef:z,valueSetForClearRef:R,filteredTgtOptionsRef:d,filteredSrcOptionsRef:u,targetOptionsRef:s,canNotSelectAnythingRef:b,canBeClearedRef:v,allCheckedRef:f,srcPatternRef:c,tgtPatternRef:r,mergedSrcFilterableRef:h,handleSrcFilterUpdateValue:B,handleTgtFilterUpdateValue:I}}const _t=k("transfer",`
 width: 100%;
 font-size: var(--n-font-size);
 height: 300px;
 display: flex;
 flex-wrap: nowrap;
 word-break: break-word;
`,[U("disabled",[k("transfer-list",[k("transfer-list-header",[y("title",`
 color: var(--n-header-text-color-disabled);
 `),y("extra",`
 color: var(--n-header-extra-text-color-disabled);
 `)])])]),k("transfer-list",`
 flex: 1;
 min-width: 0;
 height: inherit;
 display: flex;
 flex-direction: column;
 background-clip: padding-box;
 position: relative;
 transition: background-color .3s var(--n-bezier);
 background-color: var(--n-list-color);
 `,[U("source",`
 border-top-left-radius: var(--n-border-radius);
 border-bottom-left-radius: var(--n-border-radius);
 `,[y("border","border-right: 1px solid var(--n-divider-color);")]),U("target",`
 border-top-right-radius: var(--n-border-radius);
 border-bottom-right-radius: var(--n-border-radius);
 `,[y("border","border-left: none;")]),y("border",`
 padding: 0 12px;
 border: 1px solid var(--n-border-color);
 transition: border-color .3s var(--n-bezier);
 pointer-events: none;
 border-radius: inherit;
 position: absolute;
 left: 0;
 right: 0;
 top: 0;
 bottom: 0;
 `),k("transfer-list-header",`
 min-height: var(--n-header-height);
 box-sizing: border-box;
 display: flex;
 padding: 12px 12px 10px 12px;
 align-items: center;
 background-clip: padding-box;
 border-radius: inherit;
 border-bottom-left-radius: 0;
 border-bottom-right-radius: 0;
 line-height: 1.5;
 transition:
 border-color .3s var(--n-bezier),
 background-color .3s var(--n-bezier);
 `,[ee("> *:not(:first-child)",`
 margin-left: 8px;
 `),y("title",`
 flex: 1;
 min-width: 0;
 line-height: 1.5;
 font-size: var(--n-header-font-size);
 font-weight: var(--n-header-font-weight);
 transition: color .3s var(--n-bezier);
 color: var(--n-header-text-color);
 `),y("button",`
 position: relative;
 `),y("extra",`
 transition: color .3s var(--n-bezier);
 font-size: var(--n-extra-font-size);
 margin-right: 0;
 white-space: nowrap;
 color: var(--n-header-extra-text-color);
 `)]),k("transfer-list-body",`
 flex-basis: 0;
 flex-grow: 1;
 box-sizing: border-box;
 position: relative;
 display: flex;
 flex-direction: column;
 border-radius: inherit;
 border-top-left-radius: 0;
 border-top-right-radius: 0;
 `,[k("transfer-filter",`
 padding: 4px 12px 8px 12px;
 box-sizing: border-box;
 transition:
 border-color .3s var(--n-bezier),
 background-color .3s var(--n-bezier);
 `),k("transfer-list-flex-container",`
 flex: 1;
 position: relative;
 `,[k("scrollbar",`
 position: absolute;
 left: 0;
 right: 0;
 top: 0;
 bottom: 0;
 height: unset;
 `),k("empty",`
 position: absolute;
 left: 50%;
 top: 50%;
 transform: translateY(-50%) translateX(-50%);
 `),k("transfer-list-content",`
 padding: 0;
 margin: 0;
 position: relative;
 `,[k("transfer-list-item",`
 padding: 0 12px;
 min-height: var(--n-item-height);
 display: flex;
 align-items: center;
 color: var(--n-item-text-color);
 position: relative;
 transition: color .3s var(--n-bezier);
 `,[y("background",`
 position: absolute;
 left: 4px;
 right: 4px;
 top: 0;
 bottom: 0;
 border-radius: var(--n-border-radius);
 transition: background-color .3s var(--n-bezier);
 `),y("checkbox",`
 position: relative;
 margin-right: 8px;
 `),y("close",`
 opacity: 0;
 pointer-events: none;
 position: relative;
 transition:
 opacity .3s var(--n-bezier),
 background-color .3s var(--n-bezier),
 color .3s var(--n-bezier);
 `),y("label",`
 position: relative;
 min-width: 0;
 flex-grow: 1;
 `),U("source","cursor: pointer;"),U("disabled",`
 cursor: not-allowed;
 color: var(--n-item-text-color-disabled);
 `),Ke("disabled",[ee("&:hover",[y("background","background-color: var(--n-item-color-pending);"),y("close",`
 opacity: 1;
 pointer-events: all;
 `)])])])])])])])]),wt=Object.assign(Object.assign({},ne.props),{value:Array,defaultValue:{type:Array,default:null},options:{type:Array,default:()=>[]},disabled:{type:Boolean,default:void 0},virtualScroll:Boolean,sourceTitle:String,targetTitle:String,filterable:{type:Boolean,default:void 0},sourceFilterable:Boolean,targetFilterable:Boolean,showSelected:{type:Boolean,default:!0},sourceFilterPlaceholder:String,targetFilterPlaceholder:String,filter:{type:Function,default:(e,i)=>e?~(""+i.label).toLowerCase().indexOf((""+e).toLowerCase()):!0},size:String,renderSourceLabel:Function,renderTargetLabel:Function,renderSourceList:Function,renderTargetList:Function,"onUpdate:value":[Function,Array],onUpdateValue:[Function,Array],onChange:[Function,Array]}),Lt=L({name:"Transfer",props:wt,setup(e){const{mergedClsPrefixRef:i}=We(e),o=ne("Transfer","-transfer",_t,Tt,e,i),t=Ye(e),{mergedSizeRef:n,mergedDisabledRef:s}=t,c=x(()=>{const{value:C}=n,{self:{[A("itemHeight",C)]:w}}=o.value;return Xe(w)}),{uncontrolledValueRef:r,mergedValueRef:h,targetValueSetRef:u,valueSetForCheckAllRef:d,valueSetForUncheckAllRef:p,valueSetForClearRef:F,filteredTgtOptionsRef:z,filteredSrcOptionsRef:R,targetOptionsRef:b,canNotSelectAnythingRef:f,canBeClearedRef:v,allCheckedRef:B,srcPatternRef:I,tgtPatternRef:a,mergedSrcFilterableRef:m,handleSrcFilterUpdateValue:N,handleTgtFilterUpdateValue:_}=Ft(e);function O(C){const{onUpdateValue:w,"onUpdate:value":P,onChange:H}=e,{nTriggerFormInput:W,nTriggerFormChange:Y}=t;w&&G(w,C),P&&G(P,C),H&&G(H,C),r.value=C,W(),Y()}function ce(){O([...d.value])}function ue(){O([...p.value])}function fe(){O([...F.value])}function Q(C,w){O(C?(h.value||[]).concat(w):(h.value||[]).filter(P=>P!==w))}function he(C){O(C)}return Ge(M,{targetValueSetRef:u,mergedClsPrefixRef:i,disabledRef:s,mergedThemeRef:o,targetOptionsRef:b,canNotSelectAnythingRef:f,canBeClearedRef:v,allCheckedRef:B,srcOptionsLengthRef:x(()=>e.options.length),handleItemCheck:Q,renderSourceLabelRef:E(e,"renderSourceLabel"),renderTargetLabelRef:E(e,"renderTargetLabel"),showSelectedRef:E(e,"showSelected")}),{mergedClsPrefix:i,mergedDisabled:s,itemSize:c,isMounted:Je(),mergedTheme:o,filteredSrcOpts:R,filteredTgtOpts:z,srcPattern:I,tgtPattern:a,mergedSize:n,mergedSrcFilterable:m,handleSrcFilterUpdateValue:N,handleTgtFilterUpdateValue:_,handleSourceCheckAll:ce,handleSourceUncheckAll:ue,handleTargetClearAll:fe,handleItemCheck:Q,handleChecked:he,cssVars:x(()=>{const{value:C}=n,{common:{cubicBezierEaseInOut:w},self:{borderRadius:P,borderColor:H,listColor:W,titleTextColor:Y,titleTextColorDisabled:me,extraTextColor:ge,itemTextColor:ve,itemColorPending:pe,itemTextColorDisabled:be,titleFontWeight:xe,closeColorHover:Se,closeColorPressed:Ce,closeIconColor:Re,closeIconColorHover:ye,closeIconColorPressed:ze,closeIconSize:ke,closeSize:Te,dividerColor:Fe,extraTextColorDisabled:_e,[A("extraFontSize",C)]:we,[A("fontSize",C)]:Le,[A("titleFontSize",C)]:$e,[A("itemHeight",C)]:Oe,[A("headerHeight",C)]:Pe}}=o.value;return{"--n-bezier":w,"--n-border-color":H,"--n-border-radius":P,"--n-extra-font-size":we,"--n-font-size":Le,"--n-header-font-size":$e,"--n-header-extra-text-color":ge,"--n-header-extra-text-color-disabled":_e,"--n-header-font-weight":xe,"--n-header-text-color":Y,"--n-header-text-color-disabled":me,"--n-item-color-pending":pe,"--n-item-height":Oe,"--n-item-text-color":ve,"--n-item-text-color-disabled":be,"--n-list-color":W,"--n-header-height":Pe,"--n-close-size":Te,"--n-close-icon-size":ke,"--n-close-color-hover":Se,"--n-close-color-pressed":Ce,"--n-close-icon-color":Re,"--n-close-icon-color-hover":ye,"--n-close-icon-color-pressed":ze,"--n-divider-color":Fe}})}},render(){const{mergedClsPrefix:e,renderSourceList:i,renderTargetList:o,mergedTheme:t,mergedSrcFilterable:n,targetFilterable:s}=this;return l("div",{class:[`${e}-transfer`,this.mergedDisabled&&`${e}-transfer--disabled`],style:this.cssVars},l("div",{class:`${e}-transfer-list ${e}-transfer-list--source`},l(re,{source:!0,title:this.sourceTitle,onCheckedAll:this.handleSourceCheckAll,onClearAll:this.handleSourceUncheckAll,size:this.mergedSize}),l("div",{class:`${e}-transfer-list-body`},n?l(oe,{onUpdateValue:this.handleSrcFilterUpdateValue,value:this.srcPattern,disabled:this.mergedDisabled,placeholder:this.sourceFilterPlaceholder}):null,l("div",{class:`${e}-transfer-list-flex-container`},i?l(J,{theme:t.peers.Scrollbar,themeOverrides:t.peerOverrides.Scrollbar},{default:()=>i({onCheck:this.handleChecked,checkedOptions:this.filteredTgtOpts,pattern:this.srcPattern})}):l(ae,{source:!0,options:this.filteredSrcOpts,disabled:this.mergedDisabled,virtualScroll:this.virtualScroll,itemSize:this.itemSize}))),l("div",{class:`${e}-transfer-list__border`})),l("div",{class:`${e}-transfer-list ${e}-transfer-list--target`},l(re,{onClearAll:this.handleTargetClearAll,size:this.mergedSize,title:this.targetTitle}),l("div",{class:`${e}-transfer-list-body`},s?l(oe,{onUpdateValue:this.handleTgtFilterUpdateValue,value:this.tgtPattern,disabled:this.mergedDisabled,placeholder:this.sourceFilterPlaceholder}):null,l("div",{class:`${e}-transfer-list-flex-container`},o?l(J,{theme:t.peers.Scrollbar,themeOverrides:t.peerOverrides.Scrollbar},{default:()=>o({onCheck:this.handleChecked,checkedOptions:this.filteredTgtOpts,pattern:this.tgtPattern})}):l(ae,{options:this.filteredTgtOpts,disabled:this.mergedDisabled,virtualScroll:this.virtualScroll,itemSize:this.itemSize}))),l("div",{class:`${e}-transfer-list__border`})))}}),$t=L({__name:"ApiTransfer",setup(e){const{data:i,loading:o}=X(dt.getAllApi,{auto:!0}),t=x(()=>{var c,r,h;return(c=i.value)!=null&&c.list?(h=(r=i.value)==null?void 0:r.list)==null?void 0:h.map(u=>({...u,label:u.title,value:u.id})):[]}),n=function({option:s}){const{title:c,method:r,path:h,name:u}=s;return l(ht,{},{default:()=>`${r} ${h}`,trigger:()=>`${c} ${u}`})};return(s,c)=>(Z(),se(g(Lt),Ze({loading:g(o),options:g(t)},s.$attrs,{"virtual-scroll":"","source-filterable":"","target-filterable":"","render-source-label":n,"render-target-label":n}),null,16,["loading","options","render-source-label","render-target-label"]))}}),Ot={flex:"","justify-end":""},ie=L({__name:"RoleModal",props:{record:null},emits:["refresh"],setup(e,{emit:i}){const o=e,{record:t}=Qe(o),n=$(!1),s=$(null),{run:c,loading:r}=X(q.createRole),{run:h,loading:u}=X(q.editRole),d=$({name:"",parent_id:0,title:"",description:"",sort:100}),p={name:{required:!0,message:"请输入角色名",trigger:"blur"},code:{required:!0,message:"请输入code",trigger:"blur"}};function F(){n.value=!0,t!=null&&t.value&&(d.value={...t.value,api_ids:t.value.apis.map(b=>b.id)})}const z=x(()=>!!(t!=null&&t.value));function R(){var b;(b=s.value)==null||b.validate(async f=>{if(f)return;const{result:v}=t!=null&&t.value?await h(t==null?void 0:t.value.id,d.value):await c(d.value);v&&(n.value=!1,i("refresh"))})}return(b,f)=>(Z(),et(ot,null,[te("span",{onClick:F},[tt(b.$slots,"default")]),S(g(at),{show:n.value,"onUpdate:show":f[8]||(f[8]=v=>n.value=v),title:"角色",preset:"card","auto-focus":!1,"w-200":""},{default:T(()=>[S(g(rt),{ref_key:"formRef",ref:s,model:d.value,rules:p},{default:T(()=>[S(g(V),{path:"title",label:"角色名"},{default:T(()=>[S(g(D),{value:d.value.title,"onUpdate:value":f[0]||(f[0]=v=>d.value.title=v),placeholder:"输入角色名",onKeydown:f[1]||(f[1]=lt(it(()=>{},["prevent"]),["enter"]))},null,8,["value"])]),_:1}),S(g(V),{path:"name",label:"标识"},{default:T(()=>[S(g(D),{value:d.value.name,"onUpdate:value":f[2]||(f[2]=v=>d.value.name=v),disabled:g(z),placeholder:"请输入name"},null,8,["value","disabled"])]),_:1}),S(g(V),{path:"description",label:"描述"},{default:T(()=>[S(g(D),{value:d.value.description,"onUpdate:value":f[3]||(f[3]=v=>d.value.description=v),type:"textarea",placeholder:"请输入描述"},null,8,["value"])]),_:1}),S(g(V),{path:"parent_id",label:"父节点"},{default:T(()=>[S(st,{value:d.value.parent_id,"onUpdate:value":f[4]||(f[4]=v=>d.value.parent_id=v),clearable:"",placeholder:"请输入父节点"},null,8,["value"])]),_:1}),S(g(V),{path:"sort",label:"排序"},{default:T(()=>[S(g(xt),{value:d.value.sort,"onUpdate:value":f[5]||(f[5]=v=>d.value.sort=v),type:"textarea",placeholder:"请输入排序"},null,8,["value"])]),_:1}),S(g(V),{path:"api_ids",label:"权限"},{default:T(()=>[S($t,{value:d.value.api_ids,"onUpdate:value":f[6]||(f[6]=v=>d.value.api_ids=v)},null,8,["value"])]),_:1}),te("div",Ot,[S(g(j),{type:"primary",loading:g(r)||g(u),onClick:f[7]||(f[7]=v=>R())},{default:T(()=>[de(" 确定 ")]),_:1},8,["loading"])])]),_:1},8,["model"])]),_:1},8,["show"])],64))}}),Et=L({__name:"RolePageTree",setup(e){const{dataSource:i,loading:o,refresh:t,tableFilter:n,total:s,pagination:c}=vt(q.getRoleTreeList);async function r(u){const{result:d}=await q.delRole(u);d&&t()}const h=[{title:"序号",key:"id",width:"205",resizable:!0},{title:"角色名",key:"title",resizable:!0},{title:"标识",key:"name"},{title:"排序",key:"sort"},{title:"更新时间",key:"updated_at",width:"185"},{title:"操作",key:"action",width:"170",fixed:"right",render(u){return[l(ie,{record:u,onRefresh:t},{default:()=>l(St)}),l(Ct,{confirmText:"是否删除该角色",delFunction:()=>r(u.id)})]}}];return(u,d)=>(Z(),se(pt,{"table-filter":g(n),"onUpdate:tableFilter":d[0]||(d[0]=p=>nt(n)?n.value=p:null),data:g(i),loading:g(o),columns:h,total:g(s),pagination:g(c),"search-key":"name",tree:!0,"row-key":p=>`${p.id}-${p.name}`,"search-place-holder":"请输入角色名",onRefresh:g(t)},{header:T(()=>[S(ie,{onRefresh:g(t)},{default:T(()=>[S(g(j),{type:"primary"},{default:T(()=>[de(" 新增角色 ")]),_:1})]),_:1},8,["onRefresh"])]),_:1},8,["table-filter","data","loading","total","pagination","row-key","onRefresh"]))}});export{Et as default};
