go-log-benchmarks
=================

A go logging library benchmarks for me.

## Benchmark result

Note these benchmarks print different outputs.
Its purpose is not to choose the most performant library, but to get rough cost
of each operation in each library.

Especially BenchmarkZapLTSVProductionLog uses
[zapcore.EpochTimeEncoder](https://godoc.org/go.uber.org/zap/zapcore#EpochTimeEncoder).
It prints time as floating-point number of seconds since the Unix epoch, and this is a
low cost operation compared to printing time in ISO8609 format.

Other benchmarks (BenchmarkLTSVLog, BenchmarkStandardLog and BenchmarkZapLTSVDevelopmentLog)
uses ISO8609 time format, though there is a slight difference.
BenchmarkZapLTSVDevelopmentLog uses [zapcore.ISO8601TimeEncoder](https://godoc.org/go.uber.org/zap/zapcore#ISO8601TimeEncoder) which prints times with millisecond precision.
The other two prints times with microsecond precision.

```
$ go test -count=10 -bench . -benchmem -cpuprofile=cpu.prof | tee 20200818.log
goos: darwin
goarch: amd64
pkg: github.com/jonsen/go-log-benchmarks
BenchmarkLTSVLog-8                       	  330307	      3804 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-8                       	  335334	      3724 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-8                       	  331137	      3858 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-8                       	  335247	      3700 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-8                       	  330397	      4868 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-8                       	  339608	      3706 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-8                       	  332620	      3762 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-8                       	  358716	      3763 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-8                       	  337641	      3713 ns/op	       0 B/op	       0 allocs/op
BenchmarkLTSVLog-8                       	  324283	      3796 ns/op	       0 B/op	       0 allocs/op
BenchmarkStandardLog-8                   	  289574	      3892 ns/op	      64 B/op	       1 allocs/op
BenchmarkStandardLog-8                   	  314158	      3852 ns/op	      64 B/op	       1 allocs/op
BenchmarkStandardLog-8                   	  320150	      4252 ns/op	      64 B/op	       1 allocs/op
BenchmarkStandardLog-8                   	  311092	      4299 ns/op	      64 B/op	       1 allocs/op
BenchmarkStandardLog-8                   	  313514	      4159 ns/op	      64 B/op	       1 allocs/op
BenchmarkStandardLog-8                   	  326455	      6714 ns/op	      64 B/op	       1 allocs/op
BenchmarkStandardLog-8                   	  292818	      4434 ns/op	      64 B/op	       1 allocs/op
BenchmarkStandardLog-8                   	  300136	      4158 ns/op	      64 B/op	       1 allocs/op
BenchmarkStandardLog-8                   	  297190	      4239 ns/op	      64 B/op	       1 allocs/op
BenchmarkStandardLog-8                   	  284702	      3960 ns/op	      64 B/op	       1 allocs/op
BenchmarkZapJSONProductionLog-8          	 3775588	       305 ns/op	     131 B/op	       1 allocs/op
BenchmarkZapJSONProductionLog-8          	 3913600	       309 ns/op	     131 B/op	       1 allocs/op
BenchmarkZapJSONProductionLog-8          	 3834273	       329 ns/op	     131 B/op	       1 allocs/op
BenchmarkZapJSONProductionLog-8          	 3584894	       311 ns/op	     131 B/op	       1 allocs/op
BenchmarkZapJSONProductionLog-8          	 3972762	       445 ns/op	     131 B/op	       1 allocs/op
BenchmarkZapJSONProductionLog-8          	 3350910	       346 ns/op	     131 B/op	       1 allocs/op
BenchmarkZapJSONProductionLog-8          	 3870486	       309 ns/op	     131 B/op	       1 allocs/op
BenchmarkZapJSONProductionLog-8          	 3912213	       312 ns/op	     131 B/op	       1 allocs/op
BenchmarkZapJSONProductionLog-8          	 3936358	       304 ns/op	     131 B/op	       1 allocs/op
BenchmarkZapJSONProductionLog-8          	 3536766	       324 ns/op	     131 B/op	       1 allocs/op
BenchmarkZapJSONDevelopmentLog-8         	  165543	      6806 ns/op	     521 B/op	       9 allocs/op
BenchmarkZapJSONDevelopmentLog-8         	  173355	      7103 ns/op	     521 B/op	       9 allocs/op
BenchmarkZapJSONDevelopmentLog-8         	  176449	      7327 ns/op	     521 B/op	       9 allocs/op
BenchmarkZapJSONDevelopmentLog-8         	  177573	      7587 ns/op	     521 B/op	       9 allocs/op
BenchmarkZapJSONDevelopmentLog-8         	  169308	      9367 ns/op	     521 B/op	       9 allocs/op
BenchmarkZapJSONDevelopmentLog-8         	  175758	      8984 ns/op	     521 B/op	       9 allocs/op
BenchmarkZapJSONDevelopmentLog-8         	  163735	      7116 ns/op	     521 B/op	       9 allocs/op
BenchmarkZapJSONDevelopmentLog-8         	  164386	      6878 ns/op	     521 B/op	       9 allocs/op
BenchmarkZapJSONDevelopmentLog-8         	  175543	      6840 ns/op	     521 B/op	       9 allocs/op
BenchmarkZapJSONDevelopmentLog-8         	  170352	      7511 ns/op	     521 B/op	       9 allocs/op
BenchmarkZapLTSVProductionLog-8          	 4010349	       304 ns/op	     131 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-8          	 3921084	       296 ns/op	     131 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-8          	 3925057	       312 ns/op	     131 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-8          	 3788397	       305 ns/op	     131 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-8          	 3850113	       309 ns/op	     131 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-8          	 3844656	       304 ns/op	     131 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-8          	 3676774	       315 ns/op	     131 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-8          	 3816120	       300 ns/op	     131 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-8          	 3930429	       310 ns/op	     131 B/op	       1 allocs/op
BenchmarkZapLTSVProductionLog-8          	 3986175	       299 ns/op	     131 B/op	       1 allocs/op
BenchmarkZapLTSVDevelopmentLog-8         	  169334	      7790 ns/op	     425 B/op	       5 allocs/op
BenchmarkZapLTSVDevelopmentLog-8         	  163587	     10498 ns/op	     425 B/op	       5 allocs/op
BenchmarkZapLTSVDevelopmentLog-8         	  171408	      8696 ns/op	     425 B/op	       5 allocs/op
BenchmarkZapLTSVDevelopmentLog-8         	  171876	     13593 ns/op	     425 B/op	       5 allocs/op
BenchmarkZapLTSVDevelopmentLog-8         	  175699	      8024 ns/op	     425 B/op	       5 allocs/op
BenchmarkZapLTSVDevelopmentLog-8         	  163057	      6952 ns/op	     425 B/op	       5 allocs/op
BenchmarkZapLTSVDevelopmentLog-8         	  172374	      7168 ns/op	     425 B/op	       5 allocs/op
BenchmarkZapLTSVDevelopmentLog-8         	  163710	      6462 ns/op	     425 B/op	       5 allocs/op
BenchmarkZapLTSVDevelopmentLog-8         	  167438	      6853 ns/op	     424 B/op	       5 allocs/op
BenchmarkZapLTSVDevelopmentLog-8         	  176223	      7090 ns/op	     424 B/op	       5 allocs/op
BenchmarkZerologTimestamp-8              	  289047	      7064 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-8              	  323029	      3798 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-8              	  315751	      4115 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-8              	  321730	      6111 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-8              	  234276	      4687 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-8              	  328470	      4917 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-8              	  345350	      3856 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-8              	  326190	      4024 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-8              	  313056	      3888 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimestamp-8              	  347779	      4089 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-8   	  324110	      3536 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-8   	  343660	      4200 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-8   	  348852	      5473 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-8   	  324794	      5046 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-8   	  351553	      5847 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-8   	  331746	      4514 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-8   	  339765	      4548 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-8   	  336337	      4190 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-8   	  354754	      4798 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTimeSecondsFromEpoch-8   	  344282	      5206 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-8            	  334922	      3905 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-8            	  305438	      3722 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-8            	  340018	      3816 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-8            	  359914	      4482 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-8            	  329959	      4726 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-8            	  265886	      8268 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-8            	  339555	      4964 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-8            	  327349	      4812 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-8            	  333577	      4965 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339Time-8            	  320234	      4776 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-8        	  316159	     13877 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-8        	  328036	      3592 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-8        	  348864	      3557 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-8        	  355359	      4244 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-8        	  307557	      3852 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-8        	  320449	      3995 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-8        	  349984	      3797 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-8        	  353186	      3588 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-8        	  347164	      3820 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologRFC3339NanoTime-8        	  344419	      3759 ns/op	       0 B/op	       0 allocs/op
BenchmarkIrisGolog-8                     	  104558	     13962 ns/op	      64 B/op	       3 allocs/op
BenchmarkIrisGolog-8                     	   94627	     14036 ns/op	      64 B/op	       3 allocs/op
BenchmarkIrisGolog-8                     	  100478	     13622 ns/op	      64 B/op	       3 allocs/op
BenchmarkIrisGolog-8                     	   91135	     13362 ns/op	      64 B/op	       3 allocs/op
BenchmarkIrisGolog-8                     	   90141	     13833 ns/op	      64 B/op	       3 allocs/op
BenchmarkIrisGolog-8                     	   99751	     13209 ns/op	      64 B/op	       3 allocs/op
BenchmarkIrisGolog-8                     	  104064	     13747 ns/op	      64 B/op	       3 allocs/op
BenchmarkIrisGolog-8                     	  103485	     15244 ns/op	      64 B/op	       3 allocs/op
BenchmarkIrisGolog-8                     	   93470	     13137 ns/op	      64 B/op	       3 allocs/op
BenchmarkIrisGolog-8                     	  100106	     13623 ns/op	      64 B/op	       3 allocs/op
BenchmarkIrisGologJSON-8                 	   51818	     23636 ns/op	    4088 B/op	      38 allocs/op
BenchmarkIrisGologJSON-8                 	   52756	     24225 ns/op	    4089 B/op	      38 allocs/op
BenchmarkIrisGologJSON-8                 	   51949	     23318 ns/op	    4089 B/op	      38 allocs/op
BenchmarkIrisGologJSON-8                 	   51102	     24201 ns/op	    4089 B/op	      38 allocs/op
BenchmarkIrisGologJSON-8                 	   52485	     27865 ns/op	    4088 B/op	      38 allocs/op
BenchmarkIrisGologJSON-8                 	   52468	     23818 ns/op	    4089 B/op	      38 allocs/op
BenchmarkIrisGologJSON-8                 	   52318	     22450 ns/op	    4089 B/op	      38 allocs/op
BenchmarkIrisGologJSON-8                 	   51753	     23515 ns/op	    4088 B/op	      38 allocs/op
BenchmarkIrisGologJSON-8                 	   52064	     22774 ns/op	    4089 B/op	      38 allocs/op
BenchmarkIrisGologJSON-8                 	   52213	     24227 ns/op	    4089 B/op	      38 allocs/op
BenchmarkLTSVLogErr/StackAndTime-8       	   93499	     13049 ns/op	   16656 B/op	       7 allocs/op
BenchmarkLTSVLogErr/StackAndTime-8       	   90901	     13077 ns/op	   16656 B/op	       7 allocs/op
BenchmarkLTSVLogErr/StackAndTime-8       	   93366	     13282 ns/op	   16656 B/op	       7 allocs/op
BenchmarkLTSVLogErr/StackAndTime-8       	   93166	     12973 ns/op	   16656 B/op	       7 allocs/op
BenchmarkLTSVLogErr/StackAndTime-8       	   92756	     15776 ns/op	   16656 B/op	       7 allocs/op
BenchmarkLTSVLogErr/StackAndTime-8       	   93685	     13210 ns/op	   16656 B/op	       7 allocs/op
BenchmarkLTSVLogErr/StackAndTime-8       	   91880	     15226 ns/op	   16656 B/op	       7 allocs/op
BenchmarkLTSVLogErr/StackAndTime-8       	   92354	     13020 ns/op	   16656 B/op	       7 allocs/op
BenchmarkLTSVLogErr/StackAndTime-8       	   92228	     21373 ns/op	   16656 B/op	       7 allocs/op
BenchmarkLTSVLogErr/StackAndTime-8       	   84943	     13545 ns/op	   16656 B/op	       7 allocs/op
BenchmarkLTSVLogErr/Stack-8              	   92756	     14317 ns/op	   16656 B/op	       7 allocs/op
BenchmarkLTSVLogErr/Stack-8              	   80923	     13766 ns/op	   16656 B/op	       7 allocs/op
BenchmarkLTSVLogErr/Stack-8              	   90768	     13312 ns/op	   16656 B/op	       7 allocs/op
BenchmarkLTSVLogErr/Stack-8              	   94726	     12630 ns/op	   16656 B/op	       7 allocs/op
BenchmarkLTSVLogErr/Stack-8              	   97224	     15659 ns/op	   16656 B/op	       7 allocs/op
BenchmarkLTSVLogErr/Stack-8              	   92446	     13529 ns/op	   16656 B/op	       7 allocs/op
BenchmarkLTSVLogErr/Stack-8              	   84770	     15675 ns/op	   16656 B/op	       7 allocs/op
BenchmarkLTSVLogErr/Stack-8              	   90300	     13304 ns/op	   16656 B/op	       7 allocs/op
BenchmarkLTSVLogErr/Stack-8              	   94166	     13035 ns/op	   16656 B/op	       7 allocs/op
BenchmarkLTSVLogErr/Stack-8              	   97312	     12836 ns/op	   16656 B/op	       7 allocs/op
BenchmarkLTSVLogErr/Time-8               	  168440	      7841 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-8               	  151771	      7894 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-8               	  163294	      7929 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-8               	  146517	      7700 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-8               	  159069	      8071 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-8               	  153927	      9610 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-8               	  150934	     10751 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-8               	  153042	      7346 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-8               	  164715	      7698 ns/op	   16464 B/op	       4 allocs/op
BenchmarkLTSVLogErr/Time-8               	  166562	      7513 ns/op	   16464 B/op	       4 allocs/op
PASS
ok  	github.com/jonsen/go-log-benchmarks	236.945s
```

```
$ benchstat 20200818.log
name                           time/op
LTSVLog-8                      3.76µs ± 3%
StandardLog-8                  4.14µs ± 7%
ZapJSONProductionLog-8          317ns ± 9%
ZapJSONDevelopmentLog-8        7.15µs ± 6%
ZapLTSVProductionLog-8          305ns ± 3%
ZapLTSVDevelopmentLog-8        7.73µs ±36%
ZerologTimestamp-8             4.39µs ±39%
ZerologTimeSecondsFromEpoch-8  4.74µs ±25%
ZerologRFC3339Time-8           4.46µs ±17%
ZerologRFC3339NanoTime-8       3.80µs ±12%
IrisGolog-8                    13.6µs ± 4%
IrisGologJSON-8                23.6µs ± 5%
LTSVLogErr/StackAndTime-8      13.7µs ±15%
LTSVLogErr/Stack-8             13.8µs ±14%
LTSVLogErr/Time-8              7.75µs ± 5%

name                           alloc/op
LTSVLog-8                       0.00B
StandardLog-8                   64.0B ± 0%
ZapJSONProductionLog-8           131B ± 0%
ZapJSONDevelopmentLog-8          521B ± 0%
ZapLTSVProductionLog-8           131B ± 0%
ZapLTSVDevelopmentLog-8          425B ± 0%
ZerologTimestamp-8              0.00B
ZerologTimeSecondsFromEpoch-8   0.00B
ZerologRFC3339Time-8            0.00B
ZerologRFC3339NanoTime-8        0.00B
IrisGolog-8                     64.0B ± 0%
IrisGologJSON-8                4.09kB ± 0%
LTSVLogErr/StackAndTime-8      16.7kB ± 0%
LTSVLogErr/Stack-8             16.7kB ± 0%
LTSVLogErr/Time-8              16.5kB ± 0%

name                           allocs/op
LTSVLog-8                        0.00
StandardLog-8                    1.00 ± 0%
ZapJSONProductionLog-8           1.00 ± 0%
ZapJSONDevelopmentLog-8          9.00 ± 0%
ZapLTSVProductionLog-8           1.00 ± 0%
ZapLTSVDevelopmentLog-8          5.00 ± 0%
ZerologTimestamp-8               0.00
ZerologTimeSecondsFromEpoch-8    0.00
ZerologRFC3339Time-8             0.00
ZerologRFC3339NanoTime-8         0.00
IrisGolog-8                      3.00 ± 0%
IrisGologJSON-8                  38.0 ± 0%
LTSVLogErr/StackAndTime-8        7.00 ± 0%
LTSVLogErr/Stack-8               7.00 ± 0%
LTSVLogErr/Time-8                4.00 ± 0%
```

```
$ benchstat new.log 20170922.log | tee benchstat-20170922.log
name                           old time/op    new time/op    delta
LTSVLog-2                        2.11µs ± 0%    2.08µs ± 1%  -1.05%  (p=0.000 n=10+10)
StandardLog-2                    2.66µs ± 1%    2.66µs ± 1%    ~     (p=0.858 n=9+10)
ZapJSONProductionLog-2            410ns ±14%     393ns ±10%    ~     (p=0.148 n=10+10)
ZapJSONDevelopmentLog-2          8.14µs ± 3%    7.91µs ± 3%  -2.79%  (p=0.002 n=10+10)
ZapLTSVProductionLog-2            430ns ± 5%     391ns ±13%  -9.07%  (p=0.000 n=9+10)
ZapLTSVDevelopmentLog-2          8.49µs ± 2%    8.37µs ± 3%  -1.39%  (p=0.027 n=10+10)
ZerologTimestamp-2               2.21µs ± 2%    2.19µs ± 1%  -1.17%  (p=0.000 n=9+10)
ZerologTimeSecondsFromEpoch-2    1.94µs ± 1%    1.94µs ± 1%    ~     (p=0.929 n=10+10)
ZerologRFC3339Time-2             1.95µs ± 1%    1.97µs ± 1%  +1.18%  (p=0.002 n=10+9)
ZerologRFC3339NanoTime-2         1.96µs ± 1%    1.98µs ± 1%  +0.84%  (p=0.004 n=10+10)
LTSVLogErr/StackAndTime-2        17.8µs ± 2%    17.9µs ± 1%    ~     (p=0.353 n=10+10)
LTSVLogErr/Stack-2               17.3µs ± 2%    17.3µs ± 2%    ~     (p=0.720 n=9+10)
LTSVLogErr/Time-2                9.15µs ± 1%    9.13µs ± 1%    ~     (p=0.400 n=9+10)

name                           old alloc/op   new alloc/op   delta
LTSVLog-2                         0.00B          0.00B         ~     (all equal)
StandardLog-2                     64.0B ± 0%     64.0B ± 0%    ~     (all equal)
ZapJSONProductionLog-2             128B ± 0%      128B ± 0%    ~     (all equal)
ZapJSONDevelopmentLog-2            288B ± 0%      288B ± 0%    ~     (all equal)
ZapLTSVProductionLog-2             128B ± 0%      128B ± 0%    ~     (all equal)
ZapLTSVDevelopmentLog-2            208B ± 0%      208B ± 0%    ~     (all equal)
ZerologTimestamp-2                0.00B          0.00B         ~     (all equal)
ZerologTimeSecondsFromEpoch-2     0.00B          0.00B         ~     (all equal)
ZerologRFC3339Time-2              0.00B          0.00B         ~     (all equal)
ZerologRFC3339NanoTime-2          0.00B          0.00B         ~     (all equal)
LTSVLogErr/StackAndTime-2        16.5kB ± 0%    16.5kB ± 0%    ~     (all equal)
LTSVLogErr/Stack-2               16.5kB ± 0%    16.5kB ± 0%    ~     (all equal)
LTSVLogErr/Time-2                16.5kB ± 0%    16.5kB ± 0%    ~     (all equal)

name                           old allocs/op  new allocs/op  delta
LTSVLog-2                          0.00           0.00         ~     (all equal)
StandardLog-2                      1.00 ± 0%      1.00 ± 0%    ~     (all equal)
ZapJSONProductionLog-2             1.00 ± 0%      1.00 ± 0%    ~     (all equal)
ZapJSONDevelopmentLog-2            7.00 ± 0%      7.00 ± 0%    ~     (all equal)
ZapLTSVProductionLog-2             1.00 ± 0%      1.00 ± 0%    ~     (all equal)
ZapLTSVDevelopmentLog-2            3.00 ± 0%      3.00 ± 0%    ~     (all equal)
ZerologTimestamp-2                 0.00           0.00         ~     (all equal)
ZerologTimeSecondsFromEpoch-2      0.00           0.00         ~     (all equal)
ZerologRFC3339Time-2               0.00           0.00         ~     (all equal)
ZerologRFC3339NanoTime-2           0.00           0.00         ~     (all equal)
LTSVLogErr/StackAndTime-2          4.00 ± 0%      4.00 ± 0%    ~     (all equal)
LTSVLogErr/Stack-2                 4.00 ± 0%      4.00 ± 0%    ~     (all equal)
LTSVLogErr/Time-2                  4.00 ± 0%      4.00 ± 0%    ~     (all equal)
```

## License
MIT
