(this["webpackJsonpgobench-ui"]=this["webpackJsonpgobench-ui"]||[]).push([[5],{53:function(e,t,a){"use strict";a.r(t);var n=a(10),c=a(0),i=a.n(c),l=a(1),r=a(5),m=a(8),s=a(13),o=Object(c.lazy)((function(){return a.e(7).then(a.bind(null,54))}));t.default=function(e){var t=e.graph,a=void 0===t?{}:t,u=e.timestamp,h=Object(c.useState)([]),p=Object(n.a)(h,2),b=p[0],g=p[1],E=Object(c.useState)(m.c["1h"]),N=Object(n.a)(E,2),f=N[0],d=N[1],v=Object(c.useContext)(s.a),j=Object(l.get)(v,"status","");Object(c.useEffect)((function(){a&&a.id&&r.a.getMetrics(a.id).then((function(e){return g(e)}))}),[a]);var O=!["finished","cancel"].includes(j);return i.a.createElement("div",{className:"graph"},i.a.createElement("div",{className:"graph-header"},i.a.createElement("h5",{title:a.id||"",className:"graph-title"},Object(l.get)(a,"title","")," (",Object(l.get)(a,"unit",""),")"),O&&i.a.createElement("div",{className:"options-group"},i.a.createElement("ul",{className:"time-range-options-list"},i.a.createElement("li",{className:"time-range-option"},i.a.createElement("button",{className:f===m.c["5m"]?"active":"",onClick:function(){return d(m.c["5m"])}},"5m")),i.a.createElement("li",{className:"time-range-option"},i.a.createElement("button",{className:f===m.c["15m"]?"active":"",onClick:function(){return d(m.c["15m"])}},"15m")),i.a.createElement("li",{className:"time-range-option"},i.a.createElement("button",{className:f===m.c["30m"]?"active":"",onClick:function(){return d(m.c["30m"])}},"30m")),i.a.createElement("li",{className:"time-range-option"},i.a.createElement("button",{className:f===m.c["1h"]?"active":"",onClick:function(){return d(m.c["1h"])}},"1h")),i.a.createElement("li",{className:"time-range-option"},i.a.createElement("button",{className:f===m.c["12h"]?"active":"",onClick:function(){return d(m.c["12h"])}},"12h")),i.a.createElement("li",{className:"time-range-option"},i.a.createElement("button",{className:f===m.c["24h"]?"active":"",onClick:function(){return d(m.c["24h"])}},"24h"))))),i.a.createElement(c.Suspense,{fallback:i.a.createElement("p",null,"Loading graph...")},i.a.createElement(o,{timeRange:f,graph:a,metrics:b,timestamp:u,unit:Object(l.get)(a,"unit","")})))}}}]);
//# sourceMappingURL=5.d104a95e.chunk.js.map