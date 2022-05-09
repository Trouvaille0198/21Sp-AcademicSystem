kill -9 $(pgrep academic-system)
cd ~/Repo/21Sp-AcademicSystem
git restore *
git pull
cd ~/Repo/21Sp-AcademicSystem/backend
rm -r nohup.out
rm -r ./academic-system
# 也许是non-login模式下, ssh登录识别不了环境变量, 所以只能写全路径
/opt/go/bin/go mod tidy
/opt/gocode/bin/swag init
/opt/go/bin/go build
nohup ./academic-system >./nohup.out 2>&1 &
