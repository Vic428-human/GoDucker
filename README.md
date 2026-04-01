

## 觀念複習
- 例如透過 docker build -t my-go-docker . 建立一個新的映像，名字叫 my-go-docker，它不會自動啟動容器，也不會在 docker ps 裡多出一個新的 container
- container 是你用 docker run 從映像啟動出來的實例
- 所以未來，看到的docker desktop 的 containers 裡面有不同 container 紀錄，表示這是以前執行 docker run 時建立出來的實例

### docker build -t my-go-docker .
> 會依照 Dockerfile 的步驟，把你的 Go 專案編譯成一個映像，並存成 my-go-docker，讓你之後可以直接用 docker run 啟動它。

### 下一步：執行容器，如果你再輸入 docker run -p 8080:8080 my-go-docker
> 並把容器的 8080 port 映射到主機的 8080 port
```
// 這時候你打開瀏覽器訪問 http://localhost:8080，就能看到你的 Go 程式跑起來了。
docker run -p 8080:8080 my-go-docker
```

### 啟動 Docker Desktop
```
# Windows PowerShell
Start-Process "C:\Program Files\Docker\Docker\Docker Desktop.exe"

# 在 macOS
open -a Docker
```