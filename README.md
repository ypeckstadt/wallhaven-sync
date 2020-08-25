# Wallhaven sync

Command line application to get your Wallhaven.cc collections and sync your favorite pictures to a folder of your choosing.

## Get collections information
The collection id is necessary for syncing.

```bash
wallhaven-sync.exe list -k <api key>
```

```bash
2020/08/26 02:01:00 Found collections (label - id):
2020/08/26 02:01:00 Default - 5456
2020/08/26 02:01:00 test collection - 7403
2020/08/26 02:01:00 SFW - 7114
```

## Syncing

``` bash
wallhaven-sync.exe sync -k <api key> -o <output folder> -u <username> -c <collection id>
```

``` bash
2020/08/26 01:52:40 Syncing page 1 ...
2020/08/26 01:52:41 Saving new file 0wg61x.png
2020/08/26 01:52:42 Saving new file lq6rwr.png
2020/08/26 01:52:43 Saving new file 2e2exx.jpg
2020/08/26 01:52:46 Saving new file r2e391.png
2020/08/26 01:53:02 Syncing page 2 ...
2020/08/26 01:53:03 Saving new file 43vgyn.jpg
2020/08/26 01:53:03 Saving new file dgeqoj.jpg
2020/08/26 01:53:05 Saving new file 47z1vn.jpg
2020/08/26 01:53:05 Saving new file 13pv13.jpg
2020/08/26 01:53:05 Saving new file ox19m9.jpg
2020/08/26 01:53:05 Saving new file 438w60.jpg
2020/08/26 01:53:07 Saving new file r2g7rm.jpg
2020/08/26 01:53:07 Syncing page 3 ...
2020/08/26 01:53:14 ========================================================
2020/08/26 01:53:14 ========================================================
2020/08/26 01:53:14 242 existing wallpapers have been skipped
2020/08/26 01:53:14 31 new wallpapers have been added
2020/08/26 01:53:14 0 wallpapers have been deleted
```
