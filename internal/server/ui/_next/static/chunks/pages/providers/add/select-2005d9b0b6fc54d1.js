(self.webpackChunk_N_E=self.webpackChunk_N_E||[]).push([[49],{1981:function(n,e,r){(window.__NEXT_P=window.__NEXT_P||[]).push(["/providers/add/select",function(){return r(966)}])},7054:function(n,e,r){"use strict";var t=r(5893),i=r(5697),o=r.n(i),c=r(1664);function u(){var n,e,r=(n=["\n  cursor: pointer;\n\n  &:hover {\n    opacity: .5\n  }\n"],e||(e=n.slice(0)),Object.freeze(Object.defineProperties(n,{raw:{value:Object.freeze(e)}})));return u=function(){return r},r}var s=r(2125).ZP.div.withConfig({componentId:"sc-135ec5ea-0"})(u()),f=function(n){var e=n.previousPage,r=void 0===e?"/":e;return(0,t.jsx)(s,{children:(0,t.jsx)(c.default,{href:r,children:(0,t.jsx)("img",{src:"/close-icon.svg"})})})};f.prototype={previousPage:o().string},e.Z=f},292:function(n,e,r){"use strict";var t=r(5893),i=r(2125);function o(n,e){return e||(e=n.slice(0)),Object.freeze(Object.defineProperties(n,{raw:{value:Object.freeze(e)}}))}function c(){var n=o(["\n  & > *:not(:first-child) {\n    padding-top: 1rem;\n  }\n"]);return c=function(){return n},n}function u(){var n=o(["\n  font-weight: 200;\n  font-size: 10px;\n  line-height: 4px;\n\n  display: flex;\n  align-items: center;\n  text-transform: uppercase;\n"]);return u=function(){return n},n}function s(){var n=o(["\n  font-weight: 100;\n  font-size: 10px;\n  line-height: 12px;\n  display: flex;\n  align-items: center;\n\n  opacity: .5;\n"]);return s=function(){return n},n}var f=i.ZP.section.withConfig({componentId:"sc-412273f1-0"})(c()),a=i.ZP.div.withConfig({componentId:"sc-412273f1-1"})(u()),d=i.ZP.div.withConfig({componentId:"sc-412273f1-2"})(s());e.Z=function(n){var e=n.header,r=n.subheader;return(0,t.jsxs)(f,{children:[(0,t.jsx)(a,{children:e}),(0,t.jsx)(d,{children:r})]})}},5150:function(n,e,r){"use strict";var t=r(5893),i=r(2125),o=r(5697),c=r.n(o);function u(n,e){return e||(e=n.slice(0)),Object.freeze(Object.defineProperties(n,{raw:{value:Object.freeze(e)}}))}function s(){var n=u(["\n  display: flex;\n  flex-direction: row;\n  padding: .5rem;\n"]);return s=function(){return n},n}function f(){var n=u(["\n  padding-top: .4rem;  \n"]);return f=function(){return n},n}function a(){var n=u(["\n  padding-left: 1rem;\n  text-align: left;\n\n  & > *:not(:first-child) {\n    padding-top: .15rem;\n  }\n"]);return a=function(){return n},n}function d(){var n=u(["\n  font-weight: 300;\n  font-size: .75rem;\n  line-height: 1rem;\n  text-transform: capitalize;\n"]);return d=function(){return n},n}function p(){var n=u(["\n  font-weight: 300;\n  font-size: .5rem;\n  line-height: .75rem;\n  text-transform: uppercase;\n  color: #FFFFFF;\n  opacity: 0.3;\n"]);return p=function(){return n},n}var h=i.ZP.div.withConfig({componentId:"sc-d47f4971-0"})(s()),v=i.ZP.div.withConfig({componentId:"sc-d47f4971-1"})(f()),l=i.ZP.div.withConfig({componentId:"sc-d47f4971-2"})(a()),g=i.ZP.div.withConfig({componentId:"sc-d47f4971-3"})(d()),m=i.ZP.div.withConfig({componentId:"sc-d47f4971-4"})(p()),x=function(n){var e=n.type,r=n.name;return(0,t.jsxs)(h,{children:[(0,t.jsx)(v,{children:(0,t.jsx)("img",{src:"/".concat(e,".svg")})}),(0,t.jsxs)(l,{children:[(0,t.jsx)(g,{children:e}),(0,t.jsx)(m,{children:r})]})]})};x.prototype={type:c().string.isRequired,name:c().string.isRequired},e.Z=x},5885:function(n,e,r){"use strict";var t=r(5893),i=r(2125),o=r(5697),c=r.n(o),u=r(5150);function s(n,e){return e||(e=n.slice(0)),Object.freeze(Object.defineProperties(n,{raw:{value:Object.freeze(e)}}))}function f(){var n=s(["\n  & > *:not(:first-child) {\n    margin-top: .3rem;\n  }\n"]);return f=function(){return n},n}function a(){var n=s(["\n  width: 24rem;\n  height: 3rem;\n  background: rgba(255,255,255,0.02);\n  opacity: 1;\n  border-radius: .25rem;\n  border: none;\n  cursor: pointer;\n  color: #FFFFFF;\n\n  &:hover { opacity: .95 }\n"]);return a=function(){return n},n}var d=i.ZP.div.withConfig({componentId:"sc-d75b49e6-0"})(f()),p=i.ZP.button.withConfig({componentId:"sc-d75b49e6-1"})(a()),h=function(n){var e=n.providers;return(0,t.jsx)(d,{children:e.map((function(n,e){return(0,t.jsx)(p,{onClick:function(){return n.onClick()},children:(0,t.jsx)(u.Z,{type:n.type,name:n.name})},e)}))})};h.prototype={providers:c().arrayOf(c().shape({type:c().string,name:c().string,onClick:c().func})).isRequired},e.Z=h},966:function(n,e,r){"use strict";r.r(e);var t=r(4051),i=r.n(t),o=r(5893),c=r(1163),u=r(2125),s=r(7054),f=r(5885),a=r(292);function d(n,e,r,t,i,o,c){try{var u=n[o](c),s=u.value}catch(f){return void r(f)}u.done?e(s):Promise.resolve(s).then(t,i)}function p(n,e){return e||(e=n.slice(0)),Object.freeze(Object.defineProperties(n,{raw:{value:Object.freeze(e)}}))}function h(){var n=p(["\n  position: relative;\n"]);return h=function(){return n},n}function v(){var n=p(["\n  margin-left: auto;\n  margin-right: auto;\n  max-width: 24rem;\n  padding-top: 1.5rem;\n\n  & > *:not(:first-child) {\n    padding-top: 1.75rem;\n  }  \n"]);return v=function(){return n},n}function l(){var n=p(["\n  position: absolute;\n  top: .5rem;\n  right: .5rem;\n"]);return l=function(){return n},n}function g(){var n=p(["\n  & > *:not(:first-child) {\n    padding-top: 42px;\n  }\n"]);return g=function(){return n},n}var m=u.ZP.section.withConfig({componentId:"sc-4aecbff4-0"})(h()),x=u.ZP.section.withConfig({componentId:"sc-4aecbff4-1"})(v()),j=u.ZP.div.withConfig({componentId:"sc-4aecbff4-2"})(l()),w=u.ZP.section.withConfig({componentId:"sc-4aecbff4-3"})(g()),b=function(){var n,e=(n=i().mark((function n(){return i().wrap((function(n){for(;;)switch(n.prev=n.next){case 0:return n.next=2,c.default.push({pathname:"/providers/add/okta"},void 0,{shallow:!0});case 2:case"end":return n.stop()}}),n)})),function(){var e=this,r=arguments;return new Promise((function(t,i){var o=n.apply(e,r);function c(n){d(o,t,i,c,u,"next",n)}function u(n){d(o,t,i,c,u,"throw",n)}c(void 0)}))});return function(){return e.apply(this,arguments)}}(),P=[{type:"okta",name:"Identity Provider",onClick:function(){return b()}}];e.default=function(){return(0,o.jsxs)(m,{children:[(0,o.jsxs)(x,{children:[(0,o.jsx)(a.Z,{header:"Connect Identity Providers",subheader:"People, Groups and Machines"}),(0,o.jsx)(a.Z,{header:"Choose an Identity Provider",subheader:(0,o.jsxs)(o.Fragment,{children:["Currently there are no identity providers connected to Infra. ",(0,o.jsx)("br",{}),"Choose your IdP source below and get connected."]})}),(0,o.jsx)(w,{children:(0,o.jsx)("div",{children:(0,o.jsx)(f.Z,{providers:P})})})]}),(0,o.jsx)(j,{children:(0,o.jsx)(s.Z,{previousPage:"/providers"})})]})}}},function(n){n.O(0,[125,210,774,888,179],(function(){return e=1981,n(n.s=e);var e}));var e=n.O();_N_E=e}]);