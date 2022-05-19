kill -9 $(pgrep vue-cli-service)
cd ~/Repo/21Sp-AcademicSystem
git restore *
git pull
cd ~/Repo/21Sp-AcademicSystem/frontend
rm -r nohup.out

npm i
nohup npm run serve >./nohup.out 2>&1 &
