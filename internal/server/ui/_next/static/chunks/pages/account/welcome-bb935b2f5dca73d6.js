(self.webpackChunk_N_E=self.webpackChunk_N_E||[]).push([[749],{209:function(n,e,t){(window.__NEXT_P=window.__NEXT_P||[]).push(["/account/welcome",function(){return t(5786)}])},4240:function(n,e,t){"use strict";var r=t(5893),i=t(2125),o=t(5697),c=t.n(o);function a(n,e){return e||(e=n.slice(0)),Object.freeze(Object.defineProperties(n,{raw:{value:Object.freeze(e)}}))}function u(){var n=a(["\n  text-align: center;\n"]);return u=function(){return n},n}function s(){var n=a(["\n  font-weight: 400;\n  font-size: 22px;\n  line-height: 27px;\n  text-align: center;\n  letter-spacing: -0.035em;\n"]);return s=function(){return n},n}function f(){var n=a(["\n  font-weight: 400;\n  font-size: 11px;\n  line-height: 156.52%;\n  opacity: .5;\n  text-align: center;\n  padding: .5rem .5rem 1rem .5rem;\n"]);return f=function(){return n},n}function l(){var n=a(["\n  font-weight: 700;\n  font-size: 11px;\n  line-height: 156.52%;\n  text-align: center;\n  padding: 1rem .5rem 0rem;\n"]);return l=function(){return n},n}var p=i.ZP.div.withConfig({componentId:"sc-90b1100a-0"})(u()),d=i.ZP.div.withConfig({componentId:"sc-90b1100a-1"})(s()),h=i.ZP.div.withConfig({componentId:"sc-90b1100a-2"})(f()),g=i.ZP.div.withConfig({componentId:"sc-90b1100a-3"})(l()),v=function(n){var e=n.header,t=n.subheader,i=n.title;return(0,r.jsxs)(r.Fragment,{children:[(0,r.jsx)(p,{children:(0,r.jsx)("img",{src:"/infra-icon.svg"})}),(0,r.jsx)(d,{children:e}),i&&(0,r.jsx)(g,{children:i}),(0,r.jsx)(h,{children:t})]})};v.prototype={header:c().string,title:c().string,subheader:c().string},e.Z=v},3206:function(n,e,t){"use strict";var r=t(5893),i=t(2125),o=t(5697),c=t.n(o);function a(){var n,e,t=(n=["\n  width: ",";\n  height: ",";\n  background: linear-gradient(266.64deg, #CB56FF -53.31%, #4EB2F4 93.79%);\n  border-radius: 2px;\n  border: none;\n  color: #ffffff;\n  cursor: pointer;\n  font-size: 10px;\n  font-weight: 100;\n\n  &:hover {\n    opacity: .95;\n  }\n"],e||(e=n.slice(0)),Object.freeze(Object.defineProperties(n,{raw:{value:Object.freeze(e)}})));return a=function(){return t},t}var u=i.ZP.button.withConfig({componentId:"sc-753c6b5a-0"})(a(),(function(n){return"large"===n.size?"24rem":"12rem"}),(function(n){return"large"===n.size?"2.125rem":"1.5rem"})),s=function(n){var e=n.value,t=n.onClick,i=n.size,o=void 0===i?"large":i;return(0,r.jsx)(u,{onClick:t,size:o,children:e})};s.prototype={value:c().string.isRequired,onClick:c().func.isRequired,size:c().oneOf(["large","small"])},e.Z=s},5786:function(n,e,t){"use strict";t.r(e);var r=t(4051),i=t.n(r),o=t(5893),c=t(2125),a=t(7294),u=t(4240),s=t(3206),f=t(6843);function l(n,e,t,r,i,o,c){try{var a=n[o](c),u=a.value}catch(s){return void t(s)}a.done?e(u):Promise.resolve(u).then(r,i)}function p(n,e){return e||(e=n.slice(0)),Object.freeze(Object.defineProperties(n,{raw:{value:Object.freeze(e)}}))}function d(){var n=p(["\n  margin-left: auto;\n  margin-right: auto;\n  max-width: 24rem;\n  padding-top: 2rem;\n\n  display: grid;\n  grid-template-rows: 1fr auto;\n  min-height: 100%;\n"]);return d=function(){return n},n}function h(){var n=p(["\n  margin-left: -5rem;\n"]);return h=function(){return n},n}var g=c.ZP.section.withConfig({componentId:"sc-f75e0738-0"})(d()),v=c.ZP.img.withConfig({componentId:"sc-f75e0738-1"})(h());e.default=function(){var n=(0,a.useContext)(f.Z).setup,e=function(){var e,t=(e=i().mark((function e(){return i().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,n();case 2:case"end":return e.stop()}}),e)})),function(){var n=this,t=arguments;return new Promise((function(r,i){var o=e.apply(n,t);function c(n){l(o,r,i,c,a,"next",n)}function a(n){l(o,r,i,c,a,"throw",n)}c(void 0)}))});return function(){return t.apply(this,arguments)}}();return(0,o.jsx)(g,{children:(0,o.jsxs)("div",{children:[(0,o.jsx)(u.Z,{header:"Welcome to Infra",subheader:"Infra has been successfully installed. Please click Get Started below to obtain your Infra Access Key."}),(0,o.jsx)(v,{src:"/welcome.svg"}),(0,o.jsx)(s.Z,{onClick:e,value:"Get Started"})]})})}},2703:function(n,e,t){"use strict";var r=t(414);function i(){}function o(){}o.resetWarningCache=i,n.exports=function(){function n(n,e,t,i,o,c){if(c!==r){var a=new Error("Calling PropTypes validators directly is not supported by the `prop-types` package. Use PropTypes.checkPropTypes() to call them. Read more at http://fb.me/use-check-prop-types");throw a.name="Invariant Violation",a}}function e(){return n}n.isRequired=n;var t={array:n,bigint:n,bool:n,func:n,number:n,object:n,string:n,symbol:n,any:n,arrayOf:e,element:n,elementType:n,instanceOf:e,node:n,objectOf:e,oneOf:e,oneOfType:e,shape:e,exact:e,checkPropTypes:o,resetWarningCache:i};return t.PropTypes=t,t}},5697:function(n,e,t){n.exports=t(2703)()},414:function(n){"use strict";n.exports="SECRET_DO_NOT_PASS_THIS_OR_YOU_WILL_BE_FIRED"}},function(n){n.O(0,[125,774,888,179],(function(){return e=209,n(n.s=e);var e}));var e=n.O();_N_E=e}]);