package alert

// CveDictJa has CVE-ID key which included JPCERT alerts
var CveDictJa = map[string][]string{
	"CVE-2006-0003":    {"https://www.jpcert.or.jp/at/2007/at070016.html"},
	"CVE-2006-0005":    {"https://www.jpcert.or.jp/at/2007/at070016.html"},
	"CVE-2006-1173":    {"https://www.jpcert.or.jp/at/2006/at060008.html"},
	"CVE-2006-3014":    {"https://www.jpcert.or.jp/at/2006/at060009.html"},
	"CVE-2006-3059":    {"https://www.jpcert.or.jp/at/2006/at060009.html"},
	"CVE-2006-3086":    {"https://www.jpcert.or.jp/at/2006/at060009.html"},
	"CVE-2006-3643":    {"https://www.jpcert.or.jp/at/2007/at070016.html"},
	"CVE-2006-3730":    {"https://www.jpcert.or.jp/at/2007/at070016.html"},
	"CVE-2006-3877":    {"https://www.jpcert.or.jp/at/2007/at070005.html"},
	"CVE-2006-5198":    {"https://www.jpcert.or.jp/at/2007/at070016.html"},
	"CVE-2006-5745":    {"https://www.jpcert.or.jp/at/2007/at070016.html"},
	"CVE-2007-0015":    {"https://www.jpcert.or.jp/at/2007/at070016.html"},
	"CVE-2007-0038":    {"https://www.jpcert.or.jp/at/2007/at070016.html"},
	"CVE-2008-4609":    {"https://www.jpcert.or.jp/at/2009/at090019.html"},
	"CVE-2010-0886":    {"https://www.jpcert.or.jp/at/2010/at100010.html"},
	"CVE-2010-0887":    {"https://www.jpcert.or.jp/at/2010/at100010.html"},
	"CVE-2011-1910":    {"https://www.jpcert.or.jp/at/2011/at110014.html"},
	"CVE-2011-2444":    {"https://www.jpcert.or.jp/at/2011/at110026.html"},
	"CVE-2011-2462":    {"https://www.jpcert.or.jp/at/2011/at110034.html"},
	"CVE-2011-2465":    {"https://www.jpcert.or.jp/at/2011/at110019.html"},
	"CVE-2011-3192":    {"https://www.jpcert.or.jp/at/2011/at110023.html"},
	"CVE-2011-3348":    {"https://www.jpcert.or.jp/at/2011/at110023.html"},
	"CVE-2011-3544":    {"https://www.jpcert.or.jp/at/2011/at110032.html"},
	"CVE-2011-4313":    {"https://www.jpcert.or.jp/at/2011/at110031.html"},
	"CVE-2012-0002":    {"https://www.jpcert.or.jp/at/2012/at120009.html"},
	"CVE-2012-0507":    {"https://www.jpcert.or.jp/at/2012/at120010.html"},
	"CVE-2012-0767":    {"https://www.jpcert.or.jp/at/2012/at120006.html"},
	"CVE-2012-0779":    {"https://www.jpcert.or.jp/at/2012/at120014.html"},
	"CVE-2012-0830":    {"https://www.jpcert.or.jp/at/2012/at120004.html"},
	"CVE-2012-2311":    {"https://www.jpcert.or.jp/at/2012/at120016.html"},
	"CVE-2012-4244":    {"https://www.jpcert.or.jp/at/2012/at120029.html"},
	"CVE-2012-4681":    {"https://www.jpcert.or.jp/at/2012/at120028.html"},
	"CVE-2012-4969":    {"https://www.jpcert.or.jp/at/2012/at120030.html"},
	"CVE-2012-5166":    {"https://www.jpcert.or.jp/at/2012/at120033.html"},
	"CVE-2013-0422":    {"https://www.jpcert.or.jp/at/2013/at130004.html"},
	"CVE-2013-1493":    {"https://www.jpcert.or.jp/at/2013/at130014.html"},
	"CVE-2013-2266":    {"https://www.jpcert.or.jp/at/2013/at130017.html"},
	"CVE-2013-2494":    {"https://www.jpcert.or.jp/at/2013/at130017.html"},
	"CVE-2013-3893":    {"https://www.jpcert.or.jp/at/2013/at130040.html", "https://www.jpcert.or.jp/at/2013/at130041.html"},
	"CVE-2013-3906":    {"https://www.jpcert.or.jp/at/2013/at130044.html"},
	"CVE-2013-3918":    {"https://www.jpcert.or.jp/at/2013/at130045.html"},
	"CVE-2013-3919":    {"https://www.jpcert.or.jp/at/2013/at130026.html"},
	"CVE-2013-4854":    {"https://www.jpcert.or.jp/at/2013/at130034.html"},
	"CVE-2014-0050":    {"https://www.jpcert.or.jp/at/2014/at140007.html"},
	"CVE-2014-0160":    {"https://www.jpcert.or.jp/at/2014/at140013.html"},
	"CVE-2014-0322":    {"https://www.jpcert.or.jp/at/2014/at140009.html"},
	"CVE-2014-1776":    {"https://www.jpcert.or.jp/at/2014/at140018.html", "https://www.jpcert.or.jp/at/2014/at140020.html"},
	"CVE-2014-3383":    {"https://www.jpcert.or.jp/at/2015/at150021.html"},
	"CVE-2014-3859":    {"https://www.jpcert.or.jp/at/2014/at140027.html"},
	"CVE-2014-4114":    {"https://www.jpcert.or.jp/at/2014/at140039.html"},
	"CVE-2014-6271":    {"https://www.jpcert.or.jp/at/2014/at140037.html", "https://www.jpcert.or.jp/at/2014/at140038.html"},
	"CVE-2014-6277":    {"https://www.jpcert.or.jp/at/2014/at140037.html"},
	"CVE-2014-6278":    {"https://www.jpcert.or.jp/at/2014/at140037.html"},
	"CVE-2014-6324":    {"https://www.jpcert.or.jp/at/2014/at140048.html"},
	"CVE-2014-6332":    {"https://www.jpcert.or.jp/at/2015/at150015.html"},
	"CVE-2014-6352":    {"https://www.jpcert.or.jp/at/2014/at140043.html"},
	"CVE-2014-7169":    {"https://www.jpcert.or.jp/at/2014/at140037.html"},
	"CVE-2014-7186":    {"https://www.jpcert.or.jp/at/2014/at140037.html"},
	"CVE-2014-7187":    {"https://www.jpcert.or.jp/at/2014/at140037.html"},
	"CVE-2014-8361":    {"https://www.jpcert.or.jp/at/2017/at170049.html"},
	"CVE-2014-8500":    {"https://www.jpcert.or.jp/at/2014/at140050.html"},
	"CVE-2014-9163":    {"https://www.jpcert.or.jp/at/2014/at140052.html"},
	"CVE-2015-0313":    {"https://www.jpcert.or.jp/at/2015/at150015.html"},
	"CVE-2015-1769":    {"https://www.jpcert.or.jp/at/2015/at150028.html"},
	"CVE-2015-5119":    {"https://www.jpcert.or.jp/at/2015/at150019.html"},
	"CVE-2015-5122":    {"https://www.jpcert.or.jp/at/2015/at150020.html"},
	"CVE-2015-5123":    {"https://www.jpcert.or.jp/at/2015/at150020.html"},
	"CVE-2015-5477":    {"https://www.jpcert.or.jp/at/2015/at150027.html"},
	"CVE-2015-5986":    {"https://www.jpcert.or.jp/at/2015/at150031.html"},
	"CVE-2015-6835":    {"https://www.jpcert.or.jp/at/2016/at160036.html"},
	"CVE-2015-7547":    {"https://www.jpcert.or.jp/at/2016/at160009.html"},
	"CVE-2015-7645":    {"https://www.jpcert.or.jp/at/2015/at150036.html", "https://www.jpcert.or.jp/at/2015/at150037.html"},
	"CVE-2015-8000":    {"https://www.jpcert.or.jp/at/2015/at150043.html"},
	"CVE-2015-8461":    {"https://www.jpcert.or.jp/at/2015/at150043.html"},
	"CVE-2015-8562":    {"https://www.jpcert.or.jp/at/2016/at160036.html"},
	"CVE-2015-8651":    {"https://www.jpcert.or.jp/at/2016/at160001.html"},
	"CVE-2015-8704":    {"https://www.jpcert.or.jp/at/2016/at160006.html"},
	"CVE-2015-8705":    {"https://www.jpcert.or.jp/at/2016/at160006.html"},
	"CVE-2016-0189":    {"https://www.jpcert.or.jp/at/2016/at160022.html"},
	"CVE-2016-0636":    {"https://www.jpcert.or.jp/at/2016/at160015.html"},
	"CVE-2016-0800":    {"https://www.jpcert.or.jp/at/2016/at160010.html"},
	"CVE-2016-1000109": {"https://www.jpcert.or.jp/at/2016/at160031.html"},
	"CVE-2016-1000110": {"https://www.jpcert.or.jp/at/2016/at160031.html"},
	"CVE-2016-1010":    {"https://www.jpcert.or.jp/at/2016/at160014.html"},
	"CVE-2016-1019":    {"https://www.jpcert.or.jp/at/2016/at160016.html"},
	"CVE-2016-1204":    {"https://www.jpcert.or.jp/at/2016/at160019.html"},
	"CVE-2016-1286":    {"https://www.jpcert.or.jp/at/2016/at160013.html", "https://www.jpcert.or.jp/at/2016/at160037.html"},
	"CVE-2016-2776":    {"https://www.jpcert.or.jp/at/2016/at160037.html"},
	"CVE-2016-3081":    {"https://www.jpcert.or.jp/at/2016/at160020.html"},
	"CVE-2016-3227":    {"https://www.jpcert.or.jp/at/2016/at160025.html"},
	"CVE-2016-3714":    {"https://www.jpcert.or.jp/at/2016/at160021.html"},
	"CVE-2016-3715":    {"https://www.jpcert.or.jp/at/2016/at160021.html"},
	"CVE-2016-3716":    {"https://www.jpcert.or.jp/at/2016/at160021.html"},
	"CVE-2016-3717":    {"https://www.jpcert.or.jp/at/2016/at160021.html"},
	"CVE-2016-3718":    {"https://www.jpcert.or.jp/at/2016/at160021.html"},
	"CVE-2016-4117":    {"https://www.jpcert.or.jp/at/2016/at160024.html"},
	"CVE-2016-4171":    {"https://www.jpcert.or.jp/at/2016/at160026.html"},
	"CVE-2016-4438":    {"https://www.jpcert.or.jp/at/2016/at160027.html"},
	"CVE-2016-5385":    {"https://www.jpcert.or.jp/at/2016/at160031.html"},
	"CVE-2016-5386":    {"https://www.jpcert.or.jp/at/2016/at160031.html"},
	"CVE-2016-5387":    {"https://www.jpcert.or.jp/at/2016/at160031.html"},
	"CVE-2016-5388":    {"https://www.jpcert.or.jp/at/2016/at160031.html"},
	"CVE-2016-6307":    {"https://www.jpcert.or.jp/at/2016/at160038.html"},
	"CVE-2016-6309":    {"https://www.jpcert.or.jp/at/2016/at160038.html"},
	"CVE-2016-7189":    {"https://www.jpcert.or.jp/at/2016/at160039.html"},
	"CVE-2016-7836":    {"https://www.jpcert.or.jp/at/2016/at160051.html", "https://www.jpcert.or.jp/at/2017/at170023.html"},
	"CVE-2016-7855":    {"https://www.jpcert.or.jp/at/2016/at160039.html", "https://www.jpcert.or.jp/at/2016/at160043.html"},
	"CVE-2016-7892":    {"https://www.jpcert.or.jp/at/2016/at160048.html", "https://www.jpcert.or.jp/at/2016/at160049.html"},
	"CVE-2016-8864":    {"https://www.jpcert.or.jp/at/2016/at160044.html"},
	"CVE-2016-9131":    {"https://www.jpcert.or.jp/at/2017/at170004.html"},
	"CVE-2016-9147":    {"https://www.jpcert.or.jp/at/2017/at170004.html"},
	"CVE-2016-9444":    {"https://www.jpcert.or.jp/at/2017/at170004.html"},
	"CVE-2016-9778":    {"https://www.jpcert.or.jp/at/2017/at170004.html"},
	"CVE-2017-0093":    {"https://www.jpcert.or.jp/at/2017/at170015.html"},
	"CVE-2017-0106":    {"https://www.jpcert.or.jp/at/2017/at170015.html"},
	"CVE-2017-0145":    {"https://www.jpcert.or.jp/at/2017/at170020.html"},
	"CVE-2017-0158":    {"https://www.jpcert.or.jp/at/2017/at170015.html"},
	"CVE-2017-0160":    {"https://www.jpcert.or.jp/at/2017/at170015.html"},
	"CVE-2017-0161":    {"https://www.jpcert.or.jp/at/2017/at170036.html"},
	"CVE-2017-0162":    {"https://www.jpcert.or.jp/at/2017/at170015.html"},
	"CVE-2017-0163":    {"https://www.jpcert.or.jp/at/2017/at170015.html"},
	"CVE-2017-0180":    {"https://www.jpcert.or.jp/at/2017/at170015.html"},
	"CVE-2017-0181":    {"https://www.jpcert.or.jp/at/2017/at170015.html"},
	"CVE-2017-0199":    {"https://www.jpcert.or.jp/at/2017/at170015.html"},
	"CVE-2017-0200":    {"https://www.jpcert.or.jp/at/2017/at170015.html"},
	"CVE-2017-0201":    {"https://www.jpcert.or.jp/at/2017/at170015.html"},
	"CVE-2017-0202":    {"https://www.jpcert.or.jp/at/2017/at170015.html"},
	"CVE-2017-0205":    {"https://www.jpcert.or.jp/at/2017/at170015.html"},
	"CVE-2017-0210":    {"https://www.jpcert.or.jp/at/2017/at170015.html"},
	"CVE-2017-0221":    {"https://www.jpcert.or.jp/at/2017/at170019.html"},
	"CVE-2017-0222":    {"https://www.jpcert.or.jp/at/2017/at170019.html"},
	"CVE-2017-0224":    {"https://www.jpcert.or.jp/at/2017/at170019.html"},
	"CVE-2017-0227":    {"https://www.jpcert.or.jp/at/2017/at170019.html"},
	"CVE-2017-0228":    {"https://www.jpcert.or.jp/at/2017/at170019.html"},
	"CVE-2017-0229":    {"https://www.jpcert.or.jp/at/2017/at170019.html"},
	"CVE-2017-0235":    {"https://www.jpcert.or.jp/at/2017/at170019.html"},
	"CVE-2017-0236":    {"https://www.jpcert.or.jp/at/2017/at170019.html"},
	"CVE-2017-0240":    {"https://www.jpcert.or.jp/at/2017/at170019.html"},
	"CVE-2017-0250":    {"https://www.jpcert.or.jp/at/2017/at170032.html"},
	"CVE-2017-0261":    {"https://www.jpcert.or.jp/at/2017/at170019.html"},
	"CVE-2017-0263":    {"https://www.jpcert.or.jp/at/2017/at170019.html"},
	"CVE-2017-0266":    {"https://www.jpcert.or.jp/at/2017/at170019.html"},
	"CVE-2017-0272":    {"https://www.jpcert.or.jp/at/2017/at170019.html"},
	"CVE-2017-0277":    {"https://www.jpcert.or.jp/at/2017/at170019.html"},
	"CVE-2017-0278":    {"https://www.jpcert.or.jp/at/2017/at170019.html"},
	"CVE-2017-0279":    {"https://www.jpcert.or.jp/at/2017/at170019.html"},
	"CVE-2017-0283":    {"https://www.jpcert.or.jp/at/2017/at170022.html"},
	"CVE-2017-0290":    {"https://www.jpcert.or.jp/at/2017/at170019.html"},
	"CVE-2017-0291":    {"https://www.jpcert.or.jp/at/2017/at170022.html"},
	"CVE-2017-0292":    {"https://www.jpcert.or.jp/at/2017/at170022.html"},
	"CVE-2017-0293":    {"https://www.jpcert.or.jp/at/2017/at170032.html"},
	"CVE-2017-0294":    {"https://www.jpcert.or.jp/at/2017/at170022.html"},
	"CVE-2017-0781":    {"https://www.jpcert.or.jp/at/2017/at170037.html"},
	"CVE-2017-0782":    {"https://www.jpcert.or.jp/at/2017/at170037.html"},
	"CVE-2017-0783":    {"https://www.jpcert.or.jp/at/2017/at170037.html"},
	"CVE-2017-0785":    {"https://www.jpcert.or.jp/at/2017/at170037.html"},
	"CVE-2017-1000250": {"https://www.jpcert.or.jp/at/2017/at170037.html"},
	"CVE-2017-1000251": {"https://www.jpcert.or.jp/at/2017/at170037.html"},
	"CVE-2017-10271":   {"https://www.jpcert.or.jp/at/2018/at180004.html"},
	"CVE-2017-10845":   {"https://www.jpcert.or.jp/at/2017/at170034.html"},
	"CVE-2017-10846":   {"https://www.jpcert.or.jp/at/2017/at170034.html"},
	"CVE-2017-11223":   {"https://www.jpcert.or.jp/at/2017/at170031.html"},
	"CVE-2017-11292":   {"https://www.jpcert.or.jp/at/2017/at170040.html"},
	"CVE-2017-11762":   {"https://www.jpcert.or.jp/at/2017/at170039.html"},
	"CVE-2017-11763":   {"https://www.jpcert.or.jp/at/2017/at170039.html"},
	"CVE-2017-11764":   {"https://www.jpcert.or.jp/at/2017/at170036.html"},
	"CVE-2017-11766":   {"https://www.jpcert.or.jp/at/2017/at170036.html"},
	"CVE-2017-11771":   {"https://www.jpcert.or.jp/at/2017/at170039.html"},
	"CVE-2017-11779":   {"https://www.jpcert.or.jp/at/2017/at170039.html"},
	"CVE-2017-11792":   {"https://www.jpcert.or.jp/at/2017/at170039.html"},
	"CVE-2017-11793":   {"https://www.jpcert.or.jp/at/2017/at170039.html"},
	"CVE-2017-11796":   {"https://www.jpcert.or.jp/at/2017/at170039.html"},
	"CVE-2017-11798":   {"https://www.jpcert.or.jp/at/2017/at170039.html"},
	"CVE-2017-11799":   {"https://www.jpcert.or.jp/at/2017/at170039.html"},
	"CVE-2017-11800":   {"https://www.jpcert.or.jp/at/2017/at170039.html"},
	"CVE-2017-11802":   {"https://www.jpcert.or.jp/at/2017/at170039.html"},
	"CVE-2017-11804":   {"https://www.jpcert.or.jp/at/2017/at170039.html"},
	"CVE-2017-11805":   {"https://www.jpcert.or.jp/at/2017/at170039.html"},
	"CVE-2017-11806":   {"https://www.jpcert.or.jp/at/2017/at170039.html"},
	"CVE-2017-11807":   {"https://www.jpcert.or.jp/at/2017/at170039.html"},
	"CVE-2017-11808":   {"https://www.jpcert.or.jp/at/2017/at170039.html"},
	"CVE-2017-11809":   {"https://www.jpcert.or.jp/at/2017/at170039.html"},
	"CVE-2017-11810":   {"https://www.jpcert.or.jp/at/2017/at170039.html"},
	"CVE-2017-11811":   {"https://www.jpcert.or.jp/at/2017/at170039.html"},
	"CVE-2017-11812":   {"https://www.jpcert.or.jp/at/2017/at170039.html"},
	"CVE-2017-11813":   {"https://www.jpcert.or.jp/at/2017/at170039.html"},
	"CVE-2017-11819":   {"https://www.jpcert.or.jp/at/2017/at170039.html"},
	"CVE-2017-11821":   {"https://www.jpcert.or.jp/at/2017/at170039.html"},
	"CVE-2017-11822":   {"https://www.jpcert.or.jp/at/2017/at170039.html"},
	"CVE-2017-11826":   {"https://www.jpcert.or.jp/at/2017/at170039.html"},
	"CVE-2017-11836":   {"https://www.jpcert.or.jp/at/2017/at170044.html"},
	"CVE-2017-11837":   {"https://www.jpcert.or.jp/at/2017/at170044.html"},
	"CVE-2017-11838":   {"https://www.jpcert.or.jp/at/2017/at170044.html"},
	"CVE-2017-11839":   {"https://www.jpcert.or.jp/at/2017/at170044.html"},
	"CVE-2017-11840":   {"https://www.jpcert.or.jp/at/2017/at170044.html"},
	"CVE-2017-11841":   {"https://www.jpcert.or.jp/at/2017/at170044.html"},
	"CVE-2017-11843":   {"https://www.jpcert.or.jp/at/2017/at170044.html"},
	"CVE-2017-11845":   {"https://www.jpcert.or.jp/at/2017/at170044.html"},
	"CVE-2017-11846":   {"https://www.jpcert.or.jp/at/2017/at170044.html"},
	"CVE-2017-11855":   {"https://www.jpcert.or.jp/at/2017/at170044.html"},
	"CVE-2017-11856":   {"https://www.jpcert.or.jp/at/2017/at170044.html"},
	"CVE-2017-11858":   {"https://www.jpcert.or.jp/at/2017/at170044.html"},
	"CVE-2017-11861":   {"https://www.jpcert.or.jp/at/2017/at170044.html"},
	"CVE-2017-11862":   {"https://www.jpcert.or.jp/at/2017/at170044.html"},
	"CVE-2017-11866":   {"https://www.jpcert.or.jp/at/2017/at170044.html"},
	"CVE-2017-11869":   {"https://www.jpcert.or.jp/at/2017/at170044.html"},
	"CVE-2017-11870":   {"https://www.jpcert.or.jp/at/2017/at170044.html"},
	"CVE-2017-11871":   {"https://www.jpcert.or.jp/at/2017/at170044.html"},
	"CVE-2017-11873":   {"https://www.jpcert.or.jp/at/2017/at170044.html"},
	"CVE-2017-11882":   {"https://www.jpcert.or.jp/at/2017/at170044.html"},
	"CVE-2017-11886":   {"https://www.jpcert.or.jp/at/2017/at170048.html"},
	"CVE-2017-11888":   {"https://www.jpcert.or.jp/at/2017/at170048.html"},
	"CVE-2017-11889":   {"https://www.jpcert.or.jp/at/2017/at170048.html"},
	"CVE-2017-11890":   {"https://www.jpcert.or.jp/at/2017/at170048.html"},
	"CVE-2017-11893":   {"https://www.jpcert.or.jp/at/2017/at170048.html"},
	"CVE-2017-11894":   {"https://www.jpcert.or.jp/at/2017/at170048.html"},
	"CVE-2017-11895":   {"https://www.jpcert.or.jp/at/2017/at170048.html"},
	"CVE-2017-11901":   {"https://www.jpcert.or.jp/at/2017/at170048.html"},
	"CVE-2017-11903":   {"https://www.jpcert.or.jp/at/2017/at170048.html"},
	"CVE-2017-11905":   {"https://www.jpcert.or.jp/at/2017/at170048.html"},
	"CVE-2017-11907":   {"https://www.jpcert.or.jp/at/2017/at170048.html"},
	"CVE-2017-11908":   {"https://www.jpcert.or.jp/at/2017/at170048.html"},
	"CVE-2017-11909":   {"https://www.jpcert.or.jp/at/2017/at170048.html"},
	"CVE-2017-11910":   {"https://www.jpcert.or.jp/at/2017/at170048.html"},
	"CVE-2017-11911":   {"https://www.jpcert.or.jp/at/2017/at170048.html"},
	"CVE-2017-11912":   {"https://www.jpcert.or.jp/at/2017/at170048.html"},
	"CVE-2017-11914":   {"https://www.jpcert.or.jp/at/2017/at170048.html"},
	"CVE-2017-11918":   {"https://www.jpcert.or.jp/at/2017/at170048.html"},
	"CVE-2017-11930":   {"https://www.jpcert.or.jp/at/2017/at170048.html"},
	"CVE-2017-11937":   {"https://www.jpcert.or.jp/at/2017/at170046.html"},
	"CVE-2017-12615":   {"https://www.jpcert.or.jp/at/2017/at170038.html"},
	"CVE-2017-12616":   {"https://www.jpcert.or.jp/at/2017/at170038.html"},
	"CVE-2017-12617":   {"https://www.jpcert.or.jp/at/2017/at170038.html"},
	"CVE-2017-13872":   {"https://www.jpcert.or.jp/at/2017/at170045.html"},
	"CVE-2017-14315":   {"https://www.jpcert.or.jp/at/2017/at170037.html"},
	"CVE-2017-3135":    {"https://www.jpcert.or.jp/at/2017/at170007.html"},
	"CVE-2017-3136":    {"https://www.jpcert.or.jp/at/2017/at170016.html"},
	"CVE-2017-3137":    {"https://www.jpcert.or.jp/at/2017/at170016.html"},
	"CVE-2017-3138":    {"https://www.jpcert.or.jp/at/2017/at170016.html"},
	"CVE-2017-3142":    {"https://www.jpcert.or.jp/at/2017/at170024.html"},
	"CVE-2017-3143":    {"https://www.jpcert.or.jp/at/2017/at170024.html"},
	"CVE-2017-3145":    {"https://www.jpcert.or.jp/at/2018/at180005.html"},
	"CVE-2017-5638":    {"https://www.jpcert.or.jp/at/2017/at170009.html"},
	"CVE-2017-6753":    {"https://www.jpcert.or.jp/at/2017/at170028.html"},
	"CVE-2017-8463":    {"https://www.jpcert.or.jp/at/2017/at170027.html"},
	"CVE-2017-8464":    {"https://www.jpcert.or.jp/at/2017/at170022.html"},
	"CVE-2017-8496":    {"https://www.jpcert.or.jp/at/2017/at170022.html"},
	"CVE-2017-8497":    {"https://www.jpcert.or.jp/at/2017/at170022.html"},
	"CVE-2017-8499":    {"https://www.jpcert.or.jp/at/2017/at170022.html"},
	"CVE-2017-8517":    {"https://www.jpcert.or.jp/at/2017/at170022.html"},
	"CVE-2017-8520":    {"https://www.jpcert.or.jp/at/2017/at170022.html"},
	"CVE-2017-8522":    {"https://www.jpcert.or.jp/at/2017/at170022.html"},
	"CVE-2017-8524":    {"https://www.jpcert.or.jp/at/2017/at170022.html"},
	"CVE-2017-8527":    {"https://www.jpcert.or.jp/at/2017/at170022.html"},
	"CVE-2017-8528":    {"https://www.jpcert.or.jp/at/2017/at170022.html"},
	"CVE-2017-8543":    {"https://www.jpcert.or.jp/at/2017/at170022.html"},
	"CVE-2017-8548":    {"https://www.jpcert.or.jp/at/2017/at170022.html"},
	"CVE-2017-8549":    {"https://www.jpcert.or.jp/at/2017/at170022.html"},
	"CVE-2017-8584":    {"https://www.jpcert.or.jp/at/2017/at170027.html"},
	"CVE-2017-8589":    {"https://www.jpcert.or.jp/at/2017/at170027.html"},
	"CVE-2017-8591":    {"https://www.jpcert.or.jp/at/2017/at170032.html"},
	"CVE-2017-8594":    {"https://www.jpcert.or.jp/at/2017/at170027.html"},
	"CVE-2017-8595":    {"https://www.jpcert.or.jp/at/2017/at170027.html"},
	"CVE-2017-8596":    {"https://www.jpcert.or.jp/at/2017/at170027.html"},
	"CVE-2017-8598":    {"https://www.jpcert.or.jp/at/2017/at170027.html"},
	"CVE-2017-8601":    {"https://www.jpcert.or.jp/at/2017/at170027.html"},
	"CVE-2017-8603":    {"https://www.jpcert.or.jp/at/2017/at170027.html"},
	"CVE-2017-8604":    {"https://www.jpcert.or.jp/at/2017/at170027.html"},
	"CVE-2017-8605":    {"https://www.jpcert.or.jp/at/2017/at170027.html"},
	"CVE-2017-8606":    {"https://www.jpcert.or.jp/at/2017/at170027.html"},
	"CVE-2017-8607":    {"https://www.jpcert.or.jp/at/2017/at170027.html"},
	"CVE-2017-8608":    {"https://www.jpcert.or.jp/at/2017/at170027.html"},
	"CVE-2017-8609":    {"https://www.jpcert.or.jp/at/2017/at170027.html"},
	"CVE-2017-8610":    {"https://www.jpcert.or.jp/at/2017/at170027.html"},
	"CVE-2017-8617":    {"https://www.jpcert.or.jp/at/2017/at170027.html"},
	"CVE-2017-8618":    {"https://www.jpcert.or.jp/at/2017/at170027.html"},
	"CVE-2017-8619":    {"https://www.jpcert.or.jp/at/2017/at170027.html"},
	"CVE-2017-8620":    {"https://www.jpcert.or.jp/at/2017/at170032.html"},
	"CVE-2017-8622":    {"https://www.jpcert.or.jp/at/2017/at170032.html"},
	"CVE-2017-8628":    {"https://www.jpcert.or.jp/at/2017/at170037.html"},
	"CVE-2017-8634":    {"https://www.jpcert.or.jp/at/2017/at170032.html"},
	"CVE-2017-8635":    {"https://www.jpcert.or.jp/at/2017/at170032.html"},
	"CVE-2017-8636":    {"https://www.jpcert.or.jp/at/2017/at170032.html"},
	"CVE-2017-8638":    {"https://www.jpcert.or.jp/at/2017/at170032.html"},
	"CVE-2017-8639":    {"https://www.jpcert.or.jp/at/2017/at170032.html"},
	"CVE-2017-8640":    {"https://www.jpcert.or.jp/at/2017/at170032.html"},
	"CVE-2017-8641":    {"https://www.jpcert.or.jp/at/2017/at170032.html"},
	"CVE-2017-8645":    {"https://www.jpcert.or.jp/at/2017/at170032.html"},
	"CVE-2017-8646":    {"https://www.jpcert.or.jp/at/2017/at170032.html"},
	"CVE-2017-8647":    {"https://www.jpcert.or.jp/at/2017/at170032.html"},
	"CVE-2017-8649":    {"https://www.jpcert.or.jp/at/2017/at170036.html"},
	"CVE-2017-8653":    {"https://www.jpcert.or.jp/at/2017/at170032.html"},
	"CVE-2017-8655":    {"https://www.jpcert.or.jp/at/2017/at170032.html"},
	"CVE-2017-8656":    {"https://www.jpcert.or.jp/at/2017/at170032.html"},
	"CVE-2017-8657":    {"https://www.jpcert.or.jp/at/2017/at170032.html"},
	"CVE-2017-8660":    {"https://www.jpcert.or.jp/at/2017/at170036.html"},
	"CVE-2017-8661":    {"https://www.jpcert.or.jp/at/2017/at170032.html"},
	"CVE-2017-8669":    {"https://www.jpcert.or.jp/at/2017/at170032.html"},
	"CVE-2017-8670":    {"https://www.jpcert.or.jp/at/2017/at170032.html"},
	"CVE-2017-8671":    {"https://www.jpcert.or.jp/at/2017/at170032.html"},
	"CVE-2017-8672":    {"https://www.jpcert.or.jp/at/2017/at170032.html"},
	"CVE-2017-8674":    {"https://www.jpcert.or.jp/at/2017/at170032.html"},
	"CVE-2017-8676":    {"https://www.jpcert.or.jp/at/2017/at170036.html"},
	"CVE-2017-8682":    {"https://www.jpcert.or.jp/at/2017/at170036.html"},
	"CVE-2017-8686":    {"https://www.jpcert.or.jp/at/2017/at170036.html"},
	"CVE-2017-8696":    {"https://www.jpcert.or.jp/at/2017/at170036.html"},
	"CVE-2017-8727":    {"https://www.jpcert.or.jp/at/2017/at170039.html"},
	"CVE-2017-8728":    {"https://www.jpcert.or.jp/at/2017/at170036.html"},
	"CVE-2017-8729":    {"https://www.jpcert.or.jp/at/2017/at170036.html"},
	"CVE-2017-8731":    {"https://www.jpcert.or.jp/at/2017/at170036.html"},
	"CVE-2017-8734":    {"https://www.jpcert.or.jp/at/2017/at170036.html"},
	"CVE-2017-8737":    {"https://www.jpcert.or.jp/at/2017/at170036.html"},
	"CVE-2017-8738":    {"https://www.jpcert.or.jp/at/2017/at170036.html"},
	"CVE-2017-8740":    {"https://www.jpcert.or.jp/at/2017/at170036.html"},
	"CVE-2017-8741":    {"https://www.jpcert.or.jp/at/2017/at170036.html"},
	"CVE-2017-8747":    {"https://www.jpcert.or.jp/at/2017/at170036.html"},
	"CVE-2017-8748":    {"https://www.jpcert.or.jp/at/2017/at170036.html"},
	"CVE-2017-8749":    {"https://www.jpcert.or.jp/at/2017/at170036.html"},
	"CVE-2017-8750":    {"https://www.jpcert.or.jp/at/2017/at170036.html"},
	"CVE-2017-8751":    {"https://www.jpcert.or.jp/at/2017/at170036.html"},
	"CVE-2017-8752":    {"https://www.jpcert.or.jp/at/2017/at170036.html"},
	"CVE-2017-8753":    {"https://www.jpcert.or.jp/at/2017/at170036.html"},
	"CVE-2017-8755":    {"https://www.jpcert.or.jp/at/2017/at170036.html"},
	"CVE-2017-8756":    {"https://www.jpcert.or.jp/at/2017/at170036.html"},
	"CVE-2017-8757":    {"https://www.jpcert.or.jp/at/2017/at170036.html"},
	"CVE-2017-8759":    {"https://www.jpcert.or.jp/at/2017/at170036.html"},
	"CVE-2017-9791":    {"https://www.jpcert.or.jp/at/2017/at170025.html"},
	"CVE-2017-9805":    {"https://www.jpcert.or.jp/at/2017/at170033.html"},
	"CVE-2018-0171":    {"https://www.jpcert.or.jp/at/2018/at180013.html"},
	"CVE-2018-0758":    {"https://www.jpcert.or.jp/at/2018/at180002.html"},
	"CVE-2018-0762":    {"https://www.jpcert.or.jp/at/2018/at180002.html"},
	"CVE-2018-0763":    {"https://www.jpcert.or.jp/at/2018/at180008.html"},
	"CVE-2018-0767":    {"https://www.jpcert.or.jp/at/2018/at180002.html"},
	"CVE-2018-0769":    {"https://www.jpcert.or.jp/at/2018/at180002.html"},
	"CVE-2018-0770":    {"https://www.jpcert.or.jp/at/2018/at180002.html"},
	"CVE-2018-0772":    {"https://www.jpcert.or.jp/at/2018/at180002.html"},
	"CVE-2018-0773":    {"https://www.jpcert.or.jp/at/2018/at180002.html"},
	"CVE-2018-0774":    {"https://www.jpcert.or.jp/at/2018/at180002.html"},
	"CVE-2018-0775":    {"https://www.jpcert.or.jp/at/2018/at180002.html"},
	"CVE-2018-0776":    {"https://www.jpcert.or.jp/at/2018/at180002.html"},
	"CVE-2018-0777":    {"https://www.jpcert.or.jp/at/2018/at180002.html"},
	"CVE-2018-0778":    {"https://www.jpcert.or.jp/at/2018/at180002.html"},
	"CVE-2018-0780":    {"https://www.jpcert.or.jp/at/2018/at180002.html"},
	"CVE-2018-0781":    {"https://www.jpcert.or.jp/at/2018/at180002.html"},
	"CVE-2018-0797":    {"https://www.jpcert.or.jp/at/2018/at180002.html"},
	"CVE-2018-0800":    {"https://www.jpcert.or.jp/at/2018/at180002.html"},
	"CVE-2018-0802":    {"https://www.jpcert.or.jp/at/2018/at180002.html"},
	"CVE-2018-0825":    {"https://www.jpcert.or.jp/at/2018/at180008.html"},
	"CVE-2018-0834":    {"https://www.jpcert.or.jp/at/2018/at180008.html"},
	"CVE-2018-0835":    {"https://www.jpcert.or.jp/at/2018/at180008.html"},
	"CVE-2018-0837":    {"https://www.jpcert.or.jp/at/2018/at180008.html"},
	"CVE-2018-0838":    {"https://www.jpcert.or.jp/at/2018/at180008.html"},
	"CVE-2018-0840":    {"https://www.jpcert.or.jp/at/2018/at180008.html"},
	"CVE-2018-0852":    {"https://www.jpcert.or.jp/at/2018/at180008.html"},
	"CVE-2018-0856":    {"https://www.jpcert.or.jp/at/2018/at180008.html"},
	"CVE-2018-0857":    {"https://www.jpcert.or.jp/at/2018/at180008.html"},
	"CVE-2018-0859":    {"https://www.jpcert.or.jp/at/2018/at180008.html"},
	"CVE-2018-0860":    {"https://www.jpcert.or.jp/at/2018/at180008.html"},
	"CVE-2018-0861":    {"https://www.jpcert.or.jp/at/2018/at180008.html"},
	"CVE-2018-0870":    {"https://www.jpcert.or.jp/at/2018/at180016.html"},
	"CVE-2018-0872":    {"https://www.jpcert.or.jp/at/2018/at180011.html"},
	"CVE-2018-0874":    {"https://www.jpcert.or.jp/at/2018/at180011.html"},
	"CVE-2018-0876":    {"https://www.jpcert.or.jp/at/2018/at180011.html"},
	"CVE-2018-0889":    {"https://www.jpcert.or.jp/at/2018/at180011.html"},
	"CVE-2018-0893":    {"https://www.jpcert.or.jp/at/2018/at180011.html"},
	"CVE-2018-0930":    {"https://www.jpcert.or.jp/at/2018/at180011.html"},
	"CVE-2018-0931":    {"https://www.jpcert.or.jp/at/2018/at180011.html"},
	"CVE-2018-0932":    {"https://www.jpcert.or.jp/at/2018/at180011.html"},
	"CVE-2018-0933":    {"https://www.jpcert.or.jp/at/2018/at180011.html"},
	"CVE-2018-0934":    {"https://www.jpcert.or.jp/at/2018/at180011.html"},
	"CVE-2018-0936":    {"https://www.jpcert.or.jp/at/2018/at180011.html"},
	"CVE-2018-0937":    {"https://www.jpcert.or.jp/at/2018/at180011.html"},
	"CVE-2018-0939":    {"https://www.jpcert.or.jp/at/2018/at180011.html"},
	"CVE-2018-0943":    {"https://www.jpcert.or.jp/at/2018/at180021.html"},
	"CVE-2018-0945":    {"https://www.jpcert.or.jp/at/2018/at180021.html"},
	"CVE-2018-0946":    {"https://www.jpcert.or.jp/at/2018/at180021.html"},
	"CVE-2018-0950":    {"https://www.jpcert.or.jp/at/2018/at180016.html"},
	"CVE-2018-0951":    {"https://www.jpcert.or.jp/at/2018/at180021.html"},
	"CVE-2018-0953":    {"https://www.jpcert.or.jp/at/2018/at180021.html"},
	"CVE-2018-0954":    {"https://www.jpcert.or.jp/at/2018/at180021.html"},
	"CVE-2018-0955":    {"https://www.jpcert.or.jp/at/2018/at180021.html"},
	"CVE-2018-0959":    {"https://www.jpcert.or.jp/at/2018/at180021.html"},
	"CVE-2018-0961":    {"https://www.jpcert.or.jp/at/2018/at180021.html"},
	"CVE-2018-0965":    {"https://www.jpcert.or.jp/at/2018/at180038.html"},
	"CVE-2018-0979":    {"https://www.jpcert.or.jp/at/2018/at180016.html"},
	"CVE-2018-0980":    {"https://www.jpcert.or.jp/at/2018/at180016.html"},
	"CVE-2018-0981":    {"https://www.jpcert.or.jp/at/2018/at180016.html"},
	"CVE-2018-0988":    {"https://www.jpcert.or.jp/at/2018/at180016.html"},
	"CVE-2018-0990":    {"https://www.jpcert.or.jp/at/2018/at180016.html"},
	"CVE-2018-0991":    {"https://www.jpcert.or.jp/at/2018/at180016.html"},
	"CVE-2018-0993":    {"https://www.jpcert.or.jp/at/2018/at180016.html"},
	"CVE-2018-0994":    {"https://www.jpcert.or.jp/at/2018/at180016.html"},
	"CVE-2018-0995":    {"https://www.jpcert.or.jp/at/2018/at180016.html"},
	"CVE-2018-0996":    {"https://www.jpcert.or.jp/at/2018/at180016.html"},
	"CVE-2018-1000":    {"https://www.jpcert.or.jp/at/2018/at180016.html"},
	"CVE-2018-1004":    {"https://www.jpcert.or.jp/at/2018/at180016.html"},
	"CVE-2018-1010":    {"https://www.jpcert.or.jp/at/2018/at180016.html"},
	"CVE-2018-1012":    {"https://www.jpcert.or.jp/at/2018/at180016.html"},
	"CVE-2018-1013":    {"https://www.jpcert.or.jp/at/2018/at180016.html"},
	"CVE-2018-1015":    {"https://www.jpcert.or.jp/at/2018/at180016.html"},
	"CVE-2018-1016":    {"https://www.jpcert.or.jp/at/2018/at180016.html"},
	"CVE-2018-1018":    {"https://www.jpcert.or.jp/at/2018/at180016.html"},
	"CVE-2018-1019":    {"https://www.jpcert.or.jp/at/2018/at180016.html"},
	"CVE-2018-1020":    {"https://www.jpcert.or.jp/at/2018/at180016.html"},
	"CVE-2018-1022":    {"https://www.jpcert.or.jp/at/2018/at180021.html"},
	"CVE-2018-1023":    {"https://www.jpcert.or.jp/at/2018/at180016.html"},
	"CVE-2018-11776":   {"https://www.jpcert.or.jp/at/2018/at180036.html"},
	"CVE-2018-1270":    {"https://www.jpcert.or.jp/at/2018/at180014.html"},
	"CVE-2018-1271":    {"https://www.jpcert.or.jp/at/2018/at180014.html"},
	"CVE-2018-1272":    {"https://www.jpcert.or.jp/at/2018/at180014.html"},
	"CVE-2018-1273":    {"https://www.jpcert.or.jp/at/2018/at180017.html"},
	"CVE-2018-1274":    {"https://www.jpcert.or.jp/at/2018/at180017.html"},
	"CVE-2018-1275":    {"https://www.jpcert.or.jp/at/2018/at180014.html"},
	"CVE-2018-12794":   {"https://www.jpcert.or.jp/at/2018/at180039.html"},
	"CVE-2018-1336":    {"https://www.jpcert.or.jp/at/2018/at180030.html"},
	"CVE-2018-15442":   {"https://www.jpcert.or.jp/at/2018/at180043.html"},
	"CVE-2018-15979":   {"https://www.jpcert.or.jp/at/2018/at180045.html"},
	"CVE-2018-2628":    {"https://www.jpcert.or.jp/at/2018/at180029.html"},
	"CVE-2018-2893":    {"https://www.jpcert.or.jp/at/2018/at180029.html"},
	"CVE-2018-2894":    {"https://www.jpcert.or.jp/at/2018/at180029.html"},
	"CVE-2018-2933":    {"https://www.jpcert.or.jp/at/2018/at180029.html"},
	"CVE-2018-2983":    {"https://www.jpcert.or.jp/at/2018/at180029.html"},
	"CVE-2018-2998":    {"https://www.jpcert.or.jp/at/2018/at180029.html"},
	"CVE-2018-4877":    {"https://www.jpcert.or.jp/at/2018/at180008.html"},
	"CVE-2018-4878":    {"https://www.jpcert.or.jp/at/2018/at180006.html", "https://www.jpcert.or.jp/at/2018/at180008.html"},
	"CVE-2018-4945":    {"https://www.jpcert.or.jp/at/2018/at180025.html"},
	"CVE-2018-5000":    {"https://www.jpcert.or.jp/at/2018/at180025.html"},
	"CVE-2018-5001":    {"https://www.jpcert.or.jp/at/2018/at180025.html"},
	"CVE-2018-5002":    {"https://www.jpcert.or.jp/at/2018/at180024.html", "https://www.jpcert.or.jp/at/2018/at180025.html"},
	"CVE-2018-5740":    {"https://www.jpcert.or.jp/at/2018/at180031.html"},
	"CVE-2018-7600":    {"https://www.jpcert.or.jp/at/2018/at180012.html"},
	"CVE-2018-7602":    {"https://www.jpcert.or.jp/at/2018/at180019.html"},
	"CVE-2018-8034":    {"https://www.jpcert.or.jp/at/2018/at180030.html"},
	"CVE-2018-8037":    {"https://www.jpcert.or.jp/at/2018/at180030.html"},
	"CVE-2018-8110":    {"https://www.jpcert.or.jp/at/2018/at180025.html"},
	"CVE-2018-8111":    {"https://www.jpcert.or.jp/at/2018/at180025.html"},
	"CVE-2018-8114":    {"https://www.jpcert.or.jp/at/2018/at180021.html"},
	"CVE-2018-8120":    {"https://www.jpcert.or.jp/at/2018/at180021.html"},
	"CVE-2018-8122":    {"https://www.jpcert.or.jp/at/2018/at180021.html"},
	"CVE-2018-8128":    {"https://www.jpcert.or.jp/at/2018/at180021.html"},
	"CVE-2018-8130":    {"https://www.jpcert.or.jp/at/2018/at180021.html"},
	"CVE-2018-8133":    {"https://www.jpcert.or.jp/at/2018/at180021.html"},
	"CVE-2018-8137":    {"https://www.jpcert.or.jp/at/2018/at180021.html"},
	"CVE-2018-8139":    {"https://www.jpcert.or.jp/at/2018/at180021.html"},
	"CVE-2018-8154":    {"https://www.jpcert.or.jp/at/2018/at180021.html"},
	"CVE-2018-8174":    {"https://www.jpcert.or.jp/at/2018/at180021.html"},
	"CVE-2018-8178":    {"https://www.jpcert.or.jp/at/2018/at180021.html"},
	"CVE-2018-8213":    {"https://www.jpcert.or.jp/at/2018/at180025.html"},
	"CVE-2018-8225":    {"https://www.jpcert.or.jp/at/2018/at180025.html"},
	"CVE-2018-8229":    {"https://www.jpcert.or.jp/at/2018/at180025.html"},
	"CVE-2018-8231":    {"https://www.jpcert.or.jp/at/2018/at180025.html"},
	"CVE-2018-8236":    {"https://www.jpcert.or.jp/at/2018/at180025.html"},
	"CVE-2018-8242":    {"https://www.jpcert.or.jp/at/2018/at180028.html"},
	"CVE-2018-8249":    {"https://www.jpcert.or.jp/at/2018/at180025.html"},
	"CVE-2018-8251":    {"https://www.jpcert.or.jp/at/2018/at180025.html"},
	"CVE-2018-8262":    {"https://www.jpcert.or.jp/at/2018/at180028.html"},
	"CVE-2018-8266":    {"https://www.jpcert.or.jp/at/2018/at180034.html"},
	"CVE-2018-8267":    {"https://www.jpcert.or.jp/at/2018/at180025.html"},
	"CVE-2018-8273":    {"https://www.jpcert.or.jp/at/2018/at180034.html"},
	"CVE-2018-8274":    {"https://www.jpcert.or.jp/at/2018/at180028.html"},
	"CVE-2018-8275":    {"https://www.jpcert.or.jp/at/2018/at180028.html"},
	"CVE-2018-8279":    {"https://www.jpcert.or.jp/at/2018/at180028.html"},
	"CVE-2018-8280":    {"https://www.jpcert.or.jp/at/2018/at180028.html"},
	"CVE-2018-8286":    {"https://www.jpcert.or.jp/at/2018/at180028.html"},
	"CVE-2018-8288":    {"https://www.jpcert.or.jp/at/2018/at180028.html"},
	"CVE-2018-8290":    {"https://www.jpcert.or.jp/at/2018/at180028.html"},
	"CVE-2018-8291":    {"https://www.jpcert.or.jp/at/2018/at180028.html"},
	"CVE-2018-8294":    {"https://www.jpcert.or.jp/at/2018/at180028.html"},
	"CVE-2018-8296":    {"https://www.jpcert.or.jp/at/2018/at180028.html"},
	"CVE-2018-8301":    {"https://www.jpcert.or.jp/at/2018/at180028.html"},
	"CVE-2018-8302":    {"https://www.jpcert.or.jp/at/2018/at180034.html"},
	"CVE-2018-8324":    {"https://www.jpcert.or.jp/at/2018/at180028.html"},
	"CVE-2018-8332":    {"https://www.jpcert.or.jp/at/2018/at180038.html"},
	"CVE-2018-8344":    {"https://www.jpcert.or.jp/at/2018/at180034.html"},
	"CVE-2018-8345":    {"https://www.jpcert.or.jp/at/2018/at180034.html"},
	"CVE-2018-8350":    {"https://www.jpcert.or.jp/at/2018/at180034.html"},
	"CVE-2018-8355":    {"https://www.jpcert.or.jp/at/2018/at180034.html"},
	"CVE-2018-8367":    {"https://www.jpcert.or.jp/at/2018/at180038.html"},
	"CVE-2018-8371":    {"https://www.jpcert.or.jp/at/2018/at180034.html"},
	"CVE-2018-8372":    {"https://www.jpcert.or.jp/at/2018/at180034.html"},
	"CVE-2018-8373":    {"https://www.jpcert.or.jp/at/2018/at180034.html"},
	"CVE-2018-8377":    {"https://www.jpcert.or.jp/at/2018/at180034.html"},
	"CVE-2018-8380":    {"https://www.jpcert.or.jp/at/2018/at180034.html"},
	"CVE-2018-8381":    {"https://www.jpcert.or.jp/at/2018/at180034.html"},
	"CVE-2018-8385":    {"https://www.jpcert.or.jp/at/2018/at180034.html"},
	"CVE-2018-8387":    {"https://www.jpcert.or.jp/at/2018/at180034.html"},
	"CVE-2018-8390":    {"https://www.jpcert.or.jp/at/2018/at180034.html"},
	"CVE-2018-8397":    {"https://www.jpcert.or.jp/at/2018/at180034.html"},
	"CVE-2018-8403":    {"https://www.jpcert.or.jp/at/2018/at180034.html"},
	"CVE-2018-8414":    {"https://www.jpcert.or.jp/at/2018/at180034.html"},
	"CVE-2018-8420":    {"https://www.jpcert.or.jp/at/2018/at180038.html"},
	"CVE-2018-8421":    {"https://www.jpcert.or.jp/at/2018/at180038.html"},
	"CVE-2018-8439":    {"https://www.jpcert.or.jp/at/2018/at180038.html"},
	"CVE-2018-8440":    {"https://www.jpcert.or.jp/at/2018/at180038.html"},
	"CVE-2018-8447":    {"https://www.jpcert.or.jp/at/2018/at180038.html"},
	"CVE-2018-8453":    {"https://www.jpcert.or.jp/at/2018/at180041.html"},
	"CVE-2018-8456":    {"https://www.jpcert.or.jp/at/2018/at180038.html"},
	"CVE-2018-8457":    {"https://www.jpcert.or.jp/at/2018/at180038.html"},
	"CVE-2018-8459":    {"https://www.jpcert.or.jp/at/2018/at180038.html"},
	"CVE-2018-8460":    {"https://www.jpcert.or.jp/at/2018/at180041.html"},
	"CVE-2018-8461":    {"https://www.jpcert.or.jp/at/2018/at180038.html"},
	"CVE-2018-8464":    {"https://www.jpcert.or.jp/at/2018/at180038.html"},
	"CVE-2018-8465":    {"https://www.jpcert.or.jp/at/2018/at180038.html"},
	"CVE-2018-8466":    {"https://www.jpcert.or.jp/at/2018/at180038.html"},
	"CVE-2018-8467":    {"https://www.jpcert.or.jp/at/2018/at180038.html"},
	"CVE-2018-8473":    {"https://www.jpcert.or.jp/at/2018/at180041.html"},
	"CVE-2018-8475":    {"https://www.jpcert.or.jp/at/2018/at180038.html"},
	"CVE-2018-8476":    {"https://www.jpcert.or.jp/at/2018/at180046.html"},
	"CVE-2018-8489":    {"https://www.jpcert.or.jp/at/2018/at180041.html"},
	"CVE-2018-8490":    {"https://www.jpcert.or.jp/at/2018/at180041.html"},
	"CVE-2018-8491":    {"https://www.jpcert.or.jp/at/2018/at180041.html"},
	"CVE-2018-8494":    {"https://www.jpcert.or.jp/at/2018/at180041.html"},
	"CVE-2018-8505":    {"https://www.jpcert.or.jp/at/2018/at180041.html"},
	"CVE-2018-8509":    {"https://www.jpcert.or.jp/at/2018/at180041.html"},
	"CVE-2018-8510":    {"https://www.jpcert.or.jp/at/2018/at180041.html"},
	"CVE-2018-8511":    {"https://www.jpcert.or.jp/at/2018/at180041.html"},
	"CVE-2018-8513":    {"https://www.jpcert.or.jp/at/2018/at180041.html"},
	"CVE-2018-8541":    {"https://www.jpcert.or.jp/at/2018/at180046.html"},
	"CVE-2018-8542":    {"https://www.jpcert.or.jp/at/2018/at180046.html"},
	"CVE-2018-8543":    {"https://www.jpcert.or.jp/at/2018/at180046.html"},
	"CVE-2018-8544":    {"https://www.jpcert.or.jp/at/2018/at180046.html"},
	"CVE-2018-8551":    {"https://www.jpcert.or.jp/at/2018/at180046.html"},
	"CVE-2018-8553":    {"https://www.jpcert.or.jp/at/2018/at180046.html"},
	"CVE-2018-8555":    {"https://www.jpcert.or.jp/at/2018/at180046.html"},
	"CVE-2018-8556":    {"https://www.jpcert.or.jp/at/2018/at180046.html"},
	"CVE-2018-8557":    {"https://www.jpcert.or.jp/at/2018/at180046.html"},
	"CVE-2018-8588":    {"https://www.jpcert.or.jp/at/2018/at180046.html"},
	"CVE-2018-8589":    {"https://www.jpcert.or.jp/at/2018/at180046.html"},
	"CVE-2018-8609":    {"https://www.jpcert.or.jp/at/2018/at180046.html"},
}
