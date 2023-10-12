import{M as le,i as P,cc as xe,cd as Y,ce as ft,cf as J,l as We,bG as vt,d as R,h as a,b_ as pt,e as w,t as S,v as f,a_ as mt,at as Be,z as Se,b0 as gt,az as G,N as se,b1 as bt,a as Ie,b as we,k as yt,m as C,aD as xt,f as F,w as Z,g as Le,u as ce,cg as wt,ag as Ct,n as $e,aw as St,bb as Pt,y as Ae,p as Mt,ax as zt,j as Ve,aM as ae,af as Ft,F as Tt,aN as $t,aq as De,b8 as _e,a1 as Ce,b5 as At,A as M,bg as Re,bd as Dt,a9 as _t}from"./index-3a30b39c.js";const Rt=/^(\d|\.)+$/,ke=/(\d|\.)+/;function kt(o,{c:n=1,offset:r=0,attachPx:p=!0}={}){if(typeof o=="number"){const s=(o+r)*n;return s===0?"0":`${s}px`}else if(typeof o=="string")if(Rt.test(o)){const s=(Number(o)+r)*n;return p?s===0?"0":`${s}px`:`${s}`}else{const s=ke.exec(o);return s?o.replace(ke,String((Number(s[0])+r)*n)):o}return o}function Et(o,n){return le(o,r=>{r!==void 0&&(n.value=r)}),P(()=>o.value===void 0?n.value:o.value)}function Hr(o,n){return P(()=>{for(const r of n)if(o[r]!==void 0)return o[r];return o[n[n.length-1]]})}const Wt={name:"en-US",global:{undo:"Undo",redo:"Redo",confirm:"Confirm",clear:"Clear"},Popconfirm:{positiveText:"Confirm",negativeText:"Cancel"},Cascader:{placeholder:"Please Select",loading:"Loading",loadingRequiredMessage:o=>`Please load all ${o}'s descendants before checking it.`},Time:{dateFormat:"yyyy-MM-dd",dateTimeFormat:"yyyy-MM-dd HH:mm:ss"},DatePicker:{yearFormat:"yyyy",monthFormat:"MMM",dayFormat:"eeeeee",yearTypeFormat:"yyyy",monthTypeFormat:"yyyy-MM",dateFormat:"yyyy-MM-dd",dateTimeFormat:"yyyy-MM-dd HH:mm:ss",quarterFormat:"yyyy-qqq",clear:"Clear",now:"Now",confirm:"Confirm",selectTime:"Select Time",selectDate:"Select Date",datePlaceholder:"Select Date",datetimePlaceholder:"Select Date and Time",monthPlaceholder:"Select Month",yearPlaceholder:"Select Year",quarterPlaceholder:"Select Quarter",startDatePlaceholder:"Start Date",endDatePlaceholder:"End Date",startDatetimePlaceholder:"Start Date and Time",endDatetimePlaceholder:"End Date and Time",startMonthPlaceholder:"Start Month",endMonthPlaceholder:"End Month",monthBeforeYear:!0,firstDayOfWeek:6,today:"Today"},DataTable:{checkTableAll:"Select all in the table",uncheckTableAll:"Unselect all in the table",confirm:"Confirm",clear:"Clear"},LegacyTransfer:{sourceTitle:"Source",targetTitle:"Target"},Transfer:{selectAll:"Select all",unselectAll:"Unselect all",clearAll:"Clear",total:o=>`Total ${o} items`,selected:o=>`${o} items selected`},Empty:{description:"No Data"},Select:{placeholder:"Please Select"},TimePicker:{placeholder:"Select Time",positiveText:"OK",negativeText:"Cancel",now:"Now"},Pagination:{goto:"Goto",selectionSuffix:"page"},DynamicTags:{add:"Add"},Log:{loading:"Loading"},Input:{placeholder:"Please Input"},InputNumber:{placeholder:"Please Input"},DynamicInput:{create:"Create"},ThemeEditor:{title:"Theme Editor",clearAllVars:"Clear All Variables",clearSearch:"Clear Search",filterCompName:"Filter Component Name",filterVarName:"Filter Variable Name",import:"Import",export:"Export",restore:"Reset to Default"},Image:{tipPrevious:"Previous picture (←)",tipNext:"Next picture (→)",tipCounterclockwise:"Counterclockwise",tipClockwise:"Clockwise",tipZoomOut:"Zoom out",tipZoomIn:"Zoom in",tipClose:"Close (Esc)",tipOriginalSize:"Zoom to original size"}},Bt=Wt;var It={lessThanXSeconds:{one:"less than a second",other:"less than {{count}} seconds"},xSeconds:{one:"1 second",other:"{{count}} seconds"},halfAMinute:"half a minute",lessThanXMinutes:{one:"less than a minute",other:"less than {{count}} minutes"},xMinutes:{one:"1 minute",other:"{{count}} minutes"},aboutXHours:{one:"about 1 hour",other:"about {{count}} hours"},xHours:{one:"1 hour",other:"{{count}} hours"},xDays:{one:"1 day",other:"{{count}} days"},aboutXWeeks:{one:"about 1 week",other:"about {{count}} weeks"},xWeeks:{one:"1 week",other:"{{count}} weeks"},aboutXMonths:{one:"about 1 month",other:"about {{count}} months"},xMonths:{one:"1 month",other:"{{count}} months"},aboutXYears:{one:"about 1 year",other:"about {{count}} years"},xYears:{one:"1 year",other:"{{count}} years"},overXYears:{one:"over 1 year",other:"over {{count}} years"},almostXYears:{one:"almost 1 year",other:"almost {{count}} years"}},Lt=function(n,r,p){var s,l=It[n];return typeof l=="string"?s=l:r===1?s=l.one:s=l.other.replace("{{count}}",r.toString()),p!=null&&p.addSuffix?p.comparison&&p.comparison>0?"in "+s:s+" ago":s};const Vt=Lt;var Nt={full:"EEEE, MMMM do, y",long:"MMMM do, y",medium:"MMM d, y",short:"MM/dd/yyyy"},Ot={full:"h:mm:ss a zzzz",long:"h:mm:ss a z",medium:"h:mm:ss a",short:"h:mm a"},jt={full:"{{date}} 'at' {{time}}",long:"{{date}} 'at' {{time}}",medium:"{{date}}, {{time}}",short:"{{date}}, {{time}}"},Ht={date:xe({formats:Nt,defaultWidth:"full"}),time:xe({formats:Ot,defaultWidth:"full"}),dateTime:xe({formats:jt,defaultWidth:"full"})};const Ut=Ht;var qt={lastWeek:"'last' eeee 'at' p",yesterday:"'yesterday at' p",today:"'today at' p",tomorrow:"'tomorrow at' p",nextWeek:"eeee 'at' p",other:"P"},Kt=function(n,r,p,s){return qt[n]};const Xt=Kt;var Yt={narrow:["B","A"],abbreviated:["BC","AD"],wide:["Before Christ","Anno Domini"]},Jt={narrow:["1","2","3","4"],abbreviated:["Q1","Q2","Q3","Q4"],wide:["1st quarter","2nd quarter","3rd quarter","4th quarter"]},Zt={narrow:["J","F","M","A","M","J","J","A","S","O","N","D"],abbreviated:["Jan","Feb","Mar","Apr","May","Jun","Jul","Aug","Sep","Oct","Nov","Dec"],wide:["January","February","March","April","May","June","July","August","September","October","November","December"]},Gt={narrow:["S","M","T","W","T","F","S"],short:["Su","Mo","Tu","We","Th","Fr","Sa"],abbreviated:["Sun","Mon","Tue","Wed","Thu","Fri","Sat"],wide:["Sunday","Monday","Tuesday","Wednesday","Thursday","Friday","Saturday"]},Qt={narrow:{am:"a",pm:"p",midnight:"mi",noon:"n",morning:"morning",afternoon:"afternoon",evening:"evening",night:"night"},abbreviated:{am:"AM",pm:"PM",midnight:"midnight",noon:"noon",morning:"morning",afternoon:"afternoon",evening:"evening",night:"night"},wide:{am:"a.m.",pm:"p.m.",midnight:"midnight",noon:"noon",morning:"morning",afternoon:"afternoon",evening:"evening",night:"night"}},er={narrow:{am:"a",pm:"p",midnight:"mi",noon:"n",morning:"in the morning",afternoon:"in the afternoon",evening:"in the evening",night:"at night"},abbreviated:{am:"AM",pm:"PM",midnight:"midnight",noon:"noon",morning:"in the morning",afternoon:"in the afternoon",evening:"in the evening",night:"at night"},wide:{am:"a.m.",pm:"p.m.",midnight:"midnight",noon:"noon",morning:"in the morning",afternoon:"in the afternoon",evening:"in the evening",night:"at night"}},or=function(n,r){var p=Number(n),s=p%100;if(s>20||s<10)switch(s%10){case 1:return p+"st";case 2:return p+"nd";case 3:return p+"rd"}return p+"th"},tr={ordinalNumber:or,era:Y({values:Yt,defaultWidth:"wide"}),quarter:Y({values:Jt,defaultWidth:"wide",argumentCallback:function(n){return n-1}}),month:Y({values:Zt,defaultWidth:"wide"}),day:Y({values:Gt,defaultWidth:"wide"}),dayPeriod:Y({values:Qt,defaultWidth:"wide",formattingValues:er,defaultFormattingWidth:"wide"})};const rr=tr;var nr=/^(\d+)(th|st|nd|rd)?/i,ar=/\d+/i,ir={narrow:/^(b|a)/i,abbreviated:/^(b\.?\s?c\.?|b\.?\s?c\.?\s?e\.?|a\.?\s?d\.?|c\.?\s?e\.?)/i,wide:/^(before christ|before common era|anno domini|common era)/i},lr={any:[/^b/i,/^(a|c)/i]},sr={narrow:/^[1234]/i,abbreviated:/^q[1234]/i,wide:/^[1234](th|st|nd|rd)? quarter/i},cr={any:[/1/i,/2/i,/3/i,/4/i]},dr={narrow:/^[jfmasond]/i,abbreviated:/^(jan|feb|mar|apr|may|jun|jul|aug|sep|oct|nov|dec)/i,wide:/^(january|february|march|april|may|june|july|august|september|october|november|december)/i},ur={narrow:[/^j/i,/^f/i,/^m/i,/^a/i,/^m/i,/^j/i,/^j/i,/^a/i,/^s/i,/^o/i,/^n/i,/^d/i],any:[/^ja/i,/^f/i,/^mar/i,/^ap/i,/^may/i,/^jun/i,/^jul/i,/^au/i,/^s/i,/^o/i,/^n/i,/^d/i]},hr={narrow:/^[smtwf]/i,short:/^(su|mo|tu|we|th|fr|sa)/i,abbreviated:/^(sun|mon|tue|wed|thu|fri|sat)/i,wide:/^(sunday|monday|tuesday|wednesday|thursday|friday|saturday)/i},fr={narrow:[/^s/i,/^m/i,/^t/i,/^w/i,/^t/i,/^f/i,/^s/i],any:[/^su/i,/^m/i,/^tu/i,/^w/i,/^th/i,/^f/i,/^sa/i]},vr={narrow:/^(a|p|mi|n|(in the|at) (morning|afternoon|evening|night))/i,any:/^([ap]\.?\s?m\.?|midnight|noon|(in the|at) (morning|afternoon|evening|night))/i},pr={any:{am:/^a/i,pm:/^p/i,midnight:/^mi/i,noon:/^no/i,morning:/morning/i,afternoon:/afternoon/i,evening:/evening/i,night:/night/i}},mr={ordinalNumber:ft({matchPattern:nr,parsePattern:ar,valueCallback:function(n){return parseInt(n,10)}}),era:J({matchPatterns:ir,defaultMatchWidth:"wide",parsePatterns:lr,defaultParseWidth:"any"}),quarter:J({matchPatterns:sr,defaultMatchWidth:"wide",parsePatterns:cr,defaultParseWidth:"any",valueCallback:function(n){return n+1}}),month:J({matchPatterns:dr,defaultMatchWidth:"wide",parsePatterns:ur,defaultParseWidth:"any"}),day:J({matchPatterns:hr,defaultMatchWidth:"wide",parsePatterns:fr,defaultParseWidth:"any"}),dayPeriod:J({matchPatterns:vr,defaultMatchWidth:"any",parsePatterns:pr,defaultParseWidth:"any"})};const gr=mr;var br={code:"en-US",formatDistance:Vt,formatLong:Ut,formatRelative:Xt,localize:rr,match:gr,options:{weekStartsOn:0,firstWeekContainsDate:1}};const yr=br,xr={name:"en-US",locale:yr},wr=xr;function Cr(o){const{mergedLocaleRef:n,mergedDateLocaleRef:r}=We(vt,null)||{},p=P(()=>{var l,c;return(c=(l=n==null?void 0:n.value)===null||l===void 0?void 0:l[o])!==null&&c!==void 0?c:Bt[o]});return{dateLocaleRef:P(()=>{var l;return(l=r==null?void 0:r.value)!==null&&l!==void 0?l:wr}),localeRef:p}}const Sr=R({name:"Eye",render(){return a("svg",{xmlns:"http://www.w3.org/2000/svg",viewBox:"0 0 512 512"},a("path",{d:"M255.66 112c-77.94 0-157.89 45.11-220.83 135.33a16 16 0 0 0-.27 17.77C82.92 340.8 161.8 400 255.66 400c92.84 0 173.34-59.38 221.79-135.25a16.14 16.14 0 0 0 0-17.47C428.89 172.28 347.8 112 255.66 112z",fill:"none",stroke:"currentColor","stroke-linecap":"round","stroke-linejoin":"round","stroke-width":"32"}),a("circle",{cx:"256",cy:"256",r:"80",fill:"none",stroke:"currentColor","stroke-miterlimit":"10","stroke-width":"32"}))}}),Pr=R({name:"EyeOff",render(){return a("svg",{xmlns:"http://www.w3.org/2000/svg",viewBox:"0 0 512 512"},a("path",{d:"M432 448a15.92 15.92 0 0 1-11.31-4.69l-352-352a16 16 0 0 1 22.62-22.62l352 352A16 16 0 0 1 432 448z",fill:"currentColor"}),a("path",{d:"M255.66 384c-41.49 0-81.5-12.28-118.92-36.5c-34.07-22-64.74-53.51-88.7-91v-.08c19.94-28.57 41.78-52.73 65.24-72.21a2 2 0 0 0 .14-2.94L93.5 161.38a2 2 0 0 0-2.71-.12c-24.92 21-48.05 46.76-69.08 76.92a31.92 31.92 0 0 0-.64 35.54c26.41 41.33 60.4 76.14 98.28 100.65C162 402 207.9 416 255.66 416a239.13 239.13 0 0 0 75.8-12.58a2 2 0 0 0 .77-3.31l-21.58-21.58a4 4 0 0 0-3.83-1a204.8 204.8 0 0 1-51.16 6.47z",fill:"currentColor"}),a("path",{d:"M490.84 238.6c-26.46-40.92-60.79-75.68-99.27-100.53C349 110.55 302 96 255.66 96a227.34 227.34 0 0 0-74.89 12.83a2 2 0 0 0-.75 3.31l21.55 21.55a4 4 0 0 0 3.88 1a192.82 192.82 0 0 1 50.21-6.69c40.69 0 80.58 12.43 118.55 37c34.71 22.4 65.74 53.88 89.76 91a.13.13 0 0 1 0 .16a310.72 310.72 0 0 1-64.12 72.73a2 2 0 0 0-.15 2.95l19.9 19.89a2 2 0 0 0 2.7.13a343.49 343.49 0 0 0 68.64-78.48a32.2 32.2 0 0 0-.1-34.78z",fill:"currentColor"}),a("path",{d:"M256 160a95.88 95.88 0 0 0-21.37 2.4a2 2 0 0 0-1 3.38l112.59 112.56a2 2 0 0 0 3.38-1A96 96 0 0 0 256 160z",fill:"currentColor"}),a("path",{d:"M165.78 233.66a2 2 0 0 0-3.38 1a96 96 0 0 0 115 115a2 2 0 0 0 1-3.38z",fill:"currentColor"}))}}),Mr=R({name:"ChevronDown",render(){return a("svg",{viewBox:"0 0 16 16",fill:"none",xmlns:"http://www.w3.org/2000/svg"},a("path",{d:"M3.14645 5.64645C3.34171 5.45118 3.65829 5.45118 3.85355 5.64645L8 9.79289L12.1464 5.64645C12.3417 5.45118 12.6583 5.45118 12.8536 5.64645C13.0488 5.84171 13.0488 6.15829 12.8536 6.35355L8.35355 10.8536C8.15829 11.0488 7.84171 11.0488 7.64645 10.8536L3.14645 6.35355C2.95118 6.15829 2.95118 5.84171 3.14645 5.64645Z",fill:"currentColor"}))}}),zr=pt("clear",a("svg",{viewBox:"0 0 16 16",version:"1.1",xmlns:"http://www.w3.org/2000/svg"},a("g",{stroke:"none","stroke-width":"1",fill:"none","fill-rule":"evenodd"},a("g",{fill:"currentColor","fill-rule":"nonzero"},a("path",{d:"M8,2 C11.3137085,2 14,4.6862915 14,8 C14,11.3137085 11.3137085,14 8,14 C4.6862915,14 2,11.3137085 2,8 C2,4.6862915 4.6862915,2 8,2 Z M6.5343055,5.83859116 C6.33943736,5.70359511 6.07001296,5.72288026 5.89644661,5.89644661 L5.89644661,5.89644661 L5.83859116,5.9656945 C5.70359511,6.16056264 5.72288026,6.42998704 5.89644661,6.60355339 L5.89644661,6.60355339 L7.293,8 L5.89644661,9.39644661 L5.83859116,9.4656945 C5.70359511,9.66056264 5.72288026,9.92998704 5.89644661,10.1035534 L5.89644661,10.1035534 L5.9656945,10.1614088 C6.16056264,10.2964049 6.42998704,10.2771197 6.60355339,10.1035534 L6.60355339,10.1035534 L8,8.707 L9.39644661,10.1035534 L9.4656945,10.1614088 C9.66056264,10.2964049 9.92998704,10.2771197 10.1035534,10.1035534 L10.1035534,10.1035534 L10.1614088,10.0343055 C10.2964049,9.83943736 10.2771197,9.57001296 10.1035534,9.39644661 L10.1035534,9.39644661 L8.707,8 L10.1035534,6.60355339 L10.1614088,6.5343055 C10.2964049,6.33943736 10.2771197,6.07001296 10.1035534,5.89644661 L10.1035534,5.89644661 L10.0343055,5.83859116 C9.83943736,5.70359511 9.57001296,5.72288026 9.39644661,5.89644661 L9.39644661,5.89644661 L8,7.293 L6.60355339,5.89644661 Z"}))))),Fr=w("base-clear",`
 flex-shrink: 0;
 height: 1em;
 width: 1em;
 position: relative;
`,[S(">",[f("clear",`
 font-size: var(--n-clear-size);
 height: 1em;
 width: 1em;
 cursor: pointer;
 color: var(--n-clear-color);
 transition: color .3s var(--n-bezier);
 display: flex;
 `,[S("&:hover",`
 color: var(--n-clear-color-hover)!important;
 `),S("&:active",`
 color: var(--n-clear-color-pressed)!important;
 `)]),f("placeholder",`
 display: flex;
 `),f("clear, placeholder",`
 position: absolute;
 left: 50%;
 top: 50%;
 transform: translateX(-50%) translateY(-50%);
 `,[mt({originalTransform:"translateX(-50%) translateY(-50%)",left:"50%",top:"50%"})])])]),Pe=R({name:"BaseClear",props:{clsPrefix:{type:String,required:!0},show:Boolean,onClear:Function},setup(o){return Be("-base-clear",Fr,Se(o,"clsPrefix")),{handleMouseDown(n){n.preventDefault()}}},render(){const{clsPrefix:o}=this;return a("div",{class:`${o}-base-clear`},a(gt,null,{default:()=>{var n,r;return this.show?a("div",{key:"dismiss",class:`${o}-base-clear__clear`,onClick:this.onClear,onMousedown:this.handleMouseDown,"data-clear":!0},G(this.$slots.icon,()=>[a(se,{clsPrefix:o},{default:()=>a(zr,null)})])):a("div",{key:"icon",class:`${o}-base-clear__placeholder`},(r=(n=this.$slots).placeholder)===null||r===void 0?void 0:r.call(n))}}))}}),Tr=R({name:"InternalSelectionSuffix",props:{clsPrefix:{type:String,required:!0},showArrow:{type:Boolean,default:void 0},showClear:{type:Boolean,default:void 0},loading:{type:Boolean,default:!1},onClear:Function},setup(o,{slots:n}){return()=>{const{clsPrefix:r}=o;return a(bt,{clsPrefix:r,class:`${r}-base-suffix`,strokeWidth:24,scale:.85,show:o.loading},{default:()=>o.showArrow?a(Pe,{clsPrefix:r,show:o.showClear,onClear:o.onClear},{placeholder:()=>a(se,{clsPrefix:r,class:`${r}-base-suffix__arrow`},{default:()=>G(n.default,()=>[a(Mr,null)])})}):null})}}}),$r={paddingTiny:"0 8px",paddingSmall:"0 10px",paddingMedium:"0 12px",paddingLarge:"0 14px",clearSize:"16px"},Ar=o=>{const{textColor2:n,textColor3:r,textColorDisabled:p,primaryColor:s,primaryColorHover:l,inputColor:c,inputColorDisabled:d,borderColor:v,warningColor:u,warningColorHover:h,errorColor:m,errorColorHover:y,borderRadius:A,lineHeight:z,fontSizeTiny:de,fontSizeSmall:V,fontSizeMedium:ue,fontSizeLarge:T,heightTiny:k,heightSmall:O,heightMedium:D,heightLarge:he,actionColor:_,clearColor:E,clearColorHover:$,clearColorPressed:W,placeholderColor:j,placeholderColorDisabled:H,iconColor:fe,iconColorDisabled:ve,iconColorHover:U,iconColorPressed:pe}=o;return Object.assign(Object.assign({},$r),{countTextColorDisabled:p,countTextColor:r,heightTiny:k,heightSmall:O,heightMedium:D,heightLarge:he,fontSizeTiny:de,fontSizeSmall:V,fontSizeMedium:ue,fontSizeLarge:T,lineHeight:z,lineHeightTextarea:z,borderRadius:A,iconSize:"16px",groupLabelColor:_,groupLabelTextColor:n,textColor:n,textColorDisabled:p,textDecorationColor:n,caretColor:s,placeholderColor:j,placeholderColorDisabled:H,color:c,colorDisabled:d,colorFocus:c,groupLabelBorder:`1px solid ${v}`,border:`1px solid ${v}`,borderHover:`1px solid ${l}`,borderDisabled:`1px solid ${v}`,borderFocus:`1px solid ${l}`,boxShadowFocus:`0 0 0 2px ${we(s,{alpha:.2})}`,loadingColor:s,loadingColorWarning:u,borderWarning:`1px solid ${u}`,borderHoverWarning:`1px solid ${h}`,colorFocusWarning:c,borderFocusWarning:`1px solid ${h}`,boxShadowFocusWarning:`0 0 0 2px ${we(u,{alpha:.2})}`,caretColorWarning:u,loadingColorError:m,borderError:`1px solid ${m}`,borderHoverError:`1px solid ${y}`,colorFocusError:c,borderFocusError:`1px solid ${y}`,boxShadowFocusError:`0 0 0 2px ${we(m,{alpha:.2})}`,caretColorError:m,clearColor:E,clearColorHover:$,clearColorPressed:W,iconColor:fe,iconColorDisabled:ve,iconColorHover:U,iconColorPressed:pe,suffixTextColor:n})},Dr={name:"Input",common:Ie,self:Ar},_r=Dr,Ne=yt("n-input");function Rr(o){let n=0;for(const r of o)n++;return n}function ie(o){return o===""||o==null}function kr(o){const n=C(null);function r(){const{value:l}=o;if(!(l!=null&&l.focus)){s();return}const{selectionStart:c,selectionEnd:d,value:v}=l;if(c==null||d==null){s();return}n.value={start:c,end:d,beforeText:v.slice(0,c),afterText:v.slice(d)}}function p(){var l;const{value:c}=n,{value:d}=o;if(!c||!d)return;const{value:v}=d,{start:u,beforeText:h,afterText:m}=c;let y=v.length;if(v.endsWith(m))y=v.length-m.length;else if(v.startsWith(h))y=h.length;else{const A=h[u-1],z=v.indexOf(A,u-1);z!==-1&&(y=z+1)}(l=d.setSelectionRange)===null||l===void 0||l.call(d,y,y)}function s(){n.value=null}return le(o,s),{recordCursor:r,restoreCursor:p}}const Ee=R({name:"InputWordCount",setup(o,{slots:n}){const{mergedValueRef:r,maxlengthRef:p,mergedClsPrefixRef:s,countGraphemesRef:l}=We(Ne),c=P(()=>{const{value:d}=r;return d===null||Array.isArray(d)?0:(l.value||Rr)(d)});return()=>{const{value:d}=p,{value:v}=r;return a("span",{class:`${s.value}-input-word-count`},xt(n.default,{value:v===null||Array.isArray(v)?"":v},()=>[d===void 0?c.value:`${c.value} / ${d}`]))}}}),Er=w("input",`
 max-width: 100%;
 cursor: text;
 line-height: 1.5;
 z-index: auto;
 outline: none;
 box-sizing: border-box;
 position: relative;
 display: inline-flex;
 border-radius: var(--n-border-radius);
 background-color: var(--n-color);
 transition: background-color .3s var(--n-bezier);
 font-size: var(--n-font-size);
 --n-padding-vertical: calc((var(--n-height) - 1.5 * var(--n-font-size)) / 2);
`,[f("input, textarea",`
 overflow: hidden;
 flex-grow: 1;
 position: relative;
 `),f("input-el, textarea-el, input-mirror, textarea-mirror, separator, placeholder",`
 box-sizing: border-box;
 font-size: inherit;
 line-height: 1.5;
 font-family: inherit;
 border: none;
 outline: none;
 background-color: #0000;
 text-align: inherit;
 transition:
 -webkit-text-fill-color .3s var(--n-bezier),
 caret-color .3s var(--n-bezier),
 color .3s var(--n-bezier),
 text-decoration-color .3s var(--n-bezier);
 `),f("input-el, textarea-el",`
 -webkit-appearance: none;
 scrollbar-width: none;
 width: 100%;
 min-width: 0;
 text-decoration-color: var(--n-text-decoration-color);
 color: var(--n-text-color);
 caret-color: var(--n-caret-color);
 background-color: transparent;
 `,[S("&::-webkit-scrollbar, &::-webkit-scrollbar-track-piece, &::-webkit-scrollbar-thumb",`
 width: 0;
 height: 0;
 display: none;
 `),S("&::placeholder",`
 color: #0000;
 -webkit-text-fill-color: transparent !important;
 `),S("&:-webkit-autofill ~",[f("placeholder","display: none;")])]),F("round",[Z("textarea","border-radius: calc(var(--n-height) / 2);")]),f("placeholder",`
 pointer-events: none;
 position: absolute;
 left: 0;
 right: 0;
 top: 0;
 bottom: 0;
 overflow: hidden;
 color: var(--n-placeholder-color);
 `,[S("span",`
 width: 100%;
 display: inline-block;
 `)]),F("textarea",[f("placeholder","overflow: visible;")]),Z("autosize","width: 100%;"),F("autosize",[f("textarea-el, input-el",`
 position: absolute;
 top: 0;
 left: 0;
 height: 100%;
 `)]),w("input-wrapper",`
 overflow: hidden;
 display: inline-flex;
 flex-grow: 1;
 position: relative;
 padding-left: var(--n-padding-left);
 padding-right: var(--n-padding-right);
 `),f("input-mirror",`
 padding: 0;
 height: var(--n-height);
 line-height: var(--n-height);
 overflow: hidden;
 visibility: hidden;
 position: static;
 white-space: pre;
 pointer-events: none;
 `),f("input-el",`
 padding: 0;
 height: var(--n-height);
 line-height: var(--n-height);
 `,[S("+",[f("placeholder",`
 display: flex;
 align-items: center; 
 `)])]),Z("textarea",[f("placeholder","white-space: nowrap;")]),f("eye",`
 transition: color .3s var(--n-bezier);
 `),F("textarea","width: 100%;",[w("input-word-count",`
 position: absolute;
 right: var(--n-padding-right);
 bottom: var(--n-padding-vertical);
 `),F("resizable",[w("input-wrapper",`
 resize: vertical;
 min-height: var(--n-height);
 `)]),f("textarea-el, textarea-mirror, placeholder",`
 height: 100%;
 padding-left: 0;
 padding-right: 0;
 padding-top: var(--n-padding-vertical);
 padding-bottom: var(--n-padding-vertical);
 word-break: break-word;
 display: inline-block;
 vertical-align: bottom;
 box-sizing: border-box;
 line-height: var(--n-line-height-textarea);
 margin: 0;
 resize: none;
 white-space: pre-wrap;
 `),f("textarea-mirror",`
 width: 100%;
 pointer-events: none;
 overflow: hidden;
 visibility: hidden;
 position: static;
 white-space: pre-wrap;
 overflow-wrap: break-word;
 `)]),F("pair",[f("input-el, placeholder","text-align: center;"),f("separator",`
 display: flex;
 align-items: center;
 transition: color .3s var(--n-bezier);
 color: var(--n-text-color);
 white-space: nowrap;
 `,[w("icon",`
 color: var(--n-icon-color);
 `),w("base-icon",`
 color: var(--n-icon-color);
 `)])]),F("disabled",`
 cursor: not-allowed;
 background-color: var(--n-color-disabled);
 `,[f("border","border: var(--n-border-disabled);"),f("input-el, textarea-el",`
 cursor: not-allowed;
 color: var(--n-text-color-disabled);
 text-decoration-color: var(--n-text-color-disabled);
 `),f("placeholder","color: var(--n-placeholder-color-disabled);"),f("separator","color: var(--n-text-color-disabled);",[w("icon",`
 color: var(--n-icon-color-disabled);
 `),w("base-icon",`
 color: var(--n-icon-color-disabled);
 `)]),w("input-word-count",`
 color: var(--n-count-text-color-disabled);
 `),f("suffix, prefix","color: var(--n-text-color-disabled);",[w("icon",`
 color: var(--n-icon-color-disabled);
 `),w("internal-icon",`
 color: var(--n-icon-color-disabled);
 `)])]),Z("disabled",[f("eye",`
 display: flex;
 align-items: center;
 justify-content: center;
 color: var(--n-icon-color);
 cursor: pointer;
 `,[S("&:hover",`
 color: var(--n-icon-color-hover);
 `),S("&:active",`
 color: var(--n-icon-color-pressed);
 `)]),S("&:hover",[f("state-border","border: var(--n-border-hover);")]),F("focus","background-color: var(--n-color-focus);",[f("state-border",`
 border: var(--n-border-focus);
 box-shadow: var(--n-box-shadow-focus);
 `)])]),f("border, state-border",`
 box-sizing: border-box;
 position: absolute;
 left: 0;
 right: 0;
 top: 0;
 bottom: 0;
 pointer-events: none;
 border-radius: inherit;
 border: var(--n-border);
 transition:
 box-shadow .3s var(--n-bezier),
 border-color .3s var(--n-bezier);
 `),f("state-border",`
 border-color: #0000;
 z-index: 1;
 `),f("prefix","margin-right: 4px;"),f("suffix",`
 margin-left: 4px;
 `),f("suffix, prefix",`
 transition: color .3s var(--n-bezier);
 flex-wrap: nowrap;
 flex-shrink: 0;
 line-height: var(--n-height);
 white-space: nowrap;
 display: inline-flex;
 align-items: center;
 justify-content: center;
 color: var(--n-suffix-text-color);
 `,[w("base-loading",`
 font-size: var(--n-icon-size);
 margin: 0 2px;
 color: var(--n-loading-color);
 `),w("base-clear",`
 font-size: var(--n-icon-size);
 `,[f("placeholder",[w("base-icon",`
 transition: color .3s var(--n-bezier);
 color: var(--n-icon-color);
 font-size: var(--n-icon-size);
 `)])]),S(">",[w("icon",`
 transition: color .3s var(--n-bezier);
 color: var(--n-icon-color);
 font-size: var(--n-icon-size);
 `)]),w("base-icon",`
 font-size: var(--n-icon-size);
 `)]),w("input-word-count",`
 pointer-events: none;
 line-height: 1.5;
 font-size: .85em;
 color: var(--n-count-text-color);
 transition: color .3s var(--n-bezier);
 margin-left: 4px;
 font-variant: tabular-nums;
 `),["warning","error"].map(o=>F(`${o}-status`,[Z("disabled",[w("base-loading",`
 color: var(--n-loading-color-${o})
 `),f("input-el, textarea-el",`
 caret-color: var(--n-caret-color-${o});
 `),f("state-border",`
 border: var(--n-border-${o});
 `),S("&:hover",[f("state-border",`
 border: var(--n-border-hover-${o});
 `)]),S("&:focus",`
 background-color: var(--n-color-focus-${o});
 `,[f("state-border",`
 box-shadow: var(--n-box-shadow-focus-${o});
 border: var(--n-border-focus-${o});
 `)]),F("focus",`
 background-color: var(--n-color-focus-${o});
 `,[f("state-border",`
 box-shadow: var(--n-box-shadow-focus-${o});
 border: var(--n-border-focus-${o});
 `)])])]))]),Wr=w("input",[F("disabled",[f("input-el, textarea-el",`
 -webkit-text-fill-color: var(--n-text-color-disabled);
 `)])]),Br=Object.assign(Object.assign({},ce.props),{bordered:{type:Boolean,default:void 0},type:{type:String,default:"text"},placeholder:[Array,String],defaultValue:{type:[String,Array],default:null},value:[String,Array],disabled:{type:Boolean,default:void 0},size:String,rows:{type:[Number,String],default:3},round:Boolean,minlength:[String,Number],maxlength:[String,Number],clearable:Boolean,autosize:{type:[Boolean,Object],default:!1},pair:Boolean,separator:String,readonly:{type:[String,Boolean],default:!1},passivelyActivated:Boolean,showPasswordOn:String,stateful:{type:Boolean,default:!0},autofocus:Boolean,inputProps:Object,resizable:{type:Boolean,default:!0},showCount:Boolean,loading:{type:Boolean,default:void 0},allowInput:Function,renderCount:Function,onMousedown:Function,onKeydown:Function,onKeyup:Function,onInput:[Function,Array],onFocus:[Function,Array],onBlur:[Function,Array],onClick:[Function,Array],onChange:[Function,Array],onClear:[Function,Array],countGraphemes:Function,status:String,"onUpdate:value":[Function,Array],onUpdateValue:[Function,Array],textDecoration:[String,Array],attrSize:{type:Number,default:20},onInputBlur:[Function,Array],onInputFocus:[Function,Array],onDeactivate:[Function,Array],onActivate:[Function,Array],onWrapperFocus:[Function,Array],onWrapperBlur:[Function,Array],internalDeactivateOnEnter:Boolean,internalForceFocus:Boolean,internalLoadingBeforeSuffix:Boolean,showPasswordToggle:Boolean}),Ur=R({name:"Input",props:Br,setup(o){const{mergedClsPrefixRef:n,mergedBorderedRef:r,inlineThemeDisabled:p,mergedRtlRef:s}=Le(o),l=ce("Input","-input",Er,_r,o,n);wt&&Be("-input-safari",Wr,n);const c=C(null),d=C(null),v=C(null),u=C(null),h=C(null),m=C(null),y=C(null),A=kr(y),z=C(null),{localeRef:de}=Cr("Input"),V=C(o.defaultValue),ue=Se(o,"value"),T=Et(ue,V),k=Ct(o),{mergedSizeRef:O,mergedDisabledRef:D,mergedStatusRef:he}=k,_=C(!1),E=C(!1),$=C(!1),W=C(!1);let j=null;const H=P(()=>{const{placeholder:e,pair:t}=o;return t?Array.isArray(e)?e:e===void 0?["",""]:[e,e]:e===void 0?[de.value.placeholder]:[e]}),fe=P(()=>{const{value:e}=$,{value:t}=T,{value:i}=H;return!e&&(ie(t)||Array.isArray(t)&&ie(t[0]))&&i[0]}),ve=P(()=>{const{value:e}=$,{value:t}=T,{value:i}=H;return!e&&i[1]&&(ie(t)||Array.isArray(t)&&ie(t[1]))}),U=$e(()=>o.internalForceFocus||_.value),pe=$e(()=>{if(D.value||o.readonly||!o.clearable||!U.value&&!E.value)return!1;const{value:e}=T,{value:t}=U;return o.pair?!!(Array.isArray(e)&&(e[0]||e[1]))&&(E.value||t):!!e&&(E.value||t)}),me=P(()=>{const{showPasswordOn:e}=o;if(e)return e;if(o.showPasswordToggle)return"click"}),q=C(!1),Oe=P(()=>{const{textDecoration:e}=o;return e?Array.isArray(e)?e.map(t=>({textDecoration:t})):[{textDecoration:e}]:["",""]}),Me=C(void 0),je=()=>{var e,t;if(o.type==="textarea"){const{autosize:i}=o;if(i&&(Me.value=(t=(e=z.value)===null||e===void 0?void 0:e.$el)===null||t===void 0?void 0:t.offsetWidth),!d.value||typeof i=="boolean")return;const{paddingTop:b,paddingBottom:x,lineHeight:g}=window.getComputedStyle(d.value),B=Number(b.slice(0,-2)),I=Number(x.slice(0,-2)),L=Number(g.slice(0,-2)),{value:K}=v;if(!K)return;if(i.minRows){const X=Math.max(i.minRows,1),ye=`${B+I+L*X}px`;K.style.minHeight=ye}if(i.maxRows){const X=`${B+I+L*i.maxRows}px`;K.style.maxHeight=X}}},He=P(()=>{const{maxlength:e}=o;return e===void 0?void 0:Number(e)});St(()=>{const{value:e}=T;Array.isArray(e)||be(e)});const Ue=Pt().proxy;function Q(e){const{onUpdateValue:t,"onUpdate:value":i,onInput:b}=o,{nTriggerFormInput:x}=k;t&&M(t,e),i&&M(i,e),b&&M(b,e),V.value=e,x()}function ee(e){const{onChange:t}=o,{nTriggerFormChange:i}=k;t&&M(t,e),V.value=e,i()}function qe(e){const{onBlur:t}=o,{nTriggerFormBlur:i}=k;t&&M(t,e),i()}function Ke(e){const{onFocus:t}=o,{nTriggerFormFocus:i}=k;t&&M(t,e),i()}function Xe(e){const{onClear:t}=o;t&&M(t,e)}function Ye(e){const{onInputBlur:t}=o;t&&M(t,e)}function Je(e){const{onInputFocus:t}=o;t&&M(t,e)}function Ze(){const{onDeactivate:e}=o;e&&M(e)}function Ge(){const{onActivate:e}=o;e&&M(e)}function Qe(e){const{onClick:t}=o;t&&M(t,e)}function eo(e){const{onWrapperFocus:t}=o;t&&M(t,e)}function oo(e){const{onWrapperBlur:t}=o;t&&M(t,e)}function to(){$.value=!0}function ro(e){$.value=!1,e.target===m.value?oe(e,1):oe(e,0)}function oe(e,t=0,i="input"){const b=e.target.value;if(be(b),e instanceof InputEvent&&!e.isComposing&&($.value=!1),o.type==="textarea"){const{value:g}=z;g&&g.syncUnifiedContainer()}if(j=b,$.value)return;A.recordCursor();const x=no(b);if(x)if(!o.pair)i==="input"?Q(b):ee(b);else{let{value:g}=T;Array.isArray(g)?g=[g[0],g[1]]:g=["",""],g[t]=b,i==="input"?Q(g):ee(g)}Ue.$forceUpdate(),x||De(A.restoreCursor)}function no(e){const{countGraphemes:t,maxlength:i,minlength:b}=o;if(t){let g;if(i!==void 0&&(g===void 0&&(g=t(e)),g>Number(i))||b!==void 0&&(g===void 0&&(g=t(e)),g<Number(i)))return!1}const{allowInput:x}=o;return typeof x=="function"?x(e):!0}function ao(e){Ye(e),e.relatedTarget===c.value&&Ze(),e.relatedTarget!==null&&(e.relatedTarget===h.value||e.relatedTarget===m.value||e.relatedTarget===d.value)||(W.value=!1),te(e,"blur"),y.value=null}function io(e,t){Je(e),_.value=!0,W.value=!0,Ge(),te(e,"focus"),t===0?y.value=h.value:t===1?y.value=m.value:t===2&&(y.value=d.value)}function lo(e){o.passivelyActivated&&(oo(e),te(e,"blur"))}function so(e){o.passivelyActivated&&(_.value=!0,eo(e),te(e,"focus"))}function te(e,t){e.relatedTarget!==null&&(e.relatedTarget===h.value||e.relatedTarget===m.value||e.relatedTarget===d.value||e.relatedTarget===c.value)||(t==="focus"?(Ke(e),_.value=!0):t==="blur"&&(qe(e),_.value=!1))}function co(e,t){oe(e,t,"change")}function uo(e){Qe(e)}function ho(e){Xe(e),o.pair?(Q(["",""]),ee(["",""])):(Q(""),ee(""))}function fo(e){const{onMousedown:t}=o;t&&t(e);const{tagName:i}=e.target;if(i!=="INPUT"&&i!=="TEXTAREA"){if(o.resizable){const{value:b}=c;if(b){const{left:x,top:g,width:B,height:I}=b.getBoundingClientRect(),L=14;if(x+B-L<e.clientX&&e.clientX<x+B&&g+I-L<e.clientY&&e.clientY<g+I)return}}e.preventDefault(),_.value||ze()}}function vo(){var e;E.value=!0,o.type==="textarea"&&((e=z.value)===null||e===void 0||e.handleMouseEnterWrapper())}function po(){var e;E.value=!1,o.type==="textarea"&&((e=z.value)===null||e===void 0||e.handleMouseLeaveWrapper())}function mo(){D.value||me.value==="click"&&(q.value=!q.value)}function go(e){if(D.value)return;e.preventDefault();const t=b=>{b.preventDefault(),Re("mouseup",document,t)};if(_e("mouseup",document,t),me.value!=="mousedown")return;q.value=!0;const i=()=>{q.value=!1,Re("mouseup",document,i)};_e("mouseup",document,i)}function bo(e){var t;switch((t=o.onKeydown)===null||t===void 0||t.call(o,e),e.key){case"Escape":ge();break;case"Enter":yo(e);break}}function yo(e){var t,i;if(o.passivelyActivated){const{value:b}=W;if(b){o.internalDeactivateOnEnter&&ge();return}e.preventDefault(),o.type==="textarea"?(t=d.value)===null||t===void 0||t.focus():(i=h.value)===null||i===void 0||i.focus()}}function ge(){o.passivelyActivated&&(W.value=!1,De(()=>{var e;(e=c.value)===null||e===void 0||e.focus()}))}function ze(){var e,t,i;D.value||(o.passivelyActivated?(e=c.value)===null||e===void 0||e.focus():((t=d.value)===null||t===void 0||t.focus(),(i=h.value)===null||i===void 0||i.focus()))}function xo(){var e;!((e=c.value)===null||e===void 0)&&e.contains(document.activeElement)&&document.activeElement.blur()}function wo(){var e,t;(e=d.value)===null||e===void 0||e.select(),(t=h.value)===null||t===void 0||t.select()}function Co(){D.value||(d.value?d.value.focus():h.value&&h.value.focus())}function So(){const{value:e}=c;e!=null&&e.contains(document.activeElement)&&e!==document.activeElement&&ge()}function Po(e){if(o.type==="textarea"){const{value:t}=d;t==null||t.scrollTo(e)}else{const{value:t}=h;t==null||t.scrollTo(e)}}function be(e){const{type:t,pair:i,autosize:b}=o;if(!i&&b)if(t==="textarea"){const{value:x}=v;x&&(x.textContent=(e??"")+`\r
`)}else{const{value:x}=u;x&&(e?x.textContent=e:x.innerHTML="&nbsp;")}}function Mo(){je()}const Fe=C({top:"0"});function zo(e){var t;const{scrollTop:i}=e.target;Fe.value.top=`${-i}px`,(t=z.value)===null||t===void 0||t.syncUnifiedContainer()}let re=null;Ae(()=>{const{autosize:e,type:t}=o;e&&t==="textarea"?re=le(T,i=>{!Array.isArray(i)&&i!==j&&be(i)}):re==null||re()});let ne=null;Ae(()=>{o.type==="textarea"?ne=le(T,e=>{var t;!Array.isArray(e)&&e!==j&&((t=z.value)===null||t===void 0||t.syncUnifiedContainer())}):ne==null||ne()}),Mt(Ne,{mergedValueRef:T,maxlengthRef:He,mergedClsPrefixRef:n,countGraphemesRef:Se(o,"countGraphemes")});const Fo={wrapperElRef:c,inputElRef:h,textareaElRef:d,isCompositing:$,focus:ze,blur:xo,select:wo,deactivate:So,activate:Co,scrollTo:Po},To=zt("Input",s,n),Te=P(()=>{const{value:e}=O,{common:{cubicBezierEaseInOut:t},self:{color:i,borderRadius:b,textColor:x,caretColor:g,caretColorError:B,caretColorWarning:I,textDecorationColor:L,border:K,borderDisabled:X,borderHover:ye,borderFocus:$o,placeholderColor:Ao,placeholderColorDisabled:Do,lineHeightTextarea:_o,colorDisabled:Ro,colorFocus:ko,textColorDisabled:Eo,boxShadowFocus:Wo,iconSize:Bo,colorFocusWarning:Io,boxShadowFocusWarning:Lo,borderWarning:Vo,borderFocusWarning:No,borderHoverWarning:Oo,colorFocusError:jo,boxShadowFocusError:Ho,borderError:Uo,borderFocusError:qo,borderHoverError:Ko,clearSize:Xo,clearColor:Yo,clearColorHover:Jo,clearColorPressed:Zo,iconColor:Go,iconColorDisabled:Qo,suffixTextColor:et,countTextColor:ot,countTextColorDisabled:tt,iconColorHover:rt,iconColorPressed:nt,loadingColor:at,loadingColorError:it,loadingColorWarning:lt,[Ce("padding",e)]:st,[Ce("fontSize",e)]:ct,[Ce("height",e)]:dt}}=l.value,{left:ut,right:ht}=At(st);return{"--n-bezier":t,"--n-count-text-color":ot,"--n-count-text-color-disabled":tt,"--n-color":i,"--n-font-size":ct,"--n-border-radius":b,"--n-height":dt,"--n-padding-left":ut,"--n-padding-right":ht,"--n-text-color":x,"--n-caret-color":g,"--n-text-decoration-color":L,"--n-border":K,"--n-border-disabled":X,"--n-border-hover":ye,"--n-border-focus":$o,"--n-placeholder-color":Ao,"--n-placeholder-color-disabled":Do,"--n-icon-size":Bo,"--n-line-height-textarea":_o,"--n-color-disabled":Ro,"--n-color-focus":ko,"--n-text-color-disabled":Eo,"--n-box-shadow-focus":Wo,"--n-loading-color":at,"--n-caret-color-warning":I,"--n-color-focus-warning":Io,"--n-box-shadow-focus-warning":Lo,"--n-border-warning":Vo,"--n-border-focus-warning":No,"--n-border-hover-warning":Oo,"--n-loading-color-warning":lt,"--n-caret-color-error":B,"--n-color-focus-error":jo,"--n-box-shadow-focus-error":Ho,"--n-border-error":Uo,"--n-border-focus-error":qo,"--n-border-hover-error":Ko,"--n-loading-color-error":it,"--n-clear-color":Yo,"--n-clear-size":Xo,"--n-clear-color-hover":Jo,"--n-clear-color-pressed":Zo,"--n-icon-color":Go,"--n-icon-color-hover":rt,"--n-icon-color-pressed":nt,"--n-icon-color-disabled":Qo,"--n-suffix-text-color":et}}),N=p?Ve("input",P(()=>{const{value:e}=O;return e[0]}),Te,o):void 0;return Object.assign(Object.assign({},Fo),{wrapperElRef:c,inputElRef:h,inputMirrorElRef:u,inputEl2Ref:m,textareaElRef:d,textareaMirrorElRef:v,textareaScrollbarInstRef:z,rtlEnabled:To,uncontrolledValue:V,mergedValue:T,passwordVisible:q,mergedPlaceholder:H,showPlaceholder1:fe,showPlaceholder2:ve,mergedFocus:U,isComposing:$,activated:W,showClearButton:pe,mergedSize:O,mergedDisabled:D,textDecorationStyle:Oe,mergedClsPrefix:n,mergedBordered:r,mergedShowPasswordOn:me,placeholderStyle:Fe,mergedStatus:he,textAreaScrollContainerWidth:Me,handleTextAreaScroll:zo,handleCompositionStart:to,handleCompositionEnd:ro,handleInput:oe,handleInputBlur:ao,handleInputFocus:io,handleWrapperBlur:lo,handleWrapperFocus:so,handleMouseEnter:vo,handleMouseLeave:po,handleMouseDown:fo,handleChange:co,handleClick:uo,handleClear:ho,handlePasswordToggleClick:mo,handlePasswordToggleMousedown:go,handleWrapperKeydown:bo,handleTextAreaMirrorResize:Mo,getTextareaScrollContainer:()=>d.value,mergedTheme:l,cssVars:p?void 0:Te,themeClass:N==null?void 0:N.themeClass,onRender:N==null?void 0:N.onRender})},render(){var o,n;const{mergedClsPrefix:r,mergedStatus:p,themeClass:s,type:l,countGraphemes:c,onRender:d}=this,v=this.$slots;return d==null||d(),a("div",{ref:"wrapperElRef",class:[`${r}-input`,s,p&&`${r}-input--${p}-status`,{[`${r}-input--rtl`]:this.rtlEnabled,[`${r}-input--disabled`]:this.mergedDisabled,[`${r}-input--textarea`]:l==="textarea",[`${r}-input--resizable`]:this.resizable&&!this.autosize,[`${r}-input--autosize`]:this.autosize,[`${r}-input--round`]:this.round&&l!=="textarea",[`${r}-input--pair`]:this.pair,[`${r}-input--focus`]:this.mergedFocus,[`${r}-input--stateful`]:this.stateful}],style:this.cssVars,tabindex:!this.mergedDisabled&&this.passivelyActivated&&!this.activated?0:void 0,onFocus:this.handleWrapperFocus,onBlur:this.handleWrapperBlur,onClick:this.handleClick,onMousedown:this.handleMouseDown,onMouseenter:this.handleMouseEnter,onMouseleave:this.handleMouseLeave,onCompositionstart:this.handleCompositionStart,onCompositionend:this.handleCompositionEnd,onKeyup:this.onKeyup,onKeydown:this.handleWrapperKeydown},a("div",{class:`${r}-input-wrapper`},ae(v.prefix,u=>u&&a("div",{class:`${r}-input__prefix`},u)),l==="textarea"?a(Ft,{ref:"textareaScrollbarInstRef",class:`${r}-input__textarea`,container:this.getTextareaScrollContainer,triggerDisplayManually:!0,useUnifiedContainer:!0,internalHoistYRail:!0},{default:()=>{var u,h;const{textAreaScrollContainerWidth:m}=this,y={width:this.autosize&&m&&`${m}px`};return a(Tt,null,a("textarea",Object.assign({},this.inputProps,{ref:"textareaElRef",class:[`${r}-input__textarea-el`,(u=this.inputProps)===null||u===void 0?void 0:u.class],autofocus:this.autofocus,rows:Number(this.rows),placeholder:this.placeholder,value:this.mergedValue,disabled:this.mergedDisabled,maxlength:c?void 0:this.maxlength,minlength:c?void 0:this.minlength,readonly:this.readonly,tabindex:this.passivelyActivated&&!this.activated?-1:void 0,style:[this.textDecorationStyle[0],(h=this.inputProps)===null||h===void 0?void 0:h.style,y],onBlur:this.handleInputBlur,onFocus:A=>this.handleInputFocus(A,2),onInput:this.handleInput,onChange:this.handleChange,onScroll:this.handleTextAreaScroll})),this.showPlaceholder1?a("div",{class:`${r}-input__placeholder`,style:[this.placeholderStyle,y],key:"placeholder"},this.mergedPlaceholder[0]):null,this.autosize?a($t,{onResize:this.handleTextAreaMirrorResize},{default:()=>a("div",{ref:"textareaMirrorElRef",class:`${r}-input__textarea-mirror`,key:"mirror"})}):null)}}):a("div",{class:`${r}-input__input`},a("input",Object.assign({type:l==="password"&&this.mergedShowPasswordOn&&this.passwordVisible?"text":l},this.inputProps,{ref:"inputElRef",class:[`${r}-input__input-el`,(o=this.inputProps)===null||o===void 0?void 0:o.class],style:[this.textDecorationStyle[0],(n=this.inputProps)===null||n===void 0?void 0:n.style],tabindex:this.passivelyActivated&&!this.activated?-1:void 0,placeholder:this.mergedPlaceholder[0],disabled:this.mergedDisabled,maxlength:c?void 0:this.maxlength,minlength:c?void 0:this.minlength,value:Array.isArray(this.mergedValue)?this.mergedValue[0]:this.mergedValue,readonly:this.readonly,autofocus:this.autofocus,size:this.attrSize,onBlur:this.handleInputBlur,onFocus:u=>this.handleInputFocus(u,0),onInput:u=>this.handleInput(u,0),onChange:u=>this.handleChange(u,0)})),this.showPlaceholder1?a("div",{class:`${r}-input__placeholder`},a("span",null,this.mergedPlaceholder[0])):null,this.autosize?a("div",{class:`${r}-input__input-mirror`,key:"mirror",ref:"inputMirrorElRef"}," "):null),!this.pair&&ae(v.suffix,u=>u||this.clearable||this.showCount||this.mergedShowPasswordOn||this.loading!==void 0?a("div",{class:`${r}-input__suffix`},[ae(v["clear-icon-placeholder"],h=>(this.clearable||h)&&a(Pe,{clsPrefix:r,show:this.showClearButton,onClear:this.handleClear},{placeholder:()=>h,icon:()=>{var m,y;return(y=(m=this.$slots)["clear-icon"])===null||y===void 0?void 0:y.call(m)}})),this.internalLoadingBeforeSuffix?null:u,this.loading!==void 0?a(Tr,{clsPrefix:r,loading:this.loading,showArrow:!1,showClear:!1,style:this.cssVars}):null,this.internalLoadingBeforeSuffix?u:null,this.showCount&&this.type!=="textarea"?a(Ee,null,{default:h=>{var m;return(m=v.count)===null||m===void 0?void 0:m.call(v,h)}}):null,this.mergedShowPasswordOn&&this.type==="password"?a("div",{class:`${r}-input__eye`,onMousedown:this.handlePasswordToggleMousedown,onClick:this.handlePasswordToggleClick},this.passwordVisible?G(v["password-visible-icon"],()=>[a(se,{clsPrefix:r},{default:()=>a(Sr,null)})]):G(v["password-invisible-icon"],()=>[a(se,{clsPrefix:r},{default:()=>a(Pr,null)})])):null]):null)),this.pair?a("span",{class:`${r}-input__separator`},G(v.separator,()=>[this.separator])):null,this.pair?a("div",{class:`${r}-input-wrapper`},a("div",{class:`${r}-input__input`},a("input",{ref:"inputEl2Ref",type:this.type,class:`${r}-input__input-el`,tabindex:this.passivelyActivated&&!this.activated?-1:void 0,placeholder:this.mergedPlaceholder[1],disabled:this.mergedDisabled,maxlength:c?void 0:this.maxlength,minlength:c?void 0:this.minlength,value:Array.isArray(this.mergedValue)?this.mergedValue[1]:void 0,readonly:this.readonly,style:this.textDecorationStyle[1],onBlur:this.handleInputBlur,onFocus:u=>this.handleInputFocus(u,1),onInput:u=>this.handleInput(u,1),onChange:u=>this.handleChange(u,1)}),this.showPlaceholder2?a("div",{class:`${r}-input__placeholder`},a("span",null,this.mergedPlaceholder[1])):null),ae(v.suffix,u=>(this.clearable||u)&&a("div",{class:`${r}-input__suffix`},[this.clearable&&a(Pe,{clsPrefix:r,show:this.showClearButton,onClear:this.handleClear},{icon:()=>{var h;return(h=v["clear-icon"])===null||h===void 0?void 0:h.call(v)},placeholder:()=>{var h;return(h=v["clear-icon-placeholder"])===null||h===void 0?void 0:h.call(v)}}),u]))):null,this.mergedBordered?a("div",{class:`${r}-input__border`}):null,this.mergedBordered?a("div",{class:`${r}-input__state-border`}):null,this.showCount&&l==="textarea"?a(Ee,null,{default:u=>{var h;const{renderCount:m}=this;return m?m(u):(h=v.count)===null||h===void 0?void 0:h.call(v,u)}}):null)}}),Ir=o=>{const{textColorBase:n,opacity1:r,opacity2:p,opacity3:s,opacity4:l,opacity5:c}=o;return{color:n,opacity1Depth:r,opacity2Depth:p,opacity3Depth:s,opacity4Depth:l,opacity5Depth:c}},Lr={name:"Icon",common:Ie,self:Ir},Vr=Lr,Nr=w("icon",`
 height: 1em;
 width: 1em;
 line-height: 1em;
 text-align: center;
 display: inline-block;
 position: relative;
 fill: currentColor;
 transform: translateZ(0);
`,[F("color-transition",{transition:"color .3s var(--n-bezier)"}),F("depth",{color:"var(--n-color)"},[S("svg",{opacity:"var(--n-opacity)",transition:"opacity .3s var(--n-bezier)"})]),S("svg",{height:"1em",width:"1em"})]),Or=Object.assign(Object.assign({},ce.props),{depth:[String,Number],size:[Number,String],color:String,component:Object}),qr=R({_n_icon__:!0,name:"Icon",inheritAttrs:!1,props:Or,setup(o){const{mergedClsPrefixRef:n,inlineThemeDisabled:r}=Le(o),p=ce("Icon","-icon",Nr,Vr,o,n),s=P(()=>{const{depth:c}=o,{common:{cubicBezierEaseInOut:d},self:v}=p.value;if(c!==void 0){const{color:u,[`opacity${c}Depth`]:h}=v;return{"--n-bezier":d,"--n-color":u,"--n-opacity":h}}return{"--n-bezier":d,"--n-color":"","--n-opacity":""}}),l=r?Ve("icon",P(()=>`${o.depth||"d"}`),s,o):void 0;return{mergedClsPrefix:n,mergedStyle:P(()=>{const{size:c,color:d}=o;return{fontSize:kt(c),color:d}}),cssVars:r?void 0:s,themeClass:l==null?void 0:l.themeClass,onRender:l==null?void 0:l.onRender}},render(){var o;const{$parent:n,depth:r,mergedClsPrefix:p,component:s,onRender:l,themeClass:c}=this;return!((o=n==null?void 0:n.$options)===null||o===void 0)&&o._n_icon__&&Dt("icon","don't wrap `n-icon` inside `n-icon`"),l==null||l(),a("i",_t(this.$attrs,{role:"img",class:[`${p}-icon`,c,{[`${p}-icon--depth`]:r,[`${p}-icon--color-transition`]:r!==void 0}],style:[this.cssVars,this.mergedStyle]}),s?a(s):this.$slots)}});export{Mr as C,qr as N,Hr as a,Ur as b,Cr as c,Tr as d,yr as e,kt as f,_r as i,Et as u};
