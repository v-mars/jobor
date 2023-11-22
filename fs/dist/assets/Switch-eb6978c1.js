import{a as se,b as de,e as A,v as t,a_ as O,t as P,f as l,w as U,d as ce,g as ue,u as I,ag as he,m as T,z as be,i as z,j as fe,aB as N,h as o,aM as m,a1 as p,a$ as j,ai as s,b0 as ge,b1 as ve,A as D}from "./index-bbb52c5f.js";import{u as me}from "./Icon-01f3be5d.js";const pe={buttonHeightSmall:"14px",buttonHeightMedium:"18px",buttonHeightLarge:"22px",buttonWidthSmall:"14px",buttonWidthMedium:"18px",buttonWidthLarge:"22px",buttonWidthPressedSmall:"20px",buttonWidthPressedMedium:"24px",buttonWidthPressedLarge:"28px",railHeightSmall:"18px",railHeightMedium:"22px",railHeightLarge:"26px",railWidthSmall:"32px",railWidthMedium:"40px",railWidthLarge:"48px"},we= e=>{const{primaryColor:d,opacityDisabled:f,borderRadius:n,textColor3:r}=e,w="rgba(0, 0, 0, .14)";return Object.assign(Object.assign({},pe),{iconColor:r,textColor:"white",loadingColor:d,opacityDisabled:f,railColor:w,railColorActive:d,buttonBoxShadow:"0 1px 4px 0 rgba(0, 0, 0, 0.3), inset 0 0 1px 0 rgba(0, 0, 0, 0.05)",buttonColor:"#FFF",railBorderRadiusSmall:n,railBorderRadiusMedium:n,railBorderRadiusLarge:n,buttonBorderRadiusSmall:n,buttonBorderRadiusMedium:n,buttonBorderRadiusLarge:n,boxShadowFocus:`0 0 0 2px ${de(d,{alpha:.2})}`})},xe={name:"Switch",common:se,self:we},ye=xe,ke=A("switch",`
 height: var(--n-height);
 min-width: var(--n-width);
 vertical-align: middle;
 user-select: none;
 -webkit-user-select: none;
 display: inline-flex;
 outline: none;
 justify-content: center;
 align-items: center;
`,[t("children-placeholder",`
 height: var(--n-rail-height);
 display: flex;
 flex-direction: column;
 overflow: hidden;
 pointer-events: none;
 visibility: hidden;
 `),t("rail-placeholder",`
 display: flex;
 flex-wrap: none;
 `),t("button-placeholder",`
 width: calc(1.75 * var(--n-rail-height));
 height: var(--n-rail-height);
 `),A("base-loading",`
 position: absolute;
 top: 50%;
 left: 50%;
 transform: translateX(-50%) translateY(-50%);
 font-size: calc(var(--n-button-width) - 4px);
 color: var(--n-loading-color);
 transition: color .3s var(--n-bezier);
 `,[O({left:"50%",top:"50%",originalTransform:"translateX(-50%) translateY(-50%)"})]),t("checked, unchecked",`
 transition: color .3s var(--n-bezier);
 color: var(--n-text-color);
 box-sizing: border-box;
 position: absolute;
 white-space: nowrap;
 top: 0;
 bottom: 0;
 display: flex;
 align-items: center;
 line-height: 1;
 `),t("checked",`
 right: 0;
 padding-right: calc(1.25 * var(--n-rail-height) - var(--n-offset));
 `),t("unchecked",`
 left: 0;
 justify-content: flex-end;
 padding-left: calc(1.25 * var(--n-rail-height) - var(--n-offset));
 `),P("&:focus",[t("rail",`
 box-shadow: var(--n-box-shadow-focus);
 `)]),l("round",[t("rail","border-radius: calc(var(--n-rail-height) / 2);",[t("button","border-radius: calc(var(--n-button-height) / 2);")])]),U("disabled",[U("icon",[l("rubber-band",[l("pressed",[t("rail",[t("button","max-width: var(--n-button-width-pressed);")])]),t("rail",[P("&:active",[t("button","max-width: var(--n-button-width-pressed);")])]),l("active",[l("pressed",[t("rail",[t("button","left: calc(100% - var(--n-offset) - var(--n-button-width-pressed));")])]),t("rail",[P("&:active",[t("button","left: calc(100% - var(--n-offset) - var(--n-button-width-pressed));")])])])])])]),l("active",[t("rail",[t("button","left: calc(100% - var(--n-button-width) - var(--n-offset))")])]),t("rail",`
 overflow: hidden;
 height: var(--n-rail-height);
 min-width: var(--n-rail-width);
 border-radius: var(--n-rail-border-radius);
 cursor: pointer;
 position: relative;
 transition:
 opacity .3s var(--n-bezier),
 background .3s var(--n-bezier),
 box-shadow .3s var(--n-bezier);
 background-color: var(--n-rail-color);
 `,[t("button-icon",`
 color: var(--n-icon-color);
 transition: color .3s var(--n-bezier);
 font-size: calc(var(--n-button-height) - 4px);
 position: absolute;
 left: 0;
 right: 0;
 top: 0;
 bottom: 0;
 display: flex;
 justify-content: center;
 align-items: center;
 line-height: 1;
 `,[O()]),t("button",`
 align-items: center; 
 top: var(--n-offset);
 left: var(--n-offset);
 height: var(--n-button-height);
 width: var(--n-button-width-pressed);
 max-width: var(--n-button-width);
 border-radius: var(--n-button-border-radius);
 background-color: var(--n-button-color);
 box-shadow: var(--n-button-box-shadow);
 box-sizing: border-box;
 cursor: inherit;
 content: "";
 position: absolute;
 transition:
 background-color .3s var(--n-bezier),
 left .3s var(--n-bezier),
 opacity .3s var(--n-bezier),
 max-width .3s var(--n-bezier),
 box-shadow .3s var(--n-bezier);
 `)]),l("active",[t("rail","background-color: var(--n-rail-color-active);")]),l("loading",[t("rail",`
 cursor: wait;
 `)]),l("disabled",[t("rail",`
 cursor: not-allowed;
 opacity: .5;
 `)])]),Se=Object.assign(Object.assign({},I.props),{size:{type:String,default:"medium"},value:{type:[String,Number,Boolean],default:void 0},loading:Boolean,defaultValue:{type:[String,Number,Boolean],default:!1},disabled:{type:Boolean,default:void 0},round:{type:Boolean,default:!0},"onUpdate:value":[Function,Array],onUpdateValue:[Function,Array],checkedValue:{type:[String,Number,Boolean],default:!0},uncheckedValue:{type:[String,Number,Boolean],default:!1},railStyle:Function,rubberBand:{type:Boolean,default:!0},onChange:[Function,Array]});let B;const $e=ce({name:"Switch",props:Se,setup(e){B===void 0&&(typeof CSS<"u"?typeof CSS.supports<"u"?B=CSS.supports("width","max(1px)"):B=!1:B=!0);const{mergedClsPrefixRef:d,inlineThemeDisabled:f}=ue(e),n=I("Switch","-switch",ke,ye,e,d),r=he(e),{mergedSizeRef:w,mergedDisabledRef:g}=r,k=T(e.defaultValue),$=be(e,"value"),v=me($,k),S=z(()=>v.value===e.checkedValue),x=T(!1),a=T(!1),c=z(()=>{const{railStyle:i}=e;if(i)return i({focused:a.value,checked:S.value})});function u(i){const{"onUpdate:value":R,onChange:_,onUpdateValue:V}=e,{nTriggerFormInput:F,nTriggerFormChange:W}=r;R&&D(R,i),V&&D(V,i),_&&D(_,i),k.value=i,F(),W()}function E(){const{nTriggerFormFocus:i}=r;i()}function X(){const{nTriggerFormBlur:i}=r;i()}function Y(){e.loading||g.value||(v.value!==e.checkedValue?u(e.checkedValue):u(e.uncheckedValue))}function q(){a.value=!0,E()}function G(){a.value=!1,X(),x.value=!1}function J(i){e.loading||g.value||i.key===" "&&(v.value!==e.checkedValue?u(e.checkedValue):u(e.uncheckedValue),x.value=!1)}function Q(i){e.loading||g.value||i.key===" "&&(i.preventDefault(),x.value=!0)}const K=z(()=>{const{value:i}=w,{self:{opacityDisabled:R,railColor:_,railColorActive:V,buttonBoxShadow:F,buttonColor:W,boxShadowFocus:Z,loadingColor:ee,textColor:te,iconColor:ie,[p("buttonHeight",i)]:h,[p("buttonWidth",i)]:ae,[p("buttonWidthPressed",i)]:oe,[p("railHeight",i)]:b,[p("railWidth",i)]:C,[p("railBorderRadius",i)]:ne,[p("buttonBorderRadius",i)]:re},common:{cubicBezierEaseInOut:le}}=n.value;let M,H,L;return B?(M=`calc((${b} - ${h}) / 2)`,H=`max(${b}, ${h})`,L=`max(${C}, calc(${C} + ${h} - ${b}))`):(M=j((s(b)-s(h))/2),H=j(Math.max(s(b),s(h))),L=s(b)>s(h)?C:j(s(C)+s(h)-s(b))),{"--n-bezier":le,"--n-button-border-radius":re,"--n-button-box-shadow":F,"--n-button-color":W,"--n-button-width":ae,"--n-button-width-pressed":oe,"--n-button-height":h,"--n-height":H,"--n-offset":M,"--n-opacity-disabled":R,"--n-rail-border-radius":ne,"--n-rail-color":_,"--n-rail-color-active":V,"--n-rail-height":b,"--n-rail-width":C,"--n-width":L,"--n-box-shadow-focus":Z,"--n-loading-color":ee,"--n-text-color":te,"--n-icon-color":ie}}),y=f?fe("switch",z(()=>w.value[0]),K,e):void 0;return{handleClick:Y,handleBlur:G,handleFocus:q,handleKeyup:J,handleKeydown:Q,mergedRailStyle:c,pressed:x,mergedClsPrefix:d,mergedValue:v,checked:S,mergedDisabled:g,cssVars:f?void 0:K,themeClass:y==null?void 0:y.themeClass,onRender:y==null?void 0:y.onRender}},render(){const{mergedClsPrefix:e,mergedDisabled:d,checked:f,mergedRailStyle:n,onRender:r,$slots:w}=this;r==null||r();const{checked:g,unchecked:k,icon:$,"checked-icon":v,"unchecked-icon":S}=w,x=!(N($)&&N(v)&&N(S));return o("div",{role:"switch","aria-checked":f,class:[`${e}-switch`,this.themeClass,x&&`${e}-switch--icon`,f&&`${e}-switch--active`,d&&`${e}-switch--disabled`,this.round&&`${e}-switch--round`,this.loading&&`${e}-switch--loading`,this.pressed&&`${e}-switch--pressed`,this.rubberBand&&`${e}-switch--rubber-band`],tabindex:this.mergedDisabled?void 0:0,style:this.cssVars,onClick:this.handleClick,onFocus:this.handleFocus,onBlur:this.handleBlur,onKeyup:this.handleKeyup,onKeydown:this.handleKeydown},o("div",{class:`${e}-switch__rail`,"aria-hidden":"true",style:n},m(g,a=>m(k,c=>a||c?o("div",{"aria-hidden":!0,class:`${e}-switch__children-placeholder`},o("div",{class:`${e}-switch__rail-placeholder`},o("div",{class:`${e}-switch__button-placeholder`}),a),o("div",{class:`${e}-switch__rail-placeholder`},o("div",{class:`${e}-switch__button-placeholder`}),c)):null)),o("div",{class:`${e}-switch__button`},m($,a=>m(v,c=>m(S,u=>o(ge,null,{default:()=>this.loading?o(ve,{key:"loading",clsPrefix:e,strokeWidth:20}):this.checked&&(c||a)?o("div",{class:`${e}-switch__button-icon`,key:c?"checked-icon":"icon"},c||a):!this.checked&&(u||a)?o("div",{class:`${e}-switch__button-icon`,key:u?"unchecked-icon":"icon"},u||a):null})))),m(g,a=>a&&o("div",{key:"checked",class:`${e}-switch__checked`},a)),m(k,a=>a&&o("div",{key:"unchecked",class:`${e}-switch__unchecked`},a)))))}});export{$e as N};
