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

Logs are written to `/var/log/coredns/filter.log`.   

## blacklist.txt
Websites are parsed using a regular expression: `(\w+)\.\w+$`
This means that for example adding `9gag.com` will also block `9gag.eu`, `9gag.net`, ...
