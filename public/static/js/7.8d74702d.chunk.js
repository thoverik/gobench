(this["webpackJsonpgobench-ui"]=this["webpackJsonpgobench-ui"]||[]).push([[7],{54:function(e,t,n){"use strict";n.r(t);var c=n(4),a=n(10),i=n(0),u=n.n(i),b=n(1),r=n(8),s=n(13),l=Object(i.lazy)((function(){return Promise.all([n.e(3),n.e(4)]).then(n.bind(null,55))}));t.default=function(e){var t,n=e.metrics,O=e.unit,j=e.timeRange,h=Object(i.useState)([]),o=Object(a.a)(h,2),f=o[0],g=o[1],m=Object(i.useContext)(s.a),p=Object(b.get)(m,"status",""),d=Object(b.get)(m,"timestamp",""),E=!["finished","cancel"].includes(p);Object(i.useEffect)((function(){n.length>0&&Object(r.d)(n,j,d,E).then((function(e){g(e)}))}),[n]),Object(r.h)((function(){E&&Object(r.e)(n,f).then((function(e){g(e)}))}),E?r.a:null),Object(b.get)(f,"[0].type","")===r.b.HISTOGRAM?t=Object(c.a)(Object(r.g)(Object(b.get)(f,"[0].chartData.data",[]))):Object(b.isArray)(f)&&(t=f.map((function(e){return Object(b.get)(e,"chartData",{name:e.title,data:[]})})));var k=E?Object(r.f)(t,j):t;return u.a.createElement(i.Suspense,{fallback:u.a.createElement("p",null,"Loading chart...")},u.a.createElement(l,{height:"220",series:k,unit:O}))}}}]);
//# sourceMappingURL=7.8d74702d.chunk.js.map