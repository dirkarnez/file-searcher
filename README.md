file-searcher
=============
No BS file searcher: `.\file-searcher.exe --fileName=dia-0.97.2-uninstall.exe --rootDirectory="C:\Program Files (x86)"`

### TODOs
- [x] regex pattern (regex pattern *IS NOT** glob pattern)
  - `file-searcher.exe --rootDirectory=%cd%\cmake-as-scripting --fileName=".*\.cmake"`
- [ ] fix memory issue caused by searching in large `rootDirectory` (e.g. search from the C Drive)
- [ ] super light indexing
