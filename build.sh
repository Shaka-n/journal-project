# Reset previous build
rm -rf target
mkdir target

cd journal-project-backend
go build .
mv journal-project-backend ../target/journal-project-backend

cd ../journal-project-frontend
npm install
npm run build
mv build ../target/