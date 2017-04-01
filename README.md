# curtain
***
                   .-"""-.
                  / .===. \
                  \/ 6 6 \/
                  ( \___/ )
      ________ooo__\_____/_____________
     /                                 \
    | simple http web that file upload! |
     \______________________ooo________/
                   |  |  |
                   |_ | _|
                   |  |  |
                   |__|__|
                   /-'Y'-\
                  (__/ \__)

# 一、build 
    0、 Go get github.com/akavel/rsrc
    1、 把nac.manifest 文件拷贝到当前windows项目根目录
    2、 rsrc -manifest bin/nac.manifest -o bin/nac.syso
    3、 gb build

# 二、running

    .\bin\curtain.exe
  Or   
  
    .\bin\curtain.exe -port 8080 -host 127.0.0.1