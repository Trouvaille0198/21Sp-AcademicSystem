cd ~/Repo/21Sp-AcademicSystem || exit
git restore *
git pull
cd ~/Repo/21Sp-AcademicSystem/backend || exit
swag init
go build
./academic-system
