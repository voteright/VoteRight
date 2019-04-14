go test ./database/ -coverprofile > coverageReport.txt
go test ./config/ -coverprofile >> coverageReport.txt
go test ./cmd/ -coverprofile >> coverageReport.txt  
go test ./election/ -coverprofile >> coverageReport.txt
go test ./models/ -coverprofile >> coverageReport.txt
go test ./api/ -coverprofile >> coverageReport.txt
cat coverageReport.txt
# go test ./verificationapi/