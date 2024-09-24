# Win-Sleep

__A simple utility command to put Windows PC to sleep or hybernate__

## Command Usage

```
win-sleep.exe (-s|-h) [-d]
  -s  Sleep
  -h  Hybernate
  -d  Disable wake events
```

This command just calls [SetSuspendState()](https://learn.microsoft.com/en-us/windows/win32/api/powrprof/nf-powrprof-setsuspendstate) in `PowrProf.dll`


`rundll32.exe powrprof.dll,SetSuspendState 0,1,0` does not call the procedure like `SetSupendState(0,1,0)` but `SetSuspendState("0,1,0",null,null)`, so we need this utility.
