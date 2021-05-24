# filter
Playing around with CoreDNS, writing my own plugin to filter a list of sites.

## Build
Clone the official `CoreDNS` repository.   
Add `filter:github.com/voje/filter` to `plugins.cfg`.   
Fetch the correct version:
```bash
go get github.com/voje/filter@<commit-hash>
```
Build `coredns` with `make`.   
