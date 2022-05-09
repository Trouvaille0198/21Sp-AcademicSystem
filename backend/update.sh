kill -9 $(pgrep academic-system)
cd ~/Repo/21Sp-AcademicSystem
git restore *
git pull
cd ~/Repo/21Sp-AcademicSystem/backend
rm -r nohup.out
rm -r ./academic-system
go mod tidy
swag init
go build
nohup ./academic-system >./nohup.out 2>&1 &
