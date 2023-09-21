import{d as L,h as l,c as Ve,a as Ae,ab as Be,ac as Ne,ad as Ie,k as Ue,W as q,l as K,n as Me,ae as De,m as $,af as G,N as He,z as H,i as x,e as T,f as U,v as y,t as ee,w as Ee,g as qe,u as ne,ag as je,p as Ke,ah as We,a1 as V,ai as Xe,A as Y,B as Q,D as se,a9 as Ye,J as p,T as Ge,C as Je,G as te,V as Qe,S,E as k,a4 as Ze,X as de,Y as et,F as tt,a8 as rt,aa as lt}from"./index-969e67fe.js";import{u as J}from"./useRequest-c7c4a46d.js";import{_ as at,r as j}from"./RoleSelect.vue_vue_type_script_setup_true_lang-1242dad7.js";import{s as ot}from"./sysApi-7c205671.js";import{b as it,u as nt}from"./Icon-32b66fa9.js";import{e as st,g as dt,f as ct,a as ut}from"./Dropdown-cd7deeac.js";import{c as ft,N as ht,u as mt,_ as pt}from"./useTableData-e876992d.js";import{V as gt}from"./FocusDetector-1e4512d4.js";import{i as vt,N as E}from"./Input-0a4232e1.js";import{N as bt}from"./InputNumber-b997c12f.js";import{N as xt,a as A}from"./FormItem-6aa25f6b.js";import{_ as St,a as Ct}from"./EditButton.vue_vue_type_script_setup_true_lang-5cb47f39.js";import{_ as Rt}from"./TooltipText.vue_vue_type_script_setup_true_lang-dcb6540f.js";import"./get-a0781824.js";import"./Forward-eebc3cd7.js";const yt=L({name:"Search",render(){return l("svg",{version:"1.1",xmlns:"http://www.w3.org/2000/svg",viewBox:"0 0 512 512",style:"enable-background: new 0 0 512 512"},l("path",{d:`M443.5,420.2L336.7,312.4c20.9-26.2,33.5-59.4,33.5-95.5c0-84.5-68.5-153-153.1-153S64,132.5,64,217s68.5,153,153.1,153
  c36.6,0,70.1-12.8,96.5-34.2l106.1,107.1c3.2,3.4,7.6,5.1,11.9,5.1c4.1,0,8.2-1.5,11.3-4.5C449.5,437.2,449.7,426.8,443.5,420.2z
   M217.1,337.1c-32.1,0-62.3-12.5-85-35.2c-22.7-22.7-35.2-52.9-35.2-84.9c0-32.1,12.5-62.3,35.2-84.9c22.7-22.7,52.9-35.2,85-35.2
  c32.1,0,62.3,12.5,85,35.2c22.7,22.7,35.2,52.9,35.2,84.9c0,32.1-12.5,62.3-35.2,84.9C279.4,324.6,249.2,337.1,217.1,337.1z`}))}}),zt={extraFontSizeSmall:"12px",extraFontSizeMedium:"12px",extraFontSizeLarge:"14px",titleFontSizeSmall:"14px",titleFontSizeMedium:"16px",titleFontSizeLarge:"16px",closeSize:"20px",closeIconSize:"16px",headerHeightSmall:"44px",headerHeightMedium:"44px",headerHeightLarge:"50px"},Tt=e=>{const{fontWeight:i,fontSizeLarge:o,fontSizeMedium:t,fontSizeSmall:n,heightLarge:s,heightMedium:u,borderRadius:r,cardColor:h,tableHeaderColor:d,textColor1:c,textColorDisabled:v,textColor2:F,textColor3:z,borderColor:R,hoverColor:b,closeColorHover:f,closeColorPressed:g,closeIconColor:B,closeIconColorHover:N,closeIconColorPressed:a}=e;return Object.assign(Object.assign({},zt),{itemHeightSmall:u,itemHeightMedium:u,itemHeightLarge:s,fontSizeSmall:n,fontSizeMedium:t,fontSizeLarge:o,borderRadius:r,dividerColor:R,borderColor:R,listColor:h,headerColor:Ie(h,d),titleTextColor:c,titleTextColorDisabled:v,extraTextColor:z,extraTextColorDisabled:v,itemTextColor:F,itemTextColorDisabled:v,itemColorPending:b,titleFontWeight:i,closeColorHover:f,closeColorPressed:g,closeIconColor:B,closeIconColorHover:N,closeIconColorPressed:a})},kt=Ve({name:"Transfer",common:Ae,peers:{Checkbox:ft,Scrollbar:Be,Input:vt,Empty:st,Button:Ne},self:Tt}),Ft=kt,M=Ue("n-transfer"),re=L({name:"TransferHeader",props:{size:{type:String,required:!0},source:Boolean,onCheckedAll:Function,onClearAll:Function,title:String},setup(e){const{targetOptionsRef:i,canNotSelectAnythingRef:o,canBeClearedRef:t,allCheckedRef:n,mergedThemeRef:s,disabledRef:u,mergedClsPrefixRef:r,srcOptionsLengthRef:h}=K(M),{localeRef:d}=it("Transfer");return()=>{const{source:c,onClearAll:v,onCheckedAll:F}=e,{value:z}=s,{value:R}=r,{value:b}=d,f=e.size==="large"?"small":"tiny",{title:g}=e;return l("div",{class:`${R}-transfer-list-header`},g&&l("div",{class:`${R}-transfer-list-header__title`},g),c&&l(q,{class:`${R}-transfer-list-header__button`,theme:z.peers.Button,themeOverrides:z.peerOverrides.Button,size:f,tertiary:!0,onClick:n.value?v:F,disabled:o.value||u.value},{default:()=>n.value?b.unselectAll:b.selectAll}),!c&&t.value&&l(q,{class:`${R}-transfer-list-header__button`,theme:z.peers.Button,themeOverrides:z.peerOverrides.Button,size:f,tertiary:!0,onClick:v,disabled:u.value},{default:()=>b.clearAll}),l("div",{class:`${R}-transfer-list-header__extra`},c?b.total(h.value):b.selected(i.value.length)))}}}),le=L({name:"NTransferListItem",props:{source:Boolean,label:{type:String,required:!0},value:{type:[String,Number],required:!0},disabled:Boolean,option:{type:Object,required:!0}},setup(e){const{targetValueSetRef:i,mergedClsPrefixRef:o,mergedThemeRef:t,handleItemCheck:n,renderSourceLabelRef:s,renderTargetLabelRef:u,showSelectedRef:r}=K(M),h=Me(()=>i.value.has(e.value));function d(){e.disabled||n(!h.value,e.value)}return{mergedClsPrefix:o,mergedTheme:t,checked:h,showSelected:r,renderSourceLabel:s,renderTargetLabel:u,handleClick:d}},render(){const{disabled:e,mergedTheme:i,mergedClsPrefix:o,label:t,checked:n,source:s,renderSourceLabel:u,renderTargetLabel:r}=this;return l("div",{class:[`${o}-transfer-list-item`,e&&`${o}-transfer-list-item--disabled`,s?`${o}-transfer-list-item--source`:`${o}-transfer-list-item--target`],onClick:s?this.handleClick:void 0},l("div",{class:`${o}-transfer-list-item__background`}),s&&this.showSelected&&l("div",{class:`${o}-transfer-list-item__checkbox`},l(ht,{theme:i.peers.Checkbox,themeOverrides:i.peerOverrides.Checkbox,disabled:e,checked:n})),l("div",{class:`${o}-transfer-list-item__label`,title:dt(t)},s?u?u({option:this.option}):t:r?r({option:this.option}):t),!s&&!e&&l(De,{focusable:!1,class:`${o}-transfer-list-item__close`,clsPrefix:o,onClick:this.handleClick}))}}),ae=L({name:"TransferList",props:{virtualScroll:{type:Boolean,required:!0},itemSize:{type:Number,required:!0},options:{type:Array,required:!0},disabled:{type:Boolean,required:!0},source:Boolean},setup(){const{mergedThemeRef:e,mergedClsPrefixRef:i}=K(M),o=$(null),t=$(null);function n(){var r;(r=o.value)===null||r===void 0||r.sync()}function s(){const{value:r}=t;if(!r)return null;const{listElRef:h}=r;return h}function u(){const{value:r}=t;if(!r)return null;const{itemsElRef:h}=r;return h}return{mergedTheme:e,mergedClsPrefix:i,scrollerInstRef:o,vlInstRef:t,syncVLScroller:n,scrollContainer:s,scrollContent:u}},render(){const{mergedTheme:e,options:i}=this;if(i.length===0)return l(ct,{theme:e.peers.Empty,themeOverrides:e.peerOverrides.Empty});const{mergedClsPrefix:o,virtualScroll:t,source:n,disabled:s,syncVLScroller:u}=this;return l(G,{ref:"scrollerInstRef",theme:e.peers.Scrollbar,themeOverrides:e.peerOverrides.Scrollbar,container:t?this.scrollContainer:void 0,content:t?this.scrollContent:void 0},{default:()=>t?l(gt,{ref:"vlInstRef",style:{height:"100%"},class:`${o}-transfer-list-content`,items:this.options,itemSize:this.itemSize,showScrollbar:!1,onResize:u,onScroll:u,keyField:"value"},{default:({item:r})=>{const{source:h,disabled:d}=this;return l(le,{source:h,key:r.value,value:r.value,disabled:r.disabled||d,label:r.label,option:r})}}):l("div",{class:`${o}-transfer-list-content`},i.map(r=>l(le,{source:n,key:r.value,value:r.value,disabled:r.disabled||s,label:r.label,option:r})))})}}),oe=L({name:"TransferFilter",props:{value:String,placeholder:String,disabled:Boolean,onUpdateValue:{type:Function,required:!0}},setup(){const{mergedThemeRef:e,mergedClsPrefixRef:i}=K(M);return{mergedClsPrefix:i,mergedTheme:e}},render(){const{mergedTheme:e,mergedClsPrefix:i}=this;return l("div",{class:`${i}-transfer-filter`},l(E,{value:this.value,onUpdateValue:this.onUpdateValue,disabled:this.disabled,placeholder:this.placeholder,theme:e.peers.Input,themeOverrides:e.peerOverrides.Input,clearable:!0,size:"small"},{"clear-icon-placeholder":()=>l(He,{clsPrefix:i},{default:()=>l(yt,null)})}))}});function _t(e){const i=$(e.defaultValue),o=nt(H(e,"value"),i),t=x(()=>{const a=new Map;return(e.options||[]).forEach(m=>a.set(m.value,m)),a}),n=x(()=>new Set(o.value||[])),s=x(()=>{const a=t.value,m=[];return(o.value||[]).forEach(I=>{const _=a.get(I);_&&m.push(_)}),m}),u=$(""),r=$(""),h=x(()=>e.sourceFilterable||!!e.filterable),d=x(()=>{const{showSelected:a,options:m,filter:I}=e;return h.value?m.filter(_=>I(u.value,_,"source")&&(a||!n.value.has(_.value))):a?m:m.filter(_=>!n.value.has(_.value))}),c=x(()=>{if(!e.targetFilterable)return s.value;const{filter:a}=e;return s.value.filter(m=>a(r.value,m,"target"))}),v=x(()=>{const{value:a}=o;return a===null?new Set:new Set(a)}),F=x(()=>{const a=new Set(v.value);return d.value.forEach(m=>{!m.disabled&&!a.has(m.value)&&a.add(m.value)}),a}),z=x(()=>{const a=new Set(v.value);return d.value.forEach(m=>{!m.disabled&&a.has(m.value)&&a.delete(m.value)}),a}),R=x(()=>{const a=new Set(v.value);return c.value.forEach(m=>{m.disabled||a.delete(m.value)}),a}),b=x(()=>d.value.every(a=>a.disabled)),f=x(()=>{if(!d.value.length)return!1;const a=v.value;return d.value.every(m=>m.disabled||a.has(m.value))}),g=x(()=>c.value.some(a=>!a.disabled));function B(a){u.value=a??""}function N(a){r.value=a??""}return{uncontrolledValueRef:i,mergedValueRef:o,targetValueSetRef:n,valueSetForCheckAllRef:F,valueSetForUncheckAllRef:z,valueSetForClearRef:R,filteredTgtOptionsRef:c,filteredSrcOptionsRef:d,targetOptionsRef:s,canNotSelectAnythingRef:b,canBeClearedRef:g,allCheckedRef:f,srcPatternRef:u,tgtPatternRef:r,mergedSrcFilterableRef:h,handleSrcFilterUpdateValue:B,handleTgtFilterUpdateValue:N}}const wt=T("transfer",`
 width: 100%;
 font-size: var(--n-font-size);
 height: 300px;
 display: flex;
 flex-wrap: nowrap;
 word-break: break-word;
`,[U("disabled",[T("transfer-list",[T("transfer-list-header",[y("title",`
 color: var(--n-header-text-color-disabled);
 `),y("extra",`
 color: var(--n-header-extra-text-color-disabled);
 `)])])]),T("transfer-list",`
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
 `),T("transfer-list-header",`
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
 `)]),T("transfer-list-body",`
 flex-basis: 0;
 flex-grow: 1;
 box-sizing: border-box;
 position: relative;
 display: flex;
 flex-direction: column;
 border-radius: inherit;
 border-top-left-radius: 0;
 border-top-right-radius: 0;
 `,[T("transfer-filter",`
 padding: 4px 12px 8px 12px;
 box-sizing: border-box;
 transition:
 border-color .3s var(--n-bezier),
 background-color .3s var(--n-bezier);
 `),T("transfer-list-flex-container",`
 flex: 1;
 position: relative;
 `,[T("scrollbar",`
 position: absolute;
 left: 0;
 right: 0;
 top: 0;
 bottom: 0;
 height: unset;
 `),T("empty",`
 position: absolute;
 left: 50%;
 top: 50%;
 transform: translateY(-50%) translateX(-50%);
 `),T("transfer-list-content",`
 padding: 0;
 margin: 0;
 position: relative;
 `,[T("transfer-list-item",`
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
 `),Ee("disabled",[ee("&:hover",[y("background","background-color: var(--n-item-color-pending);"),y("close",`
 opacity: 1;
 pointer-events: all;
 `)])])])])])])])]),Lt=Object.assign(Object.assign({},ne.props),{value:Array,defaultValue:{type:Array,default:null},options:{type:Array,default:()=>[]},disabled:{type:Boolean,default:void 0},virtualScroll:Boolean,sourceTitle:String,targetTitle:String,filterable:{type:Boolean,default:void 0},sourceFilterable:Boolean,targetFilterable:Boolean,showSelected:{type:Boolean,default:!0},sourceFilterPlaceholder:String,targetFilterPlaceholder:String,filter:{type:Function,default:(e,i)=>e?~(""+i.label).toLowerCase().indexOf((""+e).toLowerCase()):!0},size:String,renderSourceLabel:Function,renderTargetLabel:Function,renderSourceList:Function,renderTargetList:Function,"onUpdate:value":[Function,Array],onUpdateValue:[Function,Array],onChange:[Function,Array]}),$t=L({name:"Transfer",props:Lt,setup(e){const{mergedClsPrefixRef:i}=qe(e),o=ne("Transfer","-transfer",wt,Ft,e,i),t=je(e),{mergedSizeRef:n,mergedDisabledRef:s}=t,u=x(()=>{const{value:C}=n,{self:{[V("itemHeight",C)]:w}}=o.value;return Xe(w)}),{uncontrolledValueRef:r,mergedValueRef:h,targetValueSetRef:d,valueSetForCheckAllRef:c,valueSetForUncheckAllRef:v,valueSetForClearRef:F,filteredTgtOptionsRef:z,filteredSrcOptionsRef:R,targetOptionsRef:b,canNotSelectAnythingRef:f,canBeClearedRef:g,allCheckedRef:B,srcPatternRef:N,tgtPatternRef:a,mergedSrcFilterableRef:m,handleSrcFilterUpdateValue:I,handleTgtFilterUpdateValue:_}=_t(e);function O(C){const{onUpdateValue:w,"onUpdate:value":P,onChange:D}=e,{nTriggerFormInput:W,nTriggerFormChange:X}=t;w&&Y(w,C),P&&Y(P,C),D&&Y(D,C),r.value=C,W(),X()}function ce(){O([...c.value])}function ue(){O([...v.value])}function fe(){O([...F.value])}function Z(C,w){O(C?(h.value||[]).concat(w):(h.value||[]).filter(P=>P!==w))}function he(C){O(C)}return Ke(M,{targetValueSetRef:d,mergedClsPrefixRef:i,disabledRef:s,mergedThemeRef:o,targetOptionsRef:b,canNotSelectAnythingRef:f,canBeClearedRef:g,allCheckedRef:B,srcOptionsLengthRef:x(()=>e.options.length),handleItemCheck:Z,renderSourceLabelRef:H(e,"renderSourceLabel"),renderTargetLabelRef:H(e,"renderTargetLabel"),showSelectedRef:H(e,"showSelected")}),{mergedClsPrefix:i,mergedDisabled:s,itemSize:u,isMounted:We(),mergedTheme:o,filteredSrcOpts:R,filteredTgtOpts:z,srcPattern:N,tgtPattern:a,mergedSize:n,mergedSrcFilterable:m,handleSrcFilterUpdateValue:I,handleTgtFilterUpdateValue:_,handleSourceCheckAll:ce,handleSourceUncheckAll:ue,handleTargetClearAll:fe,handleItemCheck:Z,handleChecked:he,cssVars:x(()=>{const{value:C}=n,{common:{cubicBezierEaseInOut:w},self:{borderRadius:P,borderColor:D,listColor:W,titleTextColor:X,titleTextColorDisabled:me,extraTextColor:pe,itemTextColor:ge,itemColorPending:ve,itemTextColorDisabled:be,titleFontWeight:xe,closeColorHover:Se,closeColorPressed:Ce,closeIconColor:Re,closeIconColorHover:ye,closeIconColorPressed:ze,closeIconSize:Te,closeSize:ke,dividerColor:Fe,extraTextColorDisabled:_e,[V("extraFontSize",C)]:we,[V("fontSize",C)]:Le,[V("titleFontSize",C)]:$e,[V("itemHeight",C)]:Oe,[V("headerHeight",C)]:Pe}}=o.value;return{"--n-bezier":w,"--n-border-color":D,"--n-border-radius":P,"--n-extra-font-size":we,"--n-font-size":Le,"--n-header-font-size":$e,"--n-header-extra-text-color":pe,"--n-header-extra-text-color-disabled":_e,"--n-header-font-weight":xe,"--n-header-text-color":X,"--n-header-text-color-disabled":me,"--n-item-color-pending":ve,"--n-item-height":Oe,"--n-item-text-color":ge,"--n-item-text-color-disabled":be,"--n-list-color":W,"--n-header-height":Pe,"--n-close-size":ke,"--n-close-icon-size":Te,"--n-close-color-hover":Se,"--n-close-color-pressed":Ce,"--n-close-icon-color":Re,"--n-close-icon-color-hover":ye,"--n-close-icon-color-pressed":ze,"--n-divider-color":Fe}})}},render(){const{mergedClsPrefix:e,renderSourceList:i,renderTargetList:o,mergedTheme:t,mergedSrcFilterable:n,targetFilterable:s}=this;return l("div",{class:[`${e}-transfer`,this.mergedDisabled&&`${e}-transfer--disabled`],style:this.cssVars},l("div",{class:`${e}-transfer-list ${e}-transfer-list--source`},l(re,{source:!0,title:this.sourceTitle,onCheckedAll:this.handleSourceCheckAll,onClearAll:this.handleSourceUncheckAll,size:this.mergedSize}),l("div",{class:`${e}-transfer-list-body`},n?l(oe,{onUpdateValue:this.handleSrcFilterUpdateValue,value:this.srcPattern,disabled:this.mergedDisabled,placeholder:this.sourceFilterPlaceholder}):null,l("div",{class:`${e}-transfer-list-flex-container`},i?l(G,{theme:t.peers.Scrollbar,themeOverrides:t.peerOverrides.Scrollbar},{default:()=>i({onCheck:this.handleChecked,checkedOptions:this.filteredTgtOpts,pattern:this.srcPattern})}):l(ae,{source:!0,options:this.filteredSrcOpts,disabled:this.mergedDisabled,virtualScroll:this.virtualScroll,itemSize:this.itemSize}))),l("div",{class:`${e}-transfer-list__border`})),l("div",{class:`${e}-transfer-list ${e}-transfer-list--target`},l(re,{onClearAll:this.handleTargetClearAll,size:this.mergedSize,title:this.targetTitle}),l("div",{class:`${e}-transfer-list-body`},s?l(oe,{onUpdateValue:this.handleTgtFilterUpdateValue,value:this.tgtPattern,disabled:this.mergedDisabled,placeholder:this.sourceFilterPlaceholder}):null,l("div",{class:`${e}-transfer-list-flex-container`},o?l(G,{theme:t.peers.Scrollbar,themeOverrides:t.peerOverrides.Scrollbar},{default:()=>o({onCheck:this.handleChecked,checkedOptions:this.filteredTgtOpts,pattern:this.tgtPattern})}):l(ae,{options:this.filteredTgtOpts,disabled:this.mergedDisabled,virtualScroll:this.virtualScroll,itemSize:this.itemSize}))),l("div",{class:`${e}-transfer-list__border`})))}}),Ot=L({__name:"ApiTransfer",setup(e){const{data:i,loading:o}=J(ot.getAllApi,{auto:!0}),t=x(()=>{var u,r,h;return(u=i.value)!=null&&u.list?(h=(r=i.value)==null?void 0:r.list)==null?void 0:h.map(d=>({...d,label:d.title,value:d.id})):[]}),n=function({option:s}){const{title:u,method:r,path:h,name:d}=s;return l(ut,{},{default:()=>`${r} ${h}`,trigger:()=>`${u} ${d}`})};return(s,u)=>(Q(),se(p($t),Ye({loading:p(o),options:p(t)},s.$attrs,{"virtual-scroll":"","source-filterable":"","target-filterable":"","render-source-label":n,"render-target-label":n}),null,16,["loading","options","render-source-label","render-target-label"]))}}),Pt={flex:"","justify-end":""},ie=L({__name:"RoleModal",props:{record:null},emits:["refresh"],setup(e,{emit:i}){const o=e,{record:t}=Ge(o),n=$(!1),s=$(null),{run:u,loading:r}=J(j.createRole),{run:h,loading:d}=J(j.editRole),c=$({name:"",parent_id:0,title:"",description:"",sort:100}),v={name:{required:!0,message:"请输入角色名",trigger:"blur"},code:{required:!0,message:"请输入code",trigger:"blur"}};function F(){n.value=!0,t!=null&&t.value&&(c.value={...t.value,api_ids:t.value.apis.map(b=>b.id)})}const z=x(()=>!!(t!=null&&t.value));function R(){var b;(b=s.value)==null||b.validate(async f=>{if(f)return;const{result:g}=t!=null&&t.value?await h(t==null?void 0:t.value.id,c.value):await u(c.value);g&&(n.value=!1,i("refresh"))})}return(b,f)=>(Q(),Je(tt,null,[te("span",{onClick:F},[Qe(b.$slots,"default")]),S(p(et),{show:n.value,"onUpdate:show":f[8]||(f[8]=g=>n.value=g),title:"角色",preset:"card","auto-focus":!1,"w-200":""},{default:k(()=>[S(p(xt),{ref_key:"formRef",ref:s,model:c.value,rules:v},{default:k(()=>[S(p(A),{path:"title",label:"角色名"},{default:k(()=>[S(p(E),{value:c.value.title,"onUpdate:value":f[0]||(f[0]=g=>c.value.title=g),placeholder:"输入角色名",onKeydown:f[1]||(f[1]=Ze(rt(()=>{},["prevent"]),["enter"]))},null,8,["value"])]),_:1}),S(p(A),{path:"name",label:"标识"},{default:k(()=>[S(p(E),{value:c.value.name,"onUpdate:value":f[2]||(f[2]=g=>c.value.name=g),disabled:p(z),placeholder:"请输入name"},null,8,["value","disabled"])]),_:1}),S(p(A),{path:"description",label:"描述"},{default:k(()=>[S(p(E),{value:c.value.description,"onUpdate:value":f[3]||(f[3]=g=>c.value.description=g),type:"textarea",placeholder:"请输入描述"},null,8,["value"])]),_:1}),S(p(A),{path:"parent_id",label:"父节点"},{default:k(()=>[S(at,{value:c.value.parent_id,"onUpdate:value":f[4]||(f[4]=g=>c.value.parent_id=g),clearable:"",placeholder:"请输入父节点"},null,8,["value"])]),_:1}),S(p(A),{path:"sort",label:"排序"},{default:k(()=>[S(p(bt),{value:c.value.sort,"onUpdate:value":f[5]||(f[5]=g=>c.value.sort=g),type:"textarea",placeholder:"请输入排序"},null,8,["value"])]),_:1}),S(p(A),{path:"api_ids",label:"权限"},{default:k(()=>[S(Ot,{value:c.value.api_ids,"onUpdate:value":f[6]||(f[6]=g=>c.value.api_ids=g)},null,8,["value"])]),_:1}),te("div",Pt,[S(p(q),{type:"primary",loading:p(r)||p(d),onClick:f[7]||(f[7]=g=>R())},{default:k(()=>[de(" 确定 ")]),_:1},8,["loading"])])]),_:1},8,["model"])]),_:1},8,["show"])],64))}}),Yt=L({__name:"RolePageTree",setup(e){const{dataSource:i,loading:o,refresh:t,tableFilter:n,total:s,pagination:u}=mt(j.getRoleTreeList);async function r(d){const{result:c}=await j.delRole(d);c&&t()}const h=[{title:"序号",key:"id",width:"205",resizable:!0},{title:"角色名",key:"title",resizable:!0,render(d){return l(Rt,{label:d.description},{default:()=>d.title})}},{title:"标识",key:"name"},{title:"排序",key:"sort"},{title:"更新时间",key:"updated_at",width:"185"},{title:"操作",key:"action",width:"170",fixed:"right",render(d){return[l(ie,{record:d,onRefresh:t},{default:()=>l(St)}),l(Ct,{confirmText:"是否删除该角色",delFunction:()=>r(d.id)})]}}];return(d,c)=>(Q(),se(pt,{"table-filter":p(n),"onUpdate:tableFilter":c[0]||(c[0]=v=>lt(n)?n.value=v:null),data:p(i),loading:p(o),columns:h,total:p(s),pagination:p(u),"search-key":"name",tree:!0,"row-key":v=>`${v.id}-${v.name}`,"search-place-holder":"请输入角色名",onRefresh:p(t)},{header:k(()=>[S(ie,{onRefresh:p(t)},{default:k(()=>[S(p(q),{type:"primary"},{default:k(()=>[de(" 新增角色 ")]),_:1})]),_:1},8,["onRefresh"])]),_:1},8,["table-filter","data","loading","total","pagination","row-key","onRefresh"]))}});export{Yt as default};
