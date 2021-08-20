cd ./executable/observer
# npm install
# npm fund
npm run build
cd ..
lis --make build
docker compose build
docker compose up
docker compose down
cd ../../