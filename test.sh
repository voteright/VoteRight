go test ./database/ -coverprofile fmt > coverageReport.txt
go test ./config/ -coverprofile fmt >> coverageReport.txt
go test ./cmd/ -coverprofile fmt >> coverageReport.txt  
go test ./election/ -coverprofile fmt >> coverageReport.txt
go test ./models/ -coverprofile fmt >> coverageReport.txt
go test ./api/ -coverprofile fmt >> coverageReport.txt
cat coverageReport.txt
# go test ./verificationapi/