kill -9 $(pgrep academic-system)
cd ~/Repo/21Sp-AcademicSystem
git restore *
git pull
cd ~/Repo/21Sp-AcademicSystem/backend
# rm -r nohup.out
go mod tidy
swag init
go build
./academic-system
